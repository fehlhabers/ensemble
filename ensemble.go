package main

import (
	"errors"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
)

type Ensemble struct {
	ensembleBranch string
	repo           *git.Repository
}

var (
	branchFormat = "ensemble/%s"
)

func NewEnsemble(path string) (*Ensemble, error) {
	repository, err := git.PlainOpen(path)
	if err != nil {
		return nil, fmt.Errorf("not a git repostory")
	}

	return &Ensemble{
			repo: repository,
		},
		nil
}

func (e *Ensemble) setWantedEnsembleBranch(branch string) {
	e.ensembleBranch = fmt.Sprintf(branchFormat, branch)
}

func (e *Ensemble) Start(branch string) error {
	w, _ := e.repo.Worktree()
	_ = w.Pull(&git.PullOptions{RemoteName: "origin"})
	e.setWantedEnsembleBranch(branch)
	ensembleBranch, err := e.getLocalEnsembleBranch()
	if err != nil {
		ensembleBranch, err = e.newEnsembleBranch()
		if err != nil {
			return err
		}
	}
	fmt.Printf("git checkout %s", e.ensembleBranch)
	err = w.Checkout(&git.CheckoutOptions{
		Branch: *ensembleBranch,
	})
	if err != nil {
		return err
	}
	return nil
}

func (e *Ensemble) newEnsembleBranch() (*plumbing.ReferenceName, error) {
	headRef, err := e.repo.Head()
	if err != nil {
		return nil, err
	}
	ref := plumbing.NewHashReference(e.wantedEnsembleBranchRef(), headRef.Hash())

	if err := e.repo.Storer.SetReference(ref); err != nil {
		return nil, err
	}

	e.repo.CreateBranch(&config.Branch{
		Name:   e.ensembleBranch,
		Merge:  e.wantedEnsembleBranchRef(),
		Remote: "origin",
	})
	if err := e.repo.Push(&git.PushOptions{}); err != nil {
		return nil, err
	}
	fmt.Printf("Created new ensemble <%s>", e.ensembleBranch)
	refName := ref.Name()
	return &refName, nil
}

func (e *Ensemble) Next() error {
	worktree, err := e.repo.Worktree()
	if err != nil {
		return err
	}

	if err := worktree.AddWithOptions(&git.AddOptions{All: true}); err != nil {
		return err
	}
	status, err := worktree.Status()
	if err != nil {
		return err
	}
	fmt.Print(status.String())
	fmt.Println("Commiting changes...")
	worktree.Commit("Ensemble WiP", &git.CommitOptions{})
	fmt.Println("Pushing changes...")
	e.repo.Push(&git.PushOptions{})
	return nil
}

func (e *Ensemble) getLocalEnsembleBranch() (*plumbing.ReferenceName, error) {
	var (
		ensembleBranch plumbing.ReferenceName
		branchFound    bool
	)

	branches, err := e.repo.Branches()
	if err != nil {
		return nil, err
	}
	fmt.Println("Looking for " + e.wantedEnsembleBranchRef().String())
	_ = branches.ForEach(func(ref *plumbing.Reference) error {
		fmt.Println("Found branch " + ref.Name().String())
		if ref.Name() == e.wantedEnsembleBranchRef() {
			ensembleBranch = ref.Name()
			branchFound = true
		}
		return nil
	})
	if branchFound {
		return &ensembleBranch, nil
	} else {
		return nil, errors.New("no ensemble branch found")
	}
}

func (e *Ensemble) wantedEnsembleBranchRef() plumbing.ReferenceName {
	return plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", e.ensembleBranch))
}

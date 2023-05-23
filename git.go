package main

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
)

type GitFacade interface {
	Branches() ([]string, error)
	Commit(message string) error
	Add() error
	Push() error
	Pull() error
	Fetch() error
	CheckoutRemoteTracked(branch string) error
	Checkout(branch string) error
}

type EnsembleGitFacade struct {
	repo     *git.Repository
	workTree *git.Worktree
}

var _ GitFacade = &EnsembleGitFacade{}

func NewEnsembleGitFacade(path string) (*EnsembleGitFacade, error) {
	repo, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}
	w, err := repo.Worktree()
	if err != nil {
		return nil, err
	}

	return &EnsembleGitFacade{
		repo:     repo,
		workTree: w,
	}, nil
}

func (e *EnsembleGitFacade) Add() error {
	fmt.Println("git add -A")
	return e.workTree.AddWithOptions(&git.AddOptions{
		All: true,
	})
}

func (e *EnsembleGitFacade) Branches() ([]string, error) {
	var (
		branches = make([]string, 0)
	)

	branchIt, err := e.repo.Branches()
	if err != nil {
		return nil, err
	}

	_ = branchIt.ForEach(func(r *plumbing.Reference) error {
		branches = append(branches, string(r.Name()))
		return nil
	})

	return branches, nil
}

func (e *EnsembleGitFacade) Checkout(branch string) error {
	return e.workTree.Checkout(&git.CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName(branch),
		Create: true,
	})
}

func (e *EnsembleGitFacade) Commit(message string) error {
	fmt.Printf("git commit -m \"%s\"\n", message)
	_, err := e.workTree.Commit(message, &git.CommitOptions{})
	return err
}

func (e *EnsembleGitFacade) Fetch() error {
	fmt.Println("git fetch")
	err := e.repo.Fetch(&git.FetchOptions{})
	if err == git.NoErrAlreadyUpToDate {
		fmt.Println(err)
	} else if err != nil {
		return err
	}
	return nil
}

func (e *EnsembleGitFacade) Pull() error {
	fmt.Println("Pulling latest...")
	err := e.workTree.Pull(&git.PullOptions{})
	if err == git.NoErrAlreadyUpToDate {
		fmt.Println(err)
	} else if err != nil {
		return err
	}
	return nil
}

func (e *EnsembleGitFacade) Push() error {
	fmt.Println("git push -u origin")
	return e.repo.Push(&git.PushOptions{
		RemoteName: "origin",
	})
}

func (e *EnsembleGitFacade) CheckoutRemoteTracked(branch string) error {
	branchRef := plumbing.NewBranchReferenceName(branch)

	opts := &config.Branch{
		Name:   branch,
		Remote: "origin",
		Merge:  branchRef,
	}

	if err := e.repo.CreateBranch(opts); err != nil {
		return err
	}
	// e.Checkout(branch)
	remoteRef := plumbing.NewRemoteReferenceName("origin", branch)
	remoteBranchRef := plumbing.NewSymbolicReference(branchRef, remoteRef)
	headRef, err := e.repo.Head()
	if err != nil {
		return err
	}
	ref := plumbing.NewHashReference(branchRef, headRef.Hash())
	if err := e.repo.Storer.SetReference(ref); err != nil {
		return err
	}
	if err := e.repo.Storer.SetReference(remoteBranchRef); err != nil {
		return err
	}
	return nil
}

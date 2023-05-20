package main

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type GitFacade interface {
	Branches() ([]string, error)
	Commit(message string) error
	Add() error
	Push()
	Pull() error
	Fetch()
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

// Commit implements GitFacade
func (e *EnsembleGitFacade) Commit(message string) error {
	_, err := e.workTree.Commit(message, &git.CommitOptions{})
	return err
}

// Fetch implements GitFacade
func (e *EnsembleGitFacade) Fetch() {
	panic("unimplemented")
}

// Pull implements GitFacade
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

// Push implements GitFacade
func (e *EnsembleGitFacade) Push() {
	panic("unimplemented")
}

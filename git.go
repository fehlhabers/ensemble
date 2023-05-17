package main

import (
	"fmt"

	"github.com/go-git/go-git/v5"
)

type GitFacade interface {
	Branches()
	Commit()
	Add()
	Push()
	Pull() error
	Fetch()
	Checkout()
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

// Add implements GitFacade
func (e *EnsembleGitFacade) Add() {
	panic("unimplemented")
}

// Branches implements GitFacade
func (e *EnsembleGitFacade) Branches() {
	panic("unimplemented")
}

// Checkout implements GitFacade
func (e *EnsembleGitFacade) Checkout() {
	panic("unimplemented")
}

// Commit implements GitFacade
func (e *EnsembleGitFacade) Commit() {
	panic("unimplemented")
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

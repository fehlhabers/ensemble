package main

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type Ensemble struct {
	repo *git.Repository
}

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

func (e Ensemble) Start(branch string) error {
	headRef, err := e.repo.Head()
	if err != nil {
		return err
	}
	ref := plumbing.NewHashReference(getReferenceName(branch), headRef.Hash())

	if err := e.repo.Storer.SetReference(ref); err != nil {
		return err
	}

	e.repo.Push(&git.PushOptions{})

}

func getReferenceName(ref string) plumbing.ReferenceName {
	return plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", ref))
}

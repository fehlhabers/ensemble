package main

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
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
	refName := plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", branch))

	return e.repo.CreateBranch(&config.Branch{
		Name:        branch,
		Remote:      branch,
		Merge:       refName,
		Rebase:      "true",
		Description: "",
	})

}

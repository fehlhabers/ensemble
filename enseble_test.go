package main

import (
	"fmt"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStartNoRepo(t *testing.T) {
	_, err := NewEnsemble("../")
	assert.Error(t, err)
}

func TestStart(t *testing.T) {
	repo, err := NewEnsemble(".")
	assert.NoError(t, err)
	err = repo.Start("test")
	assert.NoError(t, err)

	branches, err := repo.repo.Branches()
	if err != nil {
		return
	}

	_ = branches.ForEach(func(reference *plumbing.Reference) error {
		fmt.Println(reference)
		return nil
	})
}

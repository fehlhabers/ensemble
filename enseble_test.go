package main

import (
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStartNoRepo(t *testing.T) {
	_, err := NewEnsemble("../")
	assert.Error(t, err)
}

func TestStart(t *testing.T) {
	var (
		branchName string = "test"
	)
	e, err := NewEnsemble(".")
	assert.NoError(t, err)

	err = e.Start(branchName)
	assert.NoError(t, err)

	branches, _ := e.repo.Branches()

	foundBranch := false
	_ = branches.ForEach(func(reference *plumbing.Reference) error {
		if getReferenceName(branchName) == reference.Name() {
			err = e.repo.Storer.RemoveReference(reference.Name())
			foundBranch = true
		}
		return nil
	})
	assert.True(t, foundBranch)
}

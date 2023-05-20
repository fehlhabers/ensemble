package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPull(t *testing.T) {
	testee, err := NewEnsembleGitFacade(".")
	require.NoError(t, err)
	err = testee.Pull()
	assert.NoError(t, err)
}

func TestBranches(t *testing.T) {
	testee, err := NewEnsembleGitFacade(".")
	require.NoError(t, err)
	branches, err := testee.Branches()
	require.NoError(t, err)
	for _, br := range branches {
		fmt.Println(br)
	}
}

func TestCheckout(t *testing.T) {
	const (
		intBranch string = "int-test"
	)

	testee, err := NewEnsembleGitFacade(".")
	require.NoError(t, err)
	err = testee.Checkout(intBranch)
	require.NoError(t, err)
	br, err := testee.Branches()
	require.NoError(t, err)
	assert.Contains(t, br, intBranch)
}

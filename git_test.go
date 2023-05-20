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
		intRef    string = "refs/heads/int-test"
	)

	testee, err := NewEnsembleGitFacade(".")
	require.NoError(t, err)
	err = testee.Checkout(intBranch)
	require.NoError(t, err)
	br, err := testee.Branches()
	require.NoError(t, err)
	assert.Contains(t, br, intRef)
}

func TestAdd(t *testing.T) {
	testee, err := NewEnsembleGitFacade(".")
	require.NoError(t, err)
	err = testee.Add()
	require.NoError(t, err)
}

func TestCommit(t *testing.T) {
	testee, err := NewEnsembleGitFacade(".")
	require.NoError(t, err)
	err = testee.Commit("int test message")
	require.NoError(t, err)
}

func TestFetch(t *testing.T) {
	testee, err := NewEnsembleGitFacade(".")
	require.NoError(t, err)
	err = testee.Fetch()
	require.NoError(t, err)
}

func TestPush(t *testing.T) {
	testee, err := NewEnsembleGitFacade(".")
	require.NoError(t, err)
	testee.Push()
}

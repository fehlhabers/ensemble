package main

import (
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

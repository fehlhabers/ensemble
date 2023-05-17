package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStart(t *testing.T) {
	t.Run("pulls, checks out new branch, and pushes branch", func(t *testing.T) {
		ge := &TestGitEnsemble{}
		ensemble := &Ensemble{
			git: ge,
		}
		ensemble.Start()
		assert.Contains(t, ge.commandsGiven, "pull")
	})
}

type TestGitEnsemble struct {
	commandsGiven []string
}

// Add implements GitFacade
func (ge *TestGitEnsemble) Add() {
	panic("unimplemented")
}

// Branches implements GitFacade
func (ge *TestGitEnsemble) Branches() {
	panic("unimplemented")
}

// Checkout implements GitFacade
func (ge *TestGitEnsemble) Checkout() {
	panic("unimplemented")
}

// Commit implements GitFacade
func (ge *TestGitEnsemble) Commit() {
	panic("unimplemented")
}

// Fetch implements GitFacade
func (ge *TestGitEnsemble) Fetch() {
	panic("unimplemented")
}

// Pull implements GitFacade
func (ge *TestGitEnsemble) Pull() error {
	ge.commandsGiven = append(ge.commandsGiven, "pull")
	return nil
}

// Push implements GitFacade
func (ge *TestGitEnsemble) Push() {
	panic("unimplemented")
}

var _ GitFacade = &TestGitEnsemble{}

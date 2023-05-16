package main

type GitFacade interface {
	Branches()
	Commit()
	Add()
	Push()
	Pull()
	Fetch()
	Checkout()
}

type EnsembleGitFacade struct {
}

var _ GitFacade = &EnsembleGitFacade{}

func NewEnsembleGitFacade() *EnsembleGitFacade {
	return &EnsembleGitFacade{}
}

// Add implements GitFacade
func (*EnsembleGitFacade) Add() {
	panic("unimplemented")
}

// Branches implements GitFacade
func (*EnsembleGitFacade) Branches() {
	panic("unimplemented")
}

// Checkout implements GitFacade
func (*EnsembleGitFacade) Checkout() {
	panic("unimplemented")
}

// Commit implements GitFacade
func (*EnsembleGitFacade) Commit() {
	panic("unimplemented")
}

// Fetch implements GitFacade
func (*EnsembleGitFacade) Fetch() {
	panic("unimplemented")
}

// Pull implements GitFacade
func (*EnsembleGitFacade) Pull() {
	panic("unimplemented")
}

// Push implements GitFacade
func (*EnsembleGitFacade) Push() {
	panic("unimplemented")
}

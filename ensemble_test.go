package main

import "testing"

func TestStart(t *testing.T) {
	ensemble := NewTestEnsemble()
	ensemble.Start()
}

func NewTestEnsemble() *Ensemble {
	return &Ensemble{}
}

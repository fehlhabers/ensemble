package main

type Ensembler interface {
	Start()
	Next()
	Done()
}

type Ensemble struct {
	git GitFacade
}

var _ Ensembler = &Ensemble{}

func NewEnsemble() *Ensemble {
	return &Ensemble{
		git: &EnsembleGitFacade{},
	}
}

// Done implements Ensembler
func (e *Ensemble) Done() {
	panic("unimplemented")
}

// Next implements Ensembler
func (e *Ensemble) Next() {
	panic("unimplemented")
}

// Start implements Ensembler
func (e *Ensemble) Start() {
	e.git.Pull()
}

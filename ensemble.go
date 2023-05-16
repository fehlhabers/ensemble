package main

type Ensembler interface {
	Start()
	Next()
	Done()
}

type Ensemble struct {
}

var _ Ensembler = &Ensemble{}

// Done implements Ensembler
func (*Ensemble) Done() {
	panic("unimplemented")
}

// Next implements Ensembler
func (*Ensemble) Next() {
	panic("unimplemented")
}

// Start implements Ensembler
func (*Ensemble) Start() {
	panic("unimplemented")
}

package main

type Operator interface {
	Apply(int, int) int
}

type Operation struct {
	Operator Operator
}

package main

import (
	"errors"
)

type Result struct {
	Value  string
	Status bool
}

func (r Result) String() string {
	return "Result: " + r.Value
}

// Random error function
func GenerateResult(t int) (Result, error) {
	if t > 0 {
		return Result{Value: "Good", Status: true}, nil
	} else {
		return Result{Value: "Bad", Status: false}, errors.New("Invalid Foo")
	}
}

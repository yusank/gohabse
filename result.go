package gohbase

import (
	"github.com/tsuna/gohbase/hrpc"
)

type Result interface {
	Scanner() hrpc.Scanner
	Result() *hrpc.Result
	Int64() int64
	Bool() bool
	Err() error
}

var _ Result = &result{}

type result struct {
	result  *hrpc.Result
	scanner hrpc.Scanner
	i64     int64 // for Increment
	b       bool  // for CheckAndPut
	err     error
}

func (r *result) Scanner() hrpc.Scanner {
	return r.scanner
}

func (r *result) Result() *hrpc.Result {
	return r.result
}

func (r *result) Int64() int64 {
	return r.i64
}

func (r *result) Bool() bool {
	return r.b
}

func (r *result) Err() error {
	return r.err
}

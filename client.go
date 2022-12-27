package gohbase

import (
	"context"

	"github.com/tsuna/gohbase/hrpc"
)

type Client interface {
	/*
	* set values
	 */
	Context(context.Context) Client
	Table(table string) Client
	Key(key string) Client
	Family(family string) Client
	Qualifier(qualifier string) Client
	Amount(amount int64) Client
	Range(startRow, stopRow string) Client
	Values(values map[string]map[string][]byte) Client
	ExpectedValue(expectedValue []byte) Client
	Options(opts ...func(hrpc.Call) error) Client
	/*
	*operations
	 */
	Scan() Result
	Get() Result
	Put() Result
	Delete() Result
	Append() Result
	Increment() Result
	CheckAndPut() Result
	Close()
}

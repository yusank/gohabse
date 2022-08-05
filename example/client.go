package main

import (
	"context"

	"github.com/yusank/gohbase"
)

func main() {
	cli := gohbase.NewClient(gohbase.Addr("127.0.0.1"))

	// put
	values := map[string]map[string][]byte{
		"cf": {
			"col": []byte("val"),
		},
	}

	result := cli.Context(context.Background()).Table("test").Key("rowKey").Values(values).Put()
	if result.Err() != nil {
		panic(result.Err())
	}
}

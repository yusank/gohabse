package main

import (
	"context"

	"github.com/yusank/gohbase"
	gohbaseadmin "github.com/yusank/gohbase/admin"
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

func admin() {
	cli := gohbaseadmin.NewClient(gohbaseadmin.Addr("127.0.0.1"))

	r := cli.Context(context.Background()).Table("test").CreateTable()
	if r.Err() != nil {
		panic(r.Err())
	}
}

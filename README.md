# gohbase

`gohbase` is a hbase client for golang. This package written based on `github.com/tsuna/gohbase` and defines some interface to operate hbase more **simple** and **convenience**. For example :

```go
result := dbClient.Context(context.Background()).
    Table(msg.TableName()).
    Key(msg.HbaseRowKey()).
    Values(msg.HbaseValues()).
    Put() // define Put, Get, Delete, Scan etc.
if result.Err() != nil {
    return result.Err()
}
// result has more methos
```

# admin

this package contains the admin interface for operate the hbase db, like create table, drop table, etc.

client:

```go
type Client interface {
    /*
     *  tables
     */
    Context(context.Context) Client
    Table(string) Client
    Families(map[string]map[string]string) Client
    SplitKeys([][]byte) Client
    /*
     * snapshot
     */
    SnapshotName(string) Client
    SnapshotSkipFlush() Client
    SnapshotVersion(int32) Client
    SnapshotOwner(string) Client
    /*
     * list
     */
    Regex(string) Client
    IncludeSysTables(bool) Client
    Namespace(string) Client
    /*
     * move region
     */
    TargetRegion(string) Client
    TargetServerName(string) Client
    /*
    * other operations
     */
    CreateTable() Result
    DeleteTable() Result
    EnableTable() Result
    DisableTable() Result
    CreateSnapshot() Result
    DeleteSnapshot() Result
    ListSnapshots() Result
    RestoreSnapshot() Result
    ClusterStatus() Result
    ListTableNames() Result
    // SetBalancer sets balancer state and returns previous state
    SetBalancer(enabled bool) Result
    // MoveRegion moves a region to a different RegionServer
    MoveRegion() Result
}
```

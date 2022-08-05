package admin

import "github.com/tsuna/gohbase/pb"

/*
 * Admin Result
 */

type Result interface {
	Snapshots() []*pb.SnapshotDescription
	TableNames() []*pb.TableName
	ClusterStatus() *pb.ClusterStatus
	Bool() bool
	Err() error
}

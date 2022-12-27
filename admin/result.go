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

type result struct {
	err           error
	snapshots     []*pb.SnapshotDescription
	tableNames    []*pb.TableName
	clusterStatus *pb.ClusterStatus
	b             bool
}

func (r *result) Snapshots() []*pb.SnapshotDescription {
	return r.snapshots
}

func (r *result) TableNames() []*pb.TableName {
	return r.tableNames
}

func (r *result) ClusterStatus() *pb.ClusterStatus {
	return r.clusterStatus
}

func (r *result) Bool() bool {
	return r.b
}

func (r *result) Err() error {
	return r.err
}

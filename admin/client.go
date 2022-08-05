package admin

import (
	"context"

	"github.com/tsuna/gohbase/hrpc"
	"github.com/tsuna/gohbase/pb"
)

type Client interface {
	// set values
	//table
	Context(context.Context) Client
	Table(string) Client
	Families(map[string]map[string]string) Client
	SplitKeys([][]byte)
	//snapshot
	SnapshotName(string) Client
	SnapshotType(*pb.SnapshotDescription_Type) Client
	SnapshotVersion(int32) Client
	SnapshotOwner(string) Client

	// operations

	CreateTable(t *hrpc.CreateTable) Result
	DeleteTable(t *hrpc.DeleteTable) Result
	EnableTable(t *hrpc.EnableTable) Result
	DisableTable(t *hrpc.DisableTable) Result
	CreateSnapshot(t *hrpc.Snapshot) Result
	DeleteSnapshot(t *hrpc.Snapshot) Result
	ListSnapshots(t *hrpc.ListSnapshots) Result
	RestoreSnapshot(t *hrpc.Snapshot) Result
	ClusterStatus() Result
	ListTableNames(t *hrpc.ListTableNames) Result
	// SetBalancer sets balancer state and returns previous state
	SetBalancer(sb *hrpc.SetBalancer) Result
	// MoveRegion moves a region to a different RegionServer
	MoveRegion(mr *hrpc.MoveRegion) Result
}

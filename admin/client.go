package admin

import (
	"github.com/tsuna/gohbase/hrpc"
)

type Client interface {
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

package admin

import (
	"context"

	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/hrpc"
	"go.uber.org/atomic"
)

type client struct {
	*session

	cli   gohbase.AdminClient
	clone *atomic.Bool
}

func (c *client) Context(ctx context.Context) Client {
	tx := c.getInstance()
	tx.WithContext(ctx)

	return tx
}

func (c *client) Table(s string) Client {
	tx := c.getInstance()
	tx.WithTable(s)

	return tx
}

func (c *client) Families(families map[string]map[string]string) Client {
	tx := c.getInstance()
	tx.WithFamilies(families)

	return tx
}

func (c *client) SplitKeys(sk [][]byte) Client {
	tx := c.getInstance()
	tx.WithSplitKeys(sk)

	return tx
}

func (c *client) SnapshotName(name string) Client {
	tx := c.getInstance()
	tx.WithSnapshotName(name)

	return tx
}

func (c *client) SnapshotSkipFlush() Client {
	tx := c.getInstance()
	tx.WithSnapshotSkipFlush()

	return tx
}

func (c *client) SnapshotVersion(v int32) Client {
	tx := c.getInstance()
	tx.WithSnapshotVersion(v)

	return tx
}

func (c *client) SnapshotOwner(owner string) Client {
	tx := c.getInstance()
	tx.WithSnapshotOwner(owner)

	return tx
}

func (c *client) Regex(s string) Client {
	tx := c.getInstance()
	tx.WithRegex(s)

	return tx
}

func (c *client) IncludeSysTables(b bool) Client {
	tx := c.getInstance()
	tx.WithIncludeSysTables(b)

	return tx
}

func (c *client) Namespace(ns string) Client {
	tx := c.getInstance()
	tx.WithNamespace(ns)

	return tx
}

func (c *client) TargetRegion(region string) Client {
	tx := c.getInstance()
	tx.WithTargetRegion(region)

	return tx
}

func (c *client) TargetServerName(name string) Client {
	tx := c.getInstance()
	tx.WithTargetServiceName(name)

	return tx
}

func (c *client) CreateTable() Result {
	tx := c.getInstance()
	if err := tx.session.validate(); err != nil {
		return &result{
			err: err,
		}
	}

	opts := make([]func(*hrpc.CreateTable), 0)
	if len(tx.splitKeys) > 0 {
		opts = append(opts, hrpc.SplitKeys(tx.splitKeys))
	}

	req := hrpc.NewCreateTable(tx.ctx, []byte(tx.table), tx.families, opts...)
	return &result{
		err: tx.cli.CreateTable(req),
	}
}

func (c *client) DeleteTable() Result {
	tx := c.getInstance()
	if err := tx.session.validate(); err != nil {
		return &result{
			err: err,
		}
	}

	req := hrpc.NewDeleteTable(tx.ctx, []byte(tx.table))
	return &result{
		err: tx.cli.DeleteTable(req),
	}
}

func (c *client) EnableTable() Result {
	tx := c.getInstance()
	if err := tx.session.validate(); err != nil {
		return &result{
			err: err,
		}
	}

	req := hrpc.NewEnableTable(tx.ctx, []byte(tx.table))
	return &result{
		err: tx.cli.EnableTable(req),
	}
}

func (c *client) DisableTable() Result {
	tx := c.getInstance()
	if err := tx.session.validate(); err != nil {
		return &result{
			err: err,
		}
	}

	req := hrpc.NewDisableTable(tx.ctx, []byte(tx.table))
	return &result{
		err: tx.cli.DisableTable(req),
	}
}

func (c *client) CreateSnapshot() Result {
	tx := c.getInstance()
	if err := tx.session.validate(); err != nil {
		return &result{
			err: err,
		}
	}

	req, err := hrpc.NewSnapshot(tx.ctx, *tx.snapshotName, tx.table, tx.opts...)
	if err != nil {
		return &result{
			err: err,
		}
	}

	return &result{
		err: tx.cli.CreateSnapshot(req),
	}
}

func (c *client) DeleteSnapshot() Result {
	tx := c.getInstance()
	if err := tx.session.validate(); err != nil {
		return &result{
			err: err,
		}
	}

	req, err := hrpc.NewSnapshot(tx.ctx, *tx.snapshotName, tx.table, tx.opts...)
	if err != nil {
		return &result{
			err: err,
		}
	}

	return &result{
		err: tx.cli.DeleteSnapshot(req),
	}
}

func (c *client) ListSnapshots() Result {
	tx := c.getInstance()
	if err := tx.session.validate(); err != nil {
		return &result{
			err: err,
		}
	}

	r := &result{}
	r.snapshots, r.err = tx.cli.ListSnapshots(hrpc.NewListSnapshots(tx.ctx))
	return r
}

func (c *client) RestoreSnapshot() Result {
	tx := c.getInstance()
	if err := tx.session.validate(); err != nil {
		return &result{
			err: err,
		}
	}

	req, err := hrpc.NewSnapshot(tx.ctx, *tx.snapshotName, tx.table, tx.opts...)
	if err != nil {
		return &result{
			err: err,
		}
	}

	return &result{
		err: tx.cli.RestoreSnapshot(req),
	}
}

func (c *client) ClusterStatus() Result {
	tx := c.getInstance()
	if err := tx.session.validate(); err != nil {
		return &result{
			err: err,
		}
	}

	r := &result{}
	r.clusterStatus, r.err = tx.cli.ClusterStatus()
	return r
}

func (c *client) ListTableNames() Result {
	tx := c.getInstance()
	if err := tx.session.validate(); err != nil {
		return &result{
			err: err,
		}
	}

	req, err := hrpc.NewListTableNames(tx.ctx, tx.opts...)
	if err != nil {
		return &result{
			err: err,
		}
	}

	r := &result{}
	r.tableNames, r.err = tx.cli.ListTableNames(req)
	return r
}

func (c *client) SetBalancer(enabled bool) Result {
	tx := c.getInstance()
	if err := tx.session.validate(); err != nil {
		return &result{
			err: err,
		}
	}

	req, err := hrpc.NewSetBalancer(tx.ctx, enabled)
	if err != nil {
		return &result{
			err: err,
		}
	}

	r := &result{}
	r.b, r.err = tx.cli.SetBalancer(req)
	return r
}

func (c *client) MoveRegion() Result {
	tx := c.getInstance()
	if err := tx.session.validate(); err != nil {
		return &result{
			err: err,
		}
	}

	req, err := hrpc.NewMoveRegion(tx.ctx, []byte(*tx.targetRegion), tx.opts...)
	if err != nil {
		return &result{
			err: err,
		}
	}

	return &result{
		err: tx.cli.MoveRegion(req),
	}
}

func (c *client) getInstance() *client {
	if c.clone.Load() {
		return c
	}

	return &client{
		cli:     c.cli,
		session: newSession(context.Background()),
		clone:   atomic.NewBool(true),
	}
}

func (c *client) Close() {
	cli, ok := c.cli.(gohbase.Client)
	if ok {
		cli.Close()
	}
}

func NewClient(opts ...Option) Client {
	o := newOption()
	o.apply(opts...)

	cli := gohbase.NewAdminClient(o.addr, o.gohbaseOpts...)
	return &client{
		cli:   cli,
		clone: atomic.NewBool(false),
	}
}

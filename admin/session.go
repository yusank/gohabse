package admin

import (
	"context"
	"fmt"

	"github.com/tsuna/gohbase/hrpc"
)

type session struct {
	ctx          context.Context
	table        string
	families     map[string]map[string]string
	splitKeys    [][]byte
	snapshotName *string
	targetRegion *string

	// options
	opts []func(hrpc.Call) error
}

func newSession(ctx context.Context) *session {
	return &session{
		ctx: ctx,
		// options
		opts: make([]func(hrpc.Call) error, 0),
	}
}

func (s *session) WithContext(ctx context.Context) *session {
	s.ctx = ctx
	return s
}

func (s *session) WithTable(table string) *session {
	s.table = table
	return s
}

func (s *session) WithFamilies(families map[string]map[string]string) *session {
	s.families = copyFamilies(families)
	return s
}

func copyFamilies(src map[string]map[string]string) map[string]map[string]string {
	target := make(map[string]map[string]string, len(src))
	for k, m := range src {
		sub := make(map[string]string, len(m))
		for s, v := range m {
			sub[s] = v
		}

		target[k] = m
	}

	return target
}

func (s *session) WithSplitKeys(sk [][]byte) *session {
	s.splitKeys = make([][]byte, 0, len(sk))
	copy(s.splitKeys, sk)

	return s
}

func (s *session) WithSnapshotName(name string) *session {
	s.snapshotName = &name

	return s
}

func (s *session) WithSnapshotSkipFlush() *session {
	s.opts = append(s.opts, hrpc.SnapshotSkipFlush())

	return s
}

func (s *session) WithSnapshotVersion(v int32) *session {
	s.opts = append(s.opts, hrpc.SnapshotVersion(v))

	return s
}

func (s *session) WithSnapshotOwner(owner string) *session {
	s.opts = append(s.opts, hrpc.SnapshotOwner(owner))

	return s
}

func (s *session) WithRegex(regex string) *session {
	s.opts = append(s.opts, hrpc.ListRegex(regex))

	return s
}

func (s *session) WithIncludeSysTables(b bool) *session {
	s.opts = append(s.opts, hrpc.ListSysTables(b))

	return s
}

func (s *session) WithNamespace(ns string) *session {
	s.opts = append(s.opts, hrpc.ListNamespace(ns))

	return s
}

func (s *session) WithTargetRegion(region string) *session {
	s.targetRegion = &region

	return s
}

func (s *session) WithTargetServiceName(name string) *session {
	s.opts = append(s.opts, hrpc.WithDestinationRegionServer(name))
	return s
}

func (s *session) WithOptions(opts ...func(hrpc.Call) error) *session {
	s.opts = opts
	return s
}

var (
	ErrNilContext = fmt.Errorf("nil context")
	ErrNilTable   = fmt.Errorf("nil table")
)

// TODO: more validation for different operations.
func (s *session) validate() error {
	if s.ctx == nil {
		return ErrNilContext
	}
	if s.table == "" {
		return ErrNilTable
	}

	return nil
}

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/zibbp/ganymede/ent/playlist"
	"github.com/zibbp/ganymede/ent/predicate"
)

// PlaylistDelete is the builder for deleting a Playlist entity.
type PlaylistDelete struct {
	config
	hooks    []Hook
	mutation *PlaylistMutation
}

// Where appends a list predicates to the PlaylistDelete builder.
func (pd *PlaylistDelete) Where(ps ...predicate.Playlist) *PlaylistDelete {
	pd.mutation.Where(ps...)
	return pd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pd *PlaylistDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, PlaylistMutation](ctx, pd.sqlExec, pd.mutation, pd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pd *PlaylistDelete) ExecX(ctx context.Context) int {
	n, err := pd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pd *PlaylistDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: playlist.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: playlist.FieldID,
			},
		},
	}
	if ps := pd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pd.mutation.done = true
	return affected, err
}

// PlaylistDeleteOne is the builder for deleting a single Playlist entity.
type PlaylistDeleteOne struct {
	pd *PlaylistDelete
}

// Exec executes the deletion query.
func (pdo *PlaylistDeleteOne) Exec(ctx context.Context) error {
	n, err := pdo.pd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{playlist.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pdo *PlaylistDeleteOne) ExecX(ctx context.Context) {
	pdo.pd.ExecX(ctx)
}

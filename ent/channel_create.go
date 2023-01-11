// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/zibbp/ganymede/ent/channel"
	"github.com/zibbp/ganymede/ent/live"
	"github.com/zibbp/ganymede/ent/vod"
)

// ChannelCreate is the builder for creating a Channel entity.
type ChannelCreate struct {
	config
	mutation *ChannelMutation
	hooks    []Hook
}

// SetExtID sets the "ext_id" field.
func (cc *ChannelCreate) SetExtID(s string) *ChannelCreate {
	cc.mutation.SetExtID(s)
	return cc
}

// SetNillableExtID sets the "ext_id" field if the given value is not nil.
func (cc *ChannelCreate) SetNillableExtID(s *string) *ChannelCreate {
	if s != nil {
		cc.SetExtID(*s)
	}
	return cc
}

// SetName sets the "name" field.
func (cc *ChannelCreate) SetName(s string) *ChannelCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetDisplayName sets the "display_name" field.
func (cc *ChannelCreate) SetDisplayName(s string) *ChannelCreate {
	cc.mutation.SetDisplayName(s)
	return cc
}

// SetImagePath sets the "image_path" field.
func (cc *ChannelCreate) SetImagePath(s string) *ChannelCreate {
	cc.mutation.SetImagePath(s)
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *ChannelCreate) SetUpdatedAt(t time.Time) *ChannelCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *ChannelCreate) SetNillableUpdatedAt(t *time.Time) *ChannelCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetCreatedAt sets the "created_at" field.
func (cc *ChannelCreate) SetCreatedAt(t time.Time) *ChannelCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *ChannelCreate) SetNillableCreatedAt(t *time.Time) *ChannelCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *ChannelCreate) SetID(u uuid.UUID) *ChannelCreate {
	cc.mutation.SetID(u)
	return cc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cc *ChannelCreate) SetNillableID(u *uuid.UUID) *ChannelCreate {
	if u != nil {
		cc.SetID(*u)
	}
	return cc
}

// AddVodIDs adds the "vods" edge to the Vod entity by IDs.
func (cc *ChannelCreate) AddVodIDs(ids ...uuid.UUID) *ChannelCreate {
	cc.mutation.AddVodIDs(ids...)
	return cc
}

// AddVods adds the "vods" edges to the Vod entity.
func (cc *ChannelCreate) AddVods(v ...*Vod) *ChannelCreate {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return cc.AddVodIDs(ids...)
}

// AddLiveIDs adds the "live" edge to the Live entity by IDs.
func (cc *ChannelCreate) AddLiveIDs(ids ...uuid.UUID) *ChannelCreate {
	cc.mutation.AddLiveIDs(ids...)
	return cc
}

// AddLive adds the "live" edges to the Live entity.
func (cc *ChannelCreate) AddLive(l ...*Live) *ChannelCreate {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return cc.AddLiveIDs(ids...)
}

// Mutation returns the ChannelMutation object of the builder.
func (cc *ChannelCreate) Mutation() *ChannelMutation {
	return cc.mutation
}

// Save creates the Channel in the database.
func (cc *ChannelCreate) Save(ctx context.Context) (*Channel, error) {
	cc.defaults()
	return withHooks[*Channel, ChannelMutation](ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ChannelCreate) SaveX(ctx context.Context) *Channel {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ChannelCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ChannelCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *ChannelCreate) defaults() {
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := channel.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := channel.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.ID(); !ok {
		v := channel.DefaultID()
		cc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ChannelCreate) check() error {
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Channel.name"`)}
	}
	if _, ok := cc.mutation.DisplayName(); !ok {
		return &ValidationError{Name: "display_name", err: errors.New(`ent: missing required field "Channel.display_name"`)}
	}
	if _, ok := cc.mutation.ImagePath(); !ok {
		return &ValidationError{Name: "image_path", err: errors.New(`ent: missing required field "Channel.image_path"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Channel.updated_at"`)}
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Channel.created_at"`)}
	}
	return nil
}

func (cc *ChannelCreate) sqlSave(ctx context.Context) (*Channel, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *ChannelCreate) createSpec() (*Channel, *sqlgraph.CreateSpec) {
	var (
		_node = &Channel{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: channel.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: channel.FieldID,
			},
		}
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cc.mutation.ExtID(); ok {
		_spec.SetField(channel.FieldExtID, field.TypeString, value)
		_node.ExtID = value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(channel.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.DisplayName(); ok {
		_spec.SetField(channel.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := cc.mutation.ImagePath(); ok {
		_spec.SetField(channel.FieldImagePath, field.TypeString, value)
		_node.ImagePath = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(channel.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(channel.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := cc.mutation.VodsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.VodsTable,
			Columns: []string{channel.VodsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: vod.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.LiveIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.LiveTable,
			Columns: []string{channel.LiveColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: live.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ChannelCreateBulk is the builder for creating many Channel entities in bulk.
type ChannelCreateBulk struct {
	config
	builders []*ChannelCreate
}

// Save creates the Channel entities in the database.
func (ccb *ChannelCreateBulk) Save(ctx context.Context) ([]*Channel, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Channel, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ChannelMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ChannelCreateBulk) SaveX(ctx context.Context) []*Channel {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ChannelCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ChannelCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}

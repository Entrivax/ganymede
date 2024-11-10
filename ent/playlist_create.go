// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/zibbp/ganymede/ent/multistreaminfo"
	"github.com/zibbp/ganymede/ent/playlist"
	"github.com/zibbp/ganymede/ent/vod"
)

// PlaylistCreate is the builder for creating a Playlist entity.
type PlaylistCreate struct {
	config
	mutation *PlaylistMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (pc *PlaylistCreate) SetName(s string) *PlaylistCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetDescription sets the "description" field.
func (pc *PlaylistCreate) SetDescription(s string) *PlaylistCreate {
	pc.mutation.SetDescription(s)
	return pc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pc *PlaylistCreate) SetNillableDescription(s *string) *PlaylistCreate {
	if s != nil {
		pc.SetDescription(*s)
	}
	return pc
}

// SetThumbnailPath sets the "thumbnail_path" field.
func (pc *PlaylistCreate) SetThumbnailPath(s string) *PlaylistCreate {
	pc.mutation.SetThumbnailPath(s)
	return pc
}

// SetNillableThumbnailPath sets the "thumbnail_path" field if the given value is not nil.
func (pc *PlaylistCreate) SetNillableThumbnailPath(s *string) *PlaylistCreate {
	if s != nil {
		pc.SetThumbnailPath(*s)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *PlaylistCreate) SetUpdatedAt(t time.Time) *PlaylistCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *PlaylistCreate) SetNillableUpdatedAt(t *time.Time) *PlaylistCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetCreatedAt sets the "created_at" field.
func (pc *PlaylistCreate) SetCreatedAt(t time.Time) *PlaylistCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *PlaylistCreate) SetNillableCreatedAt(t *time.Time) *PlaylistCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *PlaylistCreate) SetID(u uuid.UUID) *PlaylistCreate {
	pc.mutation.SetID(u)
	return pc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pc *PlaylistCreate) SetNillableID(u *uuid.UUID) *PlaylistCreate {
	if u != nil {
		pc.SetID(*u)
	}
	return pc
}

// AddVodIDs adds the "vods" edge to the Vod entity by IDs.
func (pc *PlaylistCreate) AddVodIDs(ids ...uuid.UUID) *PlaylistCreate {
	pc.mutation.AddVodIDs(ids...)
	return pc
}

// AddVods adds the "vods" edges to the Vod entity.
func (pc *PlaylistCreate) AddVods(v ...*Vod) *PlaylistCreate {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return pc.AddVodIDs(ids...)
}

// AddMultistreamInfoIDs adds the "multistream_info" edge to the MultistreamInfo entity by IDs.
func (pc *PlaylistCreate) AddMultistreamInfoIDs(ids ...int) *PlaylistCreate {
	pc.mutation.AddMultistreamInfoIDs(ids...)
	return pc
}

// AddMultistreamInfo adds the "multistream_info" edges to the MultistreamInfo entity.
func (pc *PlaylistCreate) AddMultistreamInfo(m ...*MultistreamInfo) *PlaylistCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return pc.AddMultistreamInfoIDs(ids...)
}

// Mutation returns the PlaylistMutation object of the builder.
func (pc *PlaylistCreate) Mutation() *PlaylistMutation {
	return pc.mutation
}

// Save creates the Playlist in the database.
func (pc *PlaylistCreate) Save(ctx context.Context) (*Playlist, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PlaylistCreate) SaveX(ctx context.Context) *Playlist {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PlaylistCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PlaylistCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PlaylistCreate) defaults() {
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := playlist.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := playlist.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.ID(); !ok {
		v := playlist.DefaultID()
		pc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PlaylistCreate) check() error {
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Playlist.name"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Playlist.updated_at"`)}
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Playlist.created_at"`)}
	}
	return nil
}

func (pc *PlaylistCreate) sqlSave(ctx context.Context) (*Playlist, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
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
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PlaylistCreate) createSpec() (*Playlist, *sqlgraph.CreateSpec) {
	var (
		_node = &Playlist{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(playlist.Table, sqlgraph.NewFieldSpec(playlist.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = pc.conflict
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(playlist.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.SetField(playlist.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := pc.mutation.ThumbnailPath(); ok {
		_spec.SetField(playlist.FieldThumbnailPath, field.TypeString, value)
		_node.ThumbnailPath = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(playlist.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(playlist.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := pc.mutation.VodsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   playlist.VodsTable,
			Columns: playlist.VodsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vod.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.MultistreamInfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   playlist.MultistreamInfoTable,
			Columns: []string{playlist.MultistreamInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(multistreaminfo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Playlist.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PlaylistUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (pc *PlaylistCreate) OnConflict(opts ...sql.ConflictOption) *PlaylistUpsertOne {
	pc.conflict = opts
	return &PlaylistUpsertOne{
		create: pc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Playlist.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pc *PlaylistCreate) OnConflictColumns(columns ...string) *PlaylistUpsertOne {
	pc.conflict = append(pc.conflict, sql.ConflictColumns(columns...))
	return &PlaylistUpsertOne{
		create: pc,
	}
}

type (
	// PlaylistUpsertOne is the builder for "upsert"-ing
	//  one Playlist node.
	PlaylistUpsertOne struct {
		create *PlaylistCreate
	}

	// PlaylistUpsert is the "OnConflict" setter.
	PlaylistUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *PlaylistUpsert) SetName(v string) *PlaylistUpsert {
	u.Set(playlist.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PlaylistUpsert) UpdateName() *PlaylistUpsert {
	u.SetExcluded(playlist.FieldName)
	return u
}

// SetDescription sets the "description" field.
func (u *PlaylistUpsert) SetDescription(v string) *PlaylistUpsert {
	u.Set(playlist.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *PlaylistUpsert) UpdateDescription() *PlaylistUpsert {
	u.SetExcluded(playlist.FieldDescription)
	return u
}

// ClearDescription clears the value of the "description" field.
func (u *PlaylistUpsert) ClearDescription() *PlaylistUpsert {
	u.SetNull(playlist.FieldDescription)
	return u
}

// SetThumbnailPath sets the "thumbnail_path" field.
func (u *PlaylistUpsert) SetThumbnailPath(v string) *PlaylistUpsert {
	u.Set(playlist.FieldThumbnailPath, v)
	return u
}

// UpdateThumbnailPath sets the "thumbnail_path" field to the value that was provided on create.
func (u *PlaylistUpsert) UpdateThumbnailPath() *PlaylistUpsert {
	u.SetExcluded(playlist.FieldThumbnailPath)
	return u
}

// ClearThumbnailPath clears the value of the "thumbnail_path" field.
func (u *PlaylistUpsert) ClearThumbnailPath() *PlaylistUpsert {
	u.SetNull(playlist.FieldThumbnailPath)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *PlaylistUpsert) SetUpdatedAt(v time.Time) *PlaylistUpsert {
	u.Set(playlist.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *PlaylistUpsert) UpdateUpdatedAt() *PlaylistUpsert {
	u.SetExcluded(playlist.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Playlist.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(playlist.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *PlaylistUpsertOne) UpdateNewValues() *PlaylistUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(playlist.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(playlist.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Playlist.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *PlaylistUpsertOne) Ignore() *PlaylistUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PlaylistUpsertOne) DoNothing() *PlaylistUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PlaylistCreate.OnConflict
// documentation for more info.
func (u *PlaylistUpsertOne) Update(set func(*PlaylistUpsert)) *PlaylistUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PlaylistUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *PlaylistUpsertOne) SetName(v string) *PlaylistUpsertOne {
	return u.Update(func(s *PlaylistUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PlaylistUpsertOne) UpdateName() *PlaylistUpsertOne {
	return u.Update(func(s *PlaylistUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *PlaylistUpsertOne) SetDescription(v string) *PlaylistUpsertOne {
	return u.Update(func(s *PlaylistUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *PlaylistUpsertOne) UpdateDescription() *PlaylistUpsertOne {
	return u.Update(func(s *PlaylistUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *PlaylistUpsertOne) ClearDescription() *PlaylistUpsertOne {
	return u.Update(func(s *PlaylistUpsert) {
		s.ClearDescription()
	})
}

// SetThumbnailPath sets the "thumbnail_path" field.
func (u *PlaylistUpsertOne) SetThumbnailPath(v string) *PlaylistUpsertOne {
	return u.Update(func(s *PlaylistUpsert) {
		s.SetThumbnailPath(v)
	})
}

// UpdateThumbnailPath sets the "thumbnail_path" field to the value that was provided on create.
func (u *PlaylistUpsertOne) UpdateThumbnailPath() *PlaylistUpsertOne {
	return u.Update(func(s *PlaylistUpsert) {
		s.UpdateThumbnailPath()
	})
}

// ClearThumbnailPath clears the value of the "thumbnail_path" field.
func (u *PlaylistUpsertOne) ClearThumbnailPath() *PlaylistUpsertOne {
	return u.Update(func(s *PlaylistUpsert) {
		s.ClearThumbnailPath()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *PlaylistUpsertOne) SetUpdatedAt(v time.Time) *PlaylistUpsertOne {
	return u.Update(func(s *PlaylistUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *PlaylistUpsertOne) UpdateUpdatedAt() *PlaylistUpsertOne {
	return u.Update(func(s *PlaylistUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *PlaylistUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PlaylistCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PlaylistUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *PlaylistUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: PlaylistUpsertOne.ID is not supported by MySQL driver. Use PlaylistUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *PlaylistUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// PlaylistCreateBulk is the builder for creating many Playlist entities in bulk.
type PlaylistCreateBulk struct {
	config
	err      error
	builders []*PlaylistCreate
	conflict []sql.ConflictOption
}

// Save creates the Playlist entities in the database.
func (pcb *PlaylistCreateBulk) Save(ctx context.Context) ([]*Playlist, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Playlist, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PlaylistMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = pcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PlaylistCreateBulk) SaveX(ctx context.Context) []*Playlist {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PlaylistCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PlaylistCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Playlist.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PlaylistUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (pcb *PlaylistCreateBulk) OnConflict(opts ...sql.ConflictOption) *PlaylistUpsertBulk {
	pcb.conflict = opts
	return &PlaylistUpsertBulk{
		create: pcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Playlist.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pcb *PlaylistCreateBulk) OnConflictColumns(columns ...string) *PlaylistUpsertBulk {
	pcb.conflict = append(pcb.conflict, sql.ConflictColumns(columns...))
	return &PlaylistUpsertBulk{
		create: pcb,
	}
}

// PlaylistUpsertBulk is the builder for "upsert"-ing
// a bulk of Playlist nodes.
type PlaylistUpsertBulk struct {
	create *PlaylistCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Playlist.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(playlist.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *PlaylistUpsertBulk) UpdateNewValues() *PlaylistUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(playlist.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(playlist.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Playlist.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *PlaylistUpsertBulk) Ignore() *PlaylistUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PlaylistUpsertBulk) DoNothing() *PlaylistUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PlaylistCreateBulk.OnConflict
// documentation for more info.
func (u *PlaylistUpsertBulk) Update(set func(*PlaylistUpsert)) *PlaylistUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PlaylistUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *PlaylistUpsertBulk) SetName(v string) *PlaylistUpsertBulk {
	return u.Update(func(s *PlaylistUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PlaylistUpsertBulk) UpdateName() *PlaylistUpsertBulk {
	return u.Update(func(s *PlaylistUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *PlaylistUpsertBulk) SetDescription(v string) *PlaylistUpsertBulk {
	return u.Update(func(s *PlaylistUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *PlaylistUpsertBulk) UpdateDescription() *PlaylistUpsertBulk {
	return u.Update(func(s *PlaylistUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *PlaylistUpsertBulk) ClearDescription() *PlaylistUpsertBulk {
	return u.Update(func(s *PlaylistUpsert) {
		s.ClearDescription()
	})
}

// SetThumbnailPath sets the "thumbnail_path" field.
func (u *PlaylistUpsertBulk) SetThumbnailPath(v string) *PlaylistUpsertBulk {
	return u.Update(func(s *PlaylistUpsert) {
		s.SetThumbnailPath(v)
	})
}

// UpdateThumbnailPath sets the "thumbnail_path" field to the value that was provided on create.
func (u *PlaylistUpsertBulk) UpdateThumbnailPath() *PlaylistUpsertBulk {
	return u.Update(func(s *PlaylistUpsert) {
		s.UpdateThumbnailPath()
	})
}

// ClearThumbnailPath clears the value of the "thumbnail_path" field.
func (u *PlaylistUpsertBulk) ClearThumbnailPath() *PlaylistUpsertBulk {
	return u.Update(func(s *PlaylistUpsert) {
		s.ClearThumbnailPath()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *PlaylistUpsertBulk) SetUpdatedAt(v time.Time) *PlaylistUpsertBulk {
	return u.Update(func(s *PlaylistUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *PlaylistUpsertBulk) UpdateUpdatedAt() *PlaylistUpsertBulk {
	return u.Update(func(s *PlaylistUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *PlaylistUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the PlaylistCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PlaylistCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PlaylistUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

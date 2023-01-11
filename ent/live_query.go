// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/zibbp/ganymede/ent/channel"
	"github.com/zibbp/ganymede/ent/live"
	"github.com/zibbp/ganymede/ent/predicate"
)

// LiveQuery is the builder for querying Live entities.
type LiveQuery struct {
	config
	limit       *int
	offset      *int
	unique      *bool
	order       []OrderFunc
	fields      []string
	inters      []Interceptor
	predicates  []predicate.Live
	withChannel *ChannelQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LiveQuery builder.
func (lq *LiveQuery) Where(ps ...predicate.Live) *LiveQuery {
	lq.predicates = append(lq.predicates, ps...)
	return lq
}

// Limit the number of records to be returned by this query.
func (lq *LiveQuery) Limit(limit int) *LiveQuery {
	lq.limit = &limit
	return lq
}

// Offset to start from.
func (lq *LiveQuery) Offset(offset int) *LiveQuery {
	lq.offset = &offset
	return lq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lq *LiveQuery) Unique(unique bool) *LiveQuery {
	lq.unique = &unique
	return lq
}

// Order specifies how the records should be ordered.
func (lq *LiveQuery) Order(o ...OrderFunc) *LiveQuery {
	lq.order = append(lq.order, o...)
	return lq
}

// QueryChannel chains the current query on the "channel" edge.
func (lq *LiveQuery) QueryChannel() *ChannelQuery {
	query := (&ChannelClient{config: lq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := lq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := lq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(live.Table, live.FieldID, selector),
			sqlgraph.To(channel.Table, channel.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, live.ChannelTable, live.ChannelColumn),
		)
		fromU = sqlgraph.SetNeighbors(lq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Live entity from the query.
// Returns a *NotFoundError when no Live was found.
func (lq *LiveQuery) First(ctx context.Context) (*Live, error) {
	nodes, err := lq.Limit(1).All(newQueryContext(ctx, TypeLive, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{live.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lq *LiveQuery) FirstX(ctx context.Context) *Live {
	node, err := lq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Live ID from the query.
// Returns a *NotFoundError when no Live ID was found.
func (lq *LiveQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = lq.Limit(1).IDs(newQueryContext(ctx, TypeLive, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{live.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lq *LiveQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := lq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Live entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Live entity is found.
// Returns a *NotFoundError when no Live entities are found.
func (lq *LiveQuery) Only(ctx context.Context) (*Live, error) {
	nodes, err := lq.Limit(2).All(newQueryContext(ctx, TypeLive, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{live.Label}
	default:
		return nil, &NotSingularError{live.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lq *LiveQuery) OnlyX(ctx context.Context) *Live {
	node, err := lq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Live ID in the query.
// Returns a *NotSingularError when more than one Live ID is found.
// Returns a *NotFoundError when no entities are found.
func (lq *LiveQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = lq.Limit(2).IDs(newQueryContext(ctx, TypeLive, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{live.Label}
	default:
		err = &NotSingularError{live.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lq *LiveQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := lq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Lives.
func (lq *LiveQuery) All(ctx context.Context) ([]*Live, error) {
	ctx = newQueryContext(ctx, TypeLive, "All")
	if err := lq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Live, *LiveQuery]()
	return withInterceptors[[]*Live](ctx, lq, qr, lq.inters)
}

// AllX is like All, but panics if an error occurs.
func (lq *LiveQuery) AllX(ctx context.Context) []*Live {
	nodes, err := lq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Live IDs.
func (lq *LiveQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	ctx = newQueryContext(ctx, TypeLive, "IDs")
	if err := lq.Select(live.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lq *LiveQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := lq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lq *LiveQuery) Count(ctx context.Context) (int, error) {
	ctx = newQueryContext(ctx, TypeLive, "Count")
	if err := lq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, lq, querierCount[*LiveQuery](), lq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (lq *LiveQuery) CountX(ctx context.Context) int {
	count, err := lq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lq *LiveQuery) Exist(ctx context.Context) (bool, error) {
	ctx = newQueryContext(ctx, TypeLive, "Exist")
	switch _, err := lq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (lq *LiveQuery) ExistX(ctx context.Context) bool {
	exist, err := lq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LiveQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lq *LiveQuery) Clone() *LiveQuery {
	if lq == nil {
		return nil
	}
	return &LiveQuery{
		config:      lq.config,
		limit:       lq.limit,
		offset:      lq.offset,
		order:       append([]OrderFunc{}, lq.order...),
		inters:      append([]Interceptor{}, lq.inters...),
		predicates:  append([]predicate.Live{}, lq.predicates...),
		withChannel: lq.withChannel.Clone(),
		// clone intermediate query.
		sql:    lq.sql.Clone(),
		path:   lq.path,
		unique: lq.unique,
	}
}

// WithChannel tells the query-builder to eager-load the nodes that are connected to
// the "channel" edge. The optional arguments are used to configure the query builder of the edge.
func (lq *LiveQuery) WithChannel(opts ...func(*ChannelQuery)) *LiveQuery {
	query := (&ChannelClient{config: lq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	lq.withChannel = query
	return lq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		WatchLive bool `json:"watch_live,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Live.Query().
//		GroupBy(live.FieldWatchLive).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (lq *LiveQuery) GroupBy(field string, fields ...string) *LiveGroupBy {
	lq.fields = append([]string{field}, fields...)
	grbuild := &LiveGroupBy{build: lq}
	grbuild.flds = &lq.fields
	grbuild.label = live.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		WatchLive bool `json:"watch_live,omitempty"`
//	}
//
//	client.Live.Query().
//		Select(live.FieldWatchLive).
//		Scan(ctx, &v)
func (lq *LiveQuery) Select(fields ...string) *LiveSelect {
	lq.fields = append(lq.fields, fields...)
	sbuild := &LiveSelect{LiveQuery: lq}
	sbuild.label = live.Label
	sbuild.flds, sbuild.scan = &lq.fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LiveSelect configured with the given aggregations.
func (lq *LiveQuery) Aggregate(fns ...AggregateFunc) *LiveSelect {
	return lq.Select().Aggregate(fns...)
}

func (lq *LiveQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range lq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, lq); err != nil {
				return err
			}
		}
	}
	for _, f := range lq.fields {
		if !live.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if lq.path != nil {
		prev, err := lq.path(ctx)
		if err != nil {
			return err
		}
		lq.sql = prev
	}
	return nil
}

func (lq *LiveQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Live, error) {
	var (
		nodes       = []*Live{}
		withFKs     = lq.withFKs
		_spec       = lq.querySpec()
		loadedTypes = [1]bool{
			lq.withChannel != nil,
		}
	)
	if lq.withChannel != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, live.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Live).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Live{config: lq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := lq.withChannel; query != nil {
		if err := lq.loadChannel(ctx, query, nodes, nil,
			func(n *Live, e *Channel) { n.Edges.Channel = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (lq *LiveQuery) loadChannel(ctx context.Context, query *ChannelQuery, nodes []*Live, init func(*Live), assign func(*Live, *Channel)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Live)
	for i := range nodes {
		if nodes[i].channel_live == nil {
			continue
		}
		fk := *nodes[i].channel_live
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(channel.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "channel_live" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (lq *LiveQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lq.querySpec()
	_spec.Node.Columns = lq.fields
	if len(lq.fields) > 0 {
		_spec.Unique = lq.unique != nil && *lq.unique
	}
	return sqlgraph.CountNodes(ctx, lq.driver, _spec)
}

func (lq *LiveQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   live.Table,
			Columns: live.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: live.FieldID,
			},
		},
		From:   lq.sql,
		Unique: true,
	}
	if unique := lq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := lq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, live.FieldID)
		for i := range fields {
			if fields[i] != live.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := lq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lq *LiveQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lq.driver.Dialect())
	t1 := builder.Table(live.Table)
	columns := lq.fields
	if len(columns) == 0 {
		columns = live.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lq.sql != nil {
		selector = lq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lq.unique != nil && *lq.unique {
		selector.Distinct()
	}
	for _, p := range lq.predicates {
		p(selector)
	}
	for _, p := range lq.order {
		p(selector)
	}
	if offset := lq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LiveGroupBy is the group-by builder for Live entities.
type LiveGroupBy struct {
	selector
	build *LiveQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lgb *LiveGroupBy) Aggregate(fns ...AggregateFunc) *LiveGroupBy {
	lgb.fns = append(lgb.fns, fns...)
	return lgb
}

// Scan applies the selector query and scans the result into the given value.
func (lgb *LiveGroupBy) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeLive, "GroupBy")
	if err := lgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LiveQuery, *LiveGroupBy](ctx, lgb.build, lgb, lgb.build.inters, v)
}

func (lgb *LiveGroupBy) sqlScan(ctx context.Context, root *LiveQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(lgb.fns))
	for _, fn := range lgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*lgb.flds)+len(lgb.fns))
		for _, f := range *lgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*lgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LiveSelect is the builder for selecting fields of Live entities.
type LiveSelect struct {
	*LiveQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ls *LiveSelect) Aggregate(fns ...AggregateFunc) *LiveSelect {
	ls.fns = append(ls.fns, fns...)
	return ls
}

// Scan applies the selector query and scans the result into the given value.
func (ls *LiveSelect) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeLive, "Select")
	if err := ls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LiveQuery, *LiveSelect](ctx, ls.LiveQuery, ls, ls.inters, v)
}

func (ls *LiveSelect) sqlScan(ctx context.Context, root *LiveQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ls.fns))
	for _, fn := range ls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

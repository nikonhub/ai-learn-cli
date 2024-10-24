// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/nikonhub/ai-learn-cli/ent/item"
	"github.com/nikonhub/ai-learn-cli/ent/itemprogress"
	"github.com/nikonhub/ai-learn-cli/ent/predicate"
)

// ItemProgressQuery is the builder for querying ItemProgress entities.
type ItemProgressQuery struct {
	config
	ctx        *QueryContext
	order      []itemprogress.OrderOption
	inters     []Interceptor
	predicates []predicate.ItemProgress
	withItem   *ItemQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ItemProgressQuery builder.
func (ipq *ItemProgressQuery) Where(ps ...predicate.ItemProgress) *ItemProgressQuery {
	ipq.predicates = append(ipq.predicates, ps...)
	return ipq
}

// Limit the number of records to be returned by this query.
func (ipq *ItemProgressQuery) Limit(limit int) *ItemProgressQuery {
	ipq.ctx.Limit = &limit
	return ipq
}

// Offset to start from.
func (ipq *ItemProgressQuery) Offset(offset int) *ItemProgressQuery {
	ipq.ctx.Offset = &offset
	return ipq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ipq *ItemProgressQuery) Unique(unique bool) *ItemProgressQuery {
	ipq.ctx.Unique = &unique
	return ipq
}

// Order specifies how the records should be ordered.
func (ipq *ItemProgressQuery) Order(o ...itemprogress.OrderOption) *ItemProgressQuery {
	ipq.order = append(ipq.order, o...)
	return ipq
}

// QueryItem chains the current query on the "item" edge.
func (ipq *ItemProgressQuery) QueryItem() *ItemQuery {
	query := (&ItemClient{config: ipq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ipq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ipq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(itemprogress.Table, itemprogress.FieldID, selector),
			sqlgraph.To(item.Table, item.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, itemprogress.ItemTable, itemprogress.ItemColumn),
		)
		fromU = sqlgraph.SetNeighbors(ipq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ItemProgress entity from the query.
// Returns a *NotFoundError when no ItemProgress was found.
func (ipq *ItemProgressQuery) First(ctx context.Context) (*ItemProgress, error) {
	nodes, err := ipq.Limit(1).All(setContextOp(ctx, ipq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{itemprogress.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ipq *ItemProgressQuery) FirstX(ctx context.Context) *ItemProgress {
	node, err := ipq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ItemProgress ID from the query.
// Returns a *NotFoundError when no ItemProgress ID was found.
func (ipq *ItemProgressQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ipq.Limit(1).IDs(setContextOp(ctx, ipq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{itemprogress.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ipq *ItemProgressQuery) FirstIDX(ctx context.Context) int {
	id, err := ipq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ItemProgress entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ItemProgress entity is found.
// Returns a *NotFoundError when no ItemProgress entities are found.
func (ipq *ItemProgressQuery) Only(ctx context.Context) (*ItemProgress, error) {
	nodes, err := ipq.Limit(2).All(setContextOp(ctx, ipq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{itemprogress.Label}
	default:
		return nil, &NotSingularError{itemprogress.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ipq *ItemProgressQuery) OnlyX(ctx context.Context) *ItemProgress {
	node, err := ipq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ItemProgress ID in the query.
// Returns a *NotSingularError when more than one ItemProgress ID is found.
// Returns a *NotFoundError when no entities are found.
func (ipq *ItemProgressQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ipq.Limit(2).IDs(setContextOp(ctx, ipq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{itemprogress.Label}
	default:
		err = &NotSingularError{itemprogress.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ipq *ItemProgressQuery) OnlyIDX(ctx context.Context) int {
	id, err := ipq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ItemProgresses.
func (ipq *ItemProgressQuery) All(ctx context.Context) ([]*ItemProgress, error) {
	ctx = setContextOp(ctx, ipq.ctx, ent.OpQueryAll)
	if err := ipq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ItemProgress, *ItemProgressQuery]()
	return withInterceptors[[]*ItemProgress](ctx, ipq, qr, ipq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ipq *ItemProgressQuery) AllX(ctx context.Context) []*ItemProgress {
	nodes, err := ipq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ItemProgress IDs.
func (ipq *ItemProgressQuery) IDs(ctx context.Context) (ids []int, err error) {
	if ipq.ctx.Unique == nil && ipq.path != nil {
		ipq.Unique(true)
	}
	ctx = setContextOp(ctx, ipq.ctx, ent.OpQueryIDs)
	if err = ipq.Select(itemprogress.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ipq *ItemProgressQuery) IDsX(ctx context.Context) []int {
	ids, err := ipq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ipq *ItemProgressQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ipq.ctx, ent.OpQueryCount)
	if err := ipq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ipq, querierCount[*ItemProgressQuery](), ipq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ipq *ItemProgressQuery) CountX(ctx context.Context) int {
	count, err := ipq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ipq *ItemProgressQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ipq.ctx, ent.OpQueryExist)
	switch _, err := ipq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ipq *ItemProgressQuery) ExistX(ctx context.Context) bool {
	exist, err := ipq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ItemProgressQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ipq *ItemProgressQuery) Clone() *ItemProgressQuery {
	if ipq == nil {
		return nil
	}
	return &ItemProgressQuery{
		config:     ipq.config,
		ctx:        ipq.ctx.Clone(),
		order:      append([]itemprogress.OrderOption{}, ipq.order...),
		inters:     append([]Interceptor{}, ipq.inters...),
		predicates: append([]predicate.ItemProgress{}, ipq.predicates...),
		withItem:   ipq.withItem.Clone(),
		// clone intermediate query.
		sql:  ipq.sql.Clone(),
		path: ipq.path,
	}
}

// WithItem tells the query-builder to eager-load the nodes that are connected to
// the "item" edge. The optional arguments are used to configure the query builder of the edge.
func (ipq *ItemProgressQuery) WithItem(opts ...func(*ItemQuery)) *ItemProgressQuery {
	query := (&ItemClient{config: ipq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ipq.withItem = query
	return ipq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		NextReviewDate time.Time `json:"next_review_date,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ItemProgress.Query().
//		GroupBy(itemprogress.FieldNextReviewDate).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ipq *ItemProgressQuery) GroupBy(field string, fields ...string) *ItemProgressGroupBy {
	ipq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ItemProgressGroupBy{build: ipq}
	grbuild.flds = &ipq.ctx.Fields
	grbuild.label = itemprogress.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		NextReviewDate time.Time `json:"next_review_date,omitempty"`
//	}
//
//	client.ItemProgress.Query().
//		Select(itemprogress.FieldNextReviewDate).
//		Scan(ctx, &v)
func (ipq *ItemProgressQuery) Select(fields ...string) *ItemProgressSelect {
	ipq.ctx.Fields = append(ipq.ctx.Fields, fields...)
	sbuild := &ItemProgressSelect{ItemProgressQuery: ipq}
	sbuild.label = itemprogress.Label
	sbuild.flds, sbuild.scan = &ipq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ItemProgressSelect configured with the given aggregations.
func (ipq *ItemProgressQuery) Aggregate(fns ...AggregateFunc) *ItemProgressSelect {
	return ipq.Select().Aggregate(fns...)
}

func (ipq *ItemProgressQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ipq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ipq); err != nil {
				return err
			}
		}
	}
	for _, f := range ipq.ctx.Fields {
		if !itemprogress.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ipq.path != nil {
		prev, err := ipq.path(ctx)
		if err != nil {
			return err
		}
		ipq.sql = prev
	}
	return nil
}

func (ipq *ItemProgressQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ItemProgress, error) {
	var (
		nodes       = []*ItemProgress{}
		withFKs     = ipq.withFKs
		_spec       = ipq.querySpec()
		loadedTypes = [1]bool{
			ipq.withItem != nil,
		}
	)
	if ipq.withItem != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, itemprogress.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ItemProgress).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ItemProgress{config: ipq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ipq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ipq.withItem; query != nil {
		if err := ipq.loadItem(ctx, query, nodes, nil,
			func(n *ItemProgress, e *Item) { n.Edges.Item = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ipq *ItemProgressQuery) loadItem(ctx context.Context, query *ItemQuery, nodes []*ItemProgress, init func(*ItemProgress), assign func(*ItemProgress, *Item)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*ItemProgress)
	for i := range nodes {
		if nodes[i].item_progress == nil {
			continue
		}
		fk := *nodes[i].item_progress
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(item.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "item_progress" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ipq *ItemProgressQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ipq.querySpec()
	_spec.Node.Columns = ipq.ctx.Fields
	if len(ipq.ctx.Fields) > 0 {
		_spec.Unique = ipq.ctx.Unique != nil && *ipq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ipq.driver, _spec)
}

func (ipq *ItemProgressQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(itemprogress.Table, itemprogress.Columns, sqlgraph.NewFieldSpec(itemprogress.FieldID, field.TypeInt))
	_spec.From = ipq.sql
	if unique := ipq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ipq.path != nil {
		_spec.Unique = true
	}
	if fields := ipq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, itemprogress.FieldID)
		for i := range fields {
			if fields[i] != itemprogress.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ipq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ipq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ipq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ipq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ipq *ItemProgressQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ipq.driver.Dialect())
	t1 := builder.Table(itemprogress.Table)
	columns := ipq.ctx.Fields
	if len(columns) == 0 {
		columns = itemprogress.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ipq.sql != nil {
		selector = ipq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ipq.ctx.Unique != nil && *ipq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ipq.predicates {
		p(selector)
	}
	for _, p := range ipq.order {
		p(selector)
	}
	if offset := ipq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ipq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ItemProgressGroupBy is the group-by builder for ItemProgress entities.
type ItemProgressGroupBy struct {
	selector
	build *ItemProgressQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ipgb *ItemProgressGroupBy) Aggregate(fns ...AggregateFunc) *ItemProgressGroupBy {
	ipgb.fns = append(ipgb.fns, fns...)
	return ipgb
}

// Scan applies the selector query and scans the result into the given value.
func (ipgb *ItemProgressGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ipgb.build.ctx, ent.OpQueryGroupBy)
	if err := ipgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ItemProgressQuery, *ItemProgressGroupBy](ctx, ipgb.build, ipgb, ipgb.build.inters, v)
}

func (ipgb *ItemProgressGroupBy) sqlScan(ctx context.Context, root *ItemProgressQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ipgb.fns))
	for _, fn := range ipgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ipgb.flds)+len(ipgb.fns))
		for _, f := range *ipgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ipgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ipgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ItemProgressSelect is the builder for selecting fields of ItemProgress entities.
type ItemProgressSelect struct {
	*ItemProgressQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ips *ItemProgressSelect) Aggregate(fns ...AggregateFunc) *ItemProgressSelect {
	ips.fns = append(ips.fns, fns...)
	return ips
}

// Scan applies the selector query and scans the result into the given value.
func (ips *ItemProgressSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ips.ctx, ent.OpQuerySelect)
	if err := ips.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ItemProgressQuery, *ItemProgressSelect](ctx, ips.ItemProgressQuery, ips, ips.inters, v)
}

func (ips *ItemProgressSelect) sqlScan(ctx context.Context, root *ItemProgressQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ips.fns))
	for _, fn := range ips.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ips.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ips.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

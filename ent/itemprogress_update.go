// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/nikonhub/ai-learn-cli/ent/item"
	"github.com/nikonhub/ai-learn-cli/ent/itemprogress"
	"github.com/nikonhub/ai-learn-cli/ent/predicate"
)

// ItemProgressUpdate is the builder for updating ItemProgress entities.
type ItemProgressUpdate struct {
	config
	hooks    []Hook
	mutation *ItemProgressMutation
}

// Where appends a list predicates to the ItemProgressUpdate builder.
func (ipu *ItemProgressUpdate) Where(ps ...predicate.ItemProgress) *ItemProgressUpdate {
	ipu.mutation.Where(ps...)
	return ipu
}

// SetNextReviewDate sets the "next_review_date" field.
func (ipu *ItemProgressUpdate) SetNextReviewDate(t time.Time) *ItemProgressUpdate {
	ipu.mutation.SetNextReviewDate(t)
	return ipu
}

// SetNillableNextReviewDate sets the "next_review_date" field if the given value is not nil.
func (ipu *ItemProgressUpdate) SetNillableNextReviewDate(t *time.Time) *ItemProgressUpdate {
	if t != nil {
		ipu.SetNextReviewDate(*t)
	}
	return ipu
}

// SetIntervalDays sets the "interval_days" field.
func (ipu *ItemProgressUpdate) SetIntervalDays(i int) *ItemProgressUpdate {
	ipu.mutation.ResetIntervalDays()
	ipu.mutation.SetIntervalDays(i)
	return ipu
}

// SetNillableIntervalDays sets the "interval_days" field if the given value is not nil.
func (ipu *ItemProgressUpdate) SetNillableIntervalDays(i *int) *ItemProgressUpdate {
	if i != nil {
		ipu.SetIntervalDays(*i)
	}
	return ipu
}

// AddIntervalDays adds i to the "interval_days" field.
func (ipu *ItemProgressUpdate) AddIntervalDays(i int) *ItemProgressUpdate {
	ipu.mutation.AddIntervalDays(i)
	return ipu
}

// SetEaseFactor sets the "ease_factor" field.
func (ipu *ItemProgressUpdate) SetEaseFactor(f float64) *ItemProgressUpdate {
	ipu.mutation.ResetEaseFactor()
	ipu.mutation.SetEaseFactor(f)
	return ipu
}

// SetNillableEaseFactor sets the "ease_factor" field if the given value is not nil.
func (ipu *ItemProgressUpdate) SetNillableEaseFactor(f *float64) *ItemProgressUpdate {
	if f != nil {
		ipu.SetEaseFactor(*f)
	}
	return ipu
}

// AddEaseFactor adds f to the "ease_factor" field.
func (ipu *ItemProgressUpdate) AddEaseFactor(f float64) *ItemProgressUpdate {
	ipu.mutation.AddEaseFactor(f)
	return ipu
}

// SetStreak sets the "streak" field.
func (ipu *ItemProgressUpdate) SetStreak(i int) *ItemProgressUpdate {
	ipu.mutation.ResetStreak()
	ipu.mutation.SetStreak(i)
	return ipu
}

// SetNillableStreak sets the "streak" field if the given value is not nil.
func (ipu *ItemProgressUpdate) SetNillableStreak(i *int) *ItemProgressUpdate {
	if i != nil {
		ipu.SetStreak(*i)
	}
	return ipu
}

// AddStreak adds i to the "streak" field.
func (ipu *ItemProgressUpdate) AddStreak(i int) *ItemProgressUpdate {
	ipu.mutation.AddStreak(i)
	return ipu
}

// SetItemID sets the "item" edge to the Item entity by ID.
func (ipu *ItemProgressUpdate) SetItemID(id int) *ItemProgressUpdate {
	ipu.mutation.SetItemID(id)
	return ipu
}

// SetNillableItemID sets the "item" edge to the Item entity by ID if the given value is not nil.
func (ipu *ItemProgressUpdate) SetNillableItemID(id *int) *ItemProgressUpdate {
	if id != nil {
		ipu = ipu.SetItemID(*id)
	}
	return ipu
}

// SetItem sets the "item" edge to the Item entity.
func (ipu *ItemProgressUpdate) SetItem(i *Item) *ItemProgressUpdate {
	return ipu.SetItemID(i.ID)
}

// Mutation returns the ItemProgressMutation object of the builder.
func (ipu *ItemProgressUpdate) Mutation() *ItemProgressMutation {
	return ipu.mutation
}

// ClearItem clears the "item" edge to the Item entity.
func (ipu *ItemProgressUpdate) ClearItem() *ItemProgressUpdate {
	ipu.mutation.ClearItem()
	return ipu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ipu *ItemProgressUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ipu.sqlSave, ipu.mutation, ipu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ipu *ItemProgressUpdate) SaveX(ctx context.Context) int {
	affected, err := ipu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ipu *ItemProgressUpdate) Exec(ctx context.Context) error {
	_, err := ipu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ipu *ItemProgressUpdate) ExecX(ctx context.Context) {
	if err := ipu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ipu *ItemProgressUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(itemprogress.Table, itemprogress.Columns, sqlgraph.NewFieldSpec(itemprogress.FieldID, field.TypeInt))
	if ps := ipu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ipu.mutation.NextReviewDate(); ok {
		_spec.SetField(itemprogress.FieldNextReviewDate, field.TypeTime, value)
	}
	if value, ok := ipu.mutation.IntervalDays(); ok {
		_spec.SetField(itemprogress.FieldIntervalDays, field.TypeInt, value)
	}
	if value, ok := ipu.mutation.AddedIntervalDays(); ok {
		_spec.AddField(itemprogress.FieldIntervalDays, field.TypeInt, value)
	}
	if value, ok := ipu.mutation.EaseFactor(); ok {
		_spec.SetField(itemprogress.FieldEaseFactor, field.TypeFloat64, value)
	}
	if value, ok := ipu.mutation.AddedEaseFactor(); ok {
		_spec.AddField(itemprogress.FieldEaseFactor, field.TypeFloat64, value)
	}
	if value, ok := ipu.mutation.Streak(); ok {
		_spec.SetField(itemprogress.FieldStreak, field.TypeInt, value)
	}
	if value, ok := ipu.mutation.AddedStreak(); ok {
		_spec.AddField(itemprogress.FieldStreak, field.TypeInt, value)
	}
	if ipu.mutation.ItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   itemprogress.ItemTable,
			Columns: []string{itemprogress.ItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(item.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ipu.mutation.ItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   itemprogress.ItemTable,
			Columns: []string{itemprogress.ItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(item.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ipu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{itemprogress.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ipu.mutation.done = true
	return n, nil
}

// ItemProgressUpdateOne is the builder for updating a single ItemProgress entity.
type ItemProgressUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ItemProgressMutation
}

// SetNextReviewDate sets the "next_review_date" field.
func (ipuo *ItemProgressUpdateOne) SetNextReviewDate(t time.Time) *ItemProgressUpdateOne {
	ipuo.mutation.SetNextReviewDate(t)
	return ipuo
}

// SetNillableNextReviewDate sets the "next_review_date" field if the given value is not nil.
func (ipuo *ItemProgressUpdateOne) SetNillableNextReviewDate(t *time.Time) *ItemProgressUpdateOne {
	if t != nil {
		ipuo.SetNextReviewDate(*t)
	}
	return ipuo
}

// SetIntervalDays sets the "interval_days" field.
func (ipuo *ItemProgressUpdateOne) SetIntervalDays(i int) *ItemProgressUpdateOne {
	ipuo.mutation.ResetIntervalDays()
	ipuo.mutation.SetIntervalDays(i)
	return ipuo
}

// SetNillableIntervalDays sets the "interval_days" field if the given value is not nil.
func (ipuo *ItemProgressUpdateOne) SetNillableIntervalDays(i *int) *ItemProgressUpdateOne {
	if i != nil {
		ipuo.SetIntervalDays(*i)
	}
	return ipuo
}

// AddIntervalDays adds i to the "interval_days" field.
func (ipuo *ItemProgressUpdateOne) AddIntervalDays(i int) *ItemProgressUpdateOne {
	ipuo.mutation.AddIntervalDays(i)
	return ipuo
}

// SetEaseFactor sets the "ease_factor" field.
func (ipuo *ItemProgressUpdateOne) SetEaseFactor(f float64) *ItemProgressUpdateOne {
	ipuo.mutation.ResetEaseFactor()
	ipuo.mutation.SetEaseFactor(f)
	return ipuo
}

// SetNillableEaseFactor sets the "ease_factor" field if the given value is not nil.
func (ipuo *ItemProgressUpdateOne) SetNillableEaseFactor(f *float64) *ItemProgressUpdateOne {
	if f != nil {
		ipuo.SetEaseFactor(*f)
	}
	return ipuo
}

// AddEaseFactor adds f to the "ease_factor" field.
func (ipuo *ItemProgressUpdateOne) AddEaseFactor(f float64) *ItemProgressUpdateOne {
	ipuo.mutation.AddEaseFactor(f)
	return ipuo
}

// SetStreak sets the "streak" field.
func (ipuo *ItemProgressUpdateOne) SetStreak(i int) *ItemProgressUpdateOne {
	ipuo.mutation.ResetStreak()
	ipuo.mutation.SetStreak(i)
	return ipuo
}

// SetNillableStreak sets the "streak" field if the given value is not nil.
func (ipuo *ItemProgressUpdateOne) SetNillableStreak(i *int) *ItemProgressUpdateOne {
	if i != nil {
		ipuo.SetStreak(*i)
	}
	return ipuo
}

// AddStreak adds i to the "streak" field.
func (ipuo *ItemProgressUpdateOne) AddStreak(i int) *ItemProgressUpdateOne {
	ipuo.mutation.AddStreak(i)
	return ipuo
}

// SetItemID sets the "item" edge to the Item entity by ID.
func (ipuo *ItemProgressUpdateOne) SetItemID(id int) *ItemProgressUpdateOne {
	ipuo.mutation.SetItemID(id)
	return ipuo
}

// SetNillableItemID sets the "item" edge to the Item entity by ID if the given value is not nil.
func (ipuo *ItemProgressUpdateOne) SetNillableItemID(id *int) *ItemProgressUpdateOne {
	if id != nil {
		ipuo = ipuo.SetItemID(*id)
	}
	return ipuo
}

// SetItem sets the "item" edge to the Item entity.
func (ipuo *ItemProgressUpdateOne) SetItem(i *Item) *ItemProgressUpdateOne {
	return ipuo.SetItemID(i.ID)
}

// Mutation returns the ItemProgressMutation object of the builder.
func (ipuo *ItemProgressUpdateOne) Mutation() *ItemProgressMutation {
	return ipuo.mutation
}

// ClearItem clears the "item" edge to the Item entity.
func (ipuo *ItemProgressUpdateOne) ClearItem() *ItemProgressUpdateOne {
	ipuo.mutation.ClearItem()
	return ipuo
}

// Where appends a list predicates to the ItemProgressUpdate builder.
func (ipuo *ItemProgressUpdateOne) Where(ps ...predicate.ItemProgress) *ItemProgressUpdateOne {
	ipuo.mutation.Where(ps...)
	return ipuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ipuo *ItemProgressUpdateOne) Select(field string, fields ...string) *ItemProgressUpdateOne {
	ipuo.fields = append([]string{field}, fields...)
	return ipuo
}

// Save executes the query and returns the updated ItemProgress entity.
func (ipuo *ItemProgressUpdateOne) Save(ctx context.Context) (*ItemProgress, error) {
	return withHooks(ctx, ipuo.sqlSave, ipuo.mutation, ipuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ipuo *ItemProgressUpdateOne) SaveX(ctx context.Context) *ItemProgress {
	node, err := ipuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ipuo *ItemProgressUpdateOne) Exec(ctx context.Context) error {
	_, err := ipuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ipuo *ItemProgressUpdateOne) ExecX(ctx context.Context) {
	if err := ipuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ipuo *ItemProgressUpdateOne) sqlSave(ctx context.Context) (_node *ItemProgress, err error) {
	_spec := sqlgraph.NewUpdateSpec(itemprogress.Table, itemprogress.Columns, sqlgraph.NewFieldSpec(itemprogress.FieldID, field.TypeInt))
	id, ok := ipuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ItemProgress.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ipuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, itemprogress.FieldID)
		for _, f := range fields {
			if !itemprogress.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != itemprogress.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ipuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ipuo.mutation.NextReviewDate(); ok {
		_spec.SetField(itemprogress.FieldNextReviewDate, field.TypeTime, value)
	}
	if value, ok := ipuo.mutation.IntervalDays(); ok {
		_spec.SetField(itemprogress.FieldIntervalDays, field.TypeInt, value)
	}
	if value, ok := ipuo.mutation.AddedIntervalDays(); ok {
		_spec.AddField(itemprogress.FieldIntervalDays, field.TypeInt, value)
	}
	if value, ok := ipuo.mutation.EaseFactor(); ok {
		_spec.SetField(itemprogress.FieldEaseFactor, field.TypeFloat64, value)
	}
	if value, ok := ipuo.mutation.AddedEaseFactor(); ok {
		_spec.AddField(itemprogress.FieldEaseFactor, field.TypeFloat64, value)
	}
	if value, ok := ipuo.mutation.Streak(); ok {
		_spec.SetField(itemprogress.FieldStreak, field.TypeInt, value)
	}
	if value, ok := ipuo.mutation.AddedStreak(); ok {
		_spec.AddField(itemprogress.FieldStreak, field.TypeInt, value)
	}
	if ipuo.mutation.ItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   itemprogress.ItemTable,
			Columns: []string{itemprogress.ItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(item.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ipuo.mutation.ItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   itemprogress.ItemTable,
			Columns: []string{itemprogress.ItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(item.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ItemProgress{config: ipuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ipuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{itemprogress.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ipuo.mutation.done = true
	return _node, nil
}

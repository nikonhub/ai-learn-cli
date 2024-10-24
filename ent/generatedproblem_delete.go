// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/nikonhub/ai-learn-cli/ent/generatedproblem"
	"github.com/nikonhub/ai-learn-cli/ent/predicate"
)

// GeneratedProblemDelete is the builder for deleting a GeneratedProblem entity.
type GeneratedProblemDelete struct {
	config
	hooks    []Hook
	mutation *GeneratedProblemMutation
}

// Where appends a list predicates to the GeneratedProblemDelete builder.
func (gpd *GeneratedProblemDelete) Where(ps ...predicate.GeneratedProblem) *GeneratedProblemDelete {
	gpd.mutation.Where(ps...)
	return gpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gpd *GeneratedProblemDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, gpd.sqlExec, gpd.mutation, gpd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (gpd *GeneratedProblemDelete) ExecX(ctx context.Context) int {
	n, err := gpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gpd *GeneratedProblemDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(generatedproblem.Table, sqlgraph.NewFieldSpec(generatedproblem.FieldID, field.TypeInt))
	if ps := gpd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, gpd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	gpd.mutation.done = true
	return affected, err
}

// GeneratedProblemDeleteOne is the builder for deleting a single GeneratedProblem entity.
type GeneratedProblemDeleteOne struct {
	gpd *GeneratedProblemDelete
}

// Where appends a list predicates to the GeneratedProblemDelete builder.
func (gpdo *GeneratedProblemDeleteOne) Where(ps ...predicate.GeneratedProblem) *GeneratedProblemDeleteOne {
	gpdo.gpd.mutation.Where(ps...)
	return gpdo
}

// Exec executes the deletion query.
func (gpdo *GeneratedProblemDeleteOne) Exec(ctx context.Context) error {
	n, err := gpdo.gpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{generatedproblem.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gpdo *GeneratedProblemDeleteOne) ExecX(ctx context.Context) {
	if err := gpdo.Exec(ctx); err != nil {
		panic(err)
	}
}

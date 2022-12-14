// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/order-manager/pkg/db/ent/outofgas"
	"github.com/NpoolPlatform/order-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// OutOfGasUpdate is the builder for updating OutOfGas entities.
type OutOfGasUpdate struct {
	config
	hooks     []Hook
	mutation  *OutOfGasMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the OutOfGasUpdate builder.
func (oogu *OutOfGasUpdate) Where(ps ...predicate.OutOfGas) *OutOfGasUpdate {
	oogu.mutation.Where(ps...)
	return oogu
}

// SetCreatedAt sets the "created_at" field.
func (oogu *OutOfGasUpdate) SetCreatedAt(u uint32) *OutOfGasUpdate {
	oogu.mutation.ResetCreatedAt()
	oogu.mutation.SetCreatedAt(u)
	return oogu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (oogu *OutOfGasUpdate) SetNillableCreatedAt(u *uint32) *OutOfGasUpdate {
	if u != nil {
		oogu.SetCreatedAt(*u)
	}
	return oogu
}

// AddCreatedAt adds u to the "created_at" field.
func (oogu *OutOfGasUpdate) AddCreatedAt(u int32) *OutOfGasUpdate {
	oogu.mutation.AddCreatedAt(u)
	return oogu
}

// SetUpdatedAt sets the "updated_at" field.
func (oogu *OutOfGasUpdate) SetUpdatedAt(u uint32) *OutOfGasUpdate {
	oogu.mutation.ResetUpdatedAt()
	oogu.mutation.SetUpdatedAt(u)
	return oogu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (oogu *OutOfGasUpdate) AddUpdatedAt(u int32) *OutOfGasUpdate {
	oogu.mutation.AddUpdatedAt(u)
	return oogu
}

// SetDeletedAt sets the "deleted_at" field.
func (oogu *OutOfGasUpdate) SetDeletedAt(u uint32) *OutOfGasUpdate {
	oogu.mutation.ResetDeletedAt()
	oogu.mutation.SetDeletedAt(u)
	return oogu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (oogu *OutOfGasUpdate) SetNillableDeletedAt(u *uint32) *OutOfGasUpdate {
	if u != nil {
		oogu.SetDeletedAt(*u)
	}
	return oogu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (oogu *OutOfGasUpdate) AddDeletedAt(u int32) *OutOfGasUpdate {
	oogu.mutation.AddDeletedAt(u)
	return oogu
}

// SetOrderID sets the "order_id" field.
func (oogu *OutOfGasUpdate) SetOrderID(u uuid.UUID) *OutOfGasUpdate {
	oogu.mutation.SetOrderID(u)
	return oogu
}

// SetStart sets the "start" field.
func (oogu *OutOfGasUpdate) SetStart(u uint32) *OutOfGasUpdate {
	oogu.mutation.ResetStart()
	oogu.mutation.SetStart(u)
	return oogu
}

// SetNillableStart sets the "start" field if the given value is not nil.
func (oogu *OutOfGasUpdate) SetNillableStart(u *uint32) *OutOfGasUpdate {
	if u != nil {
		oogu.SetStart(*u)
	}
	return oogu
}

// AddStart adds u to the "start" field.
func (oogu *OutOfGasUpdate) AddStart(u int32) *OutOfGasUpdate {
	oogu.mutation.AddStart(u)
	return oogu
}

// ClearStart clears the value of the "start" field.
func (oogu *OutOfGasUpdate) ClearStart() *OutOfGasUpdate {
	oogu.mutation.ClearStart()
	return oogu
}

// SetEnd sets the "end" field.
func (oogu *OutOfGasUpdate) SetEnd(u uint32) *OutOfGasUpdate {
	oogu.mutation.ResetEnd()
	oogu.mutation.SetEnd(u)
	return oogu
}

// SetNillableEnd sets the "end" field if the given value is not nil.
func (oogu *OutOfGasUpdate) SetNillableEnd(u *uint32) *OutOfGasUpdate {
	if u != nil {
		oogu.SetEnd(*u)
	}
	return oogu
}

// AddEnd adds u to the "end" field.
func (oogu *OutOfGasUpdate) AddEnd(u int32) *OutOfGasUpdate {
	oogu.mutation.AddEnd(u)
	return oogu
}

// ClearEnd clears the value of the "end" field.
func (oogu *OutOfGasUpdate) ClearEnd() *OutOfGasUpdate {
	oogu.mutation.ClearEnd()
	return oogu
}

// Mutation returns the OutOfGasMutation object of the builder.
func (oogu *OutOfGasUpdate) Mutation() *OutOfGasMutation {
	return oogu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (oogu *OutOfGasUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := oogu.defaults(); err != nil {
		return 0, err
	}
	if len(oogu.hooks) == 0 {
		affected, err = oogu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutOfGasMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			oogu.mutation = mutation
			affected, err = oogu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(oogu.hooks) - 1; i >= 0; i-- {
			if oogu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = oogu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oogu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (oogu *OutOfGasUpdate) SaveX(ctx context.Context) int {
	affected, err := oogu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (oogu *OutOfGasUpdate) Exec(ctx context.Context) error {
	_, err := oogu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oogu *OutOfGasUpdate) ExecX(ctx context.Context) {
	if err := oogu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oogu *OutOfGasUpdate) defaults() error {
	if _, ok := oogu.mutation.UpdatedAt(); !ok {
		if outofgas.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized outofgas.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := outofgas.UpdateDefaultUpdatedAt()
		oogu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (oogu *OutOfGasUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OutOfGasUpdate {
	oogu.modifiers = append(oogu.modifiers, modifiers...)
	return oogu
}

func (oogu *OutOfGasUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outofgas.Table,
			Columns: outofgas.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: outofgas.FieldID,
			},
		},
	}
	if ps := oogu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oogu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldCreatedAt,
		})
	}
	if value, ok := oogu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldCreatedAt,
		})
	}
	if value, ok := oogu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldUpdatedAt,
		})
	}
	if value, ok := oogu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldUpdatedAt,
		})
	}
	if value, ok := oogu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldDeletedAt,
		})
	}
	if value, ok := oogu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldDeletedAt,
		})
	}
	if value, ok := oogu.mutation.OrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: outofgas.FieldOrderID,
		})
	}
	if value, ok := oogu.mutation.Start(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldStart,
		})
	}
	if value, ok := oogu.mutation.AddedStart(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldStart,
		})
	}
	if oogu.mutation.StartCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: outofgas.FieldStart,
		})
	}
	if value, ok := oogu.mutation.End(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldEnd,
		})
	}
	if value, ok := oogu.mutation.AddedEnd(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldEnd,
		})
	}
	if oogu.mutation.EndCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: outofgas.FieldEnd,
		})
	}
	_spec.Modifiers = oogu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, oogu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{outofgas.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// OutOfGasUpdateOne is the builder for updating a single OutOfGas entity.
type OutOfGasUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *OutOfGasMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (ooguo *OutOfGasUpdateOne) SetCreatedAt(u uint32) *OutOfGasUpdateOne {
	ooguo.mutation.ResetCreatedAt()
	ooguo.mutation.SetCreatedAt(u)
	return ooguo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ooguo *OutOfGasUpdateOne) SetNillableCreatedAt(u *uint32) *OutOfGasUpdateOne {
	if u != nil {
		ooguo.SetCreatedAt(*u)
	}
	return ooguo
}

// AddCreatedAt adds u to the "created_at" field.
func (ooguo *OutOfGasUpdateOne) AddCreatedAt(u int32) *OutOfGasUpdateOne {
	ooguo.mutation.AddCreatedAt(u)
	return ooguo
}

// SetUpdatedAt sets the "updated_at" field.
func (ooguo *OutOfGasUpdateOne) SetUpdatedAt(u uint32) *OutOfGasUpdateOne {
	ooguo.mutation.ResetUpdatedAt()
	ooguo.mutation.SetUpdatedAt(u)
	return ooguo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (ooguo *OutOfGasUpdateOne) AddUpdatedAt(u int32) *OutOfGasUpdateOne {
	ooguo.mutation.AddUpdatedAt(u)
	return ooguo
}

// SetDeletedAt sets the "deleted_at" field.
func (ooguo *OutOfGasUpdateOne) SetDeletedAt(u uint32) *OutOfGasUpdateOne {
	ooguo.mutation.ResetDeletedAt()
	ooguo.mutation.SetDeletedAt(u)
	return ooguo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ooguo *OutOfGasUpdateOne) SetNillableDeletedAt(u *uint32) *OutOfGasUpdateOne {
	if u != nil {
		ooguo.SetDeletedAt(*u)
	}
	return ooguo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (ooguo *OutOfGasUpdateOne) AddDeletedAt(u int32) *OutOfGasUpdateOne {
	ooguo.mutation.AddDeletedAt(u)
	return ooguo
}

// SetOrderID sets the "order_id" field.
func (ooguo *OutOfGasUpdateOne) SetOrderID(u uuid.UUID) *OutOfGasUpdateOne {
	ooguo.mutation.SetOrderID(u)
	return ooguo
}

// SetStart sets the "start" field.
func (ooguo *OutOfGasUpdateOne) SetStart(u uint32) *OutOfGasUpdateOne {
	ooguo.mutation.ResetStart()
	ooguo.mutation.SetStart(u)
	return ooguo
}

// SetNillableStart sets the "start" field if the given value is not nil.
func (ooguo *OutOfGasUpdateOne) SetNillableStart(u *uint32) *OutOfGasUpdateOne {
	if u != nil {
		ooguo.SetStart(*u)
	}
	return ooguo
}

// AddStart adds u to the "start" field.
func (ooguo *OutOfGasUpdateOne) AddStart(u int32) *OutOfGasUpdateOne {
	ooguo.mutation.AddStart(u)
	return ooguo
}

// ClearStart clears the value of the "start" field.
func (ooguo *OutOfGasUpdateOne) ClearStart() *OutOfGasUpdateOne {
	ooguo.mutation.ClearStart()
	return ooguo
}

// SetEnd sets the "end" field.
func (ooguo *OutOfGasUpdateOne) SetEnd(u uint32) *OutOfGasUpdateOne {
	ooguo.mutation.ResetEnd()
	ooguo.mutation.SetEnd(u)
	return ooguo
}

// SetNillableEnd sets the "end" field if the given value is not nil.
func (ooguo *OutOfGasUpdateOne) SetNillableEnd(u *uint32) *OutOfGasUpdateOne {
	if u != nil {
		ooguo.SetEnd(*u)
	}
	return ooguo
}

// AddEnd adds u to the "end" field.
func (ooguo *OutOfGasUpdateOne) AddEnd(u int32) *OutOfGasUpdateOne {
	ooguo.mutation.AddEnd(u)
	return ooguo
}

// ClearEnd clears the value of the "end" field.
func (ooguo *OutOfGasUpdateOne) ClearEnd() *OutOfGasUpdateOne {
	ooguo.mutation.ClearEnd()
	return ooguo
}

// Mutation returns the OutOfGasMutation object of the builder.
func (ooguo *OutOfGasUpdateOne) Mutation() *OutOfGasMutation {
	return ooguo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ooguo *OutOfGasUpdateOne) Select(field string, fields ...string) *OutOfGasUpdateOne {
	ooguo.fields = append([]string{field}, fields...)
	return ooguo
}

// Save executes the query and returns the updated OutOfGas entity.
func (ooguo *OutOfGasUpdateOne) Save(ctx context.Context) (*OutOfGas, error) {
	var (
		err  error
		node *OutOfGas
	)
	if err := ooguo.defaults(); err != nil {
		return nil, err
	}
	if len(ooguo.hooks) == 0 {
		node, err = ooguo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OutOfGasMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ooguo.mutation = mutation
			node, err = ooguo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ooguo.hooks) - 1; i >= 0; i-- {
			if ooguo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ooguo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ooguo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*OutOfGas)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from OutOfGasMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ooguo *OutOfGasUpdateOne) SaveX(ctx context.Context) *OutOfGas {
	node, err := ooguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ooguo *OutOfGasUpdateOne) Exec(ctx context.Context) error {
	_, err := ooguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ooguo *OutOfGasUpdateOne) ExecX(ctx context.Context) {
	if err := ooguo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ooguo *OutOfGasUpdateOne) defaults() error {
	if _, ok := ooguo.mutation.UpdatedAt(); !ok {
		if outofgas.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized outofgas.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := outofgas.UpdateDefaultUpdatedAt()
		ooguo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ooguo *OutOfGasUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OutOfGasUpdateOne {
	ooguo.modifiers = append(ooguo.modifiers, modifiers...)
	return ooguo
}

func (ooguo *OutOfGasUpdateOne) sqlSave(ctx context.Context) (_node *OutOfGas, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outofgas.Table,
			Columns: outofgas.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: outofgas.FieldID,
			},
		},
	}
	id, ok := ooguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OutOfGas.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ooguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, outofgas.FieldID)
		for _, f := range fields {
			if !outofgas.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != outofgas.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ooguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ooguo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldCreatedAt,
		})
	}
	if value, ok := ooguo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldCreatedAt,
		})
	}
	if value, ok := ooguo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldUpdatedAt,
		})
	}
	if value, ok := ooguo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldUpdatedAt,
		})
	}
	if value, ok := ooguo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldDeletedAt,
		})
	}
	if value, ok := ooguo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldDeletedAt,
		})
	}
	if value, ok := ooguo.mutation.OrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: outofgas.FieldOrderID,
		})
	}
	if value, ok := ooguo.mutation.Start(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldStart,
		})
	}
	if value, ok := ooguo.mutation.AddedStart(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldStart,
		})
	}
	if ooguo.mutation.StartCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: outofgas.FieldStart,
		})
	}
	if value, ok := ooguo.mutation.End(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldEnd,
		})
	}
	if value, ok := ooguo.mutation.AddedEnd(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: outofgas.FieldEnd,
		})
	}
	if ooguo.mutation.EndCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: outofgas.FieldEnd,
		})
	}
	_spec.Modifiers = ooguo.modifiers
	_node = &OutOfGas{config: ooguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ooguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{outofgas.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/order-manager/pkg/db/ent/order"
	"github.com/NpoolPlatform/order-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// OrderUpdate is the builder for updating Order entities.
type OrderUpdate struct {
	config
	hooks     []Hook
	mutation  *OrderMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the OrderUpdate builder.
func (ou *OrderUpdate) Where(ps ...predicate.Order) *OrderUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetCreatedAt sets the "created_at" field.
func (ou *OrderUpdate) SetCreatedAt(u uint32) *OrderUpdate {
	ou.mutation.ResetCreatedAt()
	ou.mutation.SetCreatedAt(u)
	return ou
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableCreatedAt(u *uint32) *OrderUpdate {
	if u != nil {
		ou.SetCreatedAt(*u)
	}
	return ou
}

// AddCreatedAt adds u to the "created_at" field.
func (ou *OrderUpdate) AddCreatedAt(u int32) *OrderUpdate {
	ou.mutation.AddCreatedAt(u)
	return ou
}

// SetUpdatedAt sets the "updated_at" field.
func (ou *OrderUpdate) SetUpdatedAt(u uint32) *OrderUpdate {
	ou.mutation.ResetUpdatedAt()
	ou.mutation.SetUpdatedAt(u)
	return ou
}

// AddUpdatedAt adds u to the "updated_at" field.
func (ou *OrderUpdate) AddUpdatedAt(u int32) *OrderUpdate {
	ou.mutation.AddUpdatedAt(u)
	return ou
}

// SetDeletedAt sets the "deleted_at" field.
func (ou *OrderUpdate) SetDeletedAt(u uint32) *OrderUpdate {
	ou.mutation.ResetDeletedAt()
	ou.mutation.SetDeletedAt(u)
	return ou
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableDeletedAt(u *uint32) *OrderUpdate {
	if u != nil {
		ou.SetDeletedAt(*u)
	}
	return ou
}

// AddDeletedAt adds u to the "deleted_at" field.
func (ou *OrderUpdate) AddDeletedAt(u int32) *OrderUpdate {
	ou.mutation.AddDeletedAt(u)
	return ou
}

// SetGoodID sets the "good_id" field.
func (ou *OrderUpdate) SetGoodID(u uuid.UUID) *OrderUpdate {
	ou.mutation.SetGoodID(u)
	return ou
}

// SetAppID sets the "app_id" field.
func (ou *OrderUpdate) SetAppID(u uuid.UUID) *OrderUpdate {
	ou.mutation.SetAppID(u)
	return ou
}

// SetUserID sets the "user_id" field.
func (ou *OrderUpdate) SetUserID(u uuid.UUID) *OrderUpdate {
	ou.mutation.SetUserID(u)
	return ou
}

// SetParentOrderID sets the "parent_order_id" field.
func (ou *OrderUpdate) SetParentOrderID(u uuid.UUID) *OrderUpdate {
	ou.mutation.SetParentOrderID(u)
	return ou
}

// SetNillableParentOrderID sets the "parent_order_id" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableParentOrderID(u *uuid.UUID) *OrderUpdate {
	if u != nil {
		ou.SetParentOrderID(*u)
	}
	return ou
}

// ClearParentOrderID clears the value of the "parent_order_id" field.
func (ou *OrderUpdate) ClearParentOrderID() *OrderUpdate {
	ou.mutation.ClearParentOrderID()
	return ou
}

// SetPayWithParent sets the "pay_with_parent" field.
func (ou *OrderUpdate) SetPayWithParent(b bool) *OrderUpdate {
	ou.mutation.SetPayWithParent(b)
	return ou
}

// SetNillablePayWithParent sets the "pay_with_parent" field if the given value is not nil.
func (ou *OrderUpdate) SetNillablePayWithParent(b *bool) *OrderUpdate {
	if b != nil {
		ou.SetPayWithParent(*b)
	}
	return ou
}

// ClearPayWithParent clears the value of the "pay_with_parent" field.
func (ou *OrderUpdate) ClearPayWithParent() *OrderUpdate {
	ou.mutation.ClearPayWithParent()
	return ou
}

// SetUnits sets the "units" field.
func (ou *OrderUpdate) SetUnits(u uint32) *OrderUpdate {
	ou.mutation.ResetUnits()
	ou.mutation.SetUnits(u)
	return ou
}

// AddUnits adds u to the "units" field.
func (ou *OrderUpdate) AddUnits(u int32) *OrderUpdate {
	ou.mutation.AddUnits(u)
	return ou
}

// SetPromotionID sets the "promotion_id" field.
func (ou *OrderUpdate) SetPromotionID(u uuid.UUID) *OrderUpdate {
	ou.mutation.SetPromotionID(u)
	return ou
}

// SetNillablePromotionID sets the "promotion_id" field if the given value is not nil.
func (ou *OrderUpdate) SetNillablePromotionID(u *uuid.UUID) *OrderUpdate {
	if u != nil {
		ou.SetPromotionID(*u)
	}
	return ou
}

// ClearPromotionID clears the value of the "promotion_id" field.
func (ou *OrderUpdate) ClearPromotionID() *OrderUpdate {
	ou.mutation.ClearPromotionID()
	return ou
}

// SetDiscountCouponID sets the "discount_coupon_id" field.
func (ou *OrderUpdate) SetDiscountCouponID(u uuid.UUID) *OrderUpdate {
	ou.mutation.SetDiscountCouponID(u)
	return ou
}

// SetNillableDiscountCouponID sets the "discount_coupon_id" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableDiscountCouponID(u *uuid.UUID) *OrderUpdate {
	if u != nil {
		ou.SetDiscountCouponID(*u)
	}
	return ou
}

// ClearDiscountCouponID clears the value of the "discount_coupon_id" field.
func (ou *OrderUpdate) ClearDiscountCouponID() *OrderUpdate {
	ou.mutation.ClearDiscountCouponID()
	return ou
}

// SetUserSpecialReductionID sets the "user_special_reduction_id" field.
func (ou *OrderUpdate) SetUserSpecialReductionID(u uuid.UUID) *OrderUpdate {
	ou.mutation.SetUserSpecialReductionID(u)
	return ou
}

// SetNillableUserSpecialReductionID sets the "user_special_reduction_id" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableUserSpecialReductionID(u *uuid.UUID) *OrderUpdate {
	if u != nil {
		ou.SetUserSpecialReductionID(*u)
	}
	return ou
}

// ClearUserSpecialReductionID clears the value of the "user_special_reduction_id" field.
func (ou *OrderUpdate) ClearUserSpecialReductionID() *OrderUpdate {
	ou.mutation.ClearUserSpecialReductionID()
	return ou
}

// SetStartAt sets the "start_at" field.
func (ou *OrderUpdate) SetStartAt(u uint32) *OrderUpdate {
	ou.mutation.ResetStartAt()
	ou.mutation.SetStartAt(u)
	return ou
}

// SetNillableStartAt sets the "start_at" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableStartAt(u *uint32) *OrderUpdate {
	if u != nil {
		ou.SetStartAt(*u)
	}
	return ou
}

// AddStartAt adds u to the "start_at" field.
func (ou *OrderUpdate) AddStartAt(u int32) *OrderUpdate {
	ou.mutation.AddStartAt(u)
	return ou
}

// ClearStartAt clears the value of the "start_at" field.
func (ou *OrderUpdate) ClearStartAt() *OrderUpdate {
	ou.mutation.ClearStartAt()
	return ou
}

// SetEndAt sets the "end_at" field.
func (ou *OrderUpdate) SetEndAt(u uint32) *OrderUpdate {
	ou.mutation.ResetEndAt()
	ou.mutation.SetEndAt(u)
	return ou
}

// SetNillableEndAt sets the "end_at" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableEndAt(u *uint32) *OrderUpdate {
	if u != nil {
		ou.SetEndAt(*u)
	}
	return ou
}

// AddEndAt adds u to the "end_at" field.
func (ou *OrderUpdate) AddEndAt(u int32) *OrderUpdate {
	ou.mutation.AddEndAt(u)
	return ou
}

// ClearEndAt clears the value of the "end_at" field.
func (ou *OrderUpdate) ClearEndAt() *OrderUpdate {
	ou.mutation.ClearEndAt()
	return ou
}

// SetFixAmountCouponID sets the "fix_amount_coupon_id" field.
func (ou *OrderUpdate) SetFixAmountCouponID(u uuid.UUID) *OrderUpdate {
	ou.mutation.SetFixAmountCouponID(u)
	return ou
}

// SetNillableFixAmountCouponID sets the "fix_amount_coupon_id" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableFixAmountCouponID(u *uuid.UUID) *OrderUpdate {
	if u != nil {
		ou.SetFixAmountCouponID(*u)
	}
	return ou
}

// ClearFixAmountCouponID clears the value of the "fix_amount_coupon_id" field.
func (ou *OrderUpdate) ClearFixAmountCouponID() *OrderUpdate {
	ou.mutation.ClearFixAmountCouponID()
	return ou
}

// SetType sets the "type" field.
func (ou *OrderUpdate) SetType(s string) *OrderUpdate {
	ou.mutation.SetType(s)
	return ou
}

// SetNillableType sets the "type" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableType(s *string) *OrderUpdate {
	if s != nil {
		ou.SetType(*s)
	}
	return ou
}

// ClearType clears the value of the "type" field.
func (ou *OrderUpdate) ClearType() *OrderUpdate {
	ou.mutation.ClearType()
	return ou
}

// SetState sets the "state" field.
func (ou *OrderUpdate) SetState(s string) *OrderUpdate {
	ou.mutation.SetState(s)
	return ou
}

// SetNillableState sets the "state" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableState(s *string) *OrderUpdate {
	if s != nil {
		ou.SetState(*s)
	}
	return ou
}

// ClearState clears the value of the "state" field.
func (ou *OrderUpdate) ClearState() *OrderUpdate {
	ou.mutation.ClearState()
	return ou
}

// SetCouponIds sets the "coupon_ids" field.
func (ou *OrderUpdate) SetCouponIds(s []string) *OrderUpdate {
	ou.mutation.SetCouponIds(s)
	return ou
}

// ClearCouponIds clears the value of the "coupon_ids" field.
func (ou *OrderUpdate) ClearCouponIds() *OrderUpdate {
	ou.mutation.ClearCouponIds()
	return ou
}

// Mutation returns the OrderMutation object of the builder.
func (ou *OrderUpdate) Mutation() *OrderMutation {
	return ou.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OrderUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := ou.defaults(); err != nil {
		return 0, err
	}
	if len(ou.hooks) == 0 {
		affected, err = ou.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ou.mutation = mutation
			affected, err = ou.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ou.hooks) - 1; i >= 0; i-- {
			if ou.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ou.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ou.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OrderUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OrderUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OrderUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ou *OrderUpdate) defaults() error {
	if _, ok := ou.mutation.UpdatedAt(); !ok {
		if order.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized order.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := order.UpdateDefaultUpdatedAt()
		ou.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ou *OrderUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OrderUpdate {
	ou.modifiers = append(ou.modifiers, modifiers...)
	return ou
}

func (ou *OrderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   order.Table,
			Columns: order.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: order.FieldID,
			},
		},
	}
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldCreatedAt,
		})
	}
	if value, ok := ou.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldCreatedAt,
		})
	}
	if value, ok := ou.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldUpdatedAt,
		})
	}
	if value, ok := ou.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldUpdatedAt,
		})
	}
	if value, ok := ou.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldDeletedAt,
		})
	}
	if value, ok := ou.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldDeletedAt,
		})
	}
	if value, ok := ou.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldGoodID,
		})
	}
	if value, ok := ou.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldAppID,
		})
	}
	if value, ok := ou.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldUserID,
		})
	}
	if value, ok := ou.mutation.ParentOrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldParentOrderID,
		})
	}
	if ou.mutation.ParentOrderIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: order.FieldParentOrderID,
		})
	}
	if value, ok := ou.mutation.PayWithParent(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: order.FieldPayWithParent,
		})
	}
	if ou.mutation.PayWithParentCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: order.FieldPayWithParent,
		})
	}
	if value, ok := ou.mutation.Units(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldUnits,
		})
	}
	if value, ok := ou.mutation.AddedUnits(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldUnits,
		})
	}
	if value, ok := ou.mutation.PromotionID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldPromotionID,
		})
	}
	if ou.mutation.PromotionIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: order.FieldPromotionID,
		})
	}
	if value, ok := ou.mutation.DiscountCouponID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldDiscountCouponID,
		})
	}
	if ou.mutation.DiscountCouponIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: order.FieldDiscountCouponID,
		})
	}
	if value, ok := ou.mutation.UserSpecialReductionID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldUserSpecialReductionID,
		})
	}
	if ou.mutation.UserSpecialReductionIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: order.FieldUserSpecialReductionID,
		})
	}
	if value, ok := ou.mutation.StartAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldStartAt,
		})
	}
	if value, ok := ou.mutation.AddedStartAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldStartAt,
		})
	}
	if ou.mutation.StartAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: order.FieldStartAt,
		})
	}
	if value, ok := ou.mutation.EndAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldEndAt,
		})
	}
	if value, ok := ou.mutation.AddedEndAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldEndAt,
		})
	}
	if ou.mutation.EndAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: order.FieldEndAt,
		})
	}
	if value, ok := ou.mutation.FixAmountCouponID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldFixAmountCouponID,
		})
	}
	if ou.mutation.FixAmountCouponIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: order.FieldFixAmountCouponID,
		})
	}
	if value, ok := ou.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: order.FieldType,
		})
	}
	if ou.mutation.TypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: order.FieldType,
		})
	}
	if value, ok := ou.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: order.FieldState,
		})
	}
	if ou.mutation.StateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: order.FieldState,
		})
	}
	if value, ok := ou.mutation.CouponIds(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: order.FieldCouponIds,
		})
	}
	if ou.mutation.CouponIdsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: order.FieldCouponIds,
		})
	}
	_spec.Modifiers = ou.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// OrderUpdateOne is the builder for updating a single Order entity.
type OrderUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *OrderMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (ouo *OrderUpdateOne) SetCreatedAt(u uint32) *OrderUpdateOne {
	ouo.mutation.ResetCreatedAt()
	ouo.mutation.SetCreatedAt(u)
	return ouo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableCreatedAt(u *uint32) *OrderUpdateOne {
	if u != nil {
		ouo.SetCreatedAt(*u)
	}
	return ouo
}

// AddCreatedAt adds u to the "created_at" field.
func (ouo *OrderUpdateOne) AddCreatedAt(u int32) *OrderUpdateOne {
	ouo.mutation.AddCreatedAt(u)
	return ouo
}

// SetUpdatedAt sets the "updated_at" field.
func (ouo *OrderUpdateOne) SetUpdatedAt(u uint32) *OrderUpdateOne {
	ouo.mutation.ResetUpdatedAt()
	ouo.mutation.SetUpdatedAt(u)
	return ouo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (ouo *OrderUpdateOne) AddUpdatedAt(u int32) *OrderUpdateOne {
	ouo.mutation.AddUpdatedAt(u)
	return ouo
}

// SetDeletedAt sets the "deleted_at" field.
func (ouo *OrderUpdateOne) SetDeletedAt(u uint32) *OrderUpdateOne {
	ouo.mutation.ResetDeletedAt()
	ouo.mutation.SetDeletedAt(u)
	return ouo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableDeletedAt(u *uint32) *OrderUpdateOne {
	if u != nil {
		ouo.SetDeletedAt(*u)
	}
	return ouo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (ouo *OrderUpdateOne) AddDeletedAt(u int32) *OrderUpdateOne {
	ouo.mutation.AddDeletedAt(u)
	return ouo
}

// SetGoodID sets the "good_id" field.
func (ouo *OrderUpdateOne) SetGoodID(u uuid.UUID) *OrderUpdateOne {
	ouo.mutation.SetGoodID(u)
	return ouo
}

// SetAppID sets the "app_id" field.
func (ouo *OrderUpdateOne) SetAppID(u uuid.UUID) *OrderUpdateOne {
	ouo.mutation.SetAppID(u)
	return ouo
}

// SetUserID sets the "user_id" field.
func (ouo *OrderUpdateOne) SetUserID(u uuid.UUID) *OrderUpdateOne {
	ouo.mutation.SetUserID(u)
	return ouo
}

// SetParentOrderID sets the "parent_order_id" field.
func (ouo *OrderUpdateOne) SetParentOrderID(u uuid.UUID) *OrderUpdateOne {
	ouo.mutation.SetParentOrderID(u)
	return ouo
}

// SetNillableParentOrderID sets the "parent_order_id" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableParentOrderID(u *uuid.UUID) *OrderUpdateOne {
	if u != nil {
		ouo.SetParentOrderID(*u)
	}
	return ouo
}

// ClearParentOrderID clears the value of the "parent_order_id" field.
func (ouo *OrderUpdateOne) ClearParentOrderID() *OrderUpdateOne {
	ouo.mutation.ClearParentOrderID()
	return ouo
}

// SetPayWithParent sets the "pay_with_parent" field.
func (ouo *OrderUpdateOne) SetPayWithParent(b bool) *OrderUpdateOne {
	ouo.mutation.SetPayWithParent(b)
	return ouo
}

// SetNillablePayWithParent sets the "pay_with_parent" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillablePayWithParent(b *bool) *OrderUpdateOne {
	if b != nil {
		ouo.SetPayWithParent(*b)
	}
	return ouo
}

// ClearPayWithParent clears the value of the "pay_with_parent" field.
func (ouo *OrderUpdateOne) ClearPayWithParent() *OrderUpdateOne {
	ouo.mutation.ClearPayWithParent()
	return ouo
}

// SetUnits sets the "units" field.
func (ouo *OrderUpdateOne) SetUnits(u uint32) *OrderUpdateOne {
	ouo.mutation.ResetUnits()
	ouo.mutation.SetUnits(u)
	return ouo
}

// AddUnits adds u to the "units" field.
func (ouo *OrderUpdateOne) AddUnits(u int32) *OrderUpdateOne {
	ouo.mutation.AddUnits(u)
	return ouo
}

// SetPromotionID sets the "promotion_id" field.
func (ouo *OrderUpdateOne) SetPromotionID(u uuid.UUID) *OrderUpdateOne {
	ouo.mutation.SetPromotionID(u)
	return ouo
}

// SetNillablePromotionID sets the "promotion_id" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillablePromotionID(u *uuid.UUID) *OrderUpdateOne {
	if u != nil {
		ouo.SetPromotionID(*u)
	}
	return ouo
}

// ClearPromotionID clears the value of the "promotion_id" field.
func (ouo *OrderUpdateOne) ClearPromotionID() *OrderUpdateOne {
	ouo.mutation.ClearPromotionID()
	return ouo
}

// SetDiscountCouponID sets the "discount_coupon_id" field.
func (ouo *OrderUpdateOne) SetDiscountCouponID(u uuid.UUID) *OrderUpdateOne {
	ouo.mutation.SetDiscountCouponID(u)
	return ouo
}

// SetNillableDiscountCouponID sets the "discount_coupon_id" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableDiscountCouponID(u *uuid.UUID) *OrderUpdateOne {
	if u != nil {
		ouo.SetDiscountCouponID(*u)
	}
	return ouo
}

// ClearDiscountCouponID clears the value of the "discount_coupon_id" field.
func (ouo *OrderUpdateOne) ClearDiscountCouponID() *OrderUpdateOne {
	ouo.mutation.ClearDiscountCouponID()
	return ouo
}

// SetUserSpecialReductionID sets the "user_special_reduction_id" field.
func (ouo *OrderUpdateOne) SetUserSpecialReductionID(u uuid.UUID) *OrderUpdateOne {
	ouo.mutation.SetUserSpecialReductionID(u)
	return ouo
}

// SetNillableUserSpecialReductionID sets the "user_special_reduction_id" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableUserSpecialReductionID(u *uuid.UUID) *OrderUpdateOne {
	if u != nil {
		ouo.SetUserSpecialReductionID(*u)
	}
	return ouo
}

// ClearUserSpecialReductionID clears the value of the "user_special_reduction_id" field.
func (ouo *OrderUpdateOne) ClearUserSpecialReductionID() *OrderUpdateOne {
	ouo.mutation.ClearUserSpecialReductionID()
	return ouo
}

// SetStartAt sets the "start_at" field.
func (ouo *OrderUpdateOne) SetStartAt(u uint32) *OrderUpdateOne {
	ouo.mutation.ResetStartAt()
	ouo.mutation.SetStartAt(u)
	return ouo
}

// SetNillableStartAt sets the "start_at" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableStartAt(u *uint32) *OrderUpdateOne {
	if u != nil {
		ouo.SetStartAt(*u)
	}
	return ouo
}

// AddStartAt adds u to the "start_at" field.
func (ouo *OrderUpdateOne) AddStartAt(u int32) *OrderUpdateOne {
	ouo.mutation.AddStartAt(u)
	return ouo
}

// ClearStartAt clears the value of the "start_at" field.
func (ouo *OrderUpdateOne) ClearStartAt() *OrderUpdateOne {
	ouo.mutation.ClearStartAt()
	return ouo
}

// SetEndAt sets the "end_at" field.
func (ouo *OrderUpdateOne) SetEndAt(u uint32) *OrderUpdateOne {
	ouo.mutation.ResetEndAt()
	ouo.mutation.SetEndAt(u)
	return ouo
}

// SetNillableEndAt sets the "end_at" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableEndAt(u *uint32) *OrderUpdateOne {
	if u != nil {
		ouo.SetEndAt(*u)
	}
	return ouo
}

// AddEndAt adds u to the "end_at" field.
func (ouo *OrderUpdateOne) AddEndAt(u int32) *OrderUpdateOne {
	ouo.mutation.AddEndAt(u)
	return ouo
}

// ClearEndAt clears the value of the "end_at" field.
func (ouo *OrderUpdateOne) ClearEndAt() *OrderUpdateOne {
	ouo.mutation.ClearEndAt()
	return ouo
}

// SetFixAmountCouponID sets the "fix_amount_coupon_id" field.
func (ouo *OrderUpdateOne) SetFixAmountCouponID(u uuid.UUID) *OrderUpdateOne {
	ouo.mutation.SetFixAmountCouponID(u)
	return ouo
}

// SetNillableFixAmountCouponID sets the "fix_amount_coupon_id" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableFixAmountCouponID(u *uuid.UUID) *OrderUpdateOne {
	if u != nil {
		ouo.SetFixAmountCouponID(*u)
	}
	return ouo
}

// ClearFixAmountCouponID clears the value of the "fix_amount_coupon_id" field.
func (ouo *OrderUpdateOne) ClearFixAmountCouponID() *OrderUpdateOne {
	ouo.mutation.ClearFixAmountCouponID()
	return ouo
}

// SetType sets the "type" field.
func (ouo *OrderUpdateOne) SetType(s string) *OrderUpdateOne {
	ouo.mutation.SetType(s)
	return ouo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableType(s *string) *OrderUpdateOne {
	if s != nil {
		ouo.SetType(*s)
	}
	return ouo
}

// ClearType clears the value of the "type" field.
func (ouo *OrderUpdateOne) ClearType() *OrderUpdateOne {
	ouo.mutation.ClearType()
	return ouo
}

// SetState sets the "state" field.
func (ouo *OrderUpdateOne) SetState(s string) *OrderUpdateOne {
	ouo.mutation.SetState(s)
	return ouo
}

// SetNillableState sets the "state" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableState(s *string) *OrderUpdateOne {
	if s != nil {
		ouo.SetState(*s)
	}
	return ouo
}

// ClearState clears the value of the "state" field.
func (ouo *OrderUpdateOne) ClearState() *OrderUpdateOne {
	ouo.mutation.ClearState()
	return ouo
}

// SetCouponIds sets the "coupon_ids" field.
func (ouo *OrderUpdateOne) SetCouponIds(s []string) *OrderUpdateOne {
	ouo.mutation.SetCouponIds(s)
	return ouo
}

// ClearCouponIds clears the value of the "coupon_ids" field.
func (ouo *OrderUpdateOne) ClearCouponIds() *OrderUpdateOne {
	ouo.mutation.ClearCouponIds()
	return ouo
}

// Mutation returns the OrderMutation object of the builder.
func (ouo *OrderUpdateOne) Mutation() *OrderMutation {
	return ouo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OrderUpdateOne) Select(field string, fields ...string) *OrderUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Order entity.
func (ouo *OrderUpdateOne) Save(ctx context.Context) (*Order, error) {
	var (
		err  error
		node *Order
	)
	if err := ouo.defaults(); err != nil {
		return nil, err
	}
	if len(ouo.hooks) == 0 {
		node, err = ouo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ouo.mutation = mutation
			node, err = ouo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ouo.hooks) - 1; i >= 0; i-- {
			if ouo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ouo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ouo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Order)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from OrderMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OrderUpdateOne) SaveX(ctx context.Context) *Order {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OrderUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OrderUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ouo *OrderUpdateOne) defaults() error {
	if _, ok := ouo.mutation.UpdatedAt(); !ok {
		if order.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized order.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := order.UpdateDefaultUpdatedAt()
		ouo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ouo *OrderUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OrderUpdateOne {
	ouo.modifiers = append(ouo.modifiers, modifiers...)
	return ouo
}

func (ouo *OrderUpdateOne) sqlSave(ctx context.Context) (_node *Order, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   order.Table,
			Columns: order.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: order.FieldID,
			},
		},
	}
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Order.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, order.FieldID)
		for _, f := range fields {
			if !order.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != order.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldCreatedAt,
		})
	}
	if value, ok := ouo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldCreatedAt,
		})
	}
	if value, ok := ouo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldUpdatedAt,
		})
	}
	if value, ok := ouo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldUpdatedAt,
		})
	}
	if value, ok := ouo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldDeletedAt,
		})
	}
	if value, ok := ouo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldDeletedAt,
		})
	}
	if value, ok := ouo.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldGoodID,
		})
	}
	if value, ok := ouo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldAppID,
		})
	}
	if value, ok := ouo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldUserID,
		})
	}
	if value, ok := ouo.mutation.ParentOrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldParentOrderID,
		})
	}
	if ouo.mutation.ParentOrderIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: order.FieldParentOrderID,
		})
	}
	if value, ok := ouo.mutation.PayWithParent(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: order.FieldPayWithParent,
		})
	}
	if ouo.mutation.PayWithParentCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: order.FieldPayWithParent,
		})
	}
	if value, ok := ouo.mutation.Units(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldUnits,
		})
	}
	if value, ok := ouo.mutation.AddedUnits(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldUnits,
		})
	}
	if value, ok := ouo.mutation.PromotionID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldPromotionID,
		})
	}
	if ouo.mutation.PromotionIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: order.FieldPromotionID,
		})
	}
	if value, ok := ouo.mutation.DiscountCouponID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldDiscountCouponID,
		})
	}
	if ouo.mutation.DiscountCouponIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: order.FieldDiscountCouponID,
		})
	}
	if value, ok := ouo.mutation.UserSpecialReductionID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldUserSpecialReductionID,
		})
	}
	if ouo.mutation.UserSpecialReductionIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: order.FieldUserSpecialReductionID,
		})
	}
	if value, ok := ouo.mutation.StartAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldStartAt,
		})
	}
	if value, ok := ouo.mutation.AddedStartAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldStartAt,
		})
	}
	if ouo.mutation.StartAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: order.FieldStartAt,
		})
	}
	if value, ok := ouo.mutation.EndAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldEndAt,
		})
	}
	if value, ok := ouo.mutation.AddedEndAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldEndAt,
		})
	}
	if ouo.mutation.EndAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: order.FieldEndAt,
		})
	}
	if value, ok := ouo.mutation.FixAmountCouponID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldFixAmountCouponID,
		})
	}
	if ouo.mutation.FixAmountCouponIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: order.FieldFixAmountCouponID,
		})
	}
	if value, ok := ouo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: order.FieldType,
		})
	}
	if ouo.mutation.TypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: order.FieldType,
		})
	}
	if value, ok := ouo.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: order.FieldState,
		})
	}
	if ouo.mutation.StateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: order.FieldState,
		})
	}
	if value, ok := ouo.mutation.CouponIds(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: order.FieldCouponIds,
		})
	}
	if ouo.mutation.CouponIdsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: order.FieldCouponIds,
		})
	}
	_spec.Modifiers = ouo.modifiers
	_node = &Order{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

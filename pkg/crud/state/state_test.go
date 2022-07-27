package state

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/order-manager/pkg/db/ent"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	valuedef "github.com/NpoolPlatform/message/npool"
	npool "github.com/NpoolPlatform/message/npool/order/mgr/v1/state"
	testinit "github.com/NpoolPlatform/order-manager/pkg/testinit"
	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var entity = ent.State{
	ID:      uuid.New(),
	OrderID: uuid.New(),
	State:   npool.EState_WaitPayment.String(),
}

var (
	id      = entity.ID.String()
	orderID = entity.OrderID.String()
	status  = npool.EState(npool.EState_value[entity.State])

	req = npool.StateReq{
		ID:      &id,
		OrderID: &orderID,
		State:   &status,
	}
)

var info *ent.State

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &req)
	if assert.Nil(t, err) {
		entity.UpdatedAt = info.UpdatedAt
		entity.CreatedAt = info.CreatedAt
		assert.Equal(t, info.String(), entity.String())
	}
}

func createBulk(t *testing.T) {
	entities := []*ent.State{
		{
			ID:      uuid.New(),
			OrderID: uuid.New(),
			State:   npool.EState_Paid.String(),
		},
		{
			ID:      uuid.New(),
			OrderID: uuid.New(),
			State:   npool.EState_Canceled.String(),
		},
	}

	reqs := []*npool.StateReq{}
	for _, _entity := range entities {
		_id := _entity.ID.String()
		_orderID := _entity.OrderID.String()
		_state := npool.EState(npool.EState_value[_entity.State])

		reqs = append(reqs, &npool.StateReq{
			ID:      &_id,
			OrderID: &_orderID,
			State:   &_state,
		})
	}
	infos, err := CreateBulk(context.Background(), reqs)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), entity.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), entity.String())
	}
}

func rows(t *testing.T) {
	infos, total, err := Rows(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		}, 0, 0)
	if assert.Nil(t, err) {
		if assert.Equal(t, total, 1) {
			assert.Equal(t, infos[0].String(), entity.String())
		}
	}
}

func rowOnly(t *testing.T) {
	var err error
	info, err = RowOnly(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), entity.String())
	}
}

func count(t *testing.T) {
	count, err := Count(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, count, uint32(1))
	}
}

func exist(t *testing.T) {
	exist, err := Exist(context.Background(), entity.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existConds(t *testing.T) {
	exist, err := ExistConds(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteA(t *testing.T) {
	info, err := Delete(context.Background(), entity.ID)
	if assert.Nil(t, err) {
		entity.DeletedAt = info.DeletedAt
		assert.Equal(t, info.String(), entity.String())
	}
}

func TestState(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("create", create)
	if false {
		t.Run("createBulk", createBulk)
		t.Run("row", row)
		t.Run("rows", rows)
		t.Run("rowOnly", rowOnly)
		t.Run("exist", exist)
		t.Run("existConds", existConds)
		t.Run("count", count)
		t.Run("delete", deleteA)
	}
}

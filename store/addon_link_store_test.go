package store

import (
	"testing"
	"time"

	"github.com/abdoroot/subscription/types"
	"github.com/abdoroot/subscription/util"
	"github.com/stretchr/testify/assert"
)

func TestCreateAddonsLink(t *testing.T) {
	db, err := util.ConnectToPq()
	defer func() {
		assert.Nil(t, db.Close())
	}()
	assert.NoError(t, err)
	store := NewAddonLinkStore(db)

	param := types.AddonLink{
		AddonId: 100,
		PlanId:  50,
	}
	cl, err := store.CreateAddonsLink(param)
	assert.NoError(t, err)

	defer func() {
		if cl.ID > 0 {
			err := store.DeleteAddonsLinkById(cl.ID)
			assert.NoError(t, err)
		}
	}()

	assert.Greater(t, cl.ID, 0)
	assert.Greater(t, cl.CreatedAt, time.Time{})
}

func TestGetAddonsLinkById(t *testing.T) {
	db, err := util.ConnectToPq()
	defer func() {
		assert.Nil(t, db.Close())
	}()
	assert.NoError(t, err)
	store := NewAddonLinkStore(db)
	cl, err := createAddonLink(store)
	assert.NoError(t, err)

	defer func() {
		if cl.ID > 0 {
			err := store.DeleteAddonsLinkById(cl.ID)
			assert.NoError(t, err)
		}
	}()

	got, err := store.GetAddonsLinkById(cl.ID)
	assert.NoError(t, err)

	assert.Greater(t, got.ID, 0)
	assert.Greater(t, got.CreatedAt, time.Time{})
}

// test helper func
func createAddonLink(store *addonLinkStore) (types.AddonLink, error) {
	param := types.AddonLink{
		AddonId: 100,
		PlanId:  50,
	}
	cl, err := store.CreateAddonsLink(param)
	return cl, err
}

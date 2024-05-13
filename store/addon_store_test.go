package store

import (
	"testing"

	"github.com/abdoroot/subscription/types"
	"github.com/abdoroot/subscription/util"
	"github.com/stretchr/testify/assert"
)

func TestCreateAddon(t *testing.T) {
	s, addon, err := CreateAddon()
	defer func() {
		if addon.ID > 0 {
			assert.Nil(t, s.DeleteAddonById(addon.ID))
		}
	}()
	assert.Nil(t, err)
	assert.Greater(t, addon.ID, 0)
}

func TestUpdateAddon(t *testing.T) {
	s, addon, err := CreateAddon()
	assert.Nil(t, err)
	defer func() {
		if addon.ID > 0 {
			assert.Nil(t, s.DeleteAddonById(addon.ID))
		}
	}()

	params := types.Addon{
		Name:            "New Addon 2",
		Description:     "This is a new addon. updated",
		DisplayInWidget: true,
		PricingModelID:  2,
		UnitID:          3,
		BillingType:     "once",
		BillingCycle:    "month",
		Price:           8.99,
		LinkedToPlans:   "all_plans",
	}
	if err := s.UpdateAddonById(params, addon.ID); err != nil {
		t.Error("error UpdateAddonById ", err)
	}
}

func TestGetAddonById(t *testing.T) {
	s, addon, err := CreateAddon()
	assert.Nil(t, err)
	defer func() {
		if addon.ID > 0 {
			assert.Nil(t, s.DeleteAddonById(addon.ID))
		}
	}()
	got, err := s.GetAddonById(addon.ID)
	assert.Nil(t, err)
	assert.Greater(t, addon.ID, 0)
	assert.Equal(t, got.ID, addon.ID)
	assert.Equal(t, got.Name, addon.Name)
}

func TestGetAddonsByCompanyId(t *testing.T) {
	s, addon, err := CreateAddon()
	assert.Nil(t, err)
	defer func() {
		if addon.ID > 0 {
			assert.Nil(t, s.DeleteAddonById(addon.ID))
		}
	}()
	got, err := s.GetAddonByCompanyId(addon.ID)
	assert.Nil(t, err)
	assert.Greater(t, len(got), 0)
}

func TestGetAllAddons(t *testing.T) {
	s, addon, err := CreateAddon()
	assert.Nil(t, err)
	defer func() {
		if addon.ID > 0 {
			assert.Nil(t, s.DeleteAddonById(addon.ID))
		}
	}()
	got, err := s.GetALl()
	assert.Nil(t, err)
	assert.Greater(t, len(got), 0)
}

// testing helper func
func CreateAddon() (*addonStore, types.Addon, error) {
	db, err := util.ConnectToPq()
	if err != nil {
		return nil, types.Addon{}, err
	}
	ls := NewAddonLinkStore(db)
	s := NewAddonStore(db, ls)

	param := types.Addon{
		CompanyID:       1,
		Name:            "New Addon",
		Description:     "This is a new addon.",
		DisplayInWidget: true,
		PricingModelID:  1,
		UnitID:          1,
		BillingType:     "once",
		BillingCycle:    "month",
		Price:           9.99,
		LinkedToPlans:   "all_plans",
	}

	addon, err := s.CreateAddon(param)
	if err != nil {
		return nil, types.Addon{}, err
	}

	return s, addon, err
}

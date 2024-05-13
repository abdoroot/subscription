package store

import (
	"testing"
	"time"

	"github.com/abdoroot/subscription/types"
	"github.com/abdoroot/subscription/util"
	"github.com/stretchr/testify/assert"
)

func TestCreateCoupon(t *testing.T) {
	store, c, err := CreateCoupon()
	if err != nil {
		t.Error("error CreateCoupon ", err)
	}
	defer func() {
		if c.ID > 0 {
			store.DeleteCouponById(c.ID)
		}
	}()

	assert.Greater(t, c.ID, 0)
	assert.Greater(t, c.CreatedAt, time.Time{})
}

func TestGetAllAndGetByCompanyIdCoupon(t *testing.T) {
	store, c, err := CreateCoupon()
	if err != nil {
		t.Error("error CreateCoupon ", err)
	}
	defer func() {
		if c.ID > 0 {
			store.DeleteCouponById(c.ID)
			store.db.Close()
		}
	}()

	assert.Greater(t, c.ID, 0)
	t.Run("GetCouponByCompanyId test", func(t *testing.T) {
		got, err := store.GetCouponById(c.ID)
		assert.Nil(t, err)
		assert.Greater(t, got.ID, 0)
	})
	t.Run("GetAll test", func(t *testing.T) {
		got, err := store.GetCouponByCompanyId(c.CompanyID)
		assert.Nil(t, err)
		assert.Greater(t, len(got), 0)
	})

}

func TestUpdateAndGetCoupon(t *testing.T) {
	store, c, err := CreateCoupon()
	if err != nil {
		t.Error("error CreateCoupon ", err)
	}
	defer func() {
		if c.ID > 0 {
			store.DeleteCouponById(c.ID)
			store.db.Close()
		}
	}()

	param := types.Coupon{
		CompanyID:       1,
		Name:            "updated Coupon",
		DiscountValue:   10.99,
		DiscountType:    "percentage",
		RedemptionType:  "once",
		MaxRedemptions:  100,
		RedemptionCount: 2,
		LinkedToPlans:   "all",
		LinkedToAddons:  "all",
		ExpireDate:      time.Now().AddDate(0, 2, 0), // Expires in 2 month
	}

	t.Run("UpdateCouponById test", func(t *testing.T) {
		err = store.UpdateCouponById(param, c.ID)
		assert.Nil(t, err)
	})
	t.Run("getCouponById test", func(t *testing.T) {
		got, err := store.GetCouponById(c.ID)
		assert.Nil(t, err)
		assert.Equal(t, got.ID, c.ID)
		assert.Equal(t, got.Name, param.Name)
		assert.Equal(t, got.RedemptionCount, param.RedemptionCount)
	})
}

func CreateCoupon() (*couponStore, types.Coupon, error) {
	db, err := util.ConnectToPq()
	if err != nil {
		return nil, types.Coupon{}, nil
	}
	store := NewCouponStore(db)
	param := types.Coupon{
		CompanyID:       1,
		Name:            "Example Coupon",
		DiscountValue:   10.00,
		DiscountType:    "percentage",
		RedemptionType:  "once",
		MaxRedemptions:  100,
		RedemptionCount: 0,
		LinkedToPlans:   "all",
		LinkedToAddons:  "all",
		ExpireDate:      time.Now().AddDate(0, 1, 0), // Expires in 1 month
	}
	c, err := store.CreateCoupon(param)
	return store, c, err
}

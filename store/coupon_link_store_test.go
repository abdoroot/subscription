package store

import (
	"testing"
	"time"

	"github.com/abdoroot/subscription/types"
	"github.com/abdoroot/subscription/util"
	"github.com/stretchr/testify/assert"
)

func TestCreateCouponLink(t *testing.T) {
	s, cl, err := createCouponLink()
	defer s.DeleteCouponLinkById(cl.ID)
	if err != nil {
		t.Error("error createCouponLink ", err)
	}
	assert.Greater(t, cl.ID, 0)
	assert.Greater(t, cl.CreatedAt, time.Time{})
}

func TestGetCouponLinkById(t *testing.T) {
	s, cl, err := createCouponLink()
	defer s.DeleteCouponLinkById(cl.ID)
	if err != nil {
		t.Error("error createCouponLink ", err)
	}
	got, err := s.GetCouponLinkById(cl.ID)
	if err != nil {
		t.Error("error GetCouponLinkById ", err)
	}
	assert.Greater(t, got.ID, 0)
	assert.Greater(t, got.CreatedAt, time.Time{})
}

// help test fucn
func createCouponLink() (*couponLinkStore, types.CouponLink, error) {
	db, err := util.ConnectToPq()
	if err != nil {
		return nil, types.CouponLink{}, nil
	}
	store := NewCouponLinkStore(db)
	param := types.CouponLink{
		CouponId: 10,
		LinkType: "plan",
		LinkToId: 5,
	}
	cl, err := store.CreateCouponLink(param)
	if err != nil {
		return nil, types.CouponLink{}, err
	}
	return store, cl, nil
}

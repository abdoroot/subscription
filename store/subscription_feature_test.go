package store

import (
	"testing"

	"github.com/abdoroot/subscription/types"
	"github.com/abdoroot/subscription/util"
	"github.com/stretchr/testify/assert"
)

func TestCreateFeature(t *testing.T) {
	sfs, f, err := createSubFeature()
	_ = sfs
	if err != nil {
		t.Error("error creating createSubFeature ", err)
	}
	defer sfs.DeleteFeatureById(f.Id)
	assert.Greater(t, f.Id, 0)
}

func TestUpdateFeature(t *testing.T) {
	sfs, f, err := createSubFeature()
	if err != nil {
		t.Error("error creating createSubFeature ", err)
	}
	defer sfs.DeleteFeatureById(f.Id)
	param := types.SubscriptionFeature{
		Name: "updated feature",
	}
	if err := sfs.UpdateFeatureById(param, f.Id); err != nil {
		t.Error("error UpdateFeatureById ", err)
	}
	assert.Greater(t, f.Id, 0)
}

func TestDeleteFeature(t *testing.T) {
	sfs, f, err := createSubFeature()
	if err != nil {
		t.Error("error creating createSubFeature ", err)
	}
	defer sfs.DeleteFeatureById(f.Id)
	if err := sfs.DeleteFeatureById(f.Id); err != nil {
		t.Error("error UpdateFeatureById ", err)
	}
	assert.Greater(t, f.Id, 0)
}

func TestGetFeature(t *testing.T) {
	sfs, f, err := createSubFeature()
	if err != nil {
		t.Error("error creating createSubFeature ", err)
	}
	defer sfs.DeleteFeatureById(f.Id)
	t.Run("GetFeatureById", func(t *testing.T) {
		feature, err := sfs.GetFeatureById(f.Id)
		if err != nil {
			t.Error("error GetFeatureById ", err)
		}
		assert.Greater(t, feature.Id, 0)
	})
}

// test helper func
func createSubFeature() (*subscriptionFeatureStore, types.SubscriptionFeature, error) {
	db, err := util.ConnectToPq()
	if err != nil {
		return nil, types.SubscriptionFeature{}, err
	}
	sfs := NewSubscriptionFeatureStore(db)
	param := types.SubscriptionFeature{
		SubscriptionPlanId: 5,
		Name:               "test feature",
	}
	f, err := sfs.CreateFeature(param)
	if err != nil {
		return nil, types.SubscriptionFeature{}, err
	}
	return sfs, f, nil
}

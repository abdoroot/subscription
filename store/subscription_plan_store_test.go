package store

import (
	"testing"

	"github.com/abdoroot/subscription/types"
	"github.com/abdoroot/subscription/util"
	"github.com/stretchr/testify/assert"
)

func TestCreateSubscriptionPlan(t *testing.T) {
	s, plan, err := createSubPlan()
	if err != nil {
		t.Error("error CreateSubscriptionPlan", err)
	}
	defer s.DeletePlanById(plan.Id)
	assert.Greater(t, plan.Id, 0)
}

func TestDeletePlanById(t *testing.T) {
	s, plan, err := createSubPlan()
	if err != nil {
		t.Error("error CreateSubscriptionPlan", err)
	}
	err = s.DeletePlanById(plan.Id)
	if err != nil {
		t.Error("error DeletePlanById", err)
	}
}

func TestGetAllPlans(t *testing.T) {
	s, p, err := createSubPlan()
	if err != nil {
		t.Error("error CreateSubscriptionPlan", err)
	}
	defer s.DeletePlanById(p.Id)
	if err != nil {
		t.Error("error CreateSubscriptionPlan", err)
	}

	pls, err := s.GetAll()
	if err != nil {
		t.Error("error GetAll", err)
	}
	assert.Greater(t, len(pls), 0)
}

func TestGetPlansByCompanyId(t *testing.T) {
	s, p, err := createSubPlan()
	if err != nil {
		t.Error("error CreateSubscriptionPlan", err)
	}
	defer s.DeletePlanById(p.Id)
	if err != nil {
		t.Error("error GetPlansByCompanyId", err)
	}

	pls, err := s.GetPlansByCompanyId(p.CompanyId)
	if err != nil {
		t.Error("error GetPlansByCompanyId", err)
	}
	assert.Greater(t, len(pls), 0)
}

func TestGetPlanById(t *testing.T) {
	s, p, err := createSubPlan()
	if err != nil {
		t.Error("error CreateSubscriptionPlan", err)
	}
	defer s.DeletePlanById(p.Id)
	if err != nil {
		t.Error("error GetPlansByCompanyId", err)
	}

	pl, err := s.GetPlanById(p.Id)
	if err != nil {
		t.Error("error GetPlanById", err)
	}
	assert.Greater(t, pl.Id, 0)
}

func BenchmarkGetAll(b *testing.B) {
	s, p, err := createSubPlan()
	if err != nil {
		b.Error("error CreateSubscriptionPlan", err)
	}
	defer s.DeletePlanById(p.Id)
	for i := 0; i < b.N; i++ {
		s.GetAll()
	}
}

// help test func
func createSubPlan() (*SubscriptionPlanStore, types.SubscriptionPlan, error) {
	db, err := util.ConnectToPq()
	if err != nil {
		return nil, types.SubscriptionPlan{}, err
	}
	fs := NewSubscriptionFeatureStore(db)
	sps := NewSubscriptionPlanStore(db, fs)
	param := types.SubscriptionPlan{
		CompanyId:         1,
		ProductId:         2,
		Name:              "basic plan",
		Code:              "b-1",
		UnitId:            3,
		IsActive:          true,
		Price:             15.99,
		Type:              "service",
		BillingEvery:      "week",
		BillingEveryCount: 2,
		BillingCycle:      "auto_renewal",
		BillingCycleCount: 0,
		TrialPeriod:       0.00,
		SetupFee:          0.00,
		Description:       "",
		Features: []types.SubscriptionFeature{
			{Name: "unlimited download"},
			{Name: "Seo optmization"},
		},
	}

	pl, err := sps.CreatePlan(param)
	if err != nil {
		return nil, types.SubscriptionPlan{}, err
	}
	return sps, pl, err
}

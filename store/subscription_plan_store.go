package store

import (
	"errors"
	"fmt"
	"time"

	"github.com/abdoroot/subscription/types"
	"github.com/jmoiron/sqlx"
)

// todo
type SubscriptionPlanStore struct {
	db           *sqlx.DB
	featureStore *subscriptionFeatureStore
}

func NewSubscriptionPlanStore(db *sqlx.DB, featureStore *subscriptionFeatureStore) *SubscriptionPlanStore {
	return &SubscriptionPlanStore{
		db:           db,
		featureStore: featureStore,
	}
}

func (s *SubscriptionPlanStore) CreatePlan(param types.SubscriptionPlan) (types.SubscriptionPlan, error) {
	//first add the plan
	param.CreatedAt = time.Now()
	param.UpdatedAt = time.Now()
	if len(param.Features) > 0 {
		param.HaveFeatures = true
	}
	query := `insert into subscription_plans (company_id,product_id,name,code,unit_id,type,is_active,price,billing_every,billing_every_count,billing_cycle,billing_cycle_count,trial_period,setup_fee,description,have_features,created_at,updated_at) 
	values(:company_id,:product_id,:name,:code,:unit_id,:type,:is_active,:price,:billing_every,:billing_every_count,:billing_cycle,:billing_cycle_count,:trial_period,:setup_fee,:description,:have_features,:created_at,:updated_at) returning id;`
	row, err := s.db.NamedQuery(query, param)
	if err != nil {
		return types.SubscriptionPlan{}, err
	}
	if row.Next() {
		if err := row.Scan(&param.Id); err != nil {
			return types.SubscriptionPlan{}, errors.Join(fmt.Errorf("error scanning"), err)
		}
	}

	//second add the features
	if len(param.Features) > 0 {
		for i, f := range param.Features {
			f.SubscriptionPlanId = param.Id
			sf, err := s.featureStore.CreateFeature(f)
			if err != nil {
				return types.SubscriptionPlan{}, err
			}
			param.Features[i] = sf
		}
	}
	return param, nil
}

func (s *SubscriptionPlanStore) UpdatePlanById(param types.SubscriptionPlan, id int) error {
	param.Id = id
	param.UpdatedAt = time.Now()
	query := `update subscription_plans set product_id= :product_id,name= :name,code= :code,unit_id= :unit_id,type= :type,is_active= :is_active, price= :price,billing_every= :billing_every,billing_every_count= :billing_every_count,billing_cycle = :billing_cycle ,billing_cycle_count= :billing_cycle_count,trial_period = :trial_period,setup_fee = :setup_fee,description = :description,have_features = :have_features,updated_at = :updated_at where id= :id;`
	_, err := s.db.NamedExec(query, param)
	if err != nil {
		return err
	}
	return nil
}

func (s *SubscriptionPlanStore) GetPlanById(id int) (types.SubscriptionPlan, error) {
	var plan types.SubscriptionPlan
	query := `select * from subscription_plans where id =$1;`
	if err := s.db.Get(&plan, query, id); err != nil {
		return types.SubscriptionPlan{}, err
	}

	if plan.HaveFeatures {
		pfs, err := s.featureStore.GetFeaturesByPlanId(plan.Id)
		if err != nil {
			return types.SubscriptionPlan{}, err
		}
		plan.Features = pfs
	}

	return plan, nil
}

func (s *SubscriptionPlanStore) GetPlansByCompanyId(id int) ([]types.SubscriptionPlan, error) {
	var plans []types.SubscriptionPlan
	query := `select * from subscription_plans where company_id =$1;`
	if err := s.db.Select(&plans, query, id); err != nil {
		return []types.SubscriptionPlan{}, err
	}

	if len(plans) > 0 {
		for i := 0; i < len(plans); i++ {
			if plans[i].HaveFeatures {
				pfs, err := s.featureStore.GetFeaturesByPlanId(plans[i].Id)
				if err != nil {
					continue
				}
				plans[i].Features = pfs
			}
		}
	}
	return plans, nil
}

func (s *SubscriptionPlanStore) GetAll() ([]types.SubscriptionPlan, error) {
	var plans []types.SubscriptionPlan
	query := `select * from subscription_plans;`
	if err := s.db.Select(&plans, query); err != nil {
		return []types.SubscriptionPlan{}, err
	}

	if len(plans) > 0 {
		for i := 0; i < len(plans); i++ {
			if plans[i].HaveFeatures {
				pfs, err := s.featureStore.GetFeaturesByPlanId(plans[i].Id)
				if err != nil {
					continue
				}
				plans[i].Features = pfs
			}
		}
	}
	return plans, nil
}

func (s *SubscriptionPlanStore) DeletePlanById(id int) error {
	//First delete subscriptionPlanFeature
	err := s.featureStore.DeleteFeaturesByPlandId(id)
	if err != nil {
		return err
	}
	//Secound delete subscription plan
	query := `delete from subscription_plans where id=$1`
	if _, err = s.db.Exec(query, id); err != nil {
		return err
	}
	return nil
}

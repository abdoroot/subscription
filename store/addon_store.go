package store

import (
	"time"

	"github.com/abdoroot/subscription/types"
	"github.com/jmoiron/sqlx"
)

type addonStore struct {
	db        *sqlx.DB
	linkstore *addonLinkStore
}

func NewAddonStore(db *sqlx.DB, linkstore *addonLinkStore) *addonStore {
	return &addonStore{
		db:        db,
		linkstore: linkstore,
	}
}

func (s *addonStore) CreateAddon(param types.Addon) (types.Addon, error) {
	param.CreatedAt = time.Now()
	param.UpdatedAt = time.Now()
	query := `INSERT INTO addons(company_id, name, description, display_in_widget, pricing_model_id, unit_id, billing_type, billing_cycle, price, linked_to_plans, created_at, updated_at) 
			   VALUES (:company_id, :name, :description, :display_in_widget, :pricing_model_id, :unit_id, :billing_type, :billing_cycle, :price, :linked_to_plans, :created_at, :updated_at) 
			   RETURNING id`

	row, err := s.db.NamedQuery(query, param)
	if err != nil {
		return types.Addon{}, err
	}
	if row.Next() {
		if err := row.Scan(&param.ID); err != nil {
			return types.Addon{}, err
		}
	}
	return param, nil
}

func (s *addonStore) UpdateAddonById(param types.Addon, id int) error {
	param.UpdatedAt = time.Now()
	param.ID = id
	query := `UPDATE addons
			  SET name = :name, 
			      description = :description, 
			      display_in_widget = :display_in_widget, 
			      pricing_model_id = :pricing_model_id, 
			      unit_id = :unit_id, 
			      billing_type = :billing_type, 
			      billing_cycle = :billing_cycle, 
			      price = :price, 
			      linked_to_plans = :linked_to_plans, 
			      updated_at = :updated_at
			  WHERE id = :id`

	_, err := s.db.NamedExec(query, param)
	return err
}

func (s *addonStore) GetAddonById(id int) (types.Addon, error) {
	var addon types.Addon
	query := `SELECT * FROM addons WHERE id = $1`
	err := s.db.Get(&addon, query, id)
	return addon, err
}

func (s *addonStore) GetAddonByCompanyId(id int) ([]types.Addon, error) {
	var addons []types.Addon
	query := `SELECT * FROM addons WHERE company_id = $1`
	err := s.db.Select(&addons, query, id)
	return addons, err
}

func (s *addonStore) GetALl() ([]types.Addon, error) {
	var addons []types.Addon
	query := `SELECT * FROM addons`
	err := s.db.Select(&addons, query)
	return addons, err
}

func (s *addonStore) DeleteAddonById(id int) error {
	//First Delete Links
	err := s.linkstore.DeleteAddonLinkByAddonId(id)
	if err != nil {
		return err
	}
	//Second Delete addon
	query := `DELETE FROM addons WHERE id = :id;`
	_, err = s.db.NamedExec(query, map[string]any{"id": id})
	if err != nil {
		return err
	}
	return nil
}

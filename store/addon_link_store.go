package store

import (
	"time"

	"github.com/abdoroot/subscription/types"
	"github.com/jmoiron/sqlx"
)

type addonLinkStore struct {
	db *sqlx.DB
}

func NewAddonLinkStore(db *sqlx.DB) *addonLinkStore {
	return &addonLinkStore{
		db: db,
	}
}

func (s *addonLinkStore) CreateAddonsLink(param types.AddonLink) (types.AddonLink, error) {
	param.CreatedAt = time.Now()
	param.UpdatedAt = time.Now()
	query := `insert into addons_links(addon_id,plan_id,created_at,updated_at) 
	values(:addon_id,:plan_id,:created_at,:updated_at) returning id
	`
	row, err := s.db.NamedQuery(query, param)
	if err != nil {
		return types.AddonLink{}, err
	}

	if row.Next() {
		if err := row.Scan(&param.ID); err != nil {
			return types.AddonLink{}, err
		}
	}
	return param, err
}

func (s *addonLinkStore) GetAddonsLinkById(id int) (types.AddonLink, error) {
	var link types.AddonLink
	query := `select * from addons_links  where id= $1`
	if err := s.db.Get(&link, query, id); err != nil {
		return types.AddonLink{}, err
	}
	return link, nil
}

func (s *addonLinkStore) GetAddonsLinkByPlanId(id int) (types.AddonLink, error) {
	var link types.AddonLink
	query := `select * from addons_links  where addon_id= $1`
	if err := s.db.Get(&link, query, id); err != nil {
		return types.AddonLink{}, err
	}
	return link, nil
}

func (s *addonLinkStore) DeleteAddonsLinkById(id int) error {
	query := `DELETE FROM addons_links WHERE id = :id;`
	_, err := s.db.NamedExec(query, map[string]any{"id": id})
	if err != nil {
		return err
	}
	return nil
}

func (s *addonLinkStore) DeleteAddonLinkByPlanId(id int) error {
	query := `DELETE FROM addons_links WHERE plan_id = :id;`
	_, err := s.db.NamedExec(query, map[string]any{"id": id})
	if err != nil {
		return err
	}
	return nil
}

func (s *addonLinkStore) DeleteAddonLinkByAddonId(id int) error {
	query := `DELETE FROM addons_links WHERE addon_id = :id;`
	_, err := s.db.NamedExec(query, map[string]any{"id": id})
	if err != nil {
		return err
	}
	return nil
}

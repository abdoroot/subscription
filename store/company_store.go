package store

import (
	"fmt"
	"time"

	"github.com/abdoroot/subscription/types"
	"github.com/abdoroot/subscription/util"
	"github.com/jmoiron/sqlx"
)

type companyStore struct {
	db *sqlx.DB
}

func NewCompanyStore(db *sqlx.DB) *companyStore {
	return &companyStore{
		db: db,
	}
}

func (s *userStore) CreateCompany(param types.Company) (types.Company, error) {
	param.CreatedAt = time.Now().UTC()
	query := `insert into companies (country_id,city_id,name,is_vat_registered,tax_registration_number,vat_registration_number,currency_id,billing_type,created_at) 
	values (:country_id,:city_id,:name,:is_vat_registered,:tax_registration_number,:vat_registration_number,:currency_id,:billing_type,:created_at) RETURNING id;`
	row, err := s.db.NamedQuery(query, param)
	if err != nil {
		return types.Company{}, err
	}
	var lastInsertId int
	if row.Next() {
		row.Scan(&lastInsertId)
	}
	param.Id = int(lastInsertId)
	return param, nil
}

func (s *userStore) UpdateCompanyById(param types.Company, excludeTags ...string) (types.Company, error) {
	excludeTags = append(excludeTags, "id", "billing_type", "created_at")
	query, err := util.SqlxStructUpdateQueryBuilder(param, "companies", excludeTags...)
	if err != nil {
		return types.Company{}, err
	}
	_, err = s.db.NamedExec(query, param)
	if err != nil {
		return types.Company{}, err
	}
	return param, nil
}

func (s *userStore) DeleteCompanyById(id int) error {
	query := fmt.Sprintf("delete from companies where id=%v", id)
	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

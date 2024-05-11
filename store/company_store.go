package store

import (
	"fmt"
	"time"

	"github.com/abdoroot/subscription/types"
	"github.com/abdoroot/subscription/util"
	"github.com/jmoiron/sqlx"
)

type companyStore struct {
	db       *sqlx.DB
	rolStore *roleStore
}

func NewCompanyStore(db *sqlx.DB, rolStore *roleStore) *companyStore {
	return &companyStore{
		db:       db,
		rolStore: rolStore,
	}
}

func (s *companyStore) CreateCompany(param types.Company) (types.Company, error) {
	param.CreatedAt = time.Now().UTC()
	query := `insert into companies (country_id,city_id,name,is_vat_registered,tax_registration_number,vat_registration_number,currency_id,billing_type,created_at) 
	values (:country_id,:city_id,:name,:is_vat_registered,:tax_registration_number,:vat_registration_number,:currency_id,:billing_type,:created_at) RETURNING id;`
	row, err := s.db.NamedQuery(query, param)
	if err != nil {
		return types.Company{}, err
	}
	var lastInsertId int
	if row.Next() {
		if err := row.Scan(&lastInsertId); err != nil {
			return types.Company{}, err
		}
	}
	param.Id = int(lastInsertId)

	//Create admin staff roles
	if err = s.CreateAdminStaffRoles(lastInsertId); err != nil {
		return types.Company{}, err
	}

	return param, nil
}

func (s *companyStore) UpdateCompany(param types.Company, excludeTags ...string) (types.Company, error) {
	param.UpdatedAt = time.Now()
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

func (s *companyStore) DeleteCompanyById(id int) error {
	//First delete company related roles
	if err := s.rolStore.DeleteRolesByCompanyID(id); err != nil {
		return err
	}
	//Second delete the company
	query := fmt.Sprintf("delete from companies where id=%v", id)
	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (s *companyStore) CreateAdminStaffRoles(companyId int) error {
	// Create admin role
	adminRole := types.Role{
		CompanyID:              companyId,
		Name:                   "Admin",
		CanAddProducts:         true,
		CanUpdateProducts:      true,
		CanDeleteProducts:      true,
		CanViewProducts:        true,
		CanAddCustomers:        true,
		CanUpdateCustomers:     true,
		CanDeleteCustomers:     true,
		CanViewCustomers:       true,
		CanAddSubscriptions:    true,
		CanUpdateSubscriptions: true,
		CanDeleteSubscriptions: true,
		CanViewSubscriptions:   true,
		CanAddItems:            true,
		CanUpdateItems:         true,
		CanDeleteItems:         true,
		CanViewItems:           true,
		CanAddInvoices:         true,
		CanUpdateInvoices:      true,
		CanDeleteInvoices:      true,
		CanViewInvoices:        true,
		CanAddPayments:         true,
		CanUpdatePayments:      true,
		CanDeletePayments:      true,
		CanViewPayments:        true,
		CanViewDashboard:       true,
		CreatedAt:              time.Now().UTC(),
		UpdatedAt:              time.Now().UTC(),
	}

	// Create staff role
	staffRole := types.Role{
		CompanyID:              companyId,
		Name:                   "Staff",
		CanAddProducts:         true,
		CanUpdateProducts:      true,
		CanDeleteProducts:      false,
		CanViewProducts:        true,
		CanAddCustomers:        true,
		CanUpdateCustomers:     true,
		CanDeleteCustomers:     false,
		CanViewCustomers:       true,
		CanAddSubscriptions:    true,
		CanUpdateSubscriptions: true,
		CanDeleteSubscriptions: false,
		CanViewSubscriptions:   true,
		CanAddItems:            true,
		CanUpdateItems:         true,
		CanDeleteItems:         false,
		CanViewItems:           true,
		CanAddInvoices:         true,
		CanUpdateInvoices:      true,
		CanDeleteInvoices:      false,
		CanViewInvoices:        true,
		CanAddPayments:         true,
		CanUpdatePayments:      true,
		CanDeletePayments:      false,
		CanViewPayments:        true,
		CanViewDashboard:       true,
		CreatedAt:              time.Now().UTC(),
		UpdatedAt:              time.Now().UTC(),
	}

	// Save admin role
	_, err := s.rolStore.CreateRole(adminRole)
	if err != nil {
		return err
	}

	// Save staff role
	_, err = s.rolStore.CreateRole(staffRole)
	if err != nil {
		return err
	}

	return nil
}

func (s *companyStore) GetAll() ([]types.Company, error) {
	var companies []types.Company
	query := `SELECT * FROM companies`
	err := s.db.Select(&companies, query)
	if err != nil {
		return []types.Company{}, err
	}
	return companies, nil
}

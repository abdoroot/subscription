package store

import (
	"testing"

	"github.com/abdoroot/subscription/types"
	"github.com/abdoroot/subscription/util"
	"github.com/stretchr/testify/assert"
)

func TestCreateCompanyUpdateDelete(t *testing.T) {
	db, err := util.ConnectToPq()
	if err != nil {
		t.Error("error connect to db")
	}
	s := NewUserStore(db)
	CreateCompanyParam := types.Company{
		Id:                    8,
		CountryId:             5,
		CityId:                6,
		IsVatRegistered:       false,
		VatRegistrationNumber: "33333333",
		TaxRegistrationNumber: "222222",
		Name:                  "Abdelhadi llc 2",
		CurrencyId:            8,
		BillingType:           "once",
	}

	company, err := s.CreateCompany(CreateCompanyParam)
	if err != nil {
		t.Errorf("error CreateCompany: %v", err)
	}

	assert.Greater(t, company.Id, 0)
	assert.Equal(t, CreateCompanyParam.CountryId, company.CountryId)
	assert.Equal(t, CreateCompanyParam.CityId, company.CityId)
	assert.Equal(t, CreateCompanyParam.IsVatRegistered, company.IsVatRegistered)
	assert.Equal(t, CreateCompanyParam.VatRegistrationNumber, company.VatRegistrationNumber)
	assert.Equal(t, CreateCompanyParam.TaxRegistrationNumber, company.TaxRegistrationNumber)
	assert.Equal(t, CreateCompanyParam.Name, company.Name)
	assert.Equal(t, CreateCompanyParam.CurrencyId, company.CurrencyId)
	assert.Equal(t, CreateCompanyParam.BillingType, company.BillingType)

	t.Run("update company", func(t *testing.T) {
		UpdateCompanyParam := types.Company{
			CountryId:             company.Id,
			CityId:                6,
			IsVatRegistered:       false,
			VatRegistrationNumber: "77847",
			TaxRegistrationNumber: "226697",
			Name:                  "Abdelhadi llc",
			CurrencyId:            7,
			BillingType:           "once",
		}
		updatedCompany, err := s.UpdateCompany(UpdateCompanyParam)
		if err != nil {
			t.Errorf("error UpdateCompany: %v", err)
		}

		assert.Equal(t, updatedCompany.CountryId, UpdateCompanyParam.CountryId)
		assert.Equal(t, updatedCompany.CityId, UpdateCompanyParam.CityId)
		assert.Equal(t, updatedCompany.IsVatRegistered, UpdateCompanyParam.IsVatRegistered)
		assert.Equal(t, updatedCompany.VatRegistrationNumber, UpdateCompanyParam.VatRegistrationNumber)
		assert.Equal(t, updatedCompany.TaxRegistrationNumber, UpdateCompanyParam.TaxRegistrationNumber)
		assert.Equal(t, updatedCompany.Name, UpdateCompanyParam.Name)
		assert.Equal(t, updatedCompany.CurrencyId, UpdateCompanyParam.CurrencyId)
		assert.Equal(t, updatedCompany.BillingType, UpdateCompanyParam.BillingType)
	})

	t.Run("delete company", func(t *testing.T) {
		err := s.DeleteCompanyById(company.Id)
		if err != nil {
			t.Errorf("error DeleteCompanyById: %v", err)
		}
	})
}

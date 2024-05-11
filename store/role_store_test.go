package store_test

import (
	"testing"
	"time"

	"github.com/abdoroot/subscription/store"
	"github.com/abdoroot/subscription/types"
	"github.com/abdoroot/subscription/util"
	"github.com/stretchr/testify/assert"
)

func TestRoleStore(t *testing.T) {
	db, err := util.ConnectToPq()
	if err != nil {
		t.Fatal("error connecting to the database:", err)
	}
	roleStore := store.NewRoleStore(db)

	t.Run("CreateRole", func(t *testing.T) {
		role := types.Role{
			CompanyID:              1,
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
		createdRole, err := roleStore.CreateRole(role)
		defer roleStore.DeleteRoleByID(createdRole.ID)
		assert.NoError(t, err)
		assert.NotZero(t, createdRole.ID)
		assert.Equal(t, role.Name, createdRole.Name)
	})

	t.Run("UpdateRole", func(t *testing.T) {
		role, err := creatRole()
		if err != nil {
			t.Fatal("error creatRole", err)
		}
		defer roleStore.DeleteRoleByID(role.ID)
		updatedRole := types.Role{
			ID:                     role.ID,
			Name:                   "Updated Admin",
			CanAddProducts:         false,
			CanUpdateProducts:      false,
			CanDeleteProducts:      false,
			CanViewProducts:        false,
			CanAddCustomers:        false,
			CanUpdateCustomers:     false,
			CanDeleteCustomers:     false,
			CanViewCustomers:       false,
			CanAddSubscriptions:    false,
			CanUpdateSubscriptions: false,
			CanDeleteSubscriptions: false,
			CanViewSubscriptions:   false,
			CanAddItems:            false,
			CanUpdateItems:         false,
			CanDeleteItems:         false,
			CanViewItems:           false,
			CanAddInvoices:         false,
			CanUpdateInvoices:      false,
			CanDeleteInvoices:      false,
			CanViewInvoices:        false,
			CanAddPayments:         false,
			CanUpdatePayments:      false,
			CanDeletePayments:      false,
			CanViewPayments:        false,
			CanViewDashboard:       false,
			UpdatedAt:              time.Now().UTC(),
		}
		err = roleStore.UpdateRole(updatedRole)
		assert.NoError(t, err)
	})

	t.Run("GetRoleByID", func(t *testing.T) {
		role, err := creatRole()
		if err != nil {
			t.Fatal("error creatRole", err)
		}
		defer roleStore.DeleteRoleByID(role.ID)
		r, err := roleStore.GetRoleByID(role.ID)
		if err != nil {
			t.Fatal("error GetRoleByID:", err)
		}
		assert.NotNil(t, r)
		assert.Equal(t, r.ID, role.ID)
	})

	t.Run("DeleteRoleByID", func(t *testing.T) {
		role, err := creatRole()
		if err != nil {
			t.Fatal("error creatRole", err)
		}
		defer roleStore.DeleteRoleByID(role.ID)
		if err = roleStore.DeleteRoleByID(role.ID); err != nil {
			t.Error("error DeleteRoleByID:", err)
		}
		assert.NoError(t, err)
	})
}

func TestGetAll(t *testing.T) {
	db, err := util.ConnectToPq()
	if err != nil {
		t.Fatal("error connecting to the database:", err)
	}
	roleStore := store.NewRoleStore(db)
	role, err := creatRole()
	if err != nil {
		t.Fatal("error creatRole", err)
	}
	defer roleStore.DeleteRoleByID(role.ID)
	roles, err := roleStore.GetAll()
	if err != nil {
		t.Fatal("error GetAll:", err)
	}
	assert.Greater(t, len(roles), 0)
}

func GetRolesByCompanyId(t *testing.T) {
	db, err := util.ConnectToPq()
	if err != nil {
		t.Fatal("error connecting to the database:", err)
	}
	roleStore := store.NewRoleStore(db)
	role, err := creatRole()
	if err != nil {
		t.Fatal("error creatRole", err)
	}
	defer roleStore.DeleteRoleByID(role.ID)
	roles, err := roleStore.GetRolesByCompanyId(role.CompanyID)
	if err != nil {
		t.Fatal("error GetRolesByCompanyId:", err)
	}
	assert.Greater(t, len(roles), 0)
}

// helper func
func creatRole() (types.Role, error) {
	db, err := util.ConnectToPq()
	if err != nil {
		return types.Role{}, err
	}
	roleStore := store.NewRoleStore(db)
	role := types.Role{
		CompanyID:              1,
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
	createdRole, err := roleStore.CreateRole(role)
	if err != nil {
		return types.Role{}, err
	}

	return createdRole, nil
}

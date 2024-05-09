package store

import (
	"testing"

	"github.com/abdoroot/subscription/types"
	"github.com/abdoroot/subscription/util"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	db, err := util.ConnectToPq()
	if err != nil {
		t.Error("error connect to db")
	}
	userStore := NewUserStore(db)

	req := types.CreateUserParam{
		CompanyId:   1,
		RoleId:      2,
		Name:        "abdelhadi",
		Email:       "abdelhadi@gmail.com",
		PhoneNumber: "0554241891",
		Password:    "11223344",
		IsActive:    true,
	}
	user, err := req.CreateUserFromRequest()
	if err != nil {
		t.Errorf("error CreateUserFromRequest: %v", err)
		t.Fail()
	}

	insertedUser, err := userStore.CreateUser(user)
	if err != nil {
		t.Errorf("error Store CreateUser: %v", err)
		t.Fail()
	}

	assert.Equal(t, 60, len(user.Password))
	assert.Equal(t, user.CompanyId, insertedUser.CompanyId)
	assert.Equal(t, user.Name, insertedUser.Name)
}

func TestUpdateUserById(t *testing.T) {
	db, err := util.ConnectToPq()
	if err != nil {
		t.Error("error connect to db")
	}
	request := types.UpdateUserParam{
		RoleId:      3,
		Name:        "abdelhadi mohamed",
		Email:       "abdelhadi.200930@gmail.com",
		PhoneNumber: "0554241896",
		Password:    "11223355",
		IsActive:    true,
	}

	userStore := NewUserStore(db)
	user, err := request.CreateUpdateRequest()
	user.Id = 1 //for testing perpose
	if err != nil {
		t.Errorf("error CreateUpdateRequest: %v", err)
	}
	UpdateUser, err := userStore.UpdateUser(user)
	_ = UpdateUser
	if err != nil {
		t.Errorf("error UpdateUser: %v", err)
	}
}

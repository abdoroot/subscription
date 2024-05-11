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
	defer userStore.DeleteUserById(insertedUser.Id)
	if err != nil {
		t.Errorf("error Store CreateUser: %v", err)
		t.Fail()
	}

	assert.Greater(t, insertedUser.Id, 0)
	assert.Equal(t, 60, len(user.Password))
	assert.Equal(t, user.CompanyId, insertedUser.CompanyId)
	assert.Equal(t, user.Name, insertedUser.Name)
}

func TestUpdateUserById(t *testing.T) {
	db, err := util.ConnectToPq()
	if err != nil {
		t.Error("error connect to db")
	}
	user, err := createUser()
	if err != nil {
		t.Error("error createUser")
	}
	userStore := NewUserStore(db)
	defer userStore.DeleteUserById(user.Id)
	request := types.UpdateUserParam{
		RoleId:      3,
		Name:        "abdelhadi mohamed",
		Email:       "abdelhadi.200930@gmail.com",
		PhoneNumber: "0554241896",
		Password:    "11223355",
		IsActive:    true,
	}

	req, err := request.CreateUpdateRequest()
	req.Id = user.Id
	if err != nil {
		t.Errorf("error CreateUpdateRequest: %v", err)
	}

	if err := userStore.UpdateUser(user); err != nil {
		t.Errorf("error UpdateUser: %v", err)
	}
}

func TestDeleteUserById(t *testing.T) {
	db, err := util.ConnectToPq()
	if err != nil {
		t.Error("error connect to db")
	}

	user, err := createUser()
	if err != nil {
		t.Error("error createUser")
	}
	userStore := NewUserStore(db)
	err = userStore.DeleteUserById(user.Id)
	assert.Nil(t, err, "DeleteUserById err")
}

func TestGetUsers(t *testing.T) {
	db, err := util.ConnectToPq()
	if err != nil {
		t.Error("error connect to db")
	}

	user, err := createUser()
	if err != nil {
		t.Error("error createUser")
	}
	assert.NotNil(t, user)
	userStore := NewUserStore(db)
	defer userStore.DeleteUserById(user.Id)
	t.Run("GetAllUsers", func(t *testing.T) {
		users, err := userStore.GetAllUsers()
		if err != nil {
			t.Error("error GetAllUsers", err)
		}
		assert.Greater(t, len(users), 0)
	})

	t.Run("GetAllUserByCompnayId", func(t *testing.T) {
		users, err := userStore.GetAllUserByCompnayId(user.CompanyId)
		if err != nil {
			t.Error("error GetAllUsers", err)
		}
		assert.Greater(t, len(users), 0)
	})

	t.Run("GetUserById", func(t *testing.T) {
		u, err := userStore.GetUserById(user.Id)
		if err != nil {
			t.Error("error GetAllUsers", err)
		}
		assert.Equal(t, u.Id, user.Id)
	})
}

// helper func
func createUser() (types.User, error) {
	db, err := util.ConnectToPq()
	if err != nil {
		return types.User{}, nil
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
		return types.User{}, nil
	}

	insertedUser, err := userStore.CreateUser(user)
	if err != nil {
		return types.User{}, nil
	}
	return insertedUser, nil
}

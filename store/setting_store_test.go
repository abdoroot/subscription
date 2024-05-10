package store

import (
	"testing"

	"github.com/abdoroot/subscription/types"
	"github.com/abdoroot/subscription/util"
	"github.com/stretchr/testify/assert"
)

var insertedSetting types.Setting

func TestCreateSetting(t *testing.T) {
	db, err := util.ConnectToPq()
	if err != nil {
		t.Error("error connect to db")
	}

	s := NewSettingStore(db)
	param := types.Setting{
		CompanyId: 2,
		Name:      "vat",
		Value:     "54779997777",
	}
	c, err := s.CreateSetting(param)
	if err != nil {
		t.Error("errorCreateSetting ", err)
	}
	assert.Nil(t, err)
	assert.Greater(t, c.Id, 0)
	assert.Equal(t, param.CompanyId, c.CompanyId)
	assert.Equal(t, param.Value, c.Value)
	insertedSetting = c
}

func TestUpdateSettingValueById(t *testing.T) {
	if insertedSetting.Id == 0 {
		t.Error("TestUpdateSettingById depend on insertedSetting run the all test")
	}

	db, err := util.ConnectToPq()
	if err != nil {
		t.Error("error connect to db")
	}

	s := NewSettingStore(db)
	err = s.UpdateSettingValueById("vat", insertedSetting.Id)
	if err != nil {
		t.Error("error UpdateSettingValueById ", err)
	}
}

func TestDeleteSettingById(t *testing.T) {
	if insertedSetting.Id == 0 {
		t.Error("TestDeleteSettingById Depend on insertedSetting run the all test")
	}

	db, err := util.ConnectToPq()
	if err != nil {
		t.Error("error connect to db")
	}

	s := NewSettingStore(db)
	err = s.DeleteSettingById(insertedSetting.Id)
	if err != nil {
		t.Error("error DeleteSettingById ", err)
	}
}

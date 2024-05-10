package fixtures

import (
	"github.com/abdoroot/subscription/store"
	"github.com/abdoroot/subscription/types"
	"github.com/abdoroot/subscription/util"
)

func AddSettings() (types.Setting, error) {
	db, err := util.ConnectToPq()
	if err != nil {
		return types.Setting{}, err
	}

	s := store.NewSettingStore(db)
	param := types.Setting{
		CompanyId: 2,
		Name:      "vat",
		Value:     "54779997777",
	}
	set, err := s.CreateSetting(param)
	if err != nil {
		return types.Setting{}, err
	}
	return set, nil
}

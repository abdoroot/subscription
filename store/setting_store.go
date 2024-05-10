package store

import (
	"time"

	"github.com/abdoroot/subscription/types"
	"github.com/jmoiron/sqlx"
)

type settingStore struct {
	db *sqlx.DB
}

func NewSettingStore(db *sqlx.DB) *settingStore {
	return &settingStore{db: db}
}

func (s *settingStore) CreateSetting(param types.Setting) (types.Setting, error) {
	param.CreatedAt = time.Now()
	query := `insert into settings (company_id,name,value,created_at) values(:company_id,:name,:value,:created_at) RETURNING id;`
	row, err := s.db.NamedQuery(query, param)
	if err != nil {
		return types.Setting{}, err
	}
	var lastInsertId int
	if row.Next() {
		row.Scan(&lastInsertId)
		param.Id = lastInsertId
	}
	return param, nil
}

package types

import "time"

type Setting struct {
	Id        int       `db:"id"`
	CompanyId int       `db:"company_id"`
	Name      string    `db:"name"`
	Value     string    `db:"value"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type UpdateSettingRequest struct {
	Id       int       `db:"id"`
	Value    string    `db:"value"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (s Setting) Validate() bool {
	return true
}

package types

import "time"

type Attachment struct {
	ID        int       `db:"id"`
	CompanyId int       `db:"company_id"`
	UserId    int       `db:"user_id"`
	FileName  string    `db:"filename"`
	MimeType  string    `db:"mime_type"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

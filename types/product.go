package types

import "time"

type Product struct {
	Id                int       `db:"id"`
	CompanyId         int       `db:"company_id"`
	Name              string    `db:"name"`
	Type              string    `db:"type"`
	UnitId            int       `db:"unit_id"`
	Price             float32   `db:"price"`
	Description       string    `db:"description"`
	ImageAttachmentId int       `db:"image_attachment_id"`
	CreatedAt         time.Time `db:"created_at"`
	UpdatedAt         time.Time `db:"updated_at"`
}

type UpdateProductRequest struct {
	Id                int       `db:"id"`
	Name              string    `db:"name"`
	Type              string    `db:"type"`
	UnitId            int       `db:"unit_id"`
	Price             float32   `db:"price"`
	Description       string    `db:"description"`
	ImageAttachmentId int       `db:"image_attachment_id"`
	UpdatedAt         time.Time `db:"updated_at"`
}

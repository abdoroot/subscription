package store

import (
	"fmt"
	"time"

	"github.com/abdoroot/subscription/types"
	"github.com/jmoiron/sqlx"
)

type attachmentStore struct {
	db *sqlx.DB
}

func NewAttachmentStore(db *sqlx.DB) *attachmentStore {
	return &attachmentStore{db: db}
}

func (s *attachmentStore) Create(param types.Attachment) (types.Attachment, error) {
	param.CreatedAt = time.Now()
	param.UpdatedAt = time.Now()
	query := `insert into attachments(company_id,user_id,filename,mime_type,created_at,updated_at) 
	values(:company_id,:user_id,:filename,:mime_type,:created_at,:updated_at) returning id;`
	row, err := s.db.NamedQuery(query, param)
	if err != nil {
		return types.Attachment{}, err
	}
	var lastInstertId int
	if row.Next() {
		if err := row.Scan(&lastInstertId); err != nil {
			return types.Attachment{}, err
		}
	}
	param.ID = lastInstertId
	return param, nil
}

func (s *attachmentStore) GetById(id int) (types.Attachment, error) {
	var attachment types.Attachment
	query := `select * from attachments where id=$1`
	if err := s.db.Get(&attachment, query, id); err != nil {
		return attachment, err
	}
	return attachment, nil
}

func (s *attachmentStore) UpdateByid(param types.Attachment, id int) error {
	param.ID = id
	param.UpdatedAt = time.Now()
	query := `update attachments set filename = :filename,mime_type = :mime_type,updated_at=:updated_at where id = :id`
	_, err := s.db.NamedQuery(query, param)
	if err != nil {
		return err
	}

	return nil
}

func (s *attachmentStore) DeleteById(id int) error {
	query := fmt.Sprintf("delete from attachments where id=%v", id)
	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (s *attachmentStore) DeleteByCompanyId(id int) error {
	query := fmt.Sprintf("delete from attachments where company_id=%v", id)
	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

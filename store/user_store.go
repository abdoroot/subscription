package store

import (
	"fmt"
	"time"

	"github.com/abdoroot/subscription/types"
	"github.com/abdoroot/subscription/util"
	"github.com/jmoiron/sqlx"
)

type userStore struct {
	db *sqlx.DB
}

func NewUserStore(db *sqlx.DB) *userStore {
	return &userStore{
		db: db,
	}
}

func (s *userStore) CreateUser(user types.User) (types.User, error) {
	user.CreatedAt = time.Now().UTC()
	query := `insert into users (company_id,role_id,name,email,password,phone_number,is_active,created_at) 
	values (:company_id, :role_id, :name, :email, :password, :phone_number,:is_active, :created_at) RETURNING id;`
	row, err := s.db.NamedQuery(query, user)
	if err != nil {
		return types.User{}, err
	}
	var lastInsertId int
	if row.Next() {
		row.Scan(&lastInsertId)
	}
	user.Id = int(lastInsertId)
	return user, nil
}

func (s *userStore) UpdateUser(user types.User, excludeTags ...string) (types.User, error) {
	user.UpdatedAt = time.Now()
	excludeTags = append(excludeTags, "id", "company_id", "created_at")
	query, err := util.SqlxStructUpdateQueryBuilder(user, "users", excludeTags...)
	if err != nil {
		return types.User{}, err
	}

	_, err = s.db.NamedExec(query, user)
	if err != nil {
		return types.User{}, err
	}
	return types.User{}, nil
}

func (s *userStore) DeleteUserById(id int) error {
	query := fmt.Sprintf("delete from users where id=%v", id)
	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

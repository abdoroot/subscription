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

func (s *userStore) GetAllUsers() ([]types.User, error) {
	query := `select * from users;`
	var users []types.User
	if err := s.db.Select(&users, query); err != nil {
		return []types.User{}, err
	}
	return users, nil
}

func (s *userStore) GetAllUserByCompnayId(id int) ([]types.User, error) {
	query := `select * from users where company_id = $1;`
	var users []types.User
	if err := s.db.Select(&users, query, id); err != nil {
		return []types.User{}, err
	}
	return users, nil
}

func (s *userStore) GetUserById(id int) (types.User, error) {
	query := `select * from users where id =$1;`
	var user types.User
	if err := s.db.Get(&user, query, id); err != nil {
		return types.User{}, err
	}
	return user, nil
}

func (s *userStore) UpdateUser(user types.User, excludeTags ...string) error {
	user.UpdatedAt = time.Now()
	excludeTags = append(excludeTags, "id", "company_id", "created_at")
	query, err := util.SqlxStructUpdateQueryBuilder(user, "users", excludeTags...)
	if err != nil {
		return err
	}

	_, err = s.db.NamedExec(query, user)
	if err != nil {
		return err
	}
	return nil
}

func (s *userStore) DeleteUserById(id int) error {
	query := fmt.Sprintf("delete from users where id=%v", id)
	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

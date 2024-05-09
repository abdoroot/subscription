package types

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost = 4
)

type User struct {
	Id          int       `db:"id"`
	CompanyId   int       `db:"company_id"`
	RoleId      int       `db:"role_id"`
	Name        string    `db:"name"`
	Email       string    `db:"email"`
	PhoneNumber string    `db:"phone_number"`
	Password    string    `db:"password" json:"-"`
	IsActive    bool      `db:"is_active"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type CreateUserParam struct {
	CompanyId   int    `json:"company_id"`
	RoleId      int    `json:"role_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	IsActive    bool   `json:""`
}

func (r CreateUserParam) Validate() bool {
	return true
}

func (r CreateUserParam) CreateUserFromRequest() (User, error) {
	EncreptedPassword, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcryptCost)
	if err != nil {
		return User{}, err
	}
	return User{
		CompanyId:   r.CompanyId,
		RoleId:      r.RoleId,
		Name:        r.Name,
		Email:       r.Email,
		PhoneNumber: r.PhoneNumber,
		Password:    string(EncreptedPassword),
		IsActive:    r.IsActive,
	}, nil
}

type UpdateUserParam struct {
	RoleId      int    `json:"role_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	IsActive    bool   `json:"is_active"`
}

func (r UpdateUserParam) CreateUpdateRequest() (User, error) {
	user := User{
		RoleId:      r.RoleId,
		Name:        r.Name,
		Email:       r.Email,
		PhoneNumber: r.PhoneNumber,
		IsActive:    r.IsActive,
		UpdatedAt:   time.Now(),
	}

	if r.Password != "" {
		EncreptedPassword, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcryptCost)
		if err != nil {
			return User{}, err
		}
		user.Password = string(EncreptedPassword)
	}
	return user, nil
}

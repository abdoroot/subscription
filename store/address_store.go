package store

import (
	"time"

	"github.com/abdoroot/subscription/types"
	"github.com/jmoiron/sqlx"
)

type addressStore struct {
	db *sqlx.DB
}

func NewAddressStore(db *sqlx.DB) *addressStore {
	return &addressStore{
		db: db,
	}
}

func (s *addressStore) CreateAddress(address types.Address) (types.Address, error) {
    address.CreatedAt = time.Now().UTC()
    address.UpdatedAt = time.Now().UTC()
    query := `
        INSERT INTO addresses (type, customer_id, country_id, city_id, line1, line2, created_at, updated_at)
        VALUES (:type, :customer_id, :country_id, :city_id, :line1, :line2, :created_at, :updated_at)
        RETURNING id;
    `
    row, err := s.db.NamedQuery(query, address)
    if err != nil {
        return types.Address{}, err
    }
    defer row.Close()

    if row.Next() {
        if err := row.Scan(&address.ID); err != nil {
            return types.Address{}, err
        }
    }

    return address, nil
}

func (s *addressStore) UpdateAddress(address types.Address) error {
	address.UpdatedAt = time.Now().UTC()
	query := `
        UPDATE addresses
        SET type = :type, customer_id = :customer_id, country_id = :country_id, city_id = :city_id,
            line1 = :line1, line2 = :line2, updated_at = :updated_at
        WHERE id = :id;
    `
	_, err := s.db.NamedExec(query, address)
	if err != nil {
		return err
	}
	return nil
}

func (s *addressStore) DeleteAddressByID(id int) error {
	query := "DELETE FROM addresses WHERE id = :id;"
	_, err := s.db.NamedExec(query, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	return nil
}

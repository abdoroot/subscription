package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSqlxStructUpdateQueryBuilder(t *testing.T) {
	s := struct {
		Id       int    `db:"id"`
		Password string `db:"-"`
		Name     string `db:"name"`
	}{
		Id:       10,
		Name:     "ahmed",
		Password: "14556",
	}
	tn := "users"
	q, err := SqlxStructUpdateQueryBuilder(s, tn)
	if err != nil {
		t.Error("GetSqlxUpdateStructQuery", err)
	}
	assert.Contains(t, q, tn)
}

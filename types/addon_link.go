package types

import "time"

type AddonLink struct {
	ID        int       `db:"id"`
	AddonId   int       `db:"addon_id"`
	PlanId    int       `db:"plan_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (c AddonLink) Validate() bool {
	return true
}

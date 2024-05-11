package types

import "time"

type SubscriptionFeature struct {
	Id                 int       `db:"id"`
	SubscriptionPlanId int       `db:"subscription_plan_id"`
	Name               string    `db:"name"`
	CreatedAt          time.Time `db:"created_at"`
	UpdatedAt          time.Time `db:"updated_at"`
}

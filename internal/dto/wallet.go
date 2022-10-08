package dto

import "time"

type Wallet struct {
	ID         string
	OwnedBy    string
	Status     string
	EnabledAt  time.Time
	DisabledAt time.Time
	Balance    int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

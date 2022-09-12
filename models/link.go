package models

import (
	"time"
)

type Link struct {
	ID string `json:"id" gorm:"primary_key,not null"`
	Address *string `json:"address"`
	Banned *bool `json:"banned,omitempty"`
	Link *string `json:"link"`
	Password *string `json:"password,omitempty"`
	Target *string `json:"target"`
	Description *string `json:"description"`
	VisitCount int `json:"visit_count"`
	ExpireAt *time.Time `json:"expire_at,omitempty"`
	Reusable *bool `json:"reusable,omitempty"`
	UserID *string `json:"user_id,omitempty"`
	IP *string `json:"ip,omitempty"`
	// Stats Stat `json:"stats,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

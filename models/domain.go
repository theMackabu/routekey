package models

type Domain struct {
	ID *string `json:"id" gorm:"primary_key,not null"`
	Address *string `json:"address"`
	Banned *bool `json:"banned,omitempty"`
	Homepage *string `json:"homepage"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

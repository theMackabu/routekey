package models

import "time"

type Stat struct {
	AllTime StatItem `json:"all_time"`
	LastDay StatItem `json:"last_day"`
	LastWeek StatItem `json:"last_week"`
	LastMonth StatItem `json:"last_month"`
	Address string `json:"address"`
	Banned bool `json:"banned,omitempty"`
	Link string `json:"link"`
	Password string `json:"password,omitempty"`
	Target string `json:"target"`
	Description string `json:"description"`
	ExpireAt time.Time `json:"expire_at,omitempty"`
	Reusable *bool `json:"reusable,omitempty"`
	UserID string `json:"user_id,omitempty"`
	IP string `json:"ip,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type StatItem struct {
	Browser Browser `json:"browser"`
	OS OS `json:"os"`
	Country Country `json:"country"`
	Referrer Referrer `json:"referrer"`
	Views    int      `json:"views"`
}

type Browser struct {
	Name string `json:"name"`
	Value int `json:"value"`
}

type OS struct {
	Name string `json:"name"`
	Value int `json:"value"`
}

type Country struct {
	Name string `json:"name"`
	Value int `json:"value"`
}

type Referrer struct {
	Name string `json:"name"`
	Value int `json:"value"`
}

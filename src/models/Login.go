package models

import "time"

type Authentication struct {
	Email    string    `json:"email,omitempty"`
	Password string    `json:"password,omitempty"`
	Tenant   string    `json:"tenant,omitempty"`
	CreateAt time.Time `json:"createAt,omitempty"`
}

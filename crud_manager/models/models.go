package models

import (
	"time"
)

type User struct {
	Id        interface{} `json:"_id,omitempty" bson:"_id,omitempty"`
	Email     string      `json:"email"`
	FirstName string      `json:"firstName,omitempty"`
	LastName  string      `json:"lastName,omitempty"`
	Password  string      `json:"password,omitempty"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt,omitempty"`
}

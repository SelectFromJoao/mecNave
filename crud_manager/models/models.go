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

type Company struct {
	Id             interface{} `json:"_id,omitempty" bson:"_id,omitempty"`
	BrandSpecialty string      `json:"BrandSpecialty,omitempty"`
	CompanyTitle   string      `json:"CompanyTitle,omitempty"`
	Description    string      `json:"Description,omitempty"`
	Email          string      `json:"Email,omitempty"`
	Localization   string      `json:"Localization,omitempty"`
	Password       string      `json:"password,omitempty"`
	Reviews        []Review    `json:"Reviews,omitempty"`
	CreatedAt      time.Time   `json:"createdAt"`
	UpdatedAt      time.Time   `json:"updatedAt,omitempty"`
}

type Review struct {
	Id        interface{} `json:"_id,omitempty" bson:"_id,omitempty"`
	Review    string      `json:"review,omitempty"`
	User      string      `json:"User,omitempty"`
	UserID    string      `json:"UserID,omitempty" bson:"UserID,omitempty"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt,omitempty"`
}

type Banner struct {
	Id          interface{} `json:"_id,omitempty" bson:"_id,omitempty"`
	BannerTitle string      `json:"BannerTitle,omitempty"`
	Banner      string      `json:"Banner" bson:"Banner"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt,omitempty"`
}

package models

import "time"

type BaseEntity struct {
	Identity
	CreatedDate
	UpdatedDate
	DeletedDate
}

type Identity struct {
	Id int64 `json:"id"`
}

type CreatedDate struct {
	CreatedAt *time.Time `json:"created_at"`
}

type UpdatedDate struct {
	UpdatedAt *time.Time `json:"updated_at"`
}

type DeletedDate struct {
	DeletedAt *time.Time `json:"deleted_at"`
}

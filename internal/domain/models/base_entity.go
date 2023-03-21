package models

import "time"

type BaseEntity struct {
	Identity
	CreatedAt
	UpdatedAt
	DeletedAt
}

type Identity struct {
	Id int64 `json:"id"`
}

type CreatedAt struct {
	CreatedAt time.Time `json:"created_at"`
}

type UpdatedAt struct {
	UpdatedAt time.Time `json:"updated_at"`
}

type DeletedAt struct {
	DeletedAt time.Time `json:"deleted_at"`
}

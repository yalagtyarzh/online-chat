package entity

import "github.com/google/uuid"

type User struct {
	UUID     uuid.UUID `json:"uuid"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Avatar   []byte    `json:"avatar"`
}

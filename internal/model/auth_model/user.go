package model

import (
	"oat431/go-fiber-snippets-vol2/pkg/common"
	"time"

	"github.com/google/uuid"
)

type User struct {
	common.BaseEntity

	AuthID    uuid.UUID `db:"auth_id" json:"auth_id"`
	Firstname string    `db:"firstname" json:"firstname"`
	Lastname  string    `db:"lastname" json:"lastname"`
	Nickname  string    `db:"nickname" json:"nickname"`
	Birthdate time.Time `db:"birthdate" json:"birthdate"`
}

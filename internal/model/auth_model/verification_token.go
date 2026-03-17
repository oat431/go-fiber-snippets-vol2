package model

import (
	"oat431/go-fiber-snippets-vol2/pkg/common"
	"time"

	"github.com/google/uuid"
)

type VerificationToken struct {
	common.BaseEntity

	AuthID    uuid.UUID `db:"auth_id" json:"auth_id"` // Foreign Key
	Token     string    `db:"token" json:"token"`
	ExpiredAt time.Time `db:"expired_at" json:"expired_at"`
	IsRevoke  bool      `db:"is_revoke" json:"is_revoke"`
}

package model

import (
	"oat431/go-fiber-snippets-vol2/pkg/common"
)

type Auth struct {
	common.BaseEntity // ฝังเหมือนเดิม!

	Username   string `db:"username" json:"username"`
	Password   string `db:"password" json:"-"`
	IsVerified bool   `db:"is_verified" json:"is_verified"`
	IsActive   bool   `db:"is_active" json:"is_active"`

	//todo: make relation myself
	User              *User              `db:"-" json:"user,omitempty"`
	VerificationToken *VerificationToken `db:"-" json:"verification_token,omitempty"`
	RefreshTokens     []RefreshToken     `db:"-" json:"refresh_tokens,omitempty"`
	Roles             []Role             `db:"-" json:"roles,omitempty"`
}

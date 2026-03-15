# Go Base Entity With Database (SQLx)

```go
package common

import (
	"time"

	"github.com/google/uuid"
)

type BaseEntity struct {
	ID        uuid.UUID  `db:"id" json:"id"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
}

```

usage:

```go
package model

import (
	"oat431/go-fiber-snippets-vol2/pkg/common"
)

type Auth struct {
	common.BaseEntity // Embedding BaseEntity for common fields

	Username   string `db:"username" json:"username"`
	Password   string `db:"password" json:"-"`
	IsVerified bool   `db:"is_verified" json:"is_verified"`
	IsActive   bool   `db:"is_active" json:"is_active"`
}

```
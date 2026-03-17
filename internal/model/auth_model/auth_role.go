package model

import "oat431/go-fiber-snippets-vol2/pkg/common"

type Role struct {
	common.BaseEntity

	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`

	// Relations
	Features []Feature `db:"-" json:"features,omitempty"`
}

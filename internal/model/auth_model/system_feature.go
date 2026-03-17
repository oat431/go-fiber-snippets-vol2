package model

import "oat431/go-fiber-snippets-vol2/pkg/common"

type Feature struct {
	common.BaseEntity

	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
}

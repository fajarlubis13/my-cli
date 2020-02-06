package model

import "time"

// HKPengiriman ...
type HKPengiriman struct {
	Oke  int64  `json:"oke" binding:"required"`
	Seep string `json:"seep"`

	CreatedAt *time.Time `json:"-"`
	CreatedBy int64      `json:"-"`
	UpdatedAt *time.Time `json:"-"`
	UpdatedBy int64      `json:"-"`
	DeletedAt *time.Time `json:"-"`
	DeletedBy int64      `json:"-"`
}

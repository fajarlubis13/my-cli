package model

import "time"

// HKPengiriman ...
type HKPengiriman struct {
	IDPengiriman int64 `json:"id_pengiriman" binding:"required"`
	IDSurat string `json:"id_surat"`
	
	CreatedAt            *time.Time `json:"-"`
	CreatedBy            int64      `json:"-"`
	UpdatedAt            *time.Time `json:"-"`
	UpdatedBy            int64      `json:"-"`
	DeletedAt            *time.Time `json:"-"`
	DeletedBy            int64      `json:"-"`
}

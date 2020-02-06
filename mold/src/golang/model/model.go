package model

import "time"

// {{ toCamel .ProjectName }} ...
type {{ toCamel .ProjectName }} struct {
	ID                   int64      `json:"id"`
	KodeJadwal           string     `json:"kode_jadwal" binding:"required"`
	Deskripsi            string     `json:"deskripsi" binding:"required"`
	IDSkemaLaporanTeknik int64      `json:"id_skema_laporan_teknik" binding:"required"`
	CreatedAt            *time.Time `json:"-"`
	CreatedBy            int64      `json:"-"`
	UpdatedAt            *time.Time `json:"-"`
	UpdatedBy            int64      `json:"-"`
	DeletedAt            *time.Time `json:"-"`
	DeletedBy            int64      `json:"-"`
}

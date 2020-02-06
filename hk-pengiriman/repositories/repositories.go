package repositories

import (
	"database/sql"
	"fmt"
	"hk-pengiriman/model"
	"strconv"

	"hk-pengiriman/helpers/database"
	"hk-pengiriman/helpers/request"

	"github.com/jmoiron/sqlx"

	//"github.com/lib/pq"
	_ "github.com/lib/pq"
)

// Repositories ...
type Repositories interface {
	CreateOne(m *model.HKPengiriman) (int64, error)
	UpdateOneByID(id int64, m *model.HKPengiriman) (int64, error)
	GetOneByID(id int64) (*model.HKPengiriman, int64, error)
	DeleteOneByID(id int64) (int64, error)
	GetAll(filter *request.QueryParameter) ([]*model.HKPengiriman, int64, error)
}

type repositories struct {
	DB *sqlx.DB
}

// NewRepositories ...
func NewRepositories() Repositories {
	return &repositories{
		DB: database.Init().SQL,
	}
}

func (v *repositories) CreateOne(m *model.HKPengiriman) (int64, error) {
	query := "INSERT INTO emonica_v2.hk_pengiriman(kode_jadwal, deskripsi, id_skema_laporan_teknik) VALUES(?, ?, ?) RETURNING id"

	if err := v.DB.QueryRowx(v.DB.Rebind(query), m.KodeJadwal, m.Deskripsi, m.IDSkemaLaporanTeknik).Scan(&m.ID); err != nil {
		return -1, err
	}

	return 1, nil
}

func (v *repositories) UpdateOneByID(id int64, m *model.HKPengiriman) (int64, error) {
	query := "UPDATE emonica_v2.hk_pengiriman SET kode_jadwal = ?, deskripsi = ?, id_skema_laporan_teknik = ? WHERE id = ?"

	res, err := v.DB.Exec(v.DB.Rebind(query), m.KodeJadwal, m.Deskripsi, m.IDSkemaLaporanTeknik, m.ID)
	if err != nil {
		return -1, err
	}
	ra, _ := res.RowsAffected()

	return ra, nil
}

func (v *repositories) GetOneByID(id int64) (*model.HKPengiriman, int64, error) {
	query := "SELECT id, kode_jadwal, deskripsi, id_skema_laporan_teknik, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by FROM emonica_v2.hk_pengiriman WHERE id = ?"

	var (
		data                 model.HKPengiriman
		kodeJadwal           sql.NullString
		deskripsi            sql.NullString
		idSkemaLaporanTeknik sql.NullInt64

		createdAt sql.NullTime
		createdBy sql.NullInt64
		updatedAt sql.NullTime
		updatedBy sql.NullInt64
		deletedAt sql.NullTime
		deletedBy sql.NullInt64
	)

	row := v.DB.QueryRowx(v.DB.Rebind(query), id)
	err := row.Scan(
		&id,
		&kodeJadwal,
		&deskripsi,
		&idSkemaLaporanTeknik,
		&createdAt,
		&createdBy,
		&updatedAt,
		&updatedBy,
		&deletedAt,
		&deletedBy,
	)
	if err != nil {
		return &data, -1, err
	}

	data.ID = id
	data.KodeJadwal = kodeJadwal.String
	data.Deskripsi = deskripsi.String
	data.IDSkemaLaporanTeknik = idSkemaLaporanTeknik.Int64

	data.CreatedAt = &createdAt.Time
	data.CreatedBy = createdBy.Int64
	data.UpdatedAt = &updatedAt.Time
	data.UpdatedBy = updatedBy.Int64
	data.DeletedAt = &deletedAt.Time
	data.DeletedBy = deletedBy.Int64

	return &data, 1, nil
}

func (v *repositories) DeleteOneByID(id int64) (int64, error) {
	query := "DELETE FROM emonica_v2.hk_pengiriman WHERE id = ?"

	res, err := v.DB.Exec(v.DB.Rebind(query), id)
	if err != nil {
		return -1, err
	}
	ra, _ := res.RowsAffected()

	return ra, nil
}

func (v *repositories) GetAll(filter *request.QueryParameter) ([]*model.HKPengiriman, int64, error) {
	var (
		searchToInt64, _ = strconv.ParseInt(filter.Search, 10, 64)
		namedQuery       = map[string]interface{}{
			"id":                      searchToInt64,
			"kode_jadwal":             filter.Search,
			"deskripsi":               filter.Search,
			"id_skema_laporan_teknik": filter.Search,
		}

		limitOffsetQuery = fmt.Sprintf(
			"LIMIT %s OFFSET %s",
			strconv.FormatInt(filter.Limit, 10),
			strconv.FormatInt(filter.Offset, 10),
		)
	)

	var (
		columns    = "id, kode_jadwal, deskripsi, id_skema_laporan_teknik, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by"
		mainQuery  = fmt.Sprintf("SELECT %s FROM emonica_v2.hk_pengiriman ", columns)
		countQuery = fmt.Sprintf("SELECT %s FROM emonica_v2.hk_pengiriman ", "COUNT(*) as count")

		firstConcat = "WHERE"
	)

	if filter.Search != "" {
		searchQuery := fmt.Sprintf("%s %s", firstConcat, `id = :id 
		OR kode_jadwal LIKE '%' || :kode_jadwal || '%' 
		OR deskripsi LIKE '%' || :deskripsi || '%' 
		OR id_skema_laporan_teknik = :id_skema_laporan_teknik`)

		mainQuery += searchQuery
		countQuery += searchQuery

		firstConcat = "AND"
	}

	query := fmt.Sprintf("%s %s %s %s", mainQuery, filter.FormatColumnFilter(firstConcat, "AND"), filter.FormatSort("ORDER BY"), limitOffsetQuery)
	rows, err := v.DB.NamedQuery(v.DB.Rebind(query), namedQuery)
	if err != nil {
		return nil, -1, err
	}
	defer rows.Close()

	var (
		result = make([]*model.HKPengiriman, 0)
	)

	for rows.Next() {
		var (
			data                 model.HKPengiriman
			id                   sql.NullInt64
			kodeJadwal           sql.NullString
			deskripsi            sql.NullString
			idSkemaLaporanTeknik sql.NullInt64

			createdAt sql.NullTime
			createdBy sql.NullInt64
			updatedAt sql.NullTime
			updatedBy sql.NullInt64
			deletedAt sql.NullTime
			deletedBy sql.NullInt64
		)

		err := rows.Scan(&id, &kodeJadwal, &deskripsi, &idSkemaLaporanTeknik, &createdAt, &createdBy, &updatedAt, &updatedBy, &deletedAt, &deletedBy)
		if err != nil {
			return nil, -1, err
		}

		data.ID = id.Int64
		data.KodeJadwal = kodeJadwal.String
		data.Deskripsi = deskripsi.String
		data.IDSkemaLaporanTeknik = idSkemaLaporanTeknik.Int64

		data.CreatedAt = &createdAt.Time
		data.CreatedBy = createdBy.Int64
		data.UpdatedAt = &updatedAt.Time
		data.UpdatedBy = updatedBy.Int64
		data.DeletedAt = &deletedAt.Time
		data.DeletedBy = deletedBy.Int64

		result = append(result, &data)
	}

	// ============================================== count data
	var count sql.NullInt64
	rowsCount, err := v.DB.NamedQuery(v.DB.Rebind(fmt.Sprintf("%s %s", countQuery, filter.FormatColumnFilter(firstConcat, "AND"))), namedQuery)
	if err != nil {
		return nil, -1, err
	}
	defer rowsCount.Close()

	for rowsCount.Next() {
		err := rowsCount.Scan(&count)
		if err != nil {
			return nil, -1, err
		}
	}
	// =========================================================

	return result, count.Int64, nil
}

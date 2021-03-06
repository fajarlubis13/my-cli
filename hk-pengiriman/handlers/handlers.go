package handlers

import (
	"hk-pengiriman/helpers/request"
	"hk-pengiriman/helpers/response"
	"hk-pengiriman/model"
	"hk-pengiriman/usecases"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// Handlers ...
type Handlers interface {
	CreateOne(c *gin.Context)
	UpdateOneByID(c *gin.Context)
	GetOneByID(c *gin.Context)
	DeleteOneByID(c *gin.Context)
	GetAll(c *gin.Context)
}

type handlers struct {
	ucase usecases.Usecases
}

// NewHandlers ...
func NewHandlers() Handlers {
	return &handlers{
		ucase: usecases.NewUsecases(),
	}
}

func (m *handlers) CreateOne(c *gin.Context) {
	var (
		data = model.HKPengiriman{}
		resp = &response.Response{}
	)
	defer resp.Serve(c)

	if err := c.ShouldBindJSON(&data); err != nil {
		resp.Err = err
		return
	}

	ra, err := m.ucase.CreateOne(&data)
	if err != nil {
		resp.Err = err
		return
	}

	resp.Body.Count = ra
	resp.Body.Payload = data
}

func (m *handlers) UpdateOneByID(c *gin.Context) {
	var (
		data = model.HKPengiriman{}
		resp = &response.Response{}
	)
	defer resp.Serve(c)

	if err := c.ShouldBindJSON(&data); err != nil {
		resp.Err = err
		return
	}

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	ra, err := m.ucase.UpdateOneByID(id, &data)
	if err != nil {
		resp.Err = err
		return
	}

	resp.Body.Count = ra
	resp.Body.Payload = data
}

func (m *handlers) GetOneByID(c *gin.Context) {
	var (
		resp = &response.Response{}
	)
	defer resp.Serve(c)

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	data, count, err := m.ucase.GetOneByID(id)
	if err != nil {
		resp.Err = err
		return
	}

	resp.Body.Count = count
	resp.Body.Payload = data
}

func (m *handlers) DeleteOneByID(c *gin.Context) {
	var (
		resp = &response.Response{}
	)
	defer resp.Serve(c)

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	ra, err := m.ucase.DeleteOneByID(id)
	if err != nil {
		resp.Err = err
		return
	}

	resp.Body.Count = ra
	resp.Body.Payload = ra
}

func (m *handlers) GetAll(c *gin.Context) {
	var (
		resp   = &response.Response{}
		filter = &request.QueryParameter{}
	)
	defer resp.Serve(c)

	filter.Search = c.Query("search")
	filter.Limit, _ = strconv.ParseInt(c.DefaultQuery("limit", "10"), 10, 64)
	filter.Offset, _ = strconv.ParseInt(c.DefaultQuery("offset", "0"), 10, 64)
	filter.SortBy = strings.Split(c.Query("sortby"), ",")

	filter.AssignColumnFilter(
		[]request.ColumnFilter{
			request.ColumnFilter{Column: "id_pengiriman", Criteria: c.QueryMap("id_pengiriman")},
			request.ColumnFilter{Column: "id_surat", Criteria: c.QueryMap("id_surat")},
		},
	)

	data, count, err := m.ucase.GetAll(filter)
	if err != nil {
		resp.Err = err
		return
	}

	resp.Body.Count = count
	resp.Body.Payload = data
}

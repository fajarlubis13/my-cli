package response

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// +===========================================================================+
// | Manage Response Body                                                      |
// +===========================================================================+

// Status ...
type Status string

// Body ...
type Body struct {
	StatusCode    int         `json:"status_code"`
	StatusMessage Status      `json:"status_message"`
	Description   string      `json:"description"`
	Count         int64       `json:"count"`
	Offset        int64       `json:"offset"`
	Limit         int64       `json:"limit"`
	Href          string      `json:"href"`
	Payload       interface{} `json:"payload"`
}

// Response ...
type Response struct {
	Body Body
	Err  error
}

const (
	// Error ...
	Error Status = "Error"

	// Success ...
	Success Status = "Success"
)

// Serve ...
func (c *Response) Serve(ctx *gin.Context) error {
	defer func() {
		ctx.JSON(http.StatusOK, c.Body)
	}()

	c.Body.Href = ctx.Request.RequestURI

	if c.Err != nil {
		c.Body.StatusMessage = Error
		c.Body.Description = c.Err.Error()
		c.Body.StatusCode = 400

		return c.Err
	}

	c.Body.StatusMessage = Success
	c.Body.StatusCode = 200
	c.Body.Limit, _ = strconv.ParseInt(ctx.DefaultQuery("limit", "10"), 10, 64)
	c.Body.Offset, _ = strconv.ParseInt(ctx.DefaultQuery("offset", "0"), 10, 64)

	return nil
}

// +===========================================================================+

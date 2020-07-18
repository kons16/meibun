package web

import (
	"github.com/labstack/echo"
	"net/http"
)

// getUserHandler は　GET /users/:id に対応
func (s *server) getUserHandler(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
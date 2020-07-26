package web

import (
	"github.com/kons16/meibun/api-server/model"
	"github.com/labstack/echo"
	"net/http"
)

// tokenからユーザー情報(id, name)を取得する
func(s *server) findUser(c echo.Context) *model.User {
	cookie, err := c.Cookie(sessionKey)
	if err == nil && cookie.Value != "" {
		user, _ := s.app.FindUserByToken(cookie.Value)
		return user
	}
	return nil
}

// getUserHandler は　GET /users/:id に対応
func (s *server) getUserHandler(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
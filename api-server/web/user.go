package web

import (
	"github.com/kons16/meibun/api-server/model"
	"github.com/labstack/echo"
)

// tokenからユーザー情報(id, name)を取得する
func (s *server) findUser(c echo.Context) *model.User {
	cookie, err := c.Cookie(sessionKey)
	if err == nil && cookie.Value != "" {
		user, _ := s.app.FindUserByToken(cookie.Value)
		return user
	}
	return nil
}

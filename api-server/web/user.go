package web

import (
	"errors"
	"github.com/kons16/meibun/api-server/model"
	"github.com/labstack/echo"
	"net/http"
)

// tokenからユーザー情報(id, name)を取得する
func(s *server) findUser(c echo.Context) (*model.User, error) {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return nil, errors.New("authHeader is null")
	}

	user, err := s.app.FindUserByToken(authHeader)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// getUserHandler は　GET /users/:id に対応
func (s *server) getUserHandler(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
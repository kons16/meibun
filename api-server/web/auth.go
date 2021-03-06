package web

import (
	"github.com/labstack/echo"
	"net/http"
	"time"
)

// willSignupHandler は GET /signup に対応
func (s *server) willSignupHandler(c echo.Context) error {
	user := s.findUser(c)

	data := map[string]interface{}{
		"Name": sessionKey,
		"User": user,
	}
	return c.JSON(http.StatusOK, data)
}

// signupHandler は POST /signup に対応
func (s *server) signupHandler(c echo.Context) error {
	params := new(struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	})
	c.Bind(params)

	if err := s.app.CreateNewUser(params.Name, params.Email, params.Password); err != nil {
		return err
	}

	user, err := s.app.FindUserByEmail(params.Email)
	if err != nil {
		return err
	}

	expiresAt := time.Now().Add(24 * time.Hour)
	token, err := s.app.CreateNewToken(user.ID, expiresAt)
	if err != nil {
		return err
	}

	var layout = "2006-01-02 15:04:05"
	return c.JSON(http.StatusOK, map[string]string{
		"Name": sessionKey, "token": token, "expiresAt": expiresAt.Format(layout)})
}

// signoutHandler は POST /signout に対応
func (s *server) signoutHandler(c echo.Context) error {
	/*
		TODO: user_sessionから該当レコードの削除処理
		cookie, err := c.Cookie(sessionKey)
		if err == nil && cookie.Value != "" {
		}
	*/

	return c.JSON(http.StatusOK, map[string]string{
		"Name": sessionKey})
}

// willSigninHandler は　GET /signin に対応
func (s *server) willSigninHandler(c echo.Context) error {
	user := s.findUser(c)

	data := map[string]interface{}{
		"Name": sessionKey,
		"User": user,
	}
	return c.JSON(http.StatusOK, data)
}

// signinHandler は　POST /signin に対応
func (s *server) signinHandler(c echo.Context) error {
	params := new(struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	})
	c.Bind(params)

	_, err := s.app.LoginUser(params.Email, params.Password)
	if err != nil {
		return err
	}

	user, err := s.app.FindUserByEmail(params.Email)
	if err != nil {
		return err
	}

	expiresAt := time.Now().Add(24 * time.Hour)
	token, err := s.app.CreateNewToken(user.ID, expiresAt)
	if err != nil {
		return err
	}

	var layout = "2006-01-02 15:04:05"
	return c.JSON(http.StatusOK, map[string]string{
		"Name": sessionKey, "token": token, "expiresAt": expiresAt.Format(layout)})
}

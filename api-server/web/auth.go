package web

import (
	"github.com/labstack/echo"
	"net/http"
	"time"
)

// testHandler は GET /test に対応
func (s *server) testHandler(c echo.Context) error {
	msg := map[string]string{
		"message": "Hello!",
	}
	return c.JSON(http.StatusOK, msg)
}

// willSignupHandler は GET /signup に対応
func (s *server) willSignupHandler(c echo.Context) error {
	return nil
}

// signupHandler は POST /signup に対応
func (s *server) signupHandler(c echo.Context) error {
	params := new(struct {
		Name		string `form:name`
		Email		string `form:email`
		Password	string `form:password`
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

	c.SetCookie(&http.Cookie{
		Name:	sessionKey,
		Value:	token,
		Expires: expiresAt,
	})

	msg := map[string]string{
		"message": "Hello!",
	}
	return c.JSON(http.StatusOK, msg)
}

func (s *server) signoutHandler(c echo.Context) error {
	return nil
}

func (s *server) willSigninHandler(c echo.Context) error {
	return nil
}

func (s *server) signinHandler(c echo.Context) error {
	return nil
}


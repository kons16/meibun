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
	// return c.JSON(http.StatusOK, nil)
	return c.Render(http.StatusOK, "signup.html", nil)
}

// signupHandler は POST /signup に対応
func (s *server) signupHandler(c echo.Context) error {
	params := new(struct {
		Name		string `json:"name" form:"name"`
		Email		string `json:"email" form:"email"`
		Password	string `json:"password" form:"password"`
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

	return c.Redirect(http.StatusFound,"/")
}

// signoutHandler は POST /signup に対応
func (s *server) signoutHandler(c echo.Context) error {
	c.SetCookie(&http.Cookie{
		Name:	sessionKey,
		Value:	"",
		Expires: time.Unix(0, 0),
	})
	return c.Redirect(http.StatusFound, "/")
}

// willSigninHandler は　GET /signin に対応
func (s *server) willSigninHandler(c echo.Context) error {
	// return c.JSON(http.StatusOK, nil)
	return c.Render(http.StatusOK, "signin.html", nil)
}

// signinHandler は　POST /signin に対応
func (s *server) signinHandler(c echo.Context) error {
	params := new(struct {
		Email		string `json:"email" form:"email"`
		Password	string `json:"password" form:"password"`
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
	c.SetCookie(&http.Cookie{
		Name:	sessionKey,
		Value:	token,
		Expires: expiresAt,
	})

	return c.Redirect(http.StatusFound, "/")
}

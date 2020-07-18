package web

import (
	"github.com/kons16/meibun/api-server/service"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"time"
)

const sessionKey = ""

type Server interface {
	Handler() *echo.Echo
}

type server struct {
	app service.MeibunApp
}

func NewServer(app service.MeibunApp) Server {
	return &server{app: app}
}

func (s *server) Handler() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "DENY",
		HSTSMaxAge:            3600,
		ContentSecurityPolicy: "default-src 'self'",
	}))

	e.GET("/test", s.testHandler)
	e.GET("/signup", s.willSignupHandler)
	e.POST("/signup", s.signupHandler)
	e.GET("/signin", s.willSigninHandler)
	e.POST("/signin", s.signinHandler)
	e.POST("/signout", s.signoutHandler)

	return e
}

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
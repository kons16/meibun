package web

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const sessionKey = ""

type Server interface {
	Handler() *echo.Echo
}

type server struct {
}

func NewServer() {
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

	e.GET("/", s.indexHandler)
	e.GET("/signup", s.willSignupHandler)
	e.POST("/signup", s.signupHandler)
	e.GET("/signin", s.willSigninHandler)
	e.POST("/signin", s.signinHandler)
	e.POST("/signout", s.signoutHandler)

	return e
}


func (s *server) indexHandler(c echo.Context) error {
	return nil
}

func (s *server) willSignupHandler(c echo.Context) error {
	return nil
}

func (s *server) signupHandler(c echo.Context) error {
	return nil
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
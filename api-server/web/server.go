package web

import (
	"github.com/kons16/meibun/api-server/service"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
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

	e.GET("/index", s.indexHandler)
	e.GET("/signup", s.willSignupHandler)
	e.POST("/signup", s.signupHandler)
	e.GET("/signin", s.willSigninHandler)
	e.POST("/signin", s.signinHandler)
	e.POST("/signout", s.signoutHandler)

	return e
}

// jsonのmessageを返す
func (s *server) indexHandler(c echo.Context) error {
	msg := map[string]string{
		"message": "Hello!",
	}
	return c.JSON(http.StatusOK, msg)
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
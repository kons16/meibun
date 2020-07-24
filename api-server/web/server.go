package web

import (
	"github.com/kons16/meibun/api-server/service"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

const sessionKey = "TEST_SESSION_KEY"

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
	// ここはあとで見直す
	e.Use(middleware.CORS())

	e.GET("/", s.indexHandler)
	e.GET("/test", s.testHandler)
	e.GET("/signup", s.willSignupHandler)
	e.POST("/signup", s.signupHandler)
	e.GET("/signin", s.willSigninHandler)
	e.POST("/signin", s.signinHandler)
	e.POST("/signout", s.signoutHandler)

	e.GET("/users/:id", s.getUserHandler)

	return e
}

// indexHandler は GET / に対応
func (s *server) indexHandler(c echo.Context) error {
	user, err := s.findUser(c)
	// authHeaderが空
	if err != nil {
		user = nil
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"User":    user,
	})
}

// testHandler は GET /test に対応
func (s *server) testHandler(c echo.Context) error {
	msg := map[string]string{
		"message": "Hello!",
	}
	return c.JSON(http.StatusOK, msg)
}

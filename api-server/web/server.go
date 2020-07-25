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

	// CORSの設定
	var allowedOrigins = []string{
		"http://localhost:3000",
	}
	e.Use(
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowCredentials: true,
			AllowOrigins:     allowedOrigins,
			AllowHeaders: []string{
				echo.HeaderAccessControlAllowHeaders,
				echo.HeaderContentType,
				echo.HeaderContentLength,
				echo.HeaderAcceptEncoding,
				echo.HeaderXCSRFToken,
				echo.HeaderAuthorization,
			},
			AllowMethods: []string{
				http.MethodGet,
				http.MethodPost,
				http.MethodDelete,
			},
			MaxAge: 86400,
		}),
	)

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			origin := c.Request().Header.Get(echo.HeaderOrigin)
			// 許可しているOriginの中で、リクエストヘッダのOriginと一致するものがあれば処理を継続
			for _, o := range allowedOrigins {
				if origin == o {
					return next(c)
				}
			}
			// 一致しているものがなかった場合は403(Forbidden)を返却する
			return c.String(http.StatusForbidden, "forbidden")
		}
	})


	e.GET("/check_user", s.checkUserHandler)
	e.GET("/test", s.testHandler)
	e.GET("/signup", s.willSignupHandler)
	e.POST("/signup", s.signupHandler)
	e.GET("/signin", s.willSigninHandler)
	e.POST("/signin", s.signinHandler)
	e.POST("/signout", s.signoutHandler)

	e.GET("/users/:id", s.getUserHandler)

	return e
}

// checkUserHandler は GET /check_user に対応
func (s *server) checkUserHandler(c echo.Context) error {
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

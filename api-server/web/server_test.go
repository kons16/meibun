package web

import (
	"github.com/golang/mock/gomock"
	"github.com/kons16/meibun/api-server/service/mock_service"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_GetSessionNameHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/get_session_name", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockService := mock_service.NewMockMeibunApp(ctrl)
	app := &server{app: mockService}
	_ = app.getSessionNameHandler(c)

	assert.Equal(t, http.StatusOK, rec.Code)
}

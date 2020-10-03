package web

import (
	"github.com/golang/mock/gomock"
	"github.com/kons16/meibun/api-server/model"
	"github.com/kons16/meibun/api-server/service/mock_service"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_FindUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_service.NewMockMeibunApp(ctrl)
	var user model.User
	user.ID = 1
	mockService.EXPECT().FindUserByToken("token").Return(&user, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/signup", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	req.AddCookie(&http.Cookie{Name: sessionKey, Value: "token"})
	c := e.NewContext(req, rec)

	app := &server{app: mockService}
	_ = app.findUser(c)

	assert.Equal(t, http.StatusOK, rec.Code)
}

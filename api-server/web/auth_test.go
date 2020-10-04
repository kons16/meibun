package web

import (
	"github.com/golang/mock/gomock"
	"github.com/kons16/meibun/api-server/model"
	"github.com/kons16/meibun/api-server/service/mock_service"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// GET /signup
func Test_WillSignupHandler(t *testing.T) {
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
	_ = app.willSignupHandler(c)

	assert.Equal(t, http.StatusOK, rec.Code)
}

// POST /signup
func Test_SignupHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_service.NewMockMeibunApp(ctrl)
	mockService.EXPECT().CreateNewUser("name", "a@a.com", "password").Return(nil)

	user := &model.User{}
	user.ID = uint(1)
	mockService.EXPECT().FindUserByEmail("a@a.com").Return(user, nil)
	mockService.EXPECT().CreateNewToken(uint(1), gomock.Any()).Return("token", nil)

	e := echo.New()
	userJSON := `{"name":"name","email":"a@a.com","password":"password"}`
	req := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	app := &server{app: mockService}
	r := app.signupHandler(c)

	body := rec.Body.String()
	// 両端の改行と{}を取り除く
	body = body[1 : len(body)-2]
	slice := strings.Split(body, ",")
	newBody := ""

	for _, str := range slice {
		key := strings.Split(str, ":")[0]
		value := strings.Split(str, ":")[1]
		if key != "\"expiresAt\"" {
			newBody += key + ": " + value
		}
	}

	if assert.NoError(t, r) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "\"Name\": \"TEST_SESSION_KEY\"\"token\": \"token\"", newBody)
	}
}

// POST /signout
func Test_SignoutHandler(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/signout", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_service.NewMockMeibunApp(ctrl)
	app := &server{app: mockService}
	r := app.signoutHandler(c)

	body := rec.Body.String()
	m := "{\"Name\":\"TEST_SESSION_KEY\"}\n"

	if assert.NoError(t, r) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, m, body)
	}
}

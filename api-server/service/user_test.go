package service

import (
	"github.com/golang/mock/gomock"
	"github.com/kons16/meibun/api-server/model"
	"github.com/kons16/meibun/api-server/repository/mock_repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_FindUserByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockRepository(ctrl)
	var user model.User
	mockRepo.EXPECT().FindUserByEmail("a@a.com").Return(&user, nil)

	app := &meibunApp{repo: mockRepo}
	u, _ := app.FindUserByEmail("a@a.com")

	assert.Equal(t, &user, u)
}

func Test_FindUserByToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockRepository(ctrl)
	var user model.User
	mockRepo.EXPECT().FindUserByToken("token").Return(&user, nil)

	app := &meibunApp{repo: mockRepo}
	u, _ := app.FindUserByToken("token")

	assert.Equal(t, &user, u)
}

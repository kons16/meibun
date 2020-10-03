package service

import (
	"github.com/golang/mock/gomock"
	"github.com/kons16/meibun/api-server/repository/mock_repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CreateNewBook(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockRepository(ctrl)
	mockRepo.EXPECT().CreateNewBook("sentence", "title", "author", 1, uint(1)).Return(nil)

	app := &meibunApp{repo: mockRepo}
	r := app.CreateNewBook("sentence", "title", "author", 1, uint(1))

	assert.Equal(t, nil, r)
}

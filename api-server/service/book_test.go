package service

import (
	"github.com/golang/mock/gomock"
	"github.com/kons16/meibun/api-server/model"
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

func Test_UpdateBook(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockRepository(ctrl)
	var book model.Book
	mockRepo.EXPECT().UpdateBook(&book).Return(nil)

	app := &meibunApp{repo: mockRepo}
	r := app.UpdateBook(&book)

	assert.Equal(t, nil, r)
}

func Test_GetAllBooksByUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockRepository(ctrl)
	var frontBook *[]model.FrontBook
	mockRepo.EXPECT().GetAllBooksByUserID(uint(1)).Return(frontBook, nil)

	app := &meibunApp{repo: mockRepo}
	r, _ := app.GetAllBooksByUserID(uint(1))

	assert.Equal(t, frontBook, r)
}

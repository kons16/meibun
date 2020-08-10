package web

import (
	"fmt"
	"github.com/kons16/meibun/api-server/model"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

// POST /post_book に対応
func (s *server) postBookHandler(c echo.Context) error {
	params := new(struct {
		Sentence	string	`json:"sentence"`
		Title		string	`json:"title"`
		Author		string	`json:"author"`
		Pages		int		`json:"pages"`
	})
	c.Bind(params)

	cookie, err := c.Cookie(sessionKey)
	if err == nil && cookie.Value != "" {
		user, _ := s.app.FindUserByToken(cookie.Value)
		err = s.app.CreateNewBook(params.Sentence, params.Title, params.Author, params.Pages, user.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

// POST /delete_book に対応
func (s *server) deleteBookHandler(c echo.Context) error {
	params := new(struct {
		BookID	int	`json:"bookID"`
	})
	c.Bind(params)

	bookID, _ := strconv.ParseUint(strconv.Itoa(params.BookID), 10, 32)

	cookie, err := c.Cookie(sessionKey)
	if err == nil && cookie.Value != "" {
		user, _ := s.app.FindUserByToken(cookie.Value)
		fmt.Println(uint(bookID))
		err = s.app.DeleteBookByBookID(uint(bookID), user.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

// GET /users/books に対応。 userIDに紐づくbooksレコードを全件取得。またそれに紐づくhart数も取得
func (s *server) getUserBooksHandler(c echo.Context) error {
	id, _ := strconv.ParseUint(c.QueryParam("id"), 10, 32)
	frontBooks, err := s.app.GetAllBooksByUserID(uint(id))
	if err != nil {
		return echo.ErrNotFound
	}

	data := map[string]interface{}{
		"Books":    frontBooks,
	}
	return c.JSON(http.StatusOK, data)
}

// POST /make_hart に対応。タップ後のハート数を返す
func (s *server) makeHartHandler(c echo.Context) error {
	params := new(struct {
		BookID	int	`json:"bookID"`
	})
	c.Bind(params)

	bookID, _ := strconv.ParseUint(strconv.Itoa(params.BookID), 10, 32)

	var hart int
	cookie, err := c.Cookie(sessionKey)
	if err == nil && cookie.Value != "" {
		user, _ := s.app.FindUserByToken(cookie.Value)
		hart, err = s.app.MakeHart(uint(bookID), user.ID)
		if err != nil {
			return err
		}
	}

	data := map[string]int{
		"hart": hart,
	}
	return c.JSON(http.StatusOK, data)
}

// POST /remove_hart に対応。ハート取り消し後のハート全件を返す
func (s *server) removeHartHandler(c echo.Context) error {
	params := new(struct {
		BookID	int	`json:"bookID"`
	})
	c.Bind(params)

	bookID, _ := strconv.ParseUint(strconv.Itoa(params.BookID), 10, 32)

	var myHartBooks *[]model.FrontBook
	var userID uint

	cookie, err := c.Cookie(sessionKey)
	if err == nil && cookie.Value != "" {
		user, _ := s.app.FindUserByToken(cookie.Value)
		myHartBooks, err = s.app.RemoveMyHart(uint(bookID), user.ID)
		userID = user.ID
		if err != nil {
			return err
		}
	}

	data := map[string]interface{}{
		"myHartBooks": myHartBooks,
		"ID": userID,
	}
	return c.JSON(http.StatusOK, data)
}

// GET /get_my_harts に対応。ハートしたbookとログインしているユーザーIDを返す
func (s *server) getMyHartsHandler(c echo.Context) error {
	var myHartBooks *[]model.FrontBook
	var userID uint

	cookie, err := c.Cookie(sessionKey)
	if err == nil && cookie.Value != "" {
		user, _ := s.app.FindUserByToken(cookie.Value)
		myHartBooks, err = s.app.GetMyHart(user.ID)
		userID = user.ID
		if err != nil {
			return err
		}
	}

	data := map[string]interface{}{
		"myHartBooks": myHartBooks,
		"ID": userID,
	}
	return c.JSON(http.StatusOK, data)
}

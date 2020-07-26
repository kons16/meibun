package web

import (
	"github.com/labstack/echo"
)

// POST /post_book に対応
func(s *server) postBookHandler(c echo.Context) error {
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

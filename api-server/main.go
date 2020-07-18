package main

import (
	"fmt"
	"github.com/kons16/meibun/api-server/repository"
	"github.com/kons16/meibun/api-server/service"
	"github.com/kons16/meibun/api-server/web"
	"os"
)

func main() {
	repo, err := repository.New("")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	app := service.NewApp(repo)

	s := web.NewServer(app)

	if err := s.Handler().Start(":" + "8000"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

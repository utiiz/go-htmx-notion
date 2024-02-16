package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/utiiz/go/notion/internal/database"
	"github.com/utiiz/go/notion/internal/templates/components"
)

func main() {
	db, err := database.OpenDB()
	if err != nil {
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return
	}

	go func() {
		user, _ := database.GetUsers(db)
		fmt.Println(user.Email)
	}()

	e := echo.New()

	e.Static("/css", "internal/static/css")
	e.Static("/js", "internal/static/js")

	e.GET("/", func(c echo.Context) error {
		user, _ := database.GetUsers(db)
		return components.User(user).Render(c.Request().Context(), c.Response())
	})

	e.Logger.Fatal(e.Start(":1234"))
}

// func main() {
// 	db, err := database.OpenDB()
// 	if err != nil {
// 		return
// 	}
// 	defer db.Close()
//
// 	err = db.Ping()
// 	if err != nil {
// 		return
// 	}
//
// 	go func() {
// 		user, _ := database.GetUsers(db)
// 		fmt.Println(user.Email)
// 	}()
//
// 	mux := http.NewServeMux()
//
// 	fs := http.FileServer(http.Dir("./internal/static/"))
// 	mux.Handle("/static/", http.StripPrefix("/static", fs))
//
// 	component := pages.Base()
// 	mux.Handle("/", templ.Handler(component))
//
// 	err = http.ListenAndServe(":1234", mux)
// 	if err != nil {
// 		fmt.Printf("%v\n", err)
// 		return
// 	}
// }

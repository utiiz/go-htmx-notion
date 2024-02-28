package main

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/utiiz/go/notion/internal/database"
	"github.com/utiiz/go/notion/internal/templates/components"
	"github.com/utiiz/go/notion/internal/templates/pages"
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

	user, _ := database.GetUser(db)
	fmt.Println(user)

	e := echo.New()

	e.Static("/css", "internal/static/css")
	e.Static("/js", "internal/static/js")

	e.GET("/", func(c echo.Context) error {
		projects, _ := database.GetProjects(db)
		return pages.Base(projects).Render(c.Request().Context(), c.Response())
	})

	e.GET("/toggle-projects", func(c echo.Context) error {
		isOpen, err := strconv.ParseBool(c.QueryParam("isOpen"))
		if err != nil {
			return err
		}
		projects, _ := database.GetProjects(db)
		return components.Foldable("Projects", projects, !isOpen, nil, "/toggle-projects").Render(c.Request().Context(), c.Response())
	})

	e.GET("/toggle-favorites", func(c echo.Context) error {
		isOpen, err := strconv.ParseBool(c.QueryParam("isOpen"))
		if err != nil {
			return err
		}
		projects, _ := database.GetFavoriteProjects(db, user)
		return components.Foldable("Favorites", projects, !isOpen, nil, "/toggle-favorites").Render(c.Request().Context(), c.Response())
	})

	e.GET("/select-project/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}
		projects, _ := database.GetProjects(db)
		return components.Foldable("Projects", projects, true, &id, "/toggle-projects").Render(c.Request().Context(), c.Response())
	})

	e.GET("/select-project-page/:project-id/:id", func(c echo.Context) error {
		projectID, err := strconv.Atoi(c.Param("project-id"))
		if err != nil {
			return err
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}
		project, _ := database.GetProject(db, projectID)
		return components.Project(*project, true, id).Render(c.Request().Context(), c.Response())
	})

	e.Logger.Fatal(e.Start(":3000"))
}

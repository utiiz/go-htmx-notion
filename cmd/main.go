package main

import (
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/utiiz/go/notion/internal/database"
	"github.com/utiiz/go/notion/internal/models"
	"github.com/utiiz/go/notion/internal/templates/components"
	"github.com/utiiz/go/notion/internal/templates/layouts"
	"github.com/utiiz/go/notion/internal/templates/pages"
)

func GetProjects(db *sqlx.DB, user *models.User, parent string) (*[]models.Project, error) {
	switch parent {
	case "favorites":
		return database.GetFavoriteProjects(db, user)
	default:
		return database.GetProjects(db)
	}
}

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
	e.Static("/assets", "internal/static/assets")

	e.GET("/", func(c echo.Context) error {
		projects, _ := database.GetProjects(db)
		return pages.Base(projects, nil).Render(c.Request().Context(), c.Response())
	})

	e.GET("/toggle-:parent", func(c echo.Context) error {
		parent := c.Param("parent")
		isOpen, err := strconv.ParseBool(c.QueryParam("isOpen"))
		if err != nil {
			return err
		}

		projects, _ := GetProjects(db, user, parent)

		return components.Foldable(parent, projects, !isOpen, nil).Render(c.Request().Context(), c.Response())
	})

	e.GET("/toggle-:parent/:id", func(c echo.Context) error {
		parent := c.Param("parent")
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}

		projects, _ := GetProjects(db, user, parent)
		project, _ := database.GetProject(db, id)

		return components.Foldable(parent, projects, true, project).Render(c.Request().Context(), c.Response())
	})

	e.GET("/select-:parent/:project-id/:id", func(c echo.Context) error {
		parent := c.Param("parent")
		projectID, err := strconv.Atoi(c.Param("project-id"))
		if err != nil {
			return err
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}
		project, _ := database.GetProject(db, projectID)
		return components.Project(*project, true, id, parent).Render(c.Request().Context(), c.Response())
	})

	e.GET("/show/:project-id/:id", func(c echo.Context) error {
		projectID, err := strconv.Atoi(c.Param("project-id"))
		if err != nil {
			return err
		}
		project, _ := database.GetProject(db, projectID)
		return layouts.Main(project).Render(c.Request().Context(), c.Response())
	})

	e.GET("/close-:parent", func(c echo.Context) error {
		parent := c.Param("parent")
		return components.Foldable(parent, nil, false, nil).Render(c.Request().Context(), c.Response())
	})

	e.GET("/:project/:page", func(c echo.Context) error {
		project_url := c.Param("project")
		// page := c.Param("page")

		projects, _ := database.GetProjects(db)
		project, _ := database.GetProjectByURL(db, project_url)
		return pages.Base(projects, project).Render(c.Request().Context(), c.Response())
	})

	e.Logger.Fatal(e.Start(":3000"))
}

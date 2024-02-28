package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/utiiz/go/notion/internal/models"
	"github.com/utiiz/go/notion/internal/utils"
)

var DRIVER_NAME = "postgres"
var HOST = "localhost"
var PORT = 5432

func OpenDB() (*sqlx.DB, error) {
	env, err := utils.GetEnv()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf(`%s://%s:%s@%s:%d/%s?sslmode=disable`, DRIVER_NAME, env.DB_USER, env.DB_PASS, HOST, PORT, env.DB_NAME)
	db, err := sqlx.Open(DRIVER_NAME, url)
	if err != nil {
		fmt.Println("ERROR HERE")
		return nil, err
	}

	return db, nil
}

func GetUser(db *sqlx.DB) (*models.User, error) {
	user := models.User{}
	query := `SELECT * FROM "USER" LIMIT 1`

	err := db.Get(&user, query)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, err
	}

	fmt.Printf("%v\n", user)

	return &user, nil
}

func GetProjects(db *sqlx.DB) (*[]models.Project, error) {
	projects := []models.Project{}
	query := `SELECT * FROM "PROJECT"`

	err := db.Select(&projects, query)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, err
	}

	fmt.Printf("%v\n", projects)

	return &projects, nil
}

func GetFavoriteProjects(db *sqlx.DB, user *models.User) (*[]models.Project, error) {
	projects := []models.Project{}
	query := `SELECT p.* FROM "PROJECT" p LEFT JOIN "FAVORITE" f ON f.project_id = p.id WHERE f.user_id = $1`

	err := db.Select(&projects, query, user.ID)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, err
	}

	fmt.Printf("%v\n", projects)

	return &projects, nil
}

func GetProject(db *sqlx.DB, id int) (*models.Project, error) {
	project := models.Project{}
	query := `SELECT * FROM "PROJECT" WHERE id = $1`

	err := db.Get(&project, query, id)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, err
	}

	fmt.Printf("%v\n", project)

	return &project, nil
}

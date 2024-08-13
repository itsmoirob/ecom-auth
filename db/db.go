package db

import (
	"database/sql"
	"fmt"

	"github.com/itsmoirob/ecom-auth/config"
	_ "github.com/lib/pq"
)

func NewPostgresStore() (*sql.DB, error) {
	//  docker run --name catch-up -e POSTGRES_PASSWORD=mysecretpassword -d -p 5432:5432 -v postgres_volume:/var/lib/postgresql/data postgres
	// connStr := "user=postgres dbname=postgres password=mysecretpassword sslmode=disable"
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", config.Envs.PostgresUser, config.Envs.PostgresDbName, config.Envs.PostgresPassword)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

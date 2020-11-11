package client

import (
	"database/sql"
	"fmt"

	//Import for compatibility with Postgres
	_ "github.com/lib/pq"

	log "github.com/sirupsen/logrus"
	"github.com/whyeasy/zally-cleaner/internal"
)

//New Creates a new DB connection client.
func New(c internal.Config) *sql.DB {

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", c.User, c.Password, c.Host, c.Port, c.Database, c.SSLMode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("Connected to host: %s and database: %s", c.Host, c.Database)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)

	return db
}

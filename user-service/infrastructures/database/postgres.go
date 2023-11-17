package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/FianGumilar/microservices/user-service/config"
	_ "github.com/lib/pq"
)

func NewPostgres(conf *config.AppConfig) *sql.DB {
	host := conf.Postgres.Host
	port := conf.Postgres.Port
	user := conf.Postgres.User
	pass := conf.Postgres.Pass
	name := conf.Postgres.Name

	url := fmt.Sprintf("host=%s port=%s user=%s password=%s, name=%s", host, port, user, pass, name)

	db, err := sql.Open("postgres", url)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	log.Println("Connecting to database...")

	return db
}

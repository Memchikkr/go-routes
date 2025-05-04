package bootstrap

import (
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

type PostgresClient struct {
	DB *sqlx.DB
}

func NewPostgresClient(env *Env) *PostgresClient {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"postgres",
		env.DBPort,
		env.DBUser,
		env.DBPassword,
		env.DBName,
	)

	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connection to PostgreSQL: ", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("PostgreSQL is not responding: ", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return &PostgresClient{DB: db}
}

func (pc *PostgresClient) Close() {
	if err := pc.DB.Close(); err != nil {
		log.Fatal("Error disconnecting from the database: ", err)
	}
}

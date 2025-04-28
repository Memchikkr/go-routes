package migrations

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Migrator struct {
	migrationsPath string
	dbURL          string
}

func NewMigrator(dbHost, dbPort, dbUser, dbPassword, dbName string) *Migrator {
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser,
		dbPassword,
		"postgres",
		dbPort,
		dbName,
	)

	return &Migrator{
		migrationsPath: "file://migrations/versions",
		dbURL:          dbURL,
	}
}

func (m *Migrator) Up() error {
	mg, err := migrate.New(m.migrationsPath, m.dbURL)
	if err != nil {
		return fmt.Errorf("migration initialization error: %v", err)
	}
	defer mg.Close()

	if err := mg.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migration application error: %v", err)
	}
	log.Println("Migrations have been successfully applied")
	return nil
}

func (m *Migrator) Down() error {
	mg, err := migrate.New(m.migrationsPath, m.dbURL)
	if err != nil {
		return fmt.Errorf("migration initialization error: %v", err)
	}
	defer mg.Close()

	if err := mg.Steps(-1); err != nil {
		return fmt.Errorf("migration rollback error: %v", err)
	}

	log.Println("Migrations have been rolled out")
	return nil
}

func (m *Migrator) Create(name string) error {
	log.Printf("Creating a migration: %s", name)
	cmd := exec.Command("migrate", "create", "-ext", "sql", "-dir", "migrations", "-seq", name)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("migration creation error: %v", err)
	}

	log.Println("Migration files have been created")
	return nil
}

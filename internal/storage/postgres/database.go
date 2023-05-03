package postgres

import (
	"database/sql"

	"github.com/Nicholas2012/task_tz/internal/storage/postgres/migrations"
	goose "github.com/pressly/goose/v3"
)

func ApplyMigrations(db *sql.DB) error {
	goose.SetBaseFS(migrations.Migrations)
	err := goose.Up(db, ".")
	return err
}

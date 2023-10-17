package migration

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func StartMigrations(db *sql.DB) error {
	err := goose.UpTo(db, "./internal/adapters/db/migration", 1)
	if err != nil {
		return err
	}

	err = goose.UpTo(db, "./internal/adapters/db/migration", 2)
	if err != nil {
		return err
	}

	err = goose.UpTo(db, "./internal/adapters/db/migration", 3)
	if err != nil {
		return err
	}

	err = goose.UpTo(db, "./internal/adapters/db/migration", 4)
	if err != nil {
		return err
	}

	return nil
}

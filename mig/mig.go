package mig

import (
	"database/sql"
	"embed"
	"fmt"
	"os"

	"github.com/delveper/mylib/app/ent"
	"github.com/pressly/goose/v3"
)

// go:embed *.sql
var FS embed.FS // will be used as mig.FS in main pkg

const (
	UpDirection   = "up"
	DownDirection = "down"
)

func Migrate(db *sql.DB, logger ent.Logger) error {
	goose.SetLogger(logger)

	goose.SetBaseFS(FS)
	defer func() {
		goose.SetBaseFS(nil)
	}()

	d := os.Getenv("DB_DIALECT")
	if err := goose.SetDialect(d); err != nil {
		return fmt.Errorf("error setting dialect: %w", err)
	}

	direction := os.Getenv("DB_MIGRATE")
	switch direction {
	case UpDirection:
		if err := goose.Up(db, "."); err != nil {
			return fmt.Errorf("error running mingrations up: %w", err)
		}
	case DownDirection:
		if err := goose.Down(db, "."); err != nil {
			return fmt.Errorf("error run mingrations down: %w", err)
		}
	default:
		return fmt.Errorf("wrong direction argument: %s", direction)
	}

	return nil
}

package postgres

import (
	"context"
	"fmt"

	errors "github.com/Red-Sock/trace-errors"
	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/sirupsen/logrus"
)

const (
	driver = "pgx"
	dir    = "./migrations"
)

type Config struct {
	Username     string
	Pwd          string
	Host         string
	Port         uint16
	DatabaseName string

	Migrate bool
}

func New(ctx context.Context, cfg Config) (TxManager, error) {
	connStr := CreateConnectionString(cfg)

	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, errors.Wrap(err, "error checking connection to redis")
	}

	if cfg.Migrate {
		err = migrate(connStr)
		if err != nil {
			return nil, errors.Wrap(err, "error applying migrations")
		}
	}

	return &tx{Conn: conn}, nil
}

func CreateConnectionString(cfg Config) string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		cfg.Username,
		cfg.Pwd,
		cfg.Host,
		cfg.Port,
		cfg.DatabaseName,
	)
}

func migrate(connString string) error {
	conn, err := goose.OpenDBWithDriver(driver, connString)
	if err != nil {
		return errors.Wrap(err, "error opening pg db connection")
	}

	goose.SetLogger(logrus.StandardLogger())

	err = goose.Up(conn, dir)
	if err != nil {
		return errors.Wrap(err, "error performing up")
	}

	return nil
}

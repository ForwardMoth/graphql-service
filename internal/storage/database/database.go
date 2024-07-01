package database

import (
	"errors"
	"fmt"
	"github.com/ForwardMoth/graphql-service/internal/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	DB *sqlx.DB
}

func Setup(dbCfg config.DBConfig) (*Database, error) {
	db, err := New(dbCfg)
	if err != nil {
		return nil, err
	}

	err = db.Migrate()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func New(dbCfg config.DBConfig) (*Database, error) {
	connectionURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbCfg.Host, dbCfg.Port, dbCfg.User, dbCfg.DBName, dbCfg.Password)

	conn, err := sqlx.Connect("postgres", connectionURL)
	if err != nil {
		return nil, fmt.Errorf("error with setting connection to database")
	}

	return &Database{DB: conn}, nil
}

func (db *Database) Migrate() error {
	driver, err := postgres.WithInstance(db.DB.DB, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"postgres", driver)
	if err != nil {
		return err
	}

	err = m.Up()

	if !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}

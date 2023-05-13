// Package adapters are the glue between components and external sources.
// # This manifest was generated by ymir. DO NOT EDIT.
package adapters

import (
	"testing"

	sqlEnt "entgo.io/ent/dialect/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/irhamsahbana/learn-ymir/pkg/infrastructure"
)

func TestWithEntdemoMySQL(t *testing.T) {
	is := assert.New(t)

	db, mock, err := sqlmock.New()
	if err != nil {
		is.Failf("failed to open stub db", "%v", err)
	}

	is.NotNil(db, "mock db is null")
	is.NotNil(mock, "sqlmock is null")

	EntdemoMySQLOpen = func(dialect, source string) (*sqlEnt.Driver, error) {
		return sqlEnt.NewDriver(dialect, sqlEnt.Conn{ExecQuerier: db}), nil
	}

	infrastructure.Configuration(
		infrastructure.WithPath("../.."),
		infrastructure.WithFilename("config.yaml"),
	).Initialize()

	adapter := &Adapter{}
	adapter.Sync(
		WithEntdemoMySQL(&EntdemoMySQL{
			NetworkDB: NetworkDB{
				Database:    infrastructure.Envs.EntdemoMySQL.Database,
				User:        infrastructure.Envs.EntdemoMySQL.User,
				Password:    infrastructure.Envs.EntdemoMySQL.Password,
				Host:        infrastructure.Envs.EntdemoMySQL.Host,
				Port:        infrastructure.Envs.EntdemoMySQL.Port,
				MaxIdleCons: infrastructure.Envs.DB.MaxIdleCons,
			},
		}),
	)

	mock.ExpectClose()

	// Asserts
	is.Nil(adapter.UnSync())
	is.Nil(mock.ExpectationsWereMet())
}


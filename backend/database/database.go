package database

import (
	"database/sql"
	"fmt"
	"todo/util"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var dbUser string = util.GetEnvironmentVariable("POSTGRES_USER")
var dbPassword string = util.GetEnvironmentVariable("POSTGRES_PASSWORD")
var dbName string = util.GetEnvironmentVariable("POSTGRES_DB")
var dbPort string = "5432"
var dbHost string = "db"
var dbDriverName = "postgres"

var migrationVersion uint = 1

const dbKey string = "database"

func InitDb() (*sql.DB, error) {
	sslDissable := "enable"
	if util.IsDebug {
		sslDissable = "disable"
	}
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbPort, dbUser, dbPassword, dbName, sslDissable)

	db, err := sql.Open(dbDriverName, connStr)

	if err != nil {
		return nil, err
	}

	instance, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	migrations, err := migrate.NewWithDatabaseInstance("file://./database/migrations/", dbDriverName, instance)
	if err != nil {
		return nil, err
	}

	// err = migrations.Migrate(migrationVersion)
	err = migrations.Force(int(migrationVersion))

	if err != nil && err != migrate.ErrNoChange {
		return nil, err
	}

	return db, nil
}

func DbMiddleware(db *sql.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		c.Set(dbKey, db)
	}
}

func GetDb(c *gin.Context) *sql.DB {
	return c.MustGet(dbKey).(*sql.DB)
}

package database

import (
	"database/sql"
	"fmt"
	"os"

	mysql "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	migrateMysql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type DatabaseConfig struct {
	DatabaseDsn           string `envconfig:"DATABASE_DSN" required:"true"`
	DatabasesMigrationUrl string `envconfig:"DATABASE_MIGRATIONS_URL" default:"file://database/migrations"`
}

func Connect() (*sqlx.DB, error) {
	config, err := loadConfig(".env", ".env.development")
	if err != nil {
		return nil, err
	}

	database, err := sqlx.Open("mysql", config.DatabaseDsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open a database connection: %w", err)
	}

	err = migrateDatabase(config, database.DB)

	return database, err
}

func loadConfig(filenames ...string) (*DatabaseConfig, error) {
	existingFiles := make([]string, 0)
	for _, filename := range filenames {
		if _, err := os.Stat(filename); err == nil {
			existingFiles = append(existingFiles, filename)
		}
	}

	if len(existingFiles) != 0 {
		err := godotenv.Load(existingFiles...)
		if err != nil {
			return nil, fmt.Errorf("failed to load env file: %w", err)
		}
	}

	var config DatabaseConfig
	err := envconfig.Process("", &config)
	if err != nil {
		return nil, fmt.Errorf("failed to load configuration for database: %w", err)
	}

	return &config, nil
}

func migrateDatabase(config *DatabaseConfig, instance *sql.DB) error {
	driver, err := migrateMysql.WithInstance(instance, &migrateMysql.Config{})
	if err != nil {
		return fmt.Errorf("failed to instantiate database driver for migration service: %w", err)
	}

	parsedDsn, err := mysql.ParseDSN(config.DatabaseDsn)
	if err != nil {
		return fmt.Errorf("failed to parse database dsn: %w", err)
	}

	migrateService, err := migrate.NewWithDatabaseInstance(config.DatabasesMigrationUrl, parsedDsn.DBName, driver)
	// migrate, err := migrate.New(config.DatabasesMigrationUrl, "msyql://"+config.DatabaseDsn)
	if err != nil {
		return fmt.Errorf("failed to instantiate database migration service: %w", err)
	}

	err = migrateService.Up()
	if err != migrate.ErrNoChange {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	return nil
}

package migrator

import (
	"github.com/dashroots/migrate"
	"github.com/dashroots/migrate/database/postgres"

	// load the GitHub source for the migrate library
	_ "github.com/dashroots/migrate/source/github"
)

// Run brings the database up to the specified migration
func Run(repo string, connection string, version uint) (uint, error) {
	postgres.DefaultMigrationsTable = "migrations"

	m, err := migrate.New(repo, connection)

	if err != nil {
		return 0, err
	}

	if version != 0 {
		err = m.Migrate(version)
	} else {
		err = m.Up()
	}

	if err != nil && err != migrate.ErrNoChange {
		return 0, err
	}

	v, _, err := m.Version()

	if err != nil && err != migrate.ErrNilVersion {
		return 0, err
	}

	return v, nil
}

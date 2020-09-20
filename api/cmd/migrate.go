package cmd

import (
	"github.com/ainsleyclark/verbis/api/environment"
	"github.com/ainsleyclark/verbis/api/helpers/paths"
	"fmt"
	sqlMigrate "github.com/ainsleyclark/golang-sql-migrate"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	m *sqlMigrate.Migrate
	migrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "Runs database migrations and setups up the database.",
		Long:  `Migrate will run all the database migrations, IMplement`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if err := load(); err != nil {
				log.Panic(err)
			}
		},
	}
)

func init() {
	makeCmd.Flags().String("name", "n", "Add name to migration")
	migrateCmd.AddCommand(dropCreateCmd)
	migrateCmd.AddCommand(freshCmd)
	migrateCmd.AddCommand(makeCmd)
	migrateCmd.AddCommand(rollbackCmd)
	migrateCmd.AddCommand(upCmd)
}

// Load the migration package using the database.
func load() error {

	migrator, err := sqlMigrate.NewInstance(
		app.db.Sqlx.DB,
		environment.Env.DbDatabase,
		paths.Migration(),
		true,
	)
	if err != nil {
		return err
	}

	m = migrator

	return nil
}


// downCmd represents the down command
var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Down will drop the database entirely.",
	Long: `Down will drop the database entirely, and not recreate it,
nor will it run any pending migrations or recreate the migration 
table.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := m.Down(); err != nil {
			return err
		}
		return nil
	},
}

// dropCreateCmd represents the drop-create command
var dropCreateCmd = &cobra.Command{
	Use:   "drop-create",
	Short: "Drop & Create will drop the whole database and create it again.",
	Long: `Drop & Create will drop the whole database and create it again. 
Note it is not the same as fresh, as fresh will run all the up 
migrations over again.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := m.DropAndCreate(); err != nil {
			return err
		}
		return nil
	},
}

// freshCmd represents the fresh command
var freshCmd = &cobra.Command{
	Use:   "fresh",
	Short: "Drops database and runs all migrations.",
	Long: `Fresh drops the database entirely, creates it again 
and runs all pending migrations. Warning, run with caution if
application is in production.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := m.Fresh(); err != nil {
			return err
		}

		return nil
	},
}

// makeCmd represents the make command
var makeCmd = &cobra.Command{
	Use:   "make",
	Short: "Make will create up and down sql files based on the migration path.",
	Long: `Implement`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO Not working!
		migrationName, err := cmd.Flags().GetString("name")

		if err != nil {
			return err
		}

		if migrationName == "" {
			return fmt.Errorf("provide a migration name")
		}

		if err := m.Make(migrationName); err != nil {
			return err
		}
		return nil
	},
}

// rollbackCmd represents the roll back command
var rollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "Rollback will execute any pending files that are with the .down.sql extension.",
	Long: `Rollback will get the latest version in the database and 
execute any files that are with the .down.sql extension. If no 
.down.sql files are found it will not run.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := m.Rollback(); err != nil {
			return err
		}
		return nil
	},
}

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Up will run up migrations that have not yet been run..",
	Long: `Up looks at the currently active migration batch
number and will migrate all the way up (applying all up 
migrations). If a failure happens or a sql syntax error 
has been found, it will roll back any migrations that 
have already taken place.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := m.Up(); err != nil {
			return err
		}
		return nil
	},
}
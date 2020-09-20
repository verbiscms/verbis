package cmd

import (
	"cms/api/database/seeds"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	seedCmd = &cobra.Command{
		Use:   "seed",
		Short: "Run seeds for the database location in /api/database/seeds",
		Long: `Calling seed will run all of the seeds that are located within 
the seed path api/database/seeds. This will insert default records into the DB 
for users, options and any global configuration Verbis requires to function.`,
		RunE: func(cmd *cobra.Command, args []string) (error) {
			seeder := seeds.New(app.db.Sqlx, app.store)
			if err := seeder.Seed(); err != nil {
				return err
			}
			log.Info("Successfully seeded the database.")
			return nil
		},
	}
)

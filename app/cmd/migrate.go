package cmd

import (
	"github.com/blockfint/di-example-go/app/db"
	"github.com/blockfint/di-example-go/app/db/model"
	"github.com/blockfint/di-example-go/app/logger"
	"github.com/spf13/cobra"
)

var migrateDBCmd = &cobra.Command{
	Use:   "migrate-db",
	Short: "Migrate database",
	RunE: func(cmd *cobra.Command, args []string) error {
		lg := logger.New()

		gormDB, err := db.New(lg)
		if err != nil {
			lg.Errorf("Error Initializing DB: %+v", err)
			return err
		}

		migrations := model.NewMigrations()
		migrator := gormDB.Migrator()

		forceMigrate, _ := cmd.Flags().GetBool("force-migrate")
		if forceMigrate {
			for _, table := range migrations {
				if migrator.HasTable(table) {
					err = migrator.DropTable(table)
					if err != nil {
						lg.Errorf("DropTable error: %+v", err)
						return err
					}
				}
			}
		}

		for _, table := range migrations {
			if migrator.HasTable(table) {
				err = migrator.AutoMigrate(table)
				if err != nil {
					lg.Errorf("AutoMigrate error: %+v", err)
					return err
				}
			} else {
				err = migrator.CreateTable(table)
				if err != nil {
					lg.Errorf("CreateTable error: %+v", err)
					return err
				}
			}
		}

		if forceMigrate {
			lg.Info("Successfully force migrate-db")
		} else {
			lg.Info("Successfully migrate-db")
		}

		return nil
	},
}

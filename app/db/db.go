package db

import (
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Config = &gorm.Config{
	NamingStrategy: schema.NamingStrategy{
		TablePrefix:   "",
		SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
	},
}

func connectDB(lg *zap.SugaredLogger) (*gorm.DB, error) {
	var (
		dbDriver   = viper.GetString("DATABASE.DRIVER")
		dbHost     = viper.GetString("DATABASE.HOST")
		dbPort     = viper.GetString("DATABASE.PORT")
		dbUsername = viper.GetString("DATABASE.USERNAME")
		dbPassword = viper.GetString("DATABASE.PASSWORD")
		dbName     = viper.GetString("DATABASE.DB_NAME")
	)

	if dbDriver != "postgres" {
		lg.Fatalf("DB Driver: %s, is not implemented", dbDriver)
	}

	connString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s",
		dbHost, dbPort, dbUsername, dbName, dbPassword,
	)

	gormDB, err := gorm.Open(postgres.Open(connString), Config)
	if err != nil {
		lg.Errorf("gorm.Open error: %+v", err)
		return nil, err
	}

	err = gormDB.Exec("SELECT 1").Error
	if err != nil {
		lg.Errorf("Error connecting to database: %+v", err)
		return nil, err
	}

	return gormDB, nil
}

func New(lg *zap.SugaredLogger) (*gorm.DB, error) {
	gormDB, err := connectDB(lg)
	if err != nil {
		lg.Errorf("Error db.New(): %+v", err)
		return nil, err
	}

	return gormDB, nil
}

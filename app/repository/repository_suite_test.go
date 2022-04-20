package repository

import (
	"database/sql/driver"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/blockfint/di-example-go/app/db"
	"github.com/blockfint/di-example-go/app/logger"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gLogger "gorm.io/gorm/logger"
)

func TestRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repository Suite")
}

func getDBConfigWithSilentLogger(config gorm.Config) *gorm.Config {
	config.Logger = gLogger.Default.LogMode(gLogger.Silent)

	return &config
}

// Option 1: SQL query snapshot testing
// Use copyist to record the db query
// then compare the query in the running test with the recorded query
// This option requires a database connection (docker-compose.test.yml)
func newMockedGormDB() (*gorm.DB, error) {
	var (
		dbHost     = "localhost"
		dbPort     = "5432"
		dbUsername = "postgres"
		dbPassword = "blockfint@99"
		dbName     = "di-example-go-test"
	)

	connString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbHost, dbPort, dbUsername, dbName, dbPassword,
	)

	dialector := postgres.New(postgres.Config{
		DSN:                  connString,
		DriverName:           "copyist_postgres",
		PreferSimpleProtocol: true,
	})

	testConfig := getDBConfigWithSilentLogger(*db.Config)

	return gorm.Open(dialector, testConfig)
}

func newCustomerRepositoryWithMockedDB() (*CustomerRepository, error) {
	gormDB, err := newMockedGormDB()
	if err != nil {
		return nil, err
	}

	lg := logger.New()

	customerRepository := NewCustomerRepository(gormDB, lg)

	return customerRepository, nil
}

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

// Option 2: Test the query to match the manually written SQL query
// This option doesn't require any database connection
func newMockedDB() (*gorm.DB, sqlmock.Sqlmock, error) {
	pg, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 pg,
		PreferSimpleProtocol: true,
	})

	gormDB, err := gorm.Open(dialector, db.Config)

	return gormDB, mock, err
}

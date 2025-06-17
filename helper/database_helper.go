package helper

import (
	"cnb.cool/mliev/examples/go-web/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

var (
	db     *gorm.DB
	dbOnce sync.Once
)

// initDB initializes the database connection (private function)
func initDB() {
	var err error
	dbConfig := config.GetDatabaseConfig()

	var driver gorm.Dialector
	if dbConfig.Driver == "postgresql" {
		driver = postgres.New(
			postgres.Config{
				DSN:                  dbConfig.GetPostgreSQLDSN(),
				PreferSimpleProtocol: true, // disables implicit prepared statement usage
			})
	} else {
		// Default to MySQL
		driver = mysql.Open(dbConfig.GetMySQLDSN())
	}

	db, err = gorm.Open(driver, &gorm.Config{})

	if err != nil {
		Logger().Error(fmt.Sprintf("[db connect err:%s]", err.Error()))
		return
	}

	// Auto migrate the database schema using migration config
	migrationConfig := config.MigrationConfig{}
	migrationModels := migrationConfig.Get()

	if len(migrationModels) > 0 {
		err = db.AutoMigrate(migrationModels...)
		if err != nil {
			Logger().Error(fmt.Sprintf("[db migration err:%s]", err.Error()))
		} else {
			Logger().Info(fmt.Sprintf("[db migration success: %d models migrated]", len(migrationModels)))
		}
	}
}

// Database returns the singleton database instance
func Database() *gorm.DB {
	dbOnce.Do(initDB)
	return db
}

// GetDB returns the singleton database instance (alias for Database)
func GetDB() *gorm.DB {
	return Database()
}

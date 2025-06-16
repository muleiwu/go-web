package db

import (
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"mliev.com/template/go-web/config"
	"mliev.com/template/go-web/support/logger"
)

var (
	db     *gorm.DB
	dbOnce sync.Once
)

// initDB initializes the database connection (private function)
func initDB() {
	var err error
	dbConfig := config.GetDatabaseConfig()

	driver := mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName))

	if dbConfig.Driver == "postgresql" {
		dsn := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
			dbConfig.Username,
			dbConfig.Password,
			dbConfig.Host,
			dbConfig.Port,
			dbConfig.DBName)

		//logger.Get().Debug(fmt.Sprintf("[db connect to dsn: %s]", dsn))
		driver = postgres.New(
			postgres.Config{
				DSN:                  dsn,
				PreferSimpleProtocol: true, // disables implicit prepared statement usage
			})
	}

	db, err = gorm.Open(driver, &gorm.Config{})

	if err != nil {
		logger.Get().Error(fmt.Sprintf("[db connect err:%s]", err.Error()))
		return
	}

	// Auto migrate the database schema using migration config
	migrationModels := config.MigrationConfig{}.Get()
	if len(migrationModels) > 0 {
		err = db.AutoMigrate(migrationModels...)
		if err != nil {
			logger.Get().Error(fmt.Sprintf("[db migration err:%s]", err.Error()))
		} else {
			logger.Get().Info(fmt.Sprintf("[db migration success: %d models migrated]", len(migrationModels)))
		}
	}
}

// GetDB returns the singleton database instance
func GetDB() *gorm.DB {
	dbOnce.Do(initDB)
	return db
}

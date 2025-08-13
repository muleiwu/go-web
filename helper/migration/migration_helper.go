package migration

import (
	"fmt"

	"cnb.cool/mliev/examples/go-web/config"
	"cnb.cool/mliev/examples/go-web/helper"
)

// AutoMigrate performs database migration using the new architecture
func AutoMigrate() error {
	// Auto migrate the database schema using migration config
	migrationConfig := config.MigrationConfig{}
	migrationModels := migrationConfig.Get()

	if len(migrationModels) > 0 {
		err := helper.Database().AutoMigrate(migrationModels...)
		if err != nil {
			return fmt.Errorf("[db migration err:%s]", err.Error())
		}

		helper.Logger().Info(fmt.Sprintf("[db migration success: %d models migrated]", len(migrationModels)))
	}
	return nil
}

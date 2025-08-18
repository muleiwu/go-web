package migration

import (
	"fmt"

	"cnb.cool/mliev/examples/go-web/config"
	"cnb.cool/mliev/examples/go-web/internal/helper"
)

// AutoMigrate performs database migration using the new architecture
func AutoMigrate(helper *helper.Helper) error {
	// Auto migrate the database schema using migration config
	migrationConfig := config.Migration{}
	migrationModels := migrationConfig.Get()

	if len(migrationModels) > 0 {
		err := helper.GetDatabase().AutoMigrate(migrationModels...)
		if err != nil {
			return fmt.Errorf("[db migration err:%s]", err.Error())
		}

		helper.GetLogger().Info(fmt.Sprintf("[db migration success: %d models migrated]", len(migrationModels)))
	}
	return nil
}

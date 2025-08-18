package assembly

import (
	"sync"

	"cnb.cool/mliev/examples/go-web/helper/database"
	"cnb.cool/mliev/examples/go-web/internal/interfaces"
	"cnb.cool/mliev/examples/go-web/internal/pkg/database/impl"
)

type Database struct {
	Config interfaces.ConfigInterface
}

var (
	databaseOnce sync.Once
)

func (receiver Database) Assembly() {
	databaseOnce.Do(func() {
		database.DatabaseHelper = impl.NewDatabase(receiver.Config)
	})
}

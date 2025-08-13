package assembly

import (
	"sync"

	"cnb.cool/mliev/examples/go-web/helper/database"
	"cnb.cool/mliev/examples/go-web/internal/impl"
)

type Database struct {
}

var (
	databaseOnce sync.Once
)

func (receiver Database) Assembly() {
	databaseOnce.Do(func() {
		database.DatabaseHelper = impl.NewDatabase()
	})
}

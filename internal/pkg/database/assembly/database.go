package assembly

import (
	"sync"

	"cnb.cool/mliev/examples/go-web/internal/helper"
	"cnb.cool/mliev/examples/go-web/internal/pkg/database/impl"
)

type Database struct {
	Helper *helper.Helper
}

var (
	databaseOnce sync.Once
)

func (receiver *Database) Assembly() {
	databaseOnce.Do(func() {
		receiver.Helper.SetDatabase(impl.NewDatabase(receiver.Helper))
	})
}

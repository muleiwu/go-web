package assembly

import (
	"sync"

	"cnb.cool/mliev/examples/go-web/internal/interfaces"
	"cnb.cool/mliev/examples/go-web/internal/pkg/database/impl"
)

type Database struct {
	Helper interfaces.HelperInterface
}

var (
	databaseOnce sync.Once
)

func (receiver *Database) Assembly() {
	databaseOnce.Do(func() {
		receiver.Helper.SetDatabase(impl.NewDatabase(receiver.Helper))
	})
}

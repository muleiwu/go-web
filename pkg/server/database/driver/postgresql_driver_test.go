package driver

import (
	"testing"

	"cnb.cool/mliev/open/go-web/pkg/server/database/config"
	"gorm.io/driver/postgres"
)

func TestPostgresqlDialectorCarriesProtocolPreference(t *testing.T) {
	for _, testCase := range []struct {
		name                 string
		preferSimpleProtocol bool
	}{
		{name: "extended protocol"},
		{name: "simple protocol", preferSimpleProtocol: true},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			databaseConfig := &config.DatabaseConfig{
				Host: "127.0.0.1", Port: 5432, DBName: "test", Username: "test",
				PreferSimpleProtocol: testCase.preferSimpleProtocol,
			}
			dialector := newPostgresqlDialector(databaseConfig)
			postgresDialector, ok := dialector.(*postgres.Dialector)
			if !ok {
				t.Fatalf("dialector type = %T, want *postgres.Dialector", dialector)
			}
			if got := postgresDialector.Config.PreferSimpleProtocol; got != testCase.preferSimpleProtocol {
				t.Fatalf("PreferSimpleProtocol = %t, want %t", got, testCase.preferSimpleProtocol)
			}
		})
	}
}

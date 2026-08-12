package config

import (
	"strings"
	"testing"

	configImpl "cnb.cool/mliev/open/go-web/pkg/server/config/impl"
	mysqlDriver "github.com/go-sql-driver/mysql"
	"github.com/jackc/pgx/v5/pgconn"
)

func TestNewConfigDisablesSimpleProtocolByDefault(t *testing.T) {
	provider := configImpl.NewConfig()
	if got := NewConfig(provider).PreferSimpleProtocol; got {
		t.Fatal("PreferSimpleProtocol = true, want false by default")
	}
}

func TestNewConfigAllowsSimpleProtocolOptIn(t *testing.T) {
	provider := configImpl.NewConfig()
	provider.Set("database.prefer_simple_protocol", true)
	if got := NewConfig(provider).PreferSimpleProtocol; !got {
		t.Fatal("PreferSimpleProtocol = false, want explicit true")
	}
}

func TestMySQLDriverConfigDoesNotPromoteDBNameToParams(t *testing.T) {
	dc := &DatabaseConfig{
		Host:     "127.0.0.1",
		Port:     3306,
		DBName:   "dwz?allowAllFiles=true&parseTime=false",
		Username: "dwz",
		Password: "secret",
	}

	cfg, err := dc.MySQLDriverConfig()
	if err != nil {
		t.Fatalf("MySQLDriverConfig returned error: %v", err)
	}
	parsed, err := mysqlDriver.ParseDSN(cfg.FormatDSN())
	if err != nil {
		t.Fatalf("ParseDSN returned error: %v", err)
	}
	if parsed.AllowAllFiles {
		t.Fatal("allowAllFiles must stay disabled")
	}
	if parsed.DBName != dc.DBName {
		t.Fatalf("DBName was not preserved: %q", parsed.DBName)
	}
	if !parsed.ParseTime {
		t.Fatal("parseTime should come from trusted config, not injected DBName")
	}
}

func TestNormalizeTCPHostRejectsDSNBreakoutCharacters(t *testing.T) {
	_, err := NormalizeTCPHost("127.0.0.1)/dwz?allowAllFiles=true")
	if err == nil {
		t.Fatal("expected invalid host error")
	}
	if !strings.Contains(err.Error(), "invalid DSN characters") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestPostgreSQLDSNKeepsTimezoneSlashUnescaped(t *testing.T) {
	dc := &DatabaseConfig{
		Host:     "postgresql",
		Port:     5432,
		DBName:   "dwz",
		Username: "dwz",
		Password: "secret @",
	}

	dsn := dc.GetPostgreSQLDSN()
	if strings.Contains(dsn, "Asia%2FShanghai") {
		t.Fatalf("timezone slash was percent-encoded: %s", dsn)
	}
	if !strings.Contains(dsn, "TimeZone=Asia/Shanghai") {
		t.Fatalf("timezone query value missing from dsn: %s", dsn)
	}
	if !strings.Contains(dsn, "dwz:secret%20%40@") {
		t.Fatalf("username/password were not safely URL-encoded: %s", dsn)
	}

	cfg, err := pgconn.ParseConfig(dsn)
	if err != nil {
		t.Fatalf("pgconn.ParseConfig returned error: %v", err)
	}
	if got := cfg.RuntimeParams["TimeZone"]; got != "Asia/Shanghai" {
		t.Fatalf("unexpected parsed timezone: %q", got)
	}
}

package config

import (
	"strings"
	"testing"

	mysqlDriver "github.com/go-sql-driver/mysql"
)

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

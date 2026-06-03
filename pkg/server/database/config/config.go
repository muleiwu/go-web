package config

import (
	"fmt"
	"net"
	"net/netip"
	"net/url"
	"strconv"
	"strings"
	"time"
	"unicode"

	mysqlDriver "github.com/go-sql-driver/mysql"
	"github.com/muleiwu/gsr"
)

type DatabaseConfig struct {
	Driver   string `json:"driver"`
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DBName   string `json:"dbname"`
}

func NewConfig(config gsr.Provider) *DatabaseConfig {
	return &DatabaseConfig{
		Driver:   config.GetString("database.driver", "postgresql"),
		Host:     config.GetString("database.host", "127.0.0.1"),
		Port:     config.GetInt("database.port", 5432),
		DBName:   config.GetString("database.dbname", "test"),
		Username: config.GetString("database.username", "test"),
		Password: config.GetString("database.password", "123456"),
	}
}

func (dc *DatabaseConfig) GetMySQLDSN() string {
	cfg, err := dc.MySQLDriverConfig()
	if err != nil {
		return ""
	}
	return cfg.FormatDSN()
}

func (dc *DatabaseConfig) GetPostgreSQLDSN() string {
	host, err := NormalizeTCPHost(dc.Host)
	if err != nil {
		return ""
	}
	if err := ValidateTCPPort(dc.Port); err != nil {
		return ""
	}
	return (&url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(dc.Username, dc.Password),
		Host:     net.JoinHostPort(host, strconv.Itoa(dc.Port)),
		Path:     dc.DBName,
		RawQuery: "sslmode=disable&TimeZone=Asia/Shanghai",
	}).String()
}

func (dc *DatabaseConfig) MySQLDriverConfig() (*mysqlDriver.Config, error) {
	host, err := NormalizeTCPHost(dc.Host)
	if err != nil {
		return nil, err
	}
	if err := ValidateTCPPort(dc.Port); err != nil {
		return nil, err
	}

	cfg := mysqlDriver.NewConfig()
	cfg.User = dc.Username
	cfg.Passwd = dc.Password
	cfg.Net = "tcp"
	cfg.Addr = net.JoinHostPort(host, strconv.Itoa(dc.Port))
	cfg.DBName = dc.DBName
	cfg.ParseTime = true
	cfg.Loc = time.Local
	cfg.AllowAllFiles = false
	cfg.AllowCleartextPasswords = false
	cfg.AllowFallbackToPlaintext = false
	cfg.AllowOldPasswords = false
	if err := cfg.Apply(mysqlDriver.Charset("utf8mb4", "")); err != nil {
		return nil, err
	}
	return cfg, nil
}

func NormalizeTCPHost(host string) (string, error) {
	host = strings.TrimSpace(host)
	if host == "" {
		return "", fmt.Errorf("database host must not be empty")
	}
	if strings.HasPrefix(host, "[") && strings.HasSuffix(host, "]") {
		host = strings.TrimPrefix(strings.TrimSuffix(host, "]"), "[")
	}
	if strings.ContainsAny(host, "/\\?#@()") {
		return "", fmt.Errorf("database host contains invalid DSN characters")
	}
	for _, r := range host {
		if unicode.IsControl(r) || unicode.IsSpace(r) {
			return "", fmt.Errorf("database host contains invalid whitespace or control characters")
		}
	}
	if strings.Contains(host, ":") {
		if _, err := netip.ParseAddr(host); err != nil {
			return "", fmt.Errorf("invalid IPv6 database host")
		}
	}
	return host, nil
}

func ValidateTCPPort(port int) error {
	if port <= 0 || port > 65535 {
		return fmt.Errorf("database port must be between 1 and 65535")
	}
	return nil
}

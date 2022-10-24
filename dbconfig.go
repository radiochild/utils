package utils

import (
	"fmt"
	"strings"

	"github.com/jackc/pgx"
)

func NewDBConfig(logname, password, host, dbName string) (*pgx.ConnConfig, error) {
	connStr := ConnectionString(logname, password, host, dbName)
	connCfg, err := pgx.ParseConnectionString(connStr)
	return &connCfg, err
}

// postgresql://username:password@localhost/leap?sslmode=disable
func ConnectionString(username, password, host, dbName string) string {
	ux := []string{username, password}
	userInfo := strings.Join(ux, ":")
	baseStr := fmt.Sprintf("postgresql://%s@%s/%s", userInfo, host, dbName)
	if host == "localhost" {
		baseStr = fmt.Sprintf("%s?sslmode=disable", baseStr)
	}
	return baseStr
}

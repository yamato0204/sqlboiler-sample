package integration

import (
	"os"
)

type DBConfig struct {
	User     string
	DBName   string
	RootPass string
	Host     string
	Port     string
}

func NewDBConfig() *DBConfig {

	user := os.Getenv("MYSQL_USER_NAME")
	pass := os.Getenv("MYSQL_ROOT_PASSWORD")
	database := os.Getenv("MYSQL_DB_NAME")
	host := os.Getenv("MYSQL_HOST")

	// if user == "" || pass == "" || database == "" || host == "" {
	// 	log.Fatalf("環境変数が正しく設定されていません。user: %s, pass: %s, database: %s, host: %s", user, pass, database, host)
	// }

	return &DBConfig{
		User:     user,
		RootPass: pass,
		DBName:   database,
		Host:     host,
		Port:     "3306",
	}
}

package db

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"
)

type PostgresConfig struct {
	// minimum config untuk postgres
	Port              uint
	Host              string
	Username          string
	Password          string
	DBName            string
	MaxIdleConnection int
	MaxOpenConnection int
	MaxIdleTime       int
}

func NewDBPostges() *sql.DB {
	pt, _ := strconv.Atoi(os.Getenv("Port"))

	pgConf := PostgresConfig{
		Port:              uint(pt),
		Host:              os.Getenv("Host"),
		Username:          os.Getenv("Username"),
		Password:          os.Getenv("Password"),
		DBName:            os.Getenv("DBName"),
		MaxOpenConnection: 7,
		MaxIdleConnection: 5,
		MaxIdleTime:       int(30 * time.Minute),
	}

	connString := fmt.Sprintf(`
		host=%v
		port=%v
		user=%v
		password=%v
		dbname=%v
		sslmode=disable
	`,
		pgConf.Host,
		pgConf.Port,
		pgConf.Username,
		pgConf.Password,
		pgConf.DBName,
	)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}

	// test connection
	if err := db.Ping(); err != nil {
		panic(err)
	}

	// set extended config
	db.SetMaxIdleConns(pgConf.MaxIdleConnection)
	db.SetMaxOpenConns(pgConf.MaxOpenConnection)
	db.SetConnMaxIdleTime(time.Duration(pgConf.MaxIdleTime))

	return db

}

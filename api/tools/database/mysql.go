package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var (
	db    *sql.DB
	dbErr error
)

type dbConf struct {
	host     string
	port     string
	user     string
	password string
	database string
}

func Init() error {
	c := getDatabaseConf()
	// root:test123@tcp(192.168.56.23:3306)/potato?charset=utf8
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", c.user, c.password, c.host, c.port, c.database)
	db, dbErr = sql.Open("mysql", dsn)
	if dbErr != nil {
		return fmt.Errorf("sql open dbErr: %v\n", dbErr)
	}
	if err := db.Ping(); err != nil {
		return fmt.Errorf("db connection dbErr: %v\n", err)
	}
	return nil
}

func getDatabaseConf() *dbConf {
	return &dbConf{
		host:     viper.GetString("mysql.host"),
		port:     viper.GetString("mysql.port"),
		user:     viper.GetString("mysql.user"),
		password: viper.GetString("mysql.password"),
		database: viper.GetString("mysql.database"),
	}
}

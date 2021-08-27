package gorm

import (
	"fmt"

	"bibit-test/config"

	"github.com/jinzhu/gorm"

	// Register Gorm Mysql Driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// Register Go Sql Driver
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConfig  = config.Config.DB
	mysqlConn *gorm.DB
	err       error
)

// initialize database
func init() {
	if dbConfig.Driver == "mysql" {
		setupMysqlConn()
	}
}

// setupMysqlConn: setup mysql database connection using the configuration from config.yml
func setupMysqlConn() {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)
	mysqlConn, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	err = mysqlConn.DB().Ping()
	if err != nil {
		panic(err)
	}
}

// MysqlConn: return mysql connection from gorm ORM
func MysqlConn() *gorm.DB {
	return mysqlConn
}

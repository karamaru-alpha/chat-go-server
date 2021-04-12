package mysql

import (
	"fmt"
	"os"
	"time"

	// blank import for MySQL driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// ConnectGorm GormでDBコネクションを確保する
func ConnectGorm() *gorm.DB {
	DBMS := "mysql"
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_ROOT_PASSWORD")
	CONTAINER_NAME := os.Getenv("MYSQL_CONTAINER_NAME")
	PORT := os.Getenv("MYSQL_PORT")
	DBNAME := os.Getenv("MYSQL_DATABASE")

	CONNECT := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		USER,
		PASS,
		CONTAINER_NAME,
		PORT,
		DBNAME,
	)

	count := 0
	connection, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		for {
			if err == nil {
				fmt.Println("")
				break
			}
			fmt.Print(".")
			time.Sleep(time.Second)
			count++
			if count > 60 {
				fmt.Println("")
				fmt.Println("DB接続失敗")
				panic(err)
			}
			connection, err = gorm.Open(DBMS, CONNECT)
		}
	}

	connection.LogMode(true)
	connection.Set("gorm:table_options", "ENGINE=InnoDB")

	return connection
}

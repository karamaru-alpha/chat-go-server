package mysql

import (
	"fmt"
	"log"
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
	PROTOCOL := "tcp(chat-go-server-db:3306)"
	DBNAME := os.Getenv("MYSQL_DATABASE")

	CONNECT := fmt.Sprintf(
		"%s:%s@%s/%s?charset=utf8mb4&parseTime=true&loc=Local",
		USER,
		PASS,
		PROTOCOL,
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
			if count > 180 {
				log.Fatal("接続失敗", err)
			}
			connection, err = gorm.Open(DBMS, CONNECT)
		}
	}

	connection.LogMode(true)
	connection.Set("gorm:table_options", "ENGINE=InnoDB")

	fmt.Println("DB接続成功")
	return connection
}

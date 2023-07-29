package models

// document: https://gorm.io/zh_CN/docs/connecting_to_the_database.html
import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	dsn := "root:1003835955@qq.com@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		QueryFields: true, //打印sql
		// SkipDefaultTransaction: true, //禁用事务
	})
	fmt.Println("ddddddddd1", DB)
	if err != nil {
		fmt.Printf("数据库连接失败======error: %v\n", err)
		log.Fatal(err)
	}
}

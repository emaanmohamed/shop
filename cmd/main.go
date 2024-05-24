package main

import (
	"github.com/emaanmohamed/shop/cmd/api"
	"github.com/emaanmohamed/shop/db"
	"github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db.NewMySQLStorage(mysql.Config{
		User:                 "root",
		Passwd:               "",
		Addr:                 "localhost:3306",
		Net:                  "tcp",
		DBName:               "shop",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	server := api.NewAPIServer(":3434", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

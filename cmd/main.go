package main

import (
	"database/sql"
	"github.com/emaanmohamed/shop/cmd/api"
	"github.com/emaanmohamed/shop/config"
	"github.com/emaanmohamed/shop/db"
	"github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
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

func initStorage(db *sql.DB) {
	
}

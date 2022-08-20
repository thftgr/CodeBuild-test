package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/pterm/pterm"
	"log"
	"os"
)

var (
	Maria         *sql.DB
	mysqlUsername = os.Getenv("MYSQL_USER")
	mysqlPassword = os.Getenv("MYSQL_PASSWORD")
	mysqlHost     = os.Getenv("MYSQL_HOST")
	mysqlPort     = os.Getenv("MYSQL_PORT")
	mysqlDatabase = os.Getenv("MYSQL_SCHEMA")
	serverPort    = os.Getenv("SERVER_PORT")
)

func init() {
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysqlUsername, mysqlPassword, mysqlHost, mysqlPort, mysqlDatabase)
	log.Println(dbUrl)
	db, err := sql.Open("mysql", dbUrl)
	if Maria = db; db != nil {
		Maria.SetMaxOpenConns(100)

		pterm.Success.Println("RDBMS connected")

		if _, err = Maria.Exec("SELECT 1"); err != nil {
			pterm.Error.Println("'SELECT 1' failed.")
			panic(err)
		}
	} else {
		pterm.Error.Println("RDBMS Connect Fail", err)
		panic(err)
	}
}
func main() {
	e := echo.New()
	e.GET("/healthy", func(c echo.Context) error {
		log.Println("/healthy")
		return c.String(200, "ok")
	})
	e.Logger.Fatal(e.Start(":" + serverPort))
}

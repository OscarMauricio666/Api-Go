package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

const username string = "mauricio"
const password string = "password"
const host string = "localhost"
const port int = 3306
const database string = "go_database"

func CreateConection() {
	if connection, err := sql.Open("mysql", generateURL()); err != nil {
		panic(err)
	} else {
		db = connection
		fmt.Println("Conexión Establecida.")
	}
}

func CreateTables() {
	createTable("tickets", ticketsSchema)
	//  createTable("test1", testSchema) se debe crear su respesctivo esquema
}
func createTable(tableName, schema string) {
	if !existsTable(tableName) {
		Exec(schema)
	} else {
		truncateTable(tableName)
	}
}
func truncateTable(tableName string) {
	sql := fmt.Sprintf("TRUNCATE %s", tableName)
	Exec(sql)
}
func existsTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s' ", tableName)
	rows, _ := Query(sql)
	return rows.Next()
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		log.Println(err)
	}
	return result, err
}
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		log.Println(err)
	}
	return rows, err

}

func Ping() {
	if err := db.Ping(); err != nil {
		fmt.Println(err)
	}
}
func CloseConection() {
	db.Close()
}
func generateURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", username, password, host, port, database)
}

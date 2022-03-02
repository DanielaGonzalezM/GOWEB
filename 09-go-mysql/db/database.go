package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//username:passwoed@tcp(localhost:3306)/database
const url = "root:admin@tcp(localhost:3306)/goweb_db"

//Guarda la conneccion
var db *sql.DB

//Realiza la coneccion
func Connect() {
	conection, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Coneccion exitosa")
		db = conection
	}
}

//Cerrar coneccion
func Close() {
	db.Close()
}

//Verificar coneccion
func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

//Verifica si una tabla existe o no
func ExistTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	//rows, err := db.Query(sql)  //sin polimorfismo
	rows, err := Query(sql)
	if err != nil {
		fmt.Println("Error", err)
	}
	return rows.Next()

}

//Crea una tabla
func CreateTable(schema string, name string) {
	if !ExistTable(name) {
		_, err := db.Exec(schema)
		if err != nil {
			fmt.Println(err)
		}
	}

}

//Reiniciar el registro de una tabla
func TruncateTable(tableName string) {
	sql := fmt.Sprintf("TRUNCATE %s", tableName)
	Exec(sql)
}

//Polimorfismo de Exec
func Exec(query string, args ...interface{}) (sql.Result, error) {

	result, err := db.Exec(query, args...)
	if err != nil {
		fmt.Println(err)
	}
	return result, err
}

//Plomorfismo de query
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Println(err)
	}
	return rows, err
}

package database

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Connect()(*sql.DB, error){
	db, err := sql.Open("mysql","gosec:gosec321@tcp(server:3306)/gosec")
	if err != nil{
		return nil,err
	}
	return db,nil
}
package database

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Connect()(*sql.DB, error){
	db, err := sql.Open("mysql","root:admin321@tcp(localhost:3306)/seccoding")
	if err != nil{
		return nil,err
	}
	return db,nil
}
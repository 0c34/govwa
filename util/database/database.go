package database

import(
	"database/sql"
	"fmt"
	"govwa/util/config"
	_ "github.com/go-sql-driver/mysql"
)

func Connect()(*sql.DB, error){
	config := config.LoadConfig()
	cred := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.User,config.Password,config.Sqlhost,config.Sqlport,config.Dbname)
	db, err := sql.Open("mysql",cred)
	if err != nil{
		return nil,err
	}
	return db,nil
}
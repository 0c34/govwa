package database

import(
	"fmt"
	"log"
	"database/sql"
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

var DB *sql.DB
func CheckDatabase()bool{
	/* this function use to check if no database selected and will redirect to setup page */
	DB, err := Connect()
	if err != nil{
		log.Printf("Connection Error %s ",err.Error())
	}
	
	sql := "USE govwa"
	result, err := DB.Exec(sql)
	if err != nil{
		log.Println(err.Error())
	}
	if result == nil{
		return false
	}
	log.Println(result)
	return true
}
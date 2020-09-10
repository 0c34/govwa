package setup

import (
	"database/sql"

	"github.com/govwa/util/database"
)

const (
	DropUsersTable = `DROP TABLE IF EXISTS Users`

	CreateUsersTable = `CREATE TABLE Users (
		id int(10) NOT NULL AUTO_INCREMENT,
		uname varchar(100) NOT NULL,
		pass varchar(100) NOT NULL,
		PRIMARY KEY (id)
	  ) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1`

	InsertUsers = `INSERT INTO Users VALUES (1,'admin','9f3b6fa4703a5ba96fda0dee48ec76fc'),(2,'user1','ff1d5c0015a535b01a5d03a373bf06f6')`

	DropProfilesTable = `DROP TABLE IF EXISTS Profile`

	CreateProfilesTable = `CREATE TABLE Profile (
		profile_id int(10) NOT NULL AUTO_INCREMENT,
		user_id int(10) NOT NULL,
		full_name varchar(100) NOT NULL,
		city varchar(100) NOT NULL,
		phone_number varchar(15) NOT NULL,
		PRIMARY KEY (profile_id)
	  ) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1`

	InsertProfile = `INSERT INTO Profile VALUES (1,1,'Andro','Jakarta','08882112345'),(2,2,'Rocky','Bandung','08882112345')`
)

var DB *sql.DB
var err error

/*func init() {
	DB, err = database.Connect()
	if err != nil {
		log.Println(err.Error())
	}
}*/

func createUsersTable() error {

	DB, err = database.Connect()

	_, err = DB.Exec(DropUsersTable)
	if err != nil {
		return err
	}
	_, err = DB.Exec(CreateUsersTable)
	if err != nil {
		return err
	}
	_, err = DB.Exec(InsertUsers)
	if err != nil {
		return err
	}
	return nil
}

func createProfileTable() error {
	_, err = DB.Exec(DropProfilesTable)
	if err != nil {
		return err
	}
	_, err = DB.Exec(CreateProfilesTable)
	if err != nil {
		return err
	}
	_, err = DB.Exec(InsertProfile)
	if err != nil {
		return err
	}
	return nil
}

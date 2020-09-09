package xss

import(
	"log"
	"database/sql"

	"github.com/govwa/util/database"
)

var DB *sql.DB
var err error

func init(){
	DB, err = database.Connect()
	if err != nil{
		log.Println(err.Error())
	}
}

type Search struct{
	
}

func GetExp(term string)string{
	
	var vuln = make(map[string]string)
	
	vuln["xss"] = `Cross-Site Scripting (XSS) attacks are a type of injection, in which malicious scripts are injected into otherwise benign and trusted web sites`
	vuln["sqli"] = `A SQL injection attack consists of insertion or "injection" of a SQL query via the input data from the client to the application. `
	vuln["idor"] = `Insecure Direct Object References occur when an application provides direct access to objects based on user-supplied input`
	vuln["xxe"] = `An XML External Entity attack is a type of attack against an application that parses XML input.`
	
	text, ok := vuln[term]

	if ok{
		return text
	}else{
		return ""
	}
}
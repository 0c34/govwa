package util

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type config struct{
	User string
	Password string
	Dbname string
	Sqlhost string
	Sqlport string
	Webserver string
	Webport string
	Sessionkey string
}

var Cfg *config
var Fullurl string

func LoadConfig()*config{
	raw, err := ioutil.ReadFile("config/config.json")
	if err != nil{
		fmt.Println(err.Error())
	}
	configuration := config{}
	err = json.Unmarshal(raw, &configuration)
	return &configuration
}

func init(){
	Cfg = LoadConfig()
	url := fmt.Sprintf("%s:%s/",Cfg.Webserver,Cfg.Webport)
	Fullurl = url
}

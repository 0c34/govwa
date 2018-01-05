package config

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type Config struct{
	User string
	Password string
	Dbname string
	Sqlhost string
	Sqlport string
	Webserver string
	Webport string
	Sessionkey string
}

var Cfg *Config
var Fullurl string

func LoadConfig()*Config{
	raw, err := ioutil.ReadFile("config/config.json")
	if err != nil{
		fmt.Println(err.Error())
	}
	configuration := Config{}
	err = json.Unmarshal(raw, &configuration)
	return &configuration
}

func init(){
	Cfg = LoadConfig()
	url := fmt.Sprintf("%s:%s/",Cfg.Webserver,Cfg.Webport)
	Fullurl = url
}

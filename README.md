#### GoVWA
---
GoVWA (Go Vulnerable Web Application) is a golang web application which contain vulnerability based on OWASP top 10. This application simulate how vulnerability occure and help developer or security engineer to learn security on golang application.
#### WARNING!
---
GoVWA is a vulnerable web application, **run the application on your local or testing environment only**

#### Installation
---
#### Install golang
Install golang on you host

#### Setup from source
```
git clone https://github.com/0c34/govwa.git

git pull (to update)

```
Install dependency packages
```
go mod download 
```

#### GoVWA config
---
Modified the config.json file for database configuration

```
{
    "user": "root",
    "password": "root",
    "dbname": "govwa",
    "sqlhost": "localhost",
    "sqlport": "3306",
    "webserver": "http://localhost",
    "webport": "8888",

    "sessionkey:": "G0Vw444"
}

```
Run GoVWA 
```
go run app.go 
```
```

     ÛÛÛÛÛÛÛÛÛ           ÛÛÛÛÛ   ÛÛÛÛÛ ÛÛÛÛÛ   ÛÛÛ   ÛÛÛÛÛ   ÛÛÛÛÛÛÛÛÛ  
    ÛÛÛ°°°°°ÛÛÛ         °°ÛÛÛ   °°ÛÛÛ °°ÛÛÛ   °ÛÛÛ  °°ÛÛÛ   ÛÛÛ°°°°°ÛÛÛ 
   ÛÛÛ     °°°   ÛÛÛÛÛÛ  °ÛÛÛ    °ÛÛÛ  °ÛÛÛ   °ÛÛÛ   °ÛÛÛ  °ÛÛÛ    °ÛÛÛ 
  °ÛÛÛ          ÛÛÛ°°ÛÛÛ °ÛÛÛ    °ÛÛÛ  °ÛÛÛ   °ÛÛÛ   °ÛÛÛ  °ÛÛÛÛÛÛÛÛÛÛÛ 
  °ÛÛÛ    ÛÛÛÛÛ°ÛÛÛ °ÛÛÛ °°ÛÛÛ   ÛÛÛ   °°ÛÛÛ  ÛÛÛÛÛ  ÛÛÛ   °ÛÛÛ°°°°°ÛÛÛ 
  °°ÛÛÛ  °°ÛÛÛ °ÛÛÛ °ÛÛÛ  °°°ÛÛÛÛÛ°     °°°ÛÛÛÛÛ°ÛÛÛÛÛ°    °ÛÛÛ    °ÛÛÛ 
   °°ÛÛÛÛÛÛÛÛÛ °°ÛÛÛÛÛÛ     °°ÛÛÛ         °°ÛÛÛ °°ÛÛÛ      ÛÛÛÛÛ   ÛÛÛÛÛ
     °°°°°°°°°   °°°°°°       °°°           °°°   °°°      °°°°°   °°°°° 
Server running at port :888
Open this url http://localhost:8888/ on your browser to access GoVWA

```
Open the url to access GoVWA and follow the setup instruction to create database and tables

#### Setup from docker
```
git clone https://github.com/0c34/govwa.git

inside govwa directory:
docker-compose up --build

stop running process using
docker-compose down --remove-orphans --volumes

```

GoVWA users:

|uname|password|
|-----|--------|
|admin|govwaadmin|
|user1|govwauser1|

Explore the vulnerability.

#### Contributor
---
* Khaedir (golang programming)
* Xaquille (web design)

#### To Do

* add more vulnerabilities

Powered by [NemoSecurity](https://nemosecurity.com)





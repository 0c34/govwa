#### GoVWA
GoVWA (Go Vulnerable Web Application) is a vulnerable web application designed for pentester or programmers to learn the web application vulnerability that often occur in web applications. The vulnerabilities in GoVWA are OWASP Top 10 category. 

#### WARNING!
---
GoVWA is a vulnerable web application, **Run it only on local environment**

#### Installation
---
#### Installing golang
Golang versiong : >= 1.11 
Installing guide : https://www.digitalocean.com/community/tutorials/how-to-install-go-1-6-on-ubuntu-16-04

#### Setup
```
git clone https://github.com/0c34/govwa.git

git pull (to update)

```
#### Install dependency packages

```
go mod download 
```

#### GoVWA config
---
#### Modified the config.json file for database configuration

config.json file is located in config directory.

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

=======
Server running at port :8082
Open this URL http://192.168.56.101:8082/ on your browser to access GoVWA

```
Open the URL to access GoVWA and follow the setup instruction to create database and tables

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
* ~~XXE Vulnerability~~
* NoSQLInjection
* JSON Web API (unprotected API)
* Build Simple Android APP

Powered by [NemoSecurity](https://nemosecurity.com)

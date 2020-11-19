#### GoVWA
GoVWA (Go Vulnerable Web Application) is a web application developed to help the pentester and programmers to learn the vulnerabilities that often occur in web applications which are developed using golang. Vulnerabilities that exist in GoVWA are the most common vulnerabilities found in web applications today. So it will help programmers recognize vulnerabilities before they happen to their application. Govwa can also be an additional application of your pentest lab for learning and teaching.

#### WARNING!
---
GoVWA is a vulnerable web application, **run the application on your local or testing environment only**

#### Installation
---
#### Installing golang
If you didn't have golang installed on your system. First, install it using automation script from https://github.com/canha/golang-tools-install-script.

Follow the instruction which is provided by the author and install golang depending on your Operating System Architecture.

If successfully installed you would have a directory 'go' in your home directory. The go directory has three subdirectories (bin, pgk, src). Switch to src directory then clone govwa repository. 

#### Setup from source
```
git clone https://github.com/0c34/govwa.git

git pull (to update)

```
#### Install dependency packages
we have to install several golang packages that required by govwa

Execute those command in your terminal
```
go mod download 
```

#### GoVWA config
---
#### Modified the config.json file for database configuration

Open the file config.json which is located in the config directory. Change the configuration according to your needs.

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

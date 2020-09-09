#### GoVWA
---
GoVWA (Go Vulnerable Web Application) is a golang web application which contain several web application vulnerability based on OWASP top 10. This application simulate how vulnerability occure and help developer or security engineer to learn security on golang application.
#### WARNING!
---
Since GoVWA is a web application that contains a vulnerability, **never upload govwa to web hosting that can be accessed publicly, because it can cause your server to get hacked**. As a suggestion to use GoVWA locally.

#### How To Install GoVWA
---
#### Installing golang
If you didn't have golang installed on your system. first, install it using automation script from https://github.com/canha/golang-tools-install-script.

Follow the instruction which is provided by the author and install golang depending on your Operating System Architecture.

If successfully installed you would have directory 'go' in your home directory. the go directory has three subdirectory (bin, pgk, src). switch to src directory then clone govwa repository. 

```
git clone https://github.com/0c34/govwa.git

git pull (to update)

```
we have to install several golang package that required by govwa

Execute those command in your terminal
```
go get github.com/go-sql-driver/mysql
go get github.com/gorilla/sessions
go get github.com/julienschmidt/httprouter
```

#### GoVWA config
---
Open the file config.json which is located in config directory. Change the configuration according to your needs.

```
{
    "user": "root",
    "password": "root",
    "dbname": "govwa",
    "sqlhost": "localhost",
    "sqlport": "3306",
    "webserver": "http://192.168.56.101",
    "webport": "8082",

    "sessionkey:": "G0Vw444"
}

```
Run GoVWA 
```
govwa@ubuntu-server:~/go/src/govwa$ go run app.go 
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
Server running at port :8082
Open this url http://192.168.56.101:8082/ on your browser to access GoVWA

```
Open the url to access GoVWA and follow the setup instruction to create database and tables

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

* ~~XXE Vulnerability~~
* NoSQLInjection
* Json Web API (unprotected API)
* Build Simple Android APP

warm regards [NemoSecurity](https://nemosecurity.com)





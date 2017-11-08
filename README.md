#### GoVWA
---
GoVWA (Go Vulnerable Web Application) is a web application developed to help the pentester and programmers to learn the vulnerabilities that often occur in web applications which is developed using golang. Vulnerabilities that exist in GoVWA are the most common vulnerabilities found in web applications today. So it will help programmers recognize vulnerabilities before they happen to our app. Govwa can also be an additional application of your pentest lab for learning and teaching.

#### How To Install GoVWA
---
#### Installing golang
If you didn't have golang installed on your system. for the first install it using automation script from https://github.com/canha/golang-tools-install-script.

Follow the instruction that provided and install golang depending on your Operating System Architecture.

If successfully installed you will have a directory called 'go' in your home directory. the go directory have three subdirectory (bin, pgk, src). go to src directory and clone the govwa repository. 

```
git clone https://github.com/0c34/govwa.git

```
Before start to running govwa we need to install several golang package that required by govwa

Execute command below on you linux terminal
```
go get github.com/go-sql-driver/mysql
go get github.com/gorilla/sessions
go get github.com/julienschmidt/httprouter
```

#### GoVWA config
---
Go to govwa diretory and edit file inside directory config/config.json and change the configuration according to your needs.

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

#### Note
---
GoVWA still developing on XXE and another vulnerability. however it still has functionality that doesn't works well. if you have any suggestion feel free to email me at sulhaedir05[at]gmail[dot]com

warm regards [NemoSecurity](https://nemosecurity.com)

#### Contributor
---
* Khaedir (golang programming)
* Xaquille (web design)







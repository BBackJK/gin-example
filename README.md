# gin-example

This is **Rest API** todo example with Golang, gin framework

# Getting Start

## install 

- Go : [https://golang.org/dl/](https://golang.org/dl/)
- MySQL : [https://dev.mysql.com/downloads/](https://dev.mysql.com/downloads/)

## setting

first, git clonning
```
git clone https://github.com/BBackJK/gin-example.git
```

.env file is local environment settings.

```
PORT=8080   // start port number set.
DB_CONFIG=test:test@/test?charset=utf8&parseTime=True&loc=Local     // db connecting set. [user:password@/dbname?charset=utf8&parseTime=True&loc=Local]
CORS=http://localhost:8080      // CORS set.
```

## dependency

this example does not use dependency.

if you want use dependency, see here [https://github.com/golang/dep](https://github.com/golang/dep)

anyways, you install the package.

```
go get -u github.com/gin-gonic/gin      // -u option is lastest version. [gin framework]
go get -u github.com/joho/godotenv      // [.env file using package.]
go get -u github.com/jinzhu/gorm        // [database orm package.]
go get -u github.com/jinzhu/gorm/dialects/mysql   // [database orm mysql.]
```

and then, in **$GOPATH/src/gin-example**

```
go run main.go
```

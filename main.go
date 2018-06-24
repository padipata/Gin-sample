package main

import (
 "github.com/gin-gonic/gin"
 "net/http"
)

func main() {

    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?parseTime=true")
    if err != nil{
     log.Fatalln(err)
    }
    defer db.Close()
    
    db.SetMaxIdleConns(20)
    db.SetMaxOpenConns(20)
    
    if err := db.Ping(); err != nil{
     log.Fatalln(err)
    }
   
    router := gin.Default()
    router.GET("/", func(c *gin.Context) {
     c.String(http.StatusOK, "It works")
    })
   
    router.Run(":8000")
   }
package mysql

import (    
    "fmt"
    "strconv"
    "github.com/gin-gonic/gin"  
    mysqlPool "github.com/hopehook/golang-db/mysql"
)

type UserType struct {
    Id      int    `form:"id" json:"id" `
    Name    string `form:"name" json:"name" binding:"required"`
    Age     int    `form:"age" json:"age" binding:"required"`
}

func List(mysqlDB *mysqlPool.SQLConnPool) gin.H{
    body := make(gin.H) 
    rows, err := mysqlDB.Query("SELECT * From user")
    if err != nil {
        fmt.Printf("%s\r\n","mysql query error")
    }
    //fmt.Printf("%v\r\n",rows)
    var dbdata []gin.H
    if rows != nil {
        for _, row := range rows {
            dbdata = append(dbdata, gin.H(row))
        }
    }
    body["status"] = 200
    body["data"] = dbdata
    return body
}

func AddOne(mysqlDB *mysqlPool.SQLConnPool, c *gin.Context) gin.H{
    body := make(gin.H) 
    // 保存
    var json UserType
    if err := c.ShouldBindJSON(&json); err == nil {
        lastId, err := mysqlDB.Update("INSERT INTO user (`name`, `age`) VALUES( ?, ? )", json.Name, json.Age) // ? = placeholder
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
        body["updateCount"] = lastId
        body["status"] = "success"
    } else {
        body["status"] = err.Error()
    }
    return  body
}

func UpdateById(mysqlDB *mysqlPool.SQLConnPool, c *gin.Context) gin.H{
    userId, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        panic("userId can not empty")  
    }
    body := make(gin.H) 
    // 保存
    var json UserType
    if err := c.ShouldBindJSON(&json); err == nil {
        // 进行数据库操作
        affect, err := mysqlDB.Update("update user set `name` = ? , `age` = ? where `id` = ? ", json.Name, json.Age, userId) // ? = placeholder
        if err != nil {
            panic(err.Error()) 
        }
        body["updateCount"] = affect
        body["status"] = "success"
    } else {
        body["status"] = err.Error()
    }
    return  body
}

func DeleteById(mysqlDB *mysqlPool.SQLConnPool, c *gin.Context) gin.H{
    userId, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        panic("userId can not empty")  
    }
    body := make(gin.H) 
    affect, err := mysqlDB.Update("delete from user where `id` = ?", userId) // ? = placeholder
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    body["deleteCount"] = affect
    return  body
}
# golang-db
mysql  redis mongodb等数据库连接池的库包


准备工作
-----------

> 从零开始，搭建golang环境，到做一个mysql的增删改查的api

1.安装golang，参看：[安装golang](http://www.fancyecommerce.com/2017/12/28/centos6-%e5%ae%89%e8%a3%85-golang-1-9/)



2.安装本库包

```
go get github.com/fecshopsoft/golang-db
```

3.使用mysql  

3.1 mysql

通过web访问的方式测试，因此，安装了gin框架

```
go get github.com/gin-gonic/gin
```

main.go

```
package main

import(
    "github.com/gin-gonic/gin"
    "net/http" 
    _ "github.com/go-sql-driver/mysql" 
    mysqlPool "github.com/fecshopsoft/golang-db/mysql"
    testMysql "github.com/fecshopsoft/golang-db/test/mysql"
)

func mysqlDBPool() *mysqlPool.SQLConnPool{
    host := `127.0.0.1:3306`
    database := `go_test`
    user := `root`
    password := `xxx`
    charset := `utf8`
    // 用于设置最大打开的连接数
    maxOpenConns := 200
    // 用于设置闲置的连接数
    maxIdleConns := 100
    mysqlDB := mysqlPool.InitMySQLPool(host, database, user, password, charset, maxOpenConns, maxIdleConns)
    return mysqlDB
}

func main() { 
    mysqlDB := mysqlDBPool();
	r := gin.Default()
    v2 := r.Group("/v2")
    {
        // 查询部分
        v2.GET("/users", func(c *gin.Context) {
            data := testMysql.List(mysqlDB);
            c.JSON(http.StatusOK, data)
        })
        v2.POST("/users", func(c *gin.Context) {
            data := testMysql.AddOne(mysqlDB, c);
            c.JSON(http.StatusOK, data)
        })
        v2.PATCH("/users/:id", func(c *gin.Context) {
            data := testMysql.UpdateById(mysqlDB, c);
            c.JSON(http.StatusOK, data)
        })
        v2.DELETE("/users/:id", func(c *gin.Context) {
            data := testMysql.DeleteById(mysqlDB, c);
            c.JSON(http.StatusOK, data)
        })
    }
    r.Run("120.24.37.249:3000") // 这里改成您的ip和端口
}
```

3.2 数据库go_test插入数据


```
CREATE TABLE IF NOT EXISTS `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) DEFAULT '',
  `age` int(11) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=11 ;

--
-- 转存表中的数据 `user`
--

INSERT INTO `user` (`id`, `name`, `age`) VALUES
(1, '111', 111),
(2, 'terry', 44),
(3, 'terry', 55),
(4, 'terry', 44),
(5, 'terry', 44),
(6, 'terry', 44),
(7, 'terry', 44),
(8, 'terry', 44),
(10, 'terry', 66);
```










package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	router:=gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")
	router.GET("/login.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
		})
	})
	router.GET("/sign_up.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "sign_up.html", gin.H{
		})
	})
	router.GET("/DASHBOARD.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "DASHBOARD.html", gin.H{
		})
	})
	router.GET("/network.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "network.html", gin.H{
		})
	})
	router.GET("/blocks.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "blocks.html", gin.H{
		})
	})
	router.GET("/transactions.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "transactions.html", gin.H{
		})
	})
	router.GET("/user/name/:username/:password", getuser)
	router.GET("/user/register/query_name/:username", registeruser)
	router.GET("/user/register/do_resiter/:username/:password", registeruserpassword)
	router.Run(":8080")
} 
func getuser(c *gin.Context) {
	name := c.Param("username")
	password:= c.Param("password") 
	fmt.Println("name:", name, "pass:", password)
	result := Fabric(name, password)
	c.String(http.StatusOK, result)
}
func registeruser(c *gin.Context) {
	rename := c.Param("username")
	password:= c.Param("password") 
	fmt.Println("name:", rename, "pass:", password)
	result := RegisterQuery(rename, password)
	fmt.Println(result)
	c.String(http.StatusOK, result)
}

func registeruserpassword(c *gin.Context) {
	rename := c.Param("username")
	password:= c.Param("password") 
	fmt.Println("name:", rename, "pass:", password)
	result := RegisterPrepare(rename, password)
	fmt.Println(result)
	c.String(http.StatusOK, result)
}

func Fabric(username string, password string) string {	
	db,err := sql.Open("sqlite3","./UserDB.db")
	checkErr(err)
	var userName string
	userName = ""
	rows,err :=db.Query("SELECT USERNAME FROM USER WHERE USERNAME=? AND PASSWORD=?",username,password)
	for rows.Next() {
		err = rows.Scan(&userName)
		checkErr(err)
		fmt.Println(userName)
	}
	rows.Close()
	db.Close()
	return userName
}

func RegisterQuery(username string, password string) string {
	db,err:=sql.Open("sqlite3","./UserDB.db")
	checkErr(err)
	var userName string	
	rows,err :=db.Query("SELECT USERNAME FROM USER WHERE USERNAME=?",username)	
	for rows.Next() {
		err = rows.Scan(&userName)
		checkErr(err)
	}
	return userName;
}
	
func RegisterPrepare(username string, password string) string {
	db,err:=sql.Open("sqlite3","./UserDB.db")
	checkErr(err)	
	stmt, err := db.Prepare("insert into USER(USERNAME,PASSWORD) values(?,?)")
	checkErr(err)
	_, err = stmt.Exec(username, password)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(username)
	db.Close()
	return "11"
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}



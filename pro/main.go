package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	Username string
	Password string
}

func Getnumber() (number int) {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	return secretNumber
}
func main() {
	n := Getnumber()
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&User{})
	r := gin.Default()
	r.LoadHTMLFiles("./login.html", "./index.html", "./guess.html", "wrong.html", "register_2.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	var username string
	var password string
	var (
		rusername string
		rpassword string
		info      string
	)
	var u1 User
	r.POST("/login", func(c *gin.Context) {
		username = c.PostForm("username")
		password = c.PostForm("password")
		db.Where(map[string]interface{}{"Username": username, "Password": password}).Find(&u1)
		if u1.Username == username && u1.Password == password {
			info = "hello"
		} else {
			info = "用户名或者密码错误"
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Info": info,
		})
	})
	r.GET("/register", func(c *gin.Context) {
		c.HTML(200, "register_2.html", nil)
	})
	r.POST("/register", func(c *gin.Context) {
		rusername = c.PostForm("rusername")
		rpassword = c.PostForm("rpassword")
		u := User{rusername, rpassword}
		db.Create(&u)
		c.JSON(200, "注册成功")
	})
	r.GET("/guess", func(c *gin.Context) {
		c.HTML(200, "guess.html", gin.H{
			"Number": n,
		})
	})
	r.GET("/results", func(c *gin.Context) {
		t := c.PostForm("Dnumber")
		c.HTML(200, "wrong.html", gin.H{
			"Result": t,
			"Number": n,
		})
	})
	r.Run()
}

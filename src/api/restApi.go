package api

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"model"
	"strconv"
)

var Debug bool

func InitDb() *gorm.DB {
	var err error
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/gorm?charset=utf8")
	if err != nil {
		panic("failed to connect database")
	}
	//defer db.Close()
	db.LogMode(true)
	
	// Creating the table
	if !db.HasTable(&model.User{}) {
		db.CreateTable(&model.User{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.User{})
	}
	
	return db
}

func Login(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	
	inputName := c.Request.FormValue("inputName")
	inputPassword := c.Request.FormValue("inputPassword")
	success, user, err := model.QueryOne(inputName, inputPassword)
	if err != nil {
		log.Println(err)
		c.JSON(505, gin.H{"error": "未知错误,请耐心重试!"})
		return
	}
	
	if success && user.UserName != "" {
		
		
		cookie := &http.Cookie{
			Name:     "token",
			Value:    inputName+"_"+strconv.Itoa(int(time.Now().Unix())),
			Path:     "/",
			HttpOnly: true,
		}
		
		http.SetCookie(c.Writer, cookie)
		// Display JSON result
		log.Println("登陆成功:",user.UserName)
		c.JSON(200, gin.H{"msg":"登陆成功!"})
		return
	} else {
		// Display JSON error
		//c.JSON(404, gin.H{"error": "User not found"})
		c.JSON(404, gin.H{"error": "用户不存在或账号、密码不对"})
		return
	}
	
	// curl -i http://localhost:8080/api/v1/login
}

func Register(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	
	inputName := c.Request.FormValue("inputName")
	inputPassword := c.Request.FormValue("inputPassword")
	inputNickname := c.Request.FormValue("inputNickname")
	log.Println(inputName, inputPassword)
	
	var user = &model.User{
		UserName: inputName,
		UserPwd:  inputPassword,
		NickName: inputNickname,
	}
	
	if model.QueryUsernameExist(inputName) {
		errMsg := gin.H{"error": "用户名已存在!"}
		log.Println(errMsg)
		c.JSON(401, errMsg)
		return
	}
	err := model.InsertUserInfo(user)
	if err != nil {
		log.Println(err)
		c.JSON(505, gin.H{"error": "系统错误,请稍后耐心重试!"})
		return
	}
	
	if user.UserName != "" || user.UserPwd != "" {
		cookie := &http.Cookie{
			Name:     "token",
			Value:    "12345",
			Path:     "/",
			HttpOnly: true,
		}
		http.SetCookie(c.Writer, cookie)
		//插入新用户信息
		db.Save(&user)
		// Display JSON result
		log.Println("用户注册成功：", user.UserName)
		c.JSON(200, gin.H{"msg":"注册成功!"})
	} else {
		// Display JSON error
		//c.JSON(404, gin.H{"error": "User not found"})
		c.JSON(404, gin.H{"error": "用户注册信息有误！"})
		return
	}
	
	// curl -i http://localhost:8080/api/v1/
}

func ShowAll(c *gin.Context) {
	db := InitDb()
	// Close connection database
	defer db.Close()
	
	var users []model.User
	// SELECT * FROM users
	db.Find(&users)
	
	// Display JSON result
	c.JSON(200, users)
	
}

func NowTime() string {
	return time.Now().Format("2003-01-02 15:04:05")
}

func IsExist(username string) (bool) {
	db := InitDb()
	defer db.Close()
	
	var user model.User
	result := db.Where("user_name = ?", username).First(&user).RecordNotFound()
	return !result
}

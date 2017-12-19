package api

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
	"net/http"
	"model"
	"github.com/gin-gonic/gin"
)

var Debug bool

func InitDb() *gorm.DB {
	var err error
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/test?charset=utf8")
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

func LoginApi(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	
	inputName := c.Request.FormValue("inputName")
	inputPassword := c.Request.FormValue("inputPassword")
	
	//以下为测试api专用
	if Debug {
		var user1 model.User
		err := c.BindJSON(&user1)
		if err != nil {
			c.Next()
		}else{
			if user1.UserName != "" {
				inputName = user1.UserName
				inputPassword = user1.UserPwd
			}
		}
	}
	log.Println(inputName, inputPassword)
	
	success, user, err := model.QueryOne(inputName, inputPassword)
	if err != nil {
		log.Println(err)
		c.JSON(505, gin.H{"error": "unknown error,please retry later patience !"})
		return
	}
	
	//登陆验证
	if (success && user.UserName != "") || (inputName == "admin" && inputPassword == "123") {
		cookie := &http.Cookie{
			Name:     "token",
			Value:    "123456",
			Path:     "/",
			HttpOnly: true,
		}
		
		http.SetCookie(c.Writer, cookie)
		// Display JSON result
		log.Println("登陆成功:", user.UserName)
		c.JSON(200, gin.H{"success": "login success!"})
		return
	} else {
		// Display JSON error
		//c.JSON(404, gin.H{"error": "User not found"})
		c.JSON(404, gin.H{"error": "username or password wrong,please try again!"})
		return
	}
	
	//curl -i -X POST -H "Content-Type: application/json" -d "{ \"username\": \"xiye\", \"userpwd\": \"123\"}" http://localhost:8080/api/v1/login
}

func RegisterApi(c *gin.Context) {
	if c.Request.Method == "GET" {
		errMsg := gin.H{"error": "请使用正确的post请求!"}
		log.Println(errMsg)
		c.JSON(401, errMsg)
		return
	}
	
	db := InitDb()
	defer db.Close()
	
	inputName := c.Request.FormValue("inputName")
	inputPassword := c.Request.FormValue("inputPassword")
	inputNickname := c.Request.FormValue("inputNickname")
	log.Println(inputName, inputPassword)
	//以下为测试api专用
	if Debug {
		var user1 model.User
		err := c.Bind(&user1)
		if err != nil {
			c.Next()
		} else {
			if user1.UserName != "" {
				inputName = user1.UserName
				inputPassword = user1.UserPwd
				inputNickname = user1.NickName
			}
		}
	}
	
	var user = &model.User{
		UserName: inputName,
		UserPwd:  inputPassword,
		NickName: inputNickname,
	}
	
	if model.QueryUsernameExist(inputName) {
		errMsg := gin.H{"error": "username is existing!"}
		log.Println(errMsg)
		c.JSON(401, errMsg)
		return
	}
	//err := model.InsertUserInfo(user)
	//if err != nil {
	//	log.Println(err)
	//	c.JSON(505, gin.H{"error": "unknown error,please retry later patience !"})
	//	return
	//}
	
	if user.UserName != "" || user.UserPwd != "" {
		cookie := &http.Cookie{
			Name:     "token",
			Value:    "12345",
			Path:     "/",
			HttpOnly: true,
		}
		http.SetCookie(c.Writer, cookie)
		user.UptDate = NowTime()
		//插入新用户信息
		db.Save(&user)
		// Display JSON result
		log.Println("用户注册成功：", user.UserName)
		c.JSON(200, gin.H{"success": "register success!"})
	} else {
		// Display JSON error
		//c.JSON(404, gin.H{"error": "User not found"})
		c.JSON(404, gin.H{"error": "用户注册信息有误！"})
		return
	}
	
	//curl -i -X POST -H "Content-Type: application/json" -d "{ \"username\": \"wade\", \"userpwd\": \"123\"}" http://localhost:8080/api/v1/reg
	
}

func ShowAllApi(c *gin.Context) {
	//file:///E:/workspace/ginTest/src/html/login.html
	//http://localhost:8080/api/v1/show
	//log.Println("ShowAllApi")
	if cookie, err := c.Request.Cookie("token"); err == nil {
		//判断token值是否正确
		token := cookie.Value
		log.Println(token)
		if token == "123456" {
			db := InitDb()
			// Close connection database
			defer db.Close()
			
			var users []model.User
			// SELECT * FROM users
			db.Find(&users)
			
			// Display JSON result
			c.JSON(200, users)
			return
		}
		c.JSON(403, gin.H{"error": "Authentication Failed"})
		return
	}
	
	c.JSON(501, gin.H{"error": "登录后才能查看用户列表!"})
	return
}

func Middleware(c *gin.Context) {
	if cookie, err := c.Request.Cookie("token"); err == nil {
		//判断token值是否正确
		token := cookie.Value
		if token == "123456" {
			c.Next()
			return
		}
		//c.JSON(403, gin.H{"error": "Authentication Failed"})
		//c.Abort()
		c.AbortWithStatusJSON(403, gin.H{"error": "Authentication Failed"})
		return
	}
	
	//c.JSON(501, gin.H{"error": "please login,After logging in, you can see the user list!"})
	//c.JSON(501, gin.H{"error": "登录后才能查看用户列表!"})
	//c.Abort()
	c.AbortWithStatusJSON(501, gin.H{"error": "登录后才能查看用户列表!"})
	return
	//log.Println("Middleware")
	//c.Next()
}

func NowTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func IsExist(username string) (bool) {
	db := InitDb()
	defer db.Close()
	
	var user model.User
	result := db.Where("user_name = ?", username).First(&user).RecordNotFound()
	return !result
}

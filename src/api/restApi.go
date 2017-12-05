package api

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strconv"
	"log"
	"time"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"model"
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
	
	//ToDo 调用登陆请求  返回结果：1.成功	2.失败
	if success && user.UserName != "" {
		cookie := &http.Cookie{
			Name:     "token",
			Value:    "12345",
			Path:     "/",
			HttpOnly: true,
		}
		
		http.SetCookie(c.Writer, cookie)
		// Display JSON result
		log.Println("用户登陆成功：", user.UserName)
		c.JSON(200, user)
		
	} else {
		// Display JSON error
		//c.JSON(404, gin.H{"error": "User not found"})
		c.JSON(404, gin.H{"error": "用户不存在或账号、密码不对"})
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
		var userReg model.User
		db.First(&user, "user_name = ?", inputName)
		c.JSON(200, gin.H{"用户注册成功success": userReg})
	} else {
		// Display JSON error
		//c.JSON(404, gin.H{"error": "User not found"})
		c.JSON(404, gin.H{"error": "用户注册信息有误！"})
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

//新增用户
func PostUser(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	
	var userRegister model.User
	c.Bind(&userRegister)
	//ToDo 判断用户名是否存在，如存则无法注册
	if IsExist(userRegister.UserName) {
		log.Println("用户名已存在:", userRegister.UserName)
		c.JSON(422, gin.H{"error": "username is exist,please change your register name"})
		return
	}
	
	if userRegister.UserName != "" && userRegister.UserPwd != "" {
		// INSERT INTO "user" (name) VALUES (user.Name);
		userRegister.UptDate = NowTime()
		db.Create(&userRegister)
		// Display error
		c.JSON(201, gin.H{"success": userRegister})
	} else {
		// Display error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}
	
	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"username\": \"wuqilong\", \"UserPwd\": \"123456\", \"nickname\": \"nicheng\", \"uptdate\": \"2017-12-01 15:34:29\" }" http://localhost:8080/api/v1/user
	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"username\": \"james\", \"UserPwd\": \"heat\", \"nickname\": \"king\"}" http://localhost:8080/api/v1/user
}

//获取所有用户
func GetUsers(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()
	
	var users []model.User
	// SELECT * FROM users
	db.Find(&users)
	
	// Display JSON result
	c.JSON(200, users)
	
	// curl -i http://localhost:8080/api/v1/user
}

//获取指定User
func GetUser(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	
	id := c.Params.ByName("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		go
			log.Println(err)
	}
	var user model.User
	// SELECT * FROM user WHERE id = 1;
	db.First(&user, userId)
	
	if user.Id != 0 {
		// Display JSON result
		c.JSON(200, user)
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "User not found"})
	}
	
	// curl -i http://localhost:8080/api/v1/user/1
}

//更新用户
func UpdateUser(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	
	// Get id user
	id := c.Params.ByName("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
	}
	var user model.User
	// SELECT * FROM user WHERE id = 1;
	db.First(&user, userId)
	
	if user.UserName != "" && user.UserPwd != "" {
		
		if user.Id != 0 {
			var newUser model.User
			c.Bind(&newUser)
			
			result := model.User{
				Id:       user.Id,
				UserName: newUser.UserName,
				UserPwd:  newUser.UserPwd,
				NickName: newUser.NickName,
				UptDate:  NowTime(),
			}
			
			db.Save(&result)
			// Display modified data in JSON message "success"
			c.JSON(200, gin.H{"success": result})
		} else {
			// Display JSON error
			c.JSON(404, gin.H{"error": "User not found"})
		}
		
	} else {
		// Display JSON error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}
	
	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"username\": \"wade\", \"UserPwd\": \"123456\", \"nickname\": \"shandianxia\", \"uptdate\": \"2017-12-01 15:39:29\" }" http://localhost:8080/api/v1/user/1
	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"username\": \"dwaen wade\", \"UserPwd\": \"god\", \"nickname\": \"shandianxia\" }" http://localhost:8080/api/v1/user/1
}

//删除用户
func DeleteUser(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	
	// Get id user
	id := c.Params.ByName("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
	}
	var user model.User
	// SELECT * FROM user WHERE id = 1;
	db.First(&user, userId)
	
	if user.Id != 0 {
		// DELETE FROM user WHERE id = user.Id
		db.Delete(&user)
		// Display JSON result
		c.JSON(200, gin.H{"success": "User #" + id + " deleted"})
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "User not found"})
	}
	
	// curl -i -X DELETE http://localhost:8080/api/v1/user/1
}

//type User struct {
//	Id       int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
//	UserName string `gorm:"not null" form:"username" json:"username"`
//	UserPwd  string `gorm:"not null" form:"userpwd" json:"userpwd"`
//	NickName string `gorm:"" form:"nickname" json:"nickname"`
//	Uptdate  string `gorm:"not null" form:"uptdate" json:"uptdate"`
//}

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

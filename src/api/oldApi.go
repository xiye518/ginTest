package api

import (
	"model"
	"log"
	"strconv"
	"gopkg.in/gin-gonic/gin.v1"
)

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


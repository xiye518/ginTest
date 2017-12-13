package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"github.com/jinzhu/gorm"
)

var db *sql.DB
/*
	1. 用户注册，并填写用户基本信息，用户基本信息包括用户名、密码、昵称，用户名不得重复
	2. 用户登录
	3. 登陆之后可以查看所有已注册的用户列表，未登录状态不能查看用户列表
*/

type User struct {
	Id       int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	UserName string `gorm:"not null" form:"username" json:"username"`
	UserPwd  string `gorm:"not null" form:"userpwd" json:"userpwd"`
	NickName string `gorm:"" form:"nickname" json:"nickname"`
	UptDate  string `gorm:"not null" form:"uptdate" json:"uptdate"`
	//Uptdate   *time.Time `gorm:"not null" form:"uptdate" json:"uptdate"`
}

func (User) TableName() string {
	return "user"
}

type Token struct {
	Id      int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Value   string `gorm:"not null" form:"value" json:"value"`
	UptDate string `gorm:"not null" form:"uptdate" json:"uptdate"`
}

func init() {
	var err error
	db, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/gorm?charset=utf8")
	//db, err = sql.Open("mysql", "jubao:jubao@tcp(localhost:3306)/gorm?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(100)
	
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	
	log.Println("db connecting success...")
}



/*
建表sql:
CREATE TABLE tb_userInfo(
    id int(9) primary key not null auto_increment,
    user_name char(50) not null,
		user_pwd char(20) not null,
		nick_name char(50),
		upt_date  datetime not null
);
*/
func Query() (users []*User, err error) {
	users = make([]*User, 0)
	rows, err := db.Query("select id,user_name,user_pwd,nick_name,upt_date FROM user")
	if err != nil {
		return users, err
	}
	defer rows.Close()
	for rows.Next() {
		var user *User = &User{}
		err = rows.Scan(&user.Id, &user.UserName, &user.UserPwd, &user.NickName, &user.UptDate)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	
	return users, nil
}

func QueryIsExist(username string) (bool, *User, error) {
	var user *User = &User{}
	rows, err := db.Query("select id,user_name,user_pwd,nick_name,upt_date user where user_name = ?", username)
	if err != nil {
		return false, user, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.UserName, &user.UserPwd, &user.NickName, &user.UptDate)
		if err != nil {
			return false, user, err
		}
		
		return true, user, nil
	}
	
	return false, user, nil
}

func QueryUsernameExist(username string) (success bool) {
	var user *User = &User{}
	rows, err := db.Query("select id,user_name,user_pwd,nick_name,upt_date from user where user_name = ?", username)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.UserName, &user.UserPwd, &user.NickName, &user.UptDate)
		if err != nil {
			return
		}
		
		success = true
		return
	}
	
	return
}

func QueryOne(username, userpwd string) (bool, *User, error) {
	var success = false
	var user *User = &User{}
	//var user User
	rows, err := db.Query("select id,user_name,user_pwd,nick_name,upt_date FROM user where user_name = ? and user_pwd = ?", username, userpwd)
	if err != nil {
		return success, user, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.UserName, &user.UserPwd, &user.NickName, &user.UptDate)
		if err != nil {
			return success, user, err
		}
		
		success = true
		return success, user, nil
	}
	
	return success, user, nil
}

func InsertUserInfo(user *User) (err error) {
	//用户注册调用	用户名、密码、昵称，用户名不得重复
	_, err = db.Exec(`insert into user  (user_name,user_pwd,nick_name,upt_date)  values(?,?,?,localtime())`, user.UserName, user.UserPwd, user.NickName)
	if err != nil {
		return err
	}
	
	return nil
}

func Register(username, userpwd, nickname string) (err error) {
	//用户注册调用	用户名、密码、昵称，用户名不得重复
	_, err = db.Exec(`insert into user  (user_name,user_pwd,nick_name,upt_date)  values(?,?,?,localtime())`, username, userpwd, nickname)
	if err != nil {
		return err
	}
	
	return nil
}

func QueryGorm(username string) (User) {
	var err error
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/gorm?charset=utf8")
	if err != nil {
		panic("failed to connect database")
	}
	//defer db.Close()
	db.LogMode(true)
	
	var user User
	//db.First(&user, "user_name = ?", username)
	db.Where("user_name = ?", username).First(&user)
	return user
}

func InsertToken() (string) {
	
	return ""
}

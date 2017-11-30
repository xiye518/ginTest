package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB
/*
	1. 用户注册，并填写用户基本信息，用户基本信息包括用户名、密码、昵称，用户名不得重复
	2. 用户登录
	3. 登陆之后可以查看所有已注册的用户列表，未登录状态不能查看用户列表
*/

func init() {
	var err error
	db, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/test?charset=utf8")
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
    username char(50) not null,
		userpwd char(20) not null,
		nickname char(50),
		uptdate  datetime not null
);
*/
func Query() (users []*User, err error) {
	users = make([]*User, 0)
	rows, err := db.Query("select id,username,userpwd,nickname,uptdate FROM tb_userinfo")
	if err != nil {
		return users, err
	}
	defer rows.Close()
	for rows.Next() {
		var user *User = &User{
		
		}
		err = rows.Scan(&user.Id, &user.UserName, &user.UserPwd, &user.NickName, &user.UptDate)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	
	return users, nil
}

func QueryIsExist(username string) (bool, *User, error) {
	var user *User
	rows, err := db.Query("select id,username,userpwd,nickname,uptdate FROM tb_userinfo where username = ?", username)
	if err != nil {
		return false, user, err
	}
	defer rows.Close()
	for rows.Next() {
		var user *User = &User{
		
		}
		err = rows.Scan(&user.Id, &user.UserName, &user.UserPwd, &user.NickName, &user.UptDate)
		if err != nil {
			return false, user, err
		}
		
		return true, user, nil
	}
	
	return false, user, nil
}

func QueryOne(username string) ( *User, error) {
	var user *User
	rows, err := db.Query("select id,username,userpwd,nickname,uptdate FROM tb_userinfo where username = ?", username)
	if err != nil {
		return  user, err
	}
	defer rows.Close()
	for rows.Next() {
		var user *User = &User{
		
		}
		err = rows.Scan(&user.Id, &user.UserName, &user.UserPwd, &user.NickName, &user.UptDate)
		if err != nil {
			return  user, err
		}
		
		return  user, nil
	}
	
	return  user, nil
}

func InsertUserInfo(user *User) (err error) {
	//用户注册调用	用户名、密码、昵称，用户名不得重复
	_, err = db.Exec(`insert into tb_userInfo  (UserName ,userPwd,nickName,uptdate)  values(?,?,?,localtime())`, user.UserName, user.UserPwd, user.NickName)
	if err != nil {
		return err
	}
	
	return nil
}

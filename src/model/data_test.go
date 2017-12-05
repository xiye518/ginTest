package model

import (
	"testing"
)

func TestQuery(t *testing.T) {
	users, err := Query()
	if err != nil {
		t.Fatal(err)
	}
	
	for _, user := range users {
		t.Logf("%v\n", user)
	}
	
}

func TestQueryIsExist(t *testing.T) {
	isExist,_, err := QueryIsExist("xiyang")
	if err != nil {
		t.Fatal(err)
	}
	
	t.Log(isExist)
}

func TestQueryUsernameExist(t *testing.T) {
	isExist:=QueryUsernameExist("xiye1")
	t.Log(isExist)
}

func TestInsertUserInfo(t *testing.T) {
	user := &User{
		UserName: "xiye",
		UserPwd:  "123",
		NickName: "godfather",
	}
	QueryIsExist(user.UserName)
	
	err := InsertUserInfo(user)
	if err != nil {
		t.Fatal(err)
	}
	
	_,user,err = QueryOne(user.UserName,user.UserPwd)
	t.Logf("新增成功，用户信息：%v",user)
}

func TestQueryOne(t *testing.T) {
	_,user,err := QueryOne("xiye","123")
	if err!=nil{
		t.Fatal(err)
	}
	t.Logf("查询成功，用户信息：%v",user)
}

func TestQueryGorm(t *testing.T) {
	user:=QueryGorm("xiye")
	t.Log(user)
}


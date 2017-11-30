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

func TestInsertUserInfo(t *testing.T) {
	user := &User{
		UserName: "xiyang",
		UserPwd:  "123456",
		NickName: "哈卖批",
	}
	err := InsertUserInfo(user)
	if err != nil {
		t.Fatal(err)
	}
	
	user,err = QueryOne(user.UserName)
	t.Logf("新增成功，用户信息：%v",user)
}

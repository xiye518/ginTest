package API

import "testing"

func TestIsExist(t *testing.T) {
	key := "james"
	result := IsExist(key)
	t.Log(key+" 已注册？", result)
	
	key = "111"
	result = IsExist(key)
	t.Log(key+" 已注册？", result)
}

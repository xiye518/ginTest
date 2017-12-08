### 1. 注册
__说明__
curl -i -X POST -H "Content-Type: application/json" -d "{ \"username\": \"admin\", \"userpwd\": \"123\"}" http://localhost:8080/api/v1/reg
	请求参数：{"username":"用户名","userpwd":"密码","nickname":"昵称"}
	请求方法：POST
	


__返回值__

	"success/error" : "返回信息"
	

__示例__

	请求:
	http://localhost:8080/api/v1/reg
	参数:
	{
		"username" : "admin"
		"password" : "123456"
		"nickname" : "xixi"
	}
	返回:
	{
    	"success": "register success!",
	}


### 2. 登录
__说明__
curl -i -X POST -H "Content-Type: application/json" -d "{ \"username\": \"admin\", \"userpwd\": \"123\"}" http://localhost:8080/api/v1/login
	请求参数：{"username":"用户名","userpwd":"密码"}
	请求方法：POST
	


__返回值__

	"success/error" : "返回信息"
	

__示例__

	请求:
	http://localhost:8080/api/v1/login
	参数:
	{
		"username" : "admin"
		"userpwd" : "123"
	}
	返回:
	{
    	"success": "login success!",
	}

### 2. 查询用户列表
__说明__
  需登录后读取判断token值才能判断是否显示用户列表
	请求参数：无
	请求方法：GET
	


__返回值__

	"success/error" : "返回信息"
	

__示例__

	请求:
	http://localhost:8080/api/v1/show
	参数:{ }
	返回:
	{
    	"error": "登录后才能查看用户列表!",
	}

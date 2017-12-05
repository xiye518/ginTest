package model

//type User struct {
//	Id       int
//	UserName string
//	UserPwd  string
//	NickName string
//	UptDate  string
//}

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


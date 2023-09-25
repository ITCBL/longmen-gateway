package model

import "time"

type User struct {
	BaseModel
	Mobile   string     `json:"mobile" gorm:"index:idx_mobile;unique;type:varchar(11);not null" description:"手机号 定义唯一索引 | 限制手机长度 | 不为空"`
	Password string     `json:"password" gorm:"type:varchar(100);not null" description:"密码信息：秘钥+加密方式+加密后的密码"`
	UserName string     `json:"user_name" gorm:"type:varchar(20);not null" description:"名称"`
	NickName string     `json:"nick_name" gorm:"type:varchar(20)" description:"昵称"`
	Birthday *time.Time `json:"birthday" gorm:"type:datetime" description:"生日 使用指针类型，避免空值不更新）"`
	Gender   string     `json:"gender" gorm:"default:male;type:varchar(6)" description:"female表示女，male表示男"`
	Role     int        `json:"role" gorm:"default:1;type:int" description:"1表示普通用户，2表示管理员"`
}

func (t *User) TableName() string {
	return "user"
}

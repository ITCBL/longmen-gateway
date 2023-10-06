package model

import (
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"gopkg.in/errgo.v2/errors"
	"longmen-gateway/global"
	"strings"
	"time"
)

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

func (t *User) GetUserByMobile(mobile string) (*User, error) {
	var user User

	result := global.DB.Where(&User{Mobile: mobile}).First(&user)
	if result.RowsAffected == 0 {
		return nil, errors.New("Not Found!")
	}

	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil

}

// 密码加密
var options = &password.Options{
	SaltLen:      16,
	Iterations:   100,
	KeyLen:       32,
	HashFunction: sha512.New,
}

func (t *User) CheckPassword(loginPassword, encryptedPassword string) (bool, error) {
	passwordInfo := strings.Split(encryptedPassword, "$")
	check := password.Verify(loginPassword, passwordInfo[2], passwordInfo[3], options)
	return check, nil
}

func (t *User) CreateUser(mobile, pwd, NickName string) (*User, error) {
	var user User
	result := global.DB.Where(&User{Mobile: mobile}).First(&user)
	if result.RowsAffected == 1 {
		return nil, errors.New("用户已存在")
	}
	user.Mobile = mobile
	user.NickName = NickName
	user.Password = PasswordEncryption(pwd) // 密码加密

	// 入库
	result = global.DB.Create(&user)
	if result.Error != nil {
		return nil, errors.New("创建用户失败")
	}
	return &user, nil
}

func PasswordEncryption(PassWord string) string {
	salt, encoderPwd := password.Encode(PassWord, options)
	newPassword := fmt.Sprintf("$pbkdf2--sha512$%s$%s", salt, encoderPwd) // 新密码=加密算法+盐值+原密码密文
	return newPassword
}

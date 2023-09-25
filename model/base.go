package model

import "time"

/*
gorm定义 primary_key、autoIncrement、type、default、not null
*/

// BaseModel 基础模型
type BaseModel struct {
	ID        int64      `json:"id" gorm:"primary_key;autoIncrement;type:bigint" description:"自增主键"`
	CreatedAt time.Time  `json:"-" gorm:"type:datetime" description:"创建时间"` // 通过column可以自定义数据字段名称，否则默认规则生成
	UpdatedAt time.Time  `json:"-" gorm:"type:datetime" description:"更新时间"`
	DeleteAt  *time.Time `json:"-" gorm:"type:datetime" description:"删除时间"`
	IsDelete  bool       `json:"-" gorm:"type:bool;default:false" description:"是否删除"`
}

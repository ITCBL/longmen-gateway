package model

import "time"

/*
gorm定义 primary_key、autoIncrement、type、default、not null、column、comment
*/

// BaseModel 基础模型
type BaseModel struct {
	ID        int64      `json:"id" gorm:"primary_key;autoIncrement;type:bigint" description:"自增主键"` // 数据字段名称可以通过【column】自定义，否则将生成默认规则
	CreatedAt time.Time  `json:"-" gorm:"type:datetime" description:"创建时间"`
	UpdatedAt time.Time  `json:"-" gorm:"type:datetime" description:"更新时间"`
	DeleteAt  *time.Time `json:"-" gorm:"type:datetime" description:"删除时间"`
	IsDelete  bool       `json:"-" gorm:"type:bool;default:false" description:"是否删除"`
}

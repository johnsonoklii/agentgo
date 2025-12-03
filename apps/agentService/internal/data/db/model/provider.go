package model

import "time"

type Provider struct {
	ID          int64      `gorm:"primarykey"`
	ProviderID  string     `gorm:"type:varchar(255);uniqueIndex;comment:'provider id'"`
	UID         string     `gorm:"type:varchar(255);index;comment:'用户ID'"`
	Name        string     `gorm:"type:varchar(255);comment:'服务商名称'"`
	Description string     `gorm:"type:text;comment:'服务商描述'"`
	Protocol    string     `gorm:"type:varchar(100);comment:'服务商类型'"`
	Status      bool       `gorm:"type:tinyint(1);comment:'状态:1-启用,0-禁用'"`
	Config      string     `gorm:"type:text;comment:'配置值'"`
	CreatedAt   time.Time  `gorm:"column:created_at"`
	UpdatedAt   time.Time  `gorm:"column:updated_at"`
	DeletedAt   *time.Time `gorm:"column:deleted_at;index"`
}

func (Provider) TableName() string {
	return "providers"
}

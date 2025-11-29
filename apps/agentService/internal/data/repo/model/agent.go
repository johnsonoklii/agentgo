package model

type Agent struct {
	Id               int64                                   `gorm:"primarykey"`
	Name             string                                  `gorm:"type:varchar(255) comment '名称'"`
	Avatar           string                                  `gorm:"type:varchar(255) comment '头像URL'"`
	Description      string                                  `gorm:"type:text comment '描述'"`
	SystemPrompt     string                                  `gorm:"type:text comment '系统提示'"`
	WelcomeMessage   string                                  `gorm:"type:text comment '欢迎信息'"`
	ToolIds          []string                                `gorm:"type:json comment '工具ID'"`
	KnowledgeBaseIds []string                                `gorm:"type:json comment '知识库ID'"`
	PublishedVersion string                                  `gorm:"type:varchar(36) comment '发布版本'"`
	Enabled          bool                                    `gorm:"type:tinyint(1) comment '是否启用'"`
	UserId           int64                                   `gorm:"type:bigint(20) comment '用户ID'"`
	ToolPresetParams map[string]map[string]map[string]string `gorm:"type:text comment '工具参数'"`
	MultiModel       bool                                    `gorm:"type:tinyint(1) comment '是否多模型'"`
	CreatedAt        int64                                   `gorm:"column:created_at"`
	UpdatedAt        int64                                   `gorm:"column:created_at"`
	DeletedAt        int64                                   `gorm:"column:created_at"`
}

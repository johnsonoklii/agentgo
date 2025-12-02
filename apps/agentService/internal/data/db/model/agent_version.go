package model

import (
	"time"
)

// AgentVersion Agent版本实体类，代表一个Agent的发布版本
type AgentVersion struct {
	ID               int64            `gorm:"primarykey"`
	AgentVersionID   string           `gorm:"type:varchar(255) comment 'agent version id'"`
	AgentID          string           `gorm:"type:varchar(36);not null;index:idx_agent_versions_agent_id;comment '关联的Agent ID'"`
	Name             string           `gorm:"type:varchar(255);not null;comment 'Agent名称'"`
	Avatar           string           `gorm:"type:varchar(255);comment 'Agent头像URL'"`
	Description      string           `gorm:"type:text;comment 'Agent描述'"`
	VersionNumber    string           `gorm:"type:varchar(20);not null;comment '版本号，如1.0.0'"`
	SystemPrompt     string           `gorm:"type:text;comment 'Agent系统提示词'"`
	WelcomeMessage   string           `gorm:"type:text;comment '欢迎消息'"`
	ToolIds          []string         `gorm:"type:json;comment 'Agent可使用的工具ID列表，JSON数组格式'"`
	KnowledgeBaseIds []string         `gorm:"type:json;comment '关联的知识库ID列表，JSON数组格式'"`
	ChangeLog        string           `gorm:"type:text;comment '版本更新日志'"`
	PublishStatus    int              `gorm:"type:tinyint;default:1;comment '发布状态：1-审核中, 2-已发布, 3-拒绝, 4-已下架'"`
	RejectReason     string           `gorm:"type:text;comment '审核拒绝原因'"`
	ReviewTime       *time.Time       `gorm:"comment '审核时间'"`
	PublishedAt      *time.Time       `gorm:"comment '发布时间'"`
	UserID           string           `gorm:"type:varchar(36);not null;index:idx_agent_versions_user_id;comment '创建者用户ID'"`
	ToolPresetParams ToolPresetParams `gorm:"type:json;comment '工具预设参数'"`
	MultiModal       bool             `gorm:"type:tinyint(1);default:false;comment '是否支持多模态'"`
	CreatedAt        time.Time        `gorm:"comment '创建时间'"`
	UpdatedAt        time.Time        `gorm:"comment '更新时间'"`
	DeletedAt        *time.Time       `gorm:"comment '逻辑删除时间'"`
}

// TableName 指定表名
func (AgentVersion) TableName() string {
	return "agent_versions"
}

// PublishStatusEnum 发布状态枚举
const (
	PublishStatusReviewing = 1 // 审核中
	PublishStatusPublished = 2 // 已发布
	PublishStatusRejected  = 3 // 拒绝
	PublishStatusOffline   = 4 // 已下架
)

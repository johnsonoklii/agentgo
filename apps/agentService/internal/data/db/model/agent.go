package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type Agent struct {
	ID               int64            `gorm:"primarykey"`
	AgentID          string           `gorm:"type:varchar(255) comment 'agent id'"`
	Name             string           `gorm:"type:varchar(255) comment '名称'"`
	Avatar           string           `gorm:"type:varchar(255) comment '头像URL'"`
	Description      string           `gorm:"type:text comment '描述'"`
	SystemPrompt     string           `gorm:"type:text comment '系统提示'"`
	WelcomeMessage   string           `gorm:"type:text comment '欢迎信息'"`
	ToolIds          []string         `gorm:"type:json comment '工具ID'"`
	KnowledgeBaseIds []string         `gorm:"type:json comment '知识库ID'"`
	PublishedVersion string           `gorm:"type:varchar(36) comment '发布版本'"`
	Enabled          bool             `gorm:"type:tinyint(1) comment '是否启用'"`
	UID              string           `gorm:"type:bigint(20) comment '用户ID'"`
	ToolPresetParams ToolPresetParams `gorm:"type:json comment '工具参数'"`
	MultiModel       bool             `gorm:"type:tinyint(1) comment '是否多模型'"`
	CreatedAt        time.Time        `gorm:"column:created_at"`
	UpdatedAt        time.Time        `gorm:"column:updated_at"`
	DeletedAt        *time.Time       `gorm:"column:deleted_at"`
}

type ToolPresetParams map[string]map[string]map[string]string

func (m ToolPresetParams) Value() (driver.Value, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return string(b), nil
}

func (m *ToolPresetParams) Scan(value interface{}) error {
	if value == nil {
		*m = ToolPresetParams{}
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("invalid scan source")
	}
	return json.Unmarshal(bytes, m)
}

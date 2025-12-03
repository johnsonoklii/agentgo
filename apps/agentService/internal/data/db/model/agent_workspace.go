package model

import "time"

type TokenOverflowStrategyEnum int

const (
	None TokenOverflowStrategyEnum = iota
	SLIDING_WINDOW
	SUMMARIZE
)

type LLMModalConfig struct {
	ModalId                   string
	Temperature               float64
	TopP                      float64
	MaxTokens                 int
	ReserveRatio              float64
	SummaryThreshold          int
	TokenOverflowStrategyEnum TokenOverflowStrategyEnum
}

// AgentWorkspace 工作区与Agent的关联关系
type AgentWorkspace struct {
	ID             int64          `gorm:"primarykey"`
	WorkspaceID    string         `gorm:"type:varchar(255);index;comment '工作区ID'"`
	AgentID        string         `gorm:"type:varchar(255);index;comment 'Agent ID'"`
	UID            string         `gorm:"type:varchar(255);index;comment '用户ID'"`
	LLMModalConfig LLMModalConfig `gorm:"type:json;comment 'LLM模型配置'"`
	CreatedAt      time.Time      `gorm:"column:created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at"`
	DeletedAt      *time.Time     `gorm:"column:deleted_at"`
}

func (AgentWorkspace) TableName() string {
	return "agent_workspace"
}

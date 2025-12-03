package entity

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
	WorkspaceID    string
	AgentID        string
	UID            string
	LLMModalConfig LLMModalConfig
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

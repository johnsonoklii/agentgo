package entity

import "time"

type Agent struct {
	AgentID          string
	Name             string
	Avatar           string
	Description      string
	SystemPrompt     string
	ToolIds          []string
	KnowledgeBaseIds []string
	PublishedVersion string
	Enabled          bool
	UID              string
	ToolPresetParams map[string]map[string]map[string]string
	MultiModel       bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

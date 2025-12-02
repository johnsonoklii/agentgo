package entity

import "time"

type AgentVersion struct {
	ID               int64
	AgentVersionID   string
	AgentID          string
	Name             string
	Avatar           string
	Description      string
	VersionNumber    string
	SystemPrompt     string
	WelcomeMessage   string
	ToolIds          []string
	KnowledgeBaseIds []string
	ChangeLog        string
	PublishStatus    int
	RejectReason     string
	ReviewTime       *time.Time
	PublishedAt      *time.Time
	UserID           string
	ToolPresetParams map[string]map[string]map[string]string
	MultiModal       bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

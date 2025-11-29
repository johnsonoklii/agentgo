package entity

type Agent struct {
	AgentId          int64
	Name             string
	Avatar           string
	Description      string
	SystemPrompt     string
	ToolIds          []string
	KnowledgeBaseIds []string
	PublishedVersion string
	Enabled          bool
	UserId           int64
	UserName         string
	ToolPresetParams map[string]map[string]map[string]string
	MultiModel       bool
	CreatedAt        int64
	UpdatedAt        int64
}

package model

import "time"

// AgentWorkspace 工作区与Agent的关联关系
type AgentWorkspace struct {
	ID          int64      `gorm:"primarykey"`
	WorkspaceID string     `gorm:"type:varchar(255);index;comment '工作区ID'"`
	AgentID     string     `gorm:"type:varchar(255);index;comment 'Agent ID'"`
	UID         string     `gorm:"type:varchar(255);index;comment '用户ID'"`
	CreatedAt   time.Time  `gorm:"column:created_at"`
	UpdatedAt   time.Time  `gorm:"column:updated_at"`
	DeletedAt   *time.Time `gorm:"column:deleted_at"`
}

func (AgentWorkspace) TableName() string {
	return "agent_workspace"
}

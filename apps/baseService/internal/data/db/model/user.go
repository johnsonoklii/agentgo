package model

type User struct {
	ID       int64  `gorm:"primarykey"`
	UID      string `gorm:"uniqueIndex;type:varchar(25);comment:用户ID"`
	Mobile   string `gorm:"index:idx_mobile;unique;type:varchar(11) comment '手机号码，用户唯一标识';not null"`
	UserName string `gorm:"type:varchar(25) comment '用户昵称'"`
	Password string `gorm:"type:varchar(100);not null "`
	Gender   int8   `gorm:"type:tinyint(1) comment '性别 1:男 2:女 0:未知'"`
	Deleted  bool   `gorm:"type:tinyint(1) comment '是否删除'"`
	Status   int8   `gorm:"type:tinyint(1) comment '状态 1:正常 2:禁用 0:未知'"`

	CreatedAt int64 `gorm:"column:created_at"`
	UpdatedAt int64 `gorm:"column:updated_at"`
}

package entity

import "time"

type User struct {
	UID string

	UserName string // nickname
	Mobile   string
	Password string
	Gender   int8
	Deleted  bool

	Token string

	CreatedAt time.Time // creation time
	UpdatedAt time.Time // update time
}

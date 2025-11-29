package entity

type User struct {
	UID string

	UserName string // nickname
	Mobile   string
	Password string
	Gender   int8
	Deleted  bool

	Token string

	CreatedAt int64 // creation time
	UpdatedAt int64 // update time
}

package model

type User struct {
	Id       string `xorm:"pk" json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func (c User) TableName() string {
	return "users"
}

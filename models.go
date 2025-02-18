package main

type Status string

const (
	Read   Status = "read"
	Done   Status = "done"
	ToRead Status = "to_read"
)

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	UserName string `json:"username"`
	Password string `json:"-"`
}
type Book struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Status Status `json:"status" gorm"default:to_read"`
	UserId int    `json:"user_id"`
}

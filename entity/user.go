package entity

type User struct {
	UserID   int64  `gorm:"user_id"`
	UserName string `gorm:"user_name"`
	Password string `gorm:"password"`
}

package user


type User struct {
	ID			int64 		`gorm:"index" json:"id"`
	UserName    string     	`gorm:"not null" json:"username"`
	Password    string     	`gorm:"not null" json:"password"`
}
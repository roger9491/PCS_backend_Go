package calendar



type Calendar struct {
	Id		int64 	`gorm:"index" json:"id"`
	Date	string 	`gorm:"not null" josn:"date"`
	Title 	string 	`gorm:"not null" json:"title"`
	Content string	`json:"content"`
	UserId  int64   `gorm:"column:userid;not null" json:"userId"`
}
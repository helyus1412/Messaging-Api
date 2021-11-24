package models

type User struct {
	ID      int `gorm:"PrimaryKey"`
	Name    string
	Message []Messages `gorm:"ForeignKey:UserID"`
}

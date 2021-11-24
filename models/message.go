package models

type Messages struct {
	ID      int `gorm:"PrimaryKey"`
	Message string
	UserID  int
	From    string
	To      string
}

package models

type User struct {
	ID      int    `gorm:"column:id; primary key"`
	Name    string `gorm:"column:name"`
	DOB     string `gorm:"column:dob"`
	Contact string `gorm:"column:contact; unique"`
}

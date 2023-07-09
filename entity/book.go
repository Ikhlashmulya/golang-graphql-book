package entity

type Book struct {
	Id          string `gorm:"primaryKey"`
	Title       string 
	Description string 
	Author      string 
}

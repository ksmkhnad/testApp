package models

type User struct {
	ID       uint   `gorm:"primary_key"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

type Product struct {
	ID          uint    `gorm:"primary_key"`
	Name        string  `gorm:"unique;not null"`
	Description string  `gorm:"not null"`
	Price       float64 `gorm:"not null"`
}

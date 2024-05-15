package model

import "time"

type Employee struct {
	Id        string    `json:"id" gorm:"column:id"`
	Name      string    `json:"name" gorm:"column:name"`
	Dob       time.Time `json:"dob" gorm:"column:dob"`
	Email     string    `json:"email" gorm:"column:email"`
	Phone     string    `json:"phone" gorm:"column:phone"`
	CitizenId string    `json:"citizen_id" gorm:"column:citizen_id"`
	Address   string    `json:"address" gorm:"column:address"`
}

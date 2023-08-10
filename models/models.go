package models

import "time"

type KeyValue struct {
	ID        uint      `gorm:"primary_key"`
	Key       string    `gorm:"type:varchar(255);unique;not null"`
	Value     string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"type:datetime;autoCreateTime"`
}

type User struct {
	ID       uint   `gorm:"primary_key"`
	Username string `gorm:"type:varchar(255);unique;not null"`
	Password string `gorm:"type:varchar(255);not null"`
	Token    string `gorm:"not null"`
}

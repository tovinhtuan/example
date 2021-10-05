package models

import (
	"time"
)

type User struct {
	Id         int64       `gorm:"type:integer"`
	FullName   string      `gorm:"type:character varying(255)"`
	UserName   string      `gorm:"type:character varying(255)"`
	Gender     string      `gorm:"type:character varying(10)"`
	Birthday   time.Time   `gorm:"type:timestamp"`
	AuthTokens []AuthToken `gorm:"foreignKey:UserId;references:Id"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
type AuthToken struct {
	Id        int64  `gorm:"type:integer"`
	UserId    int64  `gorm:"type:integer"`
	Token     string `gorm:"type:character varying(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

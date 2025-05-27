package dto

import "time"

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type LoginResponse struct {
	Username   string    `gorm:"NOT NULL unique" json:"username"`
	Email      string    `gorm:"NOT NULL unique" json:"email"`
	Avatar     string    `gorm:"NOT NULL; default:''" json:"avatar"`
	Management bool      `gorm:"default:false" json:"management"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"createdAt"`
	CreateName string    `gorm:"NOT NULL" json:"createName"`
}

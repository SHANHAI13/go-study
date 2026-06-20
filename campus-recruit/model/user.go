package model

import "time"

// 用户表
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"size:50;uniqueIndex;not null" json:"username"`
	Password  string    `gorm:"size:255;not null" json:"-"` // json:"-" 不返回密码
	Role      string    `gorm:"size:20;default:student" json:"role"` // student / company / admin
	Phone     string    `gorm:"size:20" json:"phone"`
	Email     string    `gorm:"size:100" json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// 注册请求
type RegisterReq struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=6,max=20"`
	Role     string `json:"role" binding:"required,oneof=student company"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

// 登录请求
type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 登录响应
type LoginResp struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

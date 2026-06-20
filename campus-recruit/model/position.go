package model

import "time"

// 职位表
type Position struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	CompanyID   uint      `gorm:"index;not null" json:"company_id"`
	Title       string    `gorm:"size:100;not null" json:"title"`     // 职位名称
	Description string    `gorm:"type:text" json:"description"`       // 职位描述
	Location    string    `gorm:"size:50" json:"location"`            // 工作地点
	Salary      string    `gorm:"size:50" json:"salary"`              // 薪资范围
	Tag         string    `gorm:"size:100" json:"tag"`                // 标签（实习/校招/社招）
	Status      int       `gorm:"default:1" json:"status"`            // 1=发布中 0=已下架
	ViewCount   int       `gorm:"default:0" json:"view_count"`        // 浏览量
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// 职位列表查询参数
type PositionListReq struct {
	Page     int    `form:"page" binding:"omitempty,min=1"`
	PageSize int    `form:"page_size" binding:"omitempty,min=1,max=50"`
	Keyword  string `form:"keyword"`
	Location string `form:"location"`
	Tag      string `form:"tag"`
}

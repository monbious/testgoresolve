package model

import "time"

type Department struct {
	DeptId int		`gorm:"index;AUTO_INCREMENT;column:dept_id;primary_key"`
	DeptName string		`gorm:"column:dept_name"`
	UpdateTime time.Time	`gorm:"column:update_time"`
}

func (dept *Department) TableName() string {
	return "tbl_dept"
}
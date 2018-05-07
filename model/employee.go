package model

import "time"

type Employee struct {
	EmpId int		`gorm:"index;AUTO_INCREMENT;column:emp_id;primary_key"`
	EmpName string		`gorm:"column:emp_name"`
	Gender string		`gorm:"column:gender"`
	Email string		`gorm:"column:email"`
	Dept *Department	`gorm:"-"`
	DID int			`gorm:"column:d_id"`
	Status string 		`gorm:"column:status;type:enum('deleted','normal');default:'normal'"`
	UpdateTime time.Time	`gorm:"column:update_time"`
}

func (e *Employee) TableName() string {
return "tbl_emp"
}


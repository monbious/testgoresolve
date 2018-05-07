package dao

import (

	"testgoresolve/model"
	"testgoresolve/dbUtils"
)


func GetDepartments() []model.Department {

	db := dbUtils.GetDB()
	defer db.Close()
	var depts []model.Department
	db.Find(&depts)
	return depts
}
package dao

import (
	"testgoresolve/model"
	"testgoresolve/dbUtils"
	"strings"
	"strconv"
)

func DelEmp(e *model.Employee) {
	var ids []int
	for _, id := range strings.Split(e.EmpName, "-"){
		eid,_ := strconv.Atoi(id)
		ids = append(ids, eid)
	}
	db := dbUtils.GetDB()
	defer db.Close()
	tx := db.Begin()
	db.Delete(model.Employee{},"emp_id in (?)", ids)
	tx.Commit()

}

func UpdateEmp(e *model.Employee) bool{
	db := dbUtils.GetDB()
	defer db.Close()
	tx := db.Begin()
	db.Model(e).Updates(e)
	tx.Commit()
	return true
}

func SaveEmp(e *model.Employee) bool {
	db := dbUtils.GetDB()
	defer db.Close()
	tx := db.Begin()
	db.Create(e)
	tx.Commit()
	return true
}
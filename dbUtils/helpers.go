package dbUtils

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
	"github.com/golang/glog"
	"testgoresolve/model"
	"time"
	"github.com/pborman/uuid"
	"gitlab.wallstcn.com/wscnbackend/ivankastd/toolkit/strutil"
	"testgoresolve/conf"
)

func CheckTables() {
	db:= GetDB()
	defer db.Close()
	emp := &model.Employee{}
	dept := model.Department{}
	boo := db.HasTable(emp.TableName())
	if  !boo {
		//fmt.Println("没有tbl_emp")
		db.CreateTable(&model.Employee{})
		for i := 0; i<500; i++  {
			tx := db.Begin()
			uuid := uuid.New()
			emp := &model.Employee{}

			empname := strutil.Sub(uuid,0,5)
			emp.EmpName = empname
			emp.Gender = "M"
			emp.Email = empname + "@163.com"
			emp.UpdateTime = time.Now().Local()
			emp.DID = 2
			db.Create(emp)
			tx.Commit()
		}
	}
	boo2 := db.HasTable(dept.TableName())

	if !boo2 {
		//fmt.Println("没有tbl_dept")
		db.CreateTable(&model.Department{})
		strs := []string{"测试部", "技术部", "开发部"}
		for i := 0; i < len(strs); i++ {
			tx := db.Begin()
			dept := &model.Department{
				DeptName:strs[i],
				UpdateTime:time.Now().UTC().Local(),
			}
			db.Create(dept)
			tx.Commit()
		}
	}
}

func GetDB() *gorm.DB {

	conf := conf.Config
	dbconf := conf.DB
	username := dbconf.Username
	password := dbconf.Password
	host := dbconf.Host
	port := dbconf.Port
	dbname := dbconf.DBname
	dburl := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=true", username, password, host, port, dbname)
	db, err := gorm.Open("mysql", dburl)
	if err != nil {
		glog.Errorf("open db failed :%s", err)
	}
	return db
}
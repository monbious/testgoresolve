package handler

import (
	"golang.org/x/net/context"

	"testgopb/pb/pbemp"
	"testgoresolve/model"
	"testgoresolve/dao"
	"time"
)

type Employees struct {

}

func (e *Employees) Update(ctx context.Context, req *pbemp.EmployeesReq, res *pbemp.EmployeesRes) error{

	memp := &model.Employee{
		EmpId:int(req.EmpId),
		Email:req.Email,
		Gender:req.Gender,
		DID:int(req.DeptId),
		UpdateTime:time.Now().Local(),
	}
	boo := dao.UpdateEmp(memp)
	*res = pbemp.EmployeesRes{
		Boo:boo,
	}

	return nil
}

func (e *Employees) DelEmp(ctx context.Context, req *pbemp.EmployeesReq, res *pbemp.EmployeesRes) error{

	emp := &model.Employee{
		EmpName:req.EmpName,
	}
	dao.DelEmp(emp)

	return nil
}
func (e *Employees) GetEmployees(ctx context.Context, req *pbemp.EmployeesReq, res *pbemp.EmployeesRes) error{

	return nil
}

func (e *Employees) SaveEmp(ctx context.Context, req *pbemp.EmployeesReq, res *pbemp.EmployeesRes) error {

	E := &model.Employee{
		EmpName:req.EmpName,
		Gender:req.Gender,
		Email:req.Email,
		DID:int(req.DeptId),
		UpdateTime:time.Now().Local(),
		Dept:&model.Department{
				DeptId:123,
				DeptName:"23445",
		},

	}

	boo := dao.SaveEmp(E)

	*res = pbemp.EmployeesRes{
		Boo:boo,
	}

	return nil
}
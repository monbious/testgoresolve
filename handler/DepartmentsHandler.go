package handler

import (
	"testgopb/pb/pbdept"
	"testgoresolve/dao"
	"golang.org/x/net/context"
)

type Departments struct {

}

func (d *Departments) GetDepartments(ctx context.Context, req *pbdept.DepartmentsReq, res *pbdept.DepartmentsRes) error {

	mdepts := dao.GetDepartments()
	var pdeptlist []*pbdept.Department

	if len(mdepts) > 0 {
		for _, mdept := range mdepts {
			d := &pbdept.Department{
				DeptId:int32(mdept.DeptId),
				DeptName:mdept.DeptName,
			}
			pdeptlist = append(pdeptlist, d)
		}
	}
	*res = pbdept.DepartmentsRes{
		Depts:pdeptlist,
	}

	return nil

}
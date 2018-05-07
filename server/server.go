package server

import (
	"testgopb/pb"
	"testgoresolve/handler"
	"qiniupkg.com/x/log.v7"
	"testgopb/pb/pbemp"
	"testgopb/pb/pbdept"
	"testgoresolve/conf"
)

func RunServer() {

	myService := pb.NewService(&conf.Config.ConfigServiceMicro)
	myService.Init()

	pbemp.RegisterEmployeesHandler(myService.Server(),new(handler.Employees))
	pbdept.RegisterDepartmentsHandler(myService.Server(), new(handler.Departments))

	if err := myService.Run();err != nil {
		log.Println(err)
	}
}
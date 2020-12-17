package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["BeeTestProjec/controllers:MainController"] = append(beego.GlobalControllerRouter["BeeTestProjec/controllers:MainController"],
        beego.ControllerComments{
            Method: "DealConsumeRecord",
            Router: "/block",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BeeTestProjec/controllers:PerformanceController"] = append(beego.GlobalControllerRouter["BeeTestProjec/controllers:PerformanceController"],
        beego.ControllerComments{
            Method: "Cpu",
            Router: "/cpu",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BeeTestProjec/controllers:PerformanceController"] = append(beego.GlobalControllerRouter["BeeTestProjec/controllers:PerformanceController"],
        beego.ControllerComments{
            Method: "Meminfo",
            Router: "/meminfo",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BeeTestProjec/controllers:PerformanceController"] = append(beego.GlobalControllerRouter["BeeTestProjec/controllers:PerformanceController"],
        beego.ControllerComments{
            Method: "Testcase",
            Router: "/testcase",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BeeTestProjec/controllers:ResultController"] = append(beego.GlobalControllerRouter["BeeTestProjec/controllers:ResultController"],
        beego.ControllerComments{
            Method: "ResultMeminfo",
            Router: "/ResultMeminfo",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}

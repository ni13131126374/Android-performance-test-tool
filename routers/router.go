package routers

import (
	"BeeTestProjec/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&controllers.MainController{})
    beego.Router("/", &controllers.MainController{})
	//beego.Router("/meminfo",&controllers.MeminfoController{})
	beego.Include(&controllers.PerformanceController{})
	beego.Include(&controllers.ResultController{})
	//beego.Router("/cpu",&)
}

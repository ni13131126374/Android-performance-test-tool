package main

import (
	_ "BeeTestProjec/routers"
	"github.com/astaxie/beego"
)

func main() {
	//beego.StaticDir["/static"] = "static"
	beego.SetStaticPath("views","views")
	beego.Run()
}


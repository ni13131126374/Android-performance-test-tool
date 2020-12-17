package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

// @router /block [get]
func (ma *MainController) DealConsumeRecord() {
	ma.Ctx.WriteString("Hello World!")
}
func (this *MainController) Get() {
	//this.Data["Website"] = "beego.me"
	//this.Data["Email"] = "astaxie@gmail.com"
	//this.TplName = "index.tpl"
	//dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(dir)
	this.TplName = "testone.html"
}

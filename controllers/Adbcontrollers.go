package controllers

import (
	"BeeTestProjec/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type PerformanceController struct {
	beego.Controller
}

func (this *PerformanceController) URLMapping() {
	this.Mapping("Meminfo", this.Meminfo)
	this.Mapping("Cpu", this.Cpu)
	this.Mapping("Testcase", this.Testcase)
}

type Meminfostruct struct {
	JavaHeap     int
	NativeHeap   int
	Code         int
	Stack        int
	Graphics     int
	PrivateOther int
	System       int
	TOTAL        int
	Activities   int
}

// @router /meminfo [get]
func (this *PerformanceController) Meminfo() {
	PackName := this.GetString("PackName")
	CaseName := this.GetString("Casename")
	if PackName == "" {
		this.Ctx.WriteString("jsoninfo is empty")
	}
	Meminfo := models.AdbShellDumpsysMeminfo(PackName,CaseName)
	MeminfoJson, err := json.Marshal(Meminfo)
	if err == nil {
		fmt.Println("MapToJsonDemo err: ", err)
	}
	var meminfostruct Meminfostruct
	json.Unmarshal([]byte(MeminfoJson), &meminfostruct)
	this.Data["json"] = meminfostruct
	this.ServeJSON()
	return
}

type CpuStruct struct {
	Cpunum float64
}

// @router /cpu [get]
func (this *PerformanceController) Cpu() {
	PackName := this.GetString("PackName")
	CaseName := this.GetString("Casename")
	if PackName == "" {
		this.Ctx.WriteString("jsoninfo is empty")
	}
	Meminfo := models.AdbShellTop(PackName,CaseName)
	//Cpumap:=make(map[string]float64)
	//Cpumap["Cpunum"]=Meminfo
	CpuJson := &CpuStruct{Meminfo}
	this.Data["json"] = CpuJson
	this.ServeJSON()
	return
	//this.Ctx.WriteString(PackName)
}

// @router /testcase [get]
func (this *PerformanceController) Testcase() {
	mystruct := &JSONStruct{0, "hello"}

	this.Data["json"] = mystruct
	this.ServeJSON()
}

type JSONStruct struct {
	Code int
	Msg  string
}

package controllers

import (
	"BeeTestProjec/models"
	"fmt"
	"github.com/astaxie/beego"
	"time"
)
type ResultController struct {
	beego.Controller
}
type MeminfoResultstruct struct {
	TOTALMAP  map[int]int
	Meminfolist []int
}

var strtime=time.Now().Format("2006_01_02")
func (c *ResultController) URLMapping() {
	c.Mapping("ResultMeminfo", c.ResultMeminfo)
}
//@router /ResultMeminfo [get]
func (this *ResultController) ResultMeminfo() {
	PackName := this.GetString("PackName")
	CaseName := this.GetString("Casename")
	fmt.Printf(PackName,CaseName)
	csvname:=CaseName+"_"+strtime+"_Meminfo"
	TOTALMAP,Meminfolist:=models.MeminfoDatahandle(csvname)
	mystruct := &MeminfoResultstruct{TOTALMAP,Meminfolist}
	this.Data["json"] = mystruct
	fmt.Println(mystruct)
	fmt.Println(this.Data["json"])
	this.ServeJSON()
	return
}
type cpuResultstruct struct {
	//CpuTOTAL map[int]int
	//cpuAverage string
	CpuAverage float64
	CpuTOTAL map[int]int
}
//@router /Resultcpu [get]
func (this *ResultController) Resultcpu() {
	PackName := this.GetString("PackName")
	CaseName := this.GetString("Casename")
	fmt.Printf(PackName,CaseName)
	csvname:=CaseName+"_"+strtime+"_CPU"
	cpuAverage,CpuTOTAL:=models.CpuDatahandle(csvname)
	fmt.Println(CpuTOTAL)
	mystruct1 := &cpuResultstruct{cpuAverage,CpuTOTAL}
	this.Data["json"] = mystruct1
	this.ServeJSON()
	return
}

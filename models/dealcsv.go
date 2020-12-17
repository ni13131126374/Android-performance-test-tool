package models

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)
/*
@title    MeminfoRead
@description   将内存进行解析并实时放入csv
@auth      作者：田震 时间（2020/7/18   10:57 ）
@param     输入参数名：Meminfoname  参数类型:String "解释":在adb命令中获取的内存信息属性 "adb shell dumpsys meminfo 包名"
@return    返回参数名 TOTALdistributed,Meminfodict 参数类型:map,map    "解释"
*/
func MeminfoRead(Meminfoname string) (map[int]int,map[string]int){
	file, err := os.Open("./csvfolder/"+Meminfoname+".csv")
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	TOTALdistributed :=make(map[int]int)
	//var csvlist [][]string
	maxMeminfo:=0
	minMeminfo:=0
	TotalMeminfo:=0
	iMeminfo:=0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
		}
		if record[0]=="Java Heap"{
			continue
		}
		total, err := strconv.Atoi(record[7])
		if iMeminfo==0{
			minMeminfo=total
		}
		fmt.Println(total)
		num:=total/1024/100
		TotalMeminfo=TotalMeminfo+total
		iMeminfo++
		if maxMeminfo<total{
			maxMeminfo=total
		}
		if minMeminfo>total{
			minMeminfo=total
		}
		TOTALdistributed[num]=TOTALdistributed[num]+1
	}
	AveMeminfo:=TotalMeminfo/iMeminfo/1024
	fmt.Println(maxMeminfo/1024,minMeminfo/1024,AveMeminfo)

	Meminfodict:=map[string]int{"maxMeminfo":maxMeminfo/1024,"minMeminfo":minMeminfo/1024,"AveMeminfo":AveMeminfo}

	return TOTALdistributed,Meminfodict
}
func cpuRead(cpuname string) map[int]int{
	file, err := os.Open("./csvfolder/"+cpuname+".csv")
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	TOTALdistributed :=make(map[int]int)
	maxcpu:=0.0
	mincpu:=0.0
	Totalcpu:=0.0
	icpu:=0
	//var csvlist [][]string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
		}
		if record[0]=="CPU"{
			continue
		}
		cpu, err := strconv.ParseFloat(record[0],64)
		//i, _ := strconv.Atoi(fmt.Sprintf("%1.0f", cpu))
		i:=int(math.Floor(cpu + 0.5))
		TOTALdistributed[i]=TOTALdistributed[i]+1
		if icpu==0{
			mincpu=cpu
		}
		fmt.Println(cpu)

		Totalcpu=Totalcpu+cpu
		icpu++
		if maxcpu<cpu{
			maxcpu=cpu
		}
		if mincpu>cpu{
			mincpu=cpu
		}
	}
	//AveMeminfo :=Totalcpu/icpu
	fmt.Println(maxcpu,mincpu,)
	for i,j := range TOTALdistributed{
		fmt.Println("key=",i,"value=",j)
	}
	return TOTALdistributed
}
func IntervalDistribution(Interval int,num int){


}
func main() {
	MeminfoRead("测试1_2020_08_11_Meminfo")
	cpuRead("测试1_2020_08_25_CPU")
	//fmt.Println(csvlist,len(csvlist))
	//
	//print(csvlist)
}
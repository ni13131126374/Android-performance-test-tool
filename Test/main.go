package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)
//内存相关
func MeminfoDatahandle(Meminfoname string){
	var TOTAL []int
	var Graphics []int
	file, err := os.Open("./csvfolder/"+Meminfoname+".csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}
		//recordlist:=strings.Split(record,",")
		Totalnum,err :=strconv.Atoi(record[7])
		TOTAL=append(TOTAL, Totalnum)
		Graphicsnum,err :=strconv.Atoi(record[4])
		Graphics=append(Graphics,Graphicsnum)
		//fmt.Println(record) // record has the type []string
	}
	//fmt.Println(TOTAL)
	TOTALMAP:=Interval(TOTAL,102400)
	MeminfoLen:=len(TOTAL)
	TotalSort:=Sortnumber(TOTAL)
	TotalMinvalue:=TotalSort[1]
	TotalMaxvalue:=TotalSort[MeminfoLen-1]
	TotalMedian:=TotalSort[MeminfoLen/2+1]
	GraphicsSort:=Sortnumber(Graphics)
	GraphicsMaxvalue:=GraphicsSort[MeminfoLen-1]
	DiffMax:=Diffmethod(TOTAL)[MeminfoLen-2]
	TotalAverage:= Average(TOTAL)
	fmt.Println(TotalMinvalue,TotalMaxvalue,TotalMedian,
		GraphicsMaxvalue,TOTALMAP,DiffMax,TotalAverage)
}
//cpu
func CpuDatahandle(Cpuname string){
	var CpuTOTAL []int
	var cpuFloat float64
	var cpuFloatTOTAL []float64
	file, err := os.Open("./csvfolder/"+Cpuname+".csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}

		//recordlist:=strings.Split(record,",")
		//fmt.Println(record[0])
		//CpuTOTAL=append(CpuTOTAL, strconv.ParseFloat(record[0],64))
		//fmt.Println(record) // record has the type []string
		cpuFloat,err = strconv.ParseFloat(record[0],64)
		if err == io.EOF {
			break
		}
		CpuTOTAL=append(CpuTOTAL,int(cpuFloat))
		cpuFloatTOTAL=append(cpuFloatTOTAL,cpuFloat)

	}

	cpuAverage(cpuFloatTOTAL)
	Interval(CpuTOTAL,1)
	fmt.Println(CpuTOTAL)
}

//获取cpuaverage value
func cpuAverage(TOTALList []float64) float64 {
	Sum:=0.0
	for _,TOTAL :=range TOTALList{
		Sum=Sum+TOTAL
	}
	fmt.Println(Sum/float64(len(TOTALList))-2.00)
	return Sum/float64(len(TOTALList))-2.00
}
//获取average value
func Average(TOTALList []int) int {
	Sum:=0
	for _,TOTAL :=range TOTALList{
		Sum=Sum+TOTAL
	}
	fmt.Println(Sum/len(TOTALList)-2)
	return Sum/len(TOTALList)-2
}
//获取Difflist
func Diffmethod(TOTALList []int) []int {
	var Difflist []int
	Diffnum:=0
	for _,TOTAL :=range TOTALList{
		Difflist=append(Difflist,TOTAL-Diffnum)
		Diffnum=TOTAL
	}
	Difflist=Sortnumber(Difflist)
	return Difflist
}
//获取Total区间值
func Interval(TOTALList []int, Intervalvalue int) map[int]int {
	TOTALMAP:= make(map[int]int)
	for _,TOTAL :=range TOTALList{
		v:=TOTAL/Intervalvalue
		if _, ok := TOTALMAP[v]; ok {
			TOTALMAP[v]=TOTALMAP[v]+1
		}else {
			TOTALMAP[v]=1
		}
	}
	fmt.Println(TOTALMAP)
	return TOTALMAP
}


type IntSlice []int

func (s IntSlice) Len() int {
	return len(s)
}

func (s IntSlice) Swap(i, j int){
	s[i], s[j] = s[j], s[i]
}

func (s IntSlice) Less(i, j int) bool {
	return s[i] < s[j]
}
func Sortnumber(TOTALList []int) []int {
	sort.Ints(TOTALList)
	//sort.Sort(IntSlice(TOTALList))
	//fmt.Println("After sorted: ", TOTALList)
	return TOTALList
}
func main() {
	//MeminfoDatahandle("测试1_2020_08_11Meminfo_")
	CpuDatahandle("4.6.30滑动100次_2020_11_03_CPU")


}
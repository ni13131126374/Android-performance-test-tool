package models
import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)
//内存相关
func MeminfoDatahandle(Meminfoname string) (map[int]int, []int) {
	var TOTAL []int
	var Graphics []int
	var activetnum []int
	file, err := os.Open("./csvfolder/"+Meminfoname+".csv")
	if err != nil {
		fmt.Println("Error:", err)
		return nil, nil
	}
	defer file.Close()
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return nil, nil
		}
		//recordlist:=strings.Split(record,",")
		Totalnum,err :=strconv.Atoi(record[7])
		actives,err :=strconv.Atoi(record[8])
		activetnum=append(activetnum,actives)
		TOTAL=append(TOTAL, Totalnum)
		Graphicsnum,err :=strconv.Atoi(record[4])
		Graphics=append(Graphics,Graphicsnum)
		//fmt.Println(record) // record has the type []string
	}
	sort.Ints(activetnum)
	//fmt.Println(TOTAL)
	TOTALMAP:=Interval(TOTAL,102400)
	TOTALMAP[0]=TOTALMAP[0]-1
	MeminfoLen:=len(TOTAL)
	TotalSort:=Sortnumber(TOTAL)
	TotalMinvalue:=TotalSort[1]
	TotalMaxvalue:=TotalSort[MeminfoLen-1]
	TotalMedian:=TotalSort[MeminfoLen/2+1]
	GraphicsSort:=Sortnumber(Graphics)
	GraphicsMaxvalue:=GraphicsSort[MeminfoLen-1]
	DiffMax:=Diffmethod(TOTAL)[MeminfoLen-2]
	activetMax:=activetnum[MeminfoLen-1]
	TotalAverage:= Average(TOTAL)
	Meminfolist:=[] int {TotalMaxvalue,TotalMinvalue,TotalMedian,TotalAverage,DiffMax,GraphicsMaxvalue,activetMax}
	fmt.Println(TotalMinvalue,TotalMinvalue,TotalMedian,
		GraphicsMaxvalue,TOTALMAP,DiffMax,TotalAverage)
	return TOTALMAP,Meminfolist
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
	//fmt.Println(TOTALMAP)
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
	sort.Sort(IntSlice(TOTALList))
	//fmt.Println("After sorted: ", TOTALList)
	return TOTALList
}

package models

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

//cpu
func CpuDatahandle(Cpuname string) (float64, map[int]int) {
	var CpuTOTAL []int
	var cpuFloat float64
	var cpuFloatTOTAL []float64
	file, err := os.Open("./csvfolder/"+Cpuname+".csv")
	if err != nil {
		fmt.Println("Error:", err)
		return 0, nil
	}
	defer file.Close()
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return 0, nil
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

	//cpuAverage(cpuFloatTOTAL)


	return cpuAverage(cpuFloatTOTAL) ,Interval(CpuTOTAL,1)
}
//获取average value
func cpuAverage(TOTALList []float64) float64 {
	Sum:=0.0
	for _,TOTAL :=range TOTALList{
		Sum=Sum+TOTAL
	}
	fmt.Println(Sum/(float64(len(TOTALList))-2.00))
	return Sum/(float64(len(TOTALList))-2.00)
}

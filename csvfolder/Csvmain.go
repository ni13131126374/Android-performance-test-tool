package csvfolder
//package main


import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func CsvWrite(csvName string,title []string,data []string){
	file, er := os.Open(csvName)
	defer file.Close()
	// 如果文件不存在，创建文件
	if er != nil && os.IsNotExist(er) {
		file, er := os.Create(csvName)
		if er != nil {
			panic(er)
		}
		defer file.Close()

		// 写入字段标题
		w := csv.NewWriter(file) //创建一个新的写入文件流
		// 这里必须刷新，才能将数据写入文件。
		w.Write(title)
		w.Write(data)
		w.Flush()
		fmt.Printf("if end")
	} else {
		// 如果文件存在，直接加在末尾
		txt, err := os.OpenFile(csvName, os.O_APPEND|os.O_RDWR, 0666)
		defer txt.Close()
		if err != nil {
			panic(err)
		}
		w := csv.NewWriter(txt) //创建一个新的写入文件流
		w.Write(data)
		w.Flush()
		fmt.Printf("else end")
	}

}
func Getpath(){
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
}
func main()  {
	//dir,_ := os.Getwd()
	//fmt.Println("当前路径：",dir)
	strtime:=time.Now().Format("2006_01_02")
	title := []string{"Java Heap", "Native Heap", "Code", "Stack", "Graphics", "Private Other", "System",
		"TOTAL", "Activities", "Cpu"}
	data :=[]string{"Java Heap", "Native Heap", "Code", "Stack", "Graphics", "Private Other", "System", "TOTAL", "Activities", "Cpu"}
	csvpath:="csvfolder/"+strtime+"test.csv"
	CsvWrite(csvpath,title,data)
}

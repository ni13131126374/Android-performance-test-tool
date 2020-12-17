package models

import (
	"BeeTestProjec/csvfolder"
	"fmt"
	"io/ioutil"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var strtime=time.Now().Format("2006_01_02")

//通过adb 获取当前应用内存使用信息，并以数组形式输出
//adb shell dumpsys meminfo
func AdbShellDumpsysMeminfo(PackName string,CaseName string) map[string]int {
	MyCmd := exec.Command("adb", "shell", "dumpsys", "meminfo", PackName)
	s := AdbResultsofthe(MyCmd)
	meminfomap,data := Meminfomap(s)
	title := []string{"Java Heap", "Native Heap", "Code", "Stack", "Graphics", "Private Other", "System",
		"TOTAL", "Activities"}
	csvpath:="csvfolder/"+CaseName+"_"+strtime+"_Meminfo.csv"

	csvfolder.CsvWrite(csvpath,title,data)
	//var testMap map[string]int
	//fmt.Println(meminfomap)
	return meminfomap
}
//处理Meminfo 数据为map
func Meminfomap(meminfotxt string) (map[string]int,[]string){
	meminfomap := map[string]int{
		"Java Heap":     0,
		"Native Heap":   0,
		"Code":          0,
		"Stack":         0,
		"Graphics":      0,
		"Private Other": 0,
		"System":        0,
		"TOTAL":         0,
		"Activities":    0,
	}
	meminfolist:=make([]string,9)
	for k, _ := range meminfomap {
		a := strings.Replace("key:(.*\\d)", "key", k, -1)
		//fmt.Println(a)
		r := regexp.MustCompile(a)
		Heap := strings.Split(strings.Replace(r.FindString(meminfotxt), " ", "", -1), ":")
		//fmt.Println(JavaHeap[1])
		Heapnum := Heap
		if k == "TOTAL" {
			Heapnum = strings.Split(Heap[1], "T")
			Heapnum[1] = Heapnum[0]
			//fmt.Println(Heapnum)
		} else {
			Heapnum = Heap
		}
		num, _ := strconv.Atoi(Heapnum[1])
		meminfomap[k] = num
		if k=="Java Heap"{
			//meminfolist=append(meminfolist,strconv.Itoa(num))
			meminfolist[0]=strconv.Itoa(num)
		} else if k=="Native Heap"{
			meminfolist[1]=strconv.Itoa(num)
		}else if k=="Code"{
			meminfolist[2]=strconv.Itoa(num)
		}else if k=="Stack"{
			meminfolist[3]=strconv.Itoa(num)
		}else if k=="Graphics"{
			meminfolist[4]=strconv.Itoa(num)
		}else if k=="Private Other"{
			meminfolist[5]=strconv.Itoa(num)
		}else if k=="System"{
			meminfolist[6]=strconv.Itoa(num)
		}else if k=="TOTAL"{
			meminfolist[7]=strconv.Itoa(num)
		}else if k=="Activities"{
			meminfolist[8]=strconv.Itoa(num)
		}else {
			continue
		}
	}
	meminfomap["JavaHeap"] = meminfomap["Java Heap"]
	delete(meminfomap, "Java Heap")
	meminfomap["PrivateOther"] = meminfomap["Private Other"]
	delete(meminfomap, "Java Heap")
	meminfomap["NativeHeap"] = meminfomap["Native Heap"]
	delete(meminfomap, "Native Heap")
	for key, value := range meminfomap {
		if key != "Activities" {
			valuefloat := value / 1024
			meminfomap[key] = valuefloat
		}
	}
	//csvfolder.CsvWrite("csvfolder/meminfo.csv")
	return meminfomap,meminfolist
}

//获取Adb命令执行的结果
func AdbResultsofthe(MyCmd *exec.Cmd) string {
	MyOut, _ := MyCmd.StdoutPipe()
	MyCmd.Start()
	MyBytes, _ := ioutil.ReadAll(MyOut)
	MyCmd.Wait()
	MyOut.Close()
	s := string(MyBytes)
	return s
}

//通过adb 获取当前应用cpu使用信息，并以数组形式输出
//adb shell dumpsys top -n | grep PackName
func AdbShellTop(PackName string,CaseName string) float64 {
	PackNamelist := strings.Split(PackName, ".")
	sysType := runtime.GOOS
	MyCmd := exec.Command("adb", "shell", "top", "-n", "5", "|", "grep", PackNamelist[1])
	if sysType == "windows" {
		MyCmd = exec.Command("adb", "shell", "top", "-n", "5", "|", "findstr", PackNamelist[1])
	}

	s := AdbResultsofthe(MyCmd) //输出adb运行结果
	CpuFloat := 0.0
	casetest := DeleteExtraSpace(s)              //删除字符串中的多余空格，有多个空格时，仅保留一个空格
	casetestlist := strings.Split(casetest, " ") //切割字符串
	i := 0
	for _, value := range casetestlist {

		if value == "R" || value == "S" {
			i++
			continue
		}
		if i == 1 {
			cpunum, _ := strconv.ParseFloat(value, 64)
			cpucorenum := 8.0
			test := fmt.Sprintf("%.2f", cpunum/cpucorenum)
			CpuFloat, _ = strconv.ParseFloat(test, 64)

			//CpuFloat = cpunum
			break
		}
	}
	title := []string{"CPU"}
	csvpath:="csvfolder/"+CaseName+"_"+strtime+"_CPU.csv"
	//.Println(CpuFloat)
	Cpustring:=[]string {fmt.Sprintf("%.2f", CpuFloat)}
	csvfolder.CsvWrite(csvpath,title,Cpustring)
	return CpuFloat
}

//删除字符串中的多余空格，有多个空格时，仅保留一个空格
func DeleteExtraSpace(s string) string {
	//删除字符串中的多余空格，有多个空格时，仅保留一个空格
	s1 := strings.Replace(s, "  ", " ", -1)      //替换tab为空格
	regstr := "\\s{2,}"                          //两个及两个以上空格的正则表达式
	reg, _ := regexp.Compile(regstr)             //编译正则表达式
	s2 := make([]byte, len(s1))                  //定义字符数组切片
	copy(s2, s1)                                 //将字符串复制到切片
	spc_index := reg.FindStringIndex(string(s2)) //在字符串中搜索
	for len(spc_index) > 0 {                     //找到适配项
		s2 = append(s2[:spc_index[0]+1], s2[spc_index[1]:]...) //删除多余空格
		spc_index = reg.FindStringIndex(string(s2))            //继续在字符串中搜索
	}
	return string(s2)
}

//判断手机是否休眠
//"adb shell dumpsys power | grep state=OFF"
func AdbShellDumpsysPowerOff() bool {
	flag := false
	//test:="adb shell dumpsys power | grep state=OFF"
	//MyCmd := exec.Command(test)
	sysType := runtime.GOOS
	MyCmd := exec.Command("adb", "shell", "dumpsys", "power", "|", "grep", "state=OFF")
	if sysType == "windows" {
		MyCmd = exec.Command("adb", "shell", "dumpsys", "power", "|", "findstr", "state=OFF")
	}
	fmt.Println(MyCmd.Args)
	err := MyCmd.Run()
	if err == nil {
		flag = true
	}

	return flag
}

//通过adb 查看最上层activity名字：
//adb shell dumpsys window | grep "mCurrentFocus"
//代码中不能直接执行findstr/grep过滤,改正则匹配
func AdbShellDumpsysActivityF() string {
	//MyCmd := exec.Command("adb","shell","dumpsys", "window", "|","grep", "mCurrentFocus")
	MyCmd := exec.Command("adb", "shell", "dumpsys", "window")
	//fmt.Println(MyCmd.Args)
	s := AdbResultsofthe(MyCmd)
	//正则匹配mFocusedActivity
	match := regexp.MustCompile(` mCurrentFocus=(.*?)}`).FindString(s)
	match1 := regexp.MustCompile(`com(.*?)}`).FindString(match)
	matchlist := strings.Split(match1, "}")
	//fmt.Println(matchlist[0])
	return matchlist[0]
}

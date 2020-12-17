package main

import (
	"BeeTestProjec/models"
	"bytes"
	"crypto/md5"
	"fmt"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)




const (
	//可用下面的AdbShellDumpsysActivityF函数获取包名和activity名
	APPPackageName = "cn.XXX.android"
	APP            = "cn.XXX.android/com.XXX.XXXActivity"
)
func main() {
	models.AdbShellTop("com.fenzotech.jimu","test")
	models.AdbShellDumpsysMeminfo("com.fenzotech.jimu","test")
	//fmt.Println(AdbShellDumpsysPowerOff())
	//如果手机是休眠状态，则打开电源
	//if AdbShellDumpsysPowerOff() {
	//	AdbShellInputKeyEvent("26") //power
	//}
	//进入手机主屏
	//AdbShellInputKeyEvent("26")
	//AdbShellInputKeyEvent("4") //back
	//AdbShellInputKeyEvent("3") //home
	//AdbShellDumpsysActivityF()
	/*如果APP未启动，则启动APP
	  if !strings.Contains(AdbShellDumpsysActivityF(), APPPackageName) {
	      AdbShellAmStartN(APP)
	  }
	*/
	//Tap("设置", 0)
	//TimeSleepDuration(5)
	//TapOnce(`\d我的`, 0, 3, 573)
	//AdbShellInputKeyEvent("26") //power
	//PackName :="com.fenzotech.jimu"
	//AdbShellDumpsysMeminfo(PackName)
	//sysType := runtime.GOOS
	//if sysType == "linux" {
	//	fmt.Println(sysType)
	//	// LINUX系统
	//}
	//if sysType == "darwin" {
	//	fmt.Println(sysType)
	//	// windows系统
	//}
	//if sysType == "windows" {
	//	fmt.Println(sysType)
	//	// windows系统
	//}
	//AdbShellDumpsysMeminfo("com.fenzotech.jimu")
}

//通过adb 获取当前应用cpu使用信息，并以数组形式输出
//adb shell dumpsys top -n | grep PackName
func AdbShellTop(PackName string) float64{
	PackNamelist:=strings.Split(PackName,".")
	sysType := runtime.GOOS
	MyCmd := exec.Command("adb", "shell", "top", "-n", "5" ,"|","grep", PackNamelist[1])
	if sysType == "windows" {
		MyCmd = exec.Command("adb", "shell", "top", "-n", "5" ,"|","findstr", PackNamelist[1])
	}

	s:=AdbResultsofthe(MyCmd)//输出adb运行结果
	CpuFloat:=0.0
	casetest:= DeleteExtraSpace(s)//删除字符串中的多余空格，有多个空格时，仅保留一个空格
	casetestlist:=strings.Split(casetest," ")//切割字符串
	i:=0
	for _,value :=range casetestlist{

		if (value=="R" || value=="S") {
			i++
			continue
			}
		if i==1{
			cpunum,_:= strconv.ParseFloat(value,64)
			cpucorenum:=8.0
			test:=fmt.Sprintf("%.2f", cpunum/cpucorenum)
			CpuFloat,_ = strconv.ParseFloat(test, 64)
			CpuFloat=cpunum

			break
			}
	}
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

//通过adb 获取当前应用内存使用信息，并以数组形式输出
//adb shell dumpsys meminfo
func AdbShellDumpsysMeminfo(PackName string) map[string]int{
	MyCmd := exec.Command("adb","shell","dumpsys", "meminfo", PackName)
	s:=AdbResultsofthe(MyCmd)
	meminfomap:=Meminfomap(s)
	print(meminfomap)
	//var testMap map[string]int
	return meminfomap
}
//处理Meminfo 数据为map
func Meminfomap(meminfotxt string) map[string]int  {
	meminfomap := map[string]int{
		"Java Heap":0,
		"Native Heap":0,
		"Code":0,
		"Stack":0,
		"Graphics":0,
		"Private Other":0,
		"System":0,
		"TOTAL":0,
		"Activities":0,
	}
	for k,_ := range meminfomap {
		a:= strings.Replace("key:(.*\\d)","key",k, -1)
		//fmt.Println(a)
		r := regexp.MustCompile(a)
		Heap := strings.Split(strings.Replace(r.FindString(meminfotxt)," ", "",-1),":")
		Heapnum:=Heap
		if k=="TOTAL"{

			Heapnum=strings.Split(Heap[1],"T")
			Heapnum[1]=Heapnum[0]

		}else {
			Heapnum=Heap
		}
		num, _ := strconv.Atoi(Heapnum[1])
		meminfomap[k]=num

	}
	return meminfomap
}
//获取Adb命令执行的结果
func AdbResultsofthe(MyCmd *exec.Cmd) string{
	MyOut, _ := MyCmd.StdoutPipe()
	MyCmd.Start()
	MyBytes, _ := ioutil.ReadAll(MyOut)
	MyCmd.Wait()
	MyOut.Close()
	s := string(MyBytes)
	return s
}
//模拟按键，如按下home键，键值参考；https://blog.csdn.net/shililang/article/details/14449527
//adb shell input keyevent 3
func AdbShellInputKeyEvent(s string) {
	exec.Command("adb", "shell", "input", "keyevent", s).Run()
}

//模拟屏幕点击
//有的控件死活抓不到，只能直接点击
//adb shell input tap  900 800
func AdbShellInputTap(x, y int) {
	x2 := strconv.Itoa(x)
	y2 := strconv.Itoa(y)
	exec.Command("adb", "shell", "input", "tap", x2, y2).Run()
}

//模拟滑动
//adb shell input swipe  0 0  600 600
func AdbShellInputSwipe(x1, y1, x2, y2 int) {
	xx1 := strconv.Itoa(x1)
	yy1 := strconv.Itoa(y1)
	xx2 := strconv.Itoa(x2)
	yy2 := strconv.Itoa(y2)
	exec.Command("adb", "shell", "input", "swipe", xx1, yy1, xx2, yy2).Run()
}

//模拟长按 最后一个参数1000表示1秒，可将下面某个参数由500改为501，即允许坐标点有很小的变化。
//adb shell input swipe  500 500  500 500 1000
func AdbShellInputSwipeL(x1, y1, x2, y2, t int) {
	xx1 := strconv.Itoa(x1)
	yy1 := strconv.Itoa(y1)
	xx2 := strconv.Itoa(x2)
	yy2 := strconv.Itoa(y2)
	exec.Command("adb", "shell", "swipe", "tap", xx1, yy1, xx2, yy2).Run()
}

//模拟输入“字符”
//adb shell input text "abc"
//若需输入中文，可参考：https://blog.csdn.net/slimboy123/article/details/54140029
func AdbShellInputText(s string) {
	exec.Command("adb", "shell", "input", "text", s).Run()
}

//等待几秒
func TimeSleepDuration(x int) {
	time.Sleep(time.Duration(x) * time.Second)
}

//截屏并保存到当前目录下。
//由于需在手机和电脑上复制文件，必要时可增加延时或用下面的PathExists()判断文件是否存在，如：
//time.Sleep(time.Duration(2) * time.Second)
func AdbShellScreencapPullRm() {
	exec.Command("adb", "shell", "screencap", "-p", "/sdcard/screen.png").Run()
	exec.Command("adb", "pull", "/sdcard/screen.png", ".").Run()
	exec.Command("adb", "shell", "rm", "/sdcard/screen.png").Run()
}

//根据图像中某一片矩形区域的左上点和右下点，计算该部分图像点的MD5，以便比较图像
//后来发现不必用这种原始的办法，可以用下面的AdbShellUiautomatorDump()下载手机页面可视控件的XML文件进行解析
func ReadPngPart2MD5(x1, y1, x2, y2 int) string {
	//先截图
	AdbShellScreencapPullRm()
	file, _ := os.Open("screen.png")
	defer file.Close()
	im, _ := png.Decode(file)
	//x := im.Bounds().Max.X
	//y := im.Bounds().Max.Y
	//按行扫描
	mybuff := new(bytes.Buffer)
	for j := y1; j <= y2; j++ {
		for i := x1; i <= x2; i++ {
			r, g, b, a := im.At(i, j).RGBA()
			mybuff.Write([]byte(fmt.Sprintf("%d ", r)))
			mybuff.Write([]byte(fmt.Sprintf("%d ", g)))
			mybuff.Write([]byte(fmt.Sprintf("%d ", b)))
			mybuff.Write([]byte(fmt.Sprintf("%d ", a)))
		}
	}
	ss := fmt.Sprint(md5.Sum(mybuff.Bytes()))
	//fmt.Printf("MobileMainPage[%d,%d][%d,%d]sum:\n%s", x1, y1, x2, y2, ss)
	return ss
}

//判断设备是否休眠。重要补充：注意：这里有错误，需要将exec.Command中的命令用逗号分隔，不能直接findstr，应在代码中查找
//adb shell dumpsys power | findstr "Display Power:state="
func AdbShellDumpsysPowerOff() bool {
	flag := false
	//test:="adb shell dumpsys power | grep state=OFF"
	//MyCmd := exec.Command(test)
	sysType := runtime.GOOS
	MyCmd := exec.Command("adb","shell","dumpsys", "power", "|", "grep", "state=OFF")
	if sysType == "windows" {
		MyCmd = exec.Command("adb","shell","dumpsys", "power", "|", "findstr", "state=OFF")
	}
	fmt.Println(MyCmd.Args)
	err :=MyCmd.Run()
	if err==nil{
		flag = true
	}

	return flag
}

//查看手机上应用的packageName
//adb shell pm list packages
func AdbShellPmListPackages() string {
	MyCmd := exec.Command("adb", "shell", "pm", "list", "packages")
	MyOut, _ := MyCmd.StdoutPipe()
	MyCmd.Start()
	MyBytes, _ := ioutil.ReadAll(MyOut)
	MyCmd.Wait()
	MyOut.Close()
	s := string(MyBytes)
	return s
}

//通过adb 查看最上层activity名字：
//adb shell dumpsys window | grep "mCurrentFocus"
//代码中不能直接执行findstr过滤,改正则匹配
func AdbShellDumpsysActivityF() string {
	//MyCmd := exec.Command("adb","shell","dumpsys", "window", "|","grep", "mCurrentFocus")
	MyCmd := exec.Command("adb","shell","dumpsys", "window")
	//fmt.Println(MyCmd.Args)
	s:=AdbResultsofthe(MyCmd)
	//正则匹配mFocusedActivity
	match := regexp.MustCompile(` mCurrentFocus=(.*?)}`).FindString(s)
	match1 := regexp.MustCompile(`com(.*?)}`).FindString(match)
	matchlist:=strings.Split(match1,"}")
	fmt.Println(matchlist[0])
	return matchlist[0]
}

//启动activity,如计算器com.android.calculator2/com.android.calculator2.Calculator
//adb shell am start -n 包名/包名＋类名（-n 类名,-a action,-d date,-m MIME-TYPE,-c category,-e 扩展数据,等
//如：adb shell am start -n com.android.camera/.Camera
func AdbShellAmStartN(p string) {
	exec.Command("adb", "shell", "am", "start", "-n", p).Run()
}

//获取当前应用屏幕上所有控件的信息并保存在sdcard下window_dump.xml文件里面. sdk版本16以上
//如：adb shell uiautomator dump --compressed /sdcard/window_dump.xml
//adb pull /sdcard/window_dump.xml .
//adb shell rm /sdcard/window_dump.xml
//可参考：https://blog.csdn.net/henni_719/article/details/72953251
//由于需在手机和电脑上复制文件，必要时可增加延时或用下面的PathExists()判断文件是否存在，如：
//time.Sleep(time.Duration(2) * time.Second) 但是经实测无需延时等待。
//特别提醒注意：对于可scroll的页面，只能dump出显示在屏幕上的可见的部分。即滑动页面后需重新dump。这个问题曾困扰我一天。
func AdbShellUiautomatorDump() {
	//删除当前目录下的window_dump.xml
	exec.Command("cmd", "/c", "del", "-y", "window_dump.xml").Run()
	//重新dump
	exec.Command("adb", "shell", "uiautomator", "dump", "/sdcard/window_dump.xml").Run()
	exec.Command("adb", "pull", "/sdcard/window_dump.xml", ".").Run()
	exec.Command("adb", "shell", "rm", "/sdcard/window_dump.xml").Run()
}

//用正则找xml文件中bounds的坐标点
//感觉用xml解析不如用正则查找直观，这里需要你自己写正则表达式，返回bounds的两个坐标点[x1,y1][x2,y2]
//如：x1, y1, x2, y2 :=RegXmlPoint(`<node\s+index=\"\d+\"\s+text=\"我的\".+?\[(\d+),(\d+)\]\[(\d+),(\d+)\]`)
func RegXmlPoint(s string) (x1, y1, x2, y2 int) {
	r := regexp.MustCompile(s)
	file, _ := os.Open("window_dump.xml")
	defer file.Close()
	doc, _ := ioutil.ReadAll(file)
	doc1 := string(doc)
	match := r.FindStringSubmatch(doc1)
	x1, _ = strconv.Atoi(match[1])
	y1, _ = strconv.Atoi(match[2])
	x2, _ = strconv.Atoi(match[3])
	y2, _ = strconv.Atoi(match[4])
	return x1, y1, x2, y2
}

//用法如：Tap（`设置`,0）  将打开手机设置
//用正则根据`关键词`（反引号，可包含正则）匹配xml文件中node区域，其中有bounds的坐标点,计算bounds中心点，并Tap之
//第一个参数为匹配用的关键词，第二个参数ix表示点击匹配到的第几个，0表示第一个，－1表示最后一个
//正则参考：ss := fmt.Sprintf("%s%s%s", `<node.[^>]+?`, s, `.[^>]+?\[(\d+),(\d+)\]\[(\d+),(\d+)\].+?[^>]`)
//        golang正则匹配任意汉字可用reg = regexp.MustCompile(`[\p{Han}]+`)  这里写正则费了较大功夫。
func Tap(s string, ix int) {
	//先执行AdbShellUiautomatorDump函数。
	AdbShellUiautomatorDump()
	file, _ := os.Open("window_dump.xml")
	defer file.Close()
	doc, _ := ioutil.ReadAll(file)
	doc1 := string(doc)
	ss := fmt.Sprintf("%s%s%s", `<node.[^>]+?`, s, `.[^>]+?\[(\d+),(\d+)\]\[(\d+),(\d+)\].+?>`)
	r := regexp.MustCompile(ss)
	match := r.FindAllStringSubmatch(doc1, -1)
	le := len(match)
	//匹配到1个或多个，ixx表示匹配到的第几个
	ixx := ix
	if le == 0 {
		log.Println("未匹配到:", s)
		return
	}
	if ix < 0 {
		ixx = le + ix
	}
	if ixx < 0 {
		ixx = 0
	}

	x1, _ := strconv.Atoi(fmt.Sprint(match[ixx][1]))
	y1, _ := strconv.Atoi(fmt.Sprint(match[ixx][2]))
	x2, _ := strconv.Atoi(fmt.Sprint(match[ixx][3]))
	y2, _ := strconv.Atoi(fmt.Sprint(match[ixx][4]))

	xx := (x2-x1)/2 + x1
	yy := (y2-y1)/2 + y1
	log.Println(s)
	AdbShellInputTap(xx, yy)
}

//用法如：TapOnce（`我的`,0,10,105） 可改为递归调用自身
///意思是：点击含有`我的`关键词（反引号，可包含正则）的第一个node（0表示第1个）；会打开新页面，10秒后返回后，再
//向上滑动页面，使该node的y2位置向上滚动到105px（页面上可滚动部分最上端的y1值，也就是上面不可滚动部分的y2值），使该node不可见。不能再点击。
//注意：此代码不通用，主要是向上滚动时从开始点[500，y2]滚动到结束点[500，pos]，这里的开始和结束点要根据实际选择。
func TapOnce(s string, ix, tm, pos int) {
	//先执行AdbShellUiautomatorDump函数。
	AdbShellUiautomatorDump()
	file, _ := os.Open("window_dump.xml")
	defer file.Close()
	doc, _ := ioutil.ReadAll(file)
	doc1 := string(doc)
	ss := fmt.Sprintf("%s%s%s", `<node.[^>]+?`, s, `.[^>]+?\[(\d+),(\d+)\]\[(\d+),(\d+)\].+?>`)
	r := regexp.MustCompile(ss)
	match := r.FindAllStringSubmatch(doc1, -1)
	le := len(match)
	//匹配到1个或多个，ixx表示匹配到的第几个
	ixx := ix
	if le == 0 {
		log.Println("未匹配到:", s)
		return
	}
	if ix < 0 {
		ixx = le + ix
	}
	if ixx < 0 {
		ixx = 0
	}

	x1, _ := strconv.Atoi(fmt.Sprint(match[ixx][1]))
	y1, _ := strconv.Atoi(fmt.Sprint(match[ixx][2]))
	x2, _ := strconv.Atoi(fmt.Sprint(match[ixx][3]))
	y2, _ := strconv.Atoi(fmt.Sprint(match[ixx][4]))

	xx := (x2-x1)/2 + x1
	yy := (y2-y1)/2 + y1
	log.Println(s)
	AdbShellInputTap(xx, yy)
	//此时app打开了新的内容页
	TimeSleepDuration(tm)
	AdbShellInputKeyEvent("4") //back
	TimeSleepDuration(1)
	//向上滚动
	AdbShellInputSwipe(500, y2, 500, pos)

}

//判断文件或文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//未实现
func notimp(s ...string) {
	/*
	   MyCmd := exec.Command("adb", "devices")
	   MyOut, _ := MyCmd.StdoutPipe()
	   MyCmd.Start()
	   MyBytes, _ := ioutil.ReadAll(MyOut)
	   MyCmd.Wait()
	   MyOut.Close()
	   fmt.Println(s)
	   return string(MyBytes)
	*/
}

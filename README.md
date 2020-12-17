# Android-performance-test-tool
beego +adb 编写的mac/linux下android 可视化性能测试工具
### 部署使用方法
####1、安装adb，链接手机。
####2、对beego进行打包 
    进入项目根目录执行 bee pack
    执行完毕后，会发现根目录下有一个tar.gz压缩包
####3、解压beego压缩包，并在其目录下打开终端窗口
####4、在终端运行./BeeTestProjec
    提示http server Running on http://8080 
####5、打开浏览，在浏览器打开http://localhost:8080/ 
    输入测试app包名，与测试case名 点击开始测试，就可以进行实时查看手机的内存与cpu数据波动图了
####6、测试结束后，在压缩包所在文件夹下找到csvfolder文件夹 就可以找到以 标示_时间 所命名的测试结果csv文件

                                         


package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

//水桶服的打包工具下载链接
const spigotDownUrl = "https://hub.spigotmc.org/jenkins/job/BuildTools/lastSuccessfulBuild/artifact/target/BuildTools.jar"

//当前路径
var currPath string

//保存路径
var savePath string

//BuildTools.jar保存路径
var buildToolsSavePath string

//spigot.jar保存路径
var spigotSavePath string

//craftbukkit.jar保存路径
var craftbukkitSavePath string

//目录符
var EOL string

//编译mc的版本
var mcVersion = flag.String("v", "1.11.2", "设定需要编译的水桶服版本")

//编译需要的内存
var memSize = flag.String("m", "1024", "编译时所需要的内存")

func main() {
	flag.Parse()

	//检查java是否安装
	checkJava()
	//检查系统设置环境变量
	checkOS()
	//下载水桶服
	downSpigot()
	//编译水桶服
	configure()
}

func checkOS() {
	switch runtime.GOOS {
	case "windows":
		EOL = "\\"
	case "darwin":
		EOL = "/"
	case "linux":
		EOL = "/"
	default:
		EOL = "/"
	}

	savePath = os.TempDir() + EOL
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	currPath = strings.Replace(dir, "\\", "/", -1)
	buildToolsSavePath = savePath + "BuildTools.jar"
	spigotSavePath = savePath + "spigot-" + *mcVersion + ".jar"
	craftbukkitSavePath = savePath + "craftbukkit-" + *mcVersion + ".jar"
}

func checkJava() {
	fmt.Println("检查是否安装了java程序")
	cmd := exec.Command("java", "-version")
	_, err := cmd.Output()
	if err != nil {
		log.Fatal("请安装java再进行操作")
	}
}

func downSpigot() {
	if _, err := os.Stat(buildToolsSavePath); err != nil {
		res, _ := http.Get(spigotDownUrl)
		file, _ := os.Create(buildToolsSavePath)
		io.Copy(file, res.Body)
		fmt.Println("下载BuildTools.jar完成")
	} else {
		fmt.Println("存在BuildTools.jar")
	}
}

func configure() {
	fmt.Println("开始编译水桶服...")
	fmt.Println("注: 可能因为网络原因时间会比较长, 建议使用速度快的科学上网工具")
	cmd := exec.Command("java", "-Xmx"+*memSize+"M", "-jar", buildToolsSavePath, "--rev", *mcVersion)
	if err := cmd.Run(); err != nil {
		log.Fatal("编译出错, err=", err.Error())
	}
	fmt.Println("编译完成!")

	//保存spigot的jar包
	//if err := os.Rename(spigotSavePath, currPath); err != nil {
	//    log.Fatal("保存文件失败,  err=", err.Error())
	//}
	//保存craftbukkit的jar包
	//craftbukkitSavePath := savePath + EOL + "craftbukkit-" + mcVersion + ".jar"
	//if err := os.Rename(craftbukkitSavePath, currPath); err != nil {
	//    log.Fatal("保存文件失败,  err=", err.Error())
	//}
	fmt.Printf("文件保存在当前文件夹下【spigot-%s.jar】和【craftbukkit-%s.jar】\n", *mcVersion, *mcVersion)
}

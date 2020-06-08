package main

import (
    "bufio"
    "bytes"
    "flag"
    "github.com/gogs/chardet"
    "golang.org/x/text/encoding/japanese"
    "golang.org/x/text/encoding/korean"
    "golang.org/x/text/encoding/simplifiedchinese"
    "golang.org/x/text/encoding/traditionalchinese"
    "golang.org/x/text/encoding/unicode"
    "golang.org/x/text/transform"
    "io"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "os/exec"
    "path"
    "path/filepath"
    //"runtime"
    "strconv"
    //"syscall"
    "time"
)

const (
    SpigotDownloadUrl = "https://hub.spigotmc.org/jenkins/job/BuildTools/lastSuccessfulBuild/artifact/target/BuildTools.jar"
    Version           = "1.11.2"
    MemSize           = 1024
)

var (
    version string
    memSize int
)

func main() {
    flag.StringVar(&version, "v", Version, "设定需要编译的水桶服版本")
    flag.IntVar(&memSize, "m", MemSize, "编译时所需要的内存(单位: M)")
    flag.Parse()

    //检查java是否安装
    cmd := exec.Command("java", "-version")
    _, err := cmd.Output()
    if err != nil {
        log.Fatal("请安装java再进行操作")
    }

    tempDir := os.TempDir()
    rootDir := filepath.Dir(os.Args[0])
    os.Chdir(tempDir)

    //编译水桶服
    configure(tempDir, rootDir)
}

func build(name string, args []string) error {
    cmd := exec.Command(name, args...)
    stdout, err := cmd.StdoutPipe()
    if err != nil {
        return err
    }

    cmd.Start()
    for {
        r := bufio.NewReader(stdout)
        line, _, err := r.ReadLine()
        if err != nil {
            break
        }
        log.Printf("%s\n", encodeText(line))
    }

    if err = cmd.Wait(); err != nil {
        return err
    }

    return nil
}

//判断gbk字符集
func isGBK(data []byte) bool {
    length := len(data)
    var i int = 0
    for i < length {
        if data[i] <= 0x7f {
            //编码0~127,只有一个字节的编码，兼容ASCII码
            i++
            continue
        } else {
            //大于127的使用双字节编码，落在gbk编码范围内的字符
            if data[i] >= 0x81 &&
                data[i] <= 0xfe &&
                data[i+1] >= 0x40 &&
                data[i+1] <= 0xfe &&
                data[i+1] != 0xf7 {
                i += 2
                continue
            } else {
                return false
            }
        }
    }
    return true
}

//根据字符集转换
func encodeText(data []byte) string {
    detector := chardet.NewTextDetector()
    result, _ := detector.DetectBest(data)

    buf := make([]byte, 0)
    switch result.Charset {
    case "GB18030":
        r := transform.NewReader(bytes.NewReader(data), simplifiedchinese.GB18030.NewDecoder())
        buf, _ = ioutil.ReadAll(r)
    case "Big5":
        r := transform.NewReader(bytes.NewReader(data), traditionalchinese.Big5.NewDecoder())
        buf, _ = ioutil.ReadAll(r)
    case "EUC-JP":
        r := transform.NewReader(bytes.NewReader(data), japanese.EUCJP.NewDecoder())
        buf, _ = ioutil.ReadAll(r)
    case "EUC-KR":
        r := transform.NewReader(bytes.NewReader(data), korean.EUCKR.NewDecoder())
        buf, _ = ioutil.ReadAll(r)
    default:
        if isGBK(data) {
            r := transform.NewReader(bytes.NewReader(data), simplifiedchinese.GBK.NewDecoder())
            buf, _ = ioutil.ReadAll(r)
        } else {
            r := transform.NewReader(bytes.NewReader(data), unicode.UTF8.NewDecoder())
            buf, _ = ioutil.ReadAll(r)
        }
    }

    return string(buf)
}

//跨磁盘移动文件
func rename(oldpath, newpath string) error {
    //if runtime.GOOS != "windows" {
    //    from, err := syscall.UTF16PtrFromString(oldpath)
    //    if err != nil {
    //        return err
    //    }
    //    to, err := syscall.UTF16PtrFromString(newpath)
    //    if err != nil {
    //        return err
    //    }
    //
    //    return syscall.MoveFile(from, to)
    //}
    return os.Rename(oldpath, newpath)
}

func configure(tempDir, rootDir string) {
    buildToolsSavePath := path.Join(tempDir, "BuildTools.jar")
    if _, err := os.Stat(buildToolsSavePath); err != nil {
        res, _ := http.Get(SpigotDownloadUrl)
        file, _ := os.Create(buildToolsSavePath)
        io.Copy(file, res.Body)
        log.Println("下载BuildTools.jar完成")
    } else {
        log.Println("存在BuildTools.jar")
    }

    dirname := "spigot-build" + strconv.FormatInt(time.Now().Unix(), 10)
    tempDir = path.Join(tempDir, dirname)
    os.Mkdir(dirname, os.ModePerm)
    os.Chdir(tempDir)

    log.Println("开始编译水桶服...")
    log.Printf("注: 第一次编译当前版本(%s)可能时间会比较长, 建议使用速度快的科学上网工具\n", version)
    buildArgs := []string{
        "-Xmx" + strconv.Itoa(memSize) + "M",
        "-jar",
        buildToolsSavePath,
        "--rev",
        version,
    }
    if err := build("java", buildArgs); err != nil {
        log.Fatal("编译出错, err=", err.Error())
    } else {
        log.Println("编译完成!")
    }

    //保存spigot的jar包
    spigotFilename := "spigot-" + version + ".jar"
    spigotTempPath := path.Join(tempDir, spigotFilename)
    spigotSavePath := path.Join(rootDir, spigotFilename)
    if err := rename(spigotTempPath, spigotSavePath); err != nil {
        log.Fatal("保存文件失败,  err=", err.Error())
    }
    //保存craftbukkit的jar包
    craftbukkitFilename := "craftbukkit-" + version + ".jar"
    craftbukkitTempPath := path.Join(tempDir, craftbukkitFilename)
    craftbukkitSavePath := path.Join(rootDir, craftbukkitFilename)
    if err := rename(craftbukkitTempPath, craftbukkitSavePath); err != nil {
        log.Fatal("保存文件失败,  err=", err.Error())
    }

    log.Printf("文件保存在当前文件夹下【spigot-%s.jar】和【craftbukkit-%s.jar】\n", version, version)
}

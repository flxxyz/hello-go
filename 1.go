// 类似包名
package main

// 引入扩展
import "fmt"

// 定义常量
const BOOL = true

// 全局变量
var name string = "ffff"

// 一般类型声明？？？ 没明白
type x int

// 结构体声明，还行，了解过一点点
type y struct {

}

// 声明接口
type Iz interface {

}

func defined() {
	name = "ccc"
	fmt.Println(BOOL)
}

func main() {
	defined()
	fmt.Println("111", name)
}

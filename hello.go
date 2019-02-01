package main

import (
    "fmt"
    "log"
    "reflect"
    "strconv"
)
func main() {
    fmt.Println("fuck world")
    id := float64(123)
    debug(id)

    str := strconv.FormatFloat(id, 'f', -1, 64)
    debug(str)

    i, _ := strconv.Atoi(str)
    debug(i)

}

func debug(v interface{})  {
    log.Printf("val = %+v, type = %v \n", v, reflect.TypeOf(v))
}
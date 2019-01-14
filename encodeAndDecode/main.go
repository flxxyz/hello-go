package main

import (
	"fmt"
	"strconv"
	"strings"
)

const NAME ="abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func toHex(xxx []string)  {
	//var ooo []string
	for _, s := range xxx{
		new, _ := strconv.Atoi(s)
		//new, _ := strconv.FormatInt(new, 16)
		fmt.Printf("%v", new)
		fmt.Println()
		//ooo = append(ooo, strconv.FormatInt(new, 16))
	}

	//fmt.Println(ooo)
}

func encode(str string)  {
	s := strings.Split(str, "")
	toHex(s)
}


func main()  {
	encode("1234")
}

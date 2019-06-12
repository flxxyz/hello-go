package main

import "fmt"

type Response struct {
	Index   int
	Content []byte
	Time    int64
}

type ResponseList struct {
	Data []*Response
}

func (rl *ResponseList) append(r *Response) {
	rl.Data = append(rl.Data, r)
}

func (rl *ResponseList) len() int {
	return len(rl.Data)
}

func (rl *ResponseList) echo() {
	for _, r := range rl.Data {
		fmt.Println("第", r.Index, "次")
		fmt.Println("Content=", string(r.Content))
		fmt.Println("Time=", r.Time)
		fmt.Println("-------------------------------")
		fmt.Println()
	}

	fmt.Println("完成结果=", rl.len())
}

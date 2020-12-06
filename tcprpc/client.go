package main

import (
	"fmt"
	"net/rpc"
)

type AddRequest struct {
	Left int
	Right int
}

//
type AddResponse struct {
	Result int
}

func main(){
	client, err := rpc.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	//调用远程方法的步骤
	req := AddRequest{2,3}
	resp := AddRequest{}
	// func, req, resp
	err = client.Call("Calc.Add", req, &resp)
	fmt.Println(err, resp)

	client.Close()
}

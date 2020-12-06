package main

import (
	"fmt"
	"net"
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

type Calc struct {}

// 参数1：请求对象（可以是指针/值）
// 参数2：响应对象（可以是指针）
// 返回值：error
func (c *Calc) Add(req AddRequest, resp *AddResponse) error {
	fmt.Println("calc.add")
	resp.Result = req.Left + req.Right
	return nil
}

type AAA struct {
}

func (a *AAA) Add(req AddRequest, resp *AddResponse) error {
	fmt.Println("aaa.Add")
	return nil
}

func main() {
	// 本地调用
	/*req := AddRequest{1,2}
	resp := AddResponse{}
	calc := Calc{}
	calc.Add(req, &resp)
	fmt.Println(resp.result)
	*/
	// rpc调用步骤
	// 注册
	rpc.Register(&Calc{})
	rpc.Register(&AAA{})
	// 启动服务
	listener, _ := net.Listen("tcp", "0.0.0.0:8888")
	/*for {
		conn, _ := listener.Accept()
	}
	 */
	rpc.Accept(listener)
	listener.Close()
}

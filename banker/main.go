package main

import (
	"fmt"
	"myos/banker/service"
	"strconv"
	"time"
)

func main() {
	//已有资源
	service.Available = []int{1, 6, 2, 2}
	//已分配
	service.Allocation = [][]int{{0, 0, 3, 2},
		{1, 0, 0, 1},
		{1, 3, 5, 4},
		{0, 3, 3, 2},
		{0, 0, 1, 4}}
	//需要
	service.Need = [][]int{{0, 0, 1, 2},
		{1, 7, 5, 0},
		{2, 3, 5, 6},
		{0, 6, 5, 2},
		{0, 6, 5, 6}}

	service.Request = []int{0, 0, 1, 0}
	service.Pid = 0

	fmt.Println("现有资源")
	for _, v := range service.Available {
		fmt.Print(v, " ")
	}
	fmt.Println()
	fmt.Println("已分配资源")
	for _, v := range service.Allocation {
		for _, n := range v {
			fmt.Print(n, " ")
		}
		fmt.Println()
	}
	fmt.Println("需分配资源")
	for _, v := range service.Need {
		for _, n := range v {
			fmt.Print(n, " ")
		}
		fmt.Println()
	}
	fmt.Println("请求资源")
	for _, v := range service.Request {
		fmt.Print(v, " ")
	}
	fmt.Println()
	fmt.Println("请求资源id：" + strconv.Itoa(service.Pid))
	//执行银行家算法
	service.Bank()
	time.Sleep(1000000000000000000)
}

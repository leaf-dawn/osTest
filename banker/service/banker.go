package service

import "fmt"

//可以用资源向量
var Available []int

//已分配矩阵
var Allocation [][]int

//需求矩阵
var Need [][]int

//请求向量
var Request []int

//finish向量，用于标记是否工作向量是否工作完成
var finish []bool

//请求资源进程号
var Pid int

//用于存储安全序列
var serial []int

//银行家算法
func Bank() {
	for i := 0; i < len(Request); i++ {
		//判断是否超过公布请求数
		if Request[i] > Need[Pid][i] {
			fmt.Println("所需资源超过宣布最大值")
			return
		}
	}
	for i := 0; i < len(Request); i++ {
		if Request[i] > Available[i] {
			fmt.Println("尚无足够资源，需等待")
			return
		}
	}

	//给请求向量分配资源
	for i := 0; i < len(Available); i++ {
		Available[i] = Available[i] - Request[i]
	}
	for i := 0; i < len(Allocation[0]); i++ {
		Allocation[Pid][i] = Allocation[Pid][i] + Request[i]
	}
	for i := 0; i < len(Need[0]); i++ {
		Need[Pid][i] = Need[Pid][i] - Request[i]
	}

	//银行家算法
	if Safe() {
		fmt.Println("完成分配")
		fmt.Println("安全序列：", serial)
	} else {
		fmt.Println("系统处于不安全状态，等待")
	}

}

//银行家
func Safe() bool {
	//初始化结束标记
	for i := 0; i < len(Allocation); i++ {
		finish = append(finish, false)
	}

	work := Available
	//用来记录已完成的作业数量
	num := 0
	for num != len(finish) {
		pronum := num
		//遍历所有判断是否可以分配
		for i := 0; i < len(Need); i++ {

			//如果没有分配，且可以进行分配，则分配，并添加到序列里
			if !finish[i] && canAvailable(work, Need[i]) {
				distribute(work, Allocation[i])
				finish[i] = true
				num++
				serial = append(serial, i)
			}

		}

		//如果进行一轮分配以后，num没有改变则表示没有合法序列
		if num == pronum {
			return false
		}
	}

	return true
}

//判断是否可以分配
func canAvailable(work []int, need []int) bool {
	//遍历每一个资源，判断是否可以分配
	for j := 0; j < len(work); j++ {
		//如果可以分配
		if need[j] > work[j] {
			return false
		}
	}
	return true
}

//分配
func distribute(work []int, allocation []int) {
	for i := 0; i < len(work); i++ {
		work[i] += allocation[i]
	}
}

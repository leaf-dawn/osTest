package main

import (
	"fmt"
	"myos/job_scheduling/service"
	"strconv"
)

/** 主函数 */
func main() {
	num := 0
	fmt.Println("作业调度实验")
	fmt.Println("（1）先到先服务 （2）短作业优先")
	_, _ = fmt.Scan(&num)
	scan()
	if num == 1 {
		service.Fcfs()
	} else if num == 2 {
		service.Sjf()
	}
	//结束做最后的输出
	service.FinishPrint()
}

func scan() {
	fmt.Print("请输入作业数:")
	num := 0
	_, _ = fmt.Scanln(&num)
	for i := 0; i < num; i++ {
		j := &service.Jcb{}
		fmt.Print("请输入作业No." + strconv.Itoa(i) + "名称:")
		_, _ = fmt.Scanln(&j.Name)
		fmt.Print("请输入作业到达时间:")
		_, _ = fmt.Scanln(&j.ArriveTime)
		fmt.Print("请输入作业运行时间:")
		_, _ = fmt.Scanln(&j.RunTime)
		//添加到队列
		service.JcbList = append(service.JcbList, j)
	}
}

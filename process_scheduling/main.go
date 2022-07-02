package main

import (
	"fmt"
	"myos/process_scheduling/service"
	"sort"
	"strconv"
)

/** 主函数 */
func main() {
	fmt.Println("进程调度算法实验")
	fmt.Println("（1）时间片轮转 （2）高相应比优先")
	num := 0
	_, _ = fmt.Scan(&num)
	if num == 1 {
		rrScan()
		service.Rr2()
	} else if num == 2 {
		hrrnScan()
		service.Hrrn()
	}
	//结束做最后的输出
	service.FinishPrint()
}

/** 获取输入的pcb列表，并按照到达时间排序 */
func rrScan() {
	//获取输入
	fmt.Print("请输入进程数:")
	num := 0
	_, _ = fmt.Scanln(&num)
	fmt.Print("请输入时间片:")
	_, _ = fmt.Scanln(&service.LimitTime)
	for i := 0; i < num; i++ {
		p := &service.Pcb{}
		fmt.Print("请输入进程No." + strconv.Itoa(i) + "名称:")
		_, _ = fmt.Scanln(&p.Name)
		fmt.Print("请输入进程到达时间:")
		_, _ = fmt.Scanln(&p.ArriveTime)
		fmt.Print("请输入进程运行所需时间:")
		_, _ = fmt.Scanln(&p.RunTime)
		//添加到队列
		service.PcbList = append(service.PcbList, p)
	}
	sort.Sort(service.SortPcbByArriveTime(service.PcbList))
}

/** 高相应比优先算法获取输入 */
func hrrnScan() {
	//获取输入
	fmt.Print("请输入进程数:")
	num := 0
	_, _ = fmt.Scanln(&num)
	for i := 0; i < num; i++ {
		p := &service.Pcb{}
		fmt.Print("请输入进程No." + strconv.Itoa(i) + "名称:")
		_, _ = fmt.Scanln(&p.Name)
		fmt.Print("请输入进程到达时间:")
		_, _ = fmt.Scanln(&p.ArriveTime)
		fmt.Print("请输入进程运行所需时间:")
		_, _ = fmt.Scanln(&p.RunTime)
		//添加到队列
		service.PcbList = append(service.PcbList, p)
	}
	sort.Sort(service.SortPcbByArriveTime(service.PcbList))
}

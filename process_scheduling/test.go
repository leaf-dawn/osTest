package main

import (
	"fmt"
	"sort"
)

/**
 * 进程控制块
 */
type Pcb struct {
	Name  string
	State int
	/** 到达时间 */
	ArriveTime int
	/** CPU使用时间 */
	RunTime int
	/** 还需要运行的时间 */
	ServiceTime int
	/** 结束时间 */
	FinishTime int
}

/* 用于存储输入的pcb */
var pcbList []*Pcb

/* 就绪队列 */
var readyQueue []*Pcb

/* 用于记录结束的pcb */
var finishQueue []*Pcb

/** 用来记录当前时间 */
var currentTime int = 0

/** 时间片 */
var limitTime int

//时间片轮转法
func Rr() {
	//所有pcb按到达时间到达
	sort.Sort(SortPcbByArriveTime(pcbList))

	for len(pcbList) > 0 || len(readyQueue) > 0 {
		if len(readyQueue) != 0 {
			//就绪队列中取出第一个进程
			nowPcb := readyQueue[0]
			readyQueue = readyQueue[1:]
			//使用时间片
			if nowPcb.ServiceTime-nowPcb.RunTime > limitTime {
				//如果需要使用时间小于时间片,减小该进程运行时间
				nowPcb.RunTime = nowPcb.RunTime + limitTime
				//更新当前时间
				currentTime = currentTime + limitTime
				//判断队列里面是否有进程需要进入就绪队列
				toReadyList()
				//把当前为执行完的进程插入到队列尾部
				readyQueue = append(readyQueue, nowPcb)
			} else {
				//如果在未使用完时间片，进程已经结束
				currentTime = currentTime + nowPcb.ServiceTime - nowPcb.RunTime
				nowPcb.RunTime = nowPcb.ServiceTime
				//更新结束时间
				nowPcb.FinishTime = currentTime

				finishQueue = append(finishQueue, nowPcb)
				//判断队列里面是否有进程需要进入就绪队列
				toReadyList()
			}
		} else {
			//如果就绪队列为空，更新当前时间，并添加进程到就绪队列
			currentTime = pcbList[0].ArriveTime
			toReadyList()
		}
	}
}

func FinishPrint() {

	for i := 0; i < len(finishQueue); i++ {
		// 周转时间
		p := finishQueue[i]
		turnaroundTime := p.FinishTime - p.ArriveTime
		weightWaitTime := float64(turnaroundTime) / float64(p.ServiceTime)
		t := fmt.Sprintf("进程名称:%s 到达时间:%d 结束时间:%d 周转时间:%d 带权周转时间:%f\n",
			p.Name, p.ArriveTime, p.FinishTime, turnaroundTime, weightWaitTime)
		fmt.Println(t)
	}
}

/**
 * 把到达的pcb加入到就绪队列中
 */
func toReadyList() {
	for len(pcbList) > 0 {
		if pcbList[0].ArriveTime <= currentTime {
			readyQueue = append(readyQueue, pcbList[0])
			pcbList = pcbList[1:]
		} else {
			break
		}
	}
}

/**
 * 用于对pcb列表进行排序
 * 根据开始时间
 */
type SortPcbByArriveTime []*Pcb

func (s SortPcbByArriveTime) Len() int {
	return len(s)
}

func (s SortPcbByArriveTime) Less(a int, b int) bool {
	return s[a].ArriveTime < s[b].ArriveTime
}

func (s SortPcbByArriveTime) Swap(a int, b int) {
	tmp := s[a]
	s[a] = s[b]
	s[b] = tmp
}

func main() {

	//时间片
	limitTime = 1
	//进程的队列
	pcbList = []*Pcb{
		{Name: "A", ArriveTime: 0, ServiceTime: 4},
		{Name: "B", ArriveTime: 1, ServiceTime: 3},
		{Name: "C", ArriveTime: 2, ServiceTime: 4},
		{Name: "D", ArriveTime: 3, ServiceTime: 2},
		{Name: "E", ArriveTime: 4, ServiceTime: 4},
	}

	//时间片轮转
	Rr()
	FinishPrint()
}

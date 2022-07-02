package service

import (
	"fmt"
	"sort"
)

type Jcb struct {
	//作业名称
	Name string
	//提交时间
	ArriveTime int
	//作业状态
	State int
	//运行时间s
	RunTime int
	//完成时间
	FinishTime int
	//所需资源
	RequiredResources int
}

//用于存储输入jcb,在作业调度里面由于需要做多种算法，所以初始化以后，不要改变这里的值
var JcbList []*Jcb

/* 就绪队列 */
var readyQueue []*Jcb

/* 用于记录结束的pcb */
var finishQueue []*Jcb

/** 用来记录当前时间 */
var currentTime int = 0

//先到先服务
func Fcfs() {
	sort.Sort(sortJcbByArriveTime(JcbList))
	//所有pcb按到达时间到达
	for len(JcbList) > 0 || len(readyQueue) > 0 {
		if len(readyQueue) != 0 {
			//就绪队列中取出第一个进程
			nowJcb := readyQueue[0]
			readyQueue = readyQueue[1:]
			//如果该进程已经执行完毕
			currentTime = currentTime + nowJcb.RunTime
			//更新结束时间
			nowJcb.FinishTime = currentTime
			finishQueue = append(finishQueue, nowJcb)
			//判断队列里面是否有进程需要进入就绪队列
			toReadyList()
		} else {
			//如果就绪队列为空，更新当前时间，并添加进程到就绪队列
			currentTime = JcbList[0].ArriveTime
			toReadyList()
		}
	}
}

//短作业优先
func Sjf() {
	sort.Sort(sortJcbByArriveTime(JcbList))
	//所有pcb按到达时间到达
	for len(JcbList) > 0 || len(readyQueue) > 0 {
		if len(readyQueue) != 0 {
			//就绪队列中取出第一个进程
			nowJcb := readyQueue[0]
			readyQueue = readyQueue[1:]
			//如果该进程已经执行完毕
			currentTime = currentTime + nowJcb.RunTime
			//更新结束时间
			nowJcb.FinishTime = currentTime
			finishQueue = append(finishQueue, nowJcb)
			//判断队列里面是否有进程需要进入就绪队列
			toReadyList()
			sort.Sort(sortJcbByRunTime(readyQueue))
		} else {
			//如果就绪队列为空，更新当前时间，并添加进程到就绪队列
			currentTime = JcbList[0].ArriveTime
			toReadyList()
			sort.Sort(sortJcbByRunTime(readyQueue))
		}
	}
}

func FinishPrint() {

	for i := 0; i < len(finishQueue); i++ {
		// 周转时间
		p := finishQueue[i]
		turnaroundTime := p.FinishTime - p.ArriveTime
		weightWaitTime := float64(turnaroundTime) / float64(p.RunTime)
		t := fmt.Sprintf("作业名称:%s 到达时间:%d 结束时间:%d 周转时间:%d 带权周转时间:%f\n",
			p.Name, p.ArriveTime, p.FinishTime, turnaroundTime, weightWaitTime)
		fmt.Println(t)
	}
}

/**
 * 把到达的pcb加入到就绪队列中
 */
func toReadyList() {
	for len(JcbList) > 0 {
		if JcbList[0].ArriveTime <= currentTime {
			readyQueue = append(readyQueue, JcbList[0])
			JcbList = JcbList[1:]
		} else {
			break
		}
	}
}

/**
 * 用于对jcb列表进行排序
 * 根据开始时间
 */

//根据到达时间排序
type sortJcbByArriveTime []*Jcb

func (s sortJcbByArriveTime) Len() int {
	return len(s)
}

func (s sortJcbByArriveTime) Less(a int, b int) bool {
	return s[a].ArriveTime < s[b].ArriveTime
}

func (s sortJcbByArriveTime) Swap(a int, b int) {
	tmp := s[a]
	s[a] = s[b]
	s[b] = tmp
}

//根据作业大小排序
type sortJcbByRunTime []*Jcb

func (s sortJcbByRunTime) Len() int {
	return len(s)
}

func (s sortJcbByRunTime) Less(a int, b int) bool {
	return s[a].RunTime < s[b].RunTime
}

func (s sortJcbByRunTime) Swap(a int, b int) {
	tmp := s[a]
	s[a] = s[b]
	s[b] = tmp
}

package service

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
	/** 运行时间 */
	RunTime int
	/** 服务时间 */
	ServiceTime int
	/** 结束时间 */
	FinishTime int
}

/* 用于存储输入的pcb */
var PcbList []*Pcb

/* 就绪队列 */
var readyQueue []*Pcb

/* 用于记录结束的pcb */
var finishQueue []*Pcb

/** 用来记录当前时间 */
var currentTime int = 0

/** 时间片 */
var LimitTime int

//高响应比优先
func Hrrn() {
	for _, p := range PcbList {
		p.ServiceTime = p.RunTime
	}
	//所有pcb按到达时间到达
	sort.Sort(SortPcbByArriveTime(PcbList))
	//所有pcb按到达时间到达
	for len(PcbList) > 0 || len(readyQueue) > 0 {
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
			sort.Sort(sortPcbByCorrespondenceRatio(readyQueue))
		} else {
			//如果就绪队列为空，更新当前时间，并添加进程到就绪队列
			currentTime = PcbList[0].ArriveTime
			toReadyList()
			sort.Sort(sortPcbByCorrespondenceRatio(readyQueue))
		}
	}
}

//时间片轮转法
func Rr() {
	for _, p := range PcbList {
		p.ServiceTime = p.RunTime
	}
	//所有pcb按到达时间到达
	sort.Sort(SortPcbByArriveTime(PcbList))

	for len(PcbList) > 0 || len(readyQueue) > 0 {
		if len(readyQueue) != 0 {
			//就绪队列中取出第一个进程
			nowPcb := readyQueue[0]
			readyQueue = readyQueue[1:]
			//使用时间片
			if nowPcb.RunTime > LimitTime {
				//如果需要使用时间小于时间片,减小该进程运行时间
				nowPcb.RunTime = nowPcb.RunTime - LimitTime
				//更新当前时间
				currentTime = currentTime + LimitTime
				//判断队列里面是否有进程需要进入就绪队列
				toReadyList()
				//把当前为执行完的进程插入到队列尾部
				readyQueue = append(readyQueue, nowPcb)
			} else {
				//如果该进程已经执行完毕
				currentTime = currentTime + getMin(LimitTime, nowPcb.RunTime)
				//更新结束时间
				nowPcb.FinishTime = currentTime
				finishQueue = append(finishQueue, nowPcb)
				//判断队列里面是否有进程需要进入就绪队列
				toReadyList()
			}
		} else {
			//如果就绪队列为空，更新当前时间，并添加进程到就绪队列
			currentTime = PcbList[0].ArriveTime
			toReadyList()
		}
	}
}

//时间片轮转法所有进程一次性加入就绪队列，不考虑进行先后到达
func Rr2() {
	for _, p := range PcbList {
		p.ServiceTime = p.RunTime
		p.RunTime = 0
	}
	//所有进程按到达时间到达
	readyQueue = PcbList
	sort.Sort(SortPcbByArriveTime(readyQueue))
	//当前时间块进到第一个到达的进程的时间
	currentTime = readyQueue[0].ArriveTime
	for len(readyQueue) > 0 {
		//就绪队列中取出第一个进程
		nowProcess := readyQueue[0]
		readyQueue = readyQueue[1:]
		//使用时间片
		if nowProcess.ServiceTime-nowProcess.RunTime > LimitTime {
			//如果需要使用时间小于时间片,减小该进程运行时间
			nowProcess.RunTime = nowProcess.RunTime + LimitTime
			//更新当前时间
			currentTime = currentTime + LimitTime
			//把当前为执行完的进程插入到队列尾部
			readyQueue = append(readyQueue, nowProcess)
		} else {
			//如果进程在没有使用完时间片以前已经结束
			//更新当前时间
			currentTime = currentTime + nowProcess.ServiceTime - nowProcess.RunTime
			nowProcess.RunTime = nowProcess.ServiceTime
			//更新结束时间
			nowProcess.FinishTime = currentTime
			finishQueue = append(finishQueue, nowProcess)
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
	for len(PcbList) > 0 {
		if PcbList[0].ArriveTime <= currentTime {
			readyQueue = append(readyQueue, PcbList[0])
			PcbList = PcbList[1:]
		} else {
			break
		}
	}
}

/**
 * 将结束的进程输出
 */
func getMin(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
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

type sortPcbByCorrespondenceRatio []*Pcb

func (s sortPcbByCorrespondenceRatio) Len() int {
	return len(s)
}

func (s sortPcbByCorrespondenceRatio) Less(a int, b int) bool {
	ac := 1 + float64(currentTime-s[a].ArriveTime)/float64(s[a].RunTime)
	bc := 1 + float64(currentTime-s[b].ArriveTime)/float64(s[b].RunTime)
	return ac > bc
}

func (s sortPcbByCorrespondenceRatio) Swap(a int, b int) {
	tmp := s[a]
	s[a] = s[b]
	s[b] = tmp
}

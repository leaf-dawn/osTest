package main

import "myos/job_scheduling/service"

func test1() {
	//进程的队列
	service.JcbList = []*service.Jcb{
		&service.Jcb{
			Name:       "1",
			ArriveTime: 0,
			RunTime:    1,
		},
		&service.Jcb{
			Name:       "2",
			ArriveTime: 1,
			RunTime:    100,
		},
		&service.Jcb{
			Name:       "3",
			ArriveTime: 2,
			RunTime:    100,
		},
		&service.Jcb{
			Name:       "4",
			ArriveTime: 3,
			RunTime:    1,
		},
	}
}

func test2() {
	//进程的队列
	service.JcbList = []*service.Jcb{
		&service.Jcb{
			Name:       "1",
			ArriveTime: 0,
			RunTime:    10,
		},
		&service.Jcb{
			Name:       "2",
			ArriveTime: 1,
			RunTime:    1,
		},
		&service.Jcb{
			Name:       "3",
			ArriveTime: 2,
			RunTime:    2,
		},
		&service.Jcb{
			Name:       "4",
			ArriveTime: 3,
			RunTime:    1,
		},
		&service.Jcb{
			Name:       "5",
			ArriveTime: 4,
			RunTime:    5,
		},
	}
}

func main() {
	num1 := 1
	if num1 == 1 {
		test1()
	} else if num1 == 2 {
		test2()
	}

	num := 2
	if num == 1 {
		//先到先服务
		service.Fcfs()
	} else if num == 2 {
		//短作业优先
		service.Sjf()
	}
	service.FinishPrint()
}

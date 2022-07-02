package service

import (
	"fmt"
	"strconv"
)

//空闲地址空间
type Area struct {
	//存储作业名称
	Name    string
	Address int
	Size    int
	//判断是否空闲
	Mark bool
}

//空间列表
var AreaList = []*Area{}

//用于记录循环首次适配的位置
var nowAddress int = 0

/**
 * 循环首次适配
 */
func Nf(name string, request int) int {
	index := getIndex(nowAddress)
	index++
	if index > len(AreaList) {
		index = 0
	}
	flagIndex := index

	//循环查找
	for {
		//分配
		allocation := malloc2(name, request, nowAddress)
		//判断是否分配成功
		if allocation == -1 {
			//获取下一个地址继续分配
			nowAddress = getNextAddress(nowAddress)
			if nowAddress == -1 {
				return -1
			} else {
				//分配成功，如果下一个地址的index和flagIndex重复
				if getIndex(nowAddress) == flagIndex {
					break
				}
			}
		} else {
			//分配成功，更新index
			nowAddress = getNextAddress(nowAddress)
			return allocation
		}
	}
	//分配失败
	return -1
}

/**
* 分配地址空间，首次适配
* 返回的是分配的地址位置
* -1为失败
 */
func Ff(name string, request int) int {
	address := 0
	//遍历所有空间
	for i := 0; i < len(AreaList); i++ {
		//判断空闲且未被使用就分配
		if AreaList[i].Mark && AreaList[i].Size >= request {
			malloc(name, request, i)
			address = AreaList[i].Address
			return address
		}
	}
	return -1
}

/**
 * 申请空间
* @param:name 作业名称
* @param:request 作业大小
* index : 索引
*/
func malloc(name string, request int, index int) {
	if AreaList[index].Size == request {
		AreaList[index].Name = name
		AreaList[index].Mark = false
	} else if AreaList[index].Size > request {
		//分裂出来
		area1 := &Area{Address: AreaList[index].Address, Name: name, Size: request, Mark: false}
		area2 := &Area{Address: AreaList[index].Address + request, Size: AreaList[index].Size - request, Mark: true}
		//替换
		left := AreaList[0:index]
		right := make([]*Area, len(AreaList[index+1:]))
		copy(right, AreaList[index+1:])
		AreaList = append(append(left, area1, area2), right...)
	}
}

/**
 * 申请空间
* @param:name 作业名称
* @param:request 作业大小
* nowIndex: 地址在数组的位置
* address : 起始地址
* 返回分配地址，否则为-1
*/
func malloc2(name string, request int, nowAddress int) int {
	//遍历找到index
	index := getIndex(nowAddress)
	//判断当前位置是否可以分配
	if nowAddress+request <= AreaList[index].Address+AreaList[index].Size {
		//指针在起始位置
		if nowAddress == AreaList[index].Address {
			malloc(name, request, index)
			return nowAddress
		}
		//不在起始位置，切分为2块
		if nowAddress != AreaList[index].Address && nowAddress+request == AreaList[index].Address+AreaList[index].Size {
			area1 := &Area{Address: AreaList[index].Address, Size: AreaList[index].Size - request, Mark: true}
			area2 := &Area{Address: nowAddress, Size: request, Mark: false}
			left := AreaList[0:index]
			right := make([]*Area, len(AreaList[index+1:]))
			copy(right, AreaList[index+1:])
			AreaList = append(append(left, area1, area2), right...)
			return area2.Address
		}
		//分为3块
		if nowAddress != AreaList[index].Address && nowAddress+request < AreaList[index].Address+AreaList[index].Size {
			area1 := &Area{Address: AreaList[index].Address, Size: request - AreaList[index].Address, Mark: true}
			area2 := &Area{Address: nowAddress, Size: request, Mark: false}
			area3 := &Area{Address: area2.Address + area2.Size, Size: AreaList[index].Size - area1.Size - area2.Size, Mark: true}
			left := AreaList[0:index]
			right := make([]*Area, len(AreaList[index+1:]))
			copy(right, AreaList[index+1:])
			AreaList = append(append(left, area1, area2, area3), right...)
			return area2.Address
		}

	}
	return -1
}

//通过地址寻找index
func getIndex(address int) int {
	index := 0
	for index < len(AreaList) {
		if AreaList[index].Address <= nowAddress && AreaList[index].Size+AreaList[index].Address > nowAddress {
			break
		}
		index++
	}
	return index
}

/** 释放空间 */
func Free(name string) {
	left := 0
	//遍历所有空间
	for i := 0; i < len(AreaList); i++ {
		//获取当前位置
		if name == AreaList[i].Name {
			//找到空间位置，标记为空闲
			AreaList[i].Mark = true
			//向左边合并
			for left = i; left > 0; left-- {
				if !AreaList[left-1].Mark {
					break
				}
			}
		}
	}
	//合并
	beginAddress := AreaList[left].Address
	size := 0
	var i int
	for i = left; i < len(AreaList) && AreaList[i].Mark; i++ {
		size += AreaList[i].Size
	}
	newErea := &Area{Address: beginAddress, Size: size, Mark: true}
	//删除旧存储
	AreaList = append(AreaList[0:left], AreaList[i:]...)
	//添加新存储
	AreaList = append(AreaList[0:left], append([]*Area{newErea}, AreaList[left:]...)...)
}

//输出内存结构
func Print() {
	fmt.Println("=========================")
	for _, a := range AreaList {
		if !a.Mark {
			fmt.Println("{作业" + a.Name + ":" + strconv.Itoa(a.Size) + "," + "1 }")
		} else {
			fmt.Println("{" + strconv.Itoa(a.Size) + ",0 }")
		}
	}
}
func GetStorageSituation() (int, int) {
	idle := 0
	used := 0
	for _, a := range AreaList {
		if a.Mark {
			idle += a.Size
		} else {
			used += a.Size
		}
	}
	return idle, used
}

//循环获取下一个空闲空间
func getNextAddress(nowAddress int) int {
	//下一个index位置
	index := getIndex(nowAddress)
	index++
	if index == len(AreaList) {
		index = 0
	}
	flagIndex := index
	//用于判断是否进入的循环
	for {
		if AreaList[index].Mark {
			return AreaList[index].Address
		}
		//判断是否是下一个
		index++
		if index > len(AreaList) {
			index = 0
		}
		if index == flagIndex {
			break
		}
	}
	return -1
}

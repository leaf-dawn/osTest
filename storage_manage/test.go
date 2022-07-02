package main

import (
	"myos/storage_manage/service"
	"strconv"
	"strings"
)

//这里是批量输入案例
func main() {
	//设置存储器大小
	size := 640
	service.AreaList = []*service.Area{{Address: 0, Size: size, Mark: true}}
	//模式,1为首次适配，2为循环首次
	model := 1
	//这里是输入样例子
	list := [][]string{{"1", "1", "130"},
		{"2", "1", "60"},
		{"3", "1", "100"},
		{"2", "0"},
		{"4", "1", "200"},
		{"3", "0"},
		{"1", "0"},
		{"5", "1", "140"},
		{"6", "1", "60"},
		{"7", "1", "50"},
		{"8", "1", "60"},
	}
	//获取操作
	//遍历所有
	for len(list) != 0 {
		work := list[0]
		list = list[1:]
		if strings.TrimSpace(work[1]) == "0" {
			service.Free(strings.TrimSpace(work[0]))
			service.Print()
		} else if strings.TrimSpace(work[1]) == "1" {
			//如果是需要申请空间
			if model == 1 {
				num, _ := strconv.Atoi(work[2])
				_ = service.Ff(strings.TrimSpace(work[0]), num)
				service.Print()
			} else if model == 2 {
				num, _ := strconv.Atoi(work[2])
				_ = service.Nf(strings.TrimSpace(work[0]), num)
				service.Print()
			}
		}
	}
}

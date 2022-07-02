package main

import (
	"fmt"
	"myos/storage_manage/service"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("存储器管理实验")
	fmt.Print("请输入存储大小:")
	size := 0
	_, _ = fmt.Scan(&size)
	//初始化存储
	service.AreaList = []*service.Area{{Address: 0, Size: size, Mark: true}}
	fmt.Print("（1）首次适配 （2）循环首次适配:")
	num := 0
	_, _ = fmt.Scan(&num)
	//获取操作
	//遍历所有
	for {
		//获取名称和操作
		fmt.Println("(0,1,50 :左边代表名称，右边代表操作),-1退出")
		si := ""
		_, _ = fmt.Scan(&si)
		if si == "-1" {
			break
		} else {
			nameAndOption := strings.Split(si, ",")
			if strings.TrimSpace(nameAndOption[1]) == "0" {
				service.Free(strings.TrimSpace(nameAndOption[0]))
				fmt.Println(strings.TrimSpace(nameAndOption[0]) + "释放成功")
			} else if strings.TrimSpace(nameAndOption[1]) == "1" {
				//如果是需要申请空间
				if num == 1 {
					num, _ := strconv.Atoi(nameAndOption[2])
					add := service.Ff(strings.TrimSpace(nameAndOption[0]), num)
					fmt.Println("作业：" + strings.TrimSpace(nameAndOption[0]) + "分配的地址为" + strconv.Itoa(add))
				} else if num == 2 {
					num, _ := strconv.Atoi(nameAndOption[2])
					add := service.Nf(strings.TrimSpace(nameAndOption[0]), num)
					fmt.Println("作业：" + strings.TrimSpace(nameAndOption[0]) + "分配的地址为" + strconv.Itoa(add))
				}
			}
		}
	}
}

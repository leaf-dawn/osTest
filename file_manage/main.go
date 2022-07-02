package main

import (
	"fmt"
	"myos/file_manage/service"
)

func main() {
	//初始化用户列表
	service.MFD = []*service.File{
		{Name: "fzw",
			NextFile: []*service.File{},
		},
		{Name: "admin",
			NextFile: []*service.File{},
		},
		{Name: "user",
			NextFile: []*service.File{},
		},
	}
	for {
		//用户名称
		fmt.Println("输入用户名")
		name := ""
		_, _ = fmt.Scanln(&name)
		//初始化用户列表
		service.UserName = name
		//进入当前用户
		for i := 0; i < len(service.MFD); i++ {
			if service.MFD[i].Name == name {
				service.NowFile = service.MFD[i]
				service.UFD = service.MFD[i].NextFile
				service.Path = name + service.Path
				break
			}
		}
		if service.NowFile == nil {
			fmt.Println("用户名错误")
		} else {
			break
		}
	}
	//创建文件 4 2 1
	for {
		s := ""
		fmt.Print(service.Path + ":  ")
		fmt.Scan(&s)
		if s == "dir" {
			service.List()
		}
		if s == "cd" {
			s2 := ""
			fmt.Scan(&s2)
			service.Cd(s2)
		}
		if s == "open" {
			s2 := ""
			fmt.Scan(&s2)
			service.Open(s2)
		}
		if s == "cd.." {
			service.Cd_()
		}
		if s == "create" {
			s2 := ""
			fmt.Scan(&s2)
			service.Create(s2, "", 7)
		}
		if s == "write" {
			s2 := ""
			fmt.Scan(&s2)
			service.Write(s2)
		}
		if s == "read" {
			i := 0
			fmt.Scan(&i)
			service.Read(i)
		}
		if s == "exit" {
			break
		}
		if s == "mkdir" {
			s1 := ""
			fmt.Scan(&s1)
			service.Mkdir(s1)
		}
	}

}

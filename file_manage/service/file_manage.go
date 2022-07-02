package service

import (
	"fmt"
)

type File struct {
	Name           string
	ProtectionCode int
	Content        string
	Len            int
	//是否是目录
	IsDir bool
	//下一级文件
	NextFile   []*File
	BeforeFile *File
}

type User struct {
	Username string
	Password string
}

//运行文件对象
type afd struct {
	File       *File
	ReadIndex  int
	WriteIndex int
}

//当前用户名称
var UserName string
var AFD *afd

//所有用户列表
var MFD = []*File{}

//file
var NowFile *File

//路径
var Path = ":/"
var UFD []*File

//创建文件
func Create(fileName string, content string, protectionCode int) {
	//获取当前用户
	if NowFile == nil {
		NowFile = getNowMFDUnit()
	}

	//添加文件
	file := File{
		Name:           fileName,
		Content:        content,
		ProtectionCode: protectionCode,
		Len:            len(content),
		BeforeFile:     NowFile,
		IsDir:          false,
	}
	NowFile.NextFile = append(NowFile.NextFile, &file)
	fmt.Println("添加成功")
}

func Mkdir(fileName string) {
	if NowFile == nil {
		NowFile = getNowMFDUnit()
	}

	//添加文件
	file := File{
		Name:       fileName,
		BeforeFile: NowFile,
		IsDir:      true,
	}
	NowFile.NextFile = append(NowFile.NextFile, &file)
}

//删除文件
func Delete(fileName string) {
	//获取当前用户
	if NowFile == nil {
		NowFile = getNowMFDUnit()
	}
	//获取文件并删除
	for i, f := range NowFile.NextFile {
		if f.Name == fileName {
			//进行删除
			NowFile.NextFile = append(NowFile.NextFile[0:i], NowFile.NextFile[i+1:]...)
			fmt.Println("删除成功")
			return
		}
	}
	fmt.Println("删除失败")
}

//进入文件
func Cd(fileName string) {
	//判断是否是文件
	for _, r := range NowFile.NextFile {
		if r.Name == fileName {
			if r.IsDir {
				Path = Path + "/" + r.Name
				NowFile = r
			} else {
				fmt.Println("不是文件夹")
			}
		}
	}
}

//回到上一级
func Cd_() {
	NowFile = NowFile.BeforeFile
}

//获取文件列表
func List() {
	for _, f := range NowFile.NextFile {
		if f.IsDir {
			fmt.Println("*" + f.Name)
		} else {
			fmt.Println(f.Name)
		}
	}
}

//打开文件
func Open(fileName string) {
	//获取当前用户
	if NowFile == nil {
		NowFile = getNowMFDUnit()
	}
	//获取当前文件
	f := getNowFile(fileName)
	AFD = &afd{}
	//添加afd
	AFD.File = f
	AFD.ReadIndex = 0
	AFD.WriteIndex = f.Len
}

//关闭文件
func Close() {
	//关闭文件
	AFD = nil
	fmt.Println("关闭成功")
}

//读取文件
func Read(len int) {
	//len表示要读取的长度,read指针会向后移动
	if AFD == nil {
		fmt.Println("当前未打开任何文件")
		return
	}
	//判断是否可以读取
	if AFD.File.ProtectionCode|1 == 0 {
		fmt.Println("当前文件不可读")
		return
	}
	//进行读取
	if AFD.ReadIndex+len > AFD.File.Len {
		fmt.Println("文件内容：" + string(AFD.File.Content[AFD.ReadIndex:]))
		AFD.ReadIndex = AFD.File.Len - 1
	} else {
		fmt.Println("文件内容：" + string(AFD.File.Content[AFD.ReadIndex:AFD.ReadIndex+len]))
		AFD.ReadIndex = AFD.ReadIndex + len
	}
}

//写操作
func Write(content string) {
	if AFD == nil {
		fmt.Println("当前未打开任何文件")
		return
	}
	//判断是否可以写
	if AFD.File.ProtectionCode|2 == 0 {
		fmt.Println("当前文件不可写")
		return
	}
	//添加内容
	before := []rune(AFD.File.Content[0:AFD.WriteIndex])
	after := []rune(AFD.File.Content[AFD.WriteIndex:])
	AFD.File.Content = string(append(append(before, []rune(content)...), after...))
	//移动写指针
	AFD.WriteIndex = AFD.WriteIndex + len(content)
	fmt.Println("已写入")
}

//获取当前用户目录
func getNowMFDUnit() *File {
	//获取当前用户
	var m *File
	for i := 0; i < len(MFD); i++ {
		if UserName == MFD[i].Name {
			m = MFD[i]
		}
	}
	if m == nil {
		fmt.Println("无当前用户")
	}
	return m
}

func getNowFile(fileName string) *File {
	for _, f := range NowFile.NextFile {
		if f.Name == fileName {
			return f
		}
	}
	fmt.Println("获取不含该文件")
	return nil
}

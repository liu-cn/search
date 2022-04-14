package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var example = `例如：
	想要检索出当前所有目录下所有go文件中包含request.GetName()的具体位置，并且排除带注释的代码（go代码里注释为//）需要执行：
		search -file *.go -text request.GetName() -exc //
    即可`
var SearchFile = flag.String("file", "", "要搜索的文件或者指定搜索的文件后缀，如果指定文件名，则只搜索该文件下文本，如果是 *.go 这种表示只搜索以.go后缀文件结尾的文件，")
var SearchF = flag.String("f", "", "要搜索的文件或者指定搜索的文件后缀，如果指定文件名，则只搜索该文件下文本，如果是 *.go 这种表示只搜索以.go后缀文件结尾的文件，")
var SearchPath = flag.String("path", "", "准备搜索文件的起始路径，如果不指定则默认从执行程序的目录开始搜索")
var SearchP = flag.String("p", "", "准备搜索文件的起始路径，如果不指定则默认从执行程序的目录开始搜索")
var SearchText = flag.String("text", "", "准备搜索的文本")
var SearchT = flag.String("t", "", "准备搜索的文本")
var Example = flag.String("example", "", example)
var Exclude = flag.String("exc", "", "排除指定文本 -exc // 表示排除带//的行")
var ExcludeE = flag.String("e", "", "排除指定文本 -e // 表示排除带//的行")
var ExecPath string

func main() {
	flag.Parse()

	searchText := ""

	if *SearchText == "" {
		if *SearchT == "" {
			fmt.Println("请输入要搜索的内容！")
			return
		}
		searchText = *SearchT
	} else {
		searchText = *SearchText
	}

	ExecPath, err := os.Getwd()
	if err != nil {
		fmt.Println("获取执行路径错误！")
		return
	}

	if *SearchPath == "" {
		if *SearchP == "" {
			ExecPath = strings.ReplaceAll(ExecPath, "\\search.exe", "")
		} else {
			ExecPath = *SearchP
		}
	} else {
		ExecPath = *SearchPath
	}

	searchFile := ""
	if *SearchFile == "" {
		if *SearchF == "" {
			fmt.Println("请输入要搜索的文件！")
			return
		} else {
			searchFile = *SearchF
		}
	} else {
		searchFile = *SearchFile
	}
	exc := ""
	if *Exclude == "" {
		exc = *ExcludeE
	}

	s := New(searchFile, ExecPath, searchText, exc)

	fmt.Println("搜索路径SearchPath：", ExecPath)
	fmt.Println("执行路径：", ExecPath)
	fmt.Println("搜索内容：", searchText)
	fmt.Println("搜索文件：", searchFile)
	s.Run()
}

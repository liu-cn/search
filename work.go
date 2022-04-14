package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

type Search struct {
	execText       string //排除文本
	hasExec        bool
	parseFileCount uint
	totalFileCount uint
	startTime      time.Time
	findCount      uint
	searchFile     string
	searchPath     string
	searchText     string
	regFile        bool
	suffix         string
}

func New(file, path, text, execText string) *Search {

	sea := Search{
		searchPath: path,
		searchFile: file,
		searchText: text,
	}
	if strings.Contains(file, "*") {
		sea.suffix = strings.ReplaceAll(file, "*", "")
		sea.regFile = true
	}
	sea.execText = execText
	sea.hasExec = execText != ""

	return &sea
}

func (s *Search) check(fileName string) bool {
	if s.regFile {
		return strings.HasSuffix(fileName, s.suffix)
	}
	return fileName == s.searchFile
}

func (s *Search) echoInfo() {
	//查找到以下结果
	fmt.Printf("总共搜索:%v个文件，搜索到:%v处结果,总计耗时：%v\n", s.totalFileCount, s.findCount, time.Now().Sub(s.startTime))
}

//去除前缀空格
func trimPrefixText(s *string) *string {
	f := false
	for {
		if strings.HasPrefix(*s, "\\O") {
			*s = strings.TrimPrefix(*s, "\\O")
			f = true
		} else {
			f = false
		}

		if strings.HasPrefix(*s, "\t") {
			*s = strings.TrimPrefix(*s, "\t")
			f = true
		}
		if f == false {
			goto Exit
		}
	}

Exit:
	return s
}
func (s *Search) find(path string) {

	if strings.HasSuffix(path, "\\") {
		path = path[0 : len(path)-1]
	}
	dir, err := os.ReadDir(path)
	if err != nil {
		return
	}

	for _, file := range dir {
		abs := ""
		if runtime.GOOS == "darwin" {
			abs = path + "/" + file.Name()
		} else {
			abs = path + "\\" + file.Name()
		}

		if file.IsDir() {
			s.find(abs)
		} else {
			if s.check(file.Name()) {
				s.totalFileCount++
				file, err := os.Open(abs)
				if err != nil {
					return
				}
				scanner := bufio.NewScanner(file)
				var line uint = 1
				for scanner.Scan() {
					if s.hasExec {
						if strings.Contains(scanner.Text(), s.searchText) && !strings.Contains(scanner.Text(), s.execText) {
							text := scanner.Text()
							fmt.Printf("%v:%v %v\n ", abs, line, "\t"+*trimPrefixText(&text))
							s.findCount++
						}
					} else {
						if strings.Contains(scanner.Text(), s.searchText) {
							text := scanner.Text()
							fmt.Printf("%v:%v %v\n ", abs, line, "\t"+*trimPrefixText(&text))
							s.findCount++
						}
					}
					line++
				}
				file.Close()
				s.parseFileCount++
			}
		}
	}
}

func (s *Search) Run() {
	s.startTime = time.Now()
	s.find(s.searchPath)
	s.echoInfo()
}

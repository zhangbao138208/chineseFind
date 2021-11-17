package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"syscall"
)
const (
	ORIGIN = "http://7xvxof.com1.z0.glb.clouddn.com"
	TARGET = "http://qiniu.*****.**"
)
var (
	chineseMap  = map[string]string{

	}
	enMap  = map[string]string{

	}
)

func main() {
	files := getFiles("./test")
	for _, file := range files {
		_, needHandle, err := readFile(file)
		if err != nil {
			panic(err)
		}
		if needHandle {
			//err = writeToFile(file, output)
			//if err != nil {
			//	panic(err)
			//}
			//fmt.Println(file)
		}
	}
	mjson,_ :=json.Marshal(chineseMap)

	mString :=string(mjson)
	write("zh.json",strings.Replace(mString,`\u001b[38;5;167m`,"",-1))



	mjson2,_ :=json.Marshal(enMap)

	mString2 :=string(mjson2)
	write("en.json",strings.Replace(mString2,`\u001b[38;5;167m`,"",-1))
}

func getFiles(path string) []string {
	files := make([]string, 0)
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	return files
}
func readFile(filePath string) ([]byte, bool, error) {
	f, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		return nil, false, err
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	needHandle := false
	output := make([]byte, 0)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				return output, needHandle, nil
			}
			return nil, needHandle, err
		}
		reg1 := regexp.MustCompile("[\u4e00-\u9fa5]+")
		if reg1 == nil {
			fmt.Println("regexp err")

		}
		//根据规则提取关键信息
		result1 := reg1.FindAllStringSubmatch(string(line),-1)
		//fmt.Println("line= result1 = ",string(line), result1)
		for _, ss := range result1 {
			for _, s := range ss {

				ret:=cmd(s)
				if ret != nil {

					key := strings.Replace(strings.Join(ret,"_"),`\u001b[38;5;167m`,"",-1)
					fmt.Println("len(ret)=========================",len(ret),strings.Join(ret,"_"),key)
					chineseMap[key] = s
					enMap[key] = strings.Join(ret," ")
				}
			}
		}
		//regexp.Match(`^[\u4e00-\u9fa5]`,line)

		if ok, _ := regexp.Match(ORIGIN, line); ok {
			reg := regexp.MustCompile(ORIGIN)
			newByte := reg.ReplaceAll(line, []byte(TARGET))
			output = append(output, newByte...)
			output = append(output, []byte("\n")...)
			if !needHandle {
				needHandle = true
			}
		} else {
			output = append(output, line...)
			output = append(output, []byte("\n")...)
		}
	}
	return output, needHandle, nil
}

func writeToFile(filePath string, outPut []byte) error {
	f, err := os.OpenFile(filePath, syscall.O_CREAT, 0600)
	defer f.Close()
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(f)
	_, err = writer.Write(outPut)
	if err != nil {
		return err
	}
	writer.Flush()
	return nil
}
func cmd(c string) []string {
	cmd := exec.Command("./gdict.exe", c)
	if runtime.GOOS == "windows" {
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	}
	//Start执行不会等待命令完成，Run会阻塞等待命令完成。
	//err := cmd.Start()
	//err := cmd.Run()
	//cmd.Output()函数的功能是运行命令并返回其标准输出。
	buf, err := cmd.Output()

	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			status := exitErr.Sys().(syscall.WaitStatus)
			switch {
			case status.Exited():
				fmt.Printf("Return exit error: exit code=%d\n", status.ExitStatus())
			case status.Signaled():
				fmt.Printf("Return exit error: signal code=%d\n", status.Signal())
			}
		} else {
			fmt.Printf("Return other error: %s\n", err)
		}
	} else {
		fmt.Printf("Return OK\n")
		//fmt.Println(string(buf))
		//reg,_ := regexp.Compile(`Exps:[\s\S]*1.([\s\S]*)网络释义:`)
		reg := regexp.MustCompile(`翻译:([\s\S]*)网络释义:`)
		ret := reg.FindAllStringSubmatch(string(buf),-1)
		//fmt.Printf("=========ret=%d=============\n", len(ret))
		//for i, ss := range ret {
		//	fmt.Println("======start======","i=",i,"ss=",ss, "len(ss)=", len(ss),"=====end=======")
		//	for i2, s := range ss {
		//		fmt.Println("======2start======","i2=",i2,"s=",s,"=====2end=======")
		//	}
		//}
		if len(ret) == 1 {
			str := strings.Replace(ret[0][1],"\n","",-1)
			str1 := strings.Replace(str,"	","",-1)
			str2:=strings.TrimSpace(str1)
			//fmt.Println("翻译--------:",str2,"--------")
			ss :=strings.Split(str2," ")
			if len(ss) > 2 {
				fmt.Println(ss[2:])
				for _, s := range ss[2:] {
					fmt.Println(s)
				}
				return ss[2:]
			}


		}

	}
	return nil
}
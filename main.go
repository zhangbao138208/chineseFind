package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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
	alMap = map[string] bool {

	}
	keys = []string {"title","label","placeholder"}
	JsonMap = map[string]string{}
)

func main() {
	// 打开json文件
	jsonFile, err := os.Open("./json/zh.json")

	// 最好要处理以下错误
	if err != nil {
		fmt.Println(err)
	}

	// 要记得关闭
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue,&JsonMap)

	files := getFiles("./test")
	for _, file := range files {
		output, needHandle, err := readFile(file)
		if err != nil {
			panic(err)
		}
		if needHandle {
			err = writeToFile(file, output)
			if err != nil {
				panic(err)
			}
			fmt.Println(file)
		}
	}
	//mjson,_ :=json.Marshal(chineseMap)
	//
	//mString :=string(mjson)
	//write("zh.json",strings.Replace(mString,`\u001b[38;5;167m`,"",-1))
	//
	//
	//
	//mjson2,_ :=json.Marshal(enMap)
	//
	//mString2 :=string(mjson2)
	//write("en.json",strings.Replace(mString2,`\u001b[38;5;167m`,"",-1))
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
			continue
			for _, s := range ss {
				_,ok:= alMap[s]
				if ok {
					continue
				}
                //ret:=cmd(s)
				ret := translate(s)
				//if ret != nil {
				//
				//	key := strings.Replace(strings.Join(ret,"_"),`\u001b[38;5;167m`,"",-1)
				//	fmt.Println("len(ret)=========================",len(ret),strings.Join(ret,"_"),key)
				//	chineseMap[key] = s
				//	enMap[key] = strings.Join(ret," ")
				//}
				if ret!= "" {
					chineseMap[strings.Join(strings.Split(ret," "),"_")] = s
					enMap[strings.Join(strings.Split(ret," "),"_")] = ret
					alMap[s] = true
				}
			}
		}
		//regexp.Match(`^[\u4e00-\u9fa5]`,line)
		// 不需要替换的
		fmt.Println(string(line))
		if strings.HasPrefix(string(line),"//") ||
			strings.HasPrefix(string(line),"*") ||
			strings.HasPrefix(string(line),"<!"){
			fmt.Println("===============HasPrefix=============")
           continue
		}
		for k, v := range JsonMap {
			fmt.Println("-----------------",string(line),"-----------------")
			fmt.Println("jsonMap",k,v)
			for _, key := range keys {
				pattern := fmt.Sprint(key,"=\"",v,"\"")

				if ok, _ := regexp.Match(pattern, line); ok {

					target := fmt.Sprintf(":%s=\"i18n.t('%s')\"",key,k)
					reg := regexp.MustCompile(pattern)
					newByte := reg.ReplaceAll(line, []byte(target))
					line = newByte

					if !needHandle {
						needHandle = true
					}
					fmt.Println(pattern,string(line))
				}
				pattern = fmt.Sprint(":",key,"=\"",v,"\"")
				if ok, _ := regexp.Match(pattern, line); ok {
					target := fmt.Sprintf(":%s=\"i18n.t('%s')\"",key,k)
					reg := regexp.MustCompile(pattern)
					newByte := reg.ReplaceAll(line, []byte(target))
					line = newByte

					if !needHandle {
						needHandle = true
					}
					fmt.Println(pattern,string(line))
				}
				pattern = fmt.Sprintf(":%s=\"'%s'\"",key,v)
				if ok, _ := regexp.Match(pattern, line); ok {
					target := fmt.Sprintf(":%s=\"i18n.t('%s')\"",key,k)
					reg := regexp.MustCompile(pattern)
					newByte := reg.ReplaceAll(line, []byte(target))
					line = newByte

					if !needHandle {
						needHandle = true
					}
					fmt.Println(pattern,string(line))
				}

			}
			pattern := fmt.Sprintf("\"%s\"",v)
			if ok, _ := regexp.Match(pattern, line); ok {
				target := fmt.Sprintf("i18n.t('%s')",k)
				reg := regexp.MustCompile(pattern)
				newByte := reg.ReplaceAll(line, []byte(target))
				line = newByte

				if !needHandle {
					needHandle = true
				}
				fmt.Println(pattern,string(line))
			}
			pattern = fmt.Sprintf("'%s'",v)
			if ok, _ := regexp.Match(pattern, line); ok {
				reg := regexp.MustCompile(pattern)
				target := fmt.Sprintf("i18n.t('%s')",k)
				newByte := reg.ReplaceAll(line, []byte(target))
				line = newByte

				if !needHandle {
					needHandle = true
				}
				fmt.Println(pattern,string(line))
			}
			pattern = fmt.Sprintf("%s",v)
			if ok, _ := regexp.Match(pattern, line); ok {
				reg := regexp.MustCompile(pattern)
				target := fmt.Sprintf("{{i18n.t('%s')}}",k)
				newByte := reg.ReplaceAll(line, []byte(target))
				line = newByte

				if !needHandle {
					needHandle = true
				}
				fmt.Println(pattern,string(line))
			}

		}

		output = append(output, line...)
		output = append(output, []byte("\n")...)
		fmt.Println("else===============:",string(output))

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
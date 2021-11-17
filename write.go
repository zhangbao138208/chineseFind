package main

import (
	"fmt"
	"os"
)

func write (fileName,content string)  {
	filePath := "./json/"+fileName
	f,err := os.Create(filePath)
	defer f.Close()
	if err !=nil {
		fmt.Println(err.Error())
	} else {
		_,err=f.Write([]byte(content))
	}
}

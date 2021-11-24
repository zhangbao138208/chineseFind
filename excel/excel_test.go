package excel

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io/ioutil"
	"os"
	"syscall"
	"testing"

)
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
var (
	chineseMap = map[string]string{}
	enMap      = map[string]string{}

)
func Init()  {
	// 打开json文件
	jsonFile, err := os.Open("../json/zh.json")

	// 最好要处理以下错误
	if err != nil {
		fmt.Println(err)
	}

	// 要记得关闭
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue,&chineseMap)


	// 打开json文件
	enFile, err := os.Open("../json/en.json")

	// 最好要处理以下错误
	if err != nil {
		fmt.Println(err)
	}

	// 要记得关闭
	defer enFile.Close()

	byteValue, _ = ioutil.ReadAll(enFile)
	json.Unmarshal(byteValue,&enMap)








}

func TestExcel(t *testing.T) {
	Init()
	f, err := excelize.OpenFile("D:\\Documents and Settings\\Timothy\\My Documents\\i Talk\\72973226-104764-timothy\\file\\01-质检语言包_质检业务(1).xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(f.GetSheetList())
	for _,sheet := range f.GetSheetList() {
		rows, err := f.GetRows(sheet)
		if err != nil {
			fmt.Println(err)
			return
		}
		for i, row := range rows {
			if i == 0 {
				continue
			}
			//for j, colCell := range row {
			//	fmt.Print(j,":",colCell, "\t")
			//}
			//fmt.Println()
			for K, V := range chineseMap {

				if V == row[3] {
					enMap[K] = row[4]
					f.SetCellValue(sheet,fmt.Sprint("B",i+1),K)
				}
			}
		}
	}

	mjson, _ := json.Marshal(chineseMap)

	writeToFile("zh1.json", []byte(mjson))

	mjson2, _ := json.Marshal(enMap)

	writeToFile("en1.json", mjson2)
	// Save spreadsheet by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}

}

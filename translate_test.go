package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"
)

func TestTranslate(t *testing.T) {
	translate("")
}

func TestName(t *testing.T) {
	val := "学校"
	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		// chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// navigate to a page, wait for an element, click
	var example string
	err := chromedp.Run(ctx,
		chromedp.Navigate(fmt.Sprintf(`https://translate.google.cn/?sl=auto&tl=en&text=%s&op=translate`,val)),
		// wait for footer element is visible (ie, page is loaded)
		chromedp.WaitVisible(`#yDmH0d > c-wiz > div > div.WFnNle > c-wiz > div.OlSOob > c-wiz > div.ccvoYb > div.AxqVh > div.OPPzxe > c-wiz.P6w8m.BDJ8fb > div.dePhmb`),

		// retrieve the text of the textarea
		chromedp.Text(`#yDmH0d > c-wiz > div > div.WFnNle > c-wiz > div.OlSOob > c-wiz > div.ccvoYb > div.AxqVh > div.OPPzxe > c-wiz.P6w8m.BDJ8fb > div.dePhmb > div > div.J0lOec > span.VIiyi > span > span`, &example),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Go's time.After example:\n%s", example)
}

//var keys = []string {"title","label","placeholder"}
func TestJson(t *testing.T) {
	// 打开json文件
	jsonFile, err := os.Open("./json/zh.json")

	// 最好要处理以下错误
	if err != nil {
		fmt.Println(err)
	}

	// 要记得关闭
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	jsonMap := make(map[string]string)
	json.Unmarshal(byteValue,&jsonMap)
	fmt.Println(jsonMap)
	fmt.Println(string(byteValue))
	for k, v := range jsonMap {
		fmt.Println(k,v)
	}
}
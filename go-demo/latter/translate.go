package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"bufio"
	"strings"
)

type DictRequest struct {
	TransType string `json:"trans_type"`
	Source string `json:"source"`
	UserId string `json:"user_id"`
}


type DictResponse struct {
	Rc int `json:"rc"`
	Wiki struct {
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"`
			En string `json:"en"`
		} `json:"prons"`
		Explanations []string `json:"explanations"`
		Synonym []string `json:"synonym"`
		Antonym []string `json:"antonym"`
		WqxExample [][]string `json:"wqx_example"`
		Entry string `json:"entry"`
		Type string `json:"type"`
		Related []interface{} `json:"related"`
		Source string `json:"source"`
	} `json:"dictionary"`
}


func query(word string) {
	client := &http.Client{}
	// var data = strings.NewReader(`{"trans_type":"en2zh","source":"good"}`)
	request := DictRequest{
		TransType: "en2zh",
		Source: word,
	}
	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	var data = bytes.NewReader(buf)
	req, err := http.NewRequest("POST", "https://lingocloud.caiyunapp.com/v1/dict", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Cookie", "_gcl_au=1.1.1637146438.1681824257; _ga=GA1.2.316353676.1681824255; _gid=GA1.2.117489048.1681824258; _gat_gtag_UA_185151443_2=1; _ga_65TZCJSDBD=GS1.1.1681824255.1.0.1681824260.0.0.0; _ga_R9YPR75N68=GS1.1.1681824256.1.0.1681824260.56.0.0")
	req.Header.Set("Origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("Referer", "https://fanyi.caiyunapp.com/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36")
	req.Header.Set("X-Authorization", "token:qgemv4jr1y38jyq6vhvi")
	req.Header.Set("app-name", "xy")
	req.Header.Set("device-id", "566664f560824e043a61624cc3f403f8")
	req.Header.Set("os-type", "web")
	req.Header.Set("os-version", "")
	req.Header.Set("sec-ch-ua", `"Chromium";v="112", "Google Chrome";v="112", "Not:A-Brand";v="99"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%s\n", bodyText)

	if resp.StatusCode != 200{
		log.Fatal("bad request", resp.StatusCode)
	}

	var dictResponse DictResponse
	//反序列化接收到的结果，并将结果保存在dictResponse当中
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%#v\n", dictResponse)
	fmt.Println(word, "的翻译结果为：")
	fmt.Println("音标：" ,"UK:", dictResponse.Dictionary.Prons.En, "US:", dictResponse.Dictionary.Prons.EnUs)

	for _, item := range dictResponse.Dictionary.Explanations {
		fmt.Println(item)
	}
}


func main() { 
	// if len(os.Args) != 2 {
	// 	fmt.Fprintf(os.Stderr, `usage: simpleDict WORD example: simpledict hello`)
	// 	os.Exit(1)
	// }
	// word := os.Args[1]
	// query(word)

	fmt.Println("please input a latter: ")
	reader := bufio.NewReader(os.Stdin)

		//获取到输入的内容
		word, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error1", err)
			return
		}

		//去除换行符
		word = strings.TrimSuffix(word, "\r\n")

		//将输入的内容转换为int整型
		// guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("error2", err)
			return
		}
		// fmt.Println(input)
		query(word)
}
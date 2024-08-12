package utils

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"time"
)

/**
http请求工具类
@MYT 20240810
*/

var client = http.Client{Timeout: time.Second * 10}

// ExecPost POST请求
func ExecPost(url string, body []byte, header map[string]string) []byte {
	return ExecHttp("POST", url, header, body, nil)
}

// ExecGet GET请求
func ExecGet(url string, params map[string]string, header map[string]string) []byte {
	return ExecHttp("GET", url, header, nil, params)
}

// ExecHttp 通用请求逻辑
func ExecHttp(methodType string, url string, header map[string]string, bodyMsg []byte, params map[string]string) []byte {
	var body io.Reader
	if methodType == "POST" {
		if bodyMsg != nil {
			// 读取请求体
			body = bytes.NewReader(bodyMsg)
		}
	} else if methodType == "GET" {
		if params != nil {
			// 拼接url参数
			url = url + "?"
			for k, v := range params {
				url = url + "&" + k + "=" + v
			}
		}
	}
	req, err := http.NewRequest(methodType, url, body)
	for k, v := range header {
		req.Header.Add(k, v)
	}
	// 处理异常
	if err != nil {
		log.Printf("request error.%v", err)
		return nil
	}
	// 发起请求
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("request error. %v", err)
		return nil
	}
	// 解析结果
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("response read error. %v", err)
		return nil
	}
	return res
}

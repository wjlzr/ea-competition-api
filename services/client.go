package services

import (
	"ea-competition-api/boot/log"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"go.uber.org/zap"
)

//统一请求分发
func Request(method, url string, body io.Reader) (request *http.Request, err error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Logger().Error("services http request err：", zap.Error(err))
		return request, err
	}
	if method == http.MethodPost {
		req.Header.Set("Content-Type", "application/json")
	}
	return req, nil
}

//返回参数统一处理
func ResponseHandle(request *http.Request) (content []byte, err error) {
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Logger().Error("services ResponseHandle http请求错误 err：", zap.Error(err))
		return content, err

	}
	defer resp.Body.Close()
	content, _ = ioutil.ReadAll(resp.Body)
	fmt.Printf(" \n")
	fmt.Printf("OpenApi请求值：%+v \n", resp.Request)
	fmt.Printf("OpenApi返回值：%s \n", string(content))
	return content, nil
}

package xjson

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// 远程json文件转为字符串
func JsonToString(url, key string) ([]string, map[string]int) {

	res1, err := http.Get(url)
	if err != nil {
		return nil, nil
	}
	defer res1.Body.Close()
	byteValue, _ := ioutil.ReadAll(res1.Body)

	var result map[string][]string
	_ = json.Unmarshal([]byte(byteValue), &result)
	brokerIds := result[key]

	var mapSort map[string]int = make(map[string]int, 0)
	for _, v := range brokerIds {
		ss := strings.Split(v, "-")
		var (
			num  int
			code string
		)
		if len(ss) == 2 {
			num, _ = strconv.Atoi(ss[1])
			code = ss[0]
			mapSort[code] = num
		}
	}
	return brokerIds, mapSort
}

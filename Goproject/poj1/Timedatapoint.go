//上传人：牛康力
package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

//构建模拟时间数据点结构体
type Data struct {
	T1 [100]Datapoint
	T2 [100]Datapoint
}

type Datapoint struct {
	S string `json:"start"`
	E string `json:"end"`
}

func main() {
	http.HandleFunc("/api", dataHandler)
	err := http.ListenAndServe(":8080", nil)
	//调用函数，向http发起请求，url: localhost:8080/api
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	var data = Data{}
	//确保在get方法下向http发起请求
	if r.Method == "GET" {
		data = GenerateCalculatepoint()
		//以 json 格式导出 data 数据
		data_json, _ := json.Marshal(data)
		//设置 content-type ，让客户端明晰数据格式为json
		w.Header().Set("content-type", "application/json")
		w.Write(data_json)
	}
}

func GenerateCalculatepoint() Data {
	var data = Data{}
	rand.Seed(time.Now().UnixNano())
	for k, _ := range data.T1 {
		//使用 range 遍历 T1，T2 数组并计算
		data.T1[k] = *timeCalculate("08:00")
		data.T2[k] = *timeCalculate("20:00")
	}
	return data
}

//定义 timeCalculate 函数，调用一次，返回一个数据点
func timeCalculate(beginTime string) *Datapoint {
	datapoint := new(Datapoint)
	//生成随机的 start time
	t, _ := time.Parse("15:04", beginTime)
	t = t.Add(time.Duration(rand.Intn(60)) * time.Minute)
	t = t.Add(time.Duration(rand.Intn(12)) * time.Hour)
	datapoint.S = t.Format("15:04")

	//计算相应的 end time
	t = t.Add(time.Hour)
	//加上 [0,60] 分钟
	t = t.Add(time.Duration(rand.Intn(61)) * time.Minute)
	datapoint.E = t.Format("15:04")
	return datapoint
}

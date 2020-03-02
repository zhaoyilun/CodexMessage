package main

import (
	"encoding/json"
	"fmt"
	"github.com/eoscanada/eos-go"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var api = eos.New("http://198.74.53.102:8888")
func main(){
	//write()
	go Delay(30,write)
	http.HandleFunc("/", sayhelloName) //设置访问的路由
	err := http.ListenAndServe(":8085", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	a:=read("zhaoyilun532")
	b:=converTime(a)
	fmt.Println(b)
}
func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	x:=strings.Join(r.Form["index"],"")
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
	data:=read(x)
	message:=converTime(data)
	fmt.Fprintf(w, message) //这个写入到w的是输出到客户端的
	//files:=readFile()
	//fmt.Fprintf(w, files)
}
func write(){
	file, err := os.OpenFile("./2.txt", os.O_RDWR|os.O_CREATE, 0766)
	defer file.Close()
	if err != nil {
		fmt.Println("openfile error: ",err)
	}
	table :=eos.GetTableRowsRequest{
		Code:       "rexclaimdrop",
		Scope:      "rexclaimdrop",
		Table:      "claims",
		Limit:      10000,
		JSON:       true,
	}
	tableInfo,err:=api.GetTableRows(table)
	if err!=nil{
		fmt.Println(1,err)
	}
	tableJson,err:=json.MarshalIndent(tableInfo.Rows," "," ")
	file.Write(tableJson)
}
func readFile()string{
	file, err := ioutil.ReadFile("./2.txt")

	if err != nil {
		fmt.Println("readFile error: ",err)
	}
	return string(file)
}
func read(name string)int64{
	file, err := os.Open("./2.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("Read File Error：",err);
	}
	var info []tableMessage
	decoder:=json.NewDecoder(file)
	err=decoder.Decode(&info)
	if err != nil {
		fmt.Println("decode error",err)
	}
	tableMap := make(map[string]int64)
	for _,x := range info {
		tableMap[x.Account]=x.Time
	}
	return tableMap[name]
}

type tableMessage struct {
	Account string "account"
	Time int64 `json:"lastclaim"`
}
func converTime(sec int64) string{
	return time.Unix(sec,0).Format("2006-01-02 15:04:05")
}



//一个函数用于计算延时时间
func GetDelayTime(timeSec int64)time.Duration{
	return time.Second * time.Duration(timeSec)
}
//一个函数用于实现定时任务
func Delay(timeSec int64,myFunction func()){
	t1:=time.Tick(GetDelayTime(timeSec))
	for{
		select {
		case <-t1:
			myFunction()
		}
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"github.com/eoscanada/eos-go"
	"os"
	"time"
)

var api = eos.New("http://198.74.53.102:8888")
func main(){
	args:=os.Args
	if args ==nil || len(args)<1{
		Usage()
		return
	}
	name:=args[1]
	a:=getInfo()
	timeInfo :=converTime(a[name])
	fmt.Println("你上一次领取空投的时间为：",timeInfo)
}
var Usage= func() {
	fmt.Println("请输入名称")
}
type tableMessage struct {
	Account string "account"
	Time int64 `json:"lastclaim"`
}
func converTime(sec int64) string{
	return time.Unix(sec,0).Format("2006-01-02 15:04:05")
}
func getInfo2()iint{
	a:=getInfo()
	return a
}
func getInfo()map[string]int64{
	//api := eos.New("http://198.74.53.102:8888")
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
	if err!=nil{
		fmt.Println(3,err)
	}
	tableStruct :=new([]tableMessage)
	json.Unmarshal(tableJson,&tableStruct)
	tableMap := make(map[string]int64)
	for _,x := range *tableStruct {
		tableMap[x.Account]=x.Time
	}
	return tableMap
}

//一个函数用于计算延时时间
func GetDelayTime(timeSec int64)time.Duration{
	return time.Second * time.Duration(timeSec)
}
//一个函数用于实现定时任务
func (data *iint)Delay(timeSec int64,myFunction func()iint){
	t1:=time.Tick(GetDelayTime(timeSec))
	for{
		select {
		case <-t1:
			*data=myFunction()
		}
	}
}
type iint map[string]int64
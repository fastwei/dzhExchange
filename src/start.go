package main

import (
	"exchange"
	"fmt"
	//	"strconv"
)

func main() {
	initSys()
}

/**
 *初始化码表
 *订阅行情
 *初始化各个品种的撮合队列
 *初始化撮合池
 *执行首次撮合后等待行情触发再次撮合
 *绑定端口
 */

func initSys() {
	codesArr := exchange.LoadData()
	exchange.CodesTable = codesArr
	fmt.Println("service is start now..")
	exchange.MatchPool = make(map[string]*exchange.StkMatchQueue)
	//根据商品代码创建队列
	for i := 0; i < len(codesArr); i++ {
		exchange.Stqueue = new(exchange.StkMatchQueue)
		exchange.Stqueue.InitQueue()
		co := new(exchange.CommonOrder)
		co.Oid = i
		co.Code = codesArr[i]
		co.Price = 10.2
		co.BsFlag = 0

		co2 := new(exchange.CommonOrder)
		co2.Oid = i
		co2.Code = codesArr[i]
		co2.Price = 5.3
		co2.BsFlag = 1
		//分别插入买卖队列
		exchange.Stqueue.InsertOrder(co)
		exchange.Stqueue.InsertOrder(co2)

		//fmt.Println("cur queue:", exchange.Stqueue)
		//执行首次撮合
		//globleVar.Ss.DealMatch(exchange.Stqueue, quote,globleVar.SignQuote)
		exchange.MatchPool[codesArr[i]] = exchange.Stqueue
	}
	fmt.Println("match pool size:", len(exchange.MatchPool))

	//订阅行情
	go exchange.WakeUp(codesArr)
	//开启请求处理线程
	go exchange.StartProcess(exchange.Sign)
	go exchange.PushBack()
	
	for {
//		<-exchange.SignQuote
		<-exchange.Sign
//		select {
//		case <-exchange.SignQuote:
//			fmt.Println("recv quote ok")
//		//wait to send to counter
//		case <-exchange.Sign:
//			fmt.Println("send to counter ok")
//		default:
//		}
	}

}

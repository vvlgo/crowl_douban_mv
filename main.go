package main

import (
	"strconv"
	"fmt"
	"crowl_douban_mv/engine"
	"crowl_douban_mv/scheduler"
	"crowl_douban_mv/fetcher"
)

var mvContent = make(chan string)
var mvUrl= make(chan string)

func main() {
	var i=0
	var n=0
	var num=0
	urls:=make([]string,0)
	for {
		url:=`https://movie.douban.com/j/new_search_subjects?sort=T&range=0,10&tags=&start=`+strconv.Itoa(n)
		fmt.Printf(url+"\n")
		if n<9960{
			urls=append(urls,url)
		}
		if n > 9960 {
			_, err := fetcher.FetchUrls(url)
			if err!=nil{
				num++
				s :=err.Error()
				fmt.Println(s)
				break
			}
			urls=append(urls,url)
		}

		n=(i+1)*20+1
		i++
		//time.Sleep(time.Duration(rand.Intn(100))*time.Millisecond)

	}
	fmt.Println(len(urls))
	e:=engine.ConcurrentEngine{
		Scheduler:&scheduler.QueuedScheduler{},
		WorkerCount:5,
		MvWorkerCount:4,
		Total:20*(len(urls)-num),
	}
	e.Run(urls)



}



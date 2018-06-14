package engine

import (
	"fmt"
	"crowl_douban_mv/fetcher"
	"strconv"
	"crowl_douban_mv/scheduler"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	MvWorkerCount int
	Total int
}

type Scheduler interface {
	Submit(*string)
	ConfigureMasterWorkerChan(chan string)
	WorkerReady(w chan string)
	Run()
}

func (e *ConcurrentEngine) Run(seeds []string) {
	out := make(chan []string)
	e.Scheduler.Run()



	for i := 0; i < e.WorkerCount; i++ {
		createWorker(out,e.Scheduler)
	}
	for _, r := range seeds {
		e.Scheduler.Submit(&r)
	}

	e2:=ConcurrentEngine{
		Scheduler:&scheduler.MvQueuedScheduler{},
		WorkerCount:3,
		MvWorkerCount:5,
	}
	ress:=make([]string,0)
	for {

		result := <-out
		fmt.Println(len(result))
		//for _, v := range result {
		//	fmt.Printf("url :%s\n", v)
		//	//content, _ := fetcher.Fetch(v)
		//	//fetcher.GetMovie(string(content))
		//}
		//go func() {
		//	e2.RunMv(result)
		//}()
		ress = append(ress,result...)
		if len(ress)>=(e.Total-5*20){
			break
		}

	}
	fmt.Println("================"+strconv.Itoa(len(ress)))
	e2.RunMv(ress)


}




func createWorker(out chan []string,s Scheduler) {
	in :=make(chan string)
	go func() {
		for {
			s.WorkerReady(in)
			url := <-in
			urls, err := fetcher.FetchUrls(url)
			if err != nil {
				fmt.Println(err.Error())
				//close(out)
				continue
			}


			out <- urls
		}
	}()
}

func (e *ConcurrentEngine) RunMv(seeds []string) {
	out := make(chan bool)
	e.Scheduler.Run()



	for i := 0; i < e.MvWorkerCount; i++ {
		createMvWorker(e.Scheduler,out)
	}
	for _, r := range seeds {
		e.Scheduler.Submit(&r)
	}

	for {
		result := <-out
		fmt.Println(result)

	}
}

func createMvWorker(s Scheduler,out chan bool) {
	in :=make(chan string)
	go func() {
		for {
			s.WorkerReady(in)
			url := <-in
			content, err := fetcher.Fetch(url)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			fetcher.GetMovie(string(content))
			out <- true
		}
	}()
}

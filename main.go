package main

import (
	"crowl_douban_mv/fetcher"
	"fmt"
	"strconv"
)

var mvContent  = make(chan string)
var mvUrl= make(chan string)

func main() {
	var i=0
	var n=0
	for {
		url:=`https://movie.douban.com/j/new_search_subjects?sort=T&range=0,10&tags=&start=`+strconv.Itoa(n)
		fmt.Printf(url+"\n")
		err := fetcher.FetchUrls(url)
		if err!=nil{
			fmt.Println(err)
			break
		}
		i++
		n=(i+1)*20+1
	}

}



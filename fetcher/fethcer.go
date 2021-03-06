package fetcher

import (
	"net/http"
	"fmt"
	"golang.org/x/text/transform"
	"io/ioutil"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"bufio"
	"golang.org/x/net/html/charset"
	"log"
	"encoding/json"
	"errors"
	"crowl_douban_mv/models"
	"time"
	"crowl_douban_mv/proxyaby"
)

type Mv struct {
	Directors []string `json:"directors"`
	Rate      string   `json:"rate"`
	CoverX    int64    `json:"cover_x"`
	Star      string   `json:"star"`
	Title     string   `json:"title"`
	Url       string   `json:"url"`
	Casts     []string `json:"casts"`
	Cover     string   `json:"cover"`
	Id        string   `json:"id"`
	CoverY    int64    `json:"cover_y"`
}
type MyDatas struct {
	Data []Mv `json:"data"`
}


var rateLimiter = time.Tick(200*time.Millisecond)
var rateLimiter2 = time.Tick(200*time.Millisecond)
//获取url页面内容准备获取信息
func Fetch(url string) ([]byte, error) {
	<- rateLimiter
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	resp, err := proxyaby.Client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)

}
//通过豆瓣隐藏接口直接获取json数据得到每部影视的url地址
func FetchUrls(url string) ([]string,error) {
	<-rateLimiter2

	request, _ := http.NewRequest(http.MethodGet, url, nil)
	resp, err := proxyaby.Client.Do(request)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil,fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	myDatas := MyDatas{Data: make([]Mv, 0)}
	for i, _ := range myDatas.Data {
		myDatas.Data[i].Directors = make([]string, 0)
		myDatas.Data[i].Casts = make([]string, 0)
	}

	bodyReader := bufio.NewReader(resp.Body)
	bytes, _ := ioutil.ReadAll(bodyReader)
	err = json.Unmarshal(bytes, &myDatas)
	if err != nil {
		fmt.Printf("json to mv err:%s", err)
	}
	//fmt.Println(string(bytes))
	if len(myDatas.Data) == 0 {
		return nil,errors.New("data nil")
	}
	var urls = make([]string,0)
	for _, v := range myDatas.Data {
		//fmt.Printf("%d  %s\n",i,v.Url)
		//content, err := Fetch(v.Url)
		urls=append(urls,v.Url)
		if err != nil {
			panic(err)
		}
		//getMovie(string(content))

	}

	return urls,nil

}
//页面重新解码为UTF-8
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

//获取页面内容的数据
func GetMovie(content string) {
	var movieInfo models.MovieInfo

	movieInfo.Movie_name = models.GetMovieName(content)
	//记录电影信息
	if movieInfo.Movie_name != "" {
		movieInfo.Movie_director = models.GetMovieDirector(content)
		movieInfo.Movie_main_character = models.GetMovieMainCharacters(content)
		movieInfo.Movie_type = models.GetMovieGenre(content)
		movieInfo.Movie_on_time = models.GetMovieOnTime(content)
		movieInfo.Movie_grade = models.GetMovieGrade(content)
		movieInfo.Movie_span = models.GetMovieRunningTime(content)

		models.AddMovie(&movieInfo)
	}
}


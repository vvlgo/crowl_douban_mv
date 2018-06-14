package proxyaby

import (
	"net/url"
	"net/http"
)

// 代理服务器
const ProxyServer = "http-dyn.abuyun.com:9020"

// 代理隧道验证信息
const proxyUser  = "XXXXXXXXX";
const proxyPass  = "XXXXXXXXX";


type AbuyunProxy struct {
	AppID string
	AppSecret string
}
var Client =AbuyunProxy{}.ProxyClient()
func (p AbuyunProxy) ProxyClient() http.Client {
	proxyUrl, _ := url.Parse("http://"+ p.AppID +":"+ p.AppSecret +"@"+ ProxyServer)
	return http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
}


func init(){
	Client=AbuyunProxy{AppID: proxyUser, AppSecret: proxyPass}.ProxyClient()
}

//func main()  {
//	targetUrl := "http://test.abuyun.com"
//	//targetUrl := "https://www.abuyun.com/switch-ip"
//	//targetUrl := "https://www.abuyun.com/current-ip"
//
//	// 初始化 proxy http client
//	client := AbuyunProxy{AppID: proxyUser, AppSecret: proxyPass}.ProxyClient()
//
//	request, _ := http.NewRequest("GET", targetUrl, bytes.NewBuffer([]byte(``)))
//
//	response, err := client.Do(request)
//
//	if err != nil {
//		panic("failed to connect: " + err.Error())
//	} else {
//		bodyByte, err := ioutil.ReadAll(response.Body)
//		if err != nil {
//			fmt.Println("读取 Body 时出错", err)
//			return
//		}
//		response.Body.Close()
//
//		body := string(bodyByte)
//
//		fmt.Println("Response Status:", response.Status)
//		fmt.Println("Response Header:", response.Header)
//		fmt.Println("Response Body:\n", body)
//	}
//}
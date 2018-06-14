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

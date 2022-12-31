package service

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strings"
	"time"
)

func LoginRequest(username string, password string) string {
	resp, err := soup.Get("https://account.ccnu.edu.cn/cas/login?service=http://kjyy.ccnu.edu.cn/loginall.aspx?page=")
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	links1 := doc.Find("body", "id", "cas").FindAll("script")
	js := links1[2].Attrs()["src"]
	links2 := doc.Find("div", "class", "logo").FindAll("input")
	st := links2[2].Attrs()["value"]
	jar, _ := cookiejar.New(&cookiejar.Options{})
	client := http.Client{
		Jar:     jar, //初始化cookie容器
		Timeout: 5 * time.Second,
	}
	s3 := js[26:]
	s4 := st
	url1 := fmt.Sprintf("https://account.ccnu.edu.cn/cas/login;jsessionid=%v?service=http://kjyy.ccnu.edu.cn/loginall.aspx?page=", s3)    //登录url
	text := fmt.Sprintf("username=%v&password=%v&lt=%v&execution=e1s1&_eventId=submit&submit=%E7%99%BB%E5%BD%95", username, password, s4) //登录的body
	body := strings.NewReader(text)
	req, _ := http.NewRequest("POST", url1, body)
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Length", "162")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	s3 = "JSESSIONID=" + s3
	req.Header.Set("Cookie", s3)
	req.Header.Set("Host", "account.ccnu.edu.cn")
	req.Header.Set("Origin", "https://account.ccnu.edu.cn")
	req.Header.Set("Referer", "https://account.ccnu.edu.cn/cas/login?service=http://kjyy.ccnu.edu.cn/loginall.aspx?page=")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:107.0) Gecko/20100101 Firefox/107.0")
	req.Header.Set("sec-ch-ua", "")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", "Windows")
	res, err := client.Do(req) //执行登录请求
	return res.Header.Get("Content-Length")
}

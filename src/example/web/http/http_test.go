package http

import (
	"fmt"
	"net/http"
	url2 "net/url"
	"strings"
	"testing"
)

func TestHttpGet(t *testing.T) {

	client := http.Client{}

	url := "https://www.toutiao.com/search/suggest/initial_page/"
	request, err := http.NewRequest("GET", url, nil)
	CheckErr(err)
	// cookies
	cookies := &http.Cookie{Name: "username", Value: "Steven"}
	request.AddCookie(cookies)
	res, err := client.Do(request)
	CheckErr(err)
	request.Header.Set("Accept-Language", "zh-cn")
	defer res.Body.Close()
	fmt.Printf("Header:%+v\n", request.Header)
	fmt.Printf("响应状态码:%v\n", res.StatusCode)
	if res.StatusCode == 200 {
		fmt.Println("网络请求成功")
		CheckErr(err)
	} else {
		fmt.Println("网络请求失败", res.Status)
	}
}

func CheckErr(err error) {
	defer func() {
		if ins, ok := recover().(error); ok {
			fmt.Println("程序出现异常:", ins.Error())
		}
	}()

	if err != nil {
		panic(err)
	}
}
func TestPost(t *testing.T) {

	client := http.Client{}
	data := url2.Values{
		"theCityName": {"重庆"},
	}
	// 参数转换成body
	reader := strings.NewReader(data.Encode())
	url := "http://www.webxml.com.cn/WebService.asmx/getWeatherbyCityName"
	contentType := "application/x-www-form-urlencoded"
	response, err := client.Post(url, contentType, reader)
	CheckErr(err)
	fmt.Printf("响应状态码:%v \n", response.StatusCode)
	if response.StatusCode == 200 {
		defer response.Body.Close()
		fmt.Println("网络请求成功")
		CheckErr(err)
	} else {
		fmt.Println("请求失败", response.Status)
	}
}

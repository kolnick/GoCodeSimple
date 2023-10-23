package staticweb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {

	http.ListenAndServe(":2003", http.FileServer(http.Dir("./public/")))
}

func TestHttpServer(t *testing.T) {
	// 绑定请求路径触发回调方法
	http.HandleFunc("/index", indexHandler)
	// 如果第二个参数为nil 表示使用http.DefaultServeMux进行处理
	err := http.ListenAndServe(":2003", nil)
	fmt.Println(err)

}

func indexHandler(we http.ResponseWriter, r *http.Request) {
	fmt.Println("/index--------")
	// 向浏览器输出
	we.Write([]byte("这是默认首页"))
}

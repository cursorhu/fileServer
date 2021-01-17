//main.go

package main
import (
	"net/http"
	"fmt"
	"fileServer/Day1/1-httpGet/handler" //这个路径是相对GOPATH
)

func main() {
	//HandleFunc用于注册资源路径对应的处理函数，当http请求的资源路径是与注册路径匹配，回调处理函数
	//资源路径，即浏览器访问服务端的地址: <ip>:<port><resource>中的resource，默认路径是'/' 
	http.HandleFunc("/file/upload", handler.UploadHandler)
	
	//开始监听请求
	//ListenAndServe参数分别是监听端口号，和创建监听的自定义方法，nil是调用http包内的默认方法
	//ListenAndServe会获取HandleFunc注册的方法和资源路径，如果请求匹配资源路径，就调用对应的方法。
	err := http.ListenAndServe(":9008", nil)
	if err != nil {
		fmt.Printf("Failed to start server, err: %s\n", err.Error())
	}
}
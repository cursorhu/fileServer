//handler.go
package handler

import (
	"net/http"
	"io"
	"io/ioutil"
)

//UploadHandler: 处理文件上传请求
//该方法要在main包使用，必须首字母大写定义成其他包可见。
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	//http.Request请求类型是GET:
	//GET请求一般是客户端获取服务端的主页面，读取服务端数据等场景，可以抽象成客户端“读”服务端的资源
	if r.Method == "GET" {
		//返回html主页面的数据,浏览器会显示主页面
		//ioutil.ReadFile方法读取文件，返回文件内容到data, 文件路径是server程序的相对路径
		data, err := ioutil.ReadFile("./static/view/index.html")
		//错误处理
		if err != nil {
			//io.WriteString写字符串到http.ResponseWriter类型的对象, 作为http请求的返回数据，发送给客户端
			io.WriteString(w, "internal server error")
			//直接返回
			  return
		}
		//如果读取文件正常，就写数据(字符串)到返回消息体
		io.WriteString(w, string(data))
	}
	//http.Request请求类型是POST:
	//POST可实现客户端对服务端的资源的增，删，改等操作，可以抽象成客户端“写”服务端的资源
	if r.Method == "POST" {
		//接收文件流，存储到本地
	}
}
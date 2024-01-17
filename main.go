package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func index(w http.ResponseWriter, r *http.Request) {
	var indexData IndexData
	indexData.Title = "lincoxi blog"
	indexData.Desc = "first of all"
	t := template.New("index.html")
	path, _ := os.Getwd()
	fmt.Println(path)
	//解析所有模版
	home := path + "/template/home.html"
	header := path + "/template/layout/header.html"
	footer := path + "/template/layout/footer.html"
	personal := path + "/template/layout/personal.html"
	post := path + "/template/layout/post-list.html"
	pagination := path + "/template/layout/pagination.html"
	t, _ = t.ParseFiles(path+"/template/index.html", home, header, footer, personal, post, pagination) //路径设置要正确
	//页面涉及数据，需要有定义
	t.Execute(w, indexData)
}
func main() {
	//web应用
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", index)

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}

}

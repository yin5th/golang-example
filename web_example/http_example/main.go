package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type UserInfo struct {
	Name string
	Age  int
	Sex  string
}

func main() {
	//登录
	http.HandleFunc("/login", login)
	//用户信息
	http.HandleFunc("/user", userInfo)
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		fmt.Printf("http listen failed, err:%v\n", err)
		return
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		t, err := template.ParseFiles("templates/user/login.gtpl")
		if err != nil {
			log.Fatalf("template not found, err: %v", err)
		}
		//渲染到页面
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		form := r.FormValue

		fmt.Fprintf(w, "welcom %v\n", form("username"))
	} else {
		fmt.Fprintf(w, "method error")
	}
}

func userInfo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		/*userInfo := UserInfo{
			Name: "yin5th",
			Age:  18,
			Sex:  "男",
		}*/
		userInfo := make(map[string]interface{})
		userInfo["name"] = "yin5th"
		userInfo["age"] = 15
		userInfo["sex"] = "男"
		t, err := template.ParseFiles("templates/user/info.gtpl")
		if err != nil {
			log.Fatalf("template not found, err: %v", err)
		}
		//渲染到页面
		t.Execute(w, userInfo)
		//渲染到终端
		t.Execute(os.Stdout, userInfo)
	} else if r.Method == "POST" {
		fmt.Fprintf(w, "user edit success\n")
	} else {
		fmt.Fprintf(w, "method error")
	}
}

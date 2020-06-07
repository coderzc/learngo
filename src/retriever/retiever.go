package main
/*
接口
 */

import (
	"fmt"
	"learngo/src/retriever/http"
	"learngo/src/retriever/mock"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string) string
}

func download(r Retriever) string {
	return r.Get("https://wwww.baidu.com")
}

func MyPrint(a interface{}) interface{}{
	fmt.Println(a.(string))
	return a
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Get("https:/www.baidu.com")
	return s.Post("https:/www.baidu.com")
}

func main() {
	var r Retriever = mock.Retriever{Contents: "this is a fake"}
	fmt.Println(download(r))

	var r2 Retriever = &http.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	//fmt.Println(download(r2))

	fmt.Printf("%T,%v\n", r2, r2)

	// type switch
	switch v := r2.(type) {
	case mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *http.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}

	// Type assertion 把接口装换成类型(T)
	if retriever, ok := r.(http.Retriever); ok {
		fmt.Println(retriever)
	}else {
		fmt.Println("not a mock retriever")
	}

	s := session(mock.Retriever{Contents: "this a mock.Retriever"})
	fmt.Println(s)

}

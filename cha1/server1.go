package main

import (
	"fmt"
	"log"
	"net/http"
)

var count int

/**
指针是可见的内存地址，&操作符可以返回一个变量的内存地址，
并且*操作符可以获取指针指向的变量内容，
但是在Go语言里没有指针运算，也就是不能像c语言里可以对指针进行加或减操作。
*/
func hander1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "the content is %q \n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	count++
	fmt.Fprintf(w, "the count is %d\n", count)
}

func incr(p *int) int {
	*p++ // 非常重要：只是增加p指向的变量的值，并不改变p指针！！！
	return *p
}

func counter1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "the count is %d\n", incr(&count))
}

func main() {
	http.HandleFunc("/he", hander1)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/count1", counter1)
	log.Fatal(http.ListenAndServe("localhost:9090", nil))
}

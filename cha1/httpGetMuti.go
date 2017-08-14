package main

import (
	"time"
	"net/http"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s  %s", secs, nbytes, url, b) //send content to ch
}

/**
   	goroutine尝试在一个channel上做send or receive操作时 goroutine会阻塞在调用处，
   	直到另一个goroutine往这个channel里写入、或者接收了值，这样两个goroutine才会继续执行操作channel完成之后的逻辑；

   	在这个例子中，每一个fetch函数在执行时都会往channel里发送一个值(ch <- expression)，
   	主函数接收这些值(<-ch)。这个程序中我们用main函数来接收所有fetch函数传回的字符串，可以避免在goroutine异步执行时同时结束
 */

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
		//goroutine是一种函数的并发执行方式， channel用来在goroutine之间进行参数传递；
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
/**
0.02s   111291  http://www.baidu.com
0.09s    94094  http://www.sina.cn
2.24s    55428  https://github.com
2.24s elapsed
 */
//todo add notes for channel

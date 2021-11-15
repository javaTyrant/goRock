package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	//创建一个字符串通道
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		//启动一个新的goroutine
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		//从ch里取数据.
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
	}
	secs := time.Since(start).Seconds()
	//写入到ch里.
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

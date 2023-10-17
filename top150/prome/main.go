package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var urlPath = flag.String("url", "http://10.119.104.12:9100/metrics", "请求url")
var cycle = flag.Int("cycle", 60, "压测周期次数")
var concurrency = flag.Int("number", 0, "压测并发数")
var interval = flag.Int("interval", 10, "压测请求并发间隔")

func main() {
	flag.Parse()

	//  10.119.104.10:9100
	//  10.119.104.8:9100
	start := time.Now()
	timer := time.Duration(*interval) * time.Second //并发周期
	for i := 0; i < *cycle; i++ {

		//result := collect(*urlPath, *concurrency, i+1)
		go collect(*urlPath, *concurrency, i+1)
		//fmt.Println(result)
		time.Sleep(timer)
	}
	since := time.Now().Sub(start)
	fmt.Println("压测执行结束，总耗时：", since)

}

func httpGet(url string, response chan string, limiter chan bool, wg *sync.WaitGroup) {

	// 函数执行完毕时 计数器减 -1
	defer wg.Done()
	client := http.DefaultClient
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(" create request failed.....")
		return
	}
	request.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(" send request failed,: ", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read http response failed:", err)
		return
	}
	byteSize := len(string(body))
	// 计算字符串所占的KB数
	kb := float64(byteSize) / 1024.0
	// 计算字符串所占的MB数
	mb := float64(byteSize) / (1024.0 * 1024.0)
	if resp.StatusCode == 200 {
		// 将拿到的结果，发送到参数中传递过来的channel中
		response <- fmt.Sprintf(" http get request success. request url: %s, code: %d ,resp bytes  %.4f KB,  %.4f MB \n", url, resp.StatusCode, kb, mb)
	} else {
		response <- fmt.Sprintf(" http get request failed. request url: %s, code: %d ,resp bytes  %.4f KB,  %.4f MB \n", url, resp.StatusCode, kb, mb)
	}
	// 释放一个坑位
	<-limiter
}

func collect(urlPath string, number, requestNumber int) {

	var result []string
	wg := &sync.WaitGroup{}

	// 控制并发数为10
	limiter := make(chan bool, number)
	defer close(limiter)

	// 函数内的局部变量channel, 专门用来接收函数内所有goroutine的结果
	responseChannel := make(chan string, number)

	// 为读取结果控制器创建新的WaitGroup, 需要保证控制器内的所有值都已经正确处理完毕，才能结束
	wgResponse := &sync.WaitGroup{}
	// 启动读取结果的控制器

	go func() {
		// wgResponse 计数器+1
		wgResponse.Add(1)
		// 读取结果
		for response := range responseChannel {
			// 处理结果
			result = append(result, response)
		}
		// 当 responseChannel被关闭时且channel中所有的值都已经被处理完毕后，将执行到这一行
		fmt.Println("   ---> ")
		wgResponse.Done()
	}()

	for i := 0; i < number; i++ {
		//计数器+1
		wg.Add(1)
		limiter <- true
		// 这里再启动goroutine时，将用来收集结果的局部变量channel也传递进去
		go httpGet(urlPath, responseChannel, limiter, wg)
	}
	// 等待所以协程执行完毕
	wg.Wait() // 当计数器为0时，不再阻塞
	fmt.Printf("执行压测周期次数：%d ,并发请求数：%d ", requestNumber, number)
	// 关闭接收结果channel

	close(responseChannel)
	// 等待wgResponse的计数器归零
	wgResponse.Wait()
	// 返回聚合后结果
	fmt.Println(result)
	//return result
}

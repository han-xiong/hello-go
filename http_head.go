package main

import (
	"fmt"
	"net"
	"net/http"
	time2 "time"
)

var url = []string{
	"https://www.baidu.com",
	"https://google.com",
	"https://taobao.com",
}

func main() {

	for _, v := range url {
		/**
		自定义客户端, 实现超时
		 */
		c := http.Client{
			Transport: &http.Transport{
				Dial: func(network, addr string) (net.Conn, error) {
					timeout := time2.Second * 2
					return net.DialTimeout(network, addr, timeout)
				},
			},
		}

		// resp,err:=http.Head(v)

		//使用自定义的http
		resp, err := c.Head(v)
		if err != nil {
			fmt.Printf("head %s failed, err:%v\n", v, err)
			continue
		}

		fmt.Printf("head success, status: %v\n", resp.Status)
	}
}

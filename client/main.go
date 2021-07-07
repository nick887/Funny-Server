package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

var (
	done = make(chan string)
	wg sync.WaitGroup
)

func main() {
	// 1、与服务端建立连接
	conn, err := net.Dial("tcp", "139.224.239.181:1037")
	if err != nil {
		fmt.Printf("conn server failed, err:%v\n", err)
		return
	}
	// 2、使用 conn 连接进行数据的发送和接收
	wg.Add(2)
	go handleWrite(conn, done)
	go handleRead(conn, done)
	wg.Wait()
}

func handleWrite(conn net.Conn, done chan string) {
	inputReader := bufio.NewReader(os.Stdin)
	for {
		in,err:=inputReader.ReadBytes('\n')
		if err != nil {
			fmt.Println(err)
		}
		_,err=conn.Write(in)
		if err != nil {
			fmt.Println(err)
		}
	}
	wg.Done()
}
func handleRead(conn net.Conn, done chan string) {
	for {
		time.Sleep(200*time.Millisecond)
		buf := make([]byte, 1024)
		reqLen, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error to read message because of ", err)
			return
		}
		fmt.Println(string(buf[:reqLen-1]))
	}
	wg.Done()
}

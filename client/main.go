package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

var (
	wg sync.WaitGroup
	who string
)

func main() {
	// 1、与服务端建立连接
	conn, err := net.Dial("tcp", "localhost:1037")
	if err != nil {
		fmt.Printf("conn server failed, err:%v\n", err)
		return
	}
	// 2、使用 conn 连接进行数据的发送和接收
	getName(conn)
	wg.Add(2)
	go handleWrite(conn)
	go handleRead(conn)
	wg.Wait()
}

func handleWrite(conn net.Conn) {
	inputReader := bufio.NewReader(os.Stdin)
	for {
		in,err:=inputReader.ReadBytes('\n')
		if err != nil {
			fmt.Println(err)
		}
		_,err=conn.Write(in)
		fmt.Printf("\x1bM")
		if err != nil {
			fmt.Println(err)
		}
	}
	wg.Done()
}
func handleRead(conn net.Conn) {
	for {
		//time.Sleep(200*time.Millisecond)
		buf := make([]byte, 1024)
		reqLen, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error to read message because of ", err)
			return
		}
		fmt.Printf("\n\x1bM")
		fmt.Printf(string(buf[:reqLen-1])+"\n")
		fmt.Printf(who+" : ")
	}
	wg.Done()
}

func getName(conn net.Conn)  {
	buf := make([]byte, 1024)
	reqLen,err:=conn.Read(buf)
	if err != nil {
		fmt.Println("Error to read message because of ", err)
		return
	}
	t:=strings.Split(string(buf[:reqLen-1])," ")
	who=t[len(t)-1]
}
package handler

import (
	"FunnyServer/client_writer"
	"FunnyServer/global"
	"FunnyServer/util"
	"bufio"
	"fmt"
	"net"
)

func  HandleConn(conn net.Conn)  {
	ch :=make(chan string)
	go client_writer.ClientWriter(conn,ch)

	who :=util.GenerateAliasByReplyIndexAndHoleId(util.GenerateIntFromIp(conn.RemoteAddr().String()),1037)
	fmt.Fprintln(conn,"You are "+who)
	res ,err:= global.RedisClient.LRange("chat",-100,-1).Result()
	if err != nil {
		fmt.Println(err)
	}else{
		for _,msg := range res{
			ch <- msg
		}
	}
	global.Messages <- who+" has arrived"
	global.Entering <- ch

	input :=bufio.NewScanner(conn)
	for input.Scan() {
		global.Messages <- who + " : "+input.Text()
	}

	global.Leaving <- ch
	global.Messages <- who+" has left"
	conn.Close()
}
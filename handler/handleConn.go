package handler

import (
	"FunnyServer/client_writer"
	"FunnyServer/consts"
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
	fmt.Fprintln(conn,"你是 <"+who+">")
	res ,err:= global.RedisClient.LRange("chat",-100,-1).Result()
	if err != nil {
		fmt.Println(err)
	}else{
		for _,msg := range res{
			ch <- msg
		}
	}
	global.Messages <- who+consts.COME_SUFFIX
	global.Entering <- ch

	input :=bufio.NewScanner(conn)
	for input.Scan() {
		if len(input.Text())>8&&(input.Text()[:7] == "declare") {
			who=input.Text()[8:]
		}
		global.Messages <- who + " : "+input.Text()
	}

	global.Leaving <- ch
	global.Messages <- who+consts.LEAVE_SUFFIX
	conn.Close()
}
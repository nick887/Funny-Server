package handler

import (
	"FunnyServer/client_writer"
	"FunnyServer/global"
	"FunnyServer/util"
	"bufio"
	"net"
)

func  HandleConn(conn net.Conn)  {
	ch :=make(chan string)
	go client_writer.ClientWriter(conn,ch)

	who :=util.GenerateAliasByReplyIndexAndHoleId(uint(len(global.Clients)),1037)
	ch <-"You are "+who
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
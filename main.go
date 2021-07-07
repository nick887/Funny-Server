package main

import (
	"FunnyServer/util"
	"bufio"
"fmt"
"log"
"net"
)

func main() {
	listener,err:=net.Listen("tcp","47.116.139.54:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for  {
		conn,err:=listener.Accept()
		if err!=nil{
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client chan<- string

var (
	entering = make(chan client)
	leaving = make(chan client)
	messages = make(chan string)
	clients = make(map[client]bool)
)

func broadcaster()  {
	for  {
		select {
		case msg:=<-messages:
			for cli:=range clients{
				cli<-msg
			}
		case cli:=<-entering:
			clients[cli] =true
		case cli:=<-leaving:
			delete(clients,cli)
			close(cli)
		}
	}
}

func  handleConn(conn net.Conn)  {
	ch :=make(chan string)
	go clientWriter(conn,ch)

	who :=util.GenerateAliasByReplyIndexAndHoleId(uint(len(clients)),1037)
	ch <-"You are "+who
	messages <- who+" has arrived"
	entering <- ch

	input :=bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + " : "+input.Text()
	}

	leaving <- ch
	messages <- who+" has left"
	conn.Close()
}

func clientWriter(conn net.Conn,ch <-chan string)  {
	for msg:=range ch{
		fmt.Fprintln(conn,msg)
	}
}

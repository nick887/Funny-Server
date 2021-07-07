package client_writer

import (
	"fmt"
	"net"
)

func ClientWriter(conn net.Conn,ch <-chan string)  {
	for msg:=range ch{
		fmt.Fprintln(conn,msg)
	}
}

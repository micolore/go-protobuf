package main

import (
	"fmt"
	"net"

	"github.com/golang/protobuf/proto"
	"kubrick.com/protobuf"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp4", "127.0.0.1:9000")
	listener, _ := net.ListenTCP("tcp4", addr)
	for {
		conn, _ := listener.AcceptTCP()
		go func() {
			for {
				buf := make([]byte, 512)
				_, _ = conn.Read(buf)
				newData := &protobuf.DataMessage{}
				_ = proto.Unmarshal(buf, newData)
				fmt.Println("server receive : ", newData.PointSendMssage.MsgId)
			}
		}()
	}
}

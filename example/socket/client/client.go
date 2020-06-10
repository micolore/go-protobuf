package main

import (
	"fmt"
	"net"
	"time"

	"github.com/golang/protobuf/proto"
	"kubrick.com/protobuf"
)

func main() {
	connection, _ := net.Dial("tcp", "127.0.0.1:9000")
	var targetID int32 = 1
	for {
		oldData := &protobuf.PointSendMssage{
			MsgId:    "1",
			MsgType:  "2",
			MsgData:  []byte("kubrick  socket protobuf "),
			SendUid:  "3",
			MsgStype: "4",
			MsgStime: "5",
		}

		sendData := &protobuf.DataMessage{
			MsgType:         "233",
			PointSendMssage: oldData,
		}
		data, _ := proto.Marshal(sendData)
		_, _ = connection.Write(data)
		fmt.Println("client send : ", data)
		time.Sleep(2 * time.Second)
		targetID++
	}
}

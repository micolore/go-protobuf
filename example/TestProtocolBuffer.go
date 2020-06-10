package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"kubrick.com/protobuf"
)

// TestProtocolBuffer ...
func main() {
	// SendMssage是models.pb.go的结构体
	oldData := &protobuf.SendMssage{
		MsgId:    "1",
		MsgType:  "2",
		MsgData:  []byte("kubrick protobuf"),
		SendUid:  "3",
		MsgStype: "4",
		MsgStime: "5",
	}

	data, err := proto.Marshal(oldData)
	if err != nil {
		fmt.Println("marshal error: ", err.Error())
	}
	fmt.Println("marshal data : ", data)

	newData := &protobuf.DataMessage{}
	err = proto.Unmarshal(data, newData)
	if err != nil {
		fmt.Println("unmarshal err:", err)
	}
	fmt.Println("unmarshal data : ", newData)

}

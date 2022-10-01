package tests

import (
	"go_client/server"
	"go_client/utils"
	"net"
	"testing"
	"time"

	"google.golang.org/protobuf/proto"
)

func TestPb(t *testing.T) {
	conn, err := net.DialTimeout("tcp", "127.0.0.1:9898", 5*time.Second)

	if err != nil {
		t.Error(err)
	}

	for i := 0; i < 1; i++ {
		sayHello := &server.SayReq{
			Text: "fff",
		}

		marshel, err := proto.Marshal(sayHello)

		if err != nil {
			t.Error(err)
		}

		var lenth int = len(marshel)

		var lenBytes = utils.IntToBytes(lenth)
		var msgBytes = utils.IntToBytes(1)
		var lenInt = len(lenBytes)
		var msgInt = len(msgBytes)

		var tempBuf = make([]byte, lenth+lenInt+msgInt)

		copy(tempBuf, lenBytes[0:lenInt])
		copy(tempBuf[lenInt:msgInt+lenInt], msgBytes)
		copy(tempBuf[lenInt+msgInt:], marshel)

		conn.Write(tempBuf)
	}

	defer conn.Close()
}

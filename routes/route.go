package routes

import (
	"go_client/network"
	"go_client/server"
	"go_client/utils"
	"net"

	"google.golang.org/protobuf/proto"
)

func InitRoutes() error {
	err := network.NetworkMgrInstance().RegisterMessageCallback(processHeartbeatFunction, 1)

	if err != nil {
		println(err)
	}

	return nil
}

func processHeartbeatFunction(conn net.Conn, b []byte) error {
	sayReq := &server.SayReq{}

	err := proto.Unmarshal(b, sayReq)

	if err != nil {
		return err
	}

	println("服务器给我说:%v", sayReq.Text)

	sayHello := &server.SayReq{
		Text: "Yes!I'm Alive",
	}

	marshel, err := proto.Marshal(sayHello)

	if err != nil {
		println(err)
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

	return nil
}

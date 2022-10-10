package routes

import (
	"errors"
	"go_client/network"
	"go_client/server"
	"net"

	"google.golang.org/protobuf/proto"
)

func InitRoutes() error {
	err := network.NetworkMgrInstance().RegisterMessageCallback(processHeartbeatFunction, 1)

	if err != nil {
		println(err)
	}

	err = network.NetworkMgrInstance().RegisterMessageCallback(processRegisterMessage, 2)

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
		Text: "Yes!I'm Alive,But Im robot",
	}

	marshel, err := proto.Marshal(sayHello)

	if err != nil {
		return err
	}

	err = network.NetworkMgrInstance().SendToClient(1, marshel)

	if err != nil {
		return err
	}

	return nil
}

func processRegisterMessage(conn net.Conn, b []byte) error {
	resultReply := &server.RegisterReply{}

	err := proto.Unmarshal(b, resultReply)

	if err != nil {
		return err
	}

	if resultReply.Result == server.RegisterResult_REG_DUPLICATE {
		println(err)
		return errors.New("注册失败")
	}

	println("注册成功")

	return nil
}

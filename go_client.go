package main

import (
	"go_client/network"
	"go_client/routes"
	"go_client/server"
	"go_client/signals"

	"google.golang.org/protobuf/proto"
)

func main() {
	println("启动客户端连接到服务器")

	err := routes.InitRoutes()

	if err != nil {
		println(err)
	}

	err = network.NetworkMgrInstance().StartServer("", "")

	if err != nil {
		println(err)
	}

	err = processLogic()

	if err != nil {
		panic(err)
	}

	select {
	case <-signals.ExitSig:
		break
	}
}

func processLogic() error {
	registerReq := &server.RegisterReq{
		Name: "Robot",
	}

	marshel, err := proto.Marshal(registerReq)

	if err != nil {
		return err
	}

	err = network.NetworkMgrInstance().SendToClient(2, marshel)

	if err != nil {
		return err
	}

	return nil
}

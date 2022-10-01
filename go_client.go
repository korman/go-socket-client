package main

import (
	"go_client/network"
	"go_client/routes"
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
}

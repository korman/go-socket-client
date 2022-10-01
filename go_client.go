package main

import "go_client/network"

func main() {
	println("启动客户端连接到服务器")

	err := network.NetworkMgrInstance().StartServer("", "")

	if err != nil {
		println(err)
	}
}

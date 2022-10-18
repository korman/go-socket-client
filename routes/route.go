package routes

import (
	"errors"
	"go_client/gate"
	"go_client/network"
	"net"
	"time"

	"google.golang.org/protobuf/proto"
)

var inputText string = "Hello World!"

func InitRoutes() error {
	err := network.NetworkMgrInstance().RegisterMessageCallback(processHeartbeatFunction, 1)

	if err != nil {
		println(err)
	}

	err = network.NetworkMgrInstance().RegisterMessageCallback(processRegisterMessage, 2)

	if err != nil {
		println(err)
	}

	err = network.NetworkMgrInstance().RegisterMessageCallback(processInitMapMessage, 11)

	if err != nil {
		println(err)
	}

	err = network.NetworkMgrInstance().RegisterMessageCallback(processLockNodeReplyMessage, 3)

	if err != nil {
		println(err)
	}

	err = network.NetworkMgrInstance().RegisterMessageCallback(processInputTextReplyMessage, 4)

	if err != nil {
		println(err)
	}

	return nil
}

func processHeartbeatFunction(conn net.Conn, b []byte) error {
	sayReq := &gate.SayReq{}

	err := proto.Unmarshal(b, sayReq)

	if err != nil {
		return err
	}

	println("服务器给我说:%v", sayReq.Text)

	sayHello := &gate.SayReq{
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
	resultReply := &gate.RegisterReply{}

	err := proto.Unmarshal(b, resultReply)

	if err != nil {
		return err
	}

	if resultReply.Result == gate.RegisterResult_REG_DUPLICATE {
		println(err)
		return errors.New("注册失败")
	}

	println("注册成功")

	return nil
}

func processInitMapMessage(conn net.Conn, b []byte) error {
	initMapInfo := new(gate.InitMapInfo)

	err := proto.Unmarshal(b, initMapInfo)

	if err != nil {
		return err
	}

	lockNode := &gate.Node{
		X:     1,
		Y:     1,
		State: gate.LockStatus_LOCKED_NODE,
	}

	var lockReq *gate.LockNodeReq = &gate.LockNodeReq{
		LockNode: lockNode,
	}

	sendData, err := proto.Marshal(lockReq)

	if err != nil {
		return err
	}

	return network.NetworkMgrInstance().SendToClient(3, sendData)
}

func processLockNodeReplyMessage(conn net.Conn, b []byte) error {
	lockResult := new(gate.LockNodeReply)

	err := proto.Unmarshal(b, lockResult)

	if err != nil {
		return err
	}

	if lockResult.Result != gate.LockResult_LOCK_SUCCEEDED {
		return errors.New("锁定节点失败")
	}

	time.Sleep(1 * time.Second)

	lockResult.LockedNode.Text = inputText[:1]

	inputMsg := new(gate.InputTextReq)

	inputMsg.InputNode = lockResult.LockedNode

	bt, err := proto.Marshal(inputMsg)

	if err != nil {
		return nil
	}

	return network.NetworkMgrInstance().SendToClient(4, bt)
}

func processInputTextReplyMessage(conn net.Conn, b []byte) error {
	inputResult := new(gate.InputTextReply)

	err := proto.Unmarshal(b, inputResult)

	if err != nil {
		return err
	}

	if inputResult.Result != gate.LockResult_LOCK_SUCCEEDED {
		return errors.New("输入失败")
	}

	if len(inputResult.InputedNode.Text) == len(inputText) {
		return nil
	}

	time.Sleep(1 * time.Second)

	inputResult.InputedNode.Text = inputText[:len(inputResult.InputedNode.Text)+1]

	inputMsg := new(gate.InputTextReq)

	inputMsg.InputNode = inputResult.InputedNode

	bt, err := proto.Marshal(inputMsg)

	if err != nil {
		return nil
	}

	return network.NetworkMgrInstance().SendToClient(4, bt)
}

package network

import (
	"errors"
	"fmt"
	"go_client/utils"
	"io"
	"net"
	"sync"
	"time"
)

var instanse *NetworkManager
var mutex sync.Mutex

type NetworkManager struct {
	readBytes   int
	tempBuffer  []byte
	callbackMap map[int64]func(net.Conn, []byte) error
	ClientConn  net.Conn
}

func NetworkMgrInstance() *NetworkManager {
	if instanse == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if instanse == nil {
			instanse = &NetworkManager{
				readBytes:   0,
				tempBuffer:  nil,
				callbackMap: make(map[int64]func(net.Conn, []byte) error),
				ClientConn:  nil,
			}
		}
	}
	return instanse
}

func (m *NetworkManager) StartServer(ip string, port string) error {
	println("连接到目标服务器……")

	conn, err := net.DialTimeout("tcp", "127.0.0.1:9898", 5*time.Second)

	if err != nil {
		return err
	}

	m.ClientConn = conn

	go m.connHandle(conn)

	return nil
}

func (m *NetworkManager) RegisterMessageCallback(f func(net.Conn, []byte) error, msgId int64) error {
	if f == nil {
		return errors.New("空的注册回调")
	}

	if msgId <= 0 {
		return errors.New("错误的注册号")
	}

	m.callbackMap[msgId] = f

	return nil
}

func (m *NetworkManager) SendToClient(msgId int, byDatas []byte) error {
	var lenth int = len(byDatas)

	var lenBytes = utils.IntToBytes(lenth)
	var msgBytes = utils.IntToBytes(msgId)
	var lenInt = len(lenBytes)
	var msgInt = len(msgBytes)

	var tempBuf = make([]byte, lenth+lenInt+msgInt)

	copy(tempBuf, lenBytes[0:lenInt])
	copy(tempBuf[lenInt:msgInt+lenInt], msgBytes)
	copy(tempBuf[lenInt+msgInt:], byDatas)

	_, err := m.ClientConn.Write(tempBuf)

	if err != nil {
		return err
	}

	return nil
}

func (m *NetworkManager) parseData(conn net.Conn, b []byte) ([]byte, error) {
	n := len(b)
	packLen := 0
	msgId := 0
	var err error = nil

	if n < 16 {
		return nil, errors.New("包太小")
	}

	if n >= 16 {
		packLenPack := make([]byte, 8)

		copy(packLenPack, b[:8])

		packLen = utils.BytesToInt(packLenPack)

		if packLen == 0 {
			return nil, errors.New("包大小为0")
		}

		msgIdPack := make([]byte, 8)
		copy(msgIdPack, b[8:16])

		msgId = utils.BytesToInt(msgIdPack)
		fmt.Printf("收到%d的消息", msgId)
	} else {
		err = errors.New("不够一个包")
		return b, err
	}

	if n >= packLen+16 {
		pack := make([]byte, packLen)

		copy(pack, b[16:16+packLen])

		if m.callbackMap[int64(msgId)] == nil {
			fmt.Printf("无msgId%d处理函数\n", msgId)
		} else {
			m.callbackMap[int64(msgId)](conn, pack)
		}

		if n-(packLen+16) > 0 {
			fmt.Printf("有溢出%d,此包大小%d\n", n-(packLen+16), packLen)
			err, tempBuf := m.parseData(conn, b[packLen+16:])

			if err != nil {
				return err, tempBuf
			}
		}
	} else {
		err = errors.New("不够一个包")
		return b, err
	}

	return nil, errors.New("正好")
}

func (m *NetworkManager) connHandle(conn net.Conn) {
	var buf []byte = make([]byte, 4000)

	for {
		n, err := conn.Read(buf)

		if err == io.EOF {
			println("失去连接,字节数%d", m.readBytes)
			break
		}

		var fullBuffer []byte = nil

		if m.tempBuffer == nil {
			fullBuffer = make([]byte, n)
			copy(fullBuffer, buf[0:n])
		} else {
			fullBuffer = make([]byte, n+len(m.tempBuffer))
			copy(fullBuffer, m.tempBuffer)
			copy(fullBuffer[len(m.tempBuffer):], buf[0:n])
		}

		if err != nil {
			return
		}

		tempBuf, err := m.parseData(conn, fullBuffer)

		if err != nil {
			if len(tempBuf) > 0 {
				m.tempBuffer = make([]byte, len(tempBuf))
				copy(m.tempBuffer, tempBuf)
			} else {
				m.tempBuffer = nil
			}
		} else {
			m.tempBuffer = nil
		}

		m.readBytes += n
	}
}

package network

import (
	"errors"
	"fmt"
	"go_client/server"
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
			}
		}
	}
	return instanse
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

func (m *NetworkManager) StartServer(ip string, port string) error {
	println("连接到目标服务器……")

	conn, err := net.DialTimeout("tcp", "127.0.0.1:9898", 5*time.Second)

	listener, err := net.Listen("tcp", ":9898")

	if err != nil {
		return err
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()

		if err != nil {
			println("连接失败")
			break
		}

		raddr := conn.RemoteAddr().String()
		println("有人连我了... %s", raddr)

		go m.connHandle(conn)
	}

	return nil
}

func (m *NetworkManager) parseData(conn net.Conn, b []byte, packs []*server.SayReq) ([]byte, error) {
	n := len(b)
	packLen := 0
	msgId := 0
	var err error = nil

	if n <= 16 {
		return nil, errors.New("包调小")
	}

	if n > 16 {
		packLenPack := make([]byte, 8)

		copy(packLenPack, b[:8])

		packLen = utils.BytesToInt(packLenPack)

		if packLen == 0 {
			return nil, errors.New("包大小为0")
		}

		msgIdPack := make([]byte, 8)
		copy(msgIdPack, b[8:16])

		msgId = utils.BytesToInt(msgIdPack)
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
			err, tempBuf := m.parseData(conn, b[packLen+16:], packs)

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

		packs := make([]*server.SayReq, 0)

		var fullBuffer []byte = nil

		if m.tempBuffer == nil {
			fullBuffer = make([]byte, n)
			copy(fullBuffer, buf[0:n])
		} else {
			println("有结余消息:%d", len(m.tempBuffer))
			fullBuffer = make([]byte, n+len(m.tempBuffer))
			copy(fullBuffer, m.tempBuffer)
			copy(fullBuffer[len(m.tempBuffer):], buf[0:n])
		}

		println("读取了%d字节", n)

		if err != nil {
			return
		}

		tempBuf, err := m.parseData(conn, fullBuffer, packs)

		if err != nil {
			println(err)

			println("剩下%d个字节", len(tempBuf))

			if len(tempBuf) > 0 {
				println("需要中间存储,余下%d字节", len(tempBuf))
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

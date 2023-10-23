package server

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int32

	// 在线用户列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	// 消息广播的channel
	Message chan string
}

func NewServer(ip string, port int32) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}
func (this *Server) Start() {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listener err:", err)
		return
	}
	defer listen.Close()

	// 启动监听Message的goroutine

	go this.ListenerMessage()

	for {
		// accept
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listener accept error", err)
			continue
		}
		// do handler
		go this.Handler(conn)
	}
}

func (this *Server) Handler(conn net.Conn) {
	fmt.Println("链接建立成功")
	user := newUser(conn, this)
	// 用户上线
	user.Online()

	// 监听用户是否活跃的channel
	isLive := make(chan bool)

	// 处理用户下线
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}

			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err:", err)
				return
			}
			// 提取用户的消息(去除'\n')
			msg := string(buf[:n-1])
			// 将得到的消息进行广播
			user.DoMessage(msg)

			// 有消息就是true
			isLive <- true
		}
	}()
	// 当前handler阻塞

	for {
		select {
		case <-isLive:
			// 活跃重置定时器 不做任何事情为了激活select 更新下面的定时器

		case <-time.After(time.Second * 10):
			// 已经超时
			// 将当前的User强制关闭
			user.sendMsg("你被踢了")

			//销毁资源
			close(user.Channel)
			// 关闭链接
			user.conn.Close()

			// 退出当前handler
			return
		}
	}
}

// 广播消息
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	this.Message <- sendMsg
}

// 监听Message广播消息channel的goroutine,一旦有消息就发送给全部的在线User
func (this *Server) ListenerMessage() {
	for {
		msg := <-this.Message
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.Channel <- msg
		}
		this.mapLock.Unlock()
	}
}

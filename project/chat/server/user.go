package server

import (
	"net"
	"strings"
)

type User struct {
	Name    string
	Addr    string
	Channel chan string
	conn    net.Conn
	server  *Server
}

// 创建一个用户的API

func newUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{Name: userAddr, Addr: userAddr, Channel: make(chan string), conn: conn, server: server}

	// 启动监听当前user channel的goroutine
	go user.ListenerMessage()
	return user
}

// 监听当前User channel的方法 一旦有消息 就直接发送给对端客户端

func (this *User) ListenerMessage() {
	for {
		msg := <-this.Channel
		this.conn.Write([]byte(msg + "\n"))
	}
}

func (this *User) Online() {
	// 用户上线了  将用户加入到onlineMap中
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()
	// 广播当前用户上线的消息
	this.server.BroadCast(this, "已上线")
}
func (this *User) Offline() {
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()
}

// 给当前User对应的客户端发送消息
func (this *User) sendMsg(msg string) {
	this.conn.Write([]byte(msg))
}
func (this *User) DoMessage(msg string) {

	if msg == "who" {
		// 查询当前在线用户列表
		this.server.mapLock.Lock()
		for _, user := range this.server.OnlineMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + ":" + "在线...\n"
			this.sendMsg(onlineMsg)
		}
		this.server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		// 消息格式: rename|张三
		newName := strings.Split(msg, "|")[1]
		// 判断name是否存在
		_, ok := this.server.OnlineMap[newName]
		if ok {
			this.sendMsg("当前用户名已被使用")
			return
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this
			this.server.mapLock.Unlock()
			this.Name = newName
			this.sendMsg("修改名称成功")
		}
	} else if len(msg) > 4 && msg[:3] == "to|" {
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			this.sendMsg("消息格式不正确,请使用 \"to|张三|你好啊\"格式.\n")
			return
		}
		remoteUser, ok := this.server.OnlineMap[remoteName]
		if !ok {
			this.sendMsg("该用户不存在")
			return
		}

		content := strings.Split(msg, "|")[2]
		if content == "" {
			this.sendMsg("无消息内容,请重发\n")
			return
		}
		remoteUser.sendMsg(this.Name + "对您说:" + content)

	} else {
		this.server.BroadCast(this, msg)
	}
}

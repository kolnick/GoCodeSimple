package websocket

import (
	"golang.org/x/net/websocket"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type CmdContextImpl struct {
	userId       int64
	clientIpAddr string
	Conn         *websocket.Conn
	sendMsgQ     chan protoreflect.ProtoMessage
}

func (ctx *CmdContextImpl) ClientIpAddr() string {
	return ctx.clientIpAddr
}

func (ctx *CmdContextImpl) SetClientIpAddr(clientIpAddr string) {
	ctx.clientIpAddr = clientIpAddr
}

func (ctx *CmdContextImpl) GetUserId() int64 {
	return ctx.GetUserId()
}

func (ctx *CmdContextImpl) BindUserId(val int64) {
	ctx.userId = val
}

func (ctx *CmdContextImpl) Disconnect() {
	if nil != ctx.Conn {
		_ = ctx.Conn.Close()
	}
}

func (ctx *CmdContextImpl) Write(msgObj protoreflect.ProtoMessage) {
	if nil == msgObj || nil == ctx.Conn || nil == ctx.sendMsgQ {
		return
	}
	ctx.sendMsgQ <- msgObj
}

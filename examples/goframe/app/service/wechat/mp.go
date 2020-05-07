package wechat

import (
	"fmt"
	"gf-app/boot"
	"github.com/gogf/gf/net/ghttp"
	"github.com/silenceper/wechat/message"
)

func MP(r *ghttp.Request) {
	// 接入验证
	if r.Method == "GET" {
		_ = r.ParseForm()
		r.Response.WriteStatus(200, r.FormValue("echostr"))
		return
	}

	//设置接收消息的处理方法
	MPServer := boot.WeChat.GetServer(r.Request, r.Response.Writer)
	MPServer.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
		// 判断不同的消息类型
		switch msg.MsgType {

		case message.MsgTypeText: //文本消息

		case message.MsgTypeImage: //图片消息

		case message.MsgTypeVoice: //语音消息

		case message.MsgTypeVideo: //视频消息

		case message.MsgTypeShortVideo: //小视频消息

		case message.MsgTypeLocation: //地理位置消息

		case message.MsgTypeLink: //链接消息

		case message.MsgTypeEvent: //事件推送消息

		}

		// 演示回复用户发送的消息
		text := message.NewText(msg.Content)
		return &message.Reply{
			MsgType: message.MsgTypeText,
			MsgData: text,
		}
	})

	//处理消息接收以及回复
	err := MPServer.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}

	//发送回复的消息
	_ = MPServer.Send()
}

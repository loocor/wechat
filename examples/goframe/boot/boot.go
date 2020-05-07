package boot

import (
	"github.com/gogf/gf/frame/g"
	"github.com/silenceper/wechat"
)

var (
	CONF     = g.Config()
	WCConfig = &wechat.Config{
		AppID:          CONF.GetString("wechat.AppID"),
		AppSecret:      CONF.GetString("wechat.AppSecret"),
		Token:          CONF.GetString("wechat.Token"),
		EncodingAESKey: CONF.GetString("wechat.EncodingAESKey"),
	}
	WeChat = wechat.NewWechat(WCConfig)
)

func init() {

}

package router

import (
	"gf-app/app/service/wechat"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	server := g.Server()
	server.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/", wechat.MP)
	})
}

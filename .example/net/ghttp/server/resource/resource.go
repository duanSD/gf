package main

import (
	"fmt"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gres"
	_ "github.com/gogf/gf/os/gres/testdata/data"
)

func main() {
	gres.Dump()

	v := g.View()
	v.SetPath("template/layout1")

	s := g.Server()
	s.SetIndexFolder(true)
	s.SetServerRoot("root")
	s.BindHookHandler("/*", ghttp.HOOK_BEFORE_SERVE, func(r *ghttp.Request) {
		fmt.Println(r.URL.Path, r.IsFileRequest())
	})
	s.BindHandler("/template", func(r *ghttp.Request) {
		r.Response.WriteTpl("layout.html")
	})
	s.SetPort(8199)
	s.Run()
}

package routers

import (
	"SshWebShell/controllers"
	"net/http"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	ignoreStaticPath()

	ns :=
		beego.NewNamespace("/api",
			//beego.NSBefore(auth),
			beego.NSNamespace("/v1.0",
				//登录注册
				beego.NSRouter("/reg", &controllers.UserController{}, "post:Reg"),
				beego.NSRouter("/login", &controllers.UserController{}, "post:Login"),
				//获取当前服务器状态
				beego.NSRouter("/system", &controllers.SystemController{}, "get:GetStat"),
				//管理进程
				beego.NSRouter("/process", &controllers.SystemController{}, "get:GetProcess;post:KillProcess"),
				//服务器管理
				beego.NSRouter("/server", &controllers.ServerController{}, "get:GetServer;post:ChangeServer"),
				//文件管理
				beego.NSRouter("/files", &controllers.FilesController{}, "get:GetFiles;post:GetTFiles"),
				beego.NSRouter("/upload", &controllers.FilesController{}, "post:FilesUpload"),
			),
		)

	//注册 namespace
	beego.AddNamespace(ns)
}

func ignoreStaticPath() {

	//透明static
	beego.InsertFilter("/", beego.BeforeRouter, TransparentStatic)
	beego.InsertFilter("/*", beego.BeforeRouter, TransparentStatic)
}

func TransparentStatic(ctx *context.Context) {
	orpath := ctx.Request.URL.Path
	beego.Debug("request url: ", orpath)
	//如果请求uri还有api字段,说明是指令应该取消静态资源路径重定向
	if strings.Index(orpath, "api") >= 0 {
		return
	}
	http.ServeFile(ctx.ResponseWriter, ctx.Request, "static/"+ctx.Request.URL.Path)

}

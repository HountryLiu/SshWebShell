package main

import (
	_ "SshWebShell/models"
	_ "SshWebShell/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

package controllers

import (
	"SshWebShell/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ServerController struct {
	beego.Controller
}

//将封装好的返回结构 变成json返回给前段
func (this *ServerController) RetData(res interface{}) {
	this.Data["json"] = res
	this.ServeJSON()
}

func (this *ServerController) GetServer() {

	o := orm.NewOrm()
	var server []models.Server

	qs := o.QueryTable("server")
	_, err := qs.All(&server)
	if err != nil {

		return
	}
	var res = server
	defer this.RetData(res)
	return
}

func (this *ServerController) ChangeServer() {
	res := make(map[string]interface{})
	res["errno"] = models.RECODE_OK
	res["errmsg"] = models.RecodeText(models.RECODE_OK)

	defer this.RetData(res)
	server := models.Server{}
	o := orm.NewOrm()

	op := this.GetString("oper")
	switch op {
	case "add":
		server.Ip = this.GetString("ip")
		server.Rusername = this.GetString("rusername")
		server.Rpassword = this.GetString("rpassword")
		server.Port, _ = this.GetInt("port")

		if _, err := o.Insert(&server); err != nil {
			beego.Info("insert error = ", err)
			return
		}
	case "del":
		server.Id, _ = this.GetInt("id")
		if _, err := o.Delete(&server); err != nil {
			beego.Info("delete error = ", err)
			return
		}
	case "edit":
		server.Id, _ = this.GetInt("id")
		if o.Read(&server) == nil {
			server.Ip = this.GetString("ip")
			server.Rusername = this.GetString("rusername")
			server.Rpassword = this.GetString("rpassword")
			server.Port, _ = this.GetInt("port")
			if _, err := o.Update(&server); err != nil {
				beego.Info("update error = ", err)
				return
			}
		}
	}

}

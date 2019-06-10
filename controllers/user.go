package controllers

import (
	"SshWebShell/models"
	"crypto/md5"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UserController struct {
	beego.Controller
}

//将封装好的返回结构 变成json返回给前段
func (this *UserController) RetData(res interface{}) {
	this.Data["json"] = res
	this.ServeJSON()
}

func (this *UserController) Reg() {

	//返回给前端的map结构体
	res := make(map[string]interface{})
	res["errno"] = models.RECODE_OK
	res["errmsg"] = models.RecodeText(models.RECODE_OK)

	defer this.RetData(res)

	var regRequestMap = make(map[string]interface{})

	//1 得到客户端请求的json数据 post数据
	json.Unmarshal(this.Ctx.Input.RequestBody, &regRequestMap)

	//2 判断数据的合法性
	if regRequestMap["username"] == "" || regRequestMap["password"] == "" {
		res["errno"] = models.RECODE_REQERR
		res["errmsg"] = models.RecodeText(models.RECODE_REQERR)
		return
	}

	//3 将数据存入到mysql数据库 user
	user := models.User{}
	//将password进行md5
	data := []byte(regRequestMap["password"].(string))
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	user.Password = md5str
	user.Username = regRequestMap["username"].(string)

	o := orm.NewOrm()

	_, err := o.Insert(&user)
	if err != nil {
		beego.Info("insert error = ", err)
		res["errno"] = models.RECODE_DATAEXIST
		res["errmsg"] = models.RecodeText(models.RECODE_DATAEXIST)
		return
	}

	//4 将当前的用户的信息存储到session中
	this.SetSession("username", user.Username)

	return

}

func (this *UserController) Login() {

	res := make(map[string]interface{})
	res["errno"] = models.RECODE_OK
	res["errmsg"] = models.RecodeText(models.RECODE_OK)
	defer this.RetData(res)

	var loginRequestMap = make(map[string]interface{})

	//1 得到客户端请求的json数据 post数据
	json.Unmarshal(this.Ctx.Input.RequestBody, &loginRequestMap)

	//2 判断数据的合法性
	if loginRequestMap["username"] == "" || loginRequestMap["password"] == "" {
		res["errno"] = models.RECODE_REQERR
		res["errmsg"] = models.RecodeText(models.RECODE_REQERR)
		return
	}

	//3 查询数据库得到user
	var user models.User

	o := orm.NewOrm()
	//select password from user where user.username = username
	qs := o.QueryTable("user")
	if err := qs.Filter("username", loginRequestMap["username"]).One(&user); err != nil {
		//查询失败
		res["errno"] = models.RECODE_USERERR
		res["errmsg"] = models.RecodeText(models.RECODE_USERERR)
		return
	}

	//4 对比密码 md5加密
	data := []byte(loginRequestMap["password"].(string))
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制

	if user.Password != md5str {
		res["errno"] = models.RECODE_PWDERR
		res["errmsg"] = models.RecodeText(models.RECODE_PWDERR)
		return
	}

	//5 将当前的用户的信息存储到session中
	this.SetSession("username", user.Username)

	return
}

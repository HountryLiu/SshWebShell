package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

/* 系统用户信息表 table_name = user */
type User struct {
	Id       int    `json:"id"`                              //用户编号
	Username string `orm:"size(32);unique"  json:"username"` //用户帐号
	Password string `orm:"size(128)" json:"password"`        //用户密码加密的
}

/* 远程云服务器信息表 table_name = server */
type Server struct {
	Id        int    `json:"id"`                        //远程云服务器编号
	Ip        string `orm:"size(15);unique" json:"ip"`  //远程云服务器ip
	Rusername string `orm:"size(32)"  json:"rusername"` //远程云服务器帐号
	Rpassword string `orm:"size(128)" json:"rpassword"` //远程云服务器密码
	Port      int    `json:"port"`                      //远程云服务器端口
}

func init() {

	// set default database
	orm.RegisterDataBase("default", "mysql", "root:168168@tcp(127.0.0.1:3306)/webshell?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(User), new(Server))

	// create table
	orm.RunSyncdb("default", false, true)

}

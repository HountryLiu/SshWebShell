package controllers

import (
	"SshWebShell/utils"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/pkg/sftp"
)

type FilesController struct {
	beego.Controller
}

//将封装好的返回结构 变成json返回给前段
func (this *FilesController) RetData(res interface{}) {
	this.Data["json"] = res
	this.ServeJSON()
}

//获取根目录文件
func (this *FilesController) GetFiles() {
	var res = make([]string, 0, 10)

	cmd := `ls /`

	sys_info := utils.ExecC(cmd)

	f_info := strings.Split(sys_info, "\n")
	f_info_tmp := make([]string, 0, 10)
	for _, v := range f_info[:len(f_info)-1] {
		f_info_tmp = strings.Fields(v)
		res = append(res, f_info_tmp[0])
	}
	defer this.RetData(res)
}

//获取目标目录文件
func (this *FilesController) GetTFiles() {
	var res = make([]string, 0, 10)
	path := this.GetString("path")
	cmd := `ls ` + path

	sys_info := utils.ExecC(cmd)

	f_info := strings.Split(sys_info, "\n")
	f_info_tmp := make([]string, 0, 10)
	for _, v := range f_info[:len(f_info)-1] {
		f_info_tmp = strings.Fields(v)
		res = append(res, f_info_tmp[0])
	}
	defer this.RetData(res)
}

//上传文件
func (this *FilesController) FilesUpload() {
	var (
		err        error
		sftpClient *sftp.Client
	)
	//本地路径文件夹
	localPath := `D:\ceshi`
	//远程路径
	remotePath := this.GetString("now_path")
	start := time.Now()
	sftpClient, err = utils.SftpConnect(utils.USERNAME, utils.PASSWORD, utils.HOST, utils.PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer sftpClient.Close()

	_, errStat := sftpClient.Stat(remotePath)
	if errStat != nil {
		log.Fatal(remotePath + " remote path not exists!")
	}

	_, err = ioutil.ReadDir(localPath)
	if err != nil {
		log.Fatal(localPath + " local path not exists!")
	}
	utils.UploadDirectory(sftpClient, localPath, remotePath)
	elapsed := time.Since(start)
	beego.Info("elapsed time : ", elapsed)
	this.Ctx.WriteString("Hello World!")

}

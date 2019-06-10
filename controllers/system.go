package controllers

import (
	"SshWebShell/utils"
	"strings"

	"github.com/astaxie/beego"
)

type SystemController struct {
	beego.Controller
}

type Process struct {
	Pid  string `json:"pid"`
	Name string `json:"name"`
}

//将封装好的返回结构 变成json返回给前段
func (this *SystemController) RetData(res interface{}) {
	this.Data["json"] = res
	this.ServeJSON()
}

func (this *SystemController) GetStat() {
	var res = make(map[string]string)

	defer this.RetData(res)
	//查看内存使用率
	cmd := `free -m | grep -E "^(Mem)"`
	sys_info := utils.ExecC(cmd)
	mem_info := strings.Fields(sys_info)
	res["mem_total"] = mem_info[1]
	res["mem_used"] = mem_info[2]

	//查看CPU使用率
	cmd = `top -b -n 1 | grep -E "^*(Cpu)"`
	sys_info = utils.ExecC(cmd)
	cpu_info := strings.Fields(sys_info)
	if cpu_info[7] == "id," {
		res["cpu_unused"] = "100"
	} else {
		res["cpu_unused"] = cpu_info[7]
	}

	//查看磁盘使用状况
	cmd = `df -lh | grep -E "^(/)"`
	sys_info = utils.ExecC(cmd)
	hd_info := strings.Fields(sys_info)
	hd_total := strings.Trim(hd_info[1], "G")
	hd_used := strings.Trim(hd_info[2], "G")
	res["hd_total"] = hd_total
	res["hd_used"] = hd_used

}

//查看进程	res["PID"] = 进程名
func (this *SystemController) GetProcess() {
	var res = make([]interface{}, 0, 10)

	var data = Process{}

	cmd := `ps -A`

	sys_info := utils.ExecC(cmd)

	p_info := strings.Split(sys_info, "\n")

	p_info_tmp := make([]string, 0, 10)

	for _, v := range p_info[1 : len(p_info)-1] {
		p_info_tmp = strings.Fields(v)
		data.Pid = p_info_tmp[0]
		data.Name = p_info_tmp[3]
		res = append(res, data)
	}

	defer this.RetData(res)
}

//结束指定进程
func (this *SystemController) KillProcess() {
	var res string
	pid := this.GetString("pid")

	cmd2 := `kill -9 ` + pid
	sys_info2 := utils.ExecC(cmd2)
	if sys_info2 == "" {
		res = "success"
	} else {
		res = "fail"
	}

	defer this.RetData(res)
}

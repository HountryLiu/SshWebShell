$(document).ready(function () {
	var tpl="";
	$.get("/api/v1.0/server", function(res) {
		$.each(res,function(i,val){
			tpl += `<li><a href="#" onclick=change(this) data-ip="`+val.ip+`" data-username="`+val.rusername+`" data-password="`+val.rpassword+`" data-port="`+val.port+`">`+val.ip+`</a></li>`
        });
		tpl += `<li><a href="login.html">安全退出</a>`;
		$(".server-ip").html(tpl);
	}, "json");

})

function change(obj){
	var now_ip = $(obj).attr("data-ip");
	var now_username = $(obj).attr("data-username");
	var now_password = $(obj).attr("data-password");
	var now_port = $(obj).attr("data-port");
	//localStorage.setItem("now_ip",now_ip);
	document.cookie="now_ip="+now_ip;
     $.ajax({
        type: "POST",
        dataType: "json",
        url: "/api/v1.0/server/choose" ,
		contentType: "application/json",
        data: jsonStr,
        success: function (result) {
            console.log(result);//打印服务端返回的数据(调试用)

        },
        error : function() {

        }
    });
	window.location.reload();
}
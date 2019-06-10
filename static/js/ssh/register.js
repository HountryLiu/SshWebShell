	$.fn.serializeObject = function()
	{
		var o = {};
		var a = this.serializeArray();
		$.each(a, function() {
		if (o[this.name] !== undefined) {
		    if (!o[this.name].push) {
		        o[this.name] = [o[this.name]];
		    }
		    o[this.name].push(this.value || '');
		} else {
		    o[this.name] = this.value || '';
		}
		});
		return o;
	};
	function reg() {

	var jsonObj = $("#reg_form").serializeObject();  //json对象			
	var jsonStr = JSON.stringify(jsonObj);  //json字符串
	
	if(jsonObj["password"] != jsonObj["password2"]){
		swal({
            title: "注册失败",
            text: "两次密码不一致，请重新输入！",
            type: "error"
        });
		return;
	}
		$.ajax({
                type: "POST",
                dataType: "json",
                url: "/api/v1.0/reg" ,
				contentType: "application/json",
                data: jsonStr,
                success: function (result) {
                    console.log(result);//打印服务端返回的数据(调试用)
                    if (result.errno == 0) {
                        swal({
		                    title: "注册成功",
		                    text: "欢迎使用 WebShell 系统",
		                    type: "success",		                 
		                }, function () {
		                    $(location).attr('href', 'index.html');
		                });
                    }else{
						swal({
		                    title: "注册失败",
		                    text: result.errmsg,
		                    type: "error"
		                });
					}
                },
                error : function() {
						swal({
		                    title: "网络异常，请稍后再试！",
		                    text: result.errmsg,
		                    type: "warning"
		                });
                }
            });
	}
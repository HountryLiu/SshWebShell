//写cookies
function setCookie(name,value)
{
	var Days = 30;
	var exp = new Date();
	exp.setTime(exp.getTime() + Days*24*60*60*1000);
	document.cookie = name + "="+ escape (value) + ";expires=" + exp.toGMTString();
}
//读cookies
function getCookie(name)
{
	var arr,reg=new RegExp("(^| )"+name+"=([^;]*)(;|$)");
	if(arr=document.cookie.match(reg))
	return unescape(arr[2]);
	else
	return null;
}
$(document).ready(function () {
        $('.file-box').each(function () {
            animationHover(this, 'pulse');
        });

	$.get("/api/v1.0/files", function(res) {
		$(".sk-spinner").css("visibility","hidden"); 
		var tpl="";
		for(var i=0; i<res.length; i++){
			tpl += `<div class="file-box" onclick="GetFiles('/`+res[i]+`')"><div class="file"><a><span class="corner"></span><div class="icon"><i class="fa fa-file"></i></div><div class="file-name">`+res[i]+`<br/></div></a></div></div>`
		}
		$("#Allfiles").html(tpl);	
		console.log(res);
		setCookie("now_path","/");
	}, "json");
	
	$('#submit').click(function(){
		$.ajax({  
             url : "/api/v1.0/upload",  
             data : {  
                 "now_path" : getCookie("now_path")
             },  
             type : 'POST',  
             success : function(res) { 

             },  
         });
          
	});
 });
		

function GetFiles(path){
	$(".sk-spinner").css("visibility","visible"); 
	$.ajax({  
              url : "/api/v1.0/files",  
              data : {  
                  "path" : path
              },  
              type : 'POST',  
              success : function(res) { 
			setCookie("now_path",path);
			$(".sk-spinner").css("visibility","hidden"); 
                  var tpl=`<div class="file-box" onclick="GetFiles('`+path+`/..')"><div class="file"><a><span class="corner"></span><div class="icon"><i class="fa fa-file"></i></div><div class="file-name">..<br/></div></a></div></div>`;
			for(var i=0; i<res.length; i++){
				tpl += `<div class="file-box" onclick="GetFiles('`+path+`/`+res[i]+`')"><div class="file"><a><span class="corner"></span><div class="icon"><i class="fa fa-file"></i></div><div class="file-name">`+res[i]+`<br/></div></a></div></div>`
			}
			$("#Allfiles").html(tpl);	
			console.log(res);
              },  
          });
          
}



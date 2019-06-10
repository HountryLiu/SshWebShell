$(document).ready(function () {


			//持续请求更新服务器状态
			function show(){
				$.get("/api/v1.0/system", function(res) {
				        console.log(res)
				}, "json");
			}
			setInterval(show,3000);
			$(".sk-spinner").css("visibility","visible");
			$.get("/api/v1.0/system", function(res) {
				$(".sk-spinner").css("visibility","hidden");
	            $("#sparkline1").sparkline([100-res['cpu_unused'], res['cpu_unused']], {
	                type: 'pie',
	                height: '140',
	                sliceColors: ['#1ab394', '#F5F5F5']
	            });
	
	            $("#sparkline2").sparkline([res['mem_used'], res['mem_total']], {
	                type: 'pie',
	                height: '140',
	                sliceColors: ['#ed5565', '#F5F5F5']
	            });
	
	            $("#sparkline3").sparkline([res['hd_used'], res['hd_total']], {
	                type: 'pie',
	                height: '140',
	                sliceColors: ['#ed5565', '#F5F5F5']
	            });

				$("#cpu_stat").html(new Number(100.0-res['cpu_unused']).toFixed(1)+"/100")
				$("#mem_stat").html(res['mem_used']+"/"+res['mem_total'])
				$("#hd_stat").html(res['hd_used']+"/"+res['hd_total'])


			}, "json");
			
			
			
			$.jgrid.defaults.styleUI = 'Bootstrap';
			$("#process").click(function(){
				$(".sk-spinner").css("visibility","visible");
			  	$.get("/api/v1.0/process", function(res) {
					$(".sk-spinner").css("visibility","hidden");
					_putGrid(res);
				    console.log(res)
				}, "json");
			});
           
            // Configuration for jqGrid Example 1
			//数据装入表格
   		 	var _putGrid = function (gridData) {
	            $("#table_list_1").jqGrid({
	                data: gridData,
	                datatype: "local",
	                height: 300,
	                autowidth: true,
	                shrinkToFit: true,
	                rowNum: 14,
	                colNames: ['PID', '名称', '操作'],
	                colModel: [
	                    {
	                        name: 'pid',
	                        index: 'pid',
	                        width: 180,
	                      
	                    },
	                    {
	                        name: 'name',
	                        index: 'name',
	                        width: 300
	                    },
	                    {
	                        name: 'operation',
	                        index: 'operation',
	                        sortable: false,
	                        width: 300,
	                        formatter:function (value, grid, rows, state) { 
	                        	return "<button type=\"button\" class=\"btn btn-outline btn-danger\"  onclick=\"Kill(" + rows.pid + ")\">结束</button>"
	                        }
	                    }
	                ],
	                pager: "#pager_list_1",
	                viewrecords: true,
					caption: "进程管理",
	                hidegrid: false
	            }); 
			};
			// Add responsive to jqGrid
            $(window).bind('resize', function () {
                var width = $('.jqGrid_wrapper').width();
                $('#table_list_1').setGridWidth(width);
          
            });
        });

		//结束进程
        function Kill(pid) {  
           swal({
			        title: "您确定要结束这个进程吗",
			        text: "结束该进程后可能会影响服务器的正常运行！",
			        type: "warning",
			        showCancelButton: true,
			        confirmButtonColor: "#DD6B55",
			        confirmButtonText: "确定",
			        cancelButtonText: "取消",
			        closeOnConfirm: false
			    }, function () {
					var model = jQuery('#table_list_1').jqGrid('getRowData', pid);  
			    	$.ajax({  
		                url : "/api/v1.0/process",  
		                data : {  
		                    "pid" : pid,
							"name" : model.name  
		                },  
		                type : 'POST',  
		                success : function(data) { 
		                    if (data=='success') {  
		                        $("#table_list_1").jqGrid('delRowData', pid);
		                        swal("成功！", "该进程已被结束。", "success");
		                    } else {  
	                            swal("失败！", "该进程结束失败，可能已被处理！", "error");
		                    }  
		                },  
		            });  
			    });
       }  
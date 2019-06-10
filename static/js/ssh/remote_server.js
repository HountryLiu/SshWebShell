
$(document).ready(function () {
			

				$.jgrid.defaults.styleUI = 'Bootstrap';
				// Configuration for jqGrid Example 2
	            $("#table_list_2").jqGrid({
	                url: "/api/v1.0/server",
	                datatype: "json",
	                height: 400,
	                autowidth: true,
	                shrinkToFit: true,
	                rowNum: 20,
	                rowList: [10, 20, 30],
	                colNames: ['ID', '服务器IP', '端口', '帐号', '密码'],
	                colModel: [
	                    {
	                        name: 'id',
	                        index: 'id',
	                        width: 60,
	                        sorttype: "int",
	                        search: true
	                    },
	                    {
	                        name: 'ip',
	                        index: 'ip',
	                        editable: true,
	                        width: 90
	                    },
						{
	                        name: 'port',
	                        index: 'port',
	                        editable: true,
	                        width: 60
					    },
	                    {
	                        name: 'rusername',
	                        index: 'rusername',
	                        editable: true,
	                        width: 100
	                    },
	                    {
	                        name: 'rpassword',
	                        index: 'rpassword',
	                        editable: true,
	                        width: 100
	                    }
	                ],
	                pager: "#pager_list_2",
	                viewrecords: true,
	                add: true,
	                edit: true,
					editurl : "/api/v1.0/server",
	                addtext: 'Add',
	                edittext: 'Edit',
	                hidegrid: false
	            });
 			// Add selection
            $("#table_list_2").setSelection(4, true);


            // Setup buttons
            $("#table_list_2").jqGrid('navGrid', '#pager_list_2', {
                edit: true,
                add: true,
                del: true,
                search: true
            }, {
                height: 200,
                reloadAfterSubmit: true
            });

            // Add responsive to jqGrid
            $(window).bind('resize', function () {
                var width = $('.jqGrid_wrapper').width();
                $('#table_list_1').setGridWidth(width);
                $('#table_list_2').setGridWidth(width);
            });


            

           
        });
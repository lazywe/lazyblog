$(function(){
	
	toastr.options = {
		  "closeButton": true,/*关闭按钮*/
		  "debug": false,/*调试状态*/
		  "progressBar": true,/*倒计时进度条*/
		  "positionClass": "toast-top-center",/*提示信息位置改后两个参数*/
		  "showDuration": "0",/*显示持续时间*/
		  "hideDuration": "0",/*隐藏持续时间*/
		  "timeOut": "1000",/*超时时间*/
		  "extendedTimeOut": "0",/*延长时间*/
		  "showEasing": "swing",/*显示动画*/
		  "hideEasing": "linear",/*隐藏动画*/
		  "showMethod": "fadeIn",/*显示方式,slideDown,fadeIn*/
		  "hideMethod": "fadeOut"/*隐藏方式,slideUp,fadeOut*/
		}
	
	
	/*提示信息*/
	$.toast = function(title,msg,type){
		if(type=='error'){
			toastr.error(msg,title)
		}else if(type=='success'){
			toastr.success(msg,title)
		}else if(type=='warning'){
			toastr.warning(msg,title)
		}else{
			toastr.info(msg,title)
		}
	}
	
	/*window*/
	var t=10*1000;//设定超时10秒
	$.window = function(title,url){
		if(!title){
			title='操作';
		}
		//iframe层-父子操作
		tempwindow=parent.layer.open({
			skin: 'layui-layer-molv', //样式类名
		    type: 2,
		    area: ['900px', '540px'],
		    maxmin: true,
		    title:title,
		    content: url
		});
		return false;
	}
	
	
	/*confirm*/
	$.confirm = function(title,msg,fun,fun1){
			if(!title){
				title='提示信息';
			}
			var ts=parent.layer.confirm(msg, {
			    skin: 'layui-layer-molv', //样式类名
			    title:title,
			    btn: ['确认','取消'], //按钮
			    shift: 4 //动画类型
			}, function(){
				if(fun){
					eval(fun);fun(ts);
				}else{
					layer.close(ts);
				}
			},function(ts){
				if(fun1){
					eval(fun1);fun1();
				}
			});
	}
	
	$.postform = function(fun,url){
		var index = parent.layer.load(1, {time:t,shade: [0.3,'#000']});
		var dom=$('#form');
		if(dom.length==0){
			dom = $('#form',window.top.document);
			
		}
		var data=dom.serialize();
		if(!url){
			url=dom.attr('action');
		}
		var ajaxTimeoutTest =$.ajax({
			type:"post",
			timeout:t,
			url:url,
			data:data,
			dataType:'json',
			success:function(data){
				parent.layer.close(index);
				if(data.status==1){
					if(fun){
						eval(fun);fun(data);
					}
				}else{
					$.toast('提示信息',data.info,'error');
				}
			},
			complete :function(XMLHttpRequest,status){ 
		　　　　　if(status!='success'){       
				 	  ajaxTimeoutTest.abort(); 
				 	  $.toast('提示信息',"请求异常",'warning');
					}
				parent.layer.close(index);	
			　}
		});
	}
	
	
	/*登陆*/
	var start=true;
	$('.login').bind('click',function(){
		if(start==false){
			return;
		}
		start=false;
		var dom=$(this);
		$.postform(function(e){
			if(e.status==1){
				dom.unbind('click');
				location.href=e.url;
			}
		});
	})
	
	
	$('.save').bind('click',function(){
		if(start==false){
			return;
		}
		start=false;
		var dom=$(this);
		$.postform(function(e){
			if(e.status==1){
				dom.unbind('click');
				var tlayer = parent.layer.getFrameIndex(window.name);
				if(tlayer){
					var windowDom;
					try{
						windowDom=parent.$('#content-main iframe:visible')[0].contentWindow
					}catch(e){
						windowDom=window.parent;
					}
					windowDom.location.reload();
					parent.layer.close(tlayer);
				}else{
					location.href=e.url;
				}
			}
		});
	})
	
	
	/**
	 * 排序 
	 */
	$.updatesort = function(url){
		$.postform(function(e){
			location.reload();
		},url)
	}
	
	/**
	 * 弹出窗口访问事件
	 */
	$.c = function(msg,url){
		$.confirm('',msg,function(ts){
			if(start==false){
				return;
			}
			start=false;
			var index = parent.layer.load(1, {time:t,shade: [0.3,'#000']});
			var ajaxTimeoutTest =$.ajax({
				type:"post",
				timeout:t,
				url:url,
				dataType:'json',
				success:function(data){
					if(data.status==1){
						start=false;
						parent.layer.close(ts);
						parent.layer.close(index);
						location.reload();
					}else{
						$.toast('提示信息',data.info,'error');
						parent.layer.close(index);
					}
				},
				complete :function(XMLHttpRequest,status){ 
				　　　　if(status!='success'){       
					 	  ajaxTimeoutTest.abort(); 
					 	  $.toast('提示信息',"请求异常",'warning');
						}
					parent.layer.close(index);	
				　}
			});
		})
	}
	
	/**
	 * 绑定点击事件
	 */
	$('.clickhref').bind('click',function(){
		var t=$(this);
		var fun=eval(t.attr('fun'));
		if(fun==undefined){
			alert("请添加自定义fun");
		}
		fun();
	})

	var check=false;
	$('.selectimg').bind('click',function(){
			if(check==true){
				return;
			}
			check=true;
			selectimgfun($(this));
			setTimeout(function(){
				check=false;
			},2000)
	})

	$(window).ajaxStop(function(e){
	    start=true;
	});	
	function selectimgfun(t){
		var url=t.attr('url');
		//iframe层-父子操作
		layer.open({
		    type: 2,
		    skin: 'layui-layer-lan', //样式类名
		    title:"选择图片",
		    area: ['700px', '430px'],
		    maxmin: true,
		    content:url
		});
	}
})

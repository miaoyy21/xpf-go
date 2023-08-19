package xmd

const (
	OriginAPP = "http://manorapp.pceggs.com"
	OriginWWW = "http://www.pceggs.com"
)

const MinBetGold = 10000

type RequestURL string

const (
	URLBetUserBase       RequestURL = "http://manorapp.pceggs.com/IFS/Manor28/Manor28_UserBase.ashx"
	URLBetAnalyseHistory RequestURL = "http://manorapp.pceggs.com/IFS/Manor28/Manor28_Analyse_History.ashx"
	URLBetBetting1       RequestURL = "http://manorapp.pceggs.com/IFS/Manor28/Manor28_Betting_1.ashx"
	URLBetCustomModes    RequestURL = "http://manorapp.pceggs.com/IFS/Manor28/Manor28_Custom_Modes.ashx"
	URLBetModesBetting   RequestURL = "http://manorapp.pceggs.com/IFS/Manor28/Manor28_ModesBetting.ashx"
	URLBetRiddle         RequestURL = "http://manorapp.pceggs.com/IFS/Manor28/Manor28_MyRiddleDetail.ashx"

	URLPrizeIndexList RequestURL = "http://www.pceggs.com/duobao/duobao_indexlist.aspx?roomtype=2"
	URLPrizePrize     RequestURL = "http://www.pceggs.com/duobao/doubaoajax.aspx"
)

const URLOpenPrize = `/duobao/duobao_index.aspx`

var tests = `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
    <title>蛋蛋夺宝_蛋咖</title>
    <meta name="Keywords" content="免费话费,免费奖品,手机,话费,金蛋" />
    <meta name="Description" content="蛋咖夺宝版块，这里给游戏玩家提供金蛋夺宝项目，用户可用少量金蛋参与夺宝，获得话费、手机等奖品。" />
    <link href="http://www.pceggs.com/css/base.v20150930.css" rel="stylesheet" type="text/css" />
    <link href="http://www.pceggs.com/css/daohang.css" rel="stylesheet" type="text/css" />
    <script type="text/javascript" src="/js/jquery-1.6.2.min.js"></script>
    <link href="http://www.pceggs.com/css/ui-lightness/jquery-ui-1.8.16.custom.css" rel="stylesheet"
        type="text/css" />
    <script type="text/javascript" src="/js/jquery-ui-1.8.16.custom.min.js"></script>
    <link href="http://www.pceggs.com/css/tm_fc.css" rel="stylesheet" type="text/css" />
    <link href="http://www.pceggs.com/css/pceggs.css" rel="stylesheet" type="text/css" />
    <link href="css/db/duobao.css?v=22" rel="stylesheet" type="text/css" />
    <link href="http://www.pceggs.com/css/FuCeng.css" rel="stylesheet" type="text/css" />
    <script src="js/FuCeng.js" type="text/javascript"></script>
    <style type="text/css">
<!--
  #topNav {
width: 47px;
z-index: 100;                                                     /*设置浮动层次*/
overflow: visible;
position: fixed;
bottom: 210px;
/*bottom: 330px;*/
margin-left:980px;                                                        /* 其他浏览器下定位，在这里可设置坐标*/
_position: absolute;                                       /*IE6 用absolute模拟fixed*/
_top: expression(documentElement.scrollTop + documentElement.clientHeight-230 + "px"); /*IE6 动态设置top位置*/
/* documentElement.scrollTop 设置浮动元素始终在浏览器最顶，可以加一个数值达到排版效果 */
height: 63px;
}
 
-->
</style>
    <script language="javascript" type="text/javascript">

        $(function () {
            $('#dialog0').dialog({
                autoOpen: false,
                modal: true,
                height: 160,
                width: 360
            }
					);
            $('#dialog1').dialog({
                autoOpen: false,
                modal: true,
                height: 160,
                width: 360
            }
					);

        });
        function colsedialog(d) {

            $("#" + d).dialog("close");

        }




        function opendialog(id) {
            $("#dialog" + id).dialog("open");

        }
    </script>
</head>
<body id="bb_body">
    <a name="totop"></a>
    <form name="form1" method="post" action="./duobao_index.aspx" id="form1">
     
<meta http-equiv="X-UA-Compatible" content="IE=Edge" />
<link href="/css/jt_news20190124.css?t=1" rel="stylesheet" type="text/css" />
<link href="http://www.pceggs.com/css/daohang.css" rel="stylesheet" type="text/css" />
<link href="/css/pcdd_num.css" rel="stylesheet" type="text/css" />
<link rel="stylesheet" href="../vipcj/css/index_zp.css?t=3">
<script src="/js/makeconform.js" type="text/javascript"></script>
<script type="text/javascript" src="/js/cookie.js"></script>
<script src="/IndexMainStatic/js/Head.js?t=20180308" type="text/javascript"></script>
 <script type="text/javascript" src="/js/UserClick.js?v=1"></script>
<script type="text/javascript">
    //去除IE6背景缓存
    document.execCommand("BackgroundImageCache", false, true);

    function addfavorite(a, b) {
        try {
            window.external.addFavorite('http://www.pceggs.com', "蛋咖-游戏试玩平台"); //IE浏览器
        }
        catch (e) {
            try {
                window.sidebar.addPanel("蛋咖-游戏试玩平台", 'http://www.pceggs.com', ""); //Firefox浏览器
            }
            catch (e) {
                alert("请按 Ctrl+D 键添加到收藏夹");
            }
        }
    }
    function toDesktop() {
        try {
            debugger;
            var WshShell = new ActiveXObject("WScript.Shell");
            var oUrlLink = WshShell.CreateShortcut(WshShell.SpecialFolders("Desktop") + "\\蛋咖-游戏试玩平台.url");
            oUrlLink.TargetPath = location.href;
            oUrlLink.Save();
            alert("蛋咖-游戏试玩平台快捷方式已创建到桌面！");
        }
        catch (e) {
            alert("当前IE安全级别不允许操作！");
        }
    }

    function doSomething(evt) {
        var e = (evt) ? evt : window.event; //判断浏览器的类型，在基于ie内核的浏览器中的使用cancelBubble
        if (window.event) {
            e.cancelBubble = true;
        } else {
            //e.preventDefault(); //在基于firefox内核的浏览器中支持做法stopPropagation
            e.stopPropagation();
        }
    }

    function answerqenstion() {

        ShowMsgo.show("/activity/201209/dt/dtfc/fuceng/answer.aspx", 740, 610);


    }

    function closedd() {

        ShowMsgo.ok();
    }

    function closelinqu() {
        ShowMsgo.cancel();
        window.location.reload(true);
    }

    function aaddgame() {
        ShowMsgo.show("/Gain/new_guide_game.aspx", 658, 390);
    }
    function aaddgg() {
        ShowMsgo.show("/Gain/new_guide_gg.aspx", 658, 390);
    }
    var isgain = "0";
    var autoopen = "0";
    var registdays = "89";
    var userid = "31792426";


    var tohref = window.location.href;
    //if(tohref.indexOf("bbsPost.aspx")>-1||tohref.indexOf("bbsMain.aspx")>-1||tohref.indexOf("ShowList.aspx")>-1||tohref.indexOf("ShowInfo.aspx")>-1||tohref.indexOf("duobao_index.aspx")>-1||tohref.indexOf("duobao_next.aspx")>-1||tohref.indexOf("duobao_historylist.aspx")>-1||tohref.indexOf("bbsPostShow.aspx")>-1 ){

    $(window).bind("scroll", function () {
        var top_top = $(document).scrollTop(); //卷上去的高度
        if (top_top > 0) {
            $("#topNavUP").css("display", "");
        } else {
            $("#topNavUP").css("display", "none");
        }

    })

    // }






</script>
<style type="text/css">
    a:hover{text-decoration:none;}
</style>
<script>
    function IsPC() { //是否是PC端打开
        var userAgentInfo = navigator.userAgent;
        var Agents = ["Android", "iPhone",
        "SymbianOS", "Windows Phone",
        "iPad", "iPod"
        ];
        var flag = true;
        for (var v = 0; v < Agents.length; v++) {
            if (userAgentInfo.indexOf(Agents[v]) > 0) {
                flag = false;
                break;
            }
        }
        return flag;
    }
    console.log(!IsPC())
    $(function () {
        function getClientHeight() {
            //        窗口可视区域高度
            var clientHeight = 0;
            clientHeight = window.innerHeight || document.documentElement.clientHeight || '高度不兼容~';
            return clientHeight;
        }
        addEventListener("scroll", function () {
            //        滚动的实际高度 = 可视区域高度 + 滚动高度
            var clients = getClientHeight();
            var scHeights = document.documentElement.scrollTop || document.body.scrollTop;
            var outbox = $("body").height();
            var outfooter = outbox - $(".footerbox").height() + 50;
            var outbutton = ".index_outbox ul li a.outbox_top";
            if (scHeights > 400) {
                $(outbutton).fadeIn()
            } else {
                $(outbutton).fadeOut()
            }
            if ((scHeights + clients) > 400 && scHeights < outfooter) {
                $(".index_outbox .outbox_top").css({ "position": "fixed", "bottom": "20px" })
            }
            if ((scHeights + clients + 80) > outfooter) {
                $(".index_outbox .outbox_top").css({ "position": "absolute", "marginTop": outfooter - 270, "bottom": "inherit" })
            }
        }, false)
        $(".index_outbox .outbox_top").click(function () {
            $("html,body").animate({ scrollTop: 0 }, 500);
        });

        if (!IsPC()) {
            $(".index_outbox").addClass('is_phone')
            $('body').addClass('is_phone_body')
            $('#sjb').attr('href', 'http://www.y571.com/');
        }
    });
  </script>
<script>
    $(function () {
        function getClientHeight() {
            //				窗口可视区域高度
            var clientHeight = 0;
            clientHeight = window.innerHeight || document.documentElement.clientHeight || '高度不兼容~';
            return clientHeight;
        }
        addEventListener("scroll", function () {
            //				滚动的实际高度 = 可视区域高度 + 滚动高度
            var clients = getClientHeight();
            var scHeights = document.documentElement.scrollTop || document.body.scrollTop;
            var outbox = $("body").height();
            var outfooter = outbox - $(".footerbox").height() + 50;
            var outbutton = ".index_outbox ul li a.outbox_top";
            if (scHeights > 400) {
                $(outbutton).fadeIn()
            } else {
                $(outbutton).fadeOut()
            }
            if ((scHeights + clients) > 400 && scHeights < outfooter) {
                $(".index_outbox .outbox_top").css({ "position": "fixed", "bottom": "20px" })
            }
            if ((scHeights + clients + 80) > outfooter) {
                $(".index_outbox .outbox_top").css({ "position": "absolute", "marginTop": outfooter - 270, "bottom": "inherit" })
            }
        }, false)
        $(".index_outbox .outbox_top").click(function () {
            $("html,body").animate({ scrollTop: 0 }, 500);
        });

    });
	</script>
   <script>
       var _hmt = _hmt || [];
       (function () {
           var hm = document.createElement("script");
           hm.src = "https://hm.baidu.com/hm.js?f8f6a0064a3e891522bdf044119d462a";
           var s = document.getElementsByTagName("script")[0];
           s.parentNode.insertBefore(hm, s);
       })();
</script>
 
<!--
 <div style="text-align:center;  position: fixed; width:100%;" id="fcts"><div class="tp" style="height:39px;   margin-top:36px;   "><img src="/images/public/fcts.png" width="887" height="39" /></div></div> -->
<input type="hidden" id="hid_userid" value="31792426" /> 
<input type="hidden" id="hid_deviceid" value="1FBC93E8-6456-4433-A6FF-3EA6E186347D" />
<input type="hidden" id="hid_token" value="2wttobp9rctga52aodb8fucarj3rtau47zj5dcel" />
<div class="jt">
    <div class="pc_ewmfc c_xjcc_div" style="display:none">
		<!-- 体验会员活动 -->
		<ul class="right">
           
			<!--<li class="c_xxcc"><a href="javascript:void(0)" style=" z-index: 100">
                <img src="https://pcdd-app.oss-cn-hangzhou.aliyuncs.com/advimg/20220120/2022012016192084497258.jpg" width="180px" onclick="openManor(1);UserClick(31792426,596);">
                <div id="ddzybox" style="display: none; position: fixed; bottom: 10px; right: 30px;
                    background: url(http://www0.pceggs.com/manor/images/manor_bg.png) no-repeat;
                    background-size: 100% 100%; padding: 20px 50px;">
                    <div style="width: 45px; height: 68px; background: url(http://www0.pceggs.com/manor/images/manor_close.png)  no-repeat;
                        background-size: 100% auto; position: absolute; top: -11px; right: 16px;" onclick="openManor(2)">
                    </div>
                    <iframe id="manor" src="http://manorapp.pceggs.com/Pages/Manor/Games.aspx?userid=31792426&deviceid=1FBC93E8-6456-4433-A6FF-3EA6E186347D&token=2wttobp9rctga52aodb8fucarj3rtau47zj5dcel"
                        width="375" height="667" frameborder="0" style="border-radius: 20px;"></iframe>
                </div>
            </a></li>-->

            

			<li>
              <a class="c_manor28" href="javascript:void(0)" style="margin-bottom:160px;margin-left:20px; z-index: 100">
                <img src="http://www0.pceggs.com/play/images/manor_icon1.png" width="180px" onclick="openManor1('31792426','1FBC93E8-6456-4433-A6FF-3EA6E186347D','2wttobp9rctga52aodb8fucarj3rtau47zj5dcel');UserClick(31792426,597);">
                <div id="ddzybox1" style="display: none; position: fixed; bottom: 10px; right: 30px;
                    background: url(http://www0.pceggs.com/manor/images/manor_bg.png) no-repeat;
                    background-size: 100% 100%; padding: 20px 50px;">
                    <div style="width: 45px; height: 68px; background: url(http://www0.pceggs.com/manor/images/manor_close.png)  no-repeat;
                        background-size: 100% auto; position: absolute; top: -11px; right: 16px;" onclick="openManor1('31792426','1FBC93E8-6456-4433-A6FF-3EA6E186347D','2wttobp9rctga52aodb8fucarj3rtau47zj5dcel')">
                    </div>
                    <iframe id="Iframe1" src=""
                        width="375" height="667" frameborder="0" style="border-radius: 20px;"></iframe>
                </div>
            </a>  
            </li>
            

            <!-- 小鸡猜猜活动 -->
            

            <!-- 暑期星选限时福利 -->
            <li>
              <a class="c_vip1" href="http://www.pceggs.com/activity2023/vip/vipact.aspx" style="margin-bottom:10px;margin-left:20px; z-index: 100">
                <img class="c_vip2_img" src="https://pcdd-app.oss-cn-hangzhou.aliyuncs.com/advimg/20230809/2023080915300920978222.jpg" width="180px">
              </a>  
            </li>

          

		</ul>
	</div>
	<div class="jt_top">
		<div class="jt_top_nei">
			
            <div class="left">
				<p>ID：<a href="/pgComUserInfo.aspx?userid=31792426"><span style=" color: #ff3a3a;">31792426</span></a></p>
				<dd><a href="/Logout.aspx" style="color: #aeaeae;">[退出]</a> </dd>
				<dd>|</dd>
				<dd><a href="/myaccount/myeggs.aspx?id=1">我的账户</a> </dd>
                <dd>|</dd>
                
                   <dd><a style="color: #333;" href="/myaccount/MyMoney/my_zhgk.aspx?id=1">金蛋：<b><span id="new_glod" class="c_sumeggs">76,456,706</span> <img style="margin-left: -4px;" src="http://www.pceggs.com/images/public/eggs.gif"></b></a> </dd>
				<dd>|</dd>
				<dd><a style="color: #333;" href="/myaccount/MyCashBox.aspx">小金库：<b>0 <img style="margin-left: -4px;" src="http://www.pceggs.com/images/public/eggs.gif"></b></a> </dd>
                
                <dd>|</dd>
				<dd><a href="/myaccount/MyMsg.aspx"><b>
                
                        <img style="margin-right: 2px;" src="http://www.pceggs.com/images/public/imgxx.png">
                      0</b></a> </dd>
			</div>
            
			<div class="right">
				<ul>
				    
                    
                    
                    
                    
                    
                    
                    

                    <li class="phone">
                        <a href="http://www.y571.com/" id="sjb" target="_blank"><img src="/activity/images/phone.png" style="vertical-align: middle; margin-top: -5px;"/> 手机版</a>
                        <div class="index_sjewm">
                            <img src="http://www.pceggs.com/images/NewAppHtml/ewm_app.png?rand=20210624"/>
                            <p style="line-height: 24px;">移动端更赚钱</p>
                        </div>
                        <script type="">
                            $(".jt_top_nei .right ul li.phone").hover(function () {
                                $(this).children(".index_sjewm").css("display", "block");
                            }, function () {
                                $(this).children(".index_sjewm").css("display", "none");
                            }
                            );
                        </script>
                    </li>
                    

                    <li>|</li>
                    
                    
                    <li><a href="/duobao/duobao_indexlist.aspx" class="list">夺宝</a> </li>
                    
                   
                    <!--<li>|</li>
                    
                    <li><a href="/invite/inviteindex.aspx">邀请</a> </li>
                    -->
                    <li>|</li>
                    
                    <li><a href="http://www.pceggs.com/vip/indexsx.aspx" class="list" style="font-weight:bold">星选VIP</a></li>
                    <li>|</li>
                    
                    <li><a href="/ShowPrize/ShowList.aspx?topicid=2">晒奖</a> </li>
                    
					<li>|</li>
                    
                    <li><a href="/ShowPrize/ShowList_gg.aspx?topicid=9">公告</a> </li>
                    
					<li>|</li>
                    
                    <li><a href="/services/ser_index.aspx">客服中心</a> </li>
                    
					<li>|</li>
                    
                    <li><a href="/help/h.aspx">帮助</a> </li>
                    
                    <li>|</li>
                    <li><a href="javascript:void(0)" onclick="addfavorite('','')"  rel="sidebar">收藏</a> </li>
<li>|</li>
                    <li><a href="javascript:void(0)" onclick="toDesktop()">桌面</a> </li>

				</ul>
			</div>
		</div>
	</div>
	<div class="jt_end jt_end_box">
		<a class="logos" href="http://www.pceggs.com/"><img src="/activity/images/n_logo.png"/> </a>
        
        <script src="/js/pcdd_num.js" type="text/javascript"></script>
        
        <div class="pcdd_number">
			<input name="Head2$WithdrawCount" type="hidden" id="Head2_WithdrawCount" value="16402852">
			<div class="header_xnkl">
				<b>累计游戏试玩人次</b>
                <ul style="display:none">
                        <li style="border: 1px solid #ffffff;">
                            <img src="/images/2019.gif"></li>
                    </ul>
			</div>
			<ul id="ulTotalBuy">
                
                <li class="pcdd_num">0</li>
                <li class="pcdd_nobor">,</li>
                <li class="pcdd_num">0</li>
                <li class="pcdd_num">0</li>
                <li class="pcdd_num">0</li>
                <li class="pcdd_nobor">,</li>
                <li class="pcdd_num">0</li>
                <li class="pcdd_num">0</li>
                <li class="pcdd_num">0</li>
                
            </ul>
		</div>
		<ul class="jt_nav">
            
                  <li><a href="/pceggsindex.aspx" id="shouye2" ctype="561">首页</a> </li>
            
                  <li><a href="/activityCenter/activityCenter.aspx"  ctype="566">热门活动</a> <w class="ico-num">new<em></em></w></li>
            
            <li><a href="/game/WebGame.aspx" id="yxsw2" ctype="564">游戏试玩</a> </li>
            
            <li class="userhidden" style="display:none"><a href="/game/QPGame.aspx" id="wqp2" ctype="565">益智试玩</a><w class="ico-num">hot<em></em></w> </li>
            
            <li style="display:none"><a href="/AdExperience/index1.aspx" id="ggty2" ctype="567">广告体验</a></li>
            
            <li><a href="/play/playIndex.aspx" ctype="568">休闲游戏</a> </li>
            
            <li><a href="/invite1/invite.aspx" ctype="571">邀请好友</a> </li>
            
            <li><a href="/newtb/NewTBIndex1.aspx" ctype="569">购物返利</a> </li>
            
            <li><a href="http://t3j4.pceggs.com" ctype="570">商城</a></li>
		</ul>
	</div>
    <div class="index_outbox">
        <ul>
          <li><a class="outbox_top"><img src="http://www.pceggs.com/images/outbox_top.png"/> </a> </li>
        </ul>
    </div>
    <div class="pc_ewmfc c_left_hdiv">
        <!-- 体验会员活动 -->
        <ul>
          

            <!-- 暑期星选限时福利  -->
            <li style="display:none"><a class="c_vip1" href="http://www.pceggs.com/activity2023/vip/vipact.aspx" target="_blank" style="margin-bottom:300px;margin-left:20px;"  tip=1 ><img class="c_vip1_img" src="https://pcdd-app.oss-cn-hangzhou.aliyuncs.com/advimg/20230809/2023080915290951493235.jpg" width="130px"/> </a> </li>

            <!-- 新人活动  -->
            <li style="display:none"><a class="c_isd" href="http://www.pceggs.com/NewTask20190618/NewTaskIndex.aspx" target="_blank" style="margin-bottom:300px;margin-left:20px;"  tip=1 ><img src="https://pcdd-app.oss-cn-hangzhou.aliyuncs.com/advimg/20230524/2023052417512469698568.jpg" width="130px"/> </a> </li>

            <!-- PC 618  -->
    
            <!-- 会员抽奖  -->
            <li class="c_sxvipzp showturn" style="display:none"><a href="javascript:void(0)" tip=1 style=" margin-bottom:300px;margin-left:20px;" style="cursor: pointer"><img src="https://pcdd-app.oss-cn-hangzhou.aliyuncs.com/advimg/20230423/2023042314252326825837.jpg" width="130px"/> </a> </li>

            <!-- 游戏体验卡  -->
            <li onclick="UserClick(31792426,597);"><a href="http://www.pceggs.com/activity2023/ty/ty.aspx" tip=1 target="_blank" style=" margin-left:20px;" style="cursor: pointer"><img src="https://pcdd-app.oss-cn-hangzhou.aliyuncs.com/advimg/20230420/2023042010192050821331.jpg" width="130px"/> </a> </li>

           <!-- 试玩抽奖 -->
           
                 
        
           <!-- 试玩大礼包  -->
           <li class="c_jialiangbao" onclick="UserClick(31792426,595);"><a href="http://www.pceggs.com/activity2022/swdlb/swdlb.aspx"  tip=1 target="_blank" style="margin-bottom:140px;margin-left:20px;" style="cursor: pointer"><img src="https://pcdd-app.oss-cn-hangzhou.aliyuncs.com/advimg/20230607/2023060715310764506237.jpg" width="130px"/> </a> </li>

           <!-- 数字猜猜第6期活动  -->
      <!--双11-->
         
           <!--年货节-->
           
           <!--38-->
                     <!--618-->
                     <!-- 专享福袋 -->
           
             <!-- 元宵活动 -->
           
          <!--数字猜猜活动 -->
                    <!-- 会员加送60% -->
                      <!-- 购物返利 -->
                     
        </ul>
    </div>
    <div id="totopNava" style="display:none"><div class="totopNav">
<!--头部右侧广告位-->
<!--
   <a href="http://www.pceggs.com/duobao/duobao_indexlist.aspx" target="_blank">
 <img src="/IndexMainStatic/hd223.jpg" border="0" style="cursor: pointer; margin-top: 3px;" />
 </a>

   <a href="http://t3j4.pceggs.com/" target="_blank">
 <img src="/IndexMainStatic/hd102.jpg" border="0" style="cursor: pointer; margin-top: 3px;" />
 </a>

   <a href="http://t3j4.pceggs.com/card/index.aspx" target="_blank">
 <img src="/IndexMainStatic/hd101.jpg" border="0" style="cursor: pointer; margin-top: 3px;" />
 </a>




<style>
.bgzz{ background:url(/IndexMainStatic/bg.png) no-repeat; width:78px; height:25px; color:#FFF; font-family: "微软雅黑"; font-size:14px; padding:17px 0 0 30px;margin-top: 3px;}

</style>

<div class="bgzz">800013154</div>-->
    

</div></div>
</div>
<script type="text/javascript">
    function SetVersion()
    {
        if(getCookie_top("now_version")==1)
        {
            SetCookie_top("now_version", 0, "1");
            window.location.href="/pceggsindex.aspx"; 
        }
        else
        {
            SetCookie_top("now_version", 1, "1");
            window.location.href="/index.aspx"; 
        }
    }
    function closeTop(day) {
        document.getElementById("topFloat").style.display = "none";
        document.getElementById("top_hidden").style.display="none";
        SetCookie_top('dd_top_url', '1', day);
    }
    function SetCookie_top(name, value, day)
    {
        var exp = new Date();//1天过期代码,下面是当天过期代码
        exp.setTime(exp.getTime() + parseInt(day) * 24 * 60 * 60 * 1000);
        document.cookie = name + "=" + escape(value) + ";expires=" + exp.toGMTString() + ";path=/;domain=pceggs.com";

        //var curDate = new Date();  
  
        //当前时间戳  
        //     var curTamp = curDate.getTime();  
  
        //当日凌晨的时间戳,减去一毫秒是为了防止后续得到的时间不会达到00:00:00的状态  
        //    var curWeeHours = new Date(curDate.toLocaleDateString()).getTime() - 1;  
  
        //当日已经过去的时间（毫秒）  
        //    var passedTamp = curTamp - curWeeHours;  
  
        //当日剩余时间  
        //   var leftTamp = 24 * 60 * 60 * 1000 - passedTamp;  
        //   var leftTime = new Date();  
        //   leftTime.setTime(leftTamp + curTamp);  
        //创建cookie  
        //  document.cookie = name + "=" + escape(value) + ";expires=" + leftTime.toGMTString() + ";path=/;domain=pceggs.com";
    }


    function getCookie_top(name)//取cookies函数        
    {
        var arr = document.cookie.match(new RegExp("(^| )" + name + "=([^;]*)(;|$)"));
        if (arr != null) return unescape(arr[2]); return null;

    }
    function showSiteTop(i) {
        if (i < 70) {
            var hh = i - 70;
            // document.getElementById("topFloat").style.top = hh + "px";
            i = i + 1;
            setTimeout("showSiteTop(" + i + ")", 30);
        }
    }
    var mob_status = 1;
    //alert(vip_level);
    if(mob_status != 1){
        if (getCookie_top('dd_top_url') != 1) {
            //showSiteTop(0);
            // document.getElementById("topFloat").style.display = "";
            //document.getElementById("top_hidden").style.display="";
        }
    }
    function showSiteTop2(i) {
        if (i <= 22) {
            i++;
            setTimeout("showSiteTop2(" + i + ")", 50);
        }
    }
    function showvipcj(){debugger
       $.ajax({
            type: "POST",
            data: { 'action': 'info' },
            url: "../sxvipzp.ashx",
            dataType: "json",
            error: function (XMLHttpRequest, textStatus, errorThrown) {
                debugger
            },
            success: function (data, textStatus) {
                debugger
                if (data.result == "0" && (data.isclick == "1" || data.isclick == "2")) {
                    $(".c_sxvipzp").show();
                }
                if(data.isd == "1"){
                    $(".c_isd").attr("href","http://www.pceggs.com/activity2023/newfl/newfl.aspx")
                    $(".c_isd1").attr("href","http://www.pceggs.com/activity2023/newfl/newfl.aspx")
                    $(".c_isd").closest("li").show();
                }else{
                    $(".c_isd").closest("li").hide(); 
                }
                if(data.vip1 == "1"){
                    $(".c_vip1_img").attr("src","https://pcdd-app.oss-cn-hangzhou.aliyuncs.com/advimg/20230809/2023080915290951493235.jpg?t=" + new Date().getTime());
                    $(".c_vip2_img").attr("src","https://pcdd-app.oss-cn-hangzhou.aliyuncs.com/advimg/20230809/2023080915300920978222.jpg?t=" + new Date().getTime());
                    $(".c_vip1").closest("li").show();
                }else if(data.vip1 == "2"){
                    $(".c_vip1_img").attr("src","https://pcdd-app.oss-cn-hangzhou.aliyuncs.com/advimg/20230809/2023080915290959006511.jpg?t=" + new Date().getTime());
                    $(".c_vip2_img").attr("src","https://pcdd-app.oss-cn-hangzhou.aliyuncs.com/advimg/20230809/2023080915300930418177.jpg?t=" + new Date().getTime());
                    $(".c_vip1").closest("li").show();
                } else{
                    $(".c_vip1").closest("li").show(); 
                }
            }
        });
    }
    showvipcj();

</script>
<script>

    function closeWX(i) {
        document.getElementById("pceggsmoblie" + i).style.display = "none";
    }

</script>
<script type="text/javascript" language="javascript">
    function openManor(i){ debugger
       
//        var show = $('#ddzybox').css('display');
//        $('#ddzybox').css('display',show =='block'?'none':'block');
//        if (i == 1) {
//            $(".c_manor28").attr("style", "margin-bottom:160px;margin-left:20px;z-index:1");
//        } else {
//            $(".c_manor28").attr("style", "margin-bottom:160px;margin-left:20px;"); 
//        }
       window.open('http://www.pceggs.com/game/gameindex/gameindex.aspx?gameid=3');
    }
    function openManor1(tuserid,tdeviceid,ttoken){ 
//        var show = $('#ddzybox1').css('display');
//        $('#ddzybox1').css('display',show =='block'?'none':'block');
//        $('#Iframe1').attr('src','http://manorapp.pceggs.com/Pages/Manor28/ManorIndex.aspx?userid='+tuserid+'&deviceid='+tdeviceid+'&token='+ttoken+'&pcback=1') ;
       window.open('http://www.pceggs.com/game/gameindex/gameindex.aspx?gameid=4');
    }
    $(document).ready(function () {
        $.each($(".jt_nav li a"),function(i,o){
            var t_ctype =$(o).attr("ctype");
            $(o).bind("click",function(){debugger
                $.ajax({
                    type: "POST",
                    data: "action=HeadClick&t=" + new Date()+"&ctype="+t_ctype,
                    url: "/Head_Ajax.ashx",
                    dataType: "json",
                    success: function (ret) {
                        message = ret.message;
                        if (ret.result == 0) {
                            $("#fuceng").html(ret.message);
                            $("#fuceng").css("display", "block");
                            $(".c_zh_tc").hide();
                        }
                        else {
                            $('#dialog0').dialog("open");
                            $("#pagecontent").html(message);
                        }
                    }
                });
            });
        })


       
     $.ajax({
         type: "POST",
         data: { 'action': 'getUserType' },
         url: "http://www.pceggs.com/Head_Ajax.ashx",
         dataType: "json",
         error: function (XMLHttpRequest, textStatus, errorThrown) {
         },
         success: function (data, textStatus) {
             if (data.status == 27996297) {
                 $(".userhidden").hide();
                 return;
             }
         }
     });
 });
 var r = window.location.href;
 if(r.indexOf("hb_activity.aspx")>-1){
     $("#fcts").css("display","none");
 }
 if(r.indexOf("/play/pxya.aspx")>-1||r.indexOf("prize/PrizeMain.aspx")>-1){
     $("#totopNava").css("display","block");
 }
 if(r.indexOf("/play/")>-1||r.indexOf("/gameindex/")>-1){
     $(".c_xjcc_div").show();
     $(".c_left_hdiv").hide();
 }
</script>


<!-- 星选VIP抽奖 begin -->
<script src="../vipcj/js/index_zp.js?t=3"></script>
<div class="turn" style="display: none;">
    <div class="turn-content" style="display: none;">
        <div class="turn-title"></div>
        <div class="turn-list">
            <div class="turn-msg"></div>
            <ul></ul>
        </div>
        <div class="turn-close"></div>
    </div>
    <div class="turn-award" style="display: none;">
        <p class="turn-award-title">徽章</p>
        <img src="" width="138"  height="100" alt="" class="turn-award-img">
        <p class="turn-award-msg">徽章已到账，快去兑换奖品吧！</p>
        <a href="http://www.pceggs.com/myaccount/GrowSys/RewardExChang.aspx" class="turn-award-btn"></a>
        <div class="turn-close"></div>
    </div>
</div>
<!-- 星选VIP抽奖 end -->
    <div class='div_full' id='parent_div' style="width: 100%; display: none; left: 0px;
        top: 0px; filter: alpha(opacity=40); -moz-opacity: 0.4; -khtml-opacity: 0.4;
        opacity: 0.4;">
    </div>
    <div class="doc">
        <div class="pc_step" style="padding-top: 0px">
            <img src="img/db/db_step.gif?t=1" width="950" height="56" /></div>
        <div class="pc_line">
            <div class="pc_left">
                <div class="pc_le" style="margin-bottom: -6px;">
                    <div style="position: absolute; margin-left: 872px; margin-top: 10px; _margin-top: 12px;
                        font-size: 14px;">
                        <img src="img/db/wddb_tb.jpg" align="absmiddle" />&nbsp;<a href="/myaccount/myduobao.aspx"
                            style="color: #F8E0E9; text-decoration: none;">我的夺宝</a>
                    </div>
                    <div class="header_2">
                        <ul class="nav_main">
                            <li class="nav_basic"><a  title="当前夺宝"
                                href="duobao_indexlist.aspx?roomtype=1"><strong>幸运场夺宝</strong></a>
                            </li>
                            <li class="nav_basic"><a class="current"  title="当前夺宝"
                                href="duobao_indexlist.aspx?roomtype=2"><strong>土豪场夺宝</strong></a><sup class="ico_current">当前</sup>
                            </li>
                            <li class="nav_basic"><a title="下期夺宝" href="duobao_next.aspx"><strong>下期夺宝</strong></a>
                            </li>
                            <li class="nav_basic"><a title="往期夺宝" href="duobao_historylist.aspx"><strong>往期夺宝</strong></a>
                            </li>
                        </ul>
                    </div>
                </div>
                <div class="pc_le">
                    <div class="basic_buy_group" style="width: 946px; height: auto;">
                        <div class="liebiao">
                            <ul>
                                
                                <li class="spzs">
                                    
                                    <img class="spt" src="http://www.pceggs.com/IndexMainStatic/duobao/img/20220614005/20220614005_l.jpg" alt="" onclick="window.open('/duobao/duobao_index.aspx?id=10061476')">
                                    <div class="sp_name">
                                        星巴克星礼卡100元（卡密）</div>
                                    <div class="db_xh">
                                        需要：<span>1,566万</span> 金蛋</div>
                                    <div class="db_jd">
                                        <div class="jd_bt">
                                            进度：</div>
                                        <div class="jdt_wk">
                                            <div class="jdt_jd" style="width: 89.91%;">
                                            </div>
                                        </div>
                                        <div class="jd_wz">
                                            89.91%</div>
                                    </div>
                                    
                                    <div class="an_db" onclick="checkOrder(10061476,10000);">
                                        <img src="img/db/an_cydb.png" alt=""></div>
                                    
                                </li>
                                
                                <li class="spzs">
                                    
                                    <img class="spt" src="http://www.pceggs.com/IndexMainStatic/duobao/img/20220614005/20220614005_l.jpg" alt="" onclick="window.open('/duobao/duobao_index.aspx?id=10061490')">
                                    <div class="sp_name">
                                        星巴克星礼卡100元（卡密）</div>
                                    <div class="db_xh">
                                        需要：<span>1,566万</span> 金蛋</div>
                                    <div class="db_jd">
                                        <div class="jd_bt">
                                            进度：</div>
                                        <div class="jdt_wk">
                                            <div class="jdt_jd" style="width: 49.94%;">
                                            </div>
                                        </div>
                                        <div class="jd_wz">
                                            49.94%</div>
                                    </div>
                                    
                                    <div class="an_db" onclick="checkOrder(10061490,10000);">
                                        <img src="img/db/an_cydb.png" alt=""></div>
                                    
                                </li>
                                
                                <li class="spzs">
                                    
                                    <img class="spt" src="http://www.pceggs.com/IndexMainStatic/duobao/img/20220509004/20220509004_l.jpg" alt="" onclick="window.open('/duobao/duobao_index.aspx?id=10061489')">
                                    <div class="sp_name">
                                        100元商城抵扣券</div>
                                    <div class="db_xh">
                                        需要：<span>1,468万</span> 金蛋</div>
                                    <div class="db_jd">
                                        <div class="jd_bt">
                                            进度：</div>
                                        <div class="jdt_wk">
                                            <div class="jdt_jd" style="width: 67.17%;">
                                            </div>
                                        </div>
                                        <div class="jd_wz">
                                            67.17%</div>
                                    </div>
                                    
                                    <div class="an_db" onclick="checkOrder(10061489,10000);">
                                        <img src="img/db/an_cydb.png" alt=""></div>
                                    
                                </li>
                                
                                <li class="spzs">
                                    
                                    <img class="spt" src="http://www.pceggs.com/IndexMainStatic/duobao/img/2021041307/2021041307_l.jpg" alt="" onclick="window.open('/duobao/duobao_index.aspx?id=10061499')">
                                    <div class="sp_name">
                                        天猫超市享淘卡200元</div>
                                    <div class="db_xh">
                                        需要：<span>3,726万</span> 金蛋</div>
                                    <div class="db_jd">
                                        <div class="jd_bt">
                                            进度：</div>
                                        <div class="jdt_wk">
                                            <div class="jdt_jd" style="width: 40.39%;">
                                            </div>
                                        </div>
                                        <div class="jd_wz">
                                            40.39%</div>
                                    </div>
                                    
                                    <div class="an_db" onclick="checkOrder(10061499,10000);">
                                        <img src="img/db/an_cydb.png" alt=""></div>
                                    
                                </li>
                                
                                <li class="spzs">
                                    
                                    <div class="xxzx">
                                        <img src="img/db/xxzx.png"></div>
                                    
                                    <img class="spt" src="http://www.pceggs.com/IndexMainStatic/duobao/img/20211220007/20211220007_l.jpg" alt="" onclick="window.open('/duobao/duobao_index.aspx?id=10061502')">
                                    <div class="sp_name">
                                        天猫超市享淘卡500元</div>
                                    <div class="db_xh">
                                        需要：<span>9,316万</span> 金蛋</div>
                                    <div class="db_jd">
                                        <div class="jd_bt">
                                            进度：</div>
                                        <div class="jdt_wk">
                                            <div class="jdt_jd" style="width: 16.13%;">
                                            </div>
                                        </div>
                                        <div class="jd_wz">
                                            16.13%</div>
                                    </div>
                                    
                                    <div class="an_db" onclick="checkOrder(10061502,10000);">
                                        <img src="img/db/an_cydb.png" alt=""></div>
                                    
                                </li>
                                
                                <li class="spzs">
                                    
                                    <img class="spt" src="http://www.pceggs.com/IndexMainStatic/duobao/img/20220614005/20220614005_l.jpg" alt="" onclick="window.open('/duobao/duobao_index.aspx?id=10061504')">
                                    <div class="sp_name">
                                        星巴克星礼卡100元（卡密）</div>
                                    <div class="db_xh">
                                        需要：<span>1,566万</span> 金蛋</div>
                                    <div class="db_jd">
                                        <div class="jd_bt">
                                            进度：</div>
                                        <div class="jdt_wk">
                                            <div class="jdt_jd" style="width: 0%;">
                                            </div>
                                        </div>
                                        <div class="jd_wz">
                                            0%</div>
                                    </div>
                                    
                                    <div class="an_db" onclick="checkOrder(10061504,10000);">
                                        <img src="img/db/an_cydb.png" alt=""></div>
                                    
                                </li>
                                
                            </ul>
                        </div>
                    </div>
                </div>
                <div class="pc_le">
                    <div class="site_push">
                    </div>
                </div>
            </div>
            <div class="pc_right" style="margin-top: -0.1rem">
                <div class="tbdp">
                    <div class="db_pcright_kj" id="helpid">
                        <div class="db_pcright_kjtop">
                            <div class="db_right_l">
                                夺宝规则</div>
                            <div class="db_right_r" onclick="window.open('/duobao/duobao_rule.aspx')">
                                查看>>
                            </div>
                        </div>
                        <div class="db_right_con" style="padding-top: 10px;">
                            <span>&nbsp;</span><a href="/duobao/duobao_rule.aspx" target="_blank">什么是夺宝</a><br />
                            <span>&nbsp;</span><a href="/duobao/duobao_rule.aspx" target="_blank">夺宝购买蛋拍号细则</a><br />
                            <span>&nbsp;</span><a href="/duobao/duobao_rule.aspx" target="_blank">夺宝开奖规则</a><br />
                            <span>&nbsp;</span><a href="/duobao/duobao_rule.aspx" target="_blank">如何查看蛋拍号</a><br />
                        </div>
                        <div class="db_pcright_kjbottom">
                        </div>
                    </div>
                </div>
                <div class="db_pcright_kj">
                    <div class="db_pcright_kjtop">
                        <div class="db_right_l">
                            最新中奖记录</div>
                        <div class="db_right_r">
                            <a href="duobao_historylist.aspx" class="blue_load">全部>></a></div>
                    </div>
                    <div class="db_right_con" style="height: auto;">
                         <div class='db_right_concon' > <div class='db_right_conl' ><a href='duobao_history.aspx?id=10061501' class='blue_load'><img src='http://www.pceggs.com/IndexMainStatic/duobao/img/20220614009/20220614009_s.jpg' width='78'  height='70'  /></a></div> <div class='db_right_conr'><a href='/pgComUserInfo.aspx?USERID=16672201' class='blue_load'>16672201</a><br/> 扔了：<span class='font_red'>31,660,000</span><img src='img/db/eggs.gif' width='16' height='16' align='absmiddle' /><br/> <a href='duobao_history.aspx?id=10061501' class='blue_load'>第20230811097期</a></div> </div> <div class='db_right_concon' > <div class='db_right_conl' ><a href='duobao_history.aspx?id=10061498' class='blue_load'><img src='http://www.pceggs.com/IndexMainStatic/duobao/img/20220614007/20220614007_s.jpg' width='78'  height='70'  /></a></div> <div class='db_right_conr'><a href='/pgComUserInfo.aspx?USERID=16672201' class='blue_load'>16672201</a><br/> 扔了：<span class='font_red'>6,350,000</span><img src='img/db/eggs.gif' width='16' height='16' align='absmiddle' /><br/> <a href='duobao_history.aspx?id=10061498' class='blue_load'>第20230811094期</a></div> </div> <div class='db_right_concon' > <div class='db_right_conl' ><a href='duobao_history.aspx?id=10061469' class='blue_load'><img src='http://www.pceggs.com/IndexMainStatic/duobao/img/2021041306/2021041306_s.jpg' width='78'  height='70'  /></a></div> <div class='db_right_conr'><a href='/pgComUserInfo.aspx?USERID=16672201' class='blue_load'>16672201</a><br/> 扔了：<span class='font_red'>7,460,000</span><img src='img/db/eggs.gif' width='16' height='16' align='absmiddle' /><br/> <a href='duobao_history.aspx?id=10061469' class='blue_load'>第20230811065期</a></div> </div> <div class='db_right_concon' > <div class='db_right_conl' ><a href='duobao_history.aspx?id=10061496' class='blue_load'><img src='http://www.pceggs.com/IndexMainStatic/duobao/img/2017020001/2017020001_s.jpg' width='78'  height='70'  /></a></div> <div class='db_right_conr'><a href='/pgComUserInfo.aspx?USERID=13304561' class='blue_load'>13304561</a><br/> 扔了：<span class='font_red'>9,000,000</span><img src='img/db/eggs.gif' width='16' height='16' align='absmiddle' /><br/> <a href='duobao_history.aspx?id=10061496' class='blue_load'>第20230811092期</a></div> </div> <div class='db_right_concon' > <div class='db_right_conl' ><a href='duobao_history.aspx?id=10061478' class='blue_load'><img src='http://www.pceggs.com/IndexMainStatic/duobao/img/20211220005/20211220005_s.jpg' width='78'  height='70'  /></a></div> <div class='db_right_conr'><a href='/pgComUserInfo.aspx?USERID=24556892' class='blue_load'>24556892</a><br/> 扔了：<span class='font_red'>37,900,000</span><img src='img/db/eggs.gif' width='16' height='16' align='absmiddle' /><br/> <a href='duobao_history.aspx?id=10061478' class='blue_load'>第20230811074期</a></div> </div> <div class='db_right_concon' style='border:0px;'> <div class='db_right_conl' ><a href='duobao_history.aspx?id=10061486' class='blue_load'><img src='http://www.pceggs.com/IndexMainStatic/duobao/img/20220614008/20220614008_s.jpg' width='78'  height='70'  /></a></div> <div class='db_right_conr'><a href='/pgComUserInfo.aspx?USERID=16452025' class='blue_load'>16452025</a><br/> 扔了：<span class='font_red'>31,660,000</span><img src='img/db/eggs.gif' width='16' height='16' align='absmiddle' /><br/> <a href='duobao_history.aspx?id=10061486' class='blue_load'>第20230811082期</a></div> </div>
                    </div>
                    <div class="db_pcright_kjbottom">
                    </div>
                </div>
            </div>
            <!-- *********** -->
            <div id="div_ad" style="position: absolute; z-index: 999; width: 528px; height: auto;
                display: none">
                <div class="dbfc_wk clearall">
                    <div class="dbfc_wk_top">
                        <div class="gb_an" id="close_div">
                            <a style="cursor: pointer;" onclick="return rm('div_ad');">
                                <img src="img/dbfc/gb_an.gif" border="0" /></a></div>
                    </div>
                    <div class="dbfc_wk_con clearall">
                       <div class="dbfc_wk_con_vip" style="text-align: center;padding-bottom: 5px;display:none;">
                            <a id="img_ktsxvip" target="_blank" href="http://www.pceggs.com/VIP/indexsx.aspx" style="display:none;"><img style="" width="254" height="27" src="img/dbfc/ktviptips.png"></a>
                            <img id="img_sxvip" style="display:none" width="273" height="28" src="img/dbfc/sxviptips.png">
                        </div>
                        <span id="errormsg"></span>
                        <div class="dbfc_cz" style="width: 166px; display: none;" id="dbfc_cz">
                            <a class="red_enter" href="/myaccount/mymobileindex.aspx" target="_blank"><span>立即手机认证</span></a>&nbsp;&nbsp;
                            <a class="black_enter" onclick="return rm('div_ad');"><span>返回</span></a>
                        </div>
                        <div class="dbxzsm" id="dbxzsm">
                            <div style="padding-top: 35px;">
                                幸运场用户等级对应的最多可购买每期蛋拍号的组数。</div>
                            <div style="padding-top: 90px;">
                                土豪场用户等级对应的最多可购买每期蛋拍号的组数。</div>
                            <div style="padding-top: 110px;">
                                未完成<a href="#">手机认证</a>，不能参加夺宝。</div>
                            <div style="padding-top: 18px;">
                                当天领取救济后，不能参加夺宝。</div>
                            <div style="padding-top: 18px;">
                                夺宝中奖后，<span class="dbxzsm_red">7天内</span>只能领取<span class="dbxzsm_red">5000</span>金蛋救济。</div>
                            <div style="padding-top: 16px;">
                                <span class="dbxzsm_red">作弊行为</span>一经发现将冻结账号并取消夺宝资格，购买的蛋拍号全部作废。</div>
                        </div>
                    </div>
                    <div class="dbfc_wk_bottom">
                    </div>
                </div>
            </div>
            <!-- *********** -->
            <!-- ui-dialog -->
            <div id="dialog0" title="提示" style="overflow: hidden; display: none;">
                <p style="text-align: center; line-height: 60px;" id="pagecontent">
                    金蛋是蛋咖的虚拟币，金蛋可以兑换实物奖品。</p>
            </div>
            <div id="dialog1" title="提示" style="overflow: hidden; display: none;">
                <p style="text-align: center; line-height: 60px;" id="p1">
                    去“游戏试玩中心”免费试玩游戏即可获得金蛋奖励。</p>
                <div class="annu_auto" style="height: 30px;">
                    <a class="red_enter" href="/Gain/GnGameAll.aspx"><span style="color: White">立即试玩</span></a>
                    &nbsp; &nbsp;&nbsp;&nbsp;<a class="red_enter" onclick="colsedialog('dialog1')"><span
                        style="color: White">返回</span></a>
                </div>
            </div>
            <!-- 京东E卡 星选会员专享 -->
            <div id="sxvipdb" style="display: none;">
                <div class="kt_xx">
                    <div class="names">
                        <a href="">
                            <img src="img/db/close1.png" onclick="colsedialog('sxvipdb')"></a></div>
                    <div class="nei">
                        <b>星选VIP专享</b>
                        <p>
                            开通星选VIP即可参与夺宝</p>
                        <div class="button">
                            <img src="img/db/tc-aniu2.png" onclick="govip()"></div>
                    </div>
                </div>
                <div class="kt_tc">
                </div>
            </div>
    </form>
    

<div class="nyfooter_box footerbox">
	<div class="nei">
		<div class="left">
			<p>
			<ul>
				<li><a href="http://www.pceggs.com/pchome/pc_home.aspx" target="_blank">关于蛋咖</a> </li><li>|</li>
				<li><a href="http://www.pceggs.com/help/h.aspx?id=11" target="_blank">服务条款</a> </li><li>|</li>
				<li><a href="/newggfw/gg_index.html" target="_blank">广告服务</a> </li><li>|</li>
				<li><a href="/newggfw/gg_hzkh.html" target="_blank">合作商家</a> </li><li>|</li>
				<li><a href="/pchome/pc_home.aspx" target="_blank">蛋友生活</a> </li>
			</ul>
			</p>
			<p>版权所有：杭州蛋咖网络技术有限公司 ICP证：浙B2-20090227  <img src="http://www.pceggs.com/images/head/logo/police.png"/> 浙公网安备 33010602004295号 </p>
		</div>
		
	</div>
</div>
</div>

    <script type="text/javascript" language="javascript">
        function checkOrder(id, price) {
            $("#div_ad").show();
            $(".div_full").show();
            ShowParentFC();
            sc1();
            $.ajax({
                type: "POST",
                data: { 'action': 'ordercheck', 'id': id },
                url: "doubaoajax.aspx",
                dataType: "json",
                error: function () {
                    //alert(111111);
                },
                success: function (ret) {
                    if (ret.result == "0") {
                        $("#dbxzsm").css("display", "");
                        $("#errormsg").html('<p id="errortip">' + ret.msg + '</p><div class="dbfc_cz" style="width:480px;text-align:center;">' + price + '金蛋/组，购买的组数：<input name="textfield" type="text" value="' + ret.dbnum + '" style="width:80px;" id="total"/>&nbsp;&nbsp; <a id="btn_ljgm" class="red_enter" onclick="order(' + id + ',' + price + ')"><span>立即购买</span></a> </div>');
                        if (ret.issxtip == "0") {
                            $(".dbfc_wk_con_vip").hide();
                        } else if (ret.issxtip == "1") {
                            $(".dbfc_wk_con_vip").show();
                            $("#img_sxvip").show();
                            $("#img_ktsxvip").hide();
                        } else if (ret.issxtip == "2") {
                            $(".dbfc_wk_con_vip").show();
                            $("#img_sxvip").hide();
                            $("#img_ktsxvip").show();
                        }
                    } else {
                        if (ret.result == "110") {

                            $("#errormsg").html("<p>您的账户金蛋不足，暂不能参加蛋蛋夺宝。</p>");


                        } else if (ret.result == "101") {

                            $("#errormsg").html('<p>对不起！您尚未登录，不能参加蛋蛋夺宝，请<a  onclick=\"Showlogin()\" style=\"cursor:pointer\"  >点此登录</a> ！ </p><p >您还没有帐号ID？请<a href="/newone/PgQCoinReg.aspx" class="blue_load">点此注册</a>!</p>');


                        } else {
                            $("#dbxzsm").css("display", "");

                            $("#errormsg").html("<p>" + ret.msg + "</p>");
                        }
                        if (ret.result == "103") {
                            $("#dbfc_cz").css("display", "");
                        }
                        if (ret.result == "176") {
                            $("#div_ad").hide();
                            $(".div_full").hide();
                            $("#sxvipdb").css("display", "");
                        }
                    }
                }
            });
        }

        function order(id, price) {
            var total = $("#total").val();

            var pattern = /^[0-9]*[1-9][0-9]*$/;
            if (!pattern.test(total)) {
                alert("购买数量格式错误");
                return;
            }
            if (!confirm("您确认要购买" + total + "组，共需" + (parseInt(total) * parseInt(price)) + "金蛋？")) {
                return;
            }


            $("#btn_ljgm").css("display", "none");
            $.ajax({
                type: "POST",
                data: { 'action': 'orders', 'id': id, 'total': total, 'price': price },
                url: "doubaoajax.aspx",
                dataType: "json",
                error: function () {
                    //alert(111111);
                },
                success: function (ret) {
                    if (ret.result == "0") {
                        $("#errormsg").html("<p>购买成功！您可以在我的夺宝中查看购买的蛋拍号。<a href='/myaccount/myduobaodetail.aspx?id=" + id + "' style='text-decoration:underline'>查看我的蛋拍号</a></p>");
                        $("#close_div").html("<a style=\"cursor:pointer;\" onclick=\"return rm('div_ad',1);\"><img src=\"img/dbfc/gb_an.gif\" border=\"0\" /></a>");

                    } else {


                        if (ret.result == "110" || ret.result == "1011") {

                            $("#errortip").html(ret.msg);

                            $("#btn_ljgm").css("display", "");

                        } else {
                            $("#errormsg").html("<p>" + ret.msg + "</p>");
                        }

                    }
                }
            });
        }

        //去星选会员页面
        function govip() {
            $("#sxvipdb").hide();
            window.open('http://www.pceggs.com/vip/indexsx.aspx');
        }
    </script>
</body>
</html>
`

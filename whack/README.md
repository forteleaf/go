웹핵 게임 클라이언트 도우미
whack helper

# 목적


# 원인

- 기술을 확인하기 위해서 매번 도움을 열러서 클릭해야 되는게 너무나 불편했다.

# 결과

그래서 내가 직접 만들기로 했다.

사용자보기
- http://wgame.kr/game.php?user_info=사용자이름

인벤
- http://wgame.kr/game.php?inven=1

조합
http://wgame.kr/game.php?mix=1

몬스터 부활시간
- http://wgame.kr/game.php?rank=1&m=5


<div align="right"><a class="tooltip"><input value="이벤트" type="button" class="icon_skill" style="color:black;width:50px;height:50px;background:gold;vertical-align:top" onclick="location.href='?event=1&amp;e=1'"></a>
<a class="tooltip"><input type="button" style="background-image:url('http://iwgame.cafe24.com/images/가방.gif');width:50px;height:50px;border:1px solid black" onclick="location.href='?inven=1'" class="icon_skill"></a>
<a class="tooltip"><input type="button" style="background-image:url('http://iwgame.cafe24.com/images/조합.gif');width:50px;height:50px;border:1px solid black" onclick="location.href='?mix=1'" class="icon_skill"></a>
<a class="tooltip"><input type="button" style="background-image:url('http://iwgame.cafe24.com/images/표지판.gif');width:50px;height:50px;border:1px solid black" onclick="location.href='?rank=1'" class="icon_skill"></a>
<a class="tooltip"><input type="button" style="background-image:url('http://iwgame.cafe24.com/images/괴물뱀.gif');width:50px;height:50px;border:1px solid black" onclick="location.href='?rank=1&amp;m=5'" class="icon_skill"></a>
<a class="tooltip"><input type="button" style="background-image:url('http://iwgame.cafe24.com/images/boo5.gif');width:50px;height:50px;border:1px solid black;color:blue;vertical-align:top" onclick="location.href='?b_set=1'" value="461" class="icon_skill"></a>
<a class="tooltip"><input class="icon_skill" type="button" style="background-image:url('http://iwgame.cafe24.com/images/고대%20나가.gif');width:50px;height:50px;border:1px solid black;background-repeat:no-repeat;font-size:2pt;vertical-align:top" onclick="location.href='?dungeon=1'"></a>

<script>function hkey_1(){location.href='?u=%EC%88%9C%EA%B0%84%EC%9D%B4%EB%8F%99+%EC%A3%BC%EB%AC%B8%EC%84%9C%2B108&check=cf97f03b7941970f2987a53241a88246';}</script>
<input class="icon_skill" type="button" style="background-image:url('http://iwgame.cafe24.com/images/주문서.gif');width:50px;height:50px;border:1px solid black;background-repeat:no-repeat;font-size:2pt;vertical-align:top" onclick="location.href='?u=%EC%88%9C%EA%B0%84%EC%9D%B4%EB%8F%99+%EC%A3%BC%EB%AC%B8%EC%84%9C%2B108&amp;check=cf97f03b7941970f2987a53241a88246'" value="108">

<script>function hkey_2(){location.href='?u=%ED%8E%AB+%EC%95%8C%2B1&check=cf97f03b7941970f2987a53241a88246';}</script>
<input type="button" class="icon_skill" style="background-image:url('http://iwgame.cafe24.com/images/펫알.gif');width:50px;height:50px;border:1px solid black;background-repeat:no-repeat;font-size:2pt;vertical-align:top" onclick="location.href='?u=%ED%8E%AB+%EC%95%8C%2B1&amp;check=cf97f03b7941970f2987a53241a88246'" value="1">

<script>function hkey_3(){location.href='?u=%EC%82%AC%ED%95%B4%EC%9D%98+%EB%91%90%EB%A3%A8%EB%A7%88%EB%A6%AC%2B3&check=cf97f03b7941970f2987a53241a88246';}</script><input type="button" class="icon_skill" style="background-image:url('http://iwgame.cafe24.com/images/사해의 두루마리.gif');width:50px;height:50px;border:1px solid black;background-repeat:no-repeat;font-size:2pt;vertical-align:top" onclick="location.href='?u=%EC%82%AC%ED%95%B4%EC%9D%98+%EB%91%90%EB%A3%A8%EB%A7%88%EB%A6%AC%2B3&amp;check=cf97f03b7941970f2987a53241a88246'" value="3">
<script>
function g_help(){
dish=1;
$().toasty({message: "<img src='http://iwgame.cafe24.com/images/도움말.gif' class=icon_skill><br><font size=2>검색어를 입력해주세요.</font><br><input type=text id=g_hv onkeypress=if(event.keyCode==13){location.href='help.php?help='+encodeURIComponent(this.value)}><input type=button value='검색' class=ibtn onclick=location.href='help.php?help='+encodeURIComponent(g_hv.value)>",modal: true});
}
</script>
<input type="button" class="icon_skill" style="background-repeat:no-repeat;background:url('http://iwgame.cafe24.com/images/도움말.gif')" onclick="g_help();g_hv.focus();">
<br><a onclick="location.href='?now_c=1'" onmouseover="this.style.color='blue'"><font size="1">접속중:96</font></a><br><font size="2">낚시터</font></div>
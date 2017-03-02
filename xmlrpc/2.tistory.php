<!-- 왜 이것을 include하냐면  이 xml rpc 방식의 meta Weblog api는 다른 플랫폼 간에  원격 메소드 호출을 하는데, 그런 rpc를 xml 포맷으로 서버에 요청하는 것입니다. 우리는  클라이언트가 되어 XML로 작성해서 요청하게 되는건데  라이브러리 같은걸로 요청 용이하게 메소드를 제공하고 있어서 php함수로 호출이 가능하고  큰 그림으로 보자면 서버측의 metaWeblog.newPost 메소드 호출을 요청한다고 보시면 됩니다. -->
<?
include_once 'xmlrpc.inc';

 function newPost($title, $desc) {
	define('BLOGID', 'API아이디');	// 티스토리 관리 가서 blog api 사용 체크하고 나오는 블로그 아이디
	define('USER', '로그인ID');	// 로그인 아이디
	define('PASSWORD', ' 패스워드');		// 로그인 비번
	// Create Client
	$cl = new xmlrpc_client("http://블로그.tistory.com/api");	// 블로그 주소/api

	// Prepare request
	$req = new xmlrpcmsg('metaWeblog.newPost');
	// Parameters one by one
	$req->addParam ( new xmlrpcval ( BLOGID, 'string')); // blog_id
	$req->addParam ( new xmlrpcval ( USER, 'string')); // user_name
	$req->addParam ( new xmlrpcval ( PASSWORD, 'string' )); // password

	$struct = new xmlrpcval(
		array (
		      "title" => new xmlrpcval ( $title, 'string' ),
		      "description" => new xmlrpcval ( $desc, 'string'),
		), "struct"
	);

	$req->addParam($struct);
	$req->addParam(new xmlrpcval(true, 'boolean'));
	$req->request_charset_encoding='UTF-8';
  // $newPost = array("categories" => new xmlrpcval(array(new xmlrpcval("책", "string")), "array"), "description" => new xmlrpcval("본문", "string"), "title" => new xmlrpcval("title", "string"), "dateCreated" => new xmlrpcval("20070303T19:20:30+09:00", "string"));
	return $ans=$cl->send($req);
};

newPost($get_title, $get_desc);

# 이상하게 작동하지 않는다

import xmlrpc.client

proxy = xmlrpc.client.ServerProxy("https://newsleaf.tistory.com/api")
struct = {
    "title": "XML-RPC를 이용한 글 올리기 테스트",
    "description": "글에 들어갈 내용입니다."
}

proxy.metaWeblog.newPost("forteleaf@gmail.com", "2280790" "888XATTM", struct, True)

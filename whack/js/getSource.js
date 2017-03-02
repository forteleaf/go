// 접속된 홈페이지의 document_body의 내용을 추천한다.
function get_source(document_body) {
   return document_body.innerHTML;
}

// 텍스트 내용은 get_source 를 통해 얻어와서 내부 메세지로 송신한다.
chrome.extension.sendMessage({
   action: "getSource",
   source: get_source(document.body)
});
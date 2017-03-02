// 저장을 하면 실시간으로 방영한다.

// 스크립트를 밀어넣는 코드
// Listener를 등록한다.
chrome.extension.onMessage.addListener(function (request, sender) {
      console.log("Content.js has been loaded");
      if (request.action == "getSource") {
            document.body.innerText = request.source;
      }
});
// 2. executeScript 함수가 작동하여 getSoruce.js 를 로딩 된 페이지에 주입한다.
function onWindowLoad() {
      chrome.tabs.executeScript(null, {
            file: "js/getSource.js"
      }, function () {
            if (chrome.extension.lastError) {
                  document.body.innerText = 'There was an error injecting script : \n' + chrome.extension.lastError.message;
            }
      });
}
// console.log 는 입력할 수 없다.
function sayHello() {
      // document.body.innerText = "Hello, World!";
      var text = $("naver_menu").val();
      document.body.innerText = text;
}

window.onload = sayHello

// storage
function saveChanges() {
      // Get a value saved in a form.
      var theValue = textarea.value;
      if (!theValue) {
            message('Error:No balue specified');
            return
      }
      // save it using the Chrome extension storage API.
      chrome.storage.local.set({
            'value': theValue
      }, function () {
            //Notify that we saved.
            message('Settings saved');
      });
}

function loadChange() {
      var keys = ["key1", "key2"];
      chrome.storage.sync.get(keys, function (items) {
            var value1 = items.key1;
            var value2 = items.key2;
      });
}
// window.onload = sayHello;
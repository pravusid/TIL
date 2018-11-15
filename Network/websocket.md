# WebSocket

## 웹 소켓(Web Socket)

상호작용하는 웹 서비스를 위해 WebSocket, FlashSocket, AJAX Long Polling, AJAX Multi part Streaming, IFrame, JSONP Polling 등 다양한 방법을 사용했다.
그러나 이러한 방식은 브라우저가 HTTP 요청를 보내고 웹 서버가 이 요청에 대한 HTTP 응답를 보내는 단방향 메세지 교환 '규칙'을 변경하지 않고 구현한 방식이다.

보다 쉽게 상호작용하는 웹 페이지를 만들려면 브라우저와 웹 서버 사이에 bi-directional full-duplex communication이 필요하다.
그래서 HTML5 표준안의 일부로 WebSocket API(이후 WebSocket)가 등장했다.

WebSocket은 그 이름에서 알 수 있듯이 소켓을 이용하여 자유롭게 데이터를 주고 받을 수 있다.
즉 기존의 요청-응답 관계 방식보다 더 쉽게 데이터를 교환할 수 있다.

## WebSocket 프로토콜

표준 WebSocket의 API는 W3C에서 관장하고, 프로토콜은 IETF(Internet Engineering Task Force)에서 관장한다.
그리고 WebSocket은 다른 HTTP 요청과 마찬가지로 80번 포트를 통해 웹 서버에 연결한다.

HTTP 프로토콜의 버전은 1.1이지만 다음 헤더의 예에서 볼 수 있듯이, Upgrade 헤더를 사용하여 웹 서버에 요청한다.
당연한 이야기겠지만 클라이언트인 브라우저와 마찬가지로 웹 서버도 WebSocket 기능을 지원해야한다.

```text
GET ws://websocket.example.com/ HTTP/1.1
Origin: http://example.com
Connection: Upgrade
Host: websocket.example.com
Upgrade: websocket
```

브라우저는 "Upgrade: WebSocket" 헤더 등과 함께 랜덤하게 생성한 키를 서버에 보낸다.
웹 서버는 이 키를 바탕으로 토큰을 생성한 후 브라우저에 돌려준다. 이런 과정으로 WebSocket 핸드쉐이킹이 이루어진다.

그 뒤 Protocol Overhead 방식으로 웹 서버와 브라우저가 데이터를 주고 받는다.
Protocol Overhead 방식은 여러 TCP 커넥션을 생성하지 않고 하나의 80번 포트 TCP 커넥션을 이용하고,
별도의 헤더 등으로 논리적인 데이터 흐름 단위를 이용하여 여러 개의 커넥션을 맺는 효과를 내는 방식이다.

이런 방식을 사용하기 때문에 방화벽이 있는 환경에서도 무리 없이 WebSocket을 사용할 수 있다.

```js
// Create a new WebSocket with an encrypted connection.
var socket = new WebSocket('ws://websocket.example.com');
```

> WebSocket URL은 ws 스키마를 사용한다. HTTPS와 동일한 보안 WebSocket 연결에 사용되는 wss도 있다.

## WebSocket 서버

Node.js의 경우 웹소켓 요청을 다음과 같이 받을 수 있다

```js
// WebSocket implementation
var WebSocketServer = require('websocket').server;
var http = require('http');

var server = http.createServer(function(request, response) {
  // process HTTP request.
});
server.listen(1337, function() { });

// create the server
wsServer = new WebSocketServer({
  httpServer: server
});

// WebSocket server
wsServer.on('request', function(request) {
  var connection = request.accept(null, request.origin);

  // This is the most important callback for us, we'll handle
  // all messages from users here.
  connection.on('message', function(message) {
      // Process WebSocket message
  });

  connection.on('close', function(connection) {
    // Connection closes
  });
});
```

이때 서버 응답은 다음과 같다

```text
HTTP/1.1 101 Switching Protocols
Date: Wed, 25 Oct 2017 10:07:34 GMT
Connection: Upgrade
Upgrade: WebSocket
```

# CORS (Cross Origin Resource Sharing)

<https://developer.mozilla.org/ko/docs/Web/HTTP/Access_control_CORS>

최초 리소스의 도메인과 다른 도메인에서 리소스가 요청될 경우 해당 리소스는 cross-origin HTTP 요청에 의해 요청된다.

예를 들어, `http://a.com`으로부터 전송되는 HTML 페이지가 `<img>`src 속성을 통해 `http://b.com/image.jpg`를 요청하는 경우가 있다.

그러나 보안 상의 이유로, 브라우저들은 **스크립트 내에서** 초기화되는 cross-origin HTTP 요청을 제한한다.
XMLHttpRequest는 same-origin 정책을 따르기에, XHR을 사용하는 웹 애플리케이션은 자신과 동일한 도메인으로 HTTP 요청을 보내는 것만 가능했다.

시대가 변함에 따라, XMLHttpRequest가 cross-domain 요청을 할 수 있도록 요청하여 CORS가 등장하였다.

CORS는 웹 서버에게 cross-domain 접근 제어권을 부여한다.
모던 브라우저들은 cross-origin HTTP 요청의 위험성을 완화시키기 위해 (XMLHttpRequest와 같은) API 컨테이너 내에서 CORS를 사용한다.

CORS는 일반적으로 다음의 요청에 대해 활성화 된다.

- cross-site의 방식 내에서의 XMLHttpRequest API 호출
- (CSS 내 @font-face에서의 cross-domain 폰트 사용을 위한) 웹 폰트
- WebGL 텍스쳐
- drawImage를 사용해 캔버스에 드로잉되는 이미지/비디오 프레임들
- (CSSOM 접근을 위한) 스타일시트
- (활성화된 예외 보고를 위한) 스크립트

## 접근제어 시나리오

### 간단한 요청

간단한 cross-site 요청은 다음 조건들을 만족한다

- 허용된 HTTP 메소드
  - GET
  - HEAD
  - POST

- 사용자 에이전트에 의해 자동으로 설정되는 헤더(Connection, User-Agent 등)를 제외하고, 수동 설정이 허용되는 헤더들
  - Accept
  - Accept-Language
  - Content-Language
  - Content-Type (but note the additional requirements below)
  - DPR
  - Downlink
  - Save-Data
  - Viewport-Width
  - Width

- Content-Type 헤더에 대해 허용되는 값
  - application/x-www-form-urlencoded
  - multipart/form-data
  - text/plain

간단한 요청의 경우 클라이언트와 서버간의 리퀘스트/리스폰스는 다음과 같은 형태일 것이다

```text
-- Request (http://foo.example)
1: GET /resources/public-data/ HTTP/1.1
2: Host: bar.other
3: User-Agent: Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.5; en-US; rv:1. 9.1b3pre) Gecko/20081130 Minefield/3.1b3pre
4: Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
5: Accept-Language: en-us,en;q=0.5
6: Accept-Encoding: gzip,deflate
7: Accept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7
8: Connection: keep-alive
9: Referer: http://foo.example/examples/access-control/simpleXSInvocation.html
10: Origin: http://foo.example

-- Response (http://bar.other)
13: HTTP/1.1 200 OK
14: Date: Mon, 01 Dec 2008 00:23:53 GMT
15: Server: Apache/2.0.61
16: Access-Control-Allow-Origin: *
17: Keep-Alive: timeout=2, max=100
18: Connection: Keep-Alive
19: Transfer-Encoding: chunked
20: Content-Type: application/xml
21:
22: [XML Data]
```

서버는 리소스가 cross-site 방식 으로 모든 도메인으로부터 접근 가능하다는 것을 의미하는 `Access-Control-Allow-Origin: *` 값을 응답한다.

만약 리소스 서버의 소유자가 리소스에 대한 접근을 `http://foo.example`에게만 허용하려면 다음으로 응답해야 한다:
`Access-Control-Allow-Origin: http://foo.example`

### preflighted"(사전 전달) 요청

사전 전달 요청은 실제 요청을 전송하기 앞서 안전 여부확인을 위해, 다른 도메인에 있는 리소스로 `OPTIONS` 메서드로 HTTP 요청을 전송한다.

특히, 다음과 같은 경우에 요청이 사전 전달된다:

- GET, HEAD 혹은 POST 외의 메서드를 사용하는 경우
  - PUT
  - DELETE
  - CONNECT
  - OPTIONS
  - TRACE
  - PATCH

- POST 메서드를 사용한 요청이 간단한 요청의 Content-Type 범위를 벗어나는 경우
  - application/x-www-form-urlencoded, multipart/form-data, or text/plain 이외의 다른 값을 가진 Content-Type
  - 즉, POST 요청이 서버에 application/xml 혹은 text/xml을 사용하여 XML 페이로드를 전송하게 되면, 요청은 사전 전달 됨

- 요청 내에 커스텀 헤더를 설정한 경우(예를 들자면 요청이 X-PINGOTHER와 같은 헤더를 사용하는 경우)

클라이언트와 서버 간의 전체적인 송수신은 다음과 같을 것이다

```text
-- preflighted Request
OPTIONS /resources/post-here/ HTTP/1.1
Host: bar.other
User-Agent: Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130 Minefield/3.1b3pre
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Language: en-us,en;q=0.5
Accept-Encoding: gzip,deflate
Accept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7
Connection: keep-alive
Origin: http://foo.example
Access-Control-Request-Method: POST
Access-Control-Request-Headers: X-PINGOTHER

-- Response to preflighted Request
HTTP/1.1 200 OK
Date: Mon, 01 Dec 2008 01:15:39 GMT
Server: Apache/2.0.61 (Unix)
Access-Control-Allow-Origin: http://foo.example
Access-Control-Allow-Methods: POST, GET, OPTIONS
Access-Control-Allow-Headers: X-PINGOTHER
Access-Control-Max-Age: 1728000
Vary: Accept-Encoding, Origin
Content-Encoding: gzip
Content-Length: 0
Keep-Alive: timeout=2, max=100
Connection: Keep-Alive
Content-Type: text/plain

-- Request
POST /resources/post-here/ HTTP/1.1
Host: bar.other
User-Agent: Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130 Minefield/3.1b3pre
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Language: en-us,en;q=0.5
Accept-Encoding: gzip,deflate
Accept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7
Connection: keep-alive
X-PINGOTHER: pingpong
Content-Type: text/xml; charset=UTF-8
Referer: http://foo.example/examples/preflightInvocation.html
Content-Length: 55
Origin: http://foo.example
Pragma: no-cache
Cache-Control: no-cache

<?xml version="1.0"?><person><name>Arun</name></person>

-- Response
HTTP/1.1 200 OK
Date: Mon, 01 Dec 2008 01:15:40 GMT
Server: Apache/2.0.61 (Unix)
Access-Control-Allow-Origin: http://foo.example
Vary: Accept-Encoding, Origin
Content-Encoding: gzip
Content-Length: 235
Keep-Alive: timeout=2, max=99
Connection: Keep-Alive
Content-Type: text/plain

[Some GZIP'd payload]
```

OPTIONS 메소드는 서버로 여분의 정보를 전달할 때 사용되고, 리소스 변경을 요청할 수 없다.
OPTIONS 요청과 함께, 다음 두 개의 요청이 전송되었다.

- `Access-Control-Request-Method: POST`: 실제 요청이 전달될 경우, POST 요청 메서드와 함께 전송될 것임을 서버에게 알린다
- `Access-Control-Request-Headers`: X-PINGOTHER: 헤더는 실제 요청이 전달된 경우, X-PINGOTHER라는 커스텀 헤더와 함께 전송될 것임을 서버에게 알린다

서버는 OPTIONS 요청에 포함된 위의 값을 근거로 요청을 받아들일지 여부를 응답하게 된다.

위에서는 요청 메서드(POST)와 요청 헤더(X-PINGOTHER)를 받아들일 수 있다는 것을 서버가 응답한다

- `Access-Control-Allow-Origin: http://foo.example`: 접근가능한 도메인
- `Access-Control-Allow-Methods: POST, GET, OPTIONS`: 리소스에 접근하기 위해 실행 가능한 메서드 목록
- `Access-Control-Allow-Headers: X-PINGOTHER, Content-Type`: 실제 요청과 함께 사용되도록 허가된 헤더인지 확인
- `Access-Control-Max-Age: 86400`: 사전 전달 요청에 대한 응답이 다른 사전 전달 요청을 전송하지 않고 얼마동안 캐시되어 있는지 (단위:초)

### credentialed(인증된) 요청

credentialed(인증된) 요청은 HTTP 쿠키와 HTTP Authentication 정보를 확인하게 한다.

기본적으로, cross-site XHR에서 자격 증명을 위한 정보를 전송하지 않는다.
만약 자격증명과 함께 요청이 호출되도록 하려면 xhr객체에서 `withCredentials=true`값을 사용해야 한다

클라이언트와 서버 간의 전체적인 송수신은 다음과 같을 것이다

```text
-- Request
GET /resources/access-control-with-credentials/ HTTP/1.1
Host: bar.other
User-Agent: Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130 Minefield/3.1b3pre
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Language: en-us,en;q=0.5
Accept-Encoding: gzip,deflate
Accept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.7
Connection: keep-alive
Referer: http://foo.example/examples/credential.html
Origin: http://foo.example
Cookie: pageAccess=2

-- Response
HTTP/1.1 200 OK
Date: Mon, 01 Dec 2008 01:34:52 GMT
Server: Apache/2.0.61 (Unix) PHP/4.4.7 mod_ssl/2.0.61 OpenSSL/0.9.7e mod_fastcgi/2.4.2 DAV/2 SVN/1.4.2
X-Powered-By: PHP/5.2.6
Access-Control-Allow-Origin: http://foo.example
Access-Control-Allow-Credentials: true
Cache-Control: no-cache
Pragma: no-cache
Set-Cookie: pageAccess=3; expires=Wed, 31-Dec-2008 01:34:53 GMT
Vary: Accept-Encoding, Origin
Content-Encoding: gzip
Content-Length: 106
Keep-Alive: timeout=2, max=100
Connection: Keep-Alive
Content-Type: text/plain

[text/plain payload]
```

서버에서는 인증된 요청에 응답하는 경우 `Access-Control-Allow-Credentials: true` 헤더와 함께 응답해야 한다.

또한 서버는 도메인을 특정해야만 하며, 와일드 카드를 사용할 수 없다. (`Access-Control-Allow-Origin: 도메인`)

## HTTP 응답헤더

CORS 스펙에 의해 서버가 응답하는 HTTP 헤더

### Access-Control-Allow-Origin

반환된 리소스는 다음과 같은 문법을 가진 Access-Control-Allow-Origin 헤더를 포함해야 한다

`Access-Control-Allow-Origin: <origin> | *`

origin 파라메터는 리소스에 접근하는 URI을 특정한다.
자격 증명 없는 요청에 대해, 서버는 와일드 카드(`*`)를 사용하며, 리소스에 접근하는 출처는 어떤 것이든 허용한다.

서버가 와일드카드 대신 하나의 origin을 특정하는 경우,
서버는 Origin request header의 값과 다르다는것을 알리기 위해 Origin을 응답의 Vary 헤더에 포함해야 한다.

### Access-Control-Expose-Headers

기본적으로 브라우저에게 노출이 되는 HTTP Response Header

- Cache-Control
- Content-Language
- Content-Type
- Expires
- Last-Modified
- Pragma

`Access-Control-Expose-Headers: X-My-Custom-Header, X-Another-Custom-Header`
브라우저에 노출되지 않는 X-My-Custom-Header 그리고 X-Another-Custom-Header 헤더를 허용한다.

### Access-Control-Max-Age

이 헤더는 사전 전달의 결과가 얼마나 오랫동안 캐시될 수 있는지를 설정함

`Access-Control-Max-Age: <delta-seconds>`

### Access-Control-Allow-Credentials

- xhr의 credentials 플래그가 true인 경우 요청에 대한 응답을 내보낼 수 있는지 없는지를 전달함
- preflight 요청에 대한 응답에 포함된 경우, 실제 요청이 자격 증명을 사용할 수 있는지를 전달함

자격 증명과 함께 리소스에 대한 요청이 이루어졌으나 위의 헤더가 함께 반환되지 않은 경우, 응답은 무시된다.

`Access-Control-Allow-Credentials: true | false`

### Access-Control-Allow-Methods

preflight 요청에 대한 응답으로, 리소스에 접근하는 경우 허용된 메서드 혹은 메서드들을 전달함

`Access-Control-Allow-Methods: <method>[, <method>]*`

### Access-Control-Allow-Headers

preflight 요청에 대한 응답으로, 실제 요청이 이루어지는 경우 어떤 HTTP 헤더가 사용될 수 있는지 전달함

`Access-Control-Allow-Headers: <field-name>[, <field-name>]*`

## HTTP 요청 헤더

이 섹션은 cross-origin 공유 기능을 사용하여 HTTP 요청을 호출할 때 클라이언트가 사용할 수 있는 헤더들

> cross-site XHR 기능을 사용하는 개발자들은 cross-origin 공유 요청의 헤더를 직접 설정할 필요가 없다

### Origin

cross-site 접근 요청 혹은 사전 전달 요청의 출처

`Origin: <origin>`

- 어떠한 경로 정보도 포함하지 않으며, 서버 이름만을 포함한다
- origin는 빈 문자열이 될 수 있다: 소스가 data URL인 경우
- 어떤 접근 제어 요청에서든 ORIGIN 헤더는 항상 전송된다

### Access-Control-Request-Method

사전 전달 요청 시, 실제 요청이 일어나는 경우 어떤 HTTP 메서드가 사용될 것인지 서버에 알리기 위해 사용됨

`Access-Control-Request-Method: <method>`

### Access-Control-Request-Headers

사전 전달 요청 시, 실제 요청이 일어나는 경우 어떤 HTTP 헤더가 사용될 것인지 서버에 알리기 위해 사용됨

`Access-Control-Request-Headers: <field-name>[, <field-name>]*`

# HTTP Basic

그림으로 배우는 HTTP network basic을 읽고 정리

## Web과 Network 기초

### 웹은 HTTP로 나타낸다

웹브라우저(클라이언트)와 서버는 HyperText Transfer Protocol로 통신한다

### HTTP의 역사

1989년 3월 CERN에서 고안됨

HyperText에 의해 상호 참조할 수 있는 World Wide Web을 구성하는 기술

- HTML(HyperText Markup Language): 문서 기술 언어
- HTTP: 문서 전송 프로토콜
- URL(Uniform Resource Locator): 문서 주소 지정

### 웹의 성장

- 1993년 1월 NCSA에서 Mosaic 브라우저가 개발됨
- 1994년 12월 넷스케이프 네비게이터 출시
- 1995년 마이크로소프트 인터넷 익스플로러 출시
- 이 무렵 웹서버 Apache 등장
- HTML 2.0 발행

### HTTP의 정체

- HTTP/0.9: 1990년 정식사양 이전
- HTTP/1.0: 1996년 5월 정식사양 (RFC1945) 발행
- HTTP/1.1: 1997년 1월 공개되었고 현재 가장 많이 사용됨 (RFC2068 -> RFC2616)
- HTTP/2.0: 차세대 버전, 현재 보급중

### 네트워크: TCP/IP

#### TCP/IP는 프로토콜의 집합

컴퓨터와 네트워크 기기가 상호통신을 하기위해 규약이 있어야 함. 이를 프로토콜이라 한다.

인터넷과 관련된 프로토콜을 모은것을 TCP/IP라고 한다.
TCP와 IP 프로토콜을 가리켜 TCP/IP라 부르기도 하지만,
IP 프로토콜을 사용한 통신에서 사용되고 있는 프로토콜을 총칭해서 TCP/IP라는 이름이 사용되고 있다.

#### TCP/IP의 계층

TCP/IP는 애플리케이션 계층, 트랜스포트계층, 네트워크 계층, 링크 계층의 4계층으로 구성되어 있다.

##### 애플리케이션 계층

애플리케이션 계층은 유저에게 제공되는 애플리케이션에서 사용하는 통신

예를 들어 FTP, HTTP등이 해당됨

##### 트랜스포트 계층

애플리케이션 계층에 네트워크로 접속되어 있는 컴퓨터사이 데이터 흐름을 제공함

TCP(Transmission Control Protocol)와 UDP(User Data Protocol)의 두가지가 있음

애플리케이션 계층에서 받은 데이터를 전송/수신하기 위해 변환함

##### 네트워크 계층

네트워크 계층은 네트워크 상에서 패킷의 이동을 다룬다.

수신지의 MAC 주소등을 통해 네트워크 기기간 길을 결정하는 것이 해당 계층의 역할이다.

##### 링크 계층

네트워크에 접속하는 하드웨어적인 면을 다룬다.

운영체제가 하드웨어를 제어하므로, 디바이스 드라이버, 네트워크 인터페이스카드(NIC)를 포함한다.
또한 케이블 및 커텍트와 같은 물리적인 영역 또한 포함한다.

#### TCP/IP 통신의 흐름

TCP/IP로 통신할 때 계층을 순서대로 거쳐 상대와 통신한다.
송신과 수신에서 거쳐가는 계층은 역순이다.

클라이언트

1. 애플리케이션: HTTP 클라이언트 (HTTP 데이터)
2. 트랜스포트: TCP (+ TCP 헤더)
3. 네트워크: IP (+ IP 헤더)
4. 링크: 네트워크 (+ Ethernet 헤더)

서버

1. 링크: 네트워크 (- Ethernet 헤더)
2. 네트워크: IP (- IP 헤더)
3. 트랜스포트: TCP (- TCP 헤더)
4. 애플리케이션: HTTP 서버 (HTTP 데이터)

#### IP

Internet Protocol은 네트워크 계층이다.

IP 프로토콜은 다음의 정보로 개개의 패킷을 상대방에게 전달한다

- IP주소 : 각 노드에 부여된 주소
- MAC주소: 네트워크 카드에 할당된 고유주소

IP 통신은 MAC주소에 의존해서 통신을 한다.
보통 네트워크간 중계를 통해 상대에게 도착하게 되는데 이때 ARP(Address Resolution Protocol) 프로토콜이 사용된다.

ARP 프로토콜은 수신지의 IP주소를 바탕으로 MAC주소를 조사할 수 있다.

목적지까지 중계 하는 도중 네트워크 기기는 대략의 목적지만 알고 전송되는데 이를 라우팅이라 한다.

#### TCP

TCP는 트랜스포트 계층에 해당되는데, 신뢰성 있는 바이트 스트림 서비스를 제공한다.

TCP는 용량이 큰 데이터를 TCP 세그먼트라 불리는 단위패킷으로 분해하여 관리하고, 도착여부의 정확성을 확인한다.

TCP에서 상대에게 확실하게 데이터를 보내기 위해서 Three way handshaking 이라는 방법을 사용한다.
이는 패킷을 보내고 바로 끝내지 않고, 전송여부 확인을 위해 'SYN'과 'ACK'라는 TCP 플래그를 사용하는 것이다.

- 송신측에서 최초 SYN 플래그로 상대에게 접속함과 동시에 패킷을 보냄
- 수신측에서 SYN/ACK 플래그로 송신측에 접속함과 동시에 패킷을 수신한 사실을 보냄
- 송신측이 ACK 플래그를 보내 패킷 교환이 완료되었음을 알림
- 통신이 끊어지면 같은 과정을 반복하며 패킷을 재전송한다

이외에도 신뢰성을 보증하기 위한 다양한 시스템이 있다

### DNS

DNS(Domain Name System)는 HTTP 같은 응용 계층 시스템에서 도메인 이름과 IP 주소 이름확인을 제공한다.

### URI와 URL

URI(Uniform Resource Identifier)

- Uniform: 통일된 서식으로 여러 종류의 리소스 지정방법을 동일 맥락에서 구분
- Resource: 리소스는 식별 가능한 모든것이다.
- Identifier: 식별가능한 것을 참조하는 오브젝트

URI는 리소스를 식별하기 위한 문자열 전반을 나타내는데, URL은 리소스의 네트워크상 위치를 나타낸다.
즉, URL은 URI의 subset이다.

절대 URI(필요 정보 전체를 지정) 포맷

- 스키마: `http://`
- 자격정보(옵션): `user:pass@`
- 서버주소: `www.example.com`
- 서버포트(기본 80): `:80`
- 계층적 경로: `/dir/index`
- 쿼리 문자열: `?id=1`
- 프래그먼트 식별자: `#ch1`

## HTTP 프로토콜의 구조

HTTP는 클라이언트와 서버 간에 통신을 한다

HTTP는 클라이언트로부터 Request가 송신되며 그 결과 서버로부터 Response가 돌아온다.
반드시 클라이언트측으로 부터 통신이 시작된다.

### HTTP는 상태를 유지하지 않는 프로토콜

HTTP는 상태를 계속 유지하지 않는 stateless 프로토콜이다.

HTTP에서는 새로운 request가 보내질 때마다 새로운 response가 생성된다.
프로토콜로서 과거의 request나 response 정보를 전혀 가지고 있지 않다.

이는 많은 데이터를 빠르고 확실허게 처리하기 위한 scalability를 확보하기 위해 간단히 설계되어 있는 것이다.

그러나 stateless 특징으로 처리하기 어려운 일이 증가하였다.

HTTP/1.1은 여전히 stateless이지만 Cookie라는 기술이 도입되었다.

### request URI로 리소스 식별

HTTP는 URI를 사용하여 인터넷상의 리소스를 지정한다.

클라이언트는 리소스를 호출할 때 마다 request를 송샌할 때에 request URI를 포함해야 한다.

- 모든 URI를 request URI에 포함: `GET http://google.com/index HTTP/1.1`
- Host 헤더 필드에 네트워크 로케이션 포함: `GET /index HTTP/1.1 Host:google.com`

특정 리소스가 아닌 서버 자신에게 request를 송신하는 경우 `*` 지정 가능

HTTP 서버가 지원하는 메소드 확인: `OPTIONS * HTTP/1.1`

### 서버에 목적을 말하는 HTTP 메소드

- GET 메소드: request URI로 식별된 리소스를 가져올 수 있도록 요구함
- POST 메소드: 엔티티를 전송하기 위해서 사용됨
- PUT 메소드: request중 포함된 엔티티를 request URI로 지정한 곳에 보존하도록 요구함 (파일 전송)
- HEAD 메소드: GET과 같은 기능이지만 body는 돌려주지 않는다. URI 유효성과 리소스 갱신시간 확인등의 목적으로 쓰임
- DELETE 메소드: 파일을 삭제하기 위해서 사용됨(PUT 메소드와 반대)
- OPTIONS 메소드: request URI로 지정한 리소스가 제공하고 있는 메소드를 조사하기 위해 사용
- TRACE 메소드: Web 서버에 접속해서 자신에게 통신을 되돌려 받는 loop-back을 발생시킨다. (XST등의 문제로 보통사용되지 않음)
- CONNECT 메소드: 프록시에 터널 접속 확립을 요구하여 TCP 통신을 터널링 시키기 위해 사용됨

### 지속연결

HTTP 초기버전에서는 매 통신마다 TCP에 의해 연결과 종료를 해야 할 필요가 있었다.

HTTP/1.1과 일부의 HTTP/1.0에서는 이러한 문제를 해결하기 위해 Persistent Connection이라는 방법을 고안하였다.
이 경우 어느 한쪽에서 명시적으로 연결을 종료하지 않는 이상 TCP 연결을 유지한다.

지속 연결은 여러 request를 보낼 수 있다록 HTTP pipelining을 가능하게 한다.

파이프라인화로 request 송신 후에 response를 수신할 때까지 대기하지 않고, 다음 request를 보낼수 있다.

### 쿠키를 사용한 상태관리

HTTP는 stateless 프로토콜이므로 과거 상태를 관리하지 못하는 단점이 존재한다.
이를 해결하기 위해 쿠키라는 시스템이 도입되었다.

쿠키는 서버에서 response로 보내진 Set-Cookie라는 헤더 필드에 의해 쿠키를 클라이언트에 보존하게 된다.
다음번에 클라이언트가 같은 서버로 리퀘스트를 보낼 때, 자동으로 쿠키 값을 넣어서 송신한다.

서버는 클라이언트가 보내온 쿠키로 클라이언트를 식별하고 서버상의 기록을 확인할 수 있다.

쿠키를 가지고 있지 않은 상태의 request

```text
GET /reader HTTP/1.1
Host: www.google.com
```

서버가 쿠키를 발행한 response

```text
HTTP/1.1 200 OK
Date: Thu, 12 Jul 2012 07:12:20 GMT
Server: Apache
<Set-Cookie: sid=1342077140226742; path=/;expires=Wed, => 10-Oct-12 07:12:20 GMT>
Content-Type: text/plain; charset=UTF-8
```

클라이언트가 보관하던 쿠키를 포함한 request

```text
GET /image HTTP/1.1
Host: www.google.com
Cookie: sid=1342077140226742
```

## HTTP 메시지

HTTP에서 교환하는 정보는 HTTP 메시지라고 하는데 request 메시지와 response 메시지로 나뉜다

### 리퀘스트 메시지와 리스폰스 메시지 구조

- 메시지 헤더
  - 리퀘스트 라인: 리퀘스트에 사용하는 메소드와 리퀘스트 URI, HTTP 버전 포함
  - 상태 라인: 리스폰스 결과를 나타내는 상태코드와 설명, HTTP 버전 포함
  - 헤더 필드: 리퀘스트와 리스폰스의 여러 조건과 속성들을 나타내는 각종 헤더필드 포함
    - 일반 헤더필드
    - 리퀘스트 헤더 필드
    - 리스폰스 헤더 필드
    - 엔티티 헤더 필드
  - 그 외: HTTP의 RFC에 없는 헤더 필드 (쿠키 ...)가 포함
- 개행 문자(CR + LF))
- 메시지 바디: 전송되는 데이터 (반드시 존재하지는 않음)

### 인코딩

HTTP로 데이터를 전송할 경우 인코딩을 통해 전송 효율을 높일 수도 있다.
단, 인코딩 처리를 해야하므로 CPU 등의 리소스는 보다 많이 소비하게 된다.

#### 메시지 바디와 엔티티 바디

메시지: HTTP 통신의 기본 단위로 Octet(8bit) Sequence로 구성된다

엔티티: 리퀘스트와 리스폰스의 payload로 전송되는 정보로 엔티티 헤더 필드와 엔티티 바디로 구성됨

HTTP 메시지 바디의 역할은 리퀘스트/리스폰스에 관한 엔티티 바디를 운반하는 것.
기본적으로 메시지 바디와 엔티티 바디는 같지만 전송코딩이 적용된 경우에는 달라진다.

#### Contents Coding

Contents Codings는 엔티티에 적용하는 인코딩인데, 엔티티 정보를 유지한 채로 압축한다.

다음과 같은 콘텐츠 압축이 있다

- gzip(GNU zip)
- compress(UNIX 표준 압축)
- deflate(zlib)
- identity(인코딩 없음)

#### Chunked Transfer Coding

큰 데이터를 전송하는 경우 데이터를 분할할 수 있는데, 엔티티 바디를 분할하는 기능을 청크 전송 코딩이라 한다.
이 경우 다음 청크 사이즈를 16진수를 사용해 단락을 표시하고 바디 끝에 CR+LF를 기록한다.

#### Multipart

MIME(Multipurpose Internet Mail Extension)는 메일로 텍스트, 영상, 이미지 같은 다른 데이터를 다루기 위한 기능이다.

MIME는 이미지 등의 바이너리 데이터를 ASCII 문자열에 인코딩하는 방법, 데이터 종류를 나타내는 방법등을 규정한다.

MIME의 타입인 Multipart는 하나의 메시지 바디 내부에 엔티티를 어러개 포함시켜 보낸다.

HTTP 메시지로 multipart를 사용할 때는 Content-Type 헤더 필드를 사용한다.
멀티파트 각각의 엔티티를 구분하기 위해 "boundary" 문자열을 사용하고 앞에는 `--`를 삽입한다.

멀티파트는 파트마다 헤더필드가 포함된다. 또한 파트 중간에 파트를 내부에 포함할 수도 있다.

##### multipart/form-data

Web의 Form으로 부터 파일 업로드에 사용됨
(user types "Joe Blow" in the name field)

```text
Content-Type: multipart/form-data; boundary=AaB03x

--AaB03x
Content-Disposition: form-data; name="field1"

Joe Blow
---AaB03x
Content-Disposition: form-data; name="pics"; filename="file1.txt"
Content-Type: text/plain

...(file1.txt 데이터)...
--AaB03x
```

##### multipart/byteranges

상태코드 206(Partial Content) response message가 복수 범위의 내용을 포함하는 때 사용

```text
HTTP/1.1 206 Partial Content
Date: Fri, 13 Jul 2012 02:45:26 GMT
Last-Modified: Fri, 31 Aug 2007 02:02:20 GMT
Content-Type: multipart/byteranges: boundary=THIS_STRING_SEPARATES

--THIS_STRING_SEPARATES

Content-Type: application/pdf
Content-Range: bytes 500-999/8000

...(범위내의 데이터)...
--THIS_STRING_SEPARATES

Content-Type: application/pdf
Content-Range: bytes 7000-7999/8000

...(범위내의 데이터)...
--THIS_STRING_SEPARATES
```

### 레인지 리퀘스트

엔티티의 범위를 지정해서 요청을 보내는 것을 Range Request라 한다.
Range Header Field를 사용해서 리소스 레인지를 지정한다.

`Ranage: bytes = -3000, 5001-10000`: 처음부터 3000바이트 까지, 그리고 5001 ~ 10000 바이트

- Range Request에 대한 Response는 상태코드 206(Partial Content)이 돌아온다.
- 복수 범위의 Range Request에 대해서는 `multipart/byteranges` response가 돌아온다.
- 서버가 range request를 지원하지 않으면 완전한 엔티티와 함께 상태코드 200이 돌아온다

### Content Negotiation

Content Negotiation은 클라이언트와 서버가 제공하는 리소스를 언어와 문자세트, 인코딩 방식을 기준으로 구분한다.

판단기준은 다음과 같은 헤더필드 내용에 근거한다.

- Accept
- Accept-Charset
- Accept-Encoding
- Accept-Language
- Content-Language

Content Negotiation의 종류는 다음이 있다.

- Server-driven Negotiation: 서버측에서 리퀘스트 헤더필드의 내용을 참고해서 처리함
- Agent-driven Negotiation: 클라이언트 측에서 처리, OS 종류, 브라우저 종류, User Agent에 따라 전환하는 것등이 포함됨
- Transparent Negotiation: 서버와 에이전트 방식을 혼합한 것, 각각 Negotiation을 한다

## HTTP 상태 코드

### Request 결과

클라이언트로부터 서버로 리퀘스트를 보낼 때 서버에서 처리된 상태를 알려주는 것이 상태코드이다.

리스폰스 클래스는 5개가 정의되어 있다

- 1xx (Informational): 리퀘스트를 받아 처리중
- 2xx (Success): 리퀘스트를 정상적으로 처리함
- 3xx (Redirection): 리퀘스트 완료를 위해 추가 동작 필요
- 4xx (Client Error): 클랄이언트 원인으로 서버에서 리퀘스트 이해 불능
- 5xx (Server Error): 서버에서 리퀘스트 처리 실패

### 2xx (Success)

- 200 OK: 클라이언트가 보낸 리퀘스트를 서버가 정상 처리했을 경우, HTTP 메소드에 따라 되돌아오는 데이터는 다르다
- 204 No Content: 리스폰스에 entity body를 포함하지 않는 경우
- 206 Partial Content: Range가 지정된 리퀘스트에 의해서 Content-Range로 지정된 범위의 엔티티가 포함

### 3xx (Redirection)

- 301 Moved Permanently: 요청된 리소스에 새로운 URI가 부여되어, 이후로 변경된 URI를 사용해야 함을 리스폰스한다
- 302 Found: 요청된 리소스에 새로운 URI가 부여되어 있지만, 301과 다르게 일시적인 변경이다
- 303 See Other: 요청된 리소스는 다른 URI이므로 GET 메소드를 통해 얻어야 함을 나타낸다
- 304 Not Modified: 조건부 리퀘스트에서 접근은 가능하나 조건을 만족하지 않는 경우이다. 비어있는 리스폰스 바디를 반환한다.
- 307 Temporary Redirect: 302 Found와 동일하다

301, 302, 303 리스폰스 코드가 오면 대부분의 브라우저는 POST를 GET으로 변경하고 엔티티 바디를 삭제하여 리퀘스트를 재송신한다.
303 코드에서 GET 메소드를 사용할 것을 명시하고 있으나 다른코드에서도 그런 방식으로 작동한다.

### 4xx (Client Error)

- 400 Bad Request: 리퀘스트 구문이 잘못되었음을 나타낸다
- 401 Unauthorized: 송신한 리퀘스트에 HTTP 인증이 필요하다는 것을 나타낸다(첫 번째), 인증에 실패했음을 표시한다(두 번째)
- 403 Forbidden: 요청한 리소스의 접근이 거부되었음을 나타낸다. 거부된 이유를 엔티티 바디에 포함해 반환할 수 있다.
- 404 Not Found: 요청한 리소스가 서버상에 없다는 것을 나타낸다. (혹은 이유 없이 거부할 때)

### 5xx (Server Error)

- 500 Internal Server Error: 서버에서 요청을 처리하는 도중 에러가 발생한 경우
- 503 Service Unavailable: 일시적으로 서버를 사용할 수 없어 리퀘스트가 처리되지 않는 경우. Retry-After 헤더필드 값을 반환할 수 있다.

## HTTP 웹서버

### 가상 호스트

HTTP/1.1에서는 하나의 HTTP 서버에 여러개의 웹사이트를 실행할 수 있다.

같은 IP 주소에서 다른 호스트/도메인명을 가진 여러개의 웹사이트가 실행되는 가상호스트 시스템이 있으므로,
HTTP 리퀘스트를 보내는 경우 호스트/도메인명을 완전히 포함한 URI를 지정하거나, Host 헤더필드에서 지정해야 한다.

### 프록시, 게이트웨이, 터널

#### 프록시

프록시 서버는 클라이언트로부터 받은 리퀘스트 URI를 변경하지 않고 다음 리소스를 가지고 있는 서버에 보낸다.

실제로 리소스를 가진 서버를 Origin Server라고 한다.
오리진 서버로부터 돌아온 리스폰스는 프록시 서버를 경유하여 클라이언트로 돌아간다.

HTTP 통신에서는 여러대의 프록시 서버를 경유하는 것도 가능하다.
프록시 서버를 통해 리퀘스트와 리스폰스를 중계할 때는 Via 헤더 필드에 경유한 호스트 정보를 추가해야 한다.

##### Cashing Proxy

프록시에 같은 리소스에 대한 요청이 온 경우 오리진 서버로부터 리소스를 획득하지 않고, 캐시를 리스폰스로 되돌려 줄 수 있다.

##### Transparent Proxy

리퀘스트와 리스폰스를 중계할 때 메시지 변경을 하지않는 프록시를 투명 프록시라한다.

반대로 메시지에 변경을 가하는 프록시를 비투과 프록시라고 한다.

#### 게이트웨이

게이트웨이는 프록시와 유사하지만 다음에 있는 서버가 HTTP 서버 이외의 서비스를 제공하는 경우이다.

#### 터널

더널은 다른 서버와의 통신경로를 확립한다. 이 때 HTTP 리퀘스트를 해석하지 않고 다음서버로 중계한다.

### Cache

캐시는 프록시 서버와 클라이언트의 로컬 디스크에 보관된 리소스 사본을 가리킨다.
캐시를 사용하면 통신량과 통신시간을 효율적으로 관리할 수 있다.

캐시 서버는 프록시 서버의 하나로 캐싱 프록시로 분류된다.

#### 캐시 유효기간

오리진 서버의 리소스가 갱신되는 경우 캐시는 유효하지 않게 된다.

따라서 캐시를 활용하더라도 클라이언트의 요구나 캐시 유효기간등에 의해
오리진 서버에 유효성을 확인하거나 새로운 리소스를 획득하는 경우가 있다.

#### 클라이언트 캐시

클라이언트 측의 웹 브라우저 역시 캐시를 가지고 있을 수 있다.
마찬가지로 캐시 유효성에 의해 다시 데이터를 획득하러 갈 수도 있다.

## HTTP Header

### HTTP 메시지 헤더

HTTP 리퀘스트와 리스폰스에는 반드시 메시지 헤더가 포함되어 있다.

#### 리퀘스트 HTTP 메시지

메시지 헤더는 메소드, URI, HTTP 버전, HTTP 헤더필드 등으로 구성되어 있다.

#### 리스폰스 HTTP 메시지

메시지 헤더는 HTTP버전, 상태코드, HTTP 헤더필드 등으로 구성되어 있다.

### HTTP 헤더 필드

HTTP 헤더 필드: `헤더필드명:헤더필드값`

하나의 헤더 필드가 여러 개의 필드 값을 가질 수 있으며 쉽표 `,`로 구분한다

#### 헤더 필드의 종류

- General Header Fields: 리퀘스트/리스폰스 메시지 두 곳다 사용됨
- Request Header Fields: 클라이언트에서 서버 방향, 리퀘스트 부가정보, 클라이언트 정보, 리스폰스 콘텐츠에 관한 우선순위 등...
- Response Header Fields: 서버에서 클라이언트 방향, 리스폰스 정보와 서버 정보, 클라이언트의 추가 정보 요구 등...
- Entity Header Fields: 리퀘스트/리스폰스 메시지에 포함된 엔티티에 사용되는 헤더로 콘텐츠 갱신 시간 등에 관한 정보를 부가

#### General Header Fields

##### Cache-Control

- 디렉티브로 불리는 명령을 사용하여 캐싱동작 지정

- 디렉티브는 파리미터가 있을 수도/없을 수도 있으며, 여러개를 지정하는 경우 `,`로 구분

- 캐시 리퀘스트 디렉티브 (디렉티브/파라미터/설명)
  - `no-cache`: orogin 서버에 강제적 재검증
  - `no-store`: 캐시는 리퀘스트/리스폰스의 일부분을 보존하면 안됨
  - `max-age=초`: 필수: 리스폰스 최대 보존시간, `Expires`헤더보다 우선함
  - `max-state=[초]`: 생략가능: 기한이 지난 리스폰스 수신, 최대 혀용기간 지정 가능
  - `min-fresh=초`: 필수: 지정한 시간 이상의 유효기간이 남은 리소스를 요청함
  - `no-transform`: 프록시 캐시는 엔티티 바디의 미디어 타입을 변경해서는 안됨
  - `only-if-cached`: 목적 리소스가 로컬 캐시에 있는 경우에만 리스폰스를 반환하도록 요구함

- 캐시 리스폰스 디렉티브 (디렉티브/파라미터/설명)
  - `public`: 유저에게 돌려줄 수 있는 리스폰스 캐시 가능
  - `private`: 생략가능: 특정 유저에 대해서만 리스폰스 (public과 반대)
  - `no-cache`: 생략가능: origin 서버에 유효성 재확인 없이 캐시 사용불가, 파라미터로 캐시할 수 없는 헤더필드를 명시할 수 있음
  - `no-store`: 캐시는 리퀘스트/리스폰스 일부분을 보존하면 안됨
  - `no-transform`: 프록시는 미디어 타입을 변경해서는 안됨
  - `must-revalidate`: 캐시 가능하지만 오리진 서버에 리소스 재확인 요구, 리퀘스트의 `max-state` 헤더를 무시함
  - `proxy-revalidate`: 모든 캐시서버에 대해 이후의 리퀘스트로 리스폰스를 반환할 때 반드시 유효성 재확인 요구
  - `max-age=초`: 필수: 리스폰스 최대 보존시간, `Expires`헤더보다 우선함
  - `s-max-age=초`: 필수: 여러 유저가 이용하는 공유 캐시 서버 리스폰스 최대 보존시간 (이경우 `Expires`와 `max-age`는 무시됨)

- 확장 토큰: `cache-extension`: 디렉티브를 해석할 수 있는 서버로 보낼경우만 유효

##### Connection

- 프록시에 더 이상 전송하지 않는 헤더필드(Hop-by-hop 헤더) 지정

- 지속적 접속(Keep-alive) 관리
  - HTTP/1.1에서는 Keep-alive가 디폴트임
  - 서버에서 명시적으로 접속을 끊고 싶을 때 `Connection`헤더 필드에 `Close`를 지정함

##### Date

메시지 생성 날짜를 표기하며, HTTP/1.1의 날짜 포맷은 RFC1123에 지정되어 있음

`Date: Tue, 03 Jul 2012 04:40:59 GMT`

##### Pragma

Pragma 필드는 HTTP/1.0과의 호환을 위해서만 정의되어 있음

`Pragma: no-cache`만 사용 가능하며 클라이언트 리퀘스트에서만 사용한다.
중간 서버들에 캐시된 리소스의 리스폰스가 필요없음을 알린다.

##### Trailer

HTTP/1.1에 구현되어있는 청크 전송 인코딩을 사용하고 있는 경우,
메시지 바디 뒤에 기술되어 있는 헤더 필드를 미리 전달할 수 있다.

```text
...
Trailer: Expires
...
... 메시지 바디 ...
0
Expires: Tue, 28 Sep 2004 23:59:59 GMT
```

##### Transfer-Encoding

메시지 바디의 전송 코딩 형식을 지정하는 경우 사용된다

`Transfer-Encoding: chunked`

##### Upgrade

HTTP 및 다른 프로토콜의 새로운 버전이 통신에 이용되는 경우 사용된다.

리퀘스트에서 Upgrade 헤더 필드에 명시한 프로토콜 사용을 요청한다.
요청은 인접한 서버에만 적용되므로 `Connection: Upgrade`도 함께 사용한다.

```text
GET /index.html HTTP/1.1
Upgrade: TLS/1.0
Connection: Upgrade
```

서버에서는 상태코드 101 `Switching Protocols` 리스폰스로 응답한다.

```text
HTTP/1.1 101 Switching Protocols
Upgrade: TLS/1.0, HTTP/1.1
Connection: Upgrade
```

##### Via

클라이언트-서버간의 리퀘스트/리스폰스 메시지의 경로를 알기 위해서 사용됨

프록시나 게이트웨이는 자신의 서버정보를 Via 헤더 필드에 추가한 뒤 메시지를 전송한다.
(traceroute나 메일의 Received Header와 유사)

Via 헤더 필드는 메시지 추적과 리퀘스트 루프회피등에 사용되므로 프록시를 경유하는 경우 반드시 붙여야 할 필요가 있다.

```text
GET / HTTP/1.1

GET / HTTP/1.1
Via: 1.0 gw.hackr.jr (Squid/3.1)

GET / HTTP/1.1
Via: 1.0 gw.hackr.jr (Squid/3.1), 1.1 a1.example.com (Squid/2.7)
```

##### Warning

Warning 헤더는 HTTP/1.0 Retry-After라 변경된 것으로, 리스폰스에 관한 추가정보를 전달한다.

`Warning: [경고코드][경고한호스트:포트번호]"[경고문]" ([날짜])`

HTTP/1.1에는 7개의 경고코드가 정의되어 있다(권장사항)

- 110: Response is state: 프록시가 유효기간이 지난 리소스 반환
- 111: Revalidation failed: 프록시가 리소스의 유효성 재확인에 실패함
- 112: Disconnection Operation: 프록시의 네트워크가 연결되어있지 않다
- 113: Heuristic expiration: 캐시 유효기한을 경과한 리스폰스
- 199: Miscellaneous warning: 임의 경고문
- 214: Transformation applied: 프록시가 인코딩/미디어 타입에 대응하여 처리한 경우
- 288: Miscellaneous persistent warning: 임의 경고문

#### Request Header Fields

- Accept: 유저 에이전트가 처리 가능한 미디어 타입
- Accept-Charset: 문자셋 우선 순위
- Accept-Encoding: 콘텐츠 인코딩 우선 순위
- Accept-Language: 자연어 우선 순위
- Authorization: 웹 인증을 위한 정보
- Expect: 서버에 대한 특정 동작의 기대
- From: 유저의 메일 주소
- Host: 요구된 리소스의 호스트
- If-Match: 엔티티 태그의 비교
- If-None-Match: 엔티티 태그의 비교(If-Match의 반대)
- If-Range: 리소스가 갱신되지 않은 경우에 엔티티의 바이트 범위 요구를 송신
- If-Modified-Since: 리소스의 갱신 시간 비교
- If-Unmodified-Since: 리소스의 갱신시간 비교(If-Modified-Since의 반대)
- Max-Forwards: 최대 전송 홉 수
- Proxy-Authorization: 프록시 서버의 클라이언트 인증을 위한 정보
- Range: 엔티티 바이트 범위 요구
- Referer: 리퀘스트중의 URI를 취등하는 곳
- TE: 전송 인코딩 우선순위
- User-Agent: HTTP 클라이언트 정보

#### Response Header Fields

- Accept-Ranges: 바이트 단위의 요구를 수신할 수 있는지 여부
- Age: 리소스의 지정 경과 시간
- Etag: 리소스를 특정하기 위한 정보
- Location: 클라이언트를 지정한 URI에 리다이렉트
- Proxy-Authenticate: 프록시 서버의 클라이언트 인증을 위한 정보
- Retry-After: 리퀘스트 재시행의 타이밍 요구
- Server: HTTP 서버 정보
- Vary: 프록시 서버에 대한 캐시 관리 정보
- WWW-Authenticate: 서버의 클라이언트 인증을 위한 정보

#### Entity Header Fields

- Allow: 리소스가 제공하는 HTTP 메소드
- Content-Encoding: 엔티티 바디에 적용되는 콘텐츠 인코딩
- Content-Language: 엔티티의 자연어
- Content-Length: 엔티티 바디의 사이즈(byte)
- Content-Location: 리소스에 대응하는 대체 URI
- Content-MD5: 엔티티 바디의 메시지 다이제스트
- Content-Range: 엔티티 바디의 범위 위치
- Content-Type: 엔티티 바디의 미디어 타입
- Expires: 엔티티 바디의 유효기한 날짜
- Last-Modified: 리소스의 최종 갱신 날짜

#### HTTP/1.1 이외의 헤더필드

HTTP 헤더 필드는 RFC2616에 정의된 47종류만 있는 것은 아니다.

Set-Cookie, Content-Disposition ...

비표준 헤더필드는 RFC4229 HTTP Header Field Registrations에 정리되어 있다

#### E2E 헤더와 Hop-by-hop 헤더

HTTP 헤더 필드는 캐시와 비캐시 프록시의 동작을 정의하기 위해서 두 가지 카테고리로 분류되어 있다.

##### E2E 헤더

E2E 분류 헤더는 리퀘스트나 리스폰스의 최종 수신자에게 전송된다.

##### Hop-by-hop 헤더

이 카테고리의 헤더는 한 번 전송에 대해서만 유효하고 캐시와 프록시에를 통하면 전송되지 않는 것도 있다.

사용되는 Hop-by-hop 헤더는 Connection 헤더 필드에 열거해야 한다.

HTTP/1.1의 Hop-by-hop 헤더에는 다음과 같은 것이 있다.
다음 8개의 헤더필드 이외에는 모두 E2E 헤더로 분류된다.

- Connection
- Keep-Alive
- Proxy-Authenticate
- Proxy-Authorization
- Trailer
- TE
- Transfer-Encoding
- Upgrade

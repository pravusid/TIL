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

- 개행 문자(CR + LF)

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

`Range: bytes = -3000, 5001-10000`: 처음부터 3000바이트 까지, 그리고 5001 ~ 10000 바이트

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

##### Accept

`Accept: text/html, application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8`

유저 에이전트에 처리할 수 있는 미디어 타입과 타입 우선순위를 전달하기 위해 사용됨

미디어 타입 지정은 `타입/서브타입`으로 한번에 여러 개를 설정할 수 있음

미디어 타입에 우선순위를 붙이려면 세미콜론(`;`)으로 구분하고 `q=1(기본값(최대)`로 품질지수를 표기한다.

- 텍스트 파일

  - text/html, text/plain, text/css ...
  - application/xhtml+xml, application/xml, application/json ...

- 이미지 파일

  - image/jpeg, image/gif, image/png ...

- 동영상 파일

  - vedio/mpeg, video/quicktime ...

- 바이너리 파일

  - application/octet-stream, application/zip

##### Accept-Charset

`Accept-Charset:iso-8859-5, unicode-1-1:q+0.8`

유저 에이전트에서 문자셋의 상대적 우선순위를 전달하기 위해서 사용된다.
여러개를 지정할 수 있으며, 품질지수에 의해 상대적 우선순위를 표시한다.

##### Accept-Encoding

`Accept-Encoding: gzip, deflate`

유저 에이전트가 처리할 수 있는 콘텐츠 코딩의 우선순위를 전달한다.

콘텐츠 코딩은 한 번에 여러개를 지정할 수 있으며, 품질지수에 의해서 상대적인 우선순위를 표시한다.
또한 `*`를 지정하면 모든 인코딩 포맷을 가리킨다.

- gzip: 파일 압축 GNU zip에서 생성된 인코딩 포맷(RFC1952)으로 LZ77 부호와 32비트 CRC 사용
- compress: UNIX 압축 compress의 인코딩 포맷으로 LZW다
- deflate: Zlib(RFC1950) 포맷과 deflate 압축 알고리다 의해 만들어진 인코딩 포맷 조합
- identity: 압축/변형을 하지 않는 디폴트 인코딩 포맷다

##### Accept-Language

`Accept-Language: ko-kr, en-us;q=0.7,en;q=o.3`

유저 에이전트가 처리할 수 있는 자연어 세트와 세트의 우선순위를 전달한다.

자연어 세트는 ㅎ나번에 여러개를 지정할 수 있으며, 품질지수에 의해 상대적인 우선순위를 나타낸다.

##### Authorization

`Authorization: Basic dWVUB3NIbjpwYXNzd29yZA==`

유저 에이전트의 인증 정보(크리덴셜 값)를 전달하기 위해 사용됨.

##### Expect

`Expect: 100-continue`

클라이언트가 서버에 특정 동작 요구를 전달한다.
기대하고 있는 요구에 서버가 응답하지 못하는 경우 417 Expectation Failed를 반환한다.

##### From

`From: info@hackr.jp`

유저 에이전트를 사용하고 있는 유저의 메일 주소를 전달한다.

##### Host

`Host: www.google.com`

리퀘스트한 리소스의 인터넷 호스트와 포트번호를 전달한다.
1대의 서버에 복수도메인 할당이 가능하므로, Host 헤더 필드는 HTTP/1.1에서 유일한 필수 헤더이다.

##### If-Match

`If-Match: "etag"`

조건부 리퀘스트의 하나로 서버 상의 리소스를 특정하기 위해서 엔티티 태그(ETag) 값을 전달한다.

서버는 If-Match 필드 값과 리소스 ETag 값이 일치하는 경우에만 리퀘스트를 받아들인다.
만약 일치하지 않는경우 상태코드 412 Precondition Failed를 반환한다.

If-Match 필드값으로 `*`를 지정하면 ETag 값과 상관없이 리소스가 존재하면 리퀘스트를 처리한다.

##### If-None-Match

조건부 리퀘스트의 하나로 If-Match와 반대로 동작한다.

GET과 HEAD 메소드에서 If-None-Match 헤더 필드를 사용하면 최신 리소스를 요구하는 것이 되므로
If-Modified-Since 헤더 필드를 사용하는 것과 비슷해진다.

##### If-Modified-Since

`If-Modified-Since: Thu, 15 Apr 2004 00:00:00 GMT`

조건부 리퀘스트의 하나로 리소스 갱신날짜가 필드값 시점보다 최신인경우 리퀘스트를 받아들인다.
리소스가 outdated인 경우 상태코드 304 Not Modified 리스폰스를 반환한다.

##### If-Unmodified-Since

If-Modified-Since 헤더 필드와 반대로 동작한다.

지정된 리소스가 필드 값 시점 이전에 갱신된 리퀘스트만 받는다.

##### If-Range

조건부 리퀘스트의 하나로 If-Range로 지정한 필드값(ETag or 날짜)과 지정 리소스의 Etag or 날짜가 일치하면 Range 리퀘스트로 처리한다.
일치하지 않는 경우 리소스 전체를 반환한다.

만약 서버의 리소스가 갱신되어 있는 경우, If-Range 헤더 필드를 사용하지 않은 Range 리퀘스트라면 무효한 요청이되어
상태코드 412 Precondition Failed를 반환하고 클라이언트에게 다시 리퀘스트를 요청하게 된다.

If-Range를 사용하는경우 갱신된 리소스의 경우 전체를 반환하므로 통신과정이 줄어들게 된다.

##### Max-Forwards

`Max-Forwards: 10`

TRACE 혹은 OPTIONS 메소드에 의해 리퀘스트를 할 때 거쳐갈 최대 서버 수를 10진 정수로 지정한다.

따라서 서버는 다음 서버에 리퀘스트를 전송할 때 Max-Forwards 갑셍서 1을 빼서 보내고,
값이 0인 리퀘스트를 받은 경우 리스폰스를 반환한다.

##### Proxy-Authorization

`Proxy-Authorization: Basic dGIwOjkpNLAGfFY5`

프록시 서버에서 인증요구를 받아들인 때 인증에 필요한 클라이언트의 정보를 전달한다.

클라이언트와 서버의 경우에 사용되는 Authorization 헤더 필드의 역할이 클라이언트와 프록시 사이에서 이루저이는 것이다.

##### Range

`Range: bytes=5001-10000`

리소스의 일부분만 취득하는 Range 리퀘스트를 할 때 지정 범위를 전달한다.

서버가 Range 헤더필드가 있는 리퀘스트를 처리할 수 있는 경우에는 상태코드 206 Partial Content 리스폰스를 반환하고,
처리할 수 없는 경우 상태코드 200 OK 리스폰스와 함께 리소스 전체를 반환한다.

##### Referer

리퀘스트가 발생한 본래 리소스의 URI를 전달한다.
브라우저 주소창에 직접 URI를 입력한 경우 전달되지 않는다.

리소스의 URI 쿼리에 민감정보가 포함되어 있는경우 Referer를 통해 해당 정보가 보내질 수 있다.

##### TE

`TE; gzip, deflate;q=0.5`

리스폰스로 받을 수 있는 전송 코딩의 형식과 우선순위를 전달한다.

Accept-Encoding 헤더 필드와 유사하지만, 전송 코딩에 적용된다는 점이 다르다.

TE 헤더 필드는 전송 코딩의 지정 외에 Trailer를 동반하는 Chunk 전송 인코딩 형식을 지정할 수 있다.
이 경우 `TE; Trailers`와 같이 표기한다.

##### User-Agent

`User-Agent:Mozilla/5.0 (Windows NT 6.1) AppleWebKit/535.19 (KHTML,like Gecko) Chrome/18.0.1025.162 Safari/535.195)`

리퀘스트를 생성한 브라우저와 유저 에이전트의 정보를 전달하기위한 필드이다.

로봇의 리퀘스트는 로봇 엔진의 책임자 메일주소가 들어있기도 하다.
또는 프록시를 경유한 리퀘스트의 경우 프록시 서버의 이름등이 표시되어 있기도 하다.

#### Response Header Fields

서버 측으로부터 클라이언트 측으로 송신되는 리스폰스 메시지에 적용된 헤더로
리스폰스의 부가정보나 서버의 정보, 클라이언트에 부가정보 요구등을 나타냄

##### Accept-Ranges

`Accept-Range: bytes`

서버가 리소스의 일부분만 지정해서 취득할 수 있는 Range 리퀘스트를 받을 수 있는지 여부전달

가능할 경우 `bytes`, 불가능할 경우 `none`

##### Age

`Age: 600`

얼마나 오래 전에 오리진 서버에서 리스폰스가 생성되었는지 전달한다.

리스폰스한 서버가 프록시 서버면 Age 헤더 필드는 필수 값이다

##### ETag

`ETag: "82e22293907ce725faf67773957acd12"`

엔티티 태그라고 불리며 리소스를 특정하기 위한 문자열을 전달한다.
서버는 리소스마다 ETag 값을 할당한다.

리소스가 갱신되는 경우 ETag 값도 갱싱되어야 할 필요가 있다.
ETag 값은 룰이 있지는 않고 서버에 따라 다양한 값을 할당한다.

URI만으로 캐시했던 리소스를 특정하기 어려운 경우 ETag를 참조해서 리소스를 특정할 수 있다.

- 강한 ETag 값

  - 엔티티가 아주 조금 다르더라도 반드시 값이 변함
  - `ETag: "Usagi-1234'`

- 약한 ETag 값

  - 약한 ETag 값은 리소스가 같다는 것만 나타냄
  - `ETag: W/"usagi-1234"`

##### Location

`Location: http://www.usagidesign.jp/sample.html`

리스폰스의 수신자에 대해서 Request-URI 이외의 리소스 액세스를 유도하는 경우 사용됨

기본적으로 3xx Redirection 리스폰스에 대해서 리다이렉트 URI를 기술함

##### Proxy-Authenticate

`Proxy-Authenticate: Basic realm="Usagidesign Auth"`

프록시 서버의 인증요구를 클라이언트에 전달한다.
클라이언트와 서버의 경우 WWW-Authorization 필드와 같은역할을 한다.

##### Retry-After

`Retry-After: 120`

클라이언트가 일정 시간 후에 리퀘스트를 다시 시행해야 하는지를 전달함
주로 상태코드 503 Service Unavailable, 3xx Redirect 리스폰스와 함께 사용된다

값으로는 날짜나 리스폰스 이후 지연시간(seconds)을 지정할 수 있다.

##### Server

`Server: Apache/2.2.17(Unix)`

서버에 설치되어 있는 HTTP 서버 소프트웨어를 전달함

##### Vary

`Vary: Accept-Language`

오리진 서버가 프록시 서버에 로컬 캐시를 사용하는 방법에 대한 지시를 전달함

오리진 서버로부터 Vary에 지정되었던 리스폰스를 받아들인 프록시 서버는
캐시된 때의 리퀘스트와 같은 Vary 헤더필드를 가진 리퀘스트에 대해서만 캐시를 반환한다.

리퀘스트에 Vary에 지정된 헤더 필드가 다른경우, 오리진 서버로 부터 리소스를 취득해야 한다.

##### WWW-Authenticate: 서버의 클라이언트 인증을 위한 정보

`WWW-Authenticate: Basic realm="Usagidesign Auth"`

HTTP 액세스 인증에 사용되고, Request-URI의 리소스에 적용할 수 있는
인증 스키마 (Basic / Digest)와 파라미터를 나타내는 challenge를 전달한다.

WWW-Authenticate 헤더필드는 상태코드 401 Unauthorized 리스폰스에 반드시 포함된다.

#### Entity Header Fields

##### Allow

`Allow: GET, HEAD`

Request-URI에 지정된 리소스가 제공하는 메소드의 목록

서버가 받을 수 없는 메소드를 수신한 경우에는 상태코드 405 Method Not Allowed 리스폰스와 함께
수신가능한 메소드 목록을 기술한 Allow 헤더필드를 반환한다.

##### Content-Encoding

`Content-Encoding: gzip`

엔티티 바디의 콘텐츠 코딩형식을 전달한다.

주로 4가지 콘텐츠 코딩 형식이 사용된다. (Accept-Encoding 헤더 필드 항목과 동일)

- Gzip
- Compress
- Deflate
- Identity

##### Content-Language

`Content-Language: en`

엔티티 바디에 사용된 자연어를 전달함

##### Content-Length

`Content-Length: 15000`

엔티티 바디의 크기(bytes)를 전달한다.

엔티티 바디에 전송 코딩이 사용된 경우 Content-Length 필드를 사용하면 안된다. (RFC2616 4.4)

##### Content-Location

`Content-Location: http://pravusid.kr/index.html`

메시지 바디에 대응하는 URI를 전달함

Location이 리다이렉션의 대상(혹은 새롭게 만들어진 문서의 URL)을 가르키는데 반해,
Content-Location은 더 이상의 컨텐츠 협상없이, 리소스 접근에 필요한 직접적인 URL을 지시함.

##### Content-MD5

`Content-MD5: OGFkZDUwNGVhNGY3N2MxMDIwZmQ4NTBmY2IyTY==`

메시지 바디가 변경되지 않고 도착했는지 확인하기 위해 MD5 해시값을 Base64 인코딩하여 전달한다.

리퀘스트시 콘텐츠와 함께 MD5 해시값도 변조하여 보내는 것도 가능하므로 서버측에서 원래 의도한 데이터인지 여부와는 관계없다.

##### Content-Range

`Content-Range: bytes 5001-10000/10000`

범위를 지정해서 일부분만을 리퀘스트하는 Range 리퀘스트에 대해서 리스폰스 할 때 사용된다.

필드 값은 반환하는 엔티티 범위와 전체사이즈이다.

##### Content-Type

`Content-Type: text/html; charset=UTF-8`

엔티티 바디에 포함되는 오브젝트의 미디어 타입을 전달한다. 필드 값은 타입/서브타입으로 구성된다.

##### Expires

`Expires: Wed, 04 Jul 2012 08:26:05 GMT`

리소스의 유효 기한 날짜를 전달한다.

캐시 서버가 Expires 헤더필드를 포함한 리소스를 수신한 경우 필드 값으로 지정된 날짜까지 리스폰스의 복사본으로 리퀘스트에 응답한다.
지정된 날짜가 지나면 리퀘스트가 왔을 때 오리진 서버에 리소스를 요청한다.

`Cache-Control` 헤더필드의 `max-age` 디렉티브는 `Expires` 헤더필드보다 우선순위가 높다.

##### Last-Modified

`Last-Modified: Wed, 23 May 2012 09:59:55 GMT`

리소스가 마지막으로 갱신되었던 날짜 정보를 전달한다.

#### 쿠키를 위한 헤더 필드

쿠키는 HTTP/1.1의 사양인 RFC2616에 포함된 것은 아니지만 널리 사용되고 있다.

현재 사용되는 쿠기의 사양은 RFC6256이다.
쿠키에 관련한 헤더필드는 다음의 것이 사용되고 있다.

- Set-Cookie: 리스폰스: 상태 관리 개시를 위한 쿠키 정보
- Cookie: 리퀘스트: 서버에서 수신한 쿠키 정보

##### Set-Cookie

`Set-Cookie: status-enable; expires=Tue, 05 Jul 2011 07:26:31 GMT; =>path=/;domain=.hack.jp;`

서버가 클라이언트에 대해서 상태 관리를 시작할 때 정보를 전달한다.

- NAME=VALUE

  - 쿠키에 부여된 이름과 값

- Expires=DATE

  - 브라우저가 쿠키를 송출할 수 있는 유효기한
  - 값이 없으면 세션이 유지되고 있는 동안만 유효함
  - 한번 송출한 클라이언트 쿠키를 명시적으로 삭제하는 방법은 없고, 덮어쓰는 것으로 실질적으로 삭제가능

- Path=PATH

  - 쿠키 송출 범위를 특정 디렉토리(도메인의)에 한정할 수 있다

- Domain=도메인명

  - domain을 지정하고 값을 비교할 때 후방 일치여부를 확인함

- Secure

  - `Set-Cookie: name=value; secure`
  - HTTPS에서 열렸을 때만 쿠키를 송출함

- HttpOnly

  - `Set-Cookie: name=value; HttpOnly`
  - 자바스크립트를 경유하여 쿠키를 취득하지 못하도록 함
  - XSS로부터 쿠키 도청을 막기 위함이 목적이다

##### Cookie

`Cookie: status=enable`

클라이언트가 서버로부터 수신한 쿠키를 이후의 리퀘스트에 포함해서 전달함

쿠키를 여러개 수신하고 있다면, 여러개의 쿠키를 보내는 것도 가능하다.

### 이외의 헤더 필드

HTTP 헤더 필드는 독자적으로 확장할 수 있다.
여러곳에서 지원하는 확장 헤더 필드들은 다음과 같다.

#### X-frame-Option

`X-Frame-Option: DENY`

다른 웹 사이트의 프레임에서 표시를 제어하는 리스폰스 헤더로, Click jacking 공격 방지를 목적으로 한다.

필드값으로 설정할 수 있는 값은 다음과 같다

- DENY: 거부
- SAMEORIGIN: Top-level-browsing-context가 일치하는 경우에만 허가

#### X-XSS-Protection

`X-XSS-Protection: 1`

크로스 사이트 스크립팅(XSS) 대책으로 브라우저의 XSS 보호기능을 제어하는 리스폰스 헤더이다

헤더 필드에 지정할 수 있는 값은 다음과 같다

- 0: XSS 필터를 무효화
- 1: XSS 필터를 유효화

#### DNT

`DNT: 1`

Do Not Track의 약어로 개인정보 수집을 거부하는 의사를 표시하는 리퀘스트 헤더이다.

헤더 필드에 지정할 수 있는 값은 다음과 같다

- 0: 트래킹 동의
- 1: 트래킹 거부

#### P3P

`P3P: CP="CAO DSP LAW CURa ADMa DEVa TAIa PSAa PSDa IVAa IVDa OUR BUS IND UNI COM NAV INT"`

웹사이트 상의 프라이버시 정책에 The Platform for Privacy Preferences를 사용하기 위한 리스폰스 헤더이다.

P3P 정책은 다음 순서로 실행된다

1. P3P 정책 작성
2. P3P 정책 참조 파일을 작성하여 `/w3c/p3p.xml`에 배치
3. P3P 정책으로부터 콤팩트 정책을 작성하고 HTTP 리스폰스 헤더에 출력

## HTTPS 프로토콜

### HTTP의 약점

#### 평문이므로 도청 가능

TCP/IP의 통신내용은 통신 경로 도중에 패킷을 수집하는 것만으로도 엿볼 수 있다.

정보를 보호하기 위해 통신을 암호화 하거나 콘텐츠를 암호화 할 수 있다.

통신 암호화는 SSL(Secure Socket Layer)이나 TLS(Transport Layer Security)이라는
다른 프로토콜을 조합하여 가능하다.
이렇게 SSL을 조합한 HTTP를 HTTPS(HTTP Secure)라 부른다.

콘텐츠를 암호화하기 위해서 클라이언트와 서버가 콘텐츠 암호화 복호화 기능을 수행해야한다.

#### 통신 상대를 확인하지 않음

HTTP에서는 상대가 누구인지 확인하는 처리는 없으므로 누구든지 리퀘스트를 보낼 수 있고,
리퀘스트가 오면 상대가 누구든지 리스폰스를 반환한다.

이는 다음과 같은 문제점을 유발할 수 있다

- 리퀘스트를 보낸 곳이 의도와 다르게 위장한 서버일 수 있다
- 리스폰스를 반환한 곳의 클라이언트가 리퀘스트를 보낼때 의도한 클라이언트인지 확인할 수 없다
- 통신 상대가 접근이 허가된 상대인지를 확인할 수 없다
- 어떤 곳에서 리퀘스트를 했는지 확인할 수 없다
- 의미없는 리퀘스트도 모두 수신한다 (DoS 공격일지라도)

HTTP에서는 통신 상대를 확인할 수 없지만 신뢰할 수 있는 제3자가 발행한 SSL의 증명서로 상대를 확인할 수 있다.

증명서를 이용하여 서버는 통신상대에게 통신하고자 하는 서버임을 알릴 수 있고,
클라이언트는 증명서로 본인확인을 하고 인증에 사용할 수도 있다.

#### 변조가능

HTTP는 리퀘스트나 리스폰스가 발신된 후 상대가 수신하는 사이 변조되었더라도 그 사실을 확인할 수 없다.

도중에 리퀘스트나 리스폰스를 탈취하여 변조하는 공격을 Man-in-the-Middle 공격이라고 부른다.

변조를 피해 완전성을 확인하기 위해 일반적으로 자주 사용되는 방법은,
해시값을 확인하거나 파일의 디지털 서명을 확인하는 것이다.

하지만 해시값이나 디지털 서명도 함께 변조될 수 있으므로 해결책이 될 수는 없다.

### HTTP + 암호화 + 인증 + 완전성 = HTTPS

HTTPS는 새로운 애플리케이션 계층의 프로토콜이 아니다.
HTTP 통신을 하는 소켓 부분을 SSL(Secure Socket Layer)이나 TLS(Transport Layer Security) 프로토콜로 대체한 것이다.

보통 HTTP는 직접 TCP와 통신하지만 SSL을 사용한 경우에는 HTTP는 SSL과 통신하고 SSL이 TCP와 통신하게 된다.

HTTPS는 공통키 암호화 공개키 암호 두 방식을 동시에 사용하여 암호화 한다.
공통키 암호는 안전하게 키를 교환하기 어렵고, 공개키 암호는 처리속도가 느리기 때문이다.

따라서 공통키 암호의 키를 교환할 때 공개키 암호화를 사용하고 이후로는 공통키 암호화 방식을 사용한다.

문제는 공개키가 진짜인지 아닌지를 증명할 필요도 있다는 것이다.
이 문제 해결을 위해 인증 기관(Certificate Authority)과 그 기관이 발행하는 공개키 증명서가 이용된다.

인증 기관은 클라이언트와 서버 모두가 신뢰하는 3자 기관이다.

인증 기관을 경유한 암호키 교환 순서는 다음과 같다

1. 서버의 공개키를 인증기관에 등록
2. 인증기관의 비밀키로 서버의 공개키에 디지털 서명으로 공개키 증명서를 작성 등록
3. 서버의 공개키 증명서를 입수하고 디지털 서명을 인증기관의 공개키로 검증한다
4. 서버의 공개키로 암호화 해서 메시지를 서버로 보낸다
5. 서버의 비밀키로 메시지를 복호화 한다

> 주요 인증기관의 공개키는 사전에 브라우저에 내장되어 있다

#### EV SSL 증명서

상대방이 실제 존재하는 기업인지 확인하는 역할을 하는 증명서를 EV SSL 증명서라 한다.
EV SSL 증명서로 증명된 웹사이트에 접속하면 주소창의 색이 녹색으로 변하는 것을 확인할 수 있다.

#### 클라이언트 증명서

HTTPS에서는 클라이언트 증명서도 이용할 수 있다.

그러나 클라이언트 증명서는 몇 가지 문제점이 있다.

- 유저가 클라이언트 증명서를 직접 구비해야 한다
- 유저 수 만큼 비용이 들게된다
- 클라이언트를 증명할 뿐 사용자의 존재를 증명하지는 않는다

#### 자가 인증기관

OpenSSL등을 활용하면 자체적으로 인증 기관을 구축하여 서버 증명서를 발행할 수 있다.
그러나 신뢰할 수 있는 제3자 기관이 인증하는 것이 아닌이상 증명서로 구실을 하지 못할 수 있다.

마찬가지로 메이저 인증기관이 아닌 마이너 인증기관을 이용한다면 브라우저에 따라 인증서 신뢰성 인식에 문제가 발생할 수 있다.

### HTTPS 구조

1. (C→S) Handshake: ClientHello

   - 클라이언트가 Client Hello 메시지를 보내며 SSL 통신시작
   - 메시지에는 클라이언트 SSL 버전지정, Cipher Suite(사용하는 암호화의 알고리즘, 키 사이즈 등)

2. (C←S) Hansshake: ServerHello

   - 서버가 SSL 통신이 가능한 경우 Server Hello 메시지로 응답
   - 메시지에는 SSL 버전, Cipher Suite 포함 (클라이언트에서 받은 내용에서 선택됨)

3. (C←S) Handshake: Certificate

   - 서버가 공개키 증명서가 포함된 Certificate 메시지를 송신

4. (C←S) Handshake: ServerHelloDone

   - 서버가 Server Hello Done 메시지를 송신하여 최초 SSL negotiation이 끝났음을 통지함

5. (C→S) Handshake: ClientKeyExchange

   - 최초 SSL negotiation이 종료되면 클라이언트가 Client Key Exchange 메시지로 응답
   - 메시지에는 통신을 암호화하는데 사용하는 Pre-Master secret 포함
   - 메시지는 `3`의 공개키 증명서에서 추출한 공개키로 암호화 됨

6. (C→S) ChangeCipherSpec

   - 클라이언트는 Change Cipher Spec 메시지를 송신함
   - 메시지 이후의 통신은 암호키를 사용하여 진행한다는 것을 표현하는 것

7. (C→S) Handshake: Finished

   - 클라이언트는 Finished 메시지를 송신함
   - 메시지는 접속 전체의 체크 값을 포함함
   - Negotiation이 성공했다면 서버가 이 메시지를 복호화 할 수 있다

8. (C←S) ChangeCipherSpec

   - 서버에서도 Change Cipher Spec 메시지를 송신함

9. (C←S) Handshake: Finished

   - 서버에서도 Finished 메시지를 송신함

10. (C→S) Application Data (HTTP)

    - 서버와 클라이언트의 Finished 메시지 교환이 완료되면 SSL 접속이 완료됨
    - 이후로는 애플리케이션 계층의 프로토콜(HTTP)로 통신(HTTP 리퀘스트 송신)

11. (C←S) Application Data (HTTP)

    - HTTP 리스폰스를 송신

12. (C→S) Alert: warning, close notify

    - 클라이언트가 접속을 끊으면 close notify 메시지를 송신함
    - 이후 TCP FIN 메시지를 보내 TCP 통신을 종료함

> 애플리케이션 계층의 데이터를 송신할 때 MAC(Message Authentication Code)라 불리는 메시지 다이제스트를 덧붙일 수 있다. MAC을 이용해 변조를 감지할 수 있다.

#### SSL과 TLS

SSL은 넷스케이프의 프로토콜로 현재는 IETF로 이관되었다.
SSL3.0을 기반으로하여 TLS1.0이 만들어졌고 현재 TLS1.1 TLS1.2 TLS1.3이 있다.

TLS는 SSL을 바탕으로 한 프로토콜이고 이를 총칭해서 SSL이라 부르기도 한다.

### HTTPS의 속도

HTTPS를 사용하게 되면 TCP 접속과 HTTP 리퀘스트/리스폰스외에도 SSL에 필요한 통신이 추가된다.
또한 암호화 처리를 위해서 서버와 클라이언트의 리소스를 사용하게된다.

SSL 엑셀레이터라는 하드웨어를 사용해서 이 문제를 해결하기도 한다.

## 인증

HTTP/1.1에서 사용할 수 있는 인증방식은 다음이 있다

### BASIC 인증

Basic 인증은 HTTP/1.0에 구현된 인증방식으로 일부 사용되고 있다.

Basic 인증에 사용되는 Base64 인코딩은 암호화가 아니며,
인증이 한 번 이루어지면 일반 브라우저에서 로그아웃 할 수 없다는 문제가 있다.

1. Basic 인증이 필요한 리소스에 요청이 발생하면 서버는 다음 내용을 포함해 리스폰스를 반환한다

   - 상태코드 401 Authorization Required
   - WWW-Authenticate 헤더필드에 Request-URI의 보호공간을 식별하기 위한 문자열(realm)
   - WWW-Authenticate 헤더필드에 인증 방식(Basic)

2. 상태코드 401을 수신한 클라이언트는 ID와 Password를 서버로 보낸다

   - ID와 Password를 `:`으로 연결한 문장을 `Base64`로 인코딩한다
   - Authorization 헤더필드에 인코딩한 문자열을 포함하여 리퀘스트를 보낸다

3. Authorization 헤더 필드가 포함된 리퀘스트를 받은 서버는

   - 인증 정보가 정확한지 확인한다
   - 인증 정보가 확인되면 Request-URI 리소스를 포함한 리스폰스를 반환한다

### DIGEST 인증

HTTP/1.1에 소개된 Basic 인증의 단점을 보완한 방식이다, 그러나 보안성이 높지않아 많이 사용되지 않는다.

Digest 인증에서 사용하는 챌린지 리스폰스 방식은,
최초 상대에게 인증 요구를 보내고 상대방 측에서 받은 챌린지 코드를 사용해서 리스폰스를 계산한 후 상대에게 보낸다.

1. 인증이 필요한 리소스에 요청이 발생하면 서버는 다음 내용을 포함해 리스폰스를 반환한다

   - 상태코드 401 Authorization Required
   - WWW-Authenticate 헤더필드에 Request-URI의 보호공간을 식별하기 위한 문자열(realm)
   - WWW-Authenticate 헤더필드에 챌린지 코드(nonce): 401 반환할 때마다 새로 생성됨

2. 상태코드 401을 수신한 클라이언트는 Digest 인증에 필요한 정보를 서버로 보낸다

   - Authorization 헤더필드에 username, realm, nonce, uri, response를 포함한다
   - username은 realm에서 인증가능한 사용자 이름이다
   - 프록시에 의해 Request-URI가 변경되는 경우를 대비하여 digest-uri에 복사한다
   - response는 Request-Digest라고 불리며 nonce를 활용하여 패스워드 문자열을 MD5로 계산한 것이다

3. Authorization 헤더 필드가 포함된 리퀘스트를 받은 서버는

   - 인증 정보가 정확한지 확인한다
   - 인증 정보가 확인되면 Request-URI 리소스를 포함한 리스폰스를 반환한다
   - 리스폰스의 Authentication-Info 헤더필드에 인증정보를 추가할 떄도 있다

### SSL 클라이언트 인증

인증 도중 정보가 탈취되었을 때를 방지하기 위한 대책 중 하나이다

SSL 클라이언트 인증을 위해서는 사전에 클라이언트에 클라이언트 증명서를 배포해야 한다

1. 인증이 필요한 리소스에 요청이 발생하면 서버는 클라이언트 증명서를 요구하는 Certificate Request라는 메시지를 보낸다

2. 증명서 요구를 받은 클라이언트 측에서는 클라이언트 증명서와 Client Certificate 메시지를 서버로 보낸다

3. 서버는 클라이언트 증명서를 검증하여 검증결과가 정확하다면 클라이언트의 공개키를 취득하고 HTTPS에 의한 암호를 개시한다

### 폼 베이스 인증

HTTP 프로토콜에 정의된 사양의 인증방식은 아니지만 일반적으로 사용되는 방식이다
폼 베이스 인증은 표준 사양은 없으나 일반적으로 쿠키를 사용하여 세션관리를 한다.

1. 클라이언트에서는 서버로 인증을 위해 다음 방식의 리퀘스트르 보낸다

   - ID나 패스워드 등의 자격 정보를 포함한 데이터
   - 보통은 POST 메소드가 사용되어 엔티티 바디에 정보를 저장한다
   - 입력 데이터의 송신에는 HTTPS를 이용한다

2. 서버는 유저를 식별하기 위해서 세션ID를 발행한다

   - 클라이언트로 부터 수신한 인증정보를 검증한다
   - 인증 상태를 세션ID와 연결하여 서버측에 기록한다
   - 클라이언트에 송신할 때는 Set-Cookie 헤더 필드에 세션ID(PHPSESSID / JSESSIONID ...)를 포함한다
   - 세션ID는 유저를 구별하기 위한 것으로 탈취/유추가 어렵도록 해야한다
   - XSS등의 취약성을 대비하기 위해 쿠키는 httponly 속성을 부여한다

3. 클라이언트는 서버에서 받은 세션ID를 쿠키로 저장해둔다

## HTTP의 추가 프로토콜

HTTP 규격이 만들어졌을 때의 예상과 다르게, 다방면으로 사용되면서 프로토콜의 한계를 보완하려는 움직임이 있다

### SPDY

Google이 2010년 발표한 SPDY(SPeeDY)는 HTTP의 병목현상을 해소한다는 목표로 개발되고 있다.

HTTP에서는 서버의 정보가 갱신되었는지를 알기위해서 클라이언트가 항상 서버에 확인을 요청해야 한다.
따라서 정보가 갱신되지 않은경우 불필요한 통신이 발생하게 된다.

#### HTTP의 병목

현재의 조건에서 다음과 같은 HTTP의 사양이 병목을 일으킬 수 있다

- 1개의 커넥션으로 1개의 리퀘스트만 보낼 수 있다
- 리퀘스트는 클라이언트에서만 시작할 수 있다. 리스폰스만 받는것은 불가능하다
- 리퀘스트/리스폰스 헤더를 압축하지 않은 채로 보낸다. 헤더의 정보가 많을 수록 느려진다
- 장황한 헤더를 보내며, 이를 매번 반복한다
- 데이터 압축 여부를 선택할 수 있다 (강제적이지 않다)

이를 해결 하기위해 다음의 방법을 사용할 수 있다

- Ajax로 해결: 기존의 동기식 통신에 비해 페이지의 일부만 갱신되므로 교환되는 데이터의 양이 줄어든다
- Comet으로 해결: Comet에서는 리스폰스를 보류 상태로 해 두고 서버의 콘텐츠가 갱신되었을 때 리스폰스를 반환한다

두 방식다 HTTP 프로토콜의 제약을 없앨 수 없고 나름의 단점이 존재한다.

#### SPDY 설계와 기능

SPDY는 HTTP를 완전히 바꾸는 것이아니라 TCP/IP의 애플리케이션 계층과 트랜스포트 계층 사이에 새로운 세션 계층을 추가하는 형태이다.
SPDY가 세션 계층으로 들어가서 데이터 흐름을 제어하지만 HTTP 커넥션은 존재한다.

SPDY를 사용하면 다음과 같은 기능을 사용할 수 있다

- 다중화 스트림: 단일 TCP 접속을 통해 복수의 HTTP 리퀘스트를 무제한 처리할 수 있다
- 리퀘스트 우선순위 부여: SPDY는 무제한 리퀘스트 병렬처리가 가능하고, 각 리퀘스트에 우선순위를 할당할 수 있다
- HTTP 헤더 압축: 리퀘스트/리스폰스 HTTP 헤더를 압축한다
- 서버 푸시: 서버에서 클라이언트로 데이터를 Push하는 서버푸시를 지원한다
- 서버 힌트: 서버가 클라이언트에게 리퀘스트 해야할 리소스를 제안할 수 있다

#### SPDY 결론

SPDY를 사용하려면 브라우저와 웹 서버가 이를 지원해야 한다.

또한, SPDY는 기본적으로 한 개의 도메인(IP 주소)과의 통신을 다중화 할 뿐이므로,
하나의 웹 사이트에서 복수의 도메인으로 리소스를 사용하는 경우 그 효과는 한정적이다.

### WebSocket

2011년 12월 11일 WebSocket의 단독 사양이 RFC 6455-The WebSocket Protocol로 출시되었다

#### WebSocket의 설계와 기능

웹소켓은 웹 브라우저와 웹 서버를 위한 양방향 통신 규격으로
웹소켓 프로토콜을 IETF가 책정하고 웹소켓 API를 W3C가 책정하고 있다.

주로 Ajax나 Comet에서 사용하는 XMLHttpRequest의 결점을 해결하는 것에 초점이 맞춰져 있다.

#### WebSocket 프로토콜

웹소켓은 웹 서버와 클라이언트가 한번 접속을 확립하면 그 뒤의 통신은 모두 전용프로토콜로 진행한다.

웹소켓 프로토콜의 특징은 다음과 같다

- 서버 푸시: 서버에서 클라이언트에게 데이터를 푸시할 수 있다

- 통신량 감소: 웹소켓은 한번 접속하면 유지할 수 있으므로 HTTP에 비해 빈번한 접속으로 인한 오버헤드가 없어진다

- 핸드쉐이크/리퀘스트: 웹소켓 통신을 하려면 HTTP의 Upgrade 헤더 필드를 사용해서 프로토콜을 변경하는 것으로 핸드쉐이크를 실시한다

  - Sec-WebSocket-Key 헤더: 핸드쉐이크에 필요한 키
  - Set-WebSocket-Protocol 헤더: 사용하는 서브 프로토콜(커넥션을 여러개로 구분할 때 사용)

- 핸드쉐이크/리스폰스: 웹소켓 리퀘스트에 대한 리스폰스는 상태코드 101 Switching Protocols

  - Sec-WebSocket-Accept 헤더: Sec-WebSocket-Key 헤더 값으로 부터 생성된 값이 저장된다
  - 웹소켓 연결이 확립되면 HTTP가 아닌 웹소켓 독자 데이터 프레임으로 통신한다

JavaScript에서 WebSocket 프로토콜을 사용하기위해 W3C 사양인 WebSocket 인터페이스를 사용한다

### WebDAV

Web-based Distributed Authoring and Versioning은
웹서버의 콘텐츠를 직접 복사/편집할 수 있는 HTTP/1.1 확장 프로토콜이다.

서버상의 리소스에 대해 WebDAV에서 새롭게 추가된 개념은 다음이 있다

- Collection: 여러개의 리소스를 한꺼번에 관리하기 위한 개념
- Resource: 파일이나 컬렉션을 리소스라 칭함
- Property: 리소스의 프로퍼티를 정의한 것 (`이름=값`)
- Lock: 파일을 편집할 수 없는 상태

WebDAV에서 추가된 메소드는 다음과 같다

- PROPFIND: 프로퍼티 취득
- PROPPATCH: 프로퍼티 변경
- MKCOL: 컬렉션 작성
- COPY: 리소스 및 프로퍼티 복제
- MOVE: 리소스 이동
- LOCK: 리소스 잠금
- UNLOCK: 리소스 잠금 해제

WebDAV에서 확장된 상태코드는 다음과 같다

- 102 Processing: 리퀘스트는 수신하였으나 처리중이다
- 207 Multi-Status: 복수의 상태
- 422 Unprocessable Entity: 서식은 맞지만 내용은 틀리다
- 423 Locked: 리소스가 잠겨있다
- 424 Failed Dependency: 리퀘스트와 연관된 리퀘스트가 실패했다
- 507 Insufficient Storage: 기억영역이 부족하다

## 웹 공격기술

HTTP는 보안과 관련된 기능이 없으며,
리퀘스트가 탈취되어 쿼리, 폼, 헤더, 쿠키등의 정보에 포함된 취약점 공격을 받을 수 있다

애플리케이션의 보안을 점검해야 하는 위치는 다음과 같다

- 클라이언트

- 서버

  - 서버의 입력값
  - 서버의 출력값

클라이언트 체크는 변조되거나 무효화 될 가능성이 있으므로 근본적인 보안 대책이 될 수 없다.

### XSS (Cross-Site Scripting)

취약성이 있는 웹사이트를 방문한 사용자의 브라우저에서 공격을 위한 HTML 태그나 JavaScript등을 동작시키는 공격이다.

공격자가 작성한 스크립트가 유저의 브라우저 상에서 작동하는 수동적 공격이다.

다음과 같은 경우에 발생할 수 있다

- 가짜 입력 폼 등에 의해 유저의 개인 정보를 도둑 맞는다
- 스크립트에 의해 유저의 쿠키 값이 도둑맞거나 피해자가 의도하지 않은 리퀘스트가 송신된다
- 가짜 문장이나 이미지 등이 표시된다

### SQL Injection

웹어플리케이션에서 이용하고 있는 데이터베이스에 SQL을 부적절하게 실행하는 공격이다.

SQL문에 전달될 데이터에 SQL문의 취약점을 공격할 수 있는 문자열을 넣어서 보낸다

- `--`를 사용하여 주어진 문자열 이후의 statement를 주석처리하여 무효화 할 수 있다
- `and 1 = 1`과 같은 항상 true인 조건을 보내 조건 처리를 무효화 할 수 있다

이외에도 여러가지 공격 방식이 존재한다

### OS Command Injection

웹 애플리케이션에게 부적절한 OS 명령어를 보내 실행하는 것이다.

쉘을 호출하는 함수가 있는 경우라면 공격에 노출 될 수 있다.

### HTTP Header Injection

리스폰스 헤더 필드에 개행 문자등을 삽입하여 임의의 리스폰스 헤더 필드나 바디를 추가하는 공격이다.

예를 들어, 링크를 클릭하면 Location 헤더 필드가 바뀌면서 리다이렉트 되는 경우,
쿼리 스트링에 개행문자와 공격을 위한 헤더를 추가할 수 있다

`%0D%0A(개행문자)Set-Cookie: SID=123456789`

또는 개행문자를 연속하여 두 개 입력하여(`%0D%0A%0D%0A`)
HTTP 헤더와 바디의 경계선으로 인식하도록 하고 가짜 바디를 보낼 수 있다.

### Mail Header Injection

웹 어플리케이션의 메일 송신 기능에 임의의 To, Subject 등의 메일 헤더를 추가하는 공격이다.

마찬가지로 메일헤더에서 개행문자인 (`%0D%0A`)를 입력하여 헤더추가, 본문변조등을 시도할 수 있다.

### Directory Traversal

웹 어플리케이션의 파일을 조작하는 처리에서 외부 파일이름을 지정하여 처리하는 경우에
처리가 취약하면 파일의 상대경로나, 절대경로 입력을 통해 임의의 파일/디렉토리에 액세스 할 수 있다.

예를 들어, 쿼리스트링에 직접 읽어올 파일명을 지정하는 경우 파일이름 대신 경로를 요청할 수 있다

`http://example.com/read?file=../../etc/passwd`

### Remote File Inclusion

스크립트의 일부를 다른 파일에서 읽어올 때
공격자가 지정한 외부 서버의 URL을 파일에서 읽게 하여 임의의 스크립트를 동작시키는 경우이다.

### Forced Browsing

웹 서버의 공개 디렉토리 중에서 공개의도가 없는 파일이 노출되는 경우이다

### 부적절한 에러 메시지 처리

웹 애플리케이션이나 데이터베이스의 에러메시지가 출력되어 공격자에게 유리한 내용이 표시되는 경우이다

### Open Redirect

파라미터로 리다이렉트할 URL을 지정하는 기능을 사용하는 경우 발생할 수 있다

### Session Hijack

공격자가 다른 유저의 세션ID를 탈취해서 약용하는 것으로 다른 유저로 위장한 공격을 하게된다

XSS와 같은 취약점 공격으로 쿠키에 포함된 세션ID를 탈취하여 공격자의 헤더에 세션 ID를 포함할 수 있다.

### Session Fixation

지정한 세션 ID를 유저에게 강제적으로 사용하게 하는 공격방식이다

미리 준비한 세션을 유저가 사용하도록 유도하고,
유저가 해당 세션에서 인증을 하면 인증된 세션으로 공격자가 접근하는 방식이다.

### Cross-Site Request Forgeries

인증된 유저가 개인정보나 설정정보 등을 공격자가 설치해둔 함정을 통해 작동시키는 공격이다.

- 인증된 유저 권한으로 설정 정보를 갱신
- 인증된 유저 권한으로 상품을 구입
- 인증된 유저 권한으로 게시물 작성

등이 발생할 수 있다.

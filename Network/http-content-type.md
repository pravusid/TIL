# HTTP Content-Type

- <https://tools.ietf.org/html/rfc7231#section-3.1.1.5>
- <https://ko.wikipedia.org/wiki/MIME>

`Content-Type` 헤더는 HTTP 요청 본문의 데이터 타입을 정의한다. 주로 사용되는 Content-Type은 다음과 같다.

## `application/x-www-form-urlencoded`

query string 타입의 key-value tuple로 인코딩되며, RFC 2396에 따라 예약 되지 않은 문자를 제외한 모든 문자를 16 진수 표현으로 변환해서 사용한다.

> 바이너리 데이터는 multipart/form-data 를 사용한다

```txt
POST / HTTP/1.1
Host: foo.com
Content-Type: application/x-www-form-urlencoded
Content-Length: 13

say=Hi&to=Mom
```

## `multipart/form-data`

```txt
POST /test.html HTTP/1.1
Host: example.org
Content-Type: multipart/form-data;boundary="boundary"

--boundary
Content-Disposition: form-data; name="field1"

value1
--boundary
Content-Disposition: form-data; name="field2"; filename="example.txt"

value2
--boundary--
```

## `application/json`

```txt
POST / HTTP/1.1
Host: foo.com
Content-Type: application/json
Content-Length: 35

{"request":"give me some response"}
```

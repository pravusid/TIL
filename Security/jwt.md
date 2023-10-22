# JWT (JSON Web Token)

> RFC 7519: <https://tools.ietf.org/html/rfc7519>

JWT는 당사자간 클레임을 안전하게 전송하기 위한 표준이다.
여러가지 목적으로 사용할 수 있으나 HTTP 인증에 많이 사용하고 있으며 header에 포함해서 전송한다.

- header, payload, signature 세 부분으로 구성되며 base64 인코딩 후 연결한 문자열임
- 내용은 암호화 되지 않으며, signature 부분만 암호화 된 값으로 진위여부를 판단한다
- signature 생성 알고리즘은 선택할 수 있으며, JWT header에 명시된다

## 알고리즘

- 대칭키 방식으로 HS256/512(HMAC with SHA-256/512)
- 비대칭키 방식으로 RS256/512(RSA Signature with SHA-256/512)

### 대칭키 방식

> RFC 4868: <https://tools.ietf.org/html/rfc4868#page-5>

- Block size (block size 크기 == secret 최대 크기)

  - SHA-256: 512 bits
  - SHA-384, SHA-512: 1024 bits

- Output length (해시 알고리즘 결과와 같음)

  - SHA-256: 256 bits
  - SHA-384: 384 bits
  - SHA-512: 512 bits

## Standard Fields

- <https://datatracker.ietf.org/doc/html/rfc7519#section-4>
- <https://en.wikipedia.org/wiki/JSON_Web_Token#Standard_fields>

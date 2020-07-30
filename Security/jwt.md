# JWT (JSON Web Token)

stateless + 인증에 주로 사용되며 주로 HTTP header에 포함한다

RFC 7519: <https://tools.ietf.org/html/rfc7519>

header, payload, signature 세 부분으로 구성되며 base64 인코딩 후 연결한 문자열임

내용은 암호화 되지 않으며, signature 부분만 암호화 된 값으로 진위여부를 판단한다

signature 생성 알고리즘은 선택할 수 있으며, JWT header에 명시된다

## 알고리즘

- 대칭키 방식으로 HS256/512(HMAC with SHA-256/512)
- 비대칭키 방식으로 RS256/512(RSA Signature with SHA-256/512)

### 대칭키 방식

RFC 4868: <https://tools.ietf.org/html/rfc4868#page-5>

- Block size (block size 크기 == secret 최대 크기)
  - SHA-256: 512 bits
  - SHA-384, SHA-512: 1024 bits

- Output length (해시 알고리즘 결과와 같음)
  - SHA-256: 256 bits
  - SHA-384: 384 bits
  - SHA-512: 512 bits

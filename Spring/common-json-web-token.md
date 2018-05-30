# JWT

OAuth와 JSON Web Token의 개념, 그리고 Spring Security 적용법

## JWT 개념

사용자가 누구인지 확인하는 과정을 인증(authentication)
토큰은 세션 아이디에 비해 서버측 부하를 낮출 수 있고, 분산/클라우드 기반 인프라스트럭처에 대응하기 좋음

또한 Claim 기반 토큰은 다른 인증방식에 비해서도 서버부담을 줄여준다. 서비스를 호출한 사용자 정보를 담고 있기 때문이다.

## JWT의 구조

JWT는 헤더(header), 페이로드(payload), 시그니처(signature)로 나누어짐

- 헤더는 토큰에 대한 설명을 담고 있음
- 페이로드는 JWT의 권한 정보를 갖고있음
- 시그니처는 토큰 무결성을 검증하기 위한 해시값을 담고 있음

페이로드를 디코드하면 권한정보를 담고 있는 JSON 객체를 확인할 수 있음

## 사용상 주의점

- JWT는 안전한 HttpOnly 쿠키에 저장해야 Cross-Site Scripting(XSS) 공격을 방지할 수 있음
- 쿠키를 사용해서 JWT를 전송한다면, CSRF 방어가 중요함
- 토큰을 사용해서 사용자를 인증할 때마다 항상 보안 키로 서명되어 있는지 검사해야 함
- 민감한 데이터는 JWT에 저장하면 안됨. 토큰은 일반적으로 조작을 방지하기 위한 목적으로 서명되므로 권한(claim) 데이터는 쉽게 디코드decode해서 볼 수 있음
- replay 공격에 대비하려면 nonce(jti claim), expiration time, creation time을 권한(claims)에 포함시켜야 함

## OAuth

OAuth 2.0은 인증authentication과 허가authorization를 제공하는 서비스와 상호 연동을 위한 방식이고 JWT가 다수 활용되고 있음

두 가지 토큰 타입: access token, refresh token

- 최초 인증시 두개의 토큰을 발급받는다
- 액세스 토큰 만료 기한을 짧게 두고 만료시킨다
- 엑세스 토큰이 만료되면 리프레시 토큰으로 새로운 토큰을 획득한다

## 자바에서 JWT 생성 및 파싱

자바에서는 JJWT를 사용 <https://github.com/jwtk/jjwt>

- Issuer, Subject, Expiration, ID와 같은 토큰의 내부 claim을 정의
- JWT를 암호화 서명을 해서 JWS를 생성
- JWT Compact Serialization 규칙에 따라 URL로 사용할 수 있도록 JWT를 압축

최종 JWT는 3부분으로 이루어져 있으며 지정된 키로 특정 서명 알고리즘으로 서명된어 Base64 인코딩된 문자열이 됨

### 공식문서 토큰 생성 예제

```java
Key key = MacProvider.generateKey();

String compactJws = Jwts.builder()
  .setSubject("Joe")
  .signWith(SignatureAlgorithm.HS512, key)
  .compact();
```

### 토큰 파싱 예제

```java
try {
  Jws<Claims> claims = Jwts.parser()
    .requireSubject("Joe")
    .require("hasMotorcycle", true)
    .setSigningKey(key)
    .parseClaimsJws(compactJws);
} catch (MissingClaimException e) {
  // we get here if the required claim is not present
} catch (IncorrectClaimException e) {
  // we get here if the required claim has the wrong value
}
```

파싱 도중 예외가 발생할 수 있다. JJWT의 에외들은 모두 `RuntimeExceptions`와 `JwtException`의 하위 클래스이다.

## Spring Security

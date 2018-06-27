# JWT

OAuth와 JSON Web Token의 개념, 그리고 Spring Security 적용법

## OAuth

OAuth 2은 인증 authentication과 허가 authorization를 제공하는 서비스와 상호 연동을 위한 방식이고 JWT가 다수 활용되고 있음

두 가지 토큰 타입: access token, refresh token

- 최초 인증시 두개의 토큰을 발급받는다
- 액세스 토큰 만료 기한을 짧게 두고 만료시킨다
- 엑세스 토큰이 만료되면 리프레시 토큰으로 새로운 토큰을 획득한다

### OAuth Token

access token을 제공하고 서버세어 토큰과 연관된 정보를 찾아 서비스를 제공한다.

1. 클라이언트가 토큰 발급 서버로 토큰을 요청한다. 요청시 사용자의 계정과 권한을 같이 전송한다.
2. 토큰 발급 서버는 사용자 계정의 진위를 확인한 후, token에 대한 정보를 token 저장소에 저장한다.
3. 생성된 토큰을 클라이언트가 다시 받아간다.
4. 클라이언트는 API를 호출할 때 보유한 token을 이용해서 리소스 서버의 API를 호출한다.
5. 리소스 서버는 호출이 발생하면 token 저장소에서 권한정보를 조회하고, 권한에 맞는 응답을 보낸다.

## JWT 개념

사용자가 누구인지 확인하는 과정을 인증(authentication)
토큰은 세션 아이디에 비해 서버측 부하를 낮출 수 있고, 분산/클라우드 기반 인프라스트럭처에 대응하기 좋음

또한 Claim 기반 토큰은 다른 인증방식에 비해서도 서버부담을 줄여준다. 서비스를 호출한 사용자 정보를 담고 있기 때문이다.

JSON 문자열을 BASE64 인코딩을 통해 문자열로 변환하여 전송한다.

### JWT token

1. OAuth와 마찬가지로 사용자를 인증한 후 토큰을 생성한다.
2. token 관련 정보를 저장소에 저장하지 않고, 토큰 자체에 넣어서 저장한다.
3. 이후 클라이언트와 서버의 통신과정은 동일하다.

### JWT의 구조

JWT는 헤더(header), 페이로드(payload), 시그니처(signature)로 나누어짐

- 헤더는 토큰에 대한 설명을 담고 있음
- 페이로드는 JWT의 권한 정보를 갖고있음
- 시그니처는 토큰 무결성을 검증하기 위한 해시값을 담고 있음

페이로드를 디코드하면 권한정보를 담고 있는 JSON 객체를 확인할 수 있음

### 사용상 주의점

- JWT는 안전한 HttpOnly 쿠키에 저장해야 Cross-Site Scripting(XSS) 공격을 방지할 수 있음
- 쿠키를 사용해서 JWT를 전송한다면, CSRF 방어가 중요함
- 토큰을 사용해서 사용자를 인증할 때마다 항상 보안 키로 서명되어 있는지 검사해야 함
- 민감한 데이터는 JWT에 저장하면 안됨. 토큰은 일반적으로 조작을 방지하기 위한 목적으로 서명되므로 권한(claim) 데이터는 쉽게 디코드decode 해서 볼 수 있음
- replay 공격에 대비하려면 nonce(jti claim), expiration time, creation time을 권한(claims)에 포함시켜야 함

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

## Spring Security: OAuth2 & JWT

<https://projects.spring.io/spring-security-oauth/docs/oauth2.html>

기본적으로 Authorization 서버와 Resource Server는 분리되어 있어야 함.

`AuthorizationEndpoint` authorization(인가)에 대한 서비스 요청 처리: Default URL: `/oauth/authorize`

`TokenEndpoint` access token에 대한 서비스 요청 처리: Default URL: `/oauth/token`

### 설정

oauth2 와 spring-jwt 라이브러리를 의존성에 추가한다

```groovy
compile("org.springframework.security:spring-security-jwt")
compile("org.springframework.security.oauth:spring-security-oauth2")
```

#### OAuth 정보를 위한 DB 설정

OAuth 정보저장을 위한 테이블을 생성한다.
<https://github.com/spring-projects/spring-security-oauth/blob/master/spring-security-oauth2/src/test/resources/schema.sql>

```sql
CREATE TABLE `oauth_client_details` (
  `client_id` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `resource_ids` varchar(256) COLLATE utf8_unicode_ci DEFAULT NULL,
  `client_secret` varchar(256) COLLATE utf8_unicode_ci DEFAULT NULL,
  `scope` varchar(256) COLLATE utf8_unicode_ci DEFAULT NULL,
  `authorized_grant_types` varchar(256) COLLATE utf8_unicode_ci DEFAULT NULL,
  `web_server_redirect_uri` varchar(256) COLLATE utf8_unicode_ci DEFAULT NULL,
  `authorities` varchar(256) COLLATE utf8_unicode_ci DEFAULT NULL,
  `access_token_validity` int(11) DEFAULT NULL,
  `refresh_token_validity` int(11) DEFAULT NULL,
  `additional_information` varchar(4096) COLLATE utf8_unicode_ci DEFAULT NULL,
  `autoapprove` varchar(256) COLLATE utf8_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`client_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
```

JPA 사용을 위해 `OauthClientDetails.java` 파일을 다음과 같이 작성한다.

```java
@Entity
public class OAuthClientDetails {
    @Id
    private long id;
    private String clientId;
    private String resourceIds;
    private String clientSecret;
    private String scope;
    private String authorizedGrantTypes;
    private String webServerRedirectUri;
    private String authorities;
    private int accessTokenValidity;
    private int refreshTokenValidity;
    private String additionalInformation;
    private String autoapprove;

    // getters
}
```

사용할 클라이언트 정보를 입력해둔다 (vueclient)

```sql
INSERT INTO
  oauth_client_details (
    id,
    client_id,
    client_secret,
    resource_ids,
    scope,
    authorized_grant_types,
    web_server_redirect_uri,
    authorities,
    access_token_validity,
    refresh_token_validity,
    additional_information,
    autoapprove
  )
  VALUES (
    1,
    'vueclient',
    'vueclientpwd',
    'spring-boot-application',
    'read,write',
    'authorization_code,password,implicit,refresh_token',
    null,
    'USER',
    36000,
    2592000,
    null,
    null
  );
```

#### TokenStore

토큰스토어 종류

1. `org.springframework.security.oauth2.provider.token.store.InMemoryTokenStore`: JAVA 내부에서 Map, Queue 구조의 메모리를 사용
2. `org.springframework.security.oauth2.provider.token.store.JdbcTokenStore`: JDBC 를 사용해서 DB 에 저장
3. `org.springframework.security.oauth2.provider.token.store.JwtTokenStore`: Json Web Token 을 이용
4. `org.springframework.security.oauth2.provider.token.store.redis.RedisTokenStore`: Redis 에 Token 정보를 저장

#### Token을 처리할 서버 설정

`AuthorizationServerEndpointsConfigurer`에서 허가유형에 대한 설정을 할 수 있다.

- `authenticationManager`: password 허가로 전환하기 위해서는 AuthenticationManager를 주입해야 한다
- `userDetailsService`: UserDetailsService를 주입하거나 어떤 방법으로든 글로벌하게 구성할 수 있다면 (예를 들어, GlobalAuthenticationManagerConfigurer), refresh token 허가는 계정이 여전히 활성화되어 있는지 보장하기 위해 user details에서 검증을 포함하게 된다.
- `authorizationCodeServices`: auth code 허가를 위해 인가 코드 서비스(AuthorizationCodeServices의 인스턴스)를 정의한다.
- `implicitGrantService`: implicit 허가 동안에 상태를 관리한다.
- `tokenGranter`: TokenGranter (허가 제어 전체를 포함하며 위의 다른 속성은 무시한다)

`AuthorizationSeverConfig.java`

```java
@Configuration
@EnableAuthorizationServer
public class AuthorizationSeverConfig extends AuthorizationServerConfigurerAdapter {

    private PasswordEncoder passwordEncoder;
    private DataSource dataSource;
    private AuthenticationManager authenticationManager;

    public AuthorizationSeverConfig(PasswordEncoder passwordEncoder,
            DataSource dataSource,
            AuthenticationManager authenticationManager) {
        this.passwordEncoder = passwordEncoder;
        this.dataSource = dataSource;
        this.authenticationManager = authenticationManager;
    }

    @Override
    public void configure(AuthorizationServerSecurityConfigurer securityConfig) throws Exception {
        // password encoder bean이 있다면 설정하지 않아도 적용되어 있음 (client 정보 조회시)
        securityConfig.passwordEncoder(passwordEncoder);
    }

    // client 정보를 DB로 부터 조회
    @Override
    public void configure(ClientDetailsServiceConfigurer clients) throws Exception {
        clients.withClientDetails(clientDetailsService());
    }

    @Bean
    @Primary
    public ClientDetailsService clientDetailsService() {
        return new JdbcClientDetailsService(dataSource);
    }

    @Override
    public void configure(AuthorizationServerEndpointsConfigurer endpoints) {
        endpoints.authenticationManager(authenticationManager)
                .tokenStore(tokenStore())
                .accessTokenConverter(accessTokenConverter());
    }

    // 아래의 세 메소드는 Jwt 토큰용 (인증 서버 / 리소스 서버 모두 있음)

    @Bean
    public TokenStore tokenStore() {
        return new JwtTokenStore(accessTokenConverter());
    }

    @Bean
    public JwtAccessTokenConverter accessTokenConverter() {
        JwtAccessTokenConverter converter = new JwtAccessTokenConverter();
        KeyStoreKeyFactory keyFactory =new KeyStoreKeyFactory(new ClassPathResource("private.jks"), "storepass".toCharArray());
        converter.setKeyPair(keyFactory.getKeyPair("jwtserver", "keypass".toCharArray()));
//        converter.setSigningKey("secret");
        return converter;
    }

    @Bean
    @Primary
    public DefaultTokenServices tokenService() {
        DefaultTokenServices defaultTokenServices = new DefaultTokenServices();
        defaultTokenServices.setTokenStore(tokenStore());
        defaultTokenServices.setSupportRefreshToken(true);
        return defaultTokenServices;
    }

}
```

#### Token을 받아 처리할 리소스 서버 설정

`ResourceServerConfig.java`

```java
@Configuration
@EnableResourceServer
public class ResourceServerConfig extends ResourceServerConfigurerAdapter {

    @Value("${resource.id:spring-boot-application}")
    private String resourceId;

    @Value("${security.oauth2.resource.jwt.key-value}")
    private String publicKey;

    @Override
    public void configure(HttpSecurity http) throws Exception {
        http
            .requestMatchers()
                .antMatchers("/api/**")
                .and()
            .authorizeRequests()
                .antMatchers("/api/**").hasAuthority(Authority.USER.getAuthority())
                .and()
            .exceptionHandling().accessDeniedHandler(new OAuth2AccessDeniedHandler());
    }

    @Override
    public void configure(ResourceServerSecurityConfigurer config) {
        config.resourceId(resourceId)
                .tokenStore(tokenStore());
    }

    // 아래의 세 메소드는 Jwt 토큰용 (인증 서버 / 리소스 서버 모두 있음)

    @Bean
    public TokenStore tokenStore() {
        return new JwtTokenStore(accessTokenConverter());
    }

    @Bean
    public JwtAccessTokenConverter accessTokenConverter() {
        JwtAccessTokenConverter converter = new JwtAccessTokenConverter();
        converter.setVerifierKey(publicKey);
//        converter.setSigningKey("secret");
        return converter;
    }

    @Bean
    @Primary
    public DefaultTokenServices tokenService() {
        DefaultTokenServices defaultTokenServices = new DefaultTokenServices();
        defaultTokenServices.setTokenStore(tokenStore());
        defaultTokenServices.setSupportRefreshToken(true);
        return defaultTokenServices;
    }

}
```

#### 토큰 발급 주소를 허용

`AuthorizationServerEndpointsConfigurer`에서는 `pathMapping()`를 제공

- 엔드포인트의 기본 URL 경로 (기본 구현체는 프레임워크 제공)
- 필수 커스텀 경로 (“/”로 시작)

프레임워크가 제공하는 URL 경로는

- /oauth/authorize (인가 엔드포인트)
- /oauth/token (토큰 엔드포인트)
- /oauth/confirm_access (사용자가 허가의 승인을 확인하는 POST 요청)
- /oauth/error (인가 서버에서 에러를 보여줄 때 사용)
- /oauth/check_token (액세스 토큰을 복호화 하기 위해 리소스 서버에서 사용)
- /oauth/token_key (JWT 토큰을 사용하는 경우 토큰 검증을 위한 공개키를 노출)가 있다.

인가 엔드포인트 /oauth/authorize (또는 대채된 매핑 경로)는 인증된 사용자만 접근할 수 있도록 Spring Security를 사용해 보호해야 한다.
예를 들어, 표준 Spring Security의 WebSecurityConfigurer를 사용하면 다음과 같다.

```java
public class SecurityConfig extends WebSecurityConfigurerAdapter {
    ...
    @Override
    protected void configure(HttpSecurity http) throws Exception {
        http
            .authorizeRequests()
                ...
                .antMatchers("/oauth/**").hasAuthority("USER")
                ...
    }
    ...
}
```

#### 토큰 검증 해시값 생성을 위한 비대칭키 생성

JRE에 포함된 keytool을 사용하여 비대칭키를 생성한다 (**RSA 알고리즘 사용해야 함**)

`keytool -genkeypair -alias jwtserver -keyalg RSA -keypass keypass -keystore private.jks -storepass storepass`

비밀키에서 공개키를 추출한다.

`keytool -list -rfc --keystore private.jks | openssl x509 -inform pem -pubkey -noout`

생성한 public 키 정보를 Resource 서버 설정에 넣어준다.

```yml
security:
  oauth2:
    resource:
      filter-order: 3 # 기존 Spring security 인증필터 체인 사용 하려면
      jwt:
        key-value:
          -----BEGIN PUBLIC KEY-----

          -----END PUBLIC KEY-----
```

(참고): 필터 순서: OAuth 2 Resource Filter (Debug mode에서 `filters: ArrayList` 내부 순서대로 필터가 작동함)

OAuth2 resource filter의 기본 실행 순서가 3에서 SecurityProperties.ACCESS_OVERRIDE_ORDER - 1 로 변경됨
actuator endpoints 다음 순서이고 the basic authentication filter chain 이전에 실행됨
기존의 순서인 3으로 다시 변경하려면 `security.oauth2.resource.filter-order = 3`

### CustomTokenEnhancer

#### Custom Claims in the Token

추가적인 claim 정의를 위해서 사용자 `TokenEnhancer`를 정의함

```java
public class CustomTokenEnhancer implements TokenEnhancer {
    @Override
    public OAuth2AccessToken enhance(OAuth2AccessToken accessToken, OAuth2Authentication authentication) {
        Map<String, Object> additionalInfo = new HashMap<>();
        additionalInfo.put("organization", authentication.getName() + randomAlphabetic(4));
        ((DefaultOAuth2AccessToken) accessToken).setAdditionalInformation(additionalInfo);
        return accessToken;
    }
}
```

Authorization Server 설정에서 TokenEnhancer를 연결한다

```java
    @Override
    public void configure(
    AuthorizationServerEndpointsConfigurer endpoints) throws Exception {
        TokenEnhancerChain tokenEnhancerChain = new TokenEnhancerChain();
        tokenEnhancerChain.setTokenEnhancers(
        Arrays.asList(tokenEnhancer(), accessTokenConverter()));

        endpoints.tokenStore(tokenStore())
                .tokenEnhancer(tokenEnhancerChain)
                .authenticationManager(authenticationManager);
    }

    @Bean
    public TokenEnhancer tokenEnhancer() {
        return new CustomTokenEnhancer();
    }
```

위의 설정으로 생성되는 token의 payload 형태는 다음과 같다

```json
{
    "user_name": "john",
    "scope": [
        "foo",
        "read",
        "write"
    ],
    "organization": "johnIiCh",
    "exp": 1458126622,
    "authorities": [
        "ROLE_USER"
    ],
    "jti": "e0ad1ef3-a8a5-4eef-998d-00b26bc2c53f",
    "client_id": "fooClientIdPassword"
}
```

#### Access Extra Claims on Resource Server

리소스 서버에서 CustomToken을 처리해보자

```java
    public Map<String, Object> getExtraInfo(OAuth2Authentication auth) {
        OAuth2AuthenticationDetails details = (OAuth2AuthenticationDetails) auth.getDetails();
        OAuth2AccessToken accessToken = tokenStore.readAccessToken(details.getTokenValue());
        return accessToken.getAdditionalInformation();
    }
```

`CustomAccessTokenConverter`

```java
@Component
public class CustomAccessTokenConverter extends DefaultAccessTokenConverter {
    @Override
    public OAuth2Authentication extractAuthentication(Map<String, ?> claims) {
        OAuth2Authentication authentication = super.extractAuthentication(claims);
        authentication.setDetails(claims);
        return authentication;
    }
}
```

CustomAccessTokenConverter를 위한 JwtTokenStore를 정의하자

```java
@Configuration
@EnableResourceServer
public class OAuth2ResourceServerConfigJwt extends ResourceServerConfigurerAdapter {

    @Autowired
    private CustomAccessTokenConverter customAccessTokenConverter;

    @Bean
    public TokenStore tokenStore() {
        return new JwtTokenStore(accessTokenConverter());
    }

    @Bean
    public JwtAccessTokenConverter accessTokenConverter() {
        JwtAccessTokenConverter converter = new JwtAccessTokenConverter();
        converter.setAccessTokenConverter(customAccessTokenConverter);
    }

}
```

Resource Server 에서 Authentication object 처리를 정의하자

```java
    public Map<String, Object> getExtraInfo(Authentication auth) {
        OAuth2AuthenticationDetails oauthDetails = (OAuth2AuthenticationDetails) auth.getDetails();
        return (Map<String, Object>) oauthDetails.getDecodedDetails();
    }
```

#### Authentication Details Test

추가 claim을 포함한 token 처리를 테스트 해보자

```java
@RunWith(SpringRunner.class)
@SpringBootTest(classes = ResourceServerApplication.class, webEnvironment = WebEnvironment.RANDOM_PORT)
public class AuthenticationClaimsIntegrationTest {

    @Autowired
    private JwtTokenStore tokenStore;

    @Test
    public void whenTokenDoesNotContainIssuer_thenSuccess() {
        String tokenValue = obtainAccessToken("fooClientIdPassword", "john", "123");
        OAuth2Authentication auth = tokenStore.readAuthentication(tokenValue);
        Map<String, Object> details = (Map<String, Object>) auth.getDetails();
  
        assertTrue(details.containsKey("organization"));
    }

    private String obtainAccessToken(
      String clientId, String username, String password) {
  
        Map<String, String> params = new HashMap<>();
        params.put("grant_type", "password");
        params.put("client_id", clientId);
        params.put("username", username);
        params.put("password", password);
        Response response = RestAssured.given()
          .auth().preemptive().basic(clientId, "secret")
          .and().with().params(params).when()
          .post("http://localhost:8081/spring-security-oauth-server/oauth/token");
        return response.jsonPath().getString("access_token");
    }
}
```

### 테스트

`curl -u vueclient:vueclientpwd http://localhost:8080/oauth/token -d "grant_type=password&username=user&password=1111"`

클라이언트에서 발급받은 토큰으로 요청

`curl -H "authorization: bearer {access_token}" http://localhost:8080/api/user`

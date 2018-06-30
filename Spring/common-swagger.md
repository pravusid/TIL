# Swagger

Swagger는 RESTful API를 문서화하는 도구이다.

Swagger Doc는 YAML 포맷으로 작성된다.

## 시작

의존성 추가

```groovy
compile("io.springfox:springfox-swagger2:2.8.0")
compile("io.springfox:springfox-swagger-ui:2.8.0")
```

## 설정

`SwaggerConfig` 생성

```java
@Configuration
@EnableSwagger2
public class SwaggerConfig {

    @Bean
    public Docket api() {
        return new Docket(DocumentationType.SWAGGER_2)
                .select()
                .apis(RequestHandlerSelectors.basePackage("kr.pravusid.web"))
                .paths(PathSelectors.any())
                .build();

    }
}
```

Security 경로가 root 부터 라면 다음 내용을 추가한다 (인증경로 회피)

```java
public class SecurityConfig extends WebSecurityConfigurerAdapter {
    ...
    @Override
    public void configure(WebSecurity web) throws Exception {
        web.ignoring().antMatchers("/v2/api-docs", "/swagger-resources", "/swagger-ui.html", "/swagger/**");
    }
    ...
}
```

## 사용

컨트롤러 `RequestMapping`에서 Swagger용 `Annotation`을 사용하여 사용자 임의 설정을 할 수 있다.

### DOCS

```java
@ApiOperation(value = “메시지 테스트”, httpMethod = “GET”, notes = “메시지 테스트 실행”)
@ApiResponses(value = {
    @ApiResponse(code = 400, message = “Invalid”),
    @ApiResponse(code = 200, message = “OK” )
})
@GetMapping(”/message”)
public String getMessage(
        @ApiParam(value = “키값”, required = true, defaultValue = “기본값”) String value) {
    return "msg";
}
```

### 헤더에 토큰 값 설정 (Request header)

```java
@ApiImplicitParams({
    @ApiImplicitParam(name = “Authorization”, value = “authorization header”,
            required = true, dataType = “string”, paramType = “header”)
})
```

# Thymeleaf

## 소개

## 설정

application.properties

  ```text
  spring.thymeleaf.check-template-location=true
  spring.thymeleaf.prefix=/WEB-INF/templates/
  spring.thymeleaf.suffix=.html
  spring.thymeleaf.mode=HTML
  spring.thymeleaf.encoding=UTF-8
  spring.thymeleaf.content-type=text/html
  spring.thymeleaf.cache=false
  ```

Thymeleaf는 태그 정책이 엄격해서 오타나 표준에 맞지 않는 구문이 있으면 칼같이 오류를 내뿜는다.
특히 닫는 태그 등의 HTML 표준 관련 충돌이 잦은데 이를 완화하기 위해서 의존성 패키지를 추가한다.

`spring.thymeleaf.mode=LEGACYHTML5` 적용 시
gradle dependency 추가 `compile("net.sourceforge.nekohtml:nekohtml:1.9.22")`

Spring boot thymeleaf-starter는 thymeleaf2 버전을 적용하고 있기 때문에 버전 변경이 필요하다

```groovy
ext {
    thymeleafVersion = '3.0.1.RELEASE'
}
dependencies {
  compile("org.thymeleaf:thymeleaf:${thymeleafVersion}")
  compile("org.thymeleaf:thymeleaf-spring4:${thymeleafVersion}")
  compile("org.thymeleaf.extras:thymeleaf-extras-springsecurity4:${thymeleafVersion}")
}
```

### Java Config 예시

```java
@Configuration
public class ThymeleafConfig extends WebMvcConfigurerAdapter implements ApplicationContextAware {

    @Autowired
    private MessageSource messageSource;

    @Value("${thymeleaf.templates.cache}")
    String thymeleafCache;

    private ApplicationContext applicationContext;

    public void setApplicationContext(ApplicationContext applicationContext) {
        this.applicationContext = applicationContext;
    }

    @Bean
    public ITemplateResolver templateResolver() {
        SpringResourceTemplateResolver templateResolver = new SpringResourceTemplateResolver();
        templateResolver.setApplicationContext(applicationContext);
        templateResolver.setPrefix("classpath:/templates/");
        templateResolver.setSuffix(".html");
        templateResolver.setTemplateMode("HTML");
        templateResolver.setCharacterEncoding("UTF-8");

        if (thymeleafCache.equals("true")){
            templateResolver.setCacheable(true);
        } else {
            templateResolver.setCacheable(false);
        }

        return templateResolver;
    }

    @Bean
    public SpringTemplateEngine templateEngine() {
        SpringTemplateEngine templateEngine = new SpringTemplateEngine();

        templateEngine.setEnableSpringELCompiler(true);
        templateEngine.setTemplateResolver(templateResolver());
        templateEngine.setMessageSource(messageSource);
        templateEngine.addDialect(new SpringDataDialect());

        return templateEngine;
    }

    @Bean
    public ViewResolver viewResolver() {
        ThymeleafViewResolver viewResolver = new ThymeleafViewResolver();
        viewResolver.setTemplateEngine(templateEngine());
        return viewResolver;
    }
}
```

## 기본 문법

### thymeleaf 사용 선언

`<html xmlns:th="http://www.thymeleaf.org">`

### 범위 선언

범위 내의 변수선언, 조건설정등을 할 수 있다.

`<th:block></th:block>`

### model의 attribute 출력

`<span th:text="${variable}"></span>`

`<span th:utext="${variable}"></span>`

#### concat attribute and String

`th:text="'static part' + ${attr.field}"`

`th:text="${'static part' + attr.field}"`

#### inline expression (3버전 이상부터 지원)

`[[${session.user.name}]]`

### 반복문

```html
<tr data-th-each="data : ${list}">
  <td th:text="${data.userId}"></td>
  <td th:text="${data.name}"></td>
  <td th:text="${data.email}"></td>
</tr>
```

만약 숫자 범위를 지정하여 반복한다면

  ```html
  <!-- foo.start, foo.end 자리에 상수도 사용 가능함 -->
  <th:block th:each="page : ${#numbers.sequence(__${foo.start}__, __${foo.end}__)}">
    <span th:text="${page}"></span>
  </th:block>
  ```

### 조건문

삼항 연산자도 사용 가능하다

  ```html
  <tr th:class="${row.even}? (${row.first}? 'first' : 'even') : 'odd'">
    ...
  </tr>
  ```

if문

  ```html
  <li th:if="${session.user==null}"><a href="/login">로그인</a></li>
  <li th:if="${session.user!=null}"><a href="/logout">로그아웃</a></li>
  ```

switch문

  ```html
  <div th:switch="${user.role}">
    <p th:case="'admin'">User is an administrator</p>
    <p th:case="#{roles.manager}">User is a manager</p>
    <p th:case="*">User is some other thing</p>
  </div>
  ```

## 레이아웃 (Layout)

- th:insert : th:insert를 선언한 태그를 유지하고 내부에 fragment 전체를 가져옴

  ```html
  <div th:insert="footer :: copy"></div>
  <div>
    <footer>
      &copy; 2011 The Good Thymes Virtual Grocery
    </footer>
  </div>
  ```

- th:replace : th:replace를 선언한 태그 자체가 fragment로 바뀜

  ```html
  <div th:replace="footer :: copy"></div>

  <footer>
    &copy; 2011 The Good Thymes Virtual Grocery
  </footer>
  ```

- th:include : th:inclue를 선언한 태그는 유지되고 fragment에서 최상위태그 내부의 내용만 가져옴

  ```html
  <div th:include="footer :: copy"></div>

  <div>
    &copy; 2011 The Good Thymes Virtual Grocery
  </div>
  ```

## Security

### namespace 추가

html namespace 추가 `<html xmlns:sec="http://www.thymeleaf.org/extras/spring-security">`

### 기본문법

The sec:authorize attribute renders its content when the attribute expression is evaluated to true:

  ```html
  <div sec:authorize="isAuthenticated()">
    This content is only shown to authenticated users.
  </div>
  <div sec:authorize="isAnonymous()">
    This content is only shown to anonymous.
  </div>
  <div sec:authorize="hasRole('ROLE_ADMIN')">
    This content is only shown to administrators.
  </div>
  <div sec:authorize="hasRole('ROLE_USER')">
    This content is only shown to users.
  </div>
  ```

The sec:authentication attribute is used to print logged user name and roles:

  ```html
  Logged user: <span sec:authentication="name">Bob</span>
  Roles: <span sec:authentication="principal.authorities">[ROLE_USER, ROLE_ADMIN]</span>
  ```

**다음처럼 쓸 수도 있다 : `${#authentication.principal.authorities}`**

### 표현식 목록

Expression | Description
--- | ---
hasRole([role]) | Returns true if the current principal has the specified role.
hasAnyRole([role1,role2]) | Returns true if the current principal has any of the supplied roles (given as a comma-separated list of strings)
hasAuthority([role]) | Returns true if the current principal has the specified authority.
hasAnyAuthority([role1,role2]) | Returns true if the current principal has any of the supplied Authority (given as a comma-separated list of strings)
principal | Allows direct access to the principal object representing the current user
authentication | Allows direct access to the current Authentication object obtained from the SecurityContext
permitAll | Always evaluates to true
denyAll | Always evaluates to false
isAnonymous() | Returns true if the current principal is an anonymous user
isRememberMe() | Returns true if the current principal is a remember-me user
isAuthenticated() | Returns true if the user is not anonymous
isFullyAuthenticated() | Returns true if the user is not an anonymous or a remember-me user

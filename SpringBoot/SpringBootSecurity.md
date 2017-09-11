# Spring Boot Security

## 시작

Gradle 의존성에 `compile('org.springframework.boot:spring-boot-starter-security')` 추가

## 설정

<https://docs.spring.io/spring-security/site/docs/current/guides/html5/index.html>

### SecurityConfig.java 파일 생성

```java
@EnableWebSecurity
public class SecurityConfig extends WebSecurityConfigurerAdapter {

  @Override
  protected void configure(HttpSecurity http) throws Exception {
    http
      .authorizeRequests()
        .antMatchers("/css/**", "/").permitAll()
        .antMatchers("/user/**").hasRole("USER")
        .and()
      .formLogin()
        .loginPage("/login").failureUrl("/login-error")
        .and()
      /* 로그아웃 커스텀 */
      .logout()
        .logoutUrl("/my/logout")
        .logoutSuccessUrl("/my/index")
        .logoutSuccessHandler(logoutSuccessHandler)
        .invalidateHttpSession(true)
        .addLogoutHandler(logoutHandler)
        .deleteCookies(cookieNamesToClear);
  }

}
```

### 로그인/로그아웃 폼 예제

```html
<c:url value="/login" var="loginUrl"/>
<form action="${loginUrl}" method="post">
  <c:if test="${param.error != null}">
    <p>
      Invalid username and password.
    </p>
  </c:if>
  <c:if test="${param.logout != null}">
    <p>
      You have been logged out.
    </p>
  </c:if>
  <p>
    <label for="username">Username</label>
    <input type="text" id="username" name="username"/>
  </p>
  <p>
    <label for="password">Password</label>
    <input type="password" id="password" name="password"/>
  </p>
  <input type="hidden"
    name="${_csrf.parameterName}"
    value="${_csrf.token}"/>
  <button type="submit" class="btn">Log in</button>
</form>
```
# Spring Boot Security

## 시작

Gradle 의존성에 `compile('org.springframework.boot:spring-boot-starter-security')` 추가

## 설정

<https://docs.spring.io/spring-security/site/docs/current/guides/html5/index.html>

### SecurityConfig.java 파일 생성

```java
// web기반의 Security 설정 annotation
@EnableWebSecurity
// method에 annotation으로 security 기능 적용 가능해짐
@EnableGlobalMethodSecurity(prePostEnabled = true)
public class SecurityConfig extends WebSecurityConfigurerAdapter {

  private final static String REMEMBER_ME_KEY = "KEY";
  private final static String COOKIE_NAME = "IDPRAVUS_AUTH";

  private UserDetailsService userDetailsService;

  @Autowired
  public SecurityConfig(UserDetailsService customUserDetailsService) {
    this.userDetailsService = customUserDetailsService;
  }

  // 인증설정 (in memory, JDBC ...) 하는데 쓰는 builder, UserDetailsService를 Bean으로 등록하면 별 쓸일 없는듯
  @Autowired
  public void configureGlobal(AuthenticationManagerBuilder auth) throws Exception {
    auth.userDetailsService(userDetailsService).passwordEncoder(passwordEncoder());
  }

  // 인증정보를 갖고가지 않을 주소를 설정한다
  @Override
  public void configure(WebSecurity web) throws Exception {
    web.ignoring().antMatchers("/js/**", "/css/**");
  }

  // 로그인을 위한 autenticationProvider를 설정한다
  @Override
  protected void configure(AuthenticationManagerBuilder auth) throws Exception {
      auth.authenticationProvider(authenticationProvider());
  }

  @Override
  protected void configure(HttpSecurity http) throws Exception {
    http
      // 인증정보를 설정한다
      .authorizeRequests()
        .antMatchers("/").permitAll()
        .antMatchers("/mypage/**").hasAnyAuthority(Authority.USER.getAuthority(), Authority.ADMIN.getAuthority())
        .antMatchers("/admin/**").hasAuthority(Authority.ADMIN.getAuthority())
        .antMatchers("/h2-console/**").permitAll()
        .and()
      // 추가 h2-console 관련설정
      .csrf()
        .ignoringAntMatchers("/h2-console/**")
        .and()
      .headers()
        .addHeaderWriter(new XFrameOptionsHeaderWriter(new WhiteListedAllowFromStrategy(Arrays.asList("localhost"))))
        .and()
      // 로그인 관련 설정, username과 password Column과 form name이 다르다면 명시해주자
      .formLogin()
        .loginPage("/login").failureUrl("/login?error").permitAll()
        .usernameParameter("userId")
        .passwordParameter("password")
        .defaultSuccessUrl("/")
        .and()
      // 로그아웃 관련 설정, Remember me 쿠키를 삭제하고 세션을 비활성화 한다
      .logout()
        .logoutUrl("/logout")
        .logoutSuccessUrl("/")
        // 쿠키로 remember Me 사용할 때 로그아웃시 쿠키 삭제
        // Persistent 기반에서는 로그아웃시 기본적으로 removeUserTokens(String username) Method 호출함
        .deleteCookies("JSESSIONID")
        .deleteCookies(COOKIE_NAME)
        .invalidateHttpSession(true)
        .and()
      // Error 페이지(인증 권한이 없을 때) 주소를 명시한다
      .exceptionHandling()
        .accessDeniedPage("/denied")
        .and()
      // Remember Me 관련 설정 (쿠키 기반), 서비스를 불러와 매칭하기 때문에 여기서 설정해봐야 안먹힘, Bean등록할 때 설정할 것
      .rememberMe()
        .key(REMEMBER_ME_KEY)
        .rememberMeServices(tokenBasedRememberMeServices());
      // Remember Me 관련 설정 (퍼시스턴스 기반)
      .rememberMe()
        .key(REMEMBER_ME_KEY)
        .rememberMeServices(persistentTokenBasedRememberMeServices());
  }

  // 패스워드 인코더 설정 Bean
  @Bean
  public PasswordEncoder passwordEncoder() {
    return new BCryptPasswordEncoder();
  }

  // AutenticationProvider를 bean으로 등록한다. 기본제공되는 AuthencationProvider의 구현체이다.
  // DaoAuthenticationProvider는 내부적으로 UserDetailsService를 호출해 db에서 사용자를 조회한다.
  @Bean
  public DaoAuthenticationProvider authenticationProvider() {
      DaoAuthenticationProvider authProvider = new DaoAuthenticationProvider();
      authProvider.setUserDetailsService(userDetailsService);
      authProvider.setPasswordEncoder(passwordEncoder());
      return authProvider;
  }

  // 쿠키 기반 Remember Me 설정, Https가 아니라면 SecureCookie가 작동하지 않을 수 있다.
  @Bean
  public TokenBasedRememberMeServices tokenBasedRememberMeServices() {
    TokenBasedRememberMeServices tokenBasedRememberMeServices =
        new TokenBasedRememberMeServices(REMEMBER_ME_KEY, userDetailsService);
    tokenBasedRememberMeServices.setParameter("REMEMBER_ME");
    tokenBasedRememberMeServices.setCookieName(COOKIE_NAME);
    tokenBasedRememberMeServices.setTokenValiditySeconds(60 * 60 * 24 * 30);
    tokenBasedRememberMeServices.setUseSecureCookie(false);
    return tokenBasedRememberMeServices;
  }

  // 퍼시스턴스 기반 Remember Me 설정
  @Bean
  public PersistentTokenBasedRememberMeServices persistentTokenBasedRememberMeServices(){
    PersistentTokenBasedRememberMeServices persistentTokenBasedRememberMeServices =
        new PersistentTokenBasedRememberMeServices(REMEMBER_ME_KEY, userDetailsService, persistentTokenRepository());
    persistentTokenBasedRememberMeServices.setParameter("REMEMBER_ME");
    persistentTokenBasedRememberMeServices.setTokenValiditySeconds(60 * 60 * 24 * 30);
    return persistentTokenBasedRememberMeServices;
  }

  @Bean
  public PersistentTokenRepository persistentTokenRepository(){
    return new PersistentTokenRepositoryImpl();
  }
}
```

### Persistent Remeber Me 구현

Token 정보를 저장하는 Entity 생성

  ```java
  @Entity
  public class PersistentLogins {
    @NotNull
    @OneToOne
    @JoinColumn(foreignKey=@ForeignKey(name = "fk_persistent_logins"))
    private User user; // User entity에 @OneToOne(mappedBy = "user")

    @Id
    private String series;

    @NotNull
    private String token;

    @NotNull
    private Date lastUsed;
  ```

Token Entity를 사용하기 위한 JpaRepository

  ```java
  public interface PersistentLoginsRepository extends JpaRepository<PersistentLogins, String> {
    public void deleteByUserId(Long userId);
  }
  ```

퍼시스턴스 기반 Remember Me에서 PersistentTokenRepository 구현

  ```java
  public class PersistentTokenRepositoryImpl implements PersistentTokenRepository {

    @Autowired
    private PersistentLoginsRepository persistentLoginsRepository;
    @Autowired
    private UserRepository userRepository;

    @Override
    public void createNewToken(PersistentRememberMeToken token) {
      User user = userRepository.findByUserId(token.getUsername());
      PersistentLogins persistentLogins =
          new PersistentLogins(user, token.getSeries(), token.getTokenValue(), token.getDate());
      persistentLoginsRepository.save(persistentLogins);
    }

    @Override
    public void updateToken(String series, String tokenValue, Date lastUsed) {
      PersistentLogins persistentLogins = persistentLoginsRepository.findOne(series);
      persistentLogins.update(series, tokenValue, lastUsed);
      persistentLoginsRepository.save(persistentLogins);
    }

    @Override
    public PersistentRememberMeToken getTokenForSeries(String series) {
      PersistentLogins persistentLogins = persistentLoginsRepository.findOne(series);
      return new PersistentRememberMeToken(persistentLogins.getUser().getUserId(), series,
          persistentLogins.getToken(), persistentLogins.getLastUsed());
    }

    @Transactional
    @Override
    public void removeUserTokens(String username) {
      User user = userRepository.findByUserId(username);
      persistentLoginsRepository.deleteByUserId(user.getId());
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
  <input type="hidden" name="${_csrf.parameterName}" value="${_csrf.token}"/>
  <button type="submit" class="btn">Log in</button>
</form>
```

csrf 값을 출력하는 `<input type="hidden" name="${_csrf.parameterName}" value="${_csrf.token}"/>` 영역은
security taglib를 사용하면 `<sec:csrfInput />` 으로 출력 가능하다

PUT / DELETE Method나 multipart/form-data의 경우에는 `form action` 파라미터로 `?${_csrf.parameterName}=${_csrf.token}` csrf를 보내야한다

### Authoriy 생성

Authority는 Enum으로 생성하거나, DB에 저장한다

  ```java
  public enum Authority implements GrantedAuthority {
    ADMIN,
    USER;

    @Override
    public String getAuthority() {
      return this.toString();
    }
  }
  ```

### User Entity implements UserDetails

UserDetails 인터페이스를 구현한다. 인증후 유저객체(Authentication.principal)와 대응된다

```java
@ElementCollection(fetch = FetchType.EAGER)
@Enumerated(EnumType.STRING)
@Column(name="authority")
private List<Authority> authorities;

public User() {
  authorities = new ArrayList<>();
  authorities.add(Authority.USER);
}

@Override
public Collection<? extends GrantedAuthority> getAuthorities() {
  return this.authorities;
}
```

### implements UserDetailsService

UserDetailsService 인터페이스를 구현한다. SecurityConfig에서 설정한 내용에 따라서
여기서 구현한 loadUserByUsername Method가 인증 과정을 처리한다.

```java
@Component
public class CustomUserDetailService implements UserDetailsService {
  private UserRepository userRepo;

  @Autowired
  public CustomUserDetailService(UserRepository userRepo) {
    this.userRepo = userRepo;
  }

  @Override
  public UserDetails loadUserByUsername(String userId) throws UsernameNotFoundException {
    User user = userRepo.findByUserId(userId);
    if (user == null) {
      throw new UsernameNotFoundException(userId);
    }
    return user;
  }
}
```

## 가입 처리이후 로그인

```java
Authentication authentication =
    new UsernamePasswordAuthenticationToken(userDetails, null, userDetails.getAuthorities());
SecurityContextHolder.getContext().setAuthentication(authentication);
```

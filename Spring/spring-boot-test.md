# Spring Boot Test

<https://docs.spring.io/spring-boot/docs/current/reference/html/boot-features-testing.html>

spring-boot-starter-test에는 다음 패키지들이 들어있다

1. JUnit
1. Spring Test & Spring Boot Test
1. AssertJ : a fluent assertion lib
1. Hamcrest : a lib of matcher object
1. Mockito : a Java mocking lib
1. JSONassert : an Assertion lib for JSON
1. JsonPath : XPath for JSON

spring-security 사용시 테스트 의존성을 추가해야 한다

`testCompile('org.springframework.security:spring-security-test')`

## @SpringBootTest

테스트에 사용할 ApplicationContext를 생성함 (@ContextConfiguration)

**`@RunWith(SpringRunner.class)`** 와 함께 사용해야 함

### Bean

`@SpringBootTest` 어노테이션의 classes 속성을 통해서 빈을 생성할 클래스들을 지정할 수 있음.
classes 속성 : `@Configuration`이 선언된 설정파일의 `@Bean`으로 생성되는 빈도 등록된다

만일 classes 속성을 통해서 클래스를 지정하지 않으면 애플리케이션 상에 정의된 모든 빈을 생성

### TestConfiguration

기존에 Configuration을 커스터마이징 하려면 TestConfiguration 기능을 사용함.
TestConfiguration은 ComponentScan 과정에서 생성되고, TestConfig이 속한 테스트가 실행될때 정의된 빈을 생성하여 등록한다.

ComponentScan 과정에서 생성되므로 `@SpringBootTest`의 classes 속성을 이용하여 특정 클래스만을 지정했을 경우에는 TestConfiguation은 감지되지 않는다.
그런 경우 `@Import` 어노테이션을 통해서 직접 사용할 TestConfiguration을 명시함.

### MockBean

기존에 방식대로 Mock 객체를 생성해서 테스트하는 방법도 있지만, `@MockBean` 어노테이션을 사용해서 Mock 객체를 빈으로써 등록할 수 있다.
`@MockBean`으로 bean을 주입받는다면 (@Autowired 같은 어노테이션 등을 통해서) ApplicationContext는 Mock 객체를 주입해준다.

만약 @MockBean으로 선언한 객체와 같은 이름과 타입의 bean이 등록되어있다면 해당 빈은 새로운 Mock bean으로 대체된다.

### Properties

설정파일은 application.properties(또는 application.yml)이 기본적으로 사용된다. 테스트 중에는 설정이 기존과 달라질 필요가 있는 경우가 많은데 별도의 테스트를 위한 application.properties(또는 application.yml)을 지정할 수 있다.

```java
@RunWith(SpringBoot.class)
@SpringBootTest(properties = "classpath:application-test.yml")
public class FooTest {
  ...
}
```

### Web Environment

`@SpringBootTest`의 webEnvironment 파라미터를 통해 테스트 환경 설정을 할 수 있다.

### TestRestTemplate

MockMvc는 Servlet Container를 생성하지 않는 반면,
TestRestTemplate은 Servlet Container를 사용해서 실제 서버가 동작하는 것처럼 테스트를 수행할 수 있다.

MockMvc는 서버 입장에서 구현한 API를 통해 비즈니스 로직이 문제없이 수행되는지 테스트를 할 수 있다면,
TestRestTemplate은 클라이언트 입장에서 RestTemplate을 사용하듯이 테스트를 수행한다.

```java
@RunWith(SpringRunner.class)
@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
public class RestApiTest {
  @Autowired
  private TestRestTemplate restTemplate;

  @Test
  public void test() {
    ResponseEntity<Article> response = restTemplate.getForEntity("/api/articles/1", Article.class);
    assertThat(response.getStatusCode()).isEqualTo(HttpStatus.OK);
    assertThat(response.getBody()).isNotNull();
    ...
  }
}
```

## @RestClientTest

서버간 통신에서 현재 테스트케이스가 클라이언트 사이드가 되는 상황을 테스트 할 때 사용.

요청에 반응하는 가상의 Mock 서버를 만드는 개념이다. `MockRestServiceServer` 빈을 이용해서 테스트를 한다.

## @WebMvcTest

MockMvc에 관한 설정을 자동으로 수행해주는 어노테이션이다. 스캔범위를 찾아 MockMvc를 자동으로 설정하여 빈으로 등록한다.

### scan 대상

- @Controller
- @ControllerAdvice
- @JsonComponent
- Filter
- WebMvcConfigurer
- HandlerMethodArgumentResolver

### MockMvc

MockMvc를 사용해서 테스트 진행

Examples

```java
@RunWith(SpringRunner.class)
@WebMvcTest(ArticleApiController.class)
public class ArticleApiControllerTest {
  @Autowired
  private MockMvc mvc;
  @MockBean
  private ArticleService fooService;
}
```

### 비동기 테스트

```java
@Test
public void testGetArticle() throws Exception {
  Article expected = new Article(1, "kwseo", "good", "good content", now());

  given(articleService.findOneFromRemote(eq(1))).willReturn(expected);

  MvcResult result = mvc.perform(get("/api/articles/1")).andReturn();
  mvc.perform(asyncDispatch(result))      // asyncDispatch 필요
  .andExpect(status().isOk())
  .andExpect(jsonPath("@.id").value(1));
}
```

## @JsonTest

JSON serialization과 deserialization을 테스트

`JacksonTester<Article> jackson;`을 통해 테스트를 진행

## @DataJpaTest

테스트에 해당 어노테이션이 있으면 `@Entity`가 명시된 클래스를 스캔한다.

기본적으로 `@Transactional`을 포함하고 있고, 해제하려면 `@Transactional(propagation = Propagation.NOT_SUPPORTED)`을 사용하면 된다.

in-memory db를 쓰지 않으려면 `@AutoConfigureTestDatabase(replace = Replace.NONE)` 옵션을 이용한다.

`TestEntityManager` bean으로 테스트를 실시한다.

### @JdbcTest

JDBC테스트에 이용하며 `JdbcTemplate`을 사용한다.

### @DataMongoTest

`@Entity`가 아닌 `@Document`를 스캔하며 MongoTemplate을 사용한다.

in-memory db를 사용하지 않으려면 `@DataMongoTest(excludeAutoConfiguration = EmbeddedMongoAutoConfiguration.class)` 옵션을 이용한다.

## AssertJ

static import: `import static org.assertj.core.api.Assertions.*;`

단언문 시작: `assertThat(objectUnderTest)`

### 일반 조건문

```java
import static org.assertj.core.util.Sets.newLinkedHashSet;

static Set<String> JEDIS = newLinkedHashSet("Luke", "Yoda", "Obiwan");

Condition<String> jediPower = new Condition<>(JEDIS::contains, "jedi power");
```

#### is / isNot

```java
assertThat("Yoda").is(jedi);
assertThat("Vader").isNot(jedi);

try {
  // condition not verified to show the clean error message
  assertThat("Vader").is(jedi);
} catch (AssertionError e) {
  assertThat(e).hasMessage("expecting:<'Vader'> to be:<jedi>");
}
```

#### has/doesNotHave

```java
assertThat("Yoda").has(jediPower);
assertThat("Solo").doesNotHave(jediPower);

try {
  // condition not verified to show the clean error message
  assertThat("Vader").has(jediPower);
} catch (AssertionError e) {
  assertThat(e).hasMessage("expecting:<'Vader'> to have:<jedi power>");
}
```

#### 조건 결합

- AND 조건: allOf(Condition...)
- OR 조건: anyOf(Condition...)

메소드로 결합할 수 있다

```java
assertThat("Vader").is(anyOf(jedi, sith));
assertThat("Solo").is(allOf(not(jedi), not(sith)));

static Set<String> SITHS = newHashSet("Sidious", "Vader", "Plagueis");
Condition<String> sith = new Condition<>(SITHS::contains, "sith");
```

### Collection 검사

#### 컬렉션객체(iterable/array)에 대한 검증

```java
// import static org.assertj.util.Sets.newLinkedHashSet;
assertThat(newLinkedHashSet("Luke", "Yoda")).are(jedi);
assertThat(newLinkedHashSet("Leia", "Solo")).areNot(jedi);

assertThat(newLinkedHashSet("Luke", "Yoda")).have(jediPower);
assertThat(newLinkedHashSet("Leia", "Solo")).doNotHave(jediPower);

assertThat(newLinkedHashSet("Luke", "Yoda", "Leia")).areAtLeast(2, jedi);
assertThat(newLinkedHashSet("Luke", "Yoda", "Leia")).haveAtLeast(2, jediPower);

assertThat(newLinkedHashSet("Luke", "Yoda", "Leia")).areAtMost(2, jedi);
assertThat(newLinkedHashSet("Luke", "Yoda", "Leia")).haveAtMost(2, jediPower);

assertThat(newLinkedHashSet("Luke", "Yoda", "Leia")).areExactly(2, jedi);
assertThat(newLinkedHashSet("Luke", "Yoda", "Leia")).haveExactly(2, jediPower);
```

#### 특정 요소 포함여부

```java
assertThat(beasts).contains(direwolf);
assertThat(beasts).contains(werewolf);

// 위와 같다
assertThat(beasts).contains(werewolf, direwolf);

assertThat(beasts).doesNotContain(vampire);
```

#### 컬렉션 구성요소가 모두 같은가

순서에 상관없이 요소가 모두 일치하는지 확인한다

```java
assertThat(beasts).containsExactlyInAnyOrder(direwolf, werewolf);
assertThat(beasts).containsExactlyInAnyOrder(werewolf, direwolf);
```

컬렉션 요소의 순서까지 확인한다

```java
assertThat(beasts).containsExactly(direwolf, werewolf);
```

#### 컬렉션이 특정 요소로만 이루어져 있는가

```java
List<Monster> group = Lists.newArrayList(direwolf, direwolf, werewolf, werewolf, werewolf);

assertThat(group).containsOnly(direwolf, werewolf);
assertThat(group).containsOnly(werewolf, direwolf);
```

#### 컬렉션에 특정요소가 하나만 포함되어 있는가

```java
List<Monster> group = Lists.newArrayList(direwolf, direwolf, werewolf, vampire);

assertThat(group).containsOnlyOnce(vampire);
assertThat(group).containsOnlyOnce(werewolf);

// 아래의 결과는 false이다
assertThat(group).containsOnlyOnce(direwolf);
```

#### 컬렉션 요소들이 주어진 property를 가지고 있는가

```java
assertThat(beasts).extracting(Monster::getName).containsExactly("Direwolf", "Werewolf");
```

## Mockito

- Mock은 특정 행위에 대한 반환값을 정의하여 요청이 왔을 때 이를 재현한다
- Spy는 실제 객체를 감싼 객체이며 해당 객체에 요청/반환된 내용을 알아낼 때 사용된다

### 사용

- mock()/@Mock: Mock을 생성한다
  - `@Mock` 애노테이션의 경우 `MockitoAnnotations.initMocks(this);`를 실행해서 인스턴스를 생성한다
  - `when()`/`given()` 메소드는 Mock의 행위를 지정한다

- spy()/@Spy: 실제 객체를 감싼 Spy 생성

- @InjectMocks: 의존성을 주입받아 생성할 (이미 생성된 객체의 타입에 맞춰 DI) Mock 객체이다

- verify(): mock 객체의 행위 검증을 한다

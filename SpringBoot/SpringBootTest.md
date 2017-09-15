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

## @RunWith(SpringRunner.class)

스프링 테스트를 가능하게 해주는 annotation

## @DataJpaTest

JPA만 테스트 할 수 있음

## @WebMvcTest

### scan 대상

- @Controller
- @ControllerAdvice
- @JsonComponent
- Filter
- WebMvcConfigurer
- HandlerMethodArgumentResolver

### MockMvc

MockMvc를 사용해서 테스트 진행

### Examples

```java
given(this.fooService.getBarList())
  .willReturn(Arrays.asList(foobars));

mockMvc.perform(get("/"))
  .andExpect(status().isOk())
  .andExpect(view().name("home"))
  .andExpect(model().attributeExists("foobars"))
  .andExpect(model().attribute("foobars", IsCollectionWithSize.hasSize(1)));
```

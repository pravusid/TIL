# Spring

Spring을 하나의 단어로 요약하라고 한다면 Dependency Injection Framework

## Project Initialization

### Dependencies

- spring-core
- spring-context

### Configuration

`@Configuration` 애노테이션 사용 -> 상세내용 `spring-mvc-configuration-*.md` 파일 참고

## 작동 순서

- 요청받는 클래스
  - tomcat에 요청 (web.xml)
  - Front Controller
    - DispatcherServlet : Spring
    - FilterDispatcher : Struts2
    - ActionServlet : Struts1
  - 설정 파일 위치
    - Spring : applicationContext.xml -> WEB-INF
    - (Struts2 : struts.xml -> SRC)
    - (Struts : struts-config.xml -> SRC)
  - 클래스 분석

- 요청을 제어(Model)
  - ~Controller : java
  - Model -> DAO, VO, Model(Service)

- View : ViewResolver

1. DispatcherServlet(WebApplicationContext)
2. Model 찾기 : HandlerMapping(XML 파싱)
3. Controller(Model) : req, resp <==> DAO
4. req.setAttribute()/session()
5. DispatcherServlet
6. JSP 찾기 : ViewResolver

### Application Context

IoC Container는 보다 일반적인 개념, Spring의 IoC Container

- Bean Factory++
  - Spring's AOP
  - I18n capabilities (국제화)
  - Web Application Context

## 스프링 구조

- 컨테이너 -> 클래스(컴포넌트)를 모아서 관리함 : 컴포넌트간의 의존성을 낮추어 관리하기 위해서
  - 클래스 찾기: Dependency Lookup : getBean(id)
  - 클래스 제어 : 멤버변수 / 메소드

- 컴포넌트
  - 사용자 정의 데이터형 사용(~VO, ~DTO)
  - 액션 사용 기능 : 초기 구동시 메모리 할당
    - ~Manager
    - ~Service
    - ~DAO
    - ~Model(~Controller)

- XML, Annotation `@Component, @Controller, @Repository, @Service` : Metadata

- Dependency Injection (의존 주입) : 컨테이너를 통해 setter를 호출함(값을 넣는다)
  - setter DI
    - (xml) `p:~`
    - (xml) `<property name="" value=""/>`
  - construct DI
    - 생성자 매개변수에 값 주입

- AOP : 반복적 작업에 대해 콜백함수 `@Aspect`
  - 트랜잭션, 보안, 로그

- ORM

- MVC

- 출력 : JSTL, EL / spEL

## 컨테이너

- 컨테이너는 객체의 생명주기를 관리하는 역할

- 객체의 생명주기는 XML을 통해서 제작(생성~소멸)

- 종류
  - BeanFactory : core(DI:의존성 주입, DL:클래스 찾기)
  - └ ApplicationContext : core(DI,DL), AOP, 국제화
  - └ WebApplicationContext : core(DI,DL), AOP, 국제화, web(MVC)

- 동작 (HandlerMapping, HandlerAdapter)
  1. XML 읽기
  2. XML parsing
  3. 클래스 메모리 할당 : SAX -> id, class -> Map -> Reflection
      > 생성자 DI
  4. setter DI
  5. 메소드 호출 : init-method
  6. 활용 (개발자가 원하는 메소드 활용 -> DL)
  7. 메소드 호출 : destroy-method

  ```xml
  <!--생성자 DI-->
  <bean id="sa2" class="com.idpravus.sawon.Sawon2">
    <constructor-arg value="2"/>
    <constructor-arg value="심청이"/>
    <constructor-arg value="영업부"/>
  </bean>
  <!--Setter DI-->
  <bean id="sa" class="com.idpravus.sawon.Sawon"
    scope="prototype" (메모리 할당 방식)
    p:sabun="1" p:name="홍길동" p:dept="개발부" (setter DI)
    init-method="init" (init 호출)
  />
  <!--객체값 setter DI-->
  <bean id="member" class="com.idpravus.temp.Member"
    p:id="admin"
    p:name="춘향이"
    p:addr="서울"
  />
  <bean id="mc" class="com.idpravus.temp.MemberConfig"
    p:member-ref="member"
  />
  <!--다수 객체 setter DI-->
  <bean id="mc" class="com.idpravus.temp.MemberConfig">
    <property name="list">
      <list>
        <ref bean="member1"/>
        <ref bean="member2"/>
        <ref bean="member3"/>
      </list>
    </property>
  </bean>
  ```

## DI

### DI 필요성

```java
public class ComplexBusinessService {
  SortAlgortihm sortAlgotithm = new BubbleSortAlgorithm();
}
```

서비스를 유연하게 제공할 수 없다. 서비스가 버블소트에 의존성이 발생한다.

서비스를 호출하는 곳에서 필요한 알고리즘을 보낼 수 있다. (Inversion of Control)

```java
SortAlgortihm sortAlgorithm = new BubbleSortAlgorithm();
ComplexBusinessService businessService = new ComplexBusinessService(sortAlgorithm);
```

### Spring DI

```java
// 인터페이스에 자동주입시 클래스 특정지어 명시해야함
@Autowired
@Qualifier("fb")
// 위의 두가지를 동시에 하려면 아래의 코드 사용
@Resource(name="rb")
```

## spring의 설정 XML

1. 일반 클래스 등록 : application-context.xml
1. 데이터베이스 등록 : application-datasource.xml
1. 보안 등록 : application-security.xml
1. 빅데이터 등록 : application-hadoop.xml
1. AOP : application-aop.xml
1. 웹소켓 : application-websocket.xml

> web.xml --> application-*.xml

xml은 관련된 클래스끼리 분산해서 작성하고 마지막에는 합쳐서 사용한다.

## Annotation

annotation 사용을위해서는 application-context.xml 파일에
`<context:annotation-config/>` 을 선언해야함

`<context:component-scan base-package="com.idpravus.*"/>` 패키지 일괄등록

### 메모리 할당시 annotation

계층에 따라서 구분함, 인스턴스화 시점에 차이가 있고 계층에 따라 필요한 기능 지원
(ex: Repository의 경우 JDBC Exception handling 지원...)

1. `@Controller` : MVC pattern의 Controller
2. `@Service` : Business Services
3. `@Repository` : DAO등 Persistence 관련
4. `@Component` : Generic Component

### JavaEE 표준 (JSR-330) Annotation

1. `@Autowired`: `@Inject`
2. `@Component`: `@Named`

### Scope Annotation

싱글톤(singleton) 사용하지 않을 때 annotation (요청시마다 bean 생성) : `@Scope(ConfigurableBeanFactory.SCOPE_PROTOTYPE)`

### Bean Lifecycle

- `@PostConstruct`: bean 초기화 이전 실행됨
- `@PreDestroy`: bean instance가 제거되기 이전 실행됨

## 세션

- 세션은 유저마다 하나 씩만 부여된다
- JSP에서 세션 attribute value 호출 : `${sessionScope.key}`
- 세션 삭제
  - 전체 삭제 : session.invalidate();
  - 부분삭제 : session.removeAttribute(key);
- 세션 유효시간 설정

## AOP

구현체로 AspectJ가 많이 쓰임

### AOP Annotation

- @Aspect: Class 사용 정의
- @Before(execution("`*`(return-type) `com.idpravus.dao.MyDatabase.*(..)`(package, parameter)")
- @After-Returning( ) : 정상수행
- @After-Throwing( ) : 에러 발생시
- @After( ) : finally 상황(무조건 수행)

1. JoinPoint : 시점
2. PointCut : 호출 조건 (execution 조건 명시)
3. Advice : JoinPoint + PointCut
4. Aspect : Advice가 모인것

### AOP (Aspect Oriented Programming)

- Aspect : 공통기반
- class 설정 : Target -> Method Pointcut
- 시점 (joint-point)
  - `@Before`
  - `@AfterReturning`
  - `@AfterThrowing`
  - `@After`
  - `@Around`

- servlet-context.xml (AOP @annotation 사용시)

  ```xml
  <aop:aspectj-autoproxy/>
  ```

- Aspect 정의(annotation) : Before와 Around 같이쓰면 Around -> Before -> Around 순으로 실행됨

#### Pointcut 설정

- `execution()`
- `bean()`
- `within()`
- `@annotation()`
- `()`

### AOP 예제

  ```java
  @Aspect
  @Component
  public class MyDBAspect {
    @Autowired
    private DBConnection conn;

    @Before("execution(* com.idpravus.dao.StudentDAO.st*(..))")
    public void getConnection() {
      conn.getConnection();
      System.out.println("@Before Call");
    }

    @After("execution(* com.idpravus.dao.StudentDAO.st*(..))")
    public void disconnect() {
      conn.disconnect();
      System.out.println("@After Call");
    }

    @AfterReturning(value="execution(* com.idpravus.dao.StudentDAO.st*(..))", returning="ret")
    public void afterReturn(Object ret) {
      System.out.println("Return값:"+ret);
      System.out.println("@AfterReturning Call");
    }

    @AfterThrowing(value="execution(* com.idpravus.dao.StudentDAO.st*(..))", throwing="e")
    public void afterThrow(Throwable e) {
      System.out.println(e.getMessage());
      System.out.println("@AfterThrowing Call");
    }

    @Around("execution(* com.idpravus.dao.StudentDAO.st*(..))")
    public void around(ProceedingJoinPoint p) throws Throwable  {
      Object obj = null;
      long start = System.currentTimeMillis();
      System.out.println("@Around Call");
      obj = p.proceed(); // 호출되는 Method
      System.out.println("@Around End");
      long end = System.currentTimeMillis();
      System.out.println(p.getSignature().getName()+"수행시간:"+(end-start));
    }
  }
  ```

- Aspect를 xml로 정의할 때

  ```xml
  <aop:before method="getConnection" pointcut=""/>
  <aop:after method="disconnect" pointcut=""/>
  ```

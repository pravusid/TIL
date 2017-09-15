# 프로젝트 생성

## 최초 프로젝트 생성

### 빌드툴과 버전 선택

- <https://start.spring.io>에서 프로젝트 생성 Wizard를 실행할 수 있다.

- 빌드툴은 Gradle, Spring Boot 1.5.4버전을 선택하였다

### 프로젝트에 포함할 의존패키지 선택

필요에의해 추가 제거한다.

- Web : 스프링부트 기본 패키지
- AOP
- Security
- DevTools
- Rest Repositories
- Thymeleaf : 템플릿 엔진
- JPA : 자바 ORM 표준기술 - hibernate사용예정
- MySQL : db는 MariaDB사용, hibernate사용할 때 필요없으면 이건 삭제
- lombok : getter/setter 간소화 지원 라이브러리, 다른용도는 찾아봐야...

### 배포형식은 JAR or WAR

- 배포형식을 JAR로 선택하면 구동시 WAS없어도 되고 (사실 내장되어있음) 테스트 할때 서버가동 대신 app실행으로 시작 가능하다 (사실 서버실행)
- 실행방법 (평범한 상황에서 다 되는지 테스트 해봐야 함)
  - Run as - Spring Boot App
  - gradle - Tasks quick launcher
  - $ gradle bootRun
> Spring Tool Suite를 사용중이라면 Server탭 옆의 Boot Dashboard를 활용해서 run server를 수행할 수 있다.

## 프로젝트 기초 설정

### 패키지 구분

- 패키지 구분은 root, web, service, domain으로 한다
  - root : 프로젝트 설정이 들어갈 예정 (~Application.java 파일)
  - web : 컨트롤러
  - service : 서비스 영역
  - domain : Entity

### 기본설정 (Talsist.java)

> 프로젝트를 생성하면 기본적으로 @SpringBootApplication 어노테이션이 붙어있다.

- @SpringBootApplication은 스프링부트 기본설정 3종 신기가 한 번에 적용되어 있는 것이다
  - @Configuration
  - @EnableAutoConfiguration
  - @ComponentScan

### application.properties(yml) 설정

<https://docs.spring.io/spring-boot/docs/current/reference/html/common-application-properties.html> 참고
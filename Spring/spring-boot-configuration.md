# Configuration

## build.gralde

```groovy
buildscript {
    ext {
        springBootVersion = '1.5.16.RELEASE'
        thymeleafVersion = '3.0.1.RELEASE'
    }
    repositories {
        mavenCentral()
    }
    dependencies {
        classpath("org.springframework.boot:spring-boot-gradle-plugin:${springBootVersion}")
        classpath 'io.spring.gradle:dependency-management-plugin:1.0.5.RELEASE'
    }
}

apply plugin: 'java'
apply plugin: 'org.springframework.boot'

version = '1.0.0'
sourceCompatibility = 1.8

repositories {
    mavenCentral()
}

dependencies {
    compile("org.springframework.boot:spring-boot-starter-web")
    compile("org.springframework.boot:spring-boot-starter-aop")
    compile("org.springframework.boot:spring-boot-starter-security")
    compile("org.springframework.security:spring-security-jwt")
    compile("org.springframework.security.oauth:spring-security-oauth2")
    compile("org.springframework.boot:spring-boot-starter-data-jpa")
    compile("org.thymeleaf:thymeleaf:${thymeleafVersion}")
    compile("org.thymeleaf:thymeleaf-spring4:${thymeleafVersion}")
    compile("org.thymeleaf.extras:thymeleaf-extras-springsecurity4:${thymeleafVersion}")
    compile("mysql:mysql-connector-java")
    compile("com.h2database:h2")
    runtime("org.springframework.boot:spring-boot-devtools")
    testCompile("org.springframework.boot:spring-boot-starter-test")
    testCompile("org.springframework.security:spring-security-test")
}
```

## application.yml

<https://docs.spring.io/spring-boot/docs/current/reference/html/common-application-properties.html>

```yml
spring:
  profiles.active: dev

security:
  oauth2:
    resource:
      filter-order: 3
      jwt:
        key-value:
          -----BEGIN PUBLIC KEY-----
          MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAs1HLWx//4CM7PYYDdsE7
          0Ji42U/JIjkI8jlRp+Urx4P0/I2bPVZePw9WNDjcen560fmBlqt4NNzsEhOqi1tv
          25LvknTvNrUTl4L+u/jRY0kJpGcSBP/VtqdY0Yt6u+L/05VzMSLXr9PbhDw8nyhq
          7M/Y0wv7VQiFmeV4zK5lsYN787ii3ctouAg/VWFT2V5ZP24MLBGjb3s7Ipi8Wngp
          NIz/2umG/VYfnpZIy5cpqJLyLJKFJ02fTYsGcl6I23aQPpcsHhabEsjKUpF/ck4H
          XrBmadVQz9vFWxQVkUwpbMt827Zzkf2VnqIyVzmXsuY9gfiQeWLtmRvpw8KAZcOR
          jwIDAQAB
          -----END PUBLIC KEY-----

---

spring:
  profiles: dev
  h2:
    console:
      enabled: true
  datasource:
    driver-class-name: org.h2.Driver
    url: jdbc:h2:mem:idpravus;MODE=MYSQL;DB_CLOSE_DELAY=-1;DB_CLOSE_ON_EXIT=FALSE
    username: sa
    password:
    sql-script-encoding: UTF-8

  jpa:
    database-platform: org.hibernate.dialect.H2Dialect
    show-sql: true
    hibernate:
      ddl-auto: create-drop

logging:
  config: classpath:logback-spring-debug.xml

---

spring:
  profiles: production
  datasource:
    driver-class-name: com.mysql.jdbc.Driver
    url: jdbc:mysql://localhost/dbname?autoReconnect=true&useUnicode=true&characterEncoding=utf8
    username: user
    password: pwd
    initialize: false

  jpa:
    database-platform: org.hibernate.dialect.MySQL5InnoDBDialect
    hibernate:
      ddl-auto: none

server:
  port: 80
```

### Hibernate / JPA를 사용한 데이터 초기화

`spring.jpa.hibernate.ddl-auto` 옵션을 통해서 데이터 초기화 전략을 설정할 수 있음.

옵션: `none`(기본값), `validate`, `update`, `create`, and `create-drop`(기본값-embedded db without schema manager)

> classpath 루트에 `import.sql` 파일이 있다면 시작할 때 Hibernate가 이를 실행함: `spring.datasource.data`를 정의했다면 중복실행될 수 있음

Spring Boot는 시작할 때 자동으로 classpath 루트의 `schema.sql`과 `data.sql`을 실행시킨다.

또한, Spring Boot는 `schema-${platform}.sql`과 `data-${platform}.sql` 파일이 있다면 실행시켜 데이터베이스에 맞춘 스크립트 실행이 가능하다.
플랫폼 정의는 `spring.datasource.platform`값을 따른다

Spring Boot는 `spring.datasource.initialization-mode`값에 따라 자동으로 embedded DataSource의 schema를 생성한다. (기본값: `always`)

Spring Boot는 Spring JDBC initializer 동작시 fail-fast이므로, 스크립트에서 문제가 발생하면 어플리케이션 동작에 실패한다.
`spring.datasource.continue-on-error` 설정값(기본값: `false`)을 `true`로 변경하여 종료되지 않게 할 수 있다.

데이터 초기화를 방지하려면 다음 설정을 하면 된다

```properties
spring.datasource.initialization-mode=never # Property for Spring boot 2.0
spring.datasource.initialize=false # Property for Spring boot 1.0
```

초기화되는 데이터 인코딩 설정을 위해서는 다음을 사용하면 된다

`spring.datasource.sql-script-encoding=UTF-8`

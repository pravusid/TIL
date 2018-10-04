# Configuration

## build.gralde

```groovy
buildscript {
    ext {
        springBootVersion = '1.5.13.RELEASE'
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
apply plugin: 'eclipse'
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
    compile("org.springframework.security:spring-security-jwt:1.0.9.RELEASE")
    compile("org.springframework.security.oauth:spring-security-oauth2:2.3.3.RELEASE")
    compile("org.springframework.boot:spring-boot-starter-data-jpa")
    compile("org.thymeleaf:thymeleaf:${thymeleafVersion}")
    compile("org.thymeleaf:thymeleaf-spring4:${thymeleafVersion}")
    compile("org.thymeleaf.extras:thymeleaf-extras-springsecurity4:${thymeleafVersion}")
    runtime('com.h2database:h2')
    runtime("org.springframework.boot:spring-boot-devtools")
    testCompile("org.springframework.boot:spring-boot-starter-test")
    testCompile("org.springframework.security:spring-security-test")
}

jar {
    manifest {
        attributes  'Title': 'boot-vue', 'Version': 1.0, 'Main-Class': 'kr.pravusid.WebApplication'
    }
    dependsOn configurations.runtime
    from {
        configurations.compile.collect {it.isDirectory()? it: zipTree(it)}
    }
}
```

## application.yml

<https://docs.spring.io/spring-boot/docs/current/reference/html/common-application-properties.html>

```yml
spring:
  profiles.active: dev

---

spring:
  profiles: dev
  h2:
    console:
      enabled: true
  datasource:
    url: jdbc:h2:mem:dbname;DB_CLOSE_DELAY=-1;DB_CLOSE_ON_EXIT=FALSE
    username: sa
    password:
    data: classpath:import.sql

  # JPA로 사용할 데이터베이스 명시
  jpa:
    database-platform: org.hibernate.dialect.H2Dialect # auto-detected by default
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

  jpa:
    database-platform: org.hibernate.dialect.MySQL5InnoDBDialect # auto-detected by default
    hibernate:
      ddl-auto: update

server:
  port: 80
```

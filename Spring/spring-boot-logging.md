# Spring Boot Logging

스프링 부트는 commons-logging(JCL)을 사용한다

스프링 부트 애플리케이션은 slf4j와 Logback을 사용한다

기본 의존성에 slf4j 브릿지가 포함되어 있어 결국 slf4j-Logback을 호출하게 된다

기본 로그레벨은 INFO이다

## 설정

기본레벨을 DEBUG 레벨로 변경/실행 : `java -jar app.jar --debug`

기본레벨 설정은 application.properties에서 `debug: true`

두 방법은 root 단위설정이기 때문에 패키지별 설정을 할 것을 권장

`spring.output.ansi.enable = ALWAYS, DETECT, NEVER`

`logging.config= # 로깅 설정파일 위치 ex)classpath:logback.xml for Logback`

`logging.path = ./logs/spring.log`

`logging.file = ./logs/{파일명}`

우선순위 : logging.path < logging.file

`logging.exception-conversion-word = LOG_EXCEPTION_CONVERSION_WORD` : 로깅 예외발생시 출력되는 단어

### 범위

로그 레벨 : TRACE, DEBUG, INFO, WARN, ERROR, FATAL, OFF

`logging.level.{패키지} = {로그레벨}`

### 로깅 라이브러리에 따른 설정

Logback : `logback-spring.groovy`, `logback-spring.xml`

Log4j2 : `log4j2.xml`

Java Util Logging : `logging.properties`

### logback-spring.xml

```xml
<?xml version="1.0" encoding="UTF-8"?>
<configuration>
  <property name="LOG_TEMP" value="./logs"/>
  <include resource="org/springframework/boot/logging/logback/base.xml"/>

  <appender name="dailyRollingFileAppender"
            class="ch.qos.logback.core.rolling.RollingFileAppender">
    <prudent>true</prudent>
    <rollingPolicy class="ch.qos.logback.core.rolling.TimeBasedRollingPolicy">
      <fileNamePattern>logs/application.%d{yyyy-MM-dd}.log</fileNamePattern>
      <maxHistory>30</maxHistory>
    </rollingPolicy>
    <filter class="ch.qos.logback.classic.filter.ThresholdFilter">
      <level>INFO</level>
    </filter>

    <encoder>
      <pattern>%d{yyyy:MM:dd HH:mm:ss.SSS} %-5level --- [%thread] %logger{35} : %msg %n</pattern>
    </encoder>
  </appender>

  <logger name="org.springframework.web" level="INFO"/>
  <logger name="org.thymeleaf" level="INFO"/>
  <logger name="org.hibernate.SQL" level="INFO"/>
  <logger name="org.quartz.core" level="INFO"/>
  <logger name="kr.pravusid" level="INFO"/>

  <root level="INFO">
    <appender-ref ref="dailyRollingFileAppender"/>
  </root>
</configuration>
```

### application.yml

기본적으로 클래스패스에 `logback-spring.xml`이 있으면 해당 설정대로 작동한다.
설정파일에서 별도로 설정하면 `logback-spring.xml`을 override 하지 않고 추가로 설정내용을 실행한다.

```yml
spring.profiles: logging-info
logging:
  file: logs/application.log
  level:
    org.thymeleaf: INFO
    org.springframework.web: INFO
    org.hibernate.SQL: INFO
    org.quartz.core: INFO

---
spring.profiles: logging-debug
logging:
  file: logs/application.log
  level:
    org.thymeleaf: DEBUG
    org.springframework.web: DEBUG
    org.hibernate.SQL: DEBUG
    org.quartz.core: DEBUG

---
spring.profiles: logging-debug
logging:
  config: classpath:logback-spring-debug.xml
```

### 실행

`java -jar application.jar --spring.profiles.active=logging-debug,logging-daily`

`application.yml`에 설정한 `logging.*` 관련한 로그는 `logs/application.log`로 생성된다

일일로그는 위에서 설정한 패턴에 따라 `log/application{날짜}log`로 생성된다

## 사용

`Logger logger = LoggerFactory.getLogger(this.getClass());`

`logger.debug("param: {}", param var);`

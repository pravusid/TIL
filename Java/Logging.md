# Java Logging

## Level

로그 level은 다음이 일반적이고 라이브러리에 따라 약간씩 차이날 수 있음

1. FATAL : 치명적 오류


1. ERROR : 일반 오류
2. WARN : 잠재적 위험
3. INFO : 프로그램의 상황
4. DEBUG : 디버깅 정보
5. TRACE : 세분화된 Debug 정보

## commons-logging

Log4j 이외의 로그 라이브러리가 있기 때문에 표준화를 위해서 Apache에서 만든 라이브러리

표준화 사용법만을 지원하기 때문에 로그를 출력하는 구현체(log4j와 같은)가 별도로 필요하다

### 설정

commons-logging.properties 파일 생성하고 다음 내용 설정

`org.apache.commons.logging.Log = org.apache.commons.logging.impl.Log4JLogger`

### 사용법

`Log log = LogFactory.getLog(this.getClass());`

## Log for Java (Log4j)

Apache에서 관리하는 로깅 라이브러리

### 설정

yaml 파일로 설정 , xml, json도 가능 <http://logging.apache.org/log4j/2.x/manual/configuration.html> 참고

### 사용법

`Logger log = LogManager.getLogger("{LOGGER_NAME}");`

## Simple Logging Facade for Java (slf4j)

Facade 패턴으로 만들어져있다. 다른 로그 구현체를 바인딩해서 쓸 수있다. 바인딩 하지 않으면 기본 로거로 작동한다.

commons-logging의 단점을 보완한 부분이 있다 

```java
if( logger.isDebugEnabled() ){
    logger.debug("param: " + param var);
}
```

특정 로그레벨일 때만 처리하도록 명시 하지 않아도 내부적으로 처리한다

```java
 logger.debug("param: {}", param var);
```

### 설정

simplelogger.properties

````properties
org.slf4j.simpleLogger.logFile={경로}
org.slf4j.simpleLogger.defaultLogLevel=DEBUG
````

기본 레벨은INFO이다

### 브릿지 라이브러리

?-over-slf4j.jar 형식의 라이브러리를 브릿지 라이브러리라 한다.

프레임워크 기본값에서 혹은 사용하던 로깅라이브러리 구문때문에 로깅라이브러리를 교체하였을때, 기존 라이브러리 의존성에서 오류가 날 수도 있다. 이럴경우에는 slf4j에서 제공하는 브릿지 라이브러리를 사용하면 해결 할 수 있다. 예를 들어 log4j 라이브러리 의존성이 필요한 경우라면, log4j-over-slf4j.jar 파일을 라이브러리에 추가해 주면 된다.

### 사용법

`Logger logger = LoggerFactory.getLogger(this.getClass());`

## Logback

log4j의 개발자 Ceki Gulcu가 기존 로깅 라이브러리를 대체하기 위해 slf4j와 logback을 만들었다.

## 설정

<https://logback.qos.ch/manual/index.html> 주소를 참고하자



#### 
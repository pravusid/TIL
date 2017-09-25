# Spring MVC 관련 Tomcat 설정

## Get request에서 인코딩 처리

server.xml

```xml
<!-- Encoding Filter(GET) -->
<Connector connectionTimeout="20000" port="8080" protocol="HTTP/1.1" redirectPort="8443" URIEncoding="UTF-8"/>
```

## eclipse-maven-plugin 배포 : pom.xml에 plugin 추가

- run as config에 goals : `tomcat7:redeploy`

  ```xml
  <plugin>
      <groupId>org.apache.tomcat.maven</groupId>
      <artifactId>tomcat7-maven-plugin</artifactId>
      <version>2.2</version>
      <configuration>
          <path>/</path>
          <url>http://211.238.142.123:80/manager/text</url>
          <username>itmenu</username>
          <password>unemti</password>
      </configuration>
  </plugin>
  ```

- tomcat-users.xml
  ```xml
  <role rolename="admin-gui"/>
  <role rolename="manager-gui"/>
  <role rolename="manager-script"/>
  <user username="tomcat" password="tomcat"  roles="admin-gui,manager-gui,manager-script" />
  ```

- 배포를 한 적이 있다면 다음 배포 이전에 mvn clean을 실행

## 커넥션 풀 관련 설정

[NAVER D2 DBCP 이해하기](http://d2.naver.com/helloworld/5102792)

속성 이름 | 설명
--- | ---
initialSize | BasicDataSource 클래스 생성 후 최초로 getConnection() 메서드를 호출할 때 커넥션 풀에 채워 넣을 커넥션 개수
maxActive |동시에 사용할 수 있는 최대 커넥션 개수(기본값: 8)
maxIdle | 커넥션 풀에 반납할 때 최대로 유지될 수 있는 커넥션 개수(기본값: 8)
minIdle | 최소한으로 유지할 커넥션 개수(기본값: 0)

개수 관련한 조건

1. maxActive >= initialSize
1. maxIdle >= minIdle
1. maxActive = maxIdle

> maxActive > maxIdle 설정에서 커넥션 사용 후 반납되는 상황을 가정하자. maxIdle을 넘어선다면 매번 커넥션이 생성되었다 닫히는 상황이 발생할 수 있다.
# Spring MVC Codes

## server.xml (tomcat)

```xml
<!-- Encoding Filter(GET) -->
<Connector connectionTimeout="20000" port="8080" protocol="HTTP/1.1" redirectPort="8443" URIEncoding="UTF-8"/>
```

## web.xml

```xml
<context-param>
  <param-name>contextConfigLocation</param-name>
  <param-value>/WEB-INF/configuration/root-context.xml</param-value>
</context-param>

<listener>
  <listener-class>org.springframework.web.context.ContextLoaderListener</listener-class>
</listener>
<!-- xml파일 지정되지 않으면 기본값(<servlet-name>-servlet.xml)으로 자동지정됨 -->
<!-- xml파일은 단일파일 or 와일드카드 사용해서 여러개 지정 가능하다 -->
<servlet>
  <servlet-name>dispatcher</servlet-name>
  <servlet-class>org.springframework.web.servlet.DispatcherServlet</servlet-class>
  <init-param>
    <param-name>contextConfigLocation</param-name>
    <param-value>/WEB-INF/configuration/appServlet/*-context.xml</param-value>
  </init-param>
  <!-- 서버 가동과 동시에 Servlet 인스턴스 생성 -->
  <!-- <load-on-startup>1</load-on-startup> -->
</servlet>
<servlet-mapping>
  <servlet-name>dispatcher</servlet-name>
  <url-pattern>/</url-pattern>
</servlet-mapping>

<!-- Encoding Filter(POST) -->
<filter>
  <filter-name>encodingFilter</filter-name>
  <filter-class>org.springframework.web.filter.CharacterEncodingFilter</filter-class>
  <init-param>
    <param-name>encoding</param-name>
    <param-value>UTF-8</param-value>
  </init-param>
</filter>
<filter-mapping>
  <filter-name>encodingFilter</filter-name>
  <url-pattern>/*</url-pattern>
</filter-mapping>
```

## root-context.xml

```xml
<!-- component scan -->
<context:component-scan base-package="com.idpravus.*">
  <context:exclude-filter expression="org.springframework.stereotype.Controller" type="annotation" />
</context:component-scan>

<!-- 서버정보를 properties에서 읽어온다 -->
<util:properties id="db" location="/WEB-INF/configuration/db.properties" />
<bean id="ds" class="org.apache.commons.dbcp.BasicDataSource"
  p:driverClassName="#{db['driver']}"
  p:url="#{db['url']}"
  p:username="#{db['username']}"
  p:password="#{db['password']}" />

<!-- Mapper 등록방법1 : scan 이용 -->
<bean class="org.mybatis.spring.SqlSessionFactoryBean" p:dataSource-ref="ds" />
<mybatis-spring:scan base-package="com.idpravus.*" />

<!-- Mapper 등록방법2 : 직접 등록 -->
<bean id="ssf" class="org.mybatis.spring.SqlSessionFactoryBean"
  p:dataSource-ref="ds" />
<bean id="bMapper" class="org.mybatis.spring.mapper.MapperFactoryBean"
  p:sqlSessionFactory-ref="ssf"
  p:mapperInterface="com.idpravus.databoard.dao.DataBoardMapper" />

<!-- Transaction Manager -->
<bean id="transactionManager" class="org.springframework.jdbc.datasource.DataSourceTransactionManager">
    <property name="dataSource" ref="ds"/>
</bean>
<tx:annotation-driven transaction-manager="transactionManager" />
```

## dispatcher-servlet.xml

### servlet-context.xml

```xml
<!-- component scan -->
<context:component-scan base-package="com.idpravus.*" use-default-filters="false">
  <context:include-filter expression="org.springframework.stereotype.Controller" type="annotation" />
</context:component-scan>

<!-- 리소스 처리 -->
<mvc:resources mapping="/resources/**" location="/resources/" />
<mvc:resources mapping="/img/**" location="/img/" />
<mvc:resources mapping="/js/**" location="/js/" />
<mvc:resources mapping="/css/**" location="/css/" />

<!-- @ResponseBody로 객체 반환 -->
<bean id="jacksonMessageConverter" class="org.springframework.http.converter.json.MappingJackson2HttpMessageConverter"/>

<!-- @ResponseBody로 String 처리할때 한글처리 -->
<mvc:annotation-driven>
  <mvc:message-converters>
    <bean class="org.springframework.http.converter.StringHttpMessageConverter">
      <property name="supportedMediaTypes">
        <list><value>text/html;charset=UTF-8</value></list>
      </property>
    </bean>
  </mvc:message-converters>
</mvc:annotation-driven>

<!-- ViewResolver : JSP -->
<bean id="viewResolver" class="org.springframework.web.servlet.view.InternalResourceViewResolver"
  p:prefix="/"
  p:suffix=".jsp" />

<!-- ViewResolver : apache tiles -->
<bean id="tilesViewResolver"
  class="org.springframework.web.servlet.view.UrlBasedViewResolver"
  p:order="1">
  <property name="viewClass"
    value="org.springframework.web.servlet.view.tiles2.TilesView" />
</bean>
<bean id="tilesConfigurer"
  class="org.springframework.web.servlet.view.tiles2.TilesConfigurer">
  <property name="definitions">
    <list>
      <value>/WEB-INF/configuration/appServlet/tiles-config.xml</value>
    </list>
  </property>
</bean>

<!-- 파일처리 -->
<bean id="multipartResolver" class="org.springframework.web.multipart.commons.CommonsMultipartResolver"/>
  <task:annotation-driven/>
```

### security-context.xml

```xml
<!-- bcrypt DI -->
<beans:bean id="bcryptPasswordEncoder" class="org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder" />

<!-- 인터셉터 -->
<mvc:interceptors>
  <mvc:interceptor>
    <!-- 관리자여부 확인 -->
    <mvc:mapping path="/admin/**" />
    <beans:bean class="com.idpravus.users.AdminInterceptor" />
  </mvc:interceptor>
  <mvc:interceptor>
    <!-- 로그인 여부 확인 -->
    <mvc:mapping path="/signup/after" />
    <mvc:mapping path="/signup/addinfo" />
    <beans:bean class="com.idpravus.users.UsersInterceptor" />
  </mvc:interceptor>
</mvc:interceptors>
```

## ViewResolver로 Tiles 활용

default.jsp

```jsp
<%@ taglib prefix="tiles" uri="http://tiles.apache.org/tags-tiles" %>
<tiles:insertAttribute name="header"/>
```

tiles.xml

```xml
<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE tiles-definitions PUBLIC
"-//Apache Software Foundation//DTD Tiles Configuration 3.0//EN"
"http://tiles.apache.org/dtds/tiles-config_3_0.dtd">
<tiles-definitions>
  <definition name="default" template="/WEB-INF/views/default.jsp">
    <put-attribute name="header" value="/WEB-INF/views/header.jsp"/>
    <put-attribute name="navi" value="/WEB-INF/views/navi.jsp"/>
    <put-attribute name="body" value="/WEB-INF/views/body.jsp"/>
    <put-attribute name="footer" value="/WEB-INF/views/footer.jsp"/>
  </definition>
  <definition name="*" extends="default">
    <put-attribute name="body" value="/WEB-INF/views/{1}.jsp"/>
  </definition>
  <definition name="*/*" extends="default">
    <put-attribute name="body" value="/WEB-INF/views/{1}/{2}.jsp"/>
  </definition>
  <definition name="*/*/*" extends="default">
    <put-attribute name="body" value="/WEB-INF/views/{1}/{2}/{3}.jsp"/>
  </definition>
</tiles-definitions>
```

## 게시판

### 페이징

```java
private int rowSize; // 한 화면에 표시할 행
private int total; // 전체 게시물 수
private int start; // 출력 화면의 시작 행
private int end; // 출력 화면의 종료 행

private int blockSize;
private int totalPage; // 전체 페이지 수
private int startBlock; // block의 시작
private int endBlock; // block의 마지막
private int prevBtn; // 이전 block 버튼
private int nextBtn; // 다음 block 버튼

private int page;

{
  rowSize = 15;
  blockSize = 10;
}

public Map<String, Integer> calcPage(int total) {
  this.total = total;
  if (page == 0) { page = 1; }

  end = rowSize * page;
  start = end - rowSize + 1;
  if (end>total) { end=total; }

  Map<String, Integer> map = new HashMap();
  map.put("end", end);
  map.put("start", start);

  calcBlock(map);

  return map;
}

private void calcBlock(Map<String,Integer> map) {
  totalPage = (int) (Math.ceil((float)total/rowSize));
  startBlock = page - (page - 1) % blockSize;
  endBlock = startBlock + blockSize - 1;
  if (endBlock > totalPage) { endBlock = totalPage; }
  prevBtn = (startBlock==1)? 1: startBlock-1;
  nextBtn = (endBlock==totalPage)? totalPage: endBlock+1;
}
```

### 게시판 파일 업로드

```java
@RequestMapping("main/board_insert_ok.do")
public String board_insert_ok(DataBoardVO vo) {
  List<MultipartFile> list = vo.getUpload();
  if (list.isEmpty()) {
    vo.setFilename("");
    vo.setFilesize("");
    vo.setFilecount(0);
  } else {
    StringBuffer strName = new StringBuffer();
    StringBuffer strSize = new StringBuffer();
    for(MultipartFile mf : list) {
      try {
        String fileName = mf.getOriginalFilename();
        Long fileSize = mf.getSize();
        mf.transferTo(new File("c:\\upload\\"+fileName));
        strName.append(fileName+",");
        strSize.append(fileSize+",");

      } catch (IllegalStateException e) {
        e.printStackTrace();
      } catch (IOException e) {
        e.printStackTrace();
      }
    }
    vo.setFilename(strName.substring(0, strName.length()-1));
    vo.setFilesize(strSize.substring(0, strSize.length()-1));
    vo.setFilecount(list.size());
  }
  service.dataBoardInsert(vo);
  return "redirect:board_list.do";
}
```

### 파일 다운로드

```java
@RequestMapping("main/board_download")
public void board_download(String fn, HttpServletResponse resp) {
  try {
    File file = new File("c:\\upload\\"+fn);
    resp.setHeader("Content-Disposition", "attatchment;filename="+URLEncoder.encode(fn, "utf-8"));
    resp.setContentLength((int)file.length());

    BufferedInputStream bis = new BufferedInputStream(new FileInputStream(file));
    BufferedOutputStream bos = new BufferedOutputStream(resp.getOutputStream());
    byte[] b = new byte[1024];
    while (true) {
      if (bis.read(b)==-1) {break;}
      bos.write(b);
    }
    bis.close();
    bos.close();

  } catch (Exception e) {
    e.printStackTrace();
  }
}
```

### @ResponseBody 이용

#### ResponseBody 일반 활용

```java
@RequestMapping("main/board_update_ok.do")
@ResponseBody
public String board_update_ok(DataBoardVO vo, int page) {
  boolean bChk = false;
  String send = null;
  if (bChk == true) {
    send = "<script>"
        +"location.href=\"board_content.do?no="+vo.getNo()+"&page="+page+"\";"
        +"</script>";
  } else {
    send = "<script>"
        +"alert(\"비밀번호가 일치하지 않습니다\");"
        +"history.back();"
        +"</script>";
  }
  return send;
}
```

#### @ResponseBody에서 Object를 JSON으로 변환해서 반환

pom.xml

```xml
<dependency>
    <groupId>com.fasterxml.jackson.core</groupId>
    <artifactId>jackson-databind</artifactId>
    <version>2.7.3</version>
</dependency>
```

jackson example

```java
@RequestMapping("/login")
public @ResponseBody UsersVO login(UsersVO vo) {
  vo = dao.selectUserData(email);
  return vo;
}
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
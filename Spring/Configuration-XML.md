# Spring 설정 XML편

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

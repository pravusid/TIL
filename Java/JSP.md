# JSP (Java Server Pages)

<!-- TOC -->

- [JSP (Java Server Pages)](#jsp-java-server-pages)
  - [JSP란](#jsp란)
  - [JSP 코드를 기재할 수 있는 영역](#jsp-코드를-기재할-수-있는-영역)
  - [JSP 내장객체](#jsp-내장객체)
  - [세션(session)](#세션session)
  - [쿠키 (Cookie)](#쿠키-cookie)
  - [액션태그](#액션태그)
  - [EL (표현언어)](#el-표현언어)
    - [연산자](#연산자)
  - [JSTL (JSP Standard Tag Library)](#jstl-jsp-standard-tag-library)
    - [core](#core)
    - [fmt](#fmt)
    - [기타](#기타)
  - [커스텀 태그](#커스텀-태그)
    - [커스텀 태그 구현](#커스텀-태그-구현)
  - [Servlet](#servlet)
    - [서블릿 개발절차](#서블릿-개발절차)
    - [서블릿 계보](#서블릿-계보)
    - [Servlet Mapping(web.xml)](#servlet-mappingwebxml)
    - [서블릿의 생명주기](#서블릿의-생명주기)
    - [요청 단계](#요청-단계)
  - [코드 조각 자동포함 기능](#코드-조각-자동포함-기능)
  - [MVC 패턴](#mvc-패턴)
    - [fowarding을 이용해서 request객체를 다른곳에서도 이용하기](#fowarding을-이용해서-request객체를-다른곳에서도-이용하기)
    - [MVC에서 컨트롤러의 역할](#mvc에서-컨트롤러의-역할)

<!-- /TOC -->

## JSP란

문법구성

- 이 영역들은 곧 변환될 서블릿의 한 위치를 차지하게 된다
- <%@ %> : 지시영역, @ : 지시자
  @문자를 기재하는 영역으로 현재 페이지에 대한 설정정보를 작성할 수 있다.
                ex) 인코딩방식, 파일형식(html, xml...)
- <%! %> : 선언부 : 서블릿의 멤버영역
- <% %> : 스크립틀릿 : 서블릿의 service 메서드영역, 즉 요청을 처리하는 영역

## JSP 코드를 기재할 수 있는 영역

- 지시영역 @지시자
  - page : jsp에 대한 정보
    - contentType : text/html, text/xml
    - import : 자바 라이브러리 읽기
    - errorPage
      - 400 : Bad Request
      - 404 : 파일이 없는 경우
      - 500 : 소스 번역 에러
      - 401 : 인증에 문제가 있을 때
      - 415 : 변환코드 문제 -> EUC-KR / ECU-KR
      - 200 : 정상수행

  - taglib : 자바문법(제어문) - > 태그형
    - _core_ : 제어문, url
    `<c:forEach>, <c:if>, <c:redirect>`
    - _format_
    `<fmt:formatDate> SimpleDateFormat`
    - sql -> DAO
    - xml -> JAXP, **JAXB**
    - fn -> function (String, List)

  - include
    ```java
    <%@ include %> // 정적 : 컴파일 하기전 java파일이 합쳐진다
    <% jsp:include %> // 동적 : 컴파일 후 html이 합쳐진다
    ```

## JSP 내장객체

서블릿클래스명 | JSP내장객체명
-------|-------
HttpServeletRequest | request
HttpServeletResponse | response
HttpSession | session(세션)
PageContext | pageContext
PrintWriter | out
ServletContext | application (웹사이트의 전역정보(환경설정))
ServletConfig | config
Exception | exception
Object | page

## 세션(session)

- 서버와 클라이언트간 접속상태
  - stateless : 연결이 비지속적
  - stateful : 연결상태의 유지

- 서버는 클라이언트를 구분하기 위해 일렬번호를 발생시키는데 일렬번호를 보유한 객체가 바로 session이다.
  브라우저를 닫거나, 일정시간이 경과(session timeout)한 경우 세션 종료

세션 유효시간 : web.xml 파일에서

```xml
<session-config>
  <session-timeout>x분</session-timeout>
</session-config>
```

세션 비활성화 : 로그아웃에 사용 `session.invalidate()`

## 쿠키 (Cookie)

쿠키의 요소 : 이름, 값, 유효시간, 도메인, 경로

활동 | 명령어
------- | -------
쿠키 생성 | `Cookie cookie = new Cookie("cookieName", URLEncoder.encode("value", "utf-8"));`
쿠키 값 얻기 | `Cookie[] cookies = request.getCookies();`
쿠키 삭제 | `cookie.setMaxAge(0);` // 초 단위
쿠키 도메인 설정 | `cookie.setDomain(".idpravus.com");`
쿠키 적용 웹경로 | `cookie.setPath("/data");`

## 액션태그

javaEE 에서는 서버측에서 실행될 수 있는 태그가 지원된다.
코드량을 줄이기 위해서 사용한다.(편의를 위해)

```java
<jsp:useBean id="product" class="com.fashion.product.Product"/>
<jsp:setProperty property="*" name="product"/>

// jsp 흐름과 관련
<jsp:include page="포함할 페이지" />
<jsp:forward page="이동할 페이지" />
```

## EL (표현언어)

브라우저 출력을 위해 `<%= %>`대신 사용

- 형식 `${ }`

```java
request.setAttribute("id", "admin);
-> <%=request.getAttribute("id")%>
-> ${requestScope.id} -> requestScope. 생략가능 -> ${id}

session.setAttribute("id", "admin");
-> <%=session.getAttribute("id")%>
-> ${sessionScope.id}

request.getParameter("id"); // 보통 Model에서 처리
-> ${param.id}
```

### 연산자

- 산술연산자 ( +, -, *, /, % )
  - 나누기(/)는 double로 계산 5/2 == 2.5
  - 더하기(+)에 문자열 결합 없음(jsp2.3부터 지원)
    - "aaa"+10 == "aaa10" (X)
    - ${"10"+"20"} == 30, ${null+10} == 10 (null은 0으로 취급한다)
    - / (div), % (mod)
- 논리연산자 ( &&, || ) -> true / false
- 비교연산자 ==, !=, <, >, <=, >=
- 삼항연산자 (조건)? 값1: 값2;
- empty값 -> "" -> true / false

## JSTL (JSP Standard Tag Library)

### core

제어문, url

- `<%@ taglib prefix="c" uri="http://java.sun.com/jsp/jstl/core"%>`

- set
  - `<c:set var="name" value="홍길동">`
  - java에서는 `request.setAttribute("name", "홍길동")`

- for
  - `<c:forEach var="i" begin="1" end="9" step="1"></c:forEach>`
  - `<c:forEach var="vo" items="list" varStatus="s"></c:forEach>`

- if (else 없음)
  ```java
  <c:if test="조건">처리</c:if>
  <c:if test="!조건">else처리</c:if>
  ```

- choose (다중if문)
  ```java
  <c:choose>
  <c:when test="조건">처리</c:when>
  <c:otherwise>처리</c:otherwise>
  </c:choose>
  ```

- token (StringTokenize)
  - `<c:forTokens var="color" items="red,green,blue" delims=",">${color}</c:forTokens>`

- url(url 인코딩 후 생성)
  ```java
  <c:url value="URL" [var="varName"] [scope="영역"]>
    <c:param name="이름" value="값"/>
  </c:url>
  ```
  - var속성의 변수에 url을 생성해서 저장한다.
  - param tag로 파라미터를 추가 할 수 있다.

- redirect -> sendRedirect(url)
  ```java
  <c:redirect url="URL" [context="컨텍스트경로"]>
    <c:param name="이름" value="값"/>
  </c:redirect>
  ```

### fmt

숫자, 날짜변환등의 포맷 태그

- `<%@ taglib prefix="fmt" uri="http://java.sun.com/jsp/jstl/fmt" %>`
- SimpleDateFormat : `<fmt:formatDate value="${today}" pattern="yyyy-MM-dd"/>`

### 기타

- _sql_ -> DAO
- _xml_ -> SAX
- _fn_ : 함수(String, List) -> JAVA

## 커스텀 태그

- 태그 파일은 WEB-INF/tags 폴더나 그 하위 폴더에 위치하며 .tag나 .tagx 확장자를 갖는다
- 태그들은 jsp와 마찬가지로 기본 객체를 사용 할 수 있다.
- 태그이름은 파일명과 동일하다
- 태그파일 속성설정
  속성 | 설명
  ---|---
  name | 속성의 이름
  description | 속성에 대한 설명
  required | 속성의 필수여부
  rtexprvalue | 속성값으로 표현식을 사용할 수 있는지 여부
  type | 속성값의 타입
  fragment | `<jsp:attribute>` 액션태그로 속성값을 전달 할 때
- 사용된 속성값은 .tag파일에서 ${속성명} 으로 불러와 쓸 수 있다.

### 커스텀 태그 구현

- 속성 하나를 명시한 커스텀 태그 예제

```java
// tag 파일
<%@ tag language="java" pageEncoding="UTF-8"%>
// 속성값 설정
<%@ attribute name="title" required="true" %>
<%= title %>

// jsp 파일
<%@ page contentType="text/html; charset=UTF-8" %>
<%@ taglib prefix="tf" tagdir="/WEB-INF/tags" %>
<tf:header title="글 제목"/>
```

- 속성 설정시 fragment가 true 일 경우 값 전달 시 액션태그 사용해야 함
  attribute 태그 내에서 표현식과 스크립틀릿은 사용불가

```java
<tf:태그명>
  <jsp:attribute name="속성명">${'내용'}</jsp:attribute>
</tf:태그명>
```

- 동적 속성 전달
  태그를 범용적으로 사용하기 위해서 정의하는 방식

```java
// tag 파일
<%@ tag language="java" pageEncoding="UTF-8"%>
<%@ tag dynamic-attributes="optionMap" %>
<%@ attribute name="name" required="true" %>
<%@ taglib prefix="c" uri="http://java.sun.com/jsp/jstl/core" %>
<select name="${ name }">
  <c:forEach items="${ optionMap }" var="option">
  <option value="${ option.key }">${ option.value }</option>
  </c:forEach>
</select>

// jsp 파일
<tf:select name="genre" rock="락" ballad="발라드" metal="메탈"/>
```

- 속성 값으로 태그몸체의 내용을 받을 때

```java
// 태그파일
<%@ tag body-content="scriptless" pageEncoding="utf-8" %>
/* 몸체 내용의 EL이나 액션태그, 태그등이 문자열로 처리됨(실행되지 않음)
<%@ tag body-content="tagdependent" pageEncoding="utf-8" %> */
<%@ attribute name="length" type="java.lang.Integer" %>
<%@ attribute name="trim" %> // boolean 값으로 받음
<jsp:doBody var="content" scope="page"/>
<%
  String content = (String)jspContext.getAttribute("content");
  out.print(length+content+"처리됨");
%>
// <c:forEach><jsp:doBody/></c:forEach> 로 반복출력 가능

// jsp파일
<tf:removeHtml length="15">
  <h1>몸체 내용 전송 테스트<h1>
</tf:removeHtml>
```

- 커스텀 태그에서 몸체 내용에서 사용할 변수 생성

```java
// 태그파일
<%@ variable name-given="변수명" variable-class="java.lang.Integer" scope="NESTED/AT_BEGIN/AT_END" %>
<c:set var="변수명" value="값" />
<jsp:doBody/> // 몸체 내용을 실행하기 전에 sum값을 태그를 호출한 페이지에 전달
```

## Servlet

### 서블릿 개발절차

1. 클래스 작성하고 HttpServlet
1. 적절한 메서드 오버라이드 doGet()
1. 서블릿 매핑(Mapping) : WEB-INF/web.xml

### 서블릿 계보

- Servlet
  - GenericServlet
    - HttpServlet

### Servlet Mapping(web.xml)

```xml
<servlet>
  <servlet-name>servletName</servlet-name>
  <servlet-class>com.pravusid.mvc.DispatcherSevlet</servlet-class>
  <!-- init에서 로드되는 파일들 -->
  <init-param>
    <param-name>configLocation</param-name>
    <param-value>/WEB-INF/mapping/servlet-mapping.json</param-value>
  </init-param>
</servlet>
<servlet-mapping>
  <servlet-name>servletName</servlet-name>
  <url-pattern>.*.do(요청받을 경로)</url-pattern>
</servlet-mapping>
```

### 서블릿의 생명주기

|요청처리 순서도|
|---|
|클라이언트요청 >> __[service 호출하는 별도의 thread 생성]__ 요청객체(클라이언트 정보)==___req___ // 응답객체(서버의 응답)==___resp___ >>|
|web.xml(mapping) >> 서블릿 인스턴스 생성 >> init(__서버가 정보주입__) (인스턴스 생성이후 init method 호출) >>|
|__[thread가 호출]__ service(___req___, ___resp___) (요청처리) >> doGet(___req___, ___resp___), doPost(___req___, ___resp___) >>|
|서버가 데이터 처리이후 response 객체로 클라이언트에게 스트림으로 전송 >>|
|do~~~() 메서드가 끝날때 ___req___, ___resp___, ___Thread___ 도 사라짐 >>|
|인스턴스는 서버가 종료될때 destroy() method로 사라짐|

### 요청 단계

|요청 동작 순서|
|---|
|client|
|DNS서버|
|socket(TCP/IP)|
|WebServer(apache,IIS)|
|WebApplicationServer(tomcat)|

서버에서

1. 접속 : httpd

1. 파일찾기 : html, xml, jsp

    있음 : html, xml -> 전송 / jsp -> (톰캣이)변환 후 전송

    없음 : 404error

1. 톰캣에서 java파일을 class로 생성 or 수정

1. 서블릿 인스턴스 메모리 로드 -> init(), service()

1. service()에서 클라이언트에게 응답

## 코드 조각 자동포함 기능

모든 jsp페이지에서 동일한 코드를 삽입해야 할 경우 web.xml에서 설정할 수 있다.

```xml
<jsp-config>
  <jsp-property-group>
    <url-pattern>/view/*</url-pattern>
    <include-coda>/common/fotter.jsp</include-coda>
  </jsp-property-group>
</jsp-config>
```

## MVC 패턴

- Model : 로직과 데이터 영역 (개발의 주된 영역)
- View : 디자인 영역 (보여지는 부분)
- Controller : 로직과 디자인을 분리시키기 위한 영역

javaEE에서 MVC패턴을 반영한 개발방법을 Model 2라고 한다.

### fowarding을 이용해서 request객체를 다른곳에서도 이용하기

```java
RequestDispatcher dis = request.getRequestDispatcher("./result.jsp");
// 요청객체에 원하는 데이터를 포함시키자
request.setAttribute("data", msg);
dis.forward(request, response);
```

### MVC에서 컨트롤러의 역할

- 로직과 디자인을 분리시키려면 컨트롤러라는 중간 매개가 필요하다.
- 컨트롤러의 업무 영역
  - 요청을 받는다
  - 요청을 분석한다(요청의 유형이 다양하므로)
  - 알맞는 로직객체에게 일 시킨다
  - 결과가 있다면 결과를 저장한다 (결과를 view로 전달하기 위해서)
  - 결과를 보여준다 (view에서 결과 보여줄것을 요청한다)

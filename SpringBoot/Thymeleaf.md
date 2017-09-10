# Thymeleaf

## 소개

## 설정

application.properties

  ```text
  spring.thymeleaf.check-template-location=true
  spring.thymeleaf.prefix=/WEB-INF/templates/
  spring.thymeleaf.suffix=.html
  spring.thymeleaf.mode=HTML
  spring.thymeleaf.encoding=UTF-8
  spring.thymeleaf.content-type=text/html
  spring.thymeleaf.cache=false
  ```

Thymeleaf는 태그 정책이 엄격해서 오타나 표준에 맞지않는 구문이 있으면 칼같이 오류를 내뿜는다.
특히 닫는태그등의 HTML표준 관련 충돌이 잦은데 이를 완화하기 위해서 의존성패키지를 추가한다.

`spring.thymeleaf.mode=LEGACYHTML5` 적용시

gradle dependency 추가 `compile group: 'net.sourceforge.nekohtml', name: 'nekohtml', version: '1.9.22'`

## 기본 문법

### thymeleaf 사용 선언

`<html xmlns:th="http://www.thymeleaf.org">`

### 범위 선언

범위 내의 변수선언, 조건설정등을 할 수 있다.

`<th:block></th:block>`

### model의 attribute 출력

`<span th:text="${variable}"></span>`

`<span th:utext="${variable}"></span>`

### 반복문

```html
<tr data-th-each="data : ${list}">
  <td th:text="${data.userId}"></td>
  <td th:text="${data.name}"></td>
  <td th:text="${data.email}"></td>
</tr>
```

만약 숫자 범위를 지정하여 반복한다면

  ```html
  <!-- foo.start, foo.end 자리에 상수도 사용 가능함 -->
  <th:block th:each="page : ${#numbers.sequence(__${foo.start}__, __${foo.end}__)}">
    <span th:text="${page}"></span>
  </th:block>
  ```

### 조건문

삼항 연산자도 사용 가능하다

  ```html
  <tr th:class="${row.even}? (${row.first}? 'first' : 'even') : 'odd'">
    ...
  </tr>
  ```

if문

  ```html
  <li th:if="${session.user==null}"><a href="/login">로그인</a></li>
  <li th:if="${session.user!=null}"><a href="/logout">로그아웃</a></li>
  ```

switch문

  ```html
  <div th:switch="${user.role}">
    <p th:case="'admin'">User is an administrator</p>
    <p th:case="#{roles.manager}">User is a manager</p>
    <p th:case="*">User is some other thing</p>
  </div>
  ```

## 레이아웃 (Layout)

- th:insert : th:insert를 선언한 태그를 유지하고 내부에 fragment 전체를 가져옴

  ```html
  <div th:insert="footer :: copy"></div>
  <div>
    <footer>
      &copy; 2011 The Good Thymes Virtual Grocery
    </footer>
  </div>
  ```

- th:replace : th:replace를 선언한 태그 자체가 fragment로 바뀜

  ```html
  <div th:replace="footer :: copy"></div>

  <footer>
    &copy; 2011 The Good Thymes Virtual Grocery
  </footer>
  ```

- th:include : th:inclue를 선언한 태그는 유지되고 fragment에서 최상위태그 내부의 내용만 가져옴

  ```html
  <div th:include="footer :: copy"></div>

  <div>
    &copy; 2011 The Good Thymes Virtual Grocery
  </div>
  ```

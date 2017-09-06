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

## 레이아웃

- th:insert is the simplest: it will simply insert the specified fragment as the body of its host tag.
  - `<div th:insert="footer :: copy"></div>`
  ```html
  <div>
    <footer>
      &copy; 2011 The Good Thymes Virtual Grocery
    </footer>
  </div>
  ```

- th:replace actually replaces its host tag with the specified fragment.
  - `<div th:replace="footer :: copy"></div>`
  ```html
  <footer>
    &copy; 2011 The Good Thymes Virtual Grocery
  </footer>
  ```

- th:include is similar to th:insert, but instead of inserting the fragment it only inserts the contents of this fragment.
  - `<div th:include="footer :: copy"></div>`
  ```html
  <div>
    &copy; 2011 The Good Thymes Virtual Grocery
  </div>
  ```

# JavaScript

## 변수

컴퓨터에게 일거리(데이터)를 메모리에 제공하는 방법 =  변수선언

- 기본타입
  - Number
  - String
  - Boolean
  - undefined
  > undefined -> 변수를 선언만 하고 값을 할당하지 않음. 즉, 자료형이 결정되지 않은 상태이다.
  >> (선언하지 않은 변수도 콘솔이나 기타 메세지에는 undefined라고 뜨지만, undefined라는 값을 가지는 것은 아니다.)
  - null
  > null -> 변수를 선언하고, 'null'이라는 빈 값을 할당한 경우이다.
  >> (이 '빈 값'의 경우 자료형에 따라 여러가지가 있지만, null은 객체형 데이터-ex: array, object-의 빈 값을 의미한다. 문자열(string)의 경우 '', 숫자(number)의 경우 0이 빈값이고, 이들 빈값 모두는 if문에서 false로 형 변환된다.

- 참조타입
  - Object
  - Array
  - Function
  - Regulation Expression

## 객체 - 사용자(개발자) 정의 객체

- 객체 생성
  - Object()생성자 함수 이용
    ```javascript
    var foo = new Object();
    foo.name = "foo";
    foo.age = 30;
    ```

  - 객체 리터럴 방식 이용
    ```javascript
    var foo = {
      name : "foo",
      age : 30
    }
    ```

  - 생성자 함수 이용

## 연산자

데이터를 어떻게 처리할지를 결정하는 부호

- 산술연산자 : `+ - * / %(나머지) ++ --`
- 비교연산자 : `==, <=, != ...`
  > 비교연산자 ==는 자료형이 다르면 자동 형변환으로 자료형을 강제로 맞춰서 비교하는 비교연산자입니다.
- 삼항연산자 : `(condition) ? true : false;`
- Type연산자 : `===, typeof`
  > undefined와 null(object)은 자료형이 다르니 자바스크립트 엔진에서 알아서 통일해서 둘다 값이 없는거니까 true를 반환합니다. 이 경우 === 연산자(자료형까지 비교)를 사용하면 원하는 결과를 얻을 수 있습니다.
- !!연산자 : 피연산자의 값을 boolean 값으로 반환함

## 제어문

조건을 판단해 보자

반복문 (for, while)

## 함수 : 코드의 재사용

- 함수 생성
  - 함수선언문(function statement)
    - 반드시 함수명이 정의되어 있어야 함
    - 일종의 함수 리터럴을 이용해 생성하는 방식
  ```js
  function add(x, y) {
    return x + y;
  }
  ```

  - 함수표현식(function expression) :
    함수도 하나의 값처럼 취급되므로, 생성된 함수를 바로 변수에 할당한다
  ```js
  var add = function(x, y) {
    return x + y;
  };
  ```

  - Function() 생성자 함수 : 자주 사용되지는 않음
  ```js
  var add = new Function("x", "y", "return x + y");
  ```

  - 함수 호이스팅(function hoisting)
    함수 선언이 일어나기 전에서 함수를 사용할 수 있는 상황
    모든 함수를 함수 표현식으로 정의하면 호이스팅이 발생하지 않음

- 함수 객체
  - 함수는 값으로 취급된다.
  - 리터럴에 의해 생성
  - 변수나 배열의 요소, 객체의 프로퍼티에 할당 가능
  - 함수의 인자로 전달 가능
  - 함수의 리턴값으로 리턴 가능
  - 동적으로 프로퍼티를 생성 및 할당 가능
  - 함수의 기본 프로퍼티
    - length : 함수가 정상적으로 실행될 때 기대되는 인자의 수
    - prototype === `__proto__`
    - name : 함수의 이름
    - caller : 자신을 호출한 함수
    - arguments : 인자 값

- 함수 형태
  - 콜백함수 : 등록해놓은 함수가 개발자가 아니라 시스템에 의해서 이벤트나 특정 시점에 도달했을 때 호출 되는 경우
  ```js
  <script>
  window.onload = function() {
    alert("this is callback functinon");
  }
  </script>
  ```
  - 즉시 실행 함수 : 정의와 동시에 실행하는 함수
    - 즉시 실행함수는 다시 호출 할 수 없다.
    - 최초 한 번의 실행만을 필요로 하는 초기화 코드 등에 사용 할 수 있다.
  ```js
  (function(name) {
    console.log("this is the immediate function --> " + name)
  })("foo");
  ```
  - 내부 함수 : 함수 내부에서 정의된 함수, 클로저 생성, 부모 함수 코드에서 외부의 접근을 막고 독립적인 헬퍼 함수를 구현 하는 용도
  ```js
  function parent() {
    var a = 100;
    var b = 200;

    function child() {
      var b = 300;
      console.log(a); // 100
      console.log(b); // 300
    }
    child(); // 위의 100과 300이 출력 됨
  }
  child();  // error : child is not defined
  ```
  - 함수 자체를 리턴하는 함수
  ```js
  var self = function() {
    console.log("a");
    return function () {
      console.log("b");
    }
  }
  self = self(); // a
  self(); // b
  ```

- 함수 호출과 this

- 프로토타입과 프로토타입 체이닝

### 배열 : 감속도운동

```none
원점 ---------------------------------------------------------목적지
        100m
내위치 = 내위치 + (비율)*(목표지점 - 나의 현위치);
```

## AJAX (비동기 자바스크립트 & XML)

별도의 실행부가 background에서 요청을 시도하고 응답을 받아오는 방식

`XMLHttp onreadystatechange = function()`

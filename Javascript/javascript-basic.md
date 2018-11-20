# JavaScript

- [JavaScript](#javascript)
  - [변수](#%EB%B3%80%EC%88%98)
    - [기본타입](#%EA%B8%B0%EB%B3%B8%ED%83%80%EC%9E%85)
    - [참조타입](#%EC%B0%B8%EC%A1%B0%ED%83%80%EC%9E%85)
      - [객체 프로퍼티 읽기/쓰기/갱신/삭제](#%EA%B0%9D%EC%B2%B4-%ED%94%84%EB%A1%9C%ED%8D%BC%ED%8B%B0-%EC%9D%BD%EA%B8%B0%EC%93%B0%EA%B8%B0%EA%B0%B1%EC%8B%A0%EC%82%AD%EC%A0%9C)
  - [객체 - 사용자(개발자) 정의 객체](#%EA%B0%9D%EC%B2%B4---%EC%82%AC%EC%9A%A9%EC%9E%90%EA%B0%9C%EB%B0%9C%EC%9E%90-%EC%A0%95%EC%9D%98-%EA%B0%9D%EC%B2%B4)
    - [객체 생성](#%EA%B0%9D%EC%B2%B4-%EC%83%9D%EC%84%B1)
  - [연산자](#%EC%97%B0%EC%82%B0%EC%9E%90)
  - [제어문](#%EC%A0%9C%EC%96%B4%EB%AC%B8)
    - [조건문 (if, switch)](#%EC%A1%B0%EA%B1%B4%EB%AC%B8-if-switch)
      - [if문](#if%EB%AC%B8)
      - [switch문](#switch%EB%AC%B8)
    - [반복문 (for, while)](#%EB%B0%98%EB%B3%B5%EB%AC%B8-for-while)
      - [기본 for 반복문](#%EA%B8%B0%EB%B3%B8-for-%EB%B0%98%EB%B3%B5%EB%AC%B8)
      - [forEach 반복문](#foreach-%EB%B0%98%EB%B3%B5%EB%AC%B8)
      - [for-in 반복문](#for-in-%EB%B0%98%EB%B3%B5%EB%AC%B8)
      - [for-of 반복문](#for-of-%EB%B0%98%EB%B3%B5%EB%AC%B8)
      - [do-while문](#do-while%EB%AC%B8)
      - [while문](#while%EB%AC%B8)
  - [함수 : 코드의 재사용](#%ED%95%A8%EC%88%98--%EC%BD%94%EB%93%9C%EC%9D%98-%EC%9E%AC%EC%82%AC%EC%9A%A9)
    - [함수 생성](#%ED%95%A8%EC%88%98-%EC%83%9D%EC%84%B1)
    - [함수 객체](#%ED%95%A8%EC%88%98-%EA%B0%9D%EC%B2%B4)
    - [함수 형태](#%ED%95%A8%EC%88%98-%ED%98%95%ED%83%9C)
    - [함수 호출과 this](#%ED%95%A8%EC%88%98-%ED%98%B8%EC%B6%9C%EA%B3%BC-this)
      - [arguments 객체](#arguments-%EA%B0%9D%EC%B2%B4)
      - [호출 패턴과 this 바인딩](#%ED%98%B8%EC%B6%9C-%ED%8C%A8%ED%84%B4%EA%B3%BC-this-%EB%B0%94%EC%9D%B8%EB%94%A9)
        - [객체의 메소드 호출시 this 바인딩](#%EA%B0%9D%EC%B2%B4%EC%9D%98-%EB%A9%94%EC%86%8C%EB%93%9C-%ED%98%B8%EC%B6%9C%EC%8B%9C-this-%EB%B0%94%EC%9D%B8%EB%94%A9)
        - [함수 호출시 this 바인딩](#%ED%95%A8%EC%88%98-%ED%98%B8%EC%B6%9C%EC%8B%9C-this-%EB%B0%94%EC%9D%B8%EB%94%A9)
        - [생성자 함수를 호출할시 this 바인딩](#%EC%83%9D%EC%84%B1%EC%9E%90-%ED%95%A8%EC%88%98%EB%A5%BC-%ED%98%B8%EC%B6%9C%ED%95%A0%EC%8B%9C-this-%EB%B0%94%EC%9D%B8%EB%94%A9)
          - [생성자 함수가 동작하는 방식](#%EC%83%9D%EC%84%B1%EC%9E%90-%ED%95%A8%EC%88%98%EA%B0%80-%EB%8F%99%EC%9E%91%ED%95%98%EB%8A%94-%EB%B0%A9%EC%8B%9D)
          - [객체 리터럴 방식과 생성자 함수를 통한 객체생성 차이](#%EA%B0%9D%EC%B2%B4-%EB%A6%AC%ED%84%B0%EB%9F%B4-%EB%B0%A9%EC%8B%9D%EA%B3%BC-%EC%83%9D%EC%84%B1%EC%9E%90-%ED%95%A8%EC%88%98%EB%A5%BC-%ED%86%B5%ED%95%9C-%EA%B0%9D%EC%B2%B4%EC%83%9D%EC%84%B1-%EC%B0%A8%EC%9D%B4)
          - [생성자 함수를 new 키워드 없이 호출할 경우](#%EC%83%9D%EC%84%B1%EC%9E%90-%ED%95%A8%EC%88%98%EB%A5%BC-new-%ED%82%A4%EC%9B%8C%EB%93%9C-%EC%97%86%EC%9D%B4-%ED%98%B8%EC%B6%9C%ED%95%A0-%EA%B2%BD%EC%9A%B0)
        - [call과 apply 메소드를 이용한 명시적인 this 바인딩](#call%EA%B3%BC-apply-%EB%A9%94%EC%86%8C%EB%93%9C%EB%A5%BC-%EC%9D%B4%EC%9A%A9%ED%95%9C-%EB%AA%85%EC%8B%9C%EC%A0%81%EC%9D%B8-this-%EB%B0%94%EC%9D%B8%EB%94%A9)
      - [함수 리턴](#%ED%95%A8%EC%88%98-%EB%A6%AC%ED%84%B4)
    - [프로토타입 체이닝](#%ED%94%84%EB%A1%9C%ED%86%A0%ED%83%80%EC%9E%85-%EC%B2%B4%EC%9D%B4%EB%8B%9D)
      - [프로토타입의 두 가지 의미](#%ED%94%84%EB%A1%9C%ED%86%A0%ED%83%80%EC%9E%85%EC%9D%98-%EB%91%90-%EA%B0%80%EC%A7%80-%EC%9D%98%EB%AF%B8)
      - [객체 리터럴 방식으로 생성된 객체의 프로토타입 체이닝](#%EA%B0%9D%EC%B2%B4-%EB%A6%AC%ED%84%B0%EB%9F%B4-%EB%B0%A9%EC%8B%9D%EC%9C%BC%EB%A1%9C-%EC%83%9D%EC%84%B1%EB%90%9C-%EA%B0%9D%EC%B2%B4%EC%9D%98-%ED%94%84%EB%A1%9C%ED%86%A0%ED%83%80%EC%9E%85-%EC%B2%B4%EC%9D%B4%EB%8B%9D)
      - [생성자 함수로 생성된 객체의 프로토타입 체이닝](#%EC%83%9D%EC%84%B1%EC%9E%90-%ED%95%A8%EC%88%98%EB%A1%9C-%EC%83%9D%EC%84%B1%EB%90%9C-%EA%B0%9D%EC%B2%B4%EC%9D%98-%ED%94%84%EB%A1%9C%ED%86%A0%ED%83%80%EC%9E%85-%EC%B2%B4%EC%9D%B4%EB%8B%9D)
      - [프로토타입 체이닝의 종점](#%ED%94%84%EB%A1%9C%ED%86%A0%ED%83%80%EC%9E%85-%EC%B2%B4%EC%9D%B4%EB%8B%9D%EC%9D%98-%EC%A2%85%EC%A0%90)
      - [기본 데이터 타입 확장](#%EA%B8%B0%EB%B3%B8-%EB%8D%B0%EC%9D%B4%ED%84%B0-%ED%83%80%EC%9E%85-%ED%99%95%EC%9E%A5)
      - [프로토타입도 자바스크립트 객체다](#%ED%94%84%EB%A1%9C%ED%86%A0%ED%83%80%EC%9E%85%EB%8F%84-%EC%9E%90%EB%B0%94%EC%8A%A4%ED%81%AC%EB%A6%BD%ED%8A%B8-%EA%B0%9D%EC%B2%B4%EB%8B%A4)
      - [프로토타입 메소드와 this 바인딩](#%ED%94%84%EB%A1%9C%ED%86%A0%ED%83%80%EC%9E%85-%EB%A9%94%EC%86%8C%EB%93%9C%EC%99%80-this-%EB%B0%94%EC%9D%B8%EB%94%A9)
      - [디폴트 프로토타입은 다른 객체로 변경 가능하다](#%EB%94%94%ED%8F%B4%ED%8A%B8-%ED%94%84%EB%A1%9C%ED%86%A0%ED%83%80%EC%9E%85%EC%9D%80-%EB%8B%A4%EB%A5%B8-%EA%B0%9D%EC%B2%B4%EB%A1%9C-%EB%B3%80%EA%B2%BD-%EA%B0%80%EB%8A%A5%ED%95%98%EB%8B%A4)
  - [실행 컨텍스트와 클로저](#%EC%8B%A4%ED%96%89-%EC%BB%A8%ED%85%8D%EC%8A%A4%ED%8A%B8%EC%99%80-%ED%81%B4%EB%A1%9C%EC%A0%80)
    - [실행 컨텍스트 개념](#%EC%8B%A4%ED%96%89-%EC%BB%A8%ED%85%8D%EC%8A%A4%ED%8A%B8-%EA%B0%9C%EB%85%90)
    - [실행 컨텍스트 생성 과정](#%EC%8B%A4%ED%96%89-%EC%BB%A8%ED%85%8D%EC%8A%A4%ED%8A%B8-%EC%83%9D%EC%84%B1-%EA%B3%BC%EC%A0%95)
    - [스코프 체인](#%EC%8A%A4%EC%BD%94%ED%94%84-%EC%B2%B4%EC%9D%B8)
      - [전역 실행 컨텍스트의 스코프 체인](#%EC%A0%84%EC%97%AD-%EC%8B%A4%ED%96%89-%EC%BB%A8%ED%85%8D%EC%8A%A4%ED%8A%B8%EC%9D%98-%EC%8A%A4%EC%BD%94%ED%94%84-%EC%B2%B4%EC%9D%B8)
      - [함수를 호출한 경우 생성되는 실행 컨텍스트의 스코프 체인](#%ED%95%A8%EC%88%98%EB%A5%BC-%ED%98%B8%EC%B6%9C%ED%95%9C-%EA%B2%BD%EC%9A%B0-%EC%83%9D%EC%84%B1%EB%90%98%EB%8A%94-%EC%8B%A4%ED%96%89-%EC%BB%A8%ED%85%8D%EC%8A%A4%ED%8A%B8%EC%9D%98-%EC%8A%A4%EC%BD%94%ED%94%84-%EC%B2%B4%EC%9D%B8)
      - [스코프 체인을 수정하는 키워드 with](#%EC%8A%A4%EC%BD%94%ED%94%84-%EC%B2%B4%EC%9D%B8%EC%9D%84-%EC%88%98%EC%A0%95%ED%95%98%EB%8A%94-%ED%82%A4%EC%9B%8C%EB%93%9C-with)
    - [클로저](#%ED%81%B4%EB%A1%9C%EC%A0%80)
      - [클로저의 개념](#%ED%81%B4%EB%A1%9C%EC%A0%80%EC%9D%98-%EA%B0%9C%EB%85%90)
      - [클로저의 활용](#%ED%81%B4%EB%A1%9C%EC%A0%80%EC%9D%98-%ED%99%9C%EC%9A%A9)
      - [클로저 활용시 주의사항](#%ED%81%B4%EB%A1%9C%EC%A0%80-%ED%99%9C%EC%9A%A9%EC%8B%9C-%EC%A3%BC%EC%9D%98%EC%82%AC%ED%95%AD)
  - [객체지향 프로그래밍](#%EA%B0%9D%EC%B2%B4%EC%A7%80%ED%96%A5-%ED%94%84%EB%A1%9C%EA%B7%B8%EB%9E%98%EB%B0%8D)
    - [클래스, 생성자, 메소드](#%ED%81%B4%EB%9E%98%EC%8A%A4-%EC%83%9D%EC%84%B1%EC%9E%90-%EB%A9%94%EC%86%8C%EB%93%9C)
    - [상속](#%EC%83%81%EC%86%8D)
    - [캡슐화](#%EC%BA%A1%EC%8A%90%ED%99%94)
    - [응용: 클래스의 기능을 가진 subClass 함수](#%EC%9D%91%EC%9A%A9-%ED%81%B4%EB%9E%98%EC%8A%A4%EC%9D%98-%EA%B8%B0%EB%8A%A5%EC%9D%84-%EA%B0%80%EC%A7%84-subclass-%ED%95%A8%EC%88%98)
      - [자식 클래스 생성 및 상속](#%EC%9E%90%EC%8B%9D-%ED%81%B4%EB%9E%98%EC%8A%A4-%EC%83%9D%EC%84%B1-%EB%B0%8F-%EC%83%81%EC%86%8D)
      - [subClass 활용](#subclass-%ED%99%9C%EC%9A%A9)
  - [함수형 프로그래밍](#%ED%95%A8%EC%88%98%ED%98%95-%ED%94%84%EB%A1%9C%EA%B7%B8%EB%9E%98%EB%B0%8D)
    - [함수형 프로그래밍의 개념](#%ED%95%A8%EC%88%98%ED%98%95-%ED%94%84%EB%A1%9C%EA%B7%B8%EB%9E%98%EB%B0%8D%EC%9D%98-%EA%B0%9C%EB%85%90)
    - [자바스크립트에서 함수형 프로그래밍](#%EC%9E%90%EB%B0%94%EC%8A%A4%ED%81%AC%EB%A6%BD%ED%8A%B8%EC%97%90%EC%84%9C-%ED%95%A8%EC%88%98%ED%98%95-%ED%94%84%EB%A1%9C%EA%B7%B8%EB%9E%98%EB%B0%8D)
    - [함수형 프로그래밍을 활용한 주요 함수](#%ED%95%A8%EC%88%98%ED%98%95-%ED%94%84%EB%A1%9C%EA%B7%B8%EB%9E%98%EB%B0%8D%EC%9D%84-%ED%99%9C%EC%9A%A9%ED%95%9C-%EC%A3%BC%EC%9A%94-%ED%95%A8%EC%88%98)
      - [apply](#apply)
      - [커링 (currying)](#%EC%BB%A4%EB%A7%81-currying)
      - [map](#map)
      - [reduce](#reduce)

## 변수

### 기본타입

- Number: 모든 숫자를 64비트 부동 소수점 형태로 저장
- String: character 자료형은 별도로 없음, 또한 여러 다른 언어처럼 문자열은 immutable임
- Boolean
- undefined(타입이자 값임)
  - undefined -> 변수를 선언만 하고 값을 할당하지 않음. 즉, 자료형이 결정되지 않은 상태이다.
  - (선언하지 않은 변수도 콘솔이나 기타 메세지에는 undefined라고 뜨지만, undefined라는 값을 가지는 것은 아니다.)
- null(null 값)
  - null -> 변수를 선언하고, 'null'이라는 빈 값을 할당한 경우이다.
  - 이 '빈 값'의 경우 자료형에 따라 여러가지가 있지만, null은 객체형 데이터 (array, object-의 빈 값을 의미한다)
  - 문자열(string)의 경우 '', 숫자(number)의 경우 0이 빈값이고, 이들 빈값 모두는 if문에서 false로 형 변환된다.

### 참조타입

자바스크립트에서는 5개의 기본타입을 제외한 모든 값은 객체이다

- Object
- Array
- Function
- Regulation Expression
- ...

#### 객체 프로퍼티 읽기/쓰기/갱신/삭제

프로퍼티 접근은 대괄호(`[]`) 또는 마침표(`.`) 두가지 방법을 사용한다

```js
var foo = {
  name: 'foo'
};

console.log(foo.name);
console.log(foo['name']);
```

프로퍼티 삭제는 `delete` 키워드를 이용한다

```js
delete foo.name;
```

## 객체 - 사용자(개발자) 정의 객체

### 객체 생성

- Object()생성자 함수 이용

  ```js
  var foo = new Object();
  foo.name = "foo";
  foo.age = 30;
  ```

- 객체 리터럴 방식 이용

  ```js
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
  - 비교(동등)연산자 == 는 자료형이 다르면 자동 형변환으로 자료형을 강제로 맞춰서 비교하는 비교연산자
- 삼항연산자 : `(condition) ? true : false;`
- Type연산자 : `===`(일치 연산자), `typeof`(타입 연산자)
  - undefined와 null(object)은 을 비교연산자로 비교하면 true를 반환함
  - 이 경우 === 연산자(자료형까지 비교)를 사용하면 원하는 결과를 얻을 수 있음
- `!!`연산자 : 피연산자의 값을 boolean 값으로 반환함

## 제어문

### 조건문 (if, switch)

#### if문

```js
if (조건1) {
  명령문1
} else if (조건2) {
  명령문2
  ...
} else {
  기본명령
}
```

#### switch문

```js
switch (expression) {
  case value1:
    //Statements executed when the result of expression matches value1
    [break;]
  case value2:
    //Statements executed when the result of expression matches value2
    [break;]
  ...
  case valueN:
    //Statements executed when the result of expression matches valueN
    [break;]
  default:
    //Statements executed when none of the values match the value of the expression
    [break;]
}
```

### 반복문 (for, while)

#### 기본 for 반복문

```js
for (var i = 0; i < 10; i ++) {
  console.log(i);
}
```

#### forEach 반복문

forEach 반복문은 Array 객체에서만 사용가능하다.(ES6의 Map, Set에서도 가능하다)
배열내의 각 요소에 대해 콜백함수가 적용된다.

```js
var items = ['item1', 'item2', 'item3'];

items.forEach(function(item) {
  console.log(item);
});
```

#### for-in 반복문

for in 반복문은 배열이나 객체의 속성에 대상으로 반복한다
객체의 key 값에 접근할 수 있지만, value에는 접근할 수 없다. (`객체[key]`로 접근)

```js
var obj = { a: 1, b: 2, c: 3 };

for (var key in obj) {
  console.log(obj[key]);
}
```

#### for-of 반복문

for of 반복문은 ES6에서 추가되었다.
`[Symbol.iterator]` 속성과 `iterable` 인터페이스를 구현하고 있어야 한다

```js
var iterable = [10, 20, 30];

for (var value of iterable) {
  console.log(value); // 10, 20, 30
}
```

#### do-while문

```js
do
  // sth to do
while (조건문);
```

#### while문

```js
while (조건문) {
  // sth to do
}
```

## 함수 : 코드의 재사용

### 함수 생성

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

### 함수 객체

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

### 함수 형태

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

### 함수 호출과 this

#### arguments 객체

자바스크립트는 함수 호출시 인자를 생략하더라도 에러가 발생하지 않는다.
생략한 인자값으로는 `undefined`가 할당된다.

반대로 정의된 인자 개수보다 많이 넘기면 초과된 인자값은 무시된다.

자바스크립트에서 함수를 호출할 때 인수들과 함께 암묵적으로 arguments 객체가 함수 내부로 전달된다.

arguments객체는 함수를 호출할 때 넘긴 인자들이 유사 배열형태로 전달된다.

```js
function add(a, b) {
  console.dir(arguments);
  return a + b;
}
```

arguments 객체는 세 부분으로 구성되어 있다

- 함수 호출시 넘겨진 인자 (배열): 인덱스와 값
- length 프로퍼티: 호출할 때 넘겨진 인자의 개수
- callee 프로퍼티: 현재 실행 중인 함수의 참조값

#### 호출 패턴과 this 바인딩

자바스크립트에서는 함수를 호출할 때 인자값에 더해 arguments 객체 및 this 인자가 암묵적으로 전달된다.

특히 this는 함수 호출방식에 따라 다른객체를 참조하게 되므로 유의해야 한다.

##### 객체의 메소드 호출시 this 바인딩

객체의 프로퍼티가 함수인 경우 이를 메소드라 부르고,
메소드 내부에서 사용된 this는 해당 메소드를 호출한 객체로 바인딩 된다.

##### 함수 호출시 this 바인딩

자바스크립트에서 함수를 호출하면 해당 함수 내부코드에서 사용된 this는 전역객체에 바인딩된다.

브라우저에서 자바스크립트를 실행하는 경우 전역객체는 `window` 객체가 되고,
Node.js 환경에서 실행하는 경우 `global` 객체가 된다.

함수 호출의 this바인딩 특성은 내부함수(inner function)를 호출했을 경우에도 적용되므로 사용에 주의해야 한다.

```js
var value = 100;

var myObject = {
  value : 1,
  func1: function() {
    this.value += 1;
    console.log('func1() called. this.value : ' + this.value);

    // 내부함수
    func2 = function() {
      this.value += 1;
      console.log('func2() called. this.value : ' + this.value);

      // 내부함수의 내부함수
      func3 = function() {
        this.value += 1;
        console.log('func3() called. this.value : ' + this.value);
      }

      func3();
    }

    func2();
  }
};

myObject.func1();

// Output:
// 2
// 101
// 102
```

내부함수의 this 바인딩을 의도한대로 사용하려면 부모함수의 this를 내부함수가 접근가능한 다른 변수에 저장하는 방법을 사용한다.

```js
var value = 100;

var myObject = {
  value : 1,
  func1: function() {
    var that = this;

    this.value += 1;
    console.log('func1() called. this.value : ' + this.value);딩

    // 내부함수
    func2 = function() {
      that.value += 1;
      console.log('func2() called. this.value : ' + that.value);

      // 내부함수의 내부함수
      func3 = function() {
        that.value += 1;
        console.log('func3() called. this.value : ' + that.value);
      }

      func3();
    }

    func2();
  }
};

myObject.func1();

// Output:
// 2
// 3
// 4
```

또한, call과 apply 메소드를 통해서 this 바인딩을 명시적으로 할 수 있다

##### 생성자 함수를 호출할시 this 바인딩

자바스크립트 객체를 생성하는 방법은 크게 객체 리터럴 방식이나 생성자 함수를 이용하는 두 가지 방식이 있다.

자바스크립트에서는 기존 함수에 `new` 연산자를 붙여서 호출하면 해당 함수는 생성자 함수로 동작한다.

함수가 원하지 않게 생성자 함수로 동작할 수 있으므로,
자바스크립트 스타일 가이드에서는 생성자 함수로 사용되는 함수는 첫 글자를 대문자로 쓰기를 권한다.

생성자 함수 코드 내부에서의 this는 메소드의 함수 호출방식의 this 바인딩과는 다르게 작동한다.

###### 생성자 함수가 동작하는 방식

new 연산자로 함수를 생성자로 호출하면 다음 순서로 동작한다

- 빈 객체 생성 및 this 바인딩
  - 생성자 함수 코드가 실행되기 전 빈 객체가 생성되고, 이 객체는 this로 바인딩된다
  - 엄밀히 말하면 빈 객체는 아니고, 자신을 생성한 생성자 함수의 prototype 프로퍼티가 가리키는 객체를 자신의 프로토타입으로 설정한다

- this를 통한 프로퍼티 생성
  - 함수 코드 내부에서 this를 사용해서 생성한 객체에 동적으로 프로퍼티나 메소드를 생성한다

- 생성된 객체 리턴
  - 특별히 리턴문이 없으면 this로 바인딩된 새로 생성한 객체가 반환된다 (생성자 함수가아닌 경우 기본 반환값은 undefined)
  - 리턴값이 this가 아닌경우 해당 값이 반환된다

###### 객체 리터럴 방식과 생성자 함수를 통한 객체생성 차이

객체 리터럴 방식으로 생성된 객체는 같은 형태의 객체를 재생성 할 수 없다.

또한 두 방식은 생성된 객체의 프로토타입이 다르다.

객체 리터럴 방식의 경우는 프로토타입이 `Object.prototype`이고,
생성자 함수 방식의 경우 `MyObject.prototype`으로 서로 다르다.

###### 생성자 함수를 new 키워드 없이 호출할 경우

일반함수와 생성자함수는 this 바인딩이 다르므로, 의도와 다르게 호출된 경우 오류발생 가능성이 있다.

이런겨우를 피하기 위해 객체를 생성하는 별도의 코드패턴을 사용하기도 한다.

```js
function A(arg) {
  if (!(this instanceof A)) return new A(arg);
  this.value = arg ? arg : 0;
}

var a = new A(100);
var b = A(10);

console.log(a.value); // 100
console.log(b.value); // 10
console.log(global.value); // undefined
```

##### call과 apply 메소드를 이용한 명시적인 this 바인딩

자바스크립트에서는 상황에 따라 this가 자동바인딩 되지만,
this를 특정 객체에 명시적으로 바인딩시키는 방법도 제공한다.

`Function.prototype` 객체의 메소드인 `apply()`와 `call()` 메소드를 사용하면 된다.

`call`과 `apply`는 기능이 같고 넘겨받는 인자의 형식만 다르다.

`function.apply(thisArg, argArray)`

- thisArg는 apply() 메소드를 호출한 함수 내부에서 사용한 this에 바인딩할 객체이다
- argArray 인자는 함수를 호출할 때 넘길 인자들의 배열을 가리킨다

```js
function Person(name, age, gender) {
  this.name = name;
  this.age = age;
  this.gender = gender;
}

var foo = {};

Person.apply(foo, ['foo', 30, 'man']); // this를 foo 객체에 바인딩
console.dir(foo);
```

위의 `apply()` 메소드를 `call()` 메소드로 바꾸면 다음과 같다

`Person.call(foo, 'foo', 30, 'man');`

arguments 객체는 실제 배열이 아니지만 apply() 메소드를 사용하여 배열로 변환하여 사용할 수 있다

```js
function myFunction() {
  console.dir(arguments); // 프로토타입이 Object.prototype
  var args = Array.prototype.slice.apply(arguments);
  console.dir(args); // 프로토타입이 Array.prototype
}

myFunction(1, 2, 3);
```

#### 함수 리턴

자바스크립트 함수는 항상 리턴값을 반환한다. return 문이 명시적이지 않아도 경우에 따라 기본값이 반환된다.

- 일반 함수나 메소드는 리턴값이 없으면 undefined 값 반환
- 생성자 함수에서 리턴값을 지정하지 않을경우 생성된 객체가 반환

### 프로토타입 체이닝

#### 프로토타입의 두 가지 의미

자바스크립트는 기존 객체지향 프로그래밍 언어와는 다른 프로토타입 기반의 객체지향 프로그래밍을 지원한다.

자바스크립트에서는 객체 리터럴이나 생성자 함수로 객체를 생성하는데,
이렇게 생성된 객체의 부모 객체가 바로 프로토타입 객체이다.

자바스크립트의 모든 객체는 자신의 부모인 프로토타입 객체를 가리키는 참조 링크 형태의 숨겨진 프로퍼티가 있다.
이런 링크를 암묵적 프로토타입 링크라고 부르며 이러한 링크는 모든 객체의 `[[Prototype]]` 프로퍼티에 저장된다.

앞에서 살펴본 함수객체의 prototype 프로퍼티와 숨은 프로퍼티인 `[[Prototype]]` 링크를 구분해야 한다.

자바스크립트에서 모든 객체는 (자신을 생성한 **생성자 함수의 prototype 프로퍼티**가 가리키는) 프로토타입 객체를
`[[Prototype]]` 링크로 연결하여 자신의 부모 객체로 설정한다.

`__proto__` 프로퍼티는 모든 객체에 존재하는 숨겨진 프로퍼티로 객체 자신의 프로토타입 객체를 가리키는 참조링크 정보이다.

ECMAScript에서는 이것을 `[[Prototype]]` 프로퍼티로 정하였지만,
크롬이나 파이어폭스 같은 브라우저에서는 `__proto__` 프로퍼티로 명시적으로 제공한다. 즉, 이 두개는 같은것으로 보면된다.

#### 객체 리터럴 방식으로 생성된 객체의 프로토타입 체이닝

자바스크립트에서 객체는 자신의 프로퍼티뿐만 아니라, 부모역할을 하는 프로토타입 객체의 프로퍼티도 자신의 것처럼 접근 가능하다.

이것을 가능하게 하는것이 프로토타입 체이닝이다.

```js
var myObject = {
  name: 'foo',
  sayName: function () {
    console.log('MyName is ' + this.name);
  }
};

myObject.sayName();
console.log(myObject.hasOwnProperty('name'));
console.log(myObject.hasOwnProperty('nickName'));
myObject.sayNickName();

// Output:
// My Name is foo
// true
// false
// Uncaught TypeError: Object #<Object> has no method 'sayNickName'
```

객체 리터럴로 생성한 객체는 `Object()`라는 내장 생성자 함수로 생성된 것이다.
따라서 myObject는 Object() 함수의 prototype 프로퍼티가 가리키는 `Object.prototype` 객체를 자신의 프로토타입객체로 연결한다.

프로토타입 체이닝은 특정 객체의 프로퍼티나 메소드에 접근하려고 할 때,
해당 객체에 해당 프로퍼티나 메소드가 없다면 `[[Prototype]]` 링크를 따라 자신의 부모 역할을 하는 프로토타입 객체의 프로퍼티를 차례로 검색하는 것이다.

#### 생성자 함수로 생성된 객체의 프로토타입 체이닝

생성자 함수로 생성된 객체의 프로토타입은 `MyObject.prototype` 객체이고, 이 객체는 `Object.prototype` 객체를 프로토타입 객체로 가진다.

#### 프로토타입 체이닝의 종점

`Object.prototype` 객체는 프로토타입 체이닝의 종점이다.

따라서 어떤 방식으로 객체를 생성하여도 모든 자바스크립트 객체는 프로토타입 체이닝으로
`Object.prototype` 객체의 프로퍼티와 메소드에 접근가능하고 공유 가능하다.

#### 기본 데이터 타입 확장

`Object.prototype`에 정의된 메소드들은 모든 자바스크립트 객체의 표준 메소드라고 볼 수 있다.

이와 같은 방식으로 자바스크립트의 숫자, 문자열, 배열 등에서 사용되는 표준 메소드의 경우
이들의 프로토타입인 `Number.prototype`, `String.prototype`, `Array.prototype` 등에 정의되어 있다.

물론 이들 기본 내장 프로토타입 객체 또한 `Object.prototype`을 자신의 프로토타입으로 가지며 프로토타입 체이닝으로 연결된다.

자바스크립트에서는 표준 built-in 프로토타입 객체에 사용자 정의 메소드 추가를 허용한다.

```js
String.prototype.testMethod = function() {
  console.log('This is the String.prototype.testMethod()');
};

var str = 'this is test';
str.testMethod(); // This is the String.prototype.testMehtod()

console.dir(String.prototype) // testMethod가 추가되어 있음
```

#### 프로토타입도 자바스크립트 객체다

함수가 생성될 때, 자신의 prototype 프로퍼티에 연결되는 프로토타입 객체는 constructor 프로퍼티만을 가진 객체다.

#### 프로토타입 메소드와 this 바인딩

프로토타입 객체는 메소드를 가질수 있다 (프로토타입 메소드)

만약 프로토타입 메소드 내부에서 this를 사용하면 this는 메소드를 호출한 객체에 바인딩된다.

1. 프로토타입으로 연결된 객체가 생성됨
2. 메소드 호출
3. 해당 객체에 없으므로 체이닝으로 프로토타입에서 검색
4. 프로토타입 메소드를 호출한 객체(생성된)에 this가 바인딩 됨

#### 디폴트 프로토타입은 다른 객체로 변경 가능하다

자바스크립트에서는 함수를 생성할 때 해당 함수와 연결되는 디폴트 프로토타입 객체를 다른 일반객체로 변경할 수 있다.
이러한 특성을 활용하여 객체지향의 상속을 구현할 수 있다.

생성자 함수의 프로토타입 객체가 변경되면 변경된 시점 이후에 생성된 객체들은
변경된 프로토타입 객체로 `[[Prototype]]` 링크를 연결한다.

생성자 함수의 프로토타입이 변경되기 이전에 생성된 객체들은 기존 프로토타입 객체로 `[[Prototype]]` 링크를 유지한다.

## 실행 컨텍스트와 클로저

### 실행 컨텍스트 개념

언어의 call stack은 함수를 호출할 때 해당 함수의 호출정보가 쌓여있는 스택을 의미한다.

실행 컨텍스트는 **자바스크립트 코드 블록이 실행되는 환경**이라 할 수 있다.

가장 바깥쪽(아래쪽)에 위치하는 컨텍스트를 전역 컨텍스트라고 한다(가장 먼저 실행되는 컨텍스트).

전역 실행 컨텍스트는 일반 실행 컨텍스트와는 달리, arguments 객체가 없으며 전역 객체 하나만을 포함하는 스코프 체인이 있다.

ECMAScript에서는 실행 컨텍스트가 형성되는 경우를 세 가지로 규정하고 있는데,
(1)전역코드, (2)`eval()` 함수로 실행되는 코드, (3)함수안의 코드를 실행할 경우이다.

대부분의 경우 함수로 실행 컨텍스트를 만들고 스택에 순차적으로 쌓인다.

### 실행 컨텍스트 생성 과정

- 활성객체: 실행 컨텍스트에서 필요한 여러가지 정보를 담을 객체

- arguments 프로퍼티: 활성 객체에서 arguments 객체를 참조값

- 스코프 정보: 현재 컨텍스트의 유효 범위
  - 스코프 정보는 현재 실행중인 실행 컨텍스트 안에서 연결 리스트와 유사한 형식으로 만들어짐
  - 현재 컨텍스트에서 특정 변수에 접근해야 할 경우 이 리스트를 활용한다
  - 리스트를 **스코프 체인**이라고 하는데 `[[scope]]` 프로퍼티로 참조된다

- 현재 실행 컨텍스트 내부에서 사용되는 지역변수의 생성
  - 앞서 생성된 활성객체가 변수객체로 사용됨
  - 변수 객체 안에서 호출된 함수인자는 각각의 프로퍼티가 만들어지고 그 값이 할당됨 (값이 넘겨지지 않았다면 `undefined`)
  - 변수나 내부 함수는 메모리에 인스턴스화가 이루어지나, 초기화는 변수나 함수에 해당하는 표현식이 실행되기 전까지는 이루어지지 않는다(초기화 이전까지 `undefined`)

- `this` 키워드를 사용하는 값이 할당된다. this가 참조하는 객체가 없다면 전역객체를 참조한다.

### 스코프 체인

자바스크립트도 다른 언어와 마찬가지로 스코프, 유효범위가 있다.

자바스크립트에서는 함수내의 블록은( `for() {...}`, `if() {}` ...) 유효범위가 없다.
오직 **함수만이 유효범위의 한 단위**가 된다.

이 유효 범위를 나타내는 스코프가 `[[scope]]` 프로퍼티로 각 함수 객체 내에서 연결 리스트 형식으로 관리되는데,
이를 스코프 체인이라고 한다.

#### 전역 실행 컨텍스트의 스코프 체인

```js
var num = 1;
console.log(num);
```

위의 코드는 전역 코드만 존재하므로 전역 컨텍스트와 변수객체가 만들어진다.
이 변수 객체의 스코프체인은 자기 자신만을 가진다.(변수 객체의 `[[scope]]`는 변수 객체 자신을 가리킨다)

#### 함수를 호출한 경우 생성되는 실행 컨텍스트의 스코프 체인

전역 실행 컨텍스트에서 함수가 호출된 경우를 살펴보자(`depth=1`)

`스코프 체인 = 현재 실행 컨텍스트의 변수 객체 + 상위 컨텍스트의 스코프 체인`

식별자 인식은 스코프 체인의 첫번째 변수 객체부터 시작한다.

- 식별자와 대응되는 이름을 가진 프로퍼티가 있는지 확인
- 함수 호출시 스코프 체인 가장 앞의 객체가 변수 객체이므로 이 객체의 공식인자, 내부함수, 지역변수에 대응되는지 확인
- 첫 번째 객체에 대응되는 프로퍼티를 발견하지 못하면 다음 객체로 이동
- 대응되는 이름의 프로퍼티를 찾을 때 까지 계속한다
- this는 식별자가 아닌 키워드이므로, 스코프 체인의 참조 없이 접근할 수 있다

#### 스코프 체인을 수정하는 키워드 with

`with` 구문은 객체인 표현식을 실행하면, 객체가 현재 실행 컨텍스트의 스코프 체인에 추가된다.

```js
var y = { x: 5 };

function withExamFunc() {
  var x = 10;
  var z;

  with(y) {
    z = function() {
      console.log(x); // 5
    }
  }
  z();
}

withExamFunc();
```

`withExamFunc()` 함수가 호출되면 실행 컨텍스트의 스코프 체인은 전역변수 객체와 현재 실행 컨텍스트의 변수 객체를 포함하는 범위이다.

여기에 `with` 구문의 범위 내에서 스코프 체인의 맨 앞에 전역 변수 `y`가 추가된다.

즉, z 실행 컨텍스트의 스코프 체인은 다음과 같아진다 = `[y객체, z 변수객체, withExamFunc 변수객체, 전역객체]`

### 클로저

#### 클로저의 개념

```js
function outerFunc() {
  var x = 10;
  var innerFunc = function() {
    console.log(x);
  }
  return innerFunc;
}

var inner = outerFunc();
inner(); // 10
```

자바스크립트의 함수는 일급 객체로 취급된다.

여기서 최종 반환되는 함수가 외부 함수의 지역변수에 접근하는데,
이 지역변수에 접근하려면 함수가 종료되어 외부 함수의 컨텍스트가 반환되더라도
변수 객체는 반환되는 내부 함수의 스코프 체인에 그대로 남아있어야 한다.

이것이 바로 클로저이다.

다시 설명하면 이미 생명 주기가 끝난 외부함수의 변수를 참조하는 함수를 클로저라 할 수 있다.

앞에서는 outerFunc에서 선언된 x를 참조하는 innerFunc가 클로저가 된다.
그리고 클로저로 참조되는 외부변수인 outerFunc의 x와 같은 변수를 자유변수(free variable)라 한다.

closure라는 명칭은 함수가 자유 변수에 대해 닫혀있다라는 의미이다.

#### 클로저의 활용

클로저는 성능적인 면과 자원적인 면에서 손해를 볼 수 있으므로 무차별적으로 사용해서는 안된다.

#### 클로저 활용시 주의사항

- 클로저의 프로퍼티 값은 쓰기 가능하므로 여러번 호출로 값이 변할 수 있다
- 하나의 클로저가 여러 함수 객체의 스코프 체인에 들어가 있는 경우도 있다
- 루프안에서 클로저 활용시 주의해야 한다

## 객체지향 프로그래밍

클래스 기반 객체지향 언어는 모든 인스턴스가 클래스에 정의된 대로 같은 구조이고 일반적으로 런타임에 바꿀 수 없다.

반면 프로토타입 기반의 언어는 객체의 자료구조, 메소드 등을 동적으로 바꿀 수 있다.

### 클래스, 생성자, 메소드

자바스크립트에서는 class라는 개념이 없다.(ES6에서 문법적으로 추가되었으나 프로토타입 기반의 특별한 함수이다)

자바스크립트는 기본 타입을 제외하면 모든것이 객체로 구성되어 있고, 함수 객체로 많은 것을 구현한다.
앞에서 살펴본대로 new 연산자와 생성자 함수를 사용해서 객체를 생성했다.

그러나 생성자 함수에서 객체 공통기능을 선언하게되면 (getter / setter 등의) 불필요한 메모리 사용이 발생한다.

이 문제 해결을 위해서 프로토타입에 접근해야 한다.

```js
function Person(arg) {
  this.name = arg;
}

Person.prototype.getName = function() {
  return this.name;
}

Person.prototype.setName = function(value) {
  this.name = value;
}

var me = new Person('me');
var you = new Person('you');
console.log(me.getName()); // me
console.log(you.getName()); // you
```

위의 코드는 더글라스 크락포드가 제시한 방식에 따라 다음처럼 쓸 수도 있다

```js
Function.prototype.method = function(name, func) {
  this.prototype[name] = func;
}

function Person(arg) {
  this.name = arg;
}

Person.method('getName', function() {
  return this.name;
});

Person.method('setName', function(value) {
  this.name = value;
});

Person.prototype.setName = function(value) {
  this.name = value;
}

var me = new Person('me');
var you = new Person('you');
console.log(me.getName()); // me
console.log(you.getName()); // you
```

### 상속

자바스크립트는 프로토타입 체인을 이용하여 상속을 구현할 수 있다.

```js
function createObject(o) {
  function F() {}
  F.prototype = o;
  return new F();
}

var person = {
  name: "zzoon",
  getName: function() {
    return this.name;
  },
  setName: function(value) {
    this.name = value;
  }
};

var student = createObject(person);

student.setName("me");
console.log(student.getName()); // me
```

`createObject()` 함수는 인자로 들어온 객체를 부모로 하는 자식 객체를 생성하여 반환한다.
위의 함수는 ES5에서 `Object.create()` 함수로 제공되므로 따로 구현할 필요는 없다.

### 캡슐화

```js
var Person = function(arg) {
  var name = arg;

  // Person 함수 객체의 프로토타입에 접근가능하도록 함수생성하여 반환함
  var Func = function() {}
  Func.prototype = {
    getName: function() {
      return name;
    },
    setName: function(arg) {
      name = arg;
    }
  };
  
  return Func;
}();

var me = new Person();
```

Person 함수의 private 멤버에 접근할 수 있는 메소드를 반환받는다.
다만, 접근하는 private 멤버가 객체나 배열이면 얕은 복사로 참조만을 반환하므로 이후 이를 쉽게 변경할 수 있다.

따라서 보통의 경우 객체를 반환하지 않고 객체의 주요 정보를 새로운 객체에 담아서 반환하는 방법을 많이 사용한다.
하지만 꼭 객체가 반환되어야 하는 경우에는 깊은 복사로 복사본을 만들어서 반환하는 방법을 사용하는 것이 좋다.

### 응용: 클래스의 기능을 가진 subClass 함수

다음을 이용하여 클래스 기능을 하는 함수를 구현하여 보자

- 함수의 프로토타입 체인
- extend 함수
- 인스턴스를 생성할 때 생성자 호출

#### 자식 클래스 생성 및 상속

subClass는 상속받을 클래스에 넣을 변수 및 메소드가 담긴 객체를 인자로 받아 부모함수를 상속하는 자식클래스를 만든다.

```js
var subClass = function() {
  // 클로저로 임시 함수객체가 한번만 생성되도록 함
  var F = function() {};

  return function(obj) {
    // 최상위 함수객체가 Function이어야 한다.
    var parent = this === window ? Function : this;

    // 자식 함수 객체 생성
    var child = function() {
      var _parent = child.parent;

      // 부모 생성자가 있으면 호출하고, 부모가 Function 인경우 최상위이므로 실행X
      if (_parent && _parent !== Function) {
        // 부모함수의 재귀적 호출
        _parent.apply(this, arguments);
      }

      // 생성자 호출
      if (child.prototype._init) {
        child.prototype._init.apply(this, arguments);
      }
    };

    // 프로토타입을 이용한 상속
    F.prototype = parent.prototype;
    child.prototype = new F();
    child.prototype.constructor = child;
    child.parent = parent;
    // 자식 함수객체에 subClass 함수가 있어야 하므로
    child.subClass = arguments.callee;

    // 사용자가 인자로 넣은 객체를 자식클래스에 넣어 확장함 (얕은복사)
    for (var i in obj) {
      if (obj.hasOwnProperty(i)) {
        child.prototype[i] = obj[i];
      }
    }
  }();
}
```

#### subClass 활용

subClass 함수로 상속예제를 만들어보자

```js
var personObj = {
  _init: function() {
    console.log("person init");
  },
  getName: function() {
    return this._name;
  },
  setName: function(name) {
    this._name = name;
  }
};

var studentObj = {
  _init: function() {
    console.log("student init");
  },
  getName: function() {
    return "Student Name: " + this._name;
  }
};

var Person = subClass(personObj);
var person = new Person();
person.setName("ME");
console.log(person.getName()); // ME

var Student = Person.subClass(studentObj);
var student = new Student();
student.setName("HAKSAENG");
console.log(student.getName()); // Student Name: HAKSAENG
```

## 함수형 프로그래밍

### 함수형 프로그래밍의 개념

함수형 프로그래밍은 함수의 조합으로 작업을 수행한다.

작업이 이루어지는 동안 작업에 필요한 데이터와 상태는 변하지 않는다. (no side effect)

### 자바스크립트에서 함수형 프로그래밍

자바스크립트에서도 함수형 프로그래밍이 가능하다

- 일급 객체로서의 함수
- 클로저

를 지원하기 때문이다.

함수형 프로그래밍을 활용한 예제를 작성해보자.

팩토리얼: 클로저로 숨겨지는 cache에는 팩토리얼을 연산한 값을 저장한다.

```js
var factorial = function() {
  var cache = { '0': 1 };
  var func = function(n) {
    var result = 0;

    if (typeof(cache[n]) === 'number') {
      result = cache[n];
    } else {
      result = cache[n] = n * func(n-1);
    }

    return result;
  }

  return func;
}();

console.log(factorial(10));
```

위의 패턴을 사용하여 팩토리얼과 피보나치 수열을 계산하는 함수를 받아 재사용하는 함수를 만들 수 있다.

```js
var cacher = function(cache, func) {
  var calculate = function(n) {
    if (typeof(cache[n]) === 'number') {
      result = cache[n];
    } else {
      result = cache[n] = func(calculate, n);
    }

    return result;
  }

  return calculate;
};

var fact = cacher({ '0': 1 }, function(func, n) {
  return n * func(n-1);
});

var fibo = cacher({ '0': 0, '1': 1 }, function(func, n) {
  return func(n-1) + func(n-2);
});

console.log(fact(10));
console.log(fibo(10));
```

### 함수형 프로그래밍을 활용한 주요 함수

#### apply

특정 데이터를 여러 함수를 적용하는 방식으로 작업을 수행하는 것을 적용(apply)한다라고 한다.

따라서 자바스크립트에서도 함수를 호출하는 역할을 하는 메소드를 apply라고 이름 붙이게 되었다.

#### 커링 (currying)

커링이란 특정 함수에서 정의된 인자의 일부를 넣어 고정시키고,
나머지를 인자로 받는 새로운 함수를 만드는 것을 의미한다.

커링은 자바스크립트에서 기본으로 지원하지 않으나 커링함수를 정의하여 사용할 수 있다.

```js
Function.prototype.curry = function() {
  var fn = this, args = Array.prototype.slice.call(arguments);
  return function() {
    return fn.apply(this, args.concat(Array.prototype.slice.call(arguments)));
  };
};
```

#### map

```js
Array.prototype.map = function(callback) {
  var obj = this;
  var value, mappedValue;
  var A = new Array(obj.length);

  for (var i = 0; i < obj.length; i++) {
    value = obj[i];
    mappedValue = callback.call(null, value);
    A[i] = mappedValue;
  }

  return A;
};
```

#### reduce

```js
Array.prototype.reduce = function(callback, memo) {
  var obj = this;
  var value, accumulatedValue = 0;

  for (var i = 0; i < obj.length; i++) {
    value = obj[i];
    accumulatedValue = callback.call(null, accumulatedValue, value);
  }

  return accumulatedValue;
}
```

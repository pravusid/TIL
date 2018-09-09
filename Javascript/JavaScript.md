# JavaScript

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

조건문 (if)

반복문 (for, while)

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
</script>우
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

myOjbect.func1();

// Output:
// 2
// 101
// 102
```

내부함수의 this바인딩을 의도한대로 사용하려면 부모함수의 this를 내부함수가 접근가능한 다른 변수에 저장하는 방법을 사용한다.

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

myOjbect.func1();

// Output:
// 2
// 3
// 4
```

또한, call과 apply 메소드를 통해서 this 바인딩을 명시적으로 할 수 있다

##### 생성자 함수를 호출할시 this 바인딩

자바스크립트 객체를 생성하는 방법은 크게 객체 리터럴 방식이나 생성자 함수를 이용하는 두가지 방식이 있다.

자바스크립트에서는 기존 함수에 new 연산자를 붙여서 호출하면 해당 함수는 생성자 함수로 동작한다.

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

앞에서 살펴본 함수객체의 prototype 프로퍼티와 숨은 프로퍼티인 [[Prototype]] 링크를 구분해야 한다.

자바스크립트에서 모든 객체는 (자신을 생성한 **생성자 함수의 prototype 프로퍼티**가 가리키는) 프로토타입 객체를
[[Prototype]] 링크로 연결하여 자신의 부모 객체로 설정한다.

`__proto__` 프로퍼티는 모든 객체에 존재하는 숨겨진 프로퍼티로 객체 자신의 프로토타입 객체를 가리키는 참조링크 정보이다.

ECMAScript에서는 이것을 [[Prototype]]] 프로퍼티로 정하였지만,
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
따라서 myObject는 Object() 함수의 prototype 프로퍼티가 가리키는 Object.protorype 객체를 자신의 프로토타입객체로 연결한다.

프로토타입 체이닝은 특정 객체의 프로퍼티나 메소드에 접근하려고 할 때,
해당 객체에 해당 프로퍼티나 메소드가 없다면 [[Prototype]] 링크를 따라 자신의 부모 역할을 하는 프로토타입 객체의 프로퍼티를 차례로 검색하는 것이다.

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

물론 이들 기본 내장 프로토타입 객체 또안 `Object.prototype`을 자신의 프로토타입으로 가지며 프로토타입 체이닝으로 연결된다.

자바스크립트에서는 표준 빌트인 프로토타입 객체에 사용자 정의 메소드 추가를 허용한다.

```js
String.prototype.testMethod = function() {
  console.log('This is the String.prototype.testMethod()');
};

var str = 'this is test';
str.testMethod(); // This is the String.prototype.testMehtod()

console.dir(String.prototype) // testMethod가 추가되어 있음
```

#### 프로토타입도 자바스크립트 객체다

함수가 생성돌 때, 자신의 prototype 프로퍼티에 연결되는 프로토타입 객체는 constructor 프로퍼티만을 가진 객체다.

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
변경된 프로토타입 객체로 [[Prototype]] ㄹ이크를 연결한다.

생성자 함수의 프로토타입이 변경되기 이전에 생성된 객체들은 기존 프로토타입 객체로 [[Prototype]] 링크를 유지한다.

## 실행 컨텍스트와 클로저

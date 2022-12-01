# Error in JavaScript

## Error 객체

<https://developer.mozilla.org/ko/docs/Web/JavaScript/Reference/Global_Objects/Error>

`new Error([message[, fileName[, lineNumber]]])`

- `message` (Optional): 에러 메시지
- `fileName` (Optional): fileName 속성 값. 기본값은 Error 생성자를 호출한 코드를 포함하고 있는 파일의 이름.
- `lineNumber` (Optional): lineNumber 속성 값. 기본값은 Error 생성자 호출을 포함한 줄 번호.

> `Error()` 함수를 사용하는 것과 `new Error()` 생성자를 사용하는 것은 동일한 결과를 출력함

```js
const x = Error("I was created using a function call!");
const y = new Error('I was constructed via the "new" keyword!');
```

### `Error.prototype`

<https://developer.mozilla.org/ko/docs/Web/JavaScript/Reference/Global_Objects/Error/prototype>

- `constructor`: 생성자 함수
- `message`: 에러 메시지
- `name`: 에러 이름

## 내장 Error

JavaScript에는 일반적인 Error 생성자 외에도 주요 Error 생성자가 있다

- `InternalError`: JavaScript 엔진의 내부에서 에러가 발생했음을 에러 인스턴스 생성
- `EvalError`: 전역 함수 `eval()`에서 발생하는 에러 인스턴스 생성
- `SyntaxError`: `eval()`이 코드를 분석하는 중 잘못된 구문을 만났음을 나타내는 에러 인스턴스 생성
- `RangeError`: 숫자 변수나 매개변수가 유효한 범위를 벗어났음을 나타내는 에러 인스턴스 생성
- `ReferenceError`: 잘못된 참조를 했음을 나타내는 에러 인스턴스 생성
- `TypeError`: 변수나 매개변수가 유효한 자료형이 아님을 나타내는 에러 인스턴스 생성
- `URIError`: `encodeURI()`나 `decodeURl()` 함수 처리중 발생하는 에러 인스턴스 생성

## 에러 던지기

```js
throw new Error("에러 메시지");
```

> `throw` 키워드로 에러를 사용하면 `try...catch` 구문을 통해 에러를 처리한다

`throw`로 에러가 아닌 객체를 던질 수도 있으나, `try...catch`로 에러를 처라하는 곳에서는 `Error` 객체를 기대하고 처리하므로 정상적인 에러처리가 어려워진다.

또한 `Promise` 객체 콜백의 `reject` 함수 인자에 `Error` 객체를 사용할 수 있다.

## 에러 구분하기

```js
try {
  errorFunction();
} catch (e) {
  if (e instanceof EvalError) {
    console.log(e.name);
  } else if (e instanceof RangeError) {
    console.log(e.name);
  }
}
```

## StackTrace

함수 호출이 있을 때마다 함수는 스택의 맨 위에 쌓이므로 StackTrace는 실행의 역순으로 출력된다.

실행환경에서 `console.trace()`를 통해서 스택 트레이스를 출력할 수 있다.

### CaptureStackTrace

> `captureStackTrace` 함수는 V8 엔진을 사용하는 브라우저나 Node.js 환경에서 사용가능하다.

`Error.captureStackTrace(targetObject[, constructorOpt])`

- `targetObject <Object>`
- `constructorOpt <Function>`

targetObject에 `.stack` 프로퍼티를 생성한다.
스택은 `Error.captureStackTrace()` 함수가 호출된 코드의 위치를 나타내는 문자열이다.

```js
const myObject = {};
Error.captureStackTrace(myObject);
myObject.stack; // Similar to `new Error().stack`
```

스택 트레이스의 첫 행은 `${myObject.name}: ${myObject.message}` 메시지로 시작한다.

선택적 인수 `constructorOpt`는 함수를 받는다.
값이 입력되면 `constructorOpt`의 상위 프레임을(`constructorOpt` 함수 포함) 스택 트레이스에서 숨긴다.

```js
function MyError() {
  Error.captureStackTrace(this, MyError);
}
```

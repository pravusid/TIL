# Generators

- <https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/for-await...of>
- <https://javascript.info/generators>
- <https://javascript.info/async-iterators-generators>

## 개요

제너레이터는 필요에 따라 순차적으로 여러 값을 반환(`yield`)할 수 있다.

## 제너레이터 함수

제너레이터를 생성하기위해서 `function*` 문법을 사용한다.

> `function* f(…)`, `function *f(…)`: 두 문법 모두 작동하지만 전자의 형태로 사용한다

제너레이터 함수가 호출되면 일반함수처럼 실행되지 않고 제너레이터 객체를 반환한다.

제너레이터는 `next()` 메소드를 포함한다.
`next()` 메소드가 호출되면 가장 가까운 `yield <value>`문 까지 실행한다. (value는 생략가능하고 생략되면 `undefined` 이다)
`yield`까지 실행이 끝나면 실행이 멈추고 value를 반환한다.

`next()` 메소드 실행결과는 다음 프로퍼티를 포함한 객체이다

- `value`: yielded value
- `done`: 함수 코드가 모두 실행되었다면 `true`, 이외의 경우 `false`

## `iterable`

`Symbol.iterator` 심볼은 객체에 대응하는 기본 이터레이터를 지정한다. `iterator`는 `iterable`을 반환한다.

```js
const iterableObj = {
  from: 1,
  to: 5,
  [Symbol.iterator]() {
    return {
      current: this.from,
      last: this.to,
      next() {
        if (this.current <= this.last) {
          return { done: false, value: this.current++ };
        } else {
          return { done: true };
        }
      }
    };
  }
};
```

제너레이터는 `iterable`이다.
따라서 `for..of` 반복문에서 사용하거나 spread operator 등을 사용할 수 있다.

제너레이터가 `iterable`이므로 `iterator`에서 제너레이터를 반환할 수 있다.
제너레이터를 사용하면 위의 코드와 동일한 기능을 더 간결하게 작성할 수 있다.

```js
const iterableObj = {
  from: 1,
  to: 5,
  *[Symbol.iterator]() {
    for(let value = this.from; value <= this.to; value++) {
      yield value;
    }
  }
};
```

## 제너레이터 합성

제너레이터 합성은 제너레이터를 상호포함하는 특수 기능이다.

예를 들어, 다음과 같은 일련의 숫자를 생성하는 함수가 있다고 하자.

```js
function* generateSequence(start, end) {
  for (let i = start; i <= end; i++) yield i;
}
```

한 제너레이터를 다른 제너레이터에 포함(embed)하기 위한 `yield*` 구문이 존재한다.

```js
function* generatePasswordCodes() {
  // 0..9
  yield* generateSequence(48, 57);

  // A..Z
  yield* generateSequence(65, 90);

  // a..z
  yield* generateSequence(97, 122);
}

let str = '';

for(let code of generatePasswordCodes()) {
  str += String.fromCharCode(code);
}

console.log(str); // 0..9A..Za..z
```

`yield*` 지시문은 실행을 다른 제너레이터로 위임한다.

결과는 중첩된 저너레이터에서 코드를 인라인으로 사용하는 것과 같다.

```js
function* generateAlphaNum() {
  // yield* generateSequence(48, 57);
  for (let i = 48; i <= 57; i++) yield i;

  // yield* generateSequence(65, 90);
  for (let i = 65; i <= 90; i++) yield i;

  // yield* generateSequence(97, 122);
  for (let i = 97; i <= 122; i++) yield i;
}
```

## 양방향 데이터 처리

제너레이터는 결과를 외부로 반환할 뿐만 아니라 제너레이터 내부로 값을 전달할 수도 있다.

내부로 값을 전달하기 위해서 `generator.next(arg)` 구문을 사용한다.
전달한 인자의 값은 `yield`의 결과값이 된다.

```js
function* gen() {
  let ask1 = yield "2 + 2 = ?";
  console.log(ask1); // 4

  let ask2 = yield "3 * 3 = ?"
  console.log(ask2); // 9
}

let generator = gen();
console.log(generator.next().value); // "2 + 2 = ?"
console.log(generator.next(4).value); // "3 * 3 = ?"
console.log(generator.next(9).done); // true
```

- 첫 번째 `next()`는 첫 `yield`까지 실행된다
- 결과 `'2 + 2 = ?'` 문자열을 외부로 반환한다
- 두 번째 `next(4)`는 인자 `4`를 첫 번째 `yield`로 전달하고 실행을 재개한다
- 두 번째 `yield`에 도착하면 외부로 값을 반환한다
- 세 번째 `next(9)`는 인자 `9`를 두 번째 `yield`로 전달하고 실행을 재개한다
- 함수의 마지막 까지 실행되고 `done: true`가 된다

> 각각의 `next(value)`(첫 번째 호출을 제외하고) 호출은 인자 값을 제너레이터로 전달해서 현재 `yield`의 결과가 되고 다음 `yield`의 결과를 반환받는다.

## `generator.throw`

`yield`에 오류를 전달하려면 `generator.throw(error)`를 호출하면 된다.
이 경우 `error`는 `yield`가 위치한 라인에서 `throw`된다.

```js
function* gen() {
  try {
    let result = yield "2 + 2 = ?"; // Error in this line
    console.log("이 곳은 실행되지 않음");
  } catch(e) {
    console.log(e); // Error: Some error message
  }
}

let generator = gen();
let question = generator.next().value;
generator.throw(new Error("Some error message"));
```

제너레이터 내부에서 에러가 처리되지 않으면 에러는 코드를 실행한 위치로 던져진다.

```js
function* generate() {
  let result = yield "2 + 2 = ?"; // Error in this line
}

let generator = generate();
let question = generator.next().value;

try {
  generator.throw(new Error("Some error message"));
} catch(e) {
  alert(e); // Error: Some error message
}
```

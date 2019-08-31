# Generators

<https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/for-await...of>

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

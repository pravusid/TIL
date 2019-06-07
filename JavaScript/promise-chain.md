# Promises Chaining

## 내장 reduce 함수 사용

JavaScript의 reduce를 활용하여 promise를 순서대로 실행하고 결과값을 모아서 반환

method를 `promiseFn`으로 넘길 때는 `bind()`를 사용하여 `this` 바인딩을 한다

```ts
const results = await chainPromisesAccumulator(foo.promiseFn.bind(foo), conditions);
```

method 실행이전 명령을 정의하기 위해서 method를 포함한 lambda 함수를 작성하여 넘긴다면,
`call()`을 사용하여 `this` 바인딩을 한다

```ts
const results = await chainPromisesAccumulator(() => {
  // ...
  foo.promiseFn.call(foo);
}, conditions);
```

`promiseFn`을 받아서 연쇄 실행하는 함수는 다음과 같다

```ts
export const chainPromisesAccumulator = <R, U>(
  promiseFn: (condition: U, ...args: any[]) => Promise<R>,
  conditions: U[],
  ...args: any[]
): Promise<R[]> => {
  return [...Array(conditions.length).keys()].reduce(async (promise, i) => {
    return (async (accumulator: R[]) => {
      accumulator.push(await promiseFn(conditions[i], ...args));
      return accumulator;
    })(await promise);
  }, Promise.resolve([] as R[]));
};
```

## Async 라이브러리

<https://caolan.github.io/async/docs.html>

`Collection` 카테고리의 `---Series` 함수를 사용하여 순차실행한다

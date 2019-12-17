# Promises Chaining

## 내장 reduce 함수 사용

JavaScript의 reduce를 활용하여 promise를 순서대로 실행하고 결과값을 모아서 반환

클래스의 메소드를 함수인자로 넘길 때는 `bind()`를 사용하여 `this` 바인딩을 한다

```ts
const results = await chainPromisesAccumulator(foo.someMethod.bind(foo), conditions);
```

method 실행이전 명령을 정의하기 위해서 method를 포함한 lambda 함수를 작성하여 넘긴다면, `call()`을 사용하여 `this` 바인딩을 한다

```ts
const results = await chainPromisesAccumulator(() => {
  // ...
  foo.someMethod.call(foo);
}, conditions);
```

`Promise`를 반환하는 함수를 받아서 연쇄 실행하는 함수는 다음과 같다

```ts
export const chainPromisesAccumulator = <R, U>(
  func: (condition: U, ...args: any[]) => Promise<R>,
  conditions: U[],
  ...args: any[]
): Promise<R[]> =>
  [...Array(conditions.length).keys()].reduce(
    async (promise, i) => (async (accumulator: R[]) => [...accumulator, await func(conditions[i], ...args)])(await promise),
    Promise.resolve<R[]>([]),
  );
```

테스트

```ts
import { chainPromisesAccumulator } from './chain.promises';

describe('chain promises', () => {
  it('test chain', async () => {
    const interval = 1000;
    const repeat = 3;
    const promiseFn: (arg: number) => Promise<number> = arg =>
      new Promise(resolve => {
        setTimeout(() => {
          const ret = new Date().getSeconds();
          console.log('arg', arg, 'ret', ret);
          resolve(ret);
        }, interval);
      });

    const begin = new Date().getSeconds();
    const result = await chainPromisesAccumulator(promiseFn, [...Array(repeat).keys()]);

    const expectedResult = b => {
      const ret = [];
      for (let i = 1; i <= repeat; i += 1) {
        const val = b + i * (interval / 1000);
        ret.push(val >= 60 ? val - 60 : val);
      }
      return ret;
    };
    expect(result).toEqual(expectedResult(begin));
  }, 10000);
});
```

## Async 라이브러리

<https://caolan.github.io/async/docs.html>

`Collection` 카테고리의 `---Series` 함수를 사용하여 순차실행한다

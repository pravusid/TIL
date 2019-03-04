# Promises Chaining

JavaScript의 reduce를 활용하여 promise를 순서대로 실행하고 결과값을 모아서 반환

클래스 내부의 method(`this.methodName()`)를 `promiseFn`으로 넘길 때는 `this` 바인딩을 해야 함

`const results = await chainPromisesAccumulator(promiseFn.bind(this), conditions);`

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
  }, Promise.resolve(<R[]>[]));
};
```

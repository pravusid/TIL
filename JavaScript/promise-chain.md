# Promises Chaining

JavaScript의 reduce를 활용하여 promise를 순서대로 실행하고 결과값을 모아서 반환

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

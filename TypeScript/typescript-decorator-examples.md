# TypeScript Decorator 예제

## ECMAScript Decorators

[[typescript-handbook-decorators]]

- <https://devblogs.microsoft.com/typescript/announcing-typescript-5-0/#decorators>
- <https://devblogs.microsoft.com/typescript/announcing-typescript-5-2/#decorator-metadata>

## (Experimetal) Method Hook

```ts
/* decorator */
export type HookTarget<R> = (...args: unknown[]) => Promise<R>;

export function Hook<R>(hookDescriptor: (hookTarget: HookTarget<R>) => Promise<R>) {
  return function (target: unknown, propertyKey: string, descriptor: PropertyDescriptor) {
    const original = descriptor.value;
    descriptor.value = function (...args: unknown[]) {
      return hookDescriptor(original.bind(this, ...args));
    };
  };
}

/* 적용 */
let counter = 0;

export async function count<R>(target: HookTarget<R>) {
  counter += 1;
  return target().finally(() => {
    counter -= 1;
  });
}

class Foo {
  @Hook((fn) => count(fn))
  async test(name: string): Promise<void> {
    // ...
  }
}
```

### 화살표 함수 사용시

```ts
export function Hook<R>(hookDescriptor: (hookTarget: HookTarget<R>) => Promise<R>) {
  return (target: unknown, propertyKey: string, descriptor: PropertyDescriptor) => {
    const original = descriptor.value;
    descriptor.value = (...args: unknown[]) => {
      return hookDescriptor(original.bind(target, ...args));
    };
  };
}
```

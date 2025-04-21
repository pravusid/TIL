# TypeScript Decorator 예제

## ECMAScript Decorators

[[typescript-handbook-decorators]]

- <https://devblogs.microsoft.com/typescript/announcing-typescript-5-0/#decorators>
- <https://devblogs.microsoft.com/typescript/announcing-typescript-5-2/#decorator-metadata>

## (Experimetal) Method Hook

```ts
/* decorator */
export type HookTarget<R> = (...args: unknown[]) => R | Promise<R>;

export function Hook<R>(hookDescriptor: (hookTarget: HookTarget<R>) => R | Promise<R>) {
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
  try {
    return await target();
  } finally {
    counter -= 1;
  }
}

class Foo {
  @Hook((fn) => count(fn))
  async test(name: string): Promise<void> {
    // ...
  }
}
```

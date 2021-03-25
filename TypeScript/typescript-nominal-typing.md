# Nominal Typing

타입스크립트는 기본적으로 Structural typing 구조를 채택하고 있다.

필요에 따라 Nominal Typing을 사용하고 싶을 때가 있다면 몇 가지 트릭을 적용할 수 있다.

<https://basarat.gitbook.io/typescript/main-1/nominaltyping>

- literal type
- enum
- interface

Nominal Typing을 위해 세 가지 방법이 제시되어 있지만, 인터페이스를 이용하는 방식만 살펴본다.

> This method is still used by the TypeScript compiler team, so worth mentioning.
> Using `_` prefix and a Brand suffix is a convention I strongly recommend (and the one followed by the TypeScript team).
> <https://github.com/Microsoft/TypeScript/blob/7b48a182c05ea4dea81bab73ecbbe9e013a79e99/src/compiler/types.ts#L693-L698>

방법은 간단하다

Structural typing의 호환성을 피하기 위해서 사용하지 않는 프로퍼티(`_ ~ Brand` 컨벤션)를 타입에 추가하고, 값을 설정할 때 type assertion을 사용한다.

```ts
// FOO
interface FooId extends String {
  _fooIdBrand: string; // To prevent type errors
}

// BAR
interface BarId extends String {
  _barIdBrand: string; // To prevent type errors
}

/**
 * Usage Demo
 */
var fooId: FooId;
var barId: BarId;

// Safety!
fooId = barId; // error
barId = fooId; // error
fooId = <FooId>barId; // error
barId = <BarId>fooId; // error

// Newing up
fooId = 'foo' as any;
barId = 'bar' as any;

// If you need the base string
var str: string;
str = fooId as any;
str = barId as any;
```

for me

```ts
interface FooBarId extends String {
  readonly _fooBarIdBrand: never
}
```

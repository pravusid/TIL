# TypeScript

## 기본 타입

- 부울: `boolean`: true / false
- 숫자: `number`: 자바스크립트와 마찬가지로 64비트 부동 소수점 값이다
- 문자: `string`: 큰 따옴표, 작은 따옴표, 템플릿 문자열을 위한 백 쿼트를 사용할 수 있다
- 배열: `T[]`, `Array<T>`: 두 가지 방식으로 선언할 수 있다
- 튜블: `[T, U]`: 고정된 개수의 요소와 타입(같을 필요 없음)을 표현한다
- 열거: `enum T {A, B, C}`: enumeration 타입의 요소는 순서대로 0부터 시작하는 키값을 갖는다
- `any`: 알지 못하는 변수 타입 (최상위 타입으로 쓸 수도 있다)
- `void`: 일반적으로 반환이 없는 함수의 반환타입으로 사용됨. `undefined` 또는 `null`만 할당할 수 있다
- `undefined` / `null`: 다른 모든 타입의 서브 타입이다
- `never`: 절대로 발생하지 않는 값의 타입, 다른 모든 타입의 서브 타입니다

### Type assertions

다음과 같이 `as T`의 형식으로 사용되며, 컴파일러에게 해당 타입을 알려주는 역할을 한다.

```ts
const someValue: any = "this is a string";
const strLength: number = (someValue as string).length;
```

## 인터페이스

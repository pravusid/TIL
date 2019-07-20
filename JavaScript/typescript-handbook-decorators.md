# TypeScript HandBook: Decorators

## Introduction

TypeScript 및 ES6에 클래스가 도입됨에 따라 클래스 및 클래스 멤버에 주석을 달거나 수정하는데 필요한 추가기능이 필요한 시나리오가 있다.

데코레이터는 클래스 선언 및 멤버에 대한 주석과 메타 프로그래밍 구문을 모두 추가할 수 있는 방법을 제공한다.
데코레이터는 현재 JavaScript에서 Stage-2이며 TypeScript의 `experimentalDecorators` 옵션으로 사용할 수 있다.

## Decorators

데코레이터는 클래스 선언, method, accessor, property, parameter에 첨부할 수 있는 특별한 종류의 선언이다.

데코레이터는 `@expression` 형식을 사용하고,
`expression`은 데코레이션된 선언에 대한 정보와 함께 런타임에 호출되는 함수로 평가하여야 한다.

예를 들어 데코레이터 `@sealed`를 사용하면 다음과 같이 `sealed` 함수를 작성할 수 있다.

```ts
function sealed(target) {
  // do something with 'target' ...
}
```

### Decorator Factories

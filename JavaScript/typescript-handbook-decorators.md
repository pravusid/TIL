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

데코레이터를 선언에 적용하는 방법을 활용하여 데코레이터 팩토리를 작성할 수 있다.
데코레이터 팩토리는 런타임에 데코레이터에 의해 호출 될 표현식을 반환하는 함수이다.

```ts
function color(value: string) { // this is the decorator factory
  return function (target) { // this is the decorator
    // do something with 'target' and 'value'...
  }
}
```

### Decorator Composition

여러 데코레이터를 선언에 적용할 수 있다

```ts
@f @g x
// 또는
@f
@g
x
```

여러 데코레이터가 하나의 선언에 적용되는 경우의 평가는 함수 합성과 유사하다.

- 각 데코레이터에 대한 표현식은 위에서 아래로 평가된다
- 결과들은 함수로서 아래부터 위로 호출된다

```ts
function f() {
  console.log("f(): evaluated");
  return function (target, propertyKey: string, descriptor: PropertyDescriptor) {
    console.log("f(): called");
  }
}

function g() {
  console.log("g(): evaluated");
  return function (target, propertyKey: string, descriptor: PropertyDescriptor) {
    console.log("g(): called");
  }
}

class C {
  @f()
  @g()
  method() {}
}
```

다음과 같이 출력된다

```ts
f(): evaluated
g(): evaluated
g(): called
f(): called
```

### Decorator Evaluation

클래스 내부의 다양한 선언에 대응해 데코레이터가 적용되는 순서이다

- Parameter Decorators, followed by Method, Accessor, or Property Decorators are applied for each instance member.
- Parameter Decorators, followed by Method, Accessor, or Property Decorators are applied for each static member.
- Parameter Decorators are applied for the constructor.
- Class Decorators are applied for the class

### Class Decorators

클래스 데코레이터는 클래스 선언 직전에 선언되며 클래스 생성자에 적용된다.
클래스 데코레이터는 선언파일이나 다른 amibient context(i.e. declare class)에서 사용할 수 없다.

클래스 데코레이터에 대한 표현식은 런타임에 함수로 호출되며 데코레이팅 된 클래스의 생성자가 유일한 인수로 호출된다.

클래스 데코레이터가 값을 반환하면 제공된 생성자 함수로 클래스 선언을 바꾼다.

> 새 생성자 함수를 반환하려면 원래의 프로토타입을 유지 관리해야 한다. 런타임에 데코레이터를 적용하는 로직은 이것을 하지 않는다.

다음은 `Greeter` 클래스에 적용된 `@sealed` 데코레이터 예제이다

```ts
@sealed
class Greeter {
  greeting: string;
  constructor(message: string) {
    this.greeting = message;
  }
  greet() {
    return "Hello, " + this.greeting;
  }
}
```

다음 함수선언을 통해 `@sealed` 데코레이터를 정의할 수 있다

```ts
function sealed(constructor: Function) {
  Object.seal(constructor);
  Object.seal(constructor.prototype);
}
```

`@sealed`가 실행되면 생성자와 포로토타입 모두 `seal` 처리한다

다음으로 생성자를 override하는 예제를 확인해보자

```ts
function classDecorator<T extends {new(...args:any[]):{}}>(constructor:T) {
  return class extends constructor {
    newProperty = "new property";
    hello = "override";
  }
}

@classDecorator
class Greeter {
  property = "property";
  hello: string;
  constructor(m: string) {
    this.hello = m;
  }
}

console.log(new Greeter("world"));
// 결과
// hello: "override"
// newProperty: "new property"
// property: "property"
```

### Method Decorators

메소드 데코레이터는 메소드 선언 직전 선언된다.
데코레이터는 메소드 property descriptor에 적용되며 메소드 정의를 관찰, 수정 또는 바꾸는데 사용한다.

메소드 데코레이터는 선언파일, overload, ambient context(i.e. declare class)에서 사용할 수 없다.

메소드 데코레이터의 표현식은 런타임에 다음 세가지 인수와 함께 함수로 호출된다.

- 클래스 스태틱 멤버는 생성자 함수, 인스턴스 멤버는 클래스 프로토타입
- member의 이름
- 멤버의 Property Descriptor

메소드 데코레이터가 값을 반환하면 메소드의 property descriptor로 사용된다.

> 타겟이 ES5 미만이면 property descriptor가 `undefined`이고, 데코레이터 반환값도 무시된다

`Greeter` 클래스의 메소드에 적용된 메소드 데코레이터(`@enumerable`)

```ts
class Greeter {
  greeting: string;
  constructor(message: string) {
    this.greeting = message;
  }

  @enumerable(false)
  greet() {
    return "Hello, " + this.greeting;
  }
}
```

다음 함수 선언을 사용해서 `@enumerable` 데코레이터를 정의할 수 있다

```ts
function enumerable(value: boolean) {
  return function (target: any, propertyKey: string, descriptor: PropertyDescriptor) {
    descriptor.enumerable = value;
  };
}
```

`@enumerable(false)` 데코레이터는 데코레이터 팩토리이다.
`@enumerable(false)` 데코레이터가 호출되면 property descriptor의 `enumerable` 프로퍼티를 수정한다.

### Accessor Decorators

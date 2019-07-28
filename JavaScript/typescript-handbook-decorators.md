# TypeScript HandBook: Decorators

<https://rbuckton.github.io/reflect-metadata/#syntax>

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
function color(value: string) {
  // this is the decorator factory
  return function(target) {
    // this is the decorator
    // do something with 'target' and 'value'...
  };
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
  return function(target, propertyKey: string, descriptor: PropertyDescriptor) {
    console.log("f(): called");
  };
}

function g() {
  console.log("g(): evaluated");
  return function(target, propertyKey: string, descriptor: PropertyDescriptor) {
    console.log("g(): called");
  };
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

- 메소드 다음에 나오는 파라메터 데코레이터, 각 instance 멤버에 적용되는 메소드/접근자/프로퍼티 데코레이터
- 메소드 다음에 나오는 파라메터 데코레이터, 각 static 멤버에 적용되는 메소드/접근자/프로퍼티 데코레이터
- 생성자에 적용된 파라메터 데코레이터
- 클래스에 적용되는 클래스 데코레이터

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
function classDecorator<T extends { new (...args: any[]): {} }>(
  constructor: T
) {
  return class extends constructor {
    newProperty = "new property";
    hello = "override";
  };
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

- static 멤버에는 클래스 생성자 함수, instance 멤버에는 클래스의 프로토타입
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
  return function(
    target: any,
    propertyKey: string,
    descriptor: PropertyDescriptor
  ) {
    descriptor.enumerable = value;
  };
}
```

`@enumerable(false)` 데코레이터는 데코레이터 팩토리이다.
`@enumerable(false)` 데코레이터가 호출되면 property descriptor의 `enumerable` 프로퍼티를 수정한다.

### Accessor Decorators

접근자 데코레이터는 접근자 직전 선언된다.
접근자 데코레이터는 선언파일이나 다른 ambient context에서 사용할 수 없다.

> TypeScript는 단일 멤버의 `get`, `set` 접근자에 동시에 데코레이터를 사용하지 못하도록 한다. 대신 멤버에 지정된 데코레이터는 문서 순서에 정의된 첫 번째 접근자에 적용해야 한다. 데코레이터는 property descriptor에 적용되기 때문이다.(`get`, `set` 접근자는 별도로 선언되는 것이 아니라 결합되어 있음)

접근자 데코레이터에 대한 표현식은 런타임에 다음 세 가지 인수와 함께 함수로 호출된다

- static 멤버에는 클래스 생성자 함수, instance 멤버에는 클래스의 프로토타입
- 멤버의 이름
- 멤버의 Property Descriptor

다음은 `Point` 클래스 멤버에 적용된 접근자 데코레이터 예시이다

```ts
class Point {
  private _x: number;
  private _y: number;
  constructor(x: number, y: number) {
    this._x = x;
    this._y = y;
  }

  @configurable(false)
  get x() {
    return this._x;
  }

  @configurable(false)
  get y() {
    return this._y;
  }
}
```

다음과 같은 함수 선언을 사용하여 `@configurable` 데코레이터를 정의할 수 있다

```ts
function configurable(value: boolean) {
  return function(
    target: any,
    propertyKey: string,
    descriptor: PropertyDescriptor
  ) {
    descriptor.configurable = value;
  };
}
```

### Property Decorators

프로퍼티 데코레이터는 속성 선언 직전에 선언된다.
프로퍼티 데코레이터는 선언 파일이나 다른 ambient context에서 사용할 수 없다.

프로퍼티 데코레이터에 대한 표현식은 런타임에 다음 두 개의 인수를 사용하여 함수로 호출된다.

- static 멤버에는 클래스 생성자 함수, instance 멤버에는 클래스의 프로토타입
- 멤버의 이름

> TypeScript에서 프로퍼티 데코레이터가 초기화 되는 방법 때문에, 프로퍼티 데코레이터의 인자로 Property Descriptor는 제공되지 않는다. 프로토타입의 멤버를 정의할 때 인스턴스 프로퍼티를 설명하는 메커니즘이 없으며, 속성의 initializer를 관찰하거나 수정할 방법이 없기 때문이다. 반환 값 역시 무시된다. 따라서 프로퍼티 데코레이터는 특정 이름의 프로퍼티가 클래스에 대해 선언된 것을 관찰하는 데만 사용할 수 있다.

이 정보를 사용하여 예제와 같이 프로퍼티에 대한 메타데이터를 기록할 수 있다

```ts
class Greeter {
  @format("Hello, %s")
  greeting: string;

  constructor(message: string) {
    this.greeting = message;
  }
  greet() {
    let formatString = getFormat(this, "greeting");
    return formatString.replace("%s", this.greeting);
  }
}
```

함수 선언을 사용하여 `@format` 데코레이터와 `getFormat` 함수를 정의할 수 있다

```ts
import "reflect-metadata";

const formatMetadataKey = Symbol("format");

function format(formatString: string) {
  return Reflect.metadata(formatMetadataKey, formatString);
}

function getFormat(target: any, propertyKey: string) {
  return Reflect.getMetadata(formatMetadataKey, target, propertyKey);
}
```

`@format("Hello, %s")` 데코레이터는 데코레이터 팩토리이다.
데코레이터가 호출 되면 `reflect-metadata` 라이브러리의 `Reflect.metadata` 함수를 사용하여 프로퍼티에 대한 메타데이터 항목을 추가한다.

`getFormat`이 호출되면 포맷의 metadata 값을 읽는다.

### Parameter Decorators

파라미터 데코레이터는 파라미터 선언 직전에 선언된다.
파라미터 데코레이터는 클래스 생성자 함수 또는 메소드 선언 함수에 적용된다.

파라미터 데코레이터는 선언파일, overload, ambient context에서 사용할 수 없다.

파라미터 데코레이터에 대한 표현식은 런타임에 다음 세 가지 인수와 함께 함수로 호출된다.

- static 멤버에는 클래스 생성자 함수, instance 멤버에는 클래스의 프로토타입
- 멤버의 이름
- 함수의 파라미터 목록중 해당 파라미터의 순서(index)

> 파라미터 데코레이터는 메소드에 선언된 매개변수를 관찰 하는데만 사용할 수 있다

파라미터 데코레이터의 반환 값은 무시된다

다음은 `Greeter` 클래스 멤버의 파라미터에 적용된 파라미터 데코레이터(`@required`) 예제이다

```ts
class Greeter {
  greeting: string;

  constructor(message: string) {
    this.greeting = message;
  }

  @validate
  greet(@required name: string) {
    return "Hello " + name + ", " + this.greeting;
  }
}
```

다음 함수 선언을 사용하여 `@required` 및 `@validate` 데코레이터를 정의할 수 있다.

```ts
import "reflect-metadata";

const requiredMetadataKey = Symbol("required");

function required(target: Object, propertyKey: string | symbol, parameterIndex: number) {
  let existingRequiredParameters: number[] = Reflect.getOwnMetadata(requiredMetadataKey, target, propertyKey) || [];
  existingRequiredParameters.push(parameterIndex);
  Reflect.defineMetadata(requiredMetadataKey, existingRequiredParameters, target, propertyKey);
}

function validate(target: any, propertyName: string, descriptor: TypedPropertyDescriptor<Function>) {
  let method = descriptor.value;
  descriptor.value = function () {
    let requiredParameters: number[] = Reflect.getOwnMetadata(requiredMetadataKey, target, propertyName);
    if (requiredParameters) {
      for (let parameterIndex of requiredParameters) {
        if (parameterIndex >= arguments.length || arguments[parameterIndex] === undefined) {
          throw new Error("Missing required argument.");
        }
      }
    }

    return method.apply(this, arguments);
  }
}
```

- `@required` 데코레이터는 optional이 아닌 파라미터를 표기하는 메타데이터 목록을 추가한다.
- `@validate` 데코레이터는 기존 메소드를 호출하기 전에 인수 유효성을 검증하는 함수에 `greet` 메소드를 래핑한다.

### Metadata

예제에서 실험적인 metadata API의 polyfill을 추가하는 `reflect-metadata` 라이브러리를 사용하고 있다.
하지만 이 라이브러리는 아직 ECMAScript 표준의 일부가 아니다.

TypeScript는 데코레이터가 있는 선언에 대해 특정 타입의 메타데이터를 내보내는 실험적인 지원을 포함한다.
이 실험 기능을 활성화 하려면 `emitDecoratorMetadata` 컴파일러 옵션을 활성화 해야 한다.

컴파일러 옵션이 활성화 된 경우 `reflect-metata` 라이브러리리를 불러오면, 추가적인 디자인 타임의 타입정보가 런타임에 표시된다.

```ts
import "reflect-metadata";

class Point {
  x: number;
  y: number;
}

class Line {
  private _p0: Point;
  private _p1: Point;

  @validate
  set p0(value: Point) { this._p0 = value; }
  get p0() { return this._p0; }

  @validate
  set p1(value: Point) { this._p1 = value; }
  get p1() { return this._p1; }
}

function validate<T>(target: any, propertyKey: string, descriptor: TypedPropertyDescriptor<T>) {
  let set = descriptor.set;
  descriptor.set = function (value: T) {
    let type = Reflect.getMetadata("design:type", target, propertyKey);
    if (!(value instanceof type)) {
      throw new TypeError("Invalid type.");
    }
    set.call(target, value);
  }
}
```

TypeScript 컴파일러는 `@Reflect.metadata` 데코레이터를 사용해서 디자인 타임 타입정보를 주입힌다.

다음 코드와 동일한 것으로 간주할 수 있다.

```ts
class Line {
  private _p0: Point;
  private _p1: Point;

  @validate
  @Reflect.metadata("design:type", Point)
  set p0(value: Point) { this._p0 = value; }
  get p0() { return this._p0; }

  @validate
  @Reflect.metadata("design:type", Point)
  set p1(value: Point) { this._p1 = value; }
  get p1() { return this._p1; }
}
```

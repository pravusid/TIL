# TypeScript HandBook: Type System

## 타입 추론

TypeScript의 타입추론은 변수나 멤버를 초기화하거나 파라미터의 기본값을 설정하거나 함수의 반환 타입을 결정할 때 발생한다.

### Best common type

컴파일러는 타입 추론을 위해 제공된 후보를 분석하여 가장 적합한 공통유형을 찾는다.

하지만 모든 유형의 super type이 명시적으로 존재하지 않으면 union array type으로 처리되므로(`(Rhino | Elephant | Snake)[].`) 필요하다면 super type을 명시해야 한다.

```ts
let zoo: Animal[] = [new Rhino(), new Elephant(), new Snake()];
```

### Contextual Typing

문맥상 타이핑은 위치에 의해 표현식의 타입이 암시될 때 발생한다.

아래애서 mouseEvent의 문맥이 없다면 `any` 타입으로 처리되어 `button` 프로퍼티를 사용할 수 없을 것이다.

```ts
window.onmousedown = function(mouseEvent) {
  console.log(mouseEvent.button); //<- OK
  console.log(mouseEvent.kangaroo); //<- Error!
};
```

함수가 문맥이 없는 위치에 있다면 인수는 암시적으로 `any` 유형이 되므로 위의 오류는 발생하지 않는다.
(컴파일러에서 `--noImplicitAny` 옵션을 사용하여 이를 방지할 수 있다)

문맥상 타이핑은 함수 호출의 인수, 할당문의 우측, type assertion, 객체와 배열 리터럴의 멤버와 반환문 등에 적용된다.

## 타입 호환성

TypeScript의 타입 호환성은 구조적 서브타이핑을 기반으로 한다.

아래에서 Person 클래스는 Named 인터페이스 구현을 명시하지 않았으나, 해당 인터페이스의 구조를 따르고 있다.

```ts
interface Named {
  name: string;
}

class Person {
  name: string;
}

let p: Named;
p = new Person();
```

TypeScript 구조적 타입 시스템의 기본 규칙은, y가 x의 멤버를 모두 구성한다면 x가 y와 호환된다는 것이다.

### 함수 비교

```ts
let x = (a: number) => 0;
let y = (b: number, s: string) => 0;

y = x;
x = y; // ERROR
```

두 함수가 서로 호환 가능한지 확인하기 위해 우선 매개 변수 목록을 확인한다.
함수의 매개 변수 비교는 상응하는 호환 가능한 타입의 매개변수에 대응되고, 이름은 고려되지 않는다.

y의 매개변수는 두 개이지만, x는 매개변수가 하나 밖에 없다.
하지만 JavaScript에서 함수의 매개 변수가 무시하는 것이 허용되므로 할당이 가능하다.

```ts
let x = () => ({ name: "Alice" });
let y = () => ({ name: "Alice", location: "Seattle" });

x = y;
y = x; // ERROR
```

타입시스템은 원본 함수 반환형이 대상 타입 반환형의 서브타입이 되도록 강제한다 (공변성)

#### Function Parameter Bivariance

함수 파라미터의 타입을 비교할 때, 원본 파라미터가 대상 파라미터에 할당 가능하거나 그 반대일 경우 할당이 성공한다.
하지만, 호출한 쪽에서 less specialized type argument로 함수를 호출할 수도 있는데 이는 바람직 하지 않다.

```ts
enum EventType {
  Mouse,
  Keyboard
}

interface Event {
  timestamp: number;
}
interface MouseEvent extends Event {
  x: number;
  y: number;
}
interface KeyEvent extends Event {
  keyCode: number;
}

function listenEvent(eventType: EventType, handler: (n: Event) => void) {
  /* ... */
}

// 바람직하지 않지만, 유용하고 일반적이다
listenEvent(EventType.Mouse, (e: MouseEvent) => console.log(e.x + "," + e.y));

// 바람직하지 않은 대안
listenEvent(EventType.Mouse, (e: Event) =>
  console.log((<MouseEvent>e).x + "," + (<MouseEvent>e).y)
);
listenEvent(EventType.Mouse, <(e: Event) => void>(
  ((e: MouseEvent) => console.log(e.x + "," + e.y))
));

// 명백한 오류: 완전히 호환되지 않는 타입이 강제됨
listenEvent(EventType.Mouse, (e: number) => console.log(e));
```

#### Optional Parameters and Rest Parameters

함수를 비교할 때 선택적 파라미터와 필수 파라미터는 서로 바꿔 사용할 수 있다.

원본 타입에 추가적인 선택적 파라미터는 오류가 아니며,
대상 타입의 선택적 파라미터에 원본타입에 대응하는 파라미터가 없더라도 오류가 아니다.

함수가 rest 파라미터를 가지고 있다면, 무한대의 선택적 파라미터 연속으로 처리된다.

이는 타입시스템의 관점에서 보면 적절하지 않지만,
런타임 관점에서 본다면 대부분의 함수에서 해당위치에 `undefined`를 전달하기 때문에 선택적 파라미터의 개념은 일반적으로 잘 적용된다 보기 어렵다.

```ts
function invokeLater(args: any[], callback: (...args: any[]) => void) {
  /* ... Invoke callback with 'args' ... */
}

// 부적절함: invokeLater는 아마도 정해지지 않은 수의 인수를 제공할 것이다
invokeLater([1, 2], (x, y) => console.log(x + ", " + y));

// 혼란스러움: x, y 는 실제로 필요하다
invokeLater([1, 2], (x?, y?) => console.log(x + ", " + y));
```

#### Functions with overloads

오버로딩한 함수는 원본 유형과 대상유형의 호환가능한 시그니처와 일치해야 한다

### Enums 타입 호환성

Enum 타입은 number 타입과 상호호환된다.
하지만 다른 Enum에서 가져온 Enum 값은 호환되지 않는다.

```ts
enum Status {
  Ready,
  Waiting
}
enum Color {
  Red,
  Blue,
  Green
}

let status = Status.Ready;
status = Color.Green; // ERROR
```

### Class 타입 호환성

클래스는 한 가지만 제외하고 객체 리터럴 타입이나 인터페이스와 동일하게 작동한다.

클래스유형의 두 객체를 비교할 때 인스턴스 멤버만 비교하고, 정적 멤버 및 생성자는 호환성에 영향을 주지 않는다.

```ts
class Animal {
  feet: number;
  constructor(name: string, numFeet: number) {}
}

class Size {
  feet: number;
  constructor(numFeet: number) {}
}

let a: Animal;
let s: Size;

a = s; // OK
s = a; // OK
```

#### Private and protected members in classes

클래스의 private | protected 멤버는 호환성에 영향을 준다.

이를 통해 클래스는 super class와 호환되어 할당 가능하지만, 동일한 형태의 다른 상속 계층 클래스와는 호환되지 않는다.

### Generics 타입 호환성

TypeScript는 구조적 타입 시스템이므로 타입 파라미터는 멤버의 일부로 사용될 때만 결과 타입에 영향을 준다.

```ts
interface Empty<T> {}
let x: Empty<number>;
let y: Empty<string>;

x = y; // OK, because y matches structure of x
```

위의 경우는 타입 파라미터를 다르게 사용하지 않으므로 구조적으로 호환 가능하다.

```ts
interface NotEmpty<T> {
  data: T;
}
let x: NotEmpty<number>;
let y: NotEmpty<string>;

x = y; // Error, because x and y are not compatible
```

타입 인자가 명시된 제네릭 타입은 실제로는 비-제네릭 타입처럼 작동한다.

타입 인자가 지정되지 않은 제네릭 타입의 경우 지정되지 않은 타입 인자 대신 `any`를 지정하여 호환성을 확인한다.
그리고 나서, 비-제네릭 타입과 같은 방식으로 결과 타입을 확인한다.

```ts
let identity = function<T>(x: T): T {
  // ...
};
let reverse = function<U>(y: U): U {
  // ...
};

identity = reverse; // OK, because (x: any) => any matches (y: any) => any
```

### Subtype vs Assignment

TypeScript Spec에는 subtype과 assignment는 두 가지 종류의 호환성이 있다.

두 호환성은 할당이 하위타입 규칙 호환성(`any` 혹은 숫자값에 대응하는 `enum`에 할당되거나 할당을 허용하는)을 확장할 때만 다르다.

## 고급 타입

### Intersection Types (`&`)

intersection 타입은 여러 타입을 하나로 결합한다.
intersection 타입은 주로 믹스인이나 고전적인 객체지향 틀에서 벗어난 개념에 사용되는 것을 볼 수 있다.

```ts
function extend<First, Second>(first: First, second: Second): First & Second {
  const result: Partial<First & Second> = {};
  for (const prop in first) {
    if (first.hasOwnProperty(prop)) {
      (<First>result)[prop] = first[prop];
    }
  }
  for (const prop in second) {
    if (second.hasOwnProperty(prop)) {
      (<Second>result)[prop] = second[prop];
    }
  }
  return <First & Second>result;
}

class Person {
  constructor(public name: string) {}
}

interface Loggable {
  log(name: string): void;
}

class ConsoleLogger implements Loggable {
  log(name) {
    console.log(`Hello, I'm ${name}.`);
  }
}

const jim = extend(new Person("Jim"), ConsoleLogger.prototype);
jim.log(jim.name);
```

### Union Types (`|`)

union 타입은 여러 타입중 하나가 될 수 있는 타입을 설명한다.
만약 union 타입을 가진 값이 있다면 union 타입에서 공통적으로 존재하는 멤버들만 접근할 수 있다.

```ts
interface Bird {
  fly();
  layEggs();
}
interface Fish {
  swim();
  layEggs();
}

function getSmallPet(): Fish | Bird {
  // ...
}

let pet = getSmallPet();
pet.layEggs(); // okay
pet.swim(); // errors
```

### Type Guards and Differentiating Types

union 타입은 공통으로 존재하는 멤버에만 접근 가능하므로,
존재하지만 타입 때문에 접근할 수 없는 멤버에 접근하기 위해서는 type assertion을 활용할 것이다.

```ts
let pet = getSmallPet();

if ((<Fish>pet).swim) {
  (<Fish>pet).swim();
} else {
  (<Bird>pet).fly();
}
```

#### User-Defined Type Guards

type assertion을 여러번 사용하지 않고 타입 확인을 한 번만 할 수 있다.
TypeScript의 Type Guard는 타입이 특정 범위 내에 있음을 보장하는 runtime check를 수행하는 표현식이다.

타입 가드를 정의하기 위해서는 반환 타입을 타입 서술로 표기하면 된다. (`parameterName is Type`)

```ts
function isFish(pet: Fish | Bird): pet is Fish {
  return (<Fish>pet).swim !== undefined;
}

if (isFish(pet)) {
  pet.swim();
} else {
  pet.fly();
}
```

TypeScript 컴파일러는 조건문에서 `Fish`임을 확인한다면 다른 경우는 `Bird` 타입임을 알고 있다.

#### `typeof` type guards

typename은 다음의 "number", "string", "boolean", or "symbol" 중의 하나를 사용해야 하며,
`typeof v === "typename"` 또는 `typeof v !== "typename"`의 두 가지 형태로 사용할 수 있다.

```ts
function padLeft(value: string, padding: string | number) {
  if (typeof padding === "number") {
    return Array(padding + 1).join(" ") + value;
  }
  if (typeof padding === "string") {
    return padding + value;
  }
  throw new Error(`Expected string or number, got '${padding}'.`);
}
```

#### `instanceof` type guards

`instanceof` 타입 가드는 생성성자 삼수를 사용하여 타입을 한정하는 방법이다.

`instanceof`의 오른쪽은 생성자 함수여야 하며, TypeScript는 다음 순서로 범위를 한정한다.

1. 함수의 prototype 프로퍼티의 타입이 `any`가 아닌 경우, 그 타입
2. 타입의 생성자 시그니처에의해 반환되는 union type

```ts
interface Padder {
  getPaddingString(): string;
}

class SpaceRepeatingPadder implements Padder {
  constructor(private numSpaces: number) {}
  getPaddingString() {
    return Array(this.numSpaces + 1).join(" ");
  }
}

class StringPadder implements Padder {
  constructor(private value: string) {}
  getPaddingString() {
    return this.value;
  }
}

function getRandomPadder() {
  return Math.random() < 0.5
    ? new SpaceRepeatingPadder(4)
    : new StringPadder("  ");
}

// Type is 'SpaceRepeatingPadder | StringPadder'
let padder: Padder = getRandomPadder();

if (padder instanceof SpaceRepeatingPadder) {
  padder; // type narrowed to 'SpaceRepeatingPadder'
}
if (padder instanceof StringPadder) {
  padder; // type narrowed to 'StringPadder'
}
```

### Nullable types

TypeScript에는 `null` 및 `undefined` 값을 갖는 두 특수한 타입이 있다.
기본적으로 타입체커는 `null`과 `undefined`는 모든 타입의 값으로 유효한 것으로 간주한다.

`--strictNullChecks` 플래그를 사용하면 변수를 선언할 때 `null` 또는 `undefined`를 자동으로 포함하지 않는다.
Union 타입을 사용해서 명시적으로 포함할 수는 있다.

```ts
let s = "foo";
s = null; // error, 'null' is not assignable to 'string'
let sn: string | null = "bar";
sn = null; // ok

sn = undefined; // error, 'undefined' is not assignable to 'string | null'
```

TypeScript는 JavaScript 처럼 `null`과 `undefined`를 다르게 처리한다.

#### Optional parameters and properties

`--strictNullChecks`를 사용하면 선택적 매개변수 및 선택적 프로퍼티에 `| undefined`가 추가된다

```ts
function f(x: number, y?: number) {
  return x + (y || 0);
}
f(1);
f(1, undefined);
f(1, null); // error, 'null' is not assignable to 'number | undefined'

class C {
  a: number;
  b?: number;
}
let c = new C();
c.a = undefined; // error, 'undefined' is not assignable to 'number'
c.b = undefined; // ok
c.b = null; // error, 'null' is not assignable to 'number | undefined'
```

#### Type guards and type assertions

nullable 타입은 union type으로 구현되므로 타입 가드를 사용해서 null을 제거해야 한다

```ts
function f(sn: string | null): string {
  if (sn == null) {
    return "default";
  } else {
    return sn;
  }
}
```

null 제거를 위해서 간결하게 연산자를 사용할 수 있다

```ts
function f(sn: string | null): string {
  return sn || "default";
}
```

컴파일러가 null 또는 undefined를 제거할 수 없는 경우 type assertion 연산자를 사용해서 수동으로 제거할 수 있다

null 제거를 위한 연산자는 postfix `!`이다

```ts
function fixed(name: string | null): string {
  function postfix(epithet: string) {
    return name!.charAt(0) + ". the " + epithet; // ok
  }
  name = name || "Bob";
  return postfix("great");
}
```

### Type Aliases

Type Alias는 타입의 새로운 이름을 생성한다.

인터페이스와 유사해보이지만 원시 타입 및 union, tuple 및 기타 여러 타입을 대상으로 이름을 지정할 수 있다.

```ts
type Name = string;
type NameResolver = () => string;
type NameOrResolver = Name | NameResolver;
function getName(n: NameOrResolver): Name {
  if (typeof n === "string") {
    return n;
  } else {
    return n();
  }
}
```

Aliasing은 새 타입을 생성하지 않으며 해당 타입을 가리키는 이름을 만드는 것이다.
원시타입에 대한 alias를 작성하는 것은 전혀 쓸모가 없다.

인터페이스와 마찬가지로 type alias에 generic도 사용 가능하다.

```ts
type Container<T> = { value: T };
```

또한 type alias를 정의할 때 프로퍼티에서 해당 alias를 바로 사용할 수 있다.

```ts
type Tree<T> = {
  value: T;
  left: Tree<T>;
  right: Tree<T>;
};
```

intersection 타입과 함께 사용해서 상당히 혼란스러운 타입을 만들 수도 있다

```ts
type LinkedList<T> = T & { next: LinkedList<T> };

interface Person {
  name: string;
}

var people: LinkedList<Person>;
var s = people.name;
var s = people.next.name;
var s = people.next.next.name;
var s = people.next.next.next.name;
```

그러나 별칭은 선언 우측에서 바로 사용될 수 없다

```ts
type Yikes = Array<Yikes>; // error
```

#### Interfaces vs Type Aliases

타입 별칭과 다르게 인터페이스는 어디에서나 사용될 수 있는 이름을 생성한다.
아래의 코드에서 `aliased`는 객체 리터럴 타입을 반환하지만 `interfaced`는 인터페이스를 반환한다.

```ts
type Alias = { num: number };
interface Interface {
  num: number;
}
declare function aliased(arg: Alias): Alias;
declare function interfaced(arg: Interface): Interface;
```

또한, t타입 별칭을 확장(extend)하거나 구현(implement)할 수 없으며, 다른 타입으로부터 확장/구현할 수도 없다.

소프트웨어는 개방-폐쇄원칙을 지키는 것이 이상적이므로, 확장이 필요한 경우 인터페이스를 사용해야 한다.
반면, 인터페이스를 통해서 형태를 표현할 수 없고 union 타입이나 튜플이 필요한 경우 타입 별칭이 사용된다.

### String Literal Types

문자열 리터럴 타입을 사용해서 문자열에 있어야 하는 정확한 값을 지정할 수 있다.

문자열 리터럴 타입은 union type, 타입 가드, 타입 별칭과 함께 사용하기 좋다.
그러한 기능을 함께 사용해서 문자열에서 enum과 비슷한 결과를 얻을 수 있다.

```ts
type Easing = "ease-in" | "ease-out" | "ease-in-out";
class UIElement {
  animate(dx: number, dy: number, easing: Easing) {
    if (easing === "ease-in") {
      // ...
    } else if (easing === "ease-out") {
    } else if (easing === "ease-in-out") {
    } else {
      // error! should not pass null or undefined.
    }
  }
}

let button = new UIElement();
button.animate(0, 0, "ease-in");
button.animate(0, 0, "uneasy"); // ERROR: Argument of type '"uneasy"' is not assignable to parameter of type '"ease-in" | "ease-out" | "ease-in-out"'
```

overloading을 구분하기 위해서 문자열 리터럴 타입을 비슷한 방식으로 사용할 수 있다

```ts
function createElement(tagName: "img"): HTMLImageElement;
function createElement(tagName: "input"): HTMLInputElement;
// ... more overloads ...
function createElement(tagName: string): Element {
  // ... code goes here ...
}
```

### Numeric Literal Types

TypeScript에는 숫자 리터럴 타입도 있다

```ts
function rollDice(): 1 | 2 | 3 | 4 | 5 | 6 {
  // ...
}
```

숫자 리터럴을 명시적으로 사용하는 경우는 거의 없으며, 버그를 잡기 위해 범위를 좁히는데 유용하다.

```ts
function foo(x: number) {
  if (x !== 1 || x !== 2) {
    // ERROR: Operator '!==' cannot be applied to types '1' and '2'.
    // ...
  }
}
```

x가 타입 `2`와 비교될 때는 이미 타입 `1`인 경우이므로 오류 발생

### Enum Member Types

Enum 멤버는 모든 멤버가 리터럴로 초기화 될 때 타입이된다.

싱글톤 타입을 말할 때 숫자/문자 리터럴 타입뿐만 아니라 enum 멤버 타입 둘다를 가리키는 것이다.
그러나 많은 사람들은 싱글톤 타입과 리터럴 타입을 혼용하고 있다.

### Discriminated Unions

싱글톤 타입, union 타입, 타입 가드, 타입 별칭을 결합해서 discriminated unions라고 불리는 고급 패턴을 만들 수 있다.

tagged union / algebraic data types로도 불리는 discriminated unions은 함수형 프로그래밍에 유용하다.

1. discriminant: 공통, 싱글턴 타입 프로퍼티를 갖는 타입
2. union: 해당 타입들의 union을 갖는 타입 별칭
3. 공통 프로퍼티에 타입가드 존재

```ts
interface Square {
  kind: "square";
  size: number;
}
interface Rectangle {
  kind: "rectangle";
  width: number;
  height: number;
}
interface Circle {
  kind: "circle";
  radius: number;
}

type Shape = Square | Rectangle | Circle;
```

결합하기 위해 선언한 인터페이스에는 서로 다른 문자열 리터럴 타입을 가진 `kind` 프로퍼티가 있다.
`kind` 프로퍼티를 discriminant 또는 tag라고 한다.

discriminated union은 다음과 같이 사용한다.

```ts
function area(s: Shape) {
  switch (s.kind) {
    case "square":
      return s.size * s.size;
    case "rectangle":
      return s.height * s.width;
    case "circle":
      return Math.PI * s.radius ** 2;
  }
}
```

#### Exhaustiveness checking

모든 discriminated union 변형을 커버하지 않으려는 것을 컴파일러에게 알려주고 싶을 때가 있다.

```ts
type Shape = Square | Rectangle | Circle | Triangle;
function area(s: Shape) {
  switch (s.kind) {
    case "square":
      return s.size * s.size;
    case "rectangle":
      return s.height * s.width;
    case "circle":
      return Math.PI * s.radius ** 2;
  }
  // should error here - we didn't handle case "triangle"
}
```

이를 위한 방법은 두 가지가 있다

- 컴파일러의 `--strictNullChecks` 옵션을 켜고 반환타입을 정의 (`number | undefined`)
- `never` 타입 사용

### Polymorphic this types

`this` 타입의 다형성은 포함하고 있는 클래스나 인터페이스의 서브타입을 나타낸다.

이는 F-bounded 다형성이라고 하며, 계층적 인터페이스를 쉽게 표현할 수 있게 한다.

```ts
class BasicCalculator {
  public constructor(protected value: number = 0) {}
  public currentValue(): number {
    return this.value;
  }
  public add(operand: number): this {
    this.value += operand;
    return this;
  }
  public multiply(operand: number): this {
    this.value *= operand;
    return this;
  }
  // ... other operations go here ...
}

let v = new BasicCalculator(2)
  .multiply(5)
  .add(1)
  .currentValue();
```

`this` 타입을 사용하기 때문에 변경점 없이 이전 메소드를 사용하는 새 클래스를 만들 수 있다.

```ts
class ScientificCalculator extends BasicCalculator {
  public constructor(value = 0) {
    super(value);
  }
  public sin() {
    this.value = Math.sin(this.value);
    return this;
  }
  // ... other operations go here ...
}

let v = new ScientificCalculator(2)
  .multiply(5)
  .sin()
  .add(1)
  .currentValue();
```

`this` 타입을 사용하지 않았다면 `ScientificCalculator`에서 `sin`메소드가 없는 `BasicCalculator` 타입이 반환되었을 것이다.

### Index types

인덱스 타입을 사용하면 컴파일러가 동적 프로퍼티를 검사하도록 할 수 있다.

```ts
function pluck<T, K extends keyof T>(o: T, names: K[]): T[K][] {
  return names.map(n => o[n]);
}

interface Person {
  name: string;
  age: number;
}
let person: Person = {
  name: "Jarid",
  age: 35
};
let strings: string[] = pluck(person, ["name"]); // ok, string[]
```

컴파일러는 `name`이 실제로 `Person`의 프로퍼티인지 확인한다.

**인덱스 타입 쿼리 연산자**인 `keyof T`는 public인 프로퍼티 이름의 union type이다.

즉 `let personProps: keyof Person; // 'name' | 'age'`

`keyof Person`은 `'name' | 'age'`와 완벽하게 호환된다.

**indexed access operator** `T[K]`는 `person['name']`이 `Person['name']` 타입을 가진다는 것을 의미한다.

인덱스 타입 쿼리와 마찬가지로 `T[K]`를 제네릭 컨텍스트에서 사용할 수 있다.
이를 위해 `K extends keyof T`임을 확인해야 한다.

```ts
function getProperty<T, K extends keyof T>(o: T, name: K): T[K] {
  return o[name]; // o[name] is of type T[K]
}
```

`getProperty`에서 `o[name]: T[K]`의 관계가 된다.
`T[K]`의 결과를 반환하면 컴파일러는 키의 실제 타입을 인스턴스화 하므로 `getProperty`의 반환타입은 요청한 프로퍼티에 따라 다양해진다.

```ts
let name: string = getProperty(person, "name");
let age: number = getProperty(person, "age");
let unknown = getProperty(person, "unknown"); // error, 'unknown' is not in 'name' | 'age'
```

#### Index types and string index signatures

`keyof`와 `T[K]`는 문자열 인덱스 시그니처와 상호작용한다.
만약 문자열 인덱스 시그니처가 있다면 `keyof T`는 문자열이 될 것이고 `T[string]`은 인덱스 시그니처 타입이 될 것이다.

```ts
interface Dictionary<T> {
  [key: string]: T;
}
let keys: keyof Dictionary<number>; // string
let value: Dictionary<number>["foo"]; // number
```

### Mapped types

TypeScript는 이전의 타입을 기반으로 새로운 타입을 만드는 방법인 mapped types를 제공한다.

```ts
type Partial<T> = { [P in keyof T]?: T[P] };
type Readonly<T> = { readonly [P in keyof T]: T[P] };
```

다음과 같이 사용한다

```ts
type PersonPartial = Partial<Person>;
type ReadonlyPerson = Readonly<Person>;
```

만약 멤버를 추가하고 싶다면, intersection type을 사용해야 한다

```ts
// Use this:
type PartialWithNewMember<T> = {
  [P in keyof T]?: T[P];
} & { newMember: boolean }

// **Do not** use the following!
// This is an error!
type PartialWithNewMember<T> = {
  [P in keyof T]?: T[P];
  newMember: boolean;
}
```

간단한 mapped types 예제를 보자

```ts
type Keys = "option1" | "option2";
type Flags = { [K in Keys]: boolean };
```

1. 타입변수 `K`는 각 프로퍼티에 순서대로 바인딩 됨
2. 문자열 리터럴 union인 `Keys`는 반복할 속성의 이름을 포함한다

이 예제에서 Keys는 하드코딩된 프로퍼티 이름의 목록이고 프로퍼티 타입은 항상 boolean이다.
따라서 이러한 mapped types는 다음과 같을 것이다.

```ts
type Flags = {
  option1: boolean;
  option2: boolean;
};
```

실제 적용은 위의 `Readonly` 혹은 `Partial`과 같은 형태이다.
기존의 타입이 있고, mapped types는 기존 타입의 프로퍼티를 변형할 것이다.
그곳에 `keyof` 인덱스 접근 타입이 위치한다.

```ts
type Nullable<T> = { [P in keyof T]: T[P] | null };
type Partial<T> = { [P in keyof T]?: T[P] };
```

이러한 예제에서 프로퍼티 목록은 `keyof T`이고 결과값의 타입은 `T[P]`의 변형이다.
이러한 종류의 변형은 homomorphic이므로 mapping은 `T`타입의 프로퍼티에만 적용된다.

컴파일러는 새로운 프로퍼티를 추가하기 전에 기존의 프로퍼티 수정자를 가져와 적용한다.
예를 들어, `Person.name`이 readonly 였다면, `Partial<Pserson>.name`은 readonly 이며 optional이다.

다음은 `T[P]` 타입이 `Proxy<T>` 클래스로 wrapped 되는 예제이다

```ts
type Proxy<T> = {
  get(): T;
  set(value: T): void;
};
type Proxify<T> = { [P in keyof T]: Proxy<T[P]> };
function proxify<T>(o: T): Proxify<T> {
  // ... wrap proxies ...
}
let proxyProps = proxify(props);
```

`Readonly<T>`와 `Partial<T>`는 유용하기 때문에 `Pick` 및 `Record`와 함께 타입스크립트 표준 라이브러리에 포함되어 있다.

```ts
type Pick<T, K extends keyof T> = { [P in K]: T[P] };
type Record<K extends keyof any, T> = { [P in K]: T };
```

`Readonly`, `Partial`, `Pick`은 homomorphic이지만 `Record`는 그렇지 않다.

`Record`가 homomorphic하지 않다는 것은 속성을 복사할 때 입력받은 타입을 사용하지 않는 점에서 알 수 있다.

```ts
type ThreeStringProps = Record<"prop1" | "prop2" | "prop3", string>;
```

non-homomorphic 타입은 본질적으로 새로운 속성을 생성하기 때문에 프로퍼티 수정자를 복사할 수 없다.

#### Inference from mapped types

반대로 타입의 프로퍼티를 unwrap하는 방법을 알아볼 차례이다

```ts
function unproxify<T>(t: Proxify<T>): T {
  let result = {} as T;
  for (const k in t) {
    result[k] = t[k].get();
  }
  return result;
}

let originalProps = unproxify(proxyProps);
```

unwrapping 추론은 homomorphic mapped types에서만 작동한다.
non-homomorphic mapped types인 경우 unwapping 함수에 명시적 타입 파라미터를 전달해야 한다.

### Conditional Types

타입스크립트 2.8에서 통일되지않은 타입 매핑을 표현하기 위한 기능인 조건부 타입이 추가되었다.
조건부 타입은 조건 표현식에 따라 두 유형 중 하나를 선택하게 된다.

아래의 의미는 `T`가 `U`에 할당가능할 때 타입이 `X`이고 아니면 `Y`이다.

```ts
T extends U ? X : Y
```

위의 타입은 `X` 또는 `Y`로 평가되거나, 평가가 지연되는데
이는 타입시스템에서 `T`가 항상 `U`에 할당가능하다고 결론을 내리기에 충분한 정보가 있는지 여부에 따라 결정된다.

다음은 타입이 즉시 평가되는 예제이다

```ts
declare function f<T extends boolean>(x: T): T extends true ? string : number;

// Type is 'string | number
let x = f(Math.random() < 0.5);
```

다른 예제는 중첩 조건부 타입을 사용하는 예제이다

```ts
type TypeName<T> = T extends string
  ? "string"
  : T extends number
  ? "number"
  : T extends boolean
  ? "boolean"
  : T extends undefined
  ? "undefined"
  : T extends Function
  ? "function"
  : "object";

type T0 = TypeName<string>; // "string"
type T1 = TypeName<"a">; // "string"
type T2 = TypeName<true>; // "boolean"
type T3 = TypeName<() => void>; // "function"
type T4 = TypeName<string[]>; // "object"
```

조건부 타입 평가가 지연되는 예

```ts
interface Foo {
  propA: boolean;
  propB: boolean;
}

declare function f<T>(x: T): T extends Foo ? string : number;

function foo<U>(x: U) {
  // Has type 'U extends Foo ? string : number'
  let a = f(x);

  // This assignment is allowed though!
  let b: string | number = a;
}
```

위의 변수 `a`는 아직 분기가 선택되지 않은 조건부 타입이다.
다른 코드가 `foo`를 호출하면 `U`에서 다른 타입으로 대체되고, 조건부 타입은 다시 평가되어 분기 선택여부를 결정하게 된다.

#### Distributive conditional types

checked type이 naked type 파라미터인 조건부 타입을 distributive conditional types(DCT)라 한다.

DCT는 인스턴스화 할 때 자동으로 union types에 분배된다.

예를 들어 `T`의 타입 인자가 `A | B | C`인 `T extends U ? X : Y`의 인스턴스화는
`(A extends U ? X : Y) | (B extends U ? X : Y) | (C extends U ? X : Y)`의 형태로 이루어진다.

```ts
type T10 = TypeName<string | (() => void)>; // "string" | "function"
type T12 = TypeName<string | string[] | undefined>; // "string" | "object" | "undefined"
type T11 = TypeName<string[] | number[]>; // "object"
```

DCT `T extends U ? X : Y`의 인스턴스화는, 조건부 타입들이 union type의 개별 구성요소로 해석되는 `T`를 참조한다.
(i.e `T`는 조건부 타입이 union 타입에 분산된 후 개별 구성요소를 참조한다)

또한, `X`내의 `T`에 대한 참조는 추가적인 타입 파라미터 제약 `U`를 갖는다.
(i.e `T`는 `X`내에서 `U`에 할당 가능하다고 간주된다)

```ts
type BoxedValue<T> = { value: T };
type BoxedArray<T> = { array: T[] };
type Boxed<T> = T extends any[] ? BoxedArray<T[number]> : BoxedValue<T>;

type T20 = Boxed<string>; // BoxedValue<string>;
type T21 = Boxed<number[]>; // BoxedArray<number>;
type T22 = Boxed<string | number[]>; // BoxedValue<string> | BoxedArray<number>;
```

`T`는 `Boxed<T>` 분기 내에서 추가적인 제약 `any[]`를 가질 수 있으므로, 배열의 요소 타입을 `T[number]`로 참조할 수 있다.

조건부 타입의 분산된 프로퍼티는 union 타입을 필터링 하는데 유용하게 사용된다.

```ts
type Diff<T, U> = T extends U ? never : T; // Remove types from T that are assignable to U
type Filter<T, U> = T extends U ? T : never; // Remove types from T that are not assignable to U

type T30 = Diff<"a" | "b" | "c" | "d", "a" | "c" | "f">; // "b" | "d"
type T31 = Filter<"a" | "b" | "c" | "d", "a" | "c" | "f">; // "a" | "c"
type T32 = Diff<string | number | (() => void), Function>; // string | number
type T33 = Filter<string | number | (() => void), Function>; // () => void

type NonNullable<T> = Diff<T, null | undefined>; // Remove null and undefined from T

type T34 = NonNullable<string | number | undefined>; // string | number
type T35 = NonNullable<string | string[] | null | undefined>; // string | string[]

function f1<T>(x: T, y: NonNullable<T>) {
  x = y; // Ok
  y = x; // Error
}

function f2<T extends string | undefined>(x: T, y: NonNullable<T>) {
  x = y; // Ok
  y = x; // Error
  let s1: string = x; // Error
  let s2: string = y; // Ok
}
```

조건부 타입은 mapped types와 결합할 때 특히 유용하다

```ts
type FunctionPropertyNames<T> = {
  [K in keyof T]: T[K] extends Function ? K : never
}[keyof T];
type FunctionProperties<T> = Pick<T, FunctionPropertyNames<T>>;

type NonFunctionPropertyNames<T> = {
  [K in keyof T]: T[K] extends Function ? never : K
}[keyof T];
type NonFunctionProperties<T> = Pick<T, NonFunctionPropertyNames<T>>;

interface Part {
  id: number;
  name: string;
  subparts: Part[];
  updatePart(newName: string): void;
}

type T40 = FunctionPropertyNames<Part>; // "updatePart"
type T41 = NonFunctionPropertyNames<Part>; // "id" | "name" | "subparts"
type T42 = FunctionProperties<Part>; // { updatePart(newName: string): void }
type T43 = NonFunctionProperties<Part>; // { id: number, name: string, subparts: Part[] }
```

union 및 intersection 타입과 비슷하게 조건부 타입은 재귀적으로 스스로를 참조할 수 없다.

```ts
type ElementType<T> = T extends any[] ? ElementType<T[number]> : T; // Error
```

#### Type inference in conditional types

조건부 타입의 `extends`절 내에서, `infer` 선언으로 타입 변수가 추론될 수 있음을 알릴 수 있다.
이러한 추론된 타입 변수는 조건부 타입에서 true 분기에서 참조된다.

같은 타입 변수에서 여러개의 `infer` 선언을 가질 수 있다.

```ts
type ReturnType<T> = T extends (...args: any[]) => infer R ? R : any;
```

조건부 타입은 순서대로 평가되는 연속적인 패턴 매칭을 형성하기 위해 중첩될 수 있다

```ts
type Unpacked<T> = T extends (infer U)[]
  ? U
  : T extends (...args: any[]) => infer U
  ? U
  : T extends Promise<infer U>
  ? U
  : T;

type T0 = Unpacked<string>; // string
type T1 = Unpacked<string[]>; // string
type T2 = Unpacked<() => string>; // string
type T3 = Unpacked<Promise<string>>; // string
type T4 = Unpacked<Promise<string>[]>; // Promise<string>
type T5 = Unpacked<Unpacked<Promise<string>[]>>; // string
```

다음 예제는 얼마나 많은 동일 타입 변수의 후보가 공변(co-variant)의 위치에서 추론될 수 있는 union 타입을 형성하는지를 보여준다.

```ts
type Foo<T> = T extends { a: infer U; b: infer U } ? U : never;
type T10 = Foo<{ a: string; b: string }>; // string
type T11 = Foo<{ a: string; b: number }>; // string | number
```

마찬가지로, 동일 타입 변수들의 후보가 반공변(contra-variant)의 위치에서 추론될 수 있는 intersection 타입을 형성하는지 보여준다.

```ts
type Bar<T> = T extends { a: (x: infer U) => void; b: (x: infer U) => void }
  ? U
  : never;
type T20 = Bar<{ a: (x: string) => void; b: (x: string) => void }>; // string
type T21 = Bar<{ a: (x: string) => void; b: (x: number) => void }>; // string & number
```

다수의 호출 시그니처를 가진 타입(오버로드된 함수와 같은...)으로 부터 추론은 마지막 시그니처로 부터 이루어진다.

```ts
declare function foo(x: string): number;
declare function foo(x: number): string;
declare function foo(x: string | number): string | number;
type T30 = ReturnType<typeof foo>; // string | number
```

일반적인 타입 파라미터의 제약조건절에서 `infer` 선언을 할 수 없다

```ts
type ReturnType<T extends (...args: any[]) => infer R> = R; // Error, not supported
```

그러나 제약조건에서 타입변수를 지우고 대신 조건부 타입을 지정하면 동일한 효과를 얻을 수 있다

```ts
type AnyFunction = (...args: any[]) => any;
type ReturnType<T extends AnyFunction> = T extends (...args: any[]) => infer R
  ? R
  : any;
```

## Symbols

### Introduction

ES6부터 `Symbol`은 JavaScript의 primitive 타입이다.

symbol 값은 `Symbol` 생성잘르 호출해서 만들 수 있다.

```ts
let sym1 = Symbol();
let sym2 = Symbol("key"); // optional string key
```

Symbol은 불변이며 유일한 값을 갖는다

```ts
let sym2 = Symbol("key");
let sym3 = Symbol("key");
sym2 === sym3; // false, symbols are unique
```

문자열과 마찬가지로 symbols는 오브젝트 프로퍼티의 키로 사용될 수 있다

```ts
const sym = Symbol();
let obj = {
  [sym]: "value";
};
console.log(obj[sym]); // "value"
```

Symbol은 계산된 프로퍼티 선언과 결합하여 오브젝트 프로퍼티와 클래스 멤버를 선언할 수 있다

```ts
const getClassNameSymbol = Symbol();

class C {
  [getClassNameSymbol]() {
    return "C";
  }
}

let c = new C();
let className = c[getClassNameSymbol](); // "C"
```

### Well-known Symbols

사용자 정의 심볼외에도 잘 알려진 내장 심볼들이 있다. 내장 심볼들은 언어 내부 동작을 나타내는데 사용된다.

다음은 잘 알려진 심볼의 목록이다

#### `Symbol.hasInstance`

생성자로 생성된 객체가 어떤 생성자의 인스턴스 중 하나인지 결정한다.
`instanceof` 연산자에 의해 호출된다. (`instanceof`를 확장해서 사용한다)

#### `Symbol.isConcatSpreadable`

`Array.prototype.concat`에 의해 객체가 배열 요소로 flatten 될수 있는지 여부를 나타내는 boolean 값이다

#### `Symbol.iterator`

객체의 기본 반복자를 반환하는 메소드이다.
`for-of`문에 의해서 호출된다.

#### `Symbol.match`

문자열을 정규표현식과 비교하는 졍규표현식 메소드이다.
`String.prototype.match` 메소드에 의해 호출된다.

#### `Symbol.replace`

일치하는 문자열의 하위 문자열을 대체하는 정규표현식 메소드이다.
`String.prototype.replace` 메소드에 의해 호출된다.

#### `Symbol.search`

정규표현식과 일치하는 문자열의 index를 반환하는 정규표현식 메소드이다.
`String.prototype.search` 메소드에 의해 호출된다.

#### `Symbol.species`

파생 객체를 만드는 생성자 함수의 함수값 프로퍼티이다

#### `Symbol.split`

정규표현식과 일치하는 indices에서 문자열을 분리하는 정규표현식 메소드이다.
`String.prototype.split` 메소드에 의해 호출된다.

#### `Symbol.toPrimitive`

객체를 대응하는 primitive value로 변환하는 메소드이다.
`ToPrimitive` abstract operation에서 호출된다.

#### `Symbol.toStringTag`

객체의 기본 문자열 description을 생성하는데 사용되는 문자열 값이다.
내장메소드인 `Object.prototype.toString`에메소드에 의해 호출된다.

#### `Symbol.unscopables`

특정 메소드들이 동적 스코핑에 관여되는 것을 방지함

## Iterators and Generators

객체에 `Symbol.iterator` 프로퍼티 구현이 있으면 iterable로 간주된다.
`Array`, `Map`, `Set`, `String`, `Int32Array`, `Uint32Array` 같은 built-in 타입들이 iterable이다.

객체의 `Symbol.iterator` 함수는 반복 할 값 목록을 반환한다.

### `for..of` statements

`for..of` 루프는 `Symbol.iterator` 프로퍼티를 호출하여 객체를 순회한다.

```ts
let someArray = [1, "string", false];
for (let entry of someArray) {
  console.log(entry); // 1, "string", false
}
```

### `for..of` vs `for..in` statements

`for..of`, `for..in`문 모두 list를 순회한다.
다만, `for..of`는 property 값을 반환하고 `for..in`은 객체의 key를 반환한다.

```ts
let list = [4, 5, 6];
for (let i in list) {
  console.log(i); // "0", "1", "2",
}
for (let i of list) {
  console.log(i); // "4", "5", "6"
}
```

또 다른점은 `for.in`은 모든 객체에서 작동한다는 것이다. 반면 `for..of`는 iterable만 반복한다.

```ts
let pets = new Set(["Cat", "Dog", "Hamster"]);
pets["species"] = "mammals";

for (let pet in pets) {
  console.log(pet); // "species"
}
for (let pet of pets) {
  console.log(pet); // "Cat", "Dog", "Hamster"
}
```

### 코드생성

#### ES5 / ES3

이 경우 iterator는 `Array` 타입에만 허용된다.

```ts
let numbers = [1, 2, 3];
for (let num of numbers) {
  console.log(num);
}
```

위 코드는 다음처럼 생성될 것이다

```ts
var numbers = [1, 2, 3];
for (var _i = 0; _i < numbers.length; _i++) {
  var num = numbers[_i];
  console.log(num);
}
```

# TypeScript

## 기본 타입

- 부울: `boolean`: `true` | `false`

- 숫자: `number`: 자바스크립트와 마찬가지로 64비트 부동 소수점 값이다. 16/2/8 진수를 할당할 수 있다

  ```ts
  let decimal: number = 6;
  let hex: number = 0xf00d;
  let binary: number = 0b1010;
  let octal: number = 0o744;
  ```

- 문자: `string`: 큰 따옴표, 작은 따옴표, 템플릿 문자열을 위한 백 쿼트를 사용할 수 있다

- 배열: `T[]`, `Array<T>`: 두 가지 방식으로 선언할 수 있다

- 튜플: `[T, U]`: 고정된 개수의 요소와 타입(같을 필요 없음)을 표현한다

- 열거: `enum T {A, B, C}`: enumeration 타입의 요소는 순서대로 0부터 시작하는 키값을 갖는다

- `any`: 알지 못하는 변수 타입 (최상위 타입으로 쓸 수도 있다)

- `void`
  - `undefined` 또는 `null`만 할당할 수 있다
  - 일반적으로 반환이 없는 함수의 반환타입으로 사용됨
  - 변수 선언시 타입으로 `void`는 유용하지 않다

- `undefined` / `null`
  - 다른 모든 타입의 서브 타입이다
  - `void`타입과 마찬가지로 변수의 타입으로 사용하기는 적절하지 않다
  - 컴파일러의 `--strictNullChecks` 옵션을 켜면 `void` 타입 혹은 각자의 타입인 `undefined` 또는 `null`에만 할당 가능하다
  - `undefined` vs `null`
    - `undefined`: primitive value used when a variable has **not been assigned a value** (ECMA)
    - `null`: primitive value that represents the intentional **absence of any object value** (ECMA)
    - TS coding guidelines of MS: [Use undefined. Do not use null](https://github.com/Microsoft/TypeScript/wiki/Coding-guidelines#null-and-undefined)
      - 명시적으로 nullable variable|property 할당하지 않아도 됨
      - runtime 타입검사에 유리하다
        - `typeof val === 'object'` for both null and object value
        - `typeof val === 'undefined'` for undefined.

- `never`
  - 다른 모든 타입의 서브 타입이다 하지만 `never`타입의 서브타입은 없으며 어떠한 타입도 `never`타입에 위치에 할당할 수 없다
  - 절대로 발생하지 않는 값의 타입이다
    - 항상 `Error`를 반환하는 함수(`throw new Error()` | `return error`)
    - end point에 도달하지 않는 함수(무한루프 ...)

- `object`
  - primitive 타입이 아닌 모든타입
  - `number`, `string`, `boolean`, `symbol`, `null`, or `undefined` 이외의 타입

### Type assertions

다음과 같이 `as T`의 형식으로 사용되며, 컴파일러에게 해당 타입을 알려주는 역할을 한다.

```ts
const someValue: any = "this is a string";
const strLength: number = (someValue as string).length;
```

## 변수

### var

전통적인 JavaScript 변수선언으로 var의 스코프 규칙은 함수이다

```ts
function f(shouldInitialize: boolean) {
  if (shouldInitialize) {
    var x = 10;
  }
  return x;
}

f(true);  // 10
f(false); // undefined
```

### let

let은 var와 달리 lexical-scope 규칙을 갖는다

```ts
function f(input: boolean) {
  let a = 100;
  if (input) {
    let b = a + 1;
    return b;
  }
  return b; // 오류 발생: b는 존재하지 않음
}
```

let은 루프의 일부로 선언될 때 반복마다 새로운 스코프를 생성한다.
반복 횟수가 많은 경우 상대적으로 성능 저하를 불러올 수 있다.

```ts
for (let i = 0; i < 10 ; i++) {
  setTimeout(function() { console.log(i); }, 100 * i);
}
```

결과는 다음과 같다

```text
0
1
2
3
4
5
6
7
8
9
```

### const

let과 거의 동일한 규칙을 갖고 있으나 재할당할 수 없다.(불변은 아님)

### 전개 연산자 (spread)

```ts
let first = [1, 2];
let second = [3, 4];
let bothPlus = [0, ...first, ...second, 5];
console.log(bothPlus); // [0, 1, 2, 3, 4, 5]
```

객체에도 적용 가능하지만 중복 항목이 있다면 나중에 등장한 객체가 이전의 프로퍼티를 덮어쓴다

```ts
let defaults = { food: "spicy", price: "$$", ambiance: "noisy" };
let search = { ...defaults, food: "rich" };
```

객체의 인스턴스를 전개할 때 메소드는 사라진다

```ts
class C {
  p = 12;
  m() {
  }
}
let c = new C();
let clone = { ...c };
clone.p; // 12
clone.m(); // Error!
```

### Array destructuring

```ts
let input = [1, 2];
let [first, second] = input;
console.log(first); // 1
console.log(second); // 2
```

변수 교환도 가능하다

```ts
[first, second] = [second, first];
```

전개 연산자로 나머지 항목도 가져올 수 있다

```ts
let [first, ...rest] = [1, 2, 3, 4];
console.log(first); // 1
console.log(rest); // [ 2, 3, 4 ]
```

일부 값을 생략하고 필요한 요소만 가져올 수 있다

```ts
let [first] = [1, 2, 3, 4];
console.log(first); // 1

let [, second, , fourth] = [1, 2, 3, 4];
console.log(second); // 2
console.log(fourth); // 4
```

### Object destructuring

배열 구조분해와 거의 동일하다

```ts
const { a, b } = { a: "baz", b: 101 };
```

전개 연산자로 나머지 항목을 가져올 수 있다

```ts
let { a, ...passthrough } = o;
```

프로퍼티의 이름을 바꿀 수 있다

```ts
let { a: newName1, b: newName2 } = o;
```

구조분해시 기본값을 할당할 수 있다

```ts
let { a, b = 1001 } = wholeObject;
```

## 인터페이스

타입스크립트의 인터페이스는 값의 구조에 초점을 맞추는 **Duck Typing**을 지원한다

### Property Types

```ts
interface LabelledValue {
  label: string;
}

function printLabel(labelledObj: LabelledValue) {
  console.log(labelledObj.label);
}

let myObj = { size: 10, label: "Size 10 Object" };
printLabel(myObj);
```

인터페이스의 모든 프로퍼티를 갖출필요는 없으며 이를 optional 키워드로 표시할 수 있다

```ts
interface SquareConfig {
  color?: string;
  width?: number;
}

function createSquare(config: SquareConfig): { color: string; area: number } {
  let newSquare = { color: "white", area: 100 };
  if (config.color) newSquare.color = config.color;
  if (config.width) newSquare.area = config.width * config.width;
  return newSquare;
}

let mySquare = createSquare({ color: "black" });
```

#### readonly property

인터페이스의 프로퍼티를 읽기전용으로 지정할 수도 있다

```ts
interface Point {
  readonly x: number;
  readonly y: number;
}

let p1: Point = { x: 10, y: 20 };
p1.x = 5; // Error
```

#### Excess Property Checks

타입에 없는 프로퍼티가 선언된 경우를 검출할 수 있다.
다음과 같이 type assertion을 통해 프로퍼티 초과 검사를 할 수 있다.

```ts
interface SquareConfig {
  color?: string;
  width?: number;
}

function createSquare(config: SquareConfig): { color: string; area: number } {
  // ...
}

let mySquare = createSquare({ width: 100, opacity: 0.5 } as SquareConfig);
```

### Function Types

인터페이스는 매개변수와 반환 타입만 주어진 함수선언인 call signature를 제공한다.

매개변수의 이름이 일치할 필요는 없으며, 해당 매개변수 위치와 타입을 비교한다.
또한 타입을 명시하면 contextual typing에 따라 인자 타입을 생략할 수 있다.

```ts
interface SearchFunc {
  (source: string, subString: string): boolean;
}

let mySearch: SearchFunc;
mySearch = function(src, sub) {
  let result = source.search(subString);
  return result > -1;
};
```

### Indexable Types

`obj[key]`와 같이 인덱스를 생성하는 타입을 만들 수 있다

```ts
interface StringArray {
  [index: number]: string;
}

let myArray: StringArray;
myArray = ["Bob", "Fred"];

let myStr: string = myArray[0];
```

index signature는 `string`, `number`를 사용할 수 있지만, 기본적으로 index는 `string`이다.

따라서 `number`를 사용하더라도 `string`으로 변환하여 검색하게 되므로 string과 number를 동시에 사용하게 된다면
number index signature에서 반환하는 값은 기존에 선언한 string index signature 반환값의 subtype이어야 한다.

```ts
class Animal {
  name: string;
}
class Dog extends Animal {
  breed: string;
}

// ERROR
interface NotOkay {
  [x: number]: Animal;
  [x: string]: Dog;
}
```

`obj.property`는 `obj['property']`로도 사용될 수 있으므로
string index signature가 선언된 타입의 모든 프로퍼티는 반환타입이 일치하여야 한다.

```ts
interface NumberDictionary {
  [index: string]: number;
  length: number;
  name: string; // ERROR
}
```

### Class Types

#### 인터페이스 구현

구현할 인터페이스를 명시하여 타입을 드러낼 수 있다

```ts
interface ClockInterface {
  currentTime: Date;
}

class Clock implements ClockInterface {
  currentTime: Date;
  constructor(h: number, m: number) {}
}
```

#### 생성자와 static

클래스의 생성자는 `new (...arg: any[]): T`과 같은 형식으로 표현할 수 있다 (`new` 키워드)

하지만 생성자는 static 이므로(클래스의 인스턴스가 아니므로) 타입검사에서 제외된다

```ts
interface ClockConstructor {
  new (hour: number, minute: number);
}

// ERROR
class Clock implements ClockConstructor {
  currentTime: Date;
  constructor(h: number, m: number) {}
}
```

대신 constructor signature를 가지고 있는 Class Type(아래의 DigitalClock, AnalogClock)을 해당 용도로 사용할 수 있다

```ts
interface ClockConstructor {
  new (hour: number, minute: number): ClockInterface;
}
interface ClockInterface {
  tick();
}

function createClock(ctor: ClockConstructor, hour: number, minute: number): ClockInterface {
  return new ctor(hour, minute);
}

class DigitalClock implements ClockInterface {
  constructor(h: number, m: number) {}
  tick() { ... }
}
class AnalogClock implements ClockInterface {
  constructor(h: number, m: number) {}
  tick() { ... }
}

let digital = createClock(DigitalClock, 12, 17);
let analog = createClock(AnalogClock, 7, 32);
```

### Extending Interfaces

인터페이스도 확장(상속) 가능하다

```ts
interface Shape {
  color: string;
}

interface PenStroke {
  penWidth: number;
}

interface Square extends Shape, PenStroke {
  sideLength: number;
}

let square = <Square>{};
square.color = "blue";
square.sideLength = 10;
square.penWidth = 5.0;
```

### Interfaces Extending Classes

인터페이스가 클래스도 확장(상속)가능하지만 해당 클래스의 구현을 확장하지는 않는다

클래스를 확장한 인터페이스를 구현하기 위해서는
해당 인터페이스가 확장한 클래스에 포함된 모든 멤버들을 선언해야 하며, 기존의 접근제한자 레벨도 유지해야 한다.

```ts
class Control {
  private state: any;
}

interface SelectableControl extends Control {
  select(): void;
}

class Button extends Control implements SelectableControl {
  select() {}
}

// ERROR: Control 클래스를 확장해야 함 (state 부재)
class Image implements SelectableControl {
  select() {}
}
```

## 클래스

### 상속

`extends` 키워드를 사용하여 클래스를 확장할 수 있다.

```ts
class Animal {
  name: string;
  constructor(theName: string) {
    this.name = theName;
  }

  move(distanceInMeters: number = 0) {
    console.log(`${this.name} moved ${distanceInMeters}m.`);
  }
}

class Snake extends Animal {
  constructor(name: string) {
    super(name);
  }

  move(distanceInMeters = 5) {
    console.log("Slithering...");
    super.move(distanceInMeters);
  }
}

class Horse extends Animal {
  constructor(name: string) {
    super(name);
  }

  move(distanceInMeters = 45) {
    console.log("Galloping...");
    super.move(distanceInMeters);
  }
}

const sam = new Snake("Sammy the Python");
const tom: Animal = new Horse("Tommy the Palomino");

sam.move();
// > Slithering...
// > Sammy the Python moved 5m.
tom.move(34);
// > Galloping...
// > Tommy the Palomino moved 34m.
```

### 접근 제한자

#### public

TypeScript에서 멤버들의 기본 접근 제한자는 `public`이다.

따라서 public은 명시하지 않아도 되고, 아무런 접근자가 없는 상태는 public이다.

#### private

private으로 선언된 멤버는 클래스 외부에서 접근할 수 없다.

해당 클래스를 확장한 클래스에서도 접근 제한레벨은 동일하다.

만약 private 멤버를 가진 클래스를 확장한 클래스에서 부모의 private 멤버 변수와 동일한 이름의 변수를 선언하면,
해당 변수는 부모의 멤버 변수와 별개로 존재하는 것이다.

#### protected

protected 접근 제한자는 해당 클래스 내부와 해당 클래스를 확장한 클래스에서 접근 가능하다.

생성자 역시 protected로 표시할 수 있으며, 이는 해당 클래스를 인스턴스화 할 수는 없지만 확장 할 수는 있음을 말한다.

### readonly modifier

`readonly` 키워드를 통해 읽기 전용 프로퍼티를 선언할 수 있다.
읽기 전용 프로피티는 선언시, 혹은 생성자에서 초기화 해야 한다.

```ts
class Octopus {
  readonly name: string;
  readonly numberOfLegs: number = 8;

  constructor (theName: string) {
    this.name = theName;
  }
}
let dad = new Octopus("Man with the 8 strong legs");
dad.name = "Man with the 3-piece suit"; // ERROR!
```

#### Parameter Properties

생성자 매개변수에서 접근 제한자(`public`, `protected`, `private`) or `readonly` 키워드 or 둘 다를 사용하여,
멤버변수의 선언과 할당을 동시에 할 수 있다.

```ts
class Octopus {
  readonly numberOfLegs: number = 8;

  constructor(readonly name: string) {}
}
```

### Accessors (getter/setter)

TypeScript의 getter/setter는 객체 멤버에 대한 접근을 가로채는 방법을 사용한다.

별도의 getter/setter를 지정하기 위해서 `get`, `set` 키워드를 사용한다.

```ts
let passcode = "secret passcode";

class Employee {
  private _fullName: string;

  get fullName(): string {
    return this._fullName;
  }

  set fullName(newName: string) {
    if (passcode && passcode == "secret passcode") {
      this._fullName = newName;
    }
    else {
      console.log("오류 : employee의 무단 업데이트!");
    }
  }
}

let employee = new Employee();
employee.fullName = "Bob Smith";
if (employee.fullName) {
  console.log(employee.fullName);
}
```

get 접근자는 있지만 set 접근자가 없다면 자동으로 `readonly`로 추론된다.

### Static Properties

클래스에 static 멤버도 생성할 수 있다.

```ts
class Grid {
  static origin = "hello";
}

console.log(Grid.origin); // > hello
```

### 추상 클래스

Java의 추상 클래스와 거의 동일하다.
직접적으로 인스턴스화 할 수 없으며, 확장하여 인스턴스화 해야 한다.

세부 구현을 포함할 수도 있고, `abstract` 키워드를 활용하여 추상 메소드를 선언할 수 있다.
추상 메소드 선언시 접근 제한자를 표시할 수 있다.

### Advanced Techniques

#### 생성자 함수(Constructor Functions)

TypeScript의 클래스 선언은 실제로는 여러개의 선언을 동시에 생성한다.
**클래스 인스턴스의 타입**을 생성하고, **생성자 함수**라고 부르는 값을 생성한다.

생성자 함수는 `new` 키워드가 사용되면 호출되는 함수이다.

```ts
class Greeter {
  greeting: string;
  constructor(message: string) {
    this.greeting = message;
  }
  greet() {
    return "Hello, " + this.greeting;
  }
}

let greeter: Greeter;
greeter = new Greeter("world");
```

이와 같은 클래스 선언과 생성자 함수 호출은 실제로는 아래와 같은 방식으로 진행된다.

```ts
let Greeter = (function () {
  function Greeter(message) {
    this.greeting = message;
  }
  Greeter.prototype.greet = function () {
    return "Hello, " + this.greeting;
  };
  return Greeter;
})();

let greeter;
greeter = new Greeter("world");
```

#### 생성자 함수와 타입

아래의 `greeterMaker` 변수의 타입은 `typeof Greeter`이다.
이는 인스턴스 타입이 아닌 Greeter 클래스 자체의 타입이며,
보다 정확히 말하자면 생성자 함수 타입인 Greeter라는 symbol 타입이다.

생성자 함수 타입에는 생성자와 함께 해당 클래스의 모든 static 멤버가 포함된다.

```ts
class Greeter {
  static standardGreeting = "Hello, there";
  greeting: string;
  greet() {
    if (this.greeting) {
      return "Hello, " + this.greeting;
    }
    else {
      return Greeter.standardGreeting;
    }
  }
}

let greeter1: Greeter;
greeter1 = new Greeter();
console.log(greeter1.greet());

let greeterMaker: typeof Greeter = Greeter;
greeterMaker.standardGreeting = "Hey there!";

let greeter2: Greeter = new greeterMaker();
console.log(greeter2.greet());
```

#### 클래스를 인터페이스로 사용

클래스 선언시 클래스 인스턴스의 타입이 생성되는데, 생성된 타입을 인터페이스 사용되는 위치에서 사용할 수 있다.

```ts
class Point {
  x: number;
  y: number;
}

interface Point3d extends Point {
  z: number;
}

let point3d: Point3d = { x: 1, y: 2, z: 3 };
```

## 함수

TypeScript도 이름이 있는 함수와 익명함수를 생성할 수 있다.

JavaScript와 마찬가지로 함수는 함수 외부의 변수를 capture 할 수 있다.

### Function Type

함수는 반환값에 따라 반환 타입을 추론할 수 있으므로 함수 선언시 반환타입은 선택적으로 작성할 수 있다.

```ts
const myAdd: (baseValue: number, increment: number) => number =
  function(x: number, y: number): number {
    return x + y;
  };
```

`myAdd` 변수 타입은 함수형 타입이다.
TypeScript는 타입추론을 지원하므로 할당문 양쪽 중 한쪽은 타입을 생략할 수 있다.

### Parameters

#### Optional Parameters

TypeScript 함수에서 인수의 개수는 선언된 함수의 매개변수의 수와 같아야 한다.

만약 매개변수를 선택적으로 사용하려면 매개변수 선언시에 변수의 끝에 `?`를 추가하면 된다.

```ts
function buildName(firstName: string, lastName?: string) {
  if (lastName)
    return firstName + " " + lastName;
  else
    return firstName;
}

let result1 = buildName("Bob");                  // OK
let result2 = buildName("Bob", "Adams", "Sr.");  // ERROR
```

필수 매개변수의 앞에 선택적 매개변수를 선언할 수 없다.

#### Default Parameters

함수 / 메소드를 사용할 때 매개변수를 전달하지 않거나(생략) undefined를 전달하더라도 매개변수에 할당되는 기본값을 정할 수 있다.

```ts
function buildName(firstName: string, lastName = "Smith") {
  return firstName + " " + lastName;
}

let result1 = buildName("Bob");                  // Bob Smith
let result2 = buildName("Bob", undefined);       // Bob Smith
let result3 = buildName("Bob", "Adams", "Sr.");  // ERROR: 매개변수가 너무 많음
```

선택적 매개변수와는 달리 필수 매개변수 앞이라도 기본값을 설정할 수 있다.

#### Rest Parameters

`...` 키워드를 사용하여 매개변수를 묶어서 표현할 수 있다.

```ts
function buildName(firstName: string, ...restOfName: string[]) {
    return firstName + " " + restOfName.join(" ");
}

let employeeName = buildName("Joseph", "Samuel", "Lucas", "MacKinzie");
```

### `this`

#### `this` and arrow function

JavaScript에서 `this` 바인딩은 함수 호출 환경과 관련되어 있으므로, 호출환경을 알아야한다는 문제가 있다

```ts
let deck = {
  suits: ["hearts", "spades", "clubs", "diamonds"],
  cards: Array(52),
  createCardPicker: function() {
    return function() {
      let pickedCard = Math.floor(Math.random() * 52);
      let pickedSuit = Math.floor(pickedCard / 13);

      return {suit: this.suits[pickedSuit], card: pickedCard % 13};
    }
  }
}

let cardPicker = deck.createCardPicker();
let pickedCard = cardPicker(); // ERROR: this.suits의 this는 window(strict에서는 undefined)이다
```

ES6의 Arrow Fuction에서 `this`는 함수가 생성된 곳에서 `this`를 capture 한다

```ts
let deck = {
  suits: ["hearts", "spades", "clubs", "diamonds"],
  cards: Array(52),
  createCardPicker: function() {
    // 아래 화살표 함수는 현재 위치의 this를 캡쳐한다
    return () => {
      let pickedCard = Math.floor(Math.random() * 52);
      let pickedSuit = Math.floor(pickedCard / 13);

      return {suit: this.suits[pickedSuit], card: pickedCard % 13};
    }
  }
}

let cardPicker = deck.createCardPicker();
let pickedCard = cardPicker();
```

#### `this` parameters

하지만 위의 `this`는 여전히 `any` 타입이다.

이를 해결하기 위해서 명시적으로 `this` 매개변수를 사용할 수 있다.

```ts
interface Card {
  suit: string;
  card: number;
}
interface Deck {
  suits: string[];
  cards: number[];
  createCardPicker(this: Deck): () => Card;
}
let deck: Deck = {
  suits: ["hearts", "spades", "clubs", "diamonds"],
  cards: Array(52),
  createCardPicker: function(this: Deck) {
    // 현재 함수의 this 타입을 명시하였다 (해당 메소드가 Deck 타입의 객체에서 호출될 것을 기대함)
    return () => {
      let pickedCard = Math.floor(Math.random() * 52);
      let pickedSuit = Math.floor(pickedCard / 13);

      return {suit: this.suits[pickedSuit], card: pickedCard % 13};
    }
  }
}

let cardPicker = deck.createCardPicker();
let pickedCard = cardPicker();
```

##### `this` parameters in callbacks

라이브러리를 사용할 때 전달하는 콜백함수에서 `this`를 사용한다면 오류가 발생할 수 있다.

이럴 때 `this` parameter를 사용해서 오류를 방지할 수 있다.

`this`를 `void`로 선언하면 콜백함수에서 `this`타입이 필요하지 않다는 것을 의미한다.

```ts
interface UIElement {
  addClickListener(onclick: (this: void, e: Event) => void): void;
}
```

화살표 함수는 `this`를 capture하지 않으므로 항상 `this: void`를 넘겨줄 수 있다.

```ts
class Handler {
  info: string;
  onClickGood = (e: Event) => { this.info = e.message }
}
```

다만 메소드는 핸들러의 프로토타입에 속하여 하나만 만들어지고,
화살표 함수는 Handler 타입의 객체마다 하나씩 생성된다.

### Overloads

TypeScript에서는 동일 함수에 다른 parameter type을 제공하여 Overloading을 사용할 수 있다

## 제네릭

제네릭을 통해서 여러 타입을 입력/출력할 때 타입을 유지하면서 코드를 재사용한다.

```ts
function identity<T>(arg: T): T {
  return arg;
}
```

타입변수 `T`를 추가한 함수를 사용할 때는 두 가지 방법으로 호출할 수 있다.

```ts
// 타입인수를 명시
let output = identity<string>("myString");

// 전달하는 인자를 통해 타입추론
let output = identity("myString");
```

### 제네릭 타입

```ts
function identity<T>(arg: T): T {
  return arg;
}

// 제네릭 함수의 타입은 일반 함수 타입의 앞에 제네릭 파라미터를 표기한 모양이다
let myIdentity: <T>(arg: T) => T = identity;
// 제네릭 타입 변수명을 다르게 사용할 수도 있다
let myIdentity: <U>(arg: U) => U = identity;
// 제네릭 함수의 타입을 객체 리터럴의 호출 시그니처 형태로 쓸수도 있다
let myIdentity: {<T>(arg: T): T} = identity;
```

### 제네릭 인터페이스 / 클래스

인터페이스를 사용하여 제네릭 타입을 표기할 수 있다

```ts
interface GenericIdentityFn {
    <T>(arg: T): T;
}
interface GenericIdentityFn<T> {
  (arg: T): T;
}

function identity<T>(arg: T): T {
  return arg;
}

let myIdentityA: GenericIdentityFn = identity;
let myIdentityB: GenericIdentityFn<number> = identity;
```

클래스에서도 인터페이스와 비슷한 형태로 사용할 수 있다

```ts
class GenericNumber<T> {
  zeroValue: T;
  add: (x: T, y: T) => T;
}

let myGenericNumber = new GenericNumber<number>();
myGenericNumber.zeroValue = 0;
myGenericNumber.add = function(x, y) { return x + y; };
```

### 제네릭 제약조건

`extends` 키워드를 사용하여 제네릭 타입에 대한 제약조건을 작성할 수 있다

```ts
interface Lengthwise {
  length: number;
}

function loggingIdentity<T extends Lengthwise>(arg: T): T {
  console.log(arg.length); // T 타입은 length 프로퍼티를 갖고있는 타입이다
  return arg;
}

```

#### Using Type Parameters in Generic Constraints

여러 제네릭 타입이 사용될 때 한 타입 매개변수로 다른 타입매개변수의 제약조건을 작성할 수 있다

```ts
function getProperty<T, K extends keyof T>(obj: T, key: K) {
  return obj[key];
}

let x = { a: 1, b: 2, c: 3, d: 4 };

getProperty(x, "a");
getProperty(x, "m"); // ERROR: "m"을 a | b | c | d 에 할당할 수 없음
```

#### Using Class Types in Generics

생성자 함수를 갖는 클래스 타입에서 제네릭을 사용할 수 있다. 여러 타입을 사용한 예제를 보자.

```ts
class BeeKeeper {
  hasMask: boolean;
}

class ZooKeeper {
  nametag: string;
}

class Animal {
  numLegs: number;
}

class Bee extends Animal {
  keeper: BeeKeeper;
}

class Lion extends Animal {
  keeper: ZooKeeper;
}

function createInstance<A extends Animal>(c: new () => A): A {
  return new c();
}

createInstance(Lion).keeper.nametag;
createInstance(Bee).keeper.hasMask;
```

## Enums

### Numeric Enums

enum을 선언하면 순서대로 숫자 0부터 값이 부여된다.
만약 첫 멤버의 초기값을 선언하였다면 이후의 멤버의 값은 자동 증가한다.

```ts
enum Direction {
  Up,     // 0
  Down,   // 1
  Left,   // 2
  Right,  // 3
}

enum Direction {
  Up = 1, // 1
  Down,   // 2
  Left,   // 3
  Right,  // 4
}
```

만약 한 멤버가 상수가 아닌 초기 값을 가진다면, 다른 멤버도 초기화 되어야 한다

```ts
enum E {
  A = getSomeValue(),
  B, // ERROR: A가 상수가 아니므로 B도 초기화 필요
}
```

### String Enums

```ts
enum Direction {
  Up = "UP",
  Down = "DOWN",
  Left = "LEFT",
  Right = "RIGHT",
}
```

string enum은 자동증가 하지 않지만 직렬화시 장점이 있다.

### Heterogeneous Enums

숫자와 문자 값을 섞어 쓸 수 있으나 그렇게 하지 않는 것이 좋다

```ts
enum BooleanLikeHeterogeneousEnum {
  No = 0,
  Yes = "YES",
}
```

### Computed and constant members

enum 멤버는 computed 이거나 constant인 값이다
상수형 enum 멤버는 컴파일 시간에 완전히 평가 될 수 있다.

다음과 같은 경우의 표현식은 상수 enum 표현식이고, 이외의 경우 computed 멤버이다.

1. 문자 리터럴 혹은 숫자 리터럴
2. 이전에 정의된 enum 멤버에 대한 참조(다른 enum에서 올 수 있음)
3. 괄호 내부의 상수 enum 표현식
4. 상수 enum 표현식에 사용된 단항연산자 `+`, `-`, `~` 중의 하나
5. `+`, `-`, `*`, `/`, `%`, `<<`, `>>`, `>>>`, `&`, `|`, `^` 다음 이항 연산자와 함께 사용된 피연산자 상수 enum 표현식. 상수 enum 표현식이 `NaN`이나 `Infinity`인 경우 컴파일 에러이다.

### Union enums and enum member types

계산되지 않은 상수형 enum 멤버의 특수하위집합이 있는데, 리터럴 enum 멤버이다.

리터럴 enum 멤버는 초기값이 없거나, 다음으로 초기화된 상수 enum 멤버이다

- any string literal (e.g. "foo", "bar, "baz")
- any numeric literal (e.g. 1, 100)
- a unary minus applied to any numeric literal (e.g. -1, -100)

모든 멤버가 리터럴 enum일 때 enum은 특수한 기능을 갖는다

- enum 멤버는 타입이 될 수 있다

  ```ts
  enum ShapeKind {
    Circle,
    Square,
  }

  interface Circle {
    kind: ShapeKind.Circle;
    radius: number;
  }

  interface Square {
    kind: ShapeKind.Square;
    sideLength: number;
  }
  ```

- enum 타입 자체가 각 enum 멤버의 합집합(union)이 된다

  ```ts
  enum E {
    Foo,
    Bar,
  }

  function f(x: E) {
    if (x !== E.Foo || x !== E.Bar) {
      // Error! Operator '!==' cannot be applied to types 'E.Foo' and 'E.Bar'
      // x는 이미 E 타입이다
    }
  }
  ```

### Enums at runtime

enum은 런타임에 존재하는 real objects이다

```ts
enum E {
  X, Y, Z
}

function f(obj: { X: number }) {
  return obj.X;
}

f(E); // Works, since 'E' has a property named 'X' which is a number
```

#### Reverse mappings

숫자 enum 멤버는 enum 값에서 이름으로 역 매핑을 할 수 있다.
문자 enum 멤버는 역매핑을 생성하지 않는다.

```ts
enum Enum {
    A
}
let a = Enum.A;
let nameOfA = Enum[a]; // "A"
```

타입스크립트는 다음과 같은 자바스크립트로 컴파일 될 것이다

```js
var Enum;
(function (Enum) {
    Enum[Enum["A"] = 0] = "A";
})(Enum || (Enum = {}));
var a = Enum.A;
var nameOfA = Enum[a]; // "A"
```

#### const enums

enum 사용시 추가적인 비용을 줄이기 위해서 const enums를 사용할 수 있다

const enum은 상수 enum 표현식만 사용할 수 있으면 일반 enum과 달리 컴파일 하는 동안 완전히 제거되어 사용장소에서 인라이닝 된다.

```ts
const enum Directions {
    Up,
    Down,
    Left,
    Right
}
let directions = [Directions.Up, Directions.Down, Directions.Left, Directions.Right]

// 컴파일 되면 다음과 같아질 것이다
var directions = [0 /* Up */, 1 /* Down */, 2 /* Left */, 3 /* Right */];
```

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
  console.log(mouseEvent.button);   //<- OK
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
let x = () => ({name: "Alice"});
let y = () => ({name: "Alice", location: "Seattle"});

x = y;
y = x; // ERROR
```

타입시스템은 원본 함수 반환형이 대상 타입 반환형의 서브타입이 되도록 강제한다 (공변성)

#### Function Parameter Bivariance

함수 파라미터의 타입을 비교할 때, 원본 파라미터가 대상 파라미터에 할당 가능하거나 그 반대일 경우 할당이 성공한다.
하지만, 호출한 쪽에서 less specialized type argument로 함수를 호출할 수도 있는데 이는 바람직 하지 않다.

```ts
enum EventType { Mouse, Keyboard }

interface Event { timestamp: number; }
interface MouseEvent extends Event { x: number; y: number }
interface KeyEvent extends Event { keyCode: number }

function listenEvent(eventType: EventType, handler: (n: Event) => void) {
  /* ... */
}

// 바람직하지 않지만, 유용하고 일반적이다
listenEvent(EventType.Mouse, (e: MouseEvent) => console.log(e.x + "," + e.y));

// 바람직하지 않은 대안
listenEvent(EventType.Mouse, (e: Event) => console.log((<MouseEvent>e).x + "," + (<MouseEvent>e).y));
listenEvent(EventType.Mouse, <(e: Event) => void>((e: MouseEvent) => console.log(e.x + "," + e.y)));

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
enum Status { Ready, Waiting };
enum Color { Red, Blue, Green };

let status = Status.Ready;
status = Color.Green;  // ERROR
```

### Class 타입 호환성

클래스는 한 가지만 제외하고 객체 리터럴 타입이나 인터페이스와 동일하게 작동한다.

클래스유형의 두 객체를 비교할 때 인스턴스 멤버만 비교하고, 정적 멤버 및 생성자는 호환성에 영향을 주지 않는다.

```ts
class Animal {
  feet: number;
  constructor(name: string, numFeet: number) { }
}

class Size {
  feet: number;
  constructor(numFeet: number) { }
}

let a: Animal;
let s: Size;

a = s;  // OK
s = a;  // OK
```

#### Private and protected members in classes

클래스의 private | protected 멤버는 호환성에 영향을 준다.

이를 통해 클래스는 super class와 호환되어 할당 가능하지만, 동일한 형태의 다른 상속 계층 클래스와는 호환되지 않는다.

### Generics 타입 호환성

TypeScript는 구조적 타입 시스템이므로 타입 파라미터는 멤버의 일부로 사용될 때만 결과 타입에 영향을 준다.

```ts
interface Empty<T> { }
let x: Empty<number>;
let y: Empty<string>;

x = y;  // OK, because y matches structure of x
```

위의 경우는 타입 파라미터를 다르게 사용하지 않으므로 구조적으로 호환 가능하다.

```ts
interface NotEmpty<T> {
  data: T;
}
let x: NotEmpty<number>;
let y: NotEmpty<string>;

x = y;  // Error, because x and y are not compatible
```

타입 인자가 명시된 제네릭 타입은 실제로는 비-제네릭 타입처럼 작동한다.

타입 인자가 지정되지 않은 제네릭 타입의 경우 지정되지 않은 타입 인자 대신 `any`를 지정하여 호환성을 확인한다.
그리고 나서, 비-제네릭 타입과 같은 방식으로 결과 타입을 확인한다.

```ts
let identity = function<T>(x: T): T {
  // ...
}
let reverse = function<U>(y: U): U {
  // ...
}

identity = reverse;  // OK, because (x: any) => any matches (y: any) => any
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
  constructor(public name: string) { }
}

interface Loggable {
  log(name: string): void;
}

class ConsoleLogger implements Loggable {
  log(name) {
    console.log(`Hello, I'm ${name}.`);
  }
}

const jim = extend(new Person('Jim'), ConsoleLogger.prototype);
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
pet.swim();    // errors
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
  getPaddingString(): string
}

class SpaceRepeatingPadder implements Padder {
  constructor(private numSpaces: number) { }
  getPaddingString() {
    return Array(this.numSpaces + 1).join(" ");
  }
}

class StringPadder implements Padder {
  constructor(private value: string) { }
  getPaddingString() {
    return this.value;
  }
}

function getRandomPadder() {
  return Math.random() < 0.5 ?
    new SpaceRepeatingPadder(4) :
    new StringPadder("  ");
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
  }
  else {
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
  }
  else {
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
}
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
type Alias = { num: number }
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
    }
    else if (easing === "ease-out") {
    }
    else if (easing === "ease-in-out") {
    }
    else {
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
  if (x !== 1 || x !== 2) {  // ERROR: Operator '!==' cannot be applied to types '1' and '2'.
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
    case "square": return s.size * s.size;
    case "rectangle": return s.height * s.width;
    case "circle": return Math.PI * s.radius ** 2;
  }
}
```

#### Exhaustiveness checking

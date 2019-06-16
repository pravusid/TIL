# TypeScript HandBook

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

f(true); // 10
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
for (let i = 0; i < 10; i++) {
  setTimeout(function() {
    console.log(i);
  }, 100 * i);
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
  m() {}
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

  constructor(theName: string) {
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
    } else {
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
let Greeter = (function() {
  function Greeter(message) {
    this.greeting = message;
  }
  Greeter.prototype.greet = function() {
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
    } else {
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
const myAdd: (baseValue: number, increment: number) => number = function(
  x: number,
  y: number
): number {
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
  if (lastName) return firstName + " " + lastName;
  else return firstName;
}

let result1 = buildName("Bob"); // OK
let result2 = buildName("Bob", "Adams", "Sr."); // ERROR
```

필수 매개변수의 앞에 선택적 매개변수를 선언할 수 없다.

#### Default Parameters

함수 / 메소드를 사용할 때 매개변수를 전달하지 않거나(생략) undefined를 전달하더라도 매개변수에 할당되는 기본값을 정할 수 있다.

```ts
function buildName(firstName: string, lastName = "Smith") {
  return firstName + " " + lastName;
}

let result1 = buildName("Bob"); // Bob Smith
let result2 = buildName("Bob", undefined); // Bob Smith
let result3 = buildName("Bob", "Adams", "Sr."); // ERROR: 매개변수가 너무 많음
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

      return { suit: this.suits[pickedSuit], card: pickedCard % 13 };
    };
  }
};

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

      return { suit: this.suits[pickedSuit], card: pickedCard % 13 };
    };
  }
};

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

      return { suit: this.suits[pickedSuit], card: pickedCard % 13 };
    };
  }
};

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
  onClickGood = (e: Event) => {
    this.info = e.message;
  };
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
let myIdentity: { <T>(arg: T): T } = identity;
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
myGenericNumber.add = function(x, y) {
  return x + y;
};
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
  Up, // 0
  Down, // 1
  Left, // 2
  Right // 3
}

enum Direction {
  Up = 1, // 1
  Down, // 2
  Left, // 3
  Right // 4
}
```

만약 한 멤버가 상수가 아닌 초기 값을 가진다면, 다른 멤버도 초기화 되어야 한다

```ts
enum E {
  A = getSomeValue(),
  B // ERROR: A가 상수가 아니므로 B도 초기화 필요
}
```

### String Enums

```ts
enum Direction {
  Up = "UP",
  Down = "DOWN",
  Left = "LEFT",
  Right = "RIGHT"
}
```

string enum은 자동증가 하지 않지만 직렬화시 장점이 있다.

### Heterogeneous Enums

숫자와 문자 값을 섞어 쓸 수 있으나 그렇게 하지 않는 것이 좋다

```ts
enum BooleanLikeHeterogeneousEnum {
  No = 0,
  Yes = "YES"
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
    Square
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
    Bar
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
  X,
  Y,
  Z
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
(function(Enum) {
  Enum[(Enum["A"] = 0)] = "A";
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
let directions = [
  Directions.Up,
  Directions.Down,
  Directions.Left,
  Directions.Right
];

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

#### Predefined conditional types

타입스크립트 2.8에서 `lib.d.ts`에 사전 정의된 조건부 타입들이 있다

- `Exclude<T, U>` -- Exclude from T those types that are assignable to U.
- `Extract<T, U>` -- Extract from T those types that are assignable to U.
- `NonNullable<T>` -- Exclude null and undefined from T.
- `ReturnType<T>` -- Obtain the return type of a function type.
- `InstanceType<T>` -- Obtain the instance type of a constructor function type.

```ts
type T00 = Exclude<"a" | "b" | "c" | "d", "a" | "c" | "f">; // "b" | "d"
type T01 = Extract<"a" | "b" | "c" | "d", "a" | "c" | "f">; // "a" | "c"

type T02 = Exclude<string | number | (() => void), Function>; // string | number
type T03 = Extract<string | number | (() => void), Function>; // () => void

type T04 = NonNullable<string | number | undefined>; // string | number
type T05 = NonNullable<(() => string) | string[] | null | undefined>; // (() => string) | string[]

function f1(s: string) {
  return { a: 1, b: s };
}

class C {
  x = 0;
  y = 0;
}

type T10 = ReturnType<() => string>; // string
type T11 = ReturnType<(s: string) => void>; // void
type T12 = ReturnType<<T>() => T>; // {}
type T13 = ReturnType<<T extends U, U extends number[]>() => T>; // number[]
type T14 = ReturnType<typeof f1>; // { a: number, b: string }
type T15 = ReturnType<any>; // any
type T16 = ReturnType<never>; // never
type T17 = ReturnType<string>; // Error
type T18 = ReturnType<Function>; // Error

type T20 = InstanceType<typeof C>; // C
type T21 = InstanceType<any>; // any
type T22 = InstanceType<never>; // never
type T23 = InstanceType<string>; // Error
type T24 = InstanceType<Function>; // Error
```

> `Exclude` 타입은 `Diff` 타입의 정확한 구현이다. `Diff`가 정의되어 있는 코드와 충돌을 회피하기 위해서 `Exclude`로 명명하였다. 또한 의미론적으로 더 나은 느낌을 전달한다. `Omit<T, K>` 타입은 포함되지 않았는데 `Pick<T, Exclude<keyof T, K>>`타입으로 사용할 수 있기 때문이다.

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

## Modules

ECMAScript 2015부터 JavaScript에는 모듈 개념이 있고, TypeScript도 개념을 공유한다.

모듈은 전역 스코프가 아닌 자체 스코프에서 실행된다.
이는 모듈에서 선언된 변수, 함수, 클래스 등을 `export` 양식을 사용하여 명시적으로 내보내지 않으면 외부에서 볼 수 없음을 의미한다.

반대로 다른 모듈에서 내보낸 변수, 함수, 클래스, 인터페이스 등을 사용하려면 `import` 양식을 사용하여야 한다.

모듈은 선언적이므로 모듈간의 관계는 파일 수준의 `import`, `export`로 지정된다.

모듈은 모듈 로더를 사용하여 서로를 불러온다.
런타임에서 모듈 로더는 모듈을 실행하기 전 모듈의 모든 종속성을 찾아 실행한다.

잘 알려진 모듈 로더는 Node.js의 CommonJS 모듈 로더와 웹 어플리케이션의 require.js이다.

TypeScript에서는 ES6와 마찬가지로 top-level import 또는 export가 포함된 파일을 모듈로 간주한다.
반대로, top-level import 또는 export가 없는 파일은 전역 스코프에서 사용할 수 있는 스크립트로 취급된다.

### Export

#### Exporting a declaration

`export` 키워드를 추가하여 모든 선언(변수, 함수, 클래스, 타입 별칭, 인터페이스 ...)을 내보낼 수 있다.

```ts
// Validation.ts
export interface StringValidator {
  isAcceptable(s: string): boolean;
}

// ZipCodeValidator.ts
export const numberRegexp = /^[0-9]+$/;
export class ZipCodeValidator implements StringValidator {
  isAcceptable(s: string) {
    return s.length === 5 && numberRegexp.test(s);
  }
}
```

#### Export statements

export문을 사용할 때 편리하게 다른이름을 지정할 수 있다

```ts
class ZipCodeValidator implements StringValidator {
  isAcceptable(s: string) {
    return s.length === 5 && numberRegexp.test(s);
  }
}
export { ZipCodeValidator };
export { ZipCodeValidator as mainValidator };
```

#### Re-exports

종종 모듈은 다른 모듈을 확장하고 부분적으로 일부기능을 노출한다.
Re-exports는 대상 모듈을 로컬에 불러오거나 변수로 선언하지 않고 내보내기를 한다.

```ts
// ParseIntBasedZipCodeValidator.ts
export class ParseIntBasedZipCodeValidator {
  isAcceptable(s: string) {
    return s.length === 5 && parseInt(s).toString() === s;
  }
}
// Export original validator but rename it
export {
  ZipCodeValidator as RegExpBasedZipCodeValidator
} from "./ZipCodeValidator";
```

선택적으로 모듈은 하나 이상의 모듈을 감싸고 결합한 뒤 `export * from "module"` 문법을 통해 내보낼수 있다.

```ts
export * from "./StringValidator"; // exports interface 'StringValidator'
export * from "./LettersOnlyValidator"; // exports class 'LettersOnlyValidator'
export * from "./ZipCodeValidator"; // exports class 'ZipCodeValidator'
```

### Import

내보내기 선언을 불러오려면 아래의 `import` 중 하나를 사용한다

#### Import a single export from a module

```ts
import { ZipCodeValidator } from "./ZipCodeValidator";
let myValidator = new ZipCodeValidator();
```

import를 하면서 이름을 다시 지정할 수 있다

```ts
import { ZipCodeValidator as ZCV } from "./ZipCodeValidator";
let myValidator = new ZCV();
```

#### Import the entire module into a single variable, and use it to access the module exports

```ts
import * as validator from "./ZipCodeValidator";
let myValidator = new validator.ZipCodeValidator();
```

#### Import a module for side-effects only

권장 사항은 아니지만 일부 모듈은 다른 모듈에서 사용할 수 있는 일부 전역 상태를 설정한다.
이런 모듈에는 내보내기가 없거나 모듈 사용처에서 모듈의 내보내기를 사용하지 않는 경우이다.

이런 경우 다음과 같이 불러오기를 사용한다.

```ts
import "./my-module.js";
```

### Default exports

각 모듈은 선택적으로 `default` 키워드를 통해 기본 내보내기를 할 수 있다.
모듈당 하나의 기본 내보내기만 할 수 있다.

```ts
// JQuery.d.ts
declare let $: JQuery;
export default $;

// App.ts
import $ from "JQuery";
$("button.continue").html("Next Step...");
```

클래스나 함수 선언은 기본 내보내기로 직접 작성될 수 있다. 이름 선언은 선택사항이다.

```ts
// ZipCodeValidator.ts
export default class ZipCodeValidator {
  static numberRegexp = /^[0-9]+$/;
  isAcceptable(s: string) {
    return s.length === 5 && ZipCodeValidator.numberRegexp.test(s);
  }
}

// Test.ts
import validator from "./ZipCodeValidator";
let myValidator = new validator();
```

또는

```ts
// StaticZipCodeValidator.ts
const numberRegexp = /^[0-9]+$/;
export default function(s: string) {
  return s.length === 5 && numberRegexp.test(s);
}

// Test.ts
import validate from "./StaticZipCodeValidator";
let myValidator = new validator();
```

기본 내보내기는 단순히 값일 수도 있다

```ts
// OneTwoThree.ts
export default "123";

// Log.ts
import num from "./OneTwoThree";
console.log(num); // "123"
```

### export = and import = require()

CommonJS와 AMD 모두 일반적으로 모듈의 모든 내보내기를 포함하는 내보내기 객체 개념이 있다.

또한 내보내기 객체를 사용자 지정 단일 객체로 바꾸는 기능도 지원한다.
기본 내보내기는 이 동작을 대신하는 역할을 한다. 그러나 앞의 두 동작은 호환되지 않는다.

TypeScript는 일반적인 CommonJS 및 AMD 워크플로우를 모델링하기 위해 `export =`을 지원한다.

`export =` 문법은 모듈에서 내보낼 단일 객체를 지정한다. 객체는 클래스, 인터페이스, 네임 스페이스, 함수, Enum이 될 수 있다.

`export =`을 사용해 내보낸 모듈을 가져올 때 TypeScript 고유의 `import module = require("module")`을 사용해야 한다.

```ts
// ZipCodeValidator.ts
let numberRegexp = /^[0-9]+$/;
class ZipCodeValidator {
  isAcceptable(s: string) {
    return s.length === 5 && numberRegexp.test(s);
  }
}
export = ZipCodeValidator;

// Test.ts
import zip = require("./ZipCodeValidator");
let validator = new zip();
```

### Code Generation for Modules

컴파일하는 동안 지정된 모듈 타겟에 따라 컴파일러는 Node.js(CommonJS), require.js(AMD), UMD, SystemJS 또는
ES6 모듈로드 시스템에 적합한 코드를 생성한다.

```js
// SimpleModule.ts
import m = require("mod");
export let t = m.something + 1;

// AMD / RequireJS SimpleModule.js
define(["require", "exports", "./mod"], function (require, exports, mod_1) {
  exports.t = mod_1.something + 1;
});

// CommonJS / Node SimpleModule.js
var mod_1 = require("./mod");
exports.t = mod_1.something + 1;

// UMD SimpleModule.js
(function (factory) {
  if (typeof module === "object" && typeof module.exports === "object") {
    var v = factory(require, exports); if (v !== undefined) module.exports = v;
  }
  else if (typeof define === "function" && define.amd) {
    define(["require", "exports", "./mod"], factory);
  }
})(function (require, exports) {
  var mod_1 = require("./mod");
  exports.t = mod_1.something + 1;
});

// System SimpleModule.js
System.register(["./mod"], function(exports_1) {
  var mod_1;
  var t;
  return {
    setters:[
      function (mod_1_1) {
        mod_1 = mod_1_1;
      }],
    execute: function() {
      exports_1("t", t = mod_1.something + 1);
    }
  }
});

// Native ECMAScript 2015 modules SimpleModule.js
import { something } from "./mod";
export var t = something + 1;
```

### Working with Other JavaScript Libraries

타입스크립트로 작성되지 않은 라이브러리의 형태를 기술하려면, 라이브러리가 노출하는 API를 선언해야 한다.

#### Ambient Modules

최상위 수준의 내보내기 선언을 사용하여 각 모듈을 자체 `d.ts` 파일로 정의할 수 있지만, 더 큰 `d.ts` 파일로 작성하는 것이 편리하다.

```ts
node.d.ts (simplified excerpt)

declare module "url" {
  export interface Url {
    protocol?: string;
    hostname?: string;
    pathname?: string;
  }
  export function parse(urlStr: string, parseQueryString?, slashesDenoteHost?): Url;
}

declare module "path" {
  export function normalize(p: string): string;
  export function join(...paths: any[]): string;
  export var sep: string;
}
```

`/// <reference> node.d.ts`를 사용하고 `import url = require("url")`을 사용하여 모듈을 로드할 수 있다.

```ts
/// <reference path="node.d.ts"/>
import * as URL from "url";
let myUrl = URL.parse("http://www.typescriptlang.org");
```

#### Shorthand ambient modules

새 모듈을 사용하기 전에 선언을 작성하는데 시간을 허비하지 않으려면 shorthand 선언을 사용할 수 있다.

```ts
declare module "hot-new-module";
```

shorthand 모듈을 불러오면 `any` 타입이된다.

```ts
import x, { y } from "hot-new-module";
x(y);
```

#### Wildcard module declarations

SystemJS 및 AMD와 같은 일부 모듈로더는 JavaScript가 아닌 콘텐츠를 가져올 수 있다.
이때, 특수 로딩의 의미를 나타내기 위해 접두/접미사를 사용한다.

이러한 경우를 다루기 위해 와일드카드 모듈 선언을 사용할 수 있다.

```ts
declare module "*!text" {
  const content: string;
  export default content;
}
// Some do it the other way around.
declare module "json!*" {
  const value: any;
  export default value;
}
```

이제 `*!text` 또는 `json!*`과 일치하는 항목을 가져올 수 있다.

```ts
import fileContent from "./xyz.txt!text";
import data from "json!http://example.com/data.json";
console.log(data, fileContent);
```

#### UMD modules

일부 라이브러리는 많은 모듈 로더 또는 모듈 로딩 없이 사용하도록 설계되었다.(전역변수)
이를 UMD 모듈이라고 하며, 이러한 라이브러리는 가져오기 또는 전역 변수를 통해 액세스 할 수 있다.

```ts
import { isPrime } from "math-lib";
isPrime(2);
mathLib.isPrime(2); // ERROR: can't use the global definition from inside a module
```

### Guidance for structuring modules

#### 최대한 최상위 수준으로 내보내기

모듈을 사용하는 곳에서 가능한 마찰이 적어야 한다.
너무 많은 중첩 수준을 추가하는 것은 성가시므로 대상을 구조화 하는 방법에 신중해야 한다.

모듈에서 네임스페이스를 내보내는 것은 중첩 레이어를 많이 만드는 것이다.네
네임스페이스는 모듈이 사용될 때 추가적인 간접 참조를 추가하므로 일반적으로 귀찮고 불필요하다.

내보내기를 한 정적 메소드에도 비슷한 문제가 있다.
클래스가 중첩 레이어를 추가하므로, 단순히 helper 함수를 내보내는 것을 고려해야 한다.

#### 하나의 클래스나 함수만 내보내는 경우 `export default`를 사용

```ts
// MyClass.ts
export default class SomeType {
  constructor() { ... }
}

// MyFunc.ts
export default function getThing() { return "thing"; }

// Consumer.ts
import t from "./MyClass";
import f from "./MyFunc";
let x = new t();
console.log(f());
```

기본 내보내기를 사용하면 모듈을 사용하는 곳에서는 원하는 대로 유형을 지정할 수 있으며 개체의 이름을 위한 추가 문자열이 절약된다.

#### 여러 개체를 내보내는 경우 최상위 수준에 모두 배치하여야 함

```ts
// MyThings.ts
export class SomeType {
  /* ... */
}
export function someFunc() {
  /* ... */
}
```

반대로 불러올 때는 명시적으로 이름을 나열한다

```ts
// Consumer.ts
import { SomeType, someFunc } from "./MyThings";
let x = new SomeType();
let y = someFunc();
```

#### 많은 수의 항목을 가져오는 경우 네임스페이스 가져오기 패턴을 사용

```ts
// MyLargeModule.ts
export class Dog { ... }
export class Cat { ... }
export class Tree { ... }
export class Flower { ... }

// Consumer.ts
import * as myLargeModule from "./MyLargeModule.ts";
let x = new myLargeModule.Dog();
```

#### 모듈에서 네임스페이스를 사용하지 않음

모듈 기반의 구성으로 변경하는 경우 내보내기에 네임스페이스로 레이어를 추가하는 경향이 있다.
모듈은 자체 범위를 가지며 내보낸 선언만 모듈 외부에서 볼 수 있다.

네임스페이스는 논리적으로 관련된 개체와 유형을 그룹화 할 때 편리하다.

네임스페이스는 전역 범위에서 이름이 충돌하지 않도록 하는 것이 중요하다.
이름이 같지만 네임스페이스가 다른 경우는 가능하다.
하지만, 이는 모듈의 문제가 아니다. 모듈 내에서 동일한 이름을 가진 두개의 객체를 가져서는 안된다.

## Namespaces

### 네임스페이스 소개

내부모듈은 네임스페이스로 명명된다(TypeScript 1.5이후)
내부모듈을 선언할 때 `module` 키워드가 사용된 곳이면 어디나 `namespace` 키워드를 사용할 수 있지만,
`module` 키워드 대신 사용해야 비슷한 이름으로 overloading 함으로써 사용자를 혼란스럽게 하는 상황을 방지할 수 있다.

### Namespacing

```ts
namespace Validation {
  export interface StringValidator {
    isAcceptable(s: string): boolean;
  }

  const lettersRegexp = /^[A-Za-z]+$/;
  const numberRegexp = /^[0-9]+$/;

  export class LettersOnlyValidator implements StringValidator {
    isAcceptable(s: string) {
      return lettersRegexp.test(s);
    }
  }

  export class ZipCodeValidator implements StringValidator {
    isAcceptable(s: string) {
      return s.length === 5 && numberRegexp.test(s);
    }
  }
}

// Some samples to try
let strings = ["Hello", "98052", "101"];

// Validators to use
let validators: { [s: string]: Validation.StringValidator; } = {};
validators["ZIP code"] = new Validation.ZipCodeValidator();
validators["Letters only"] = new Validation.LettersOnlyValidator();

// Show whether each string passed each validator
for (let s of strings) {
  for (let name in validators) {
    console.log(`"${ s }" - ${ validators[name].isAcceptable(s) ? "matches" : "does not match" } ${ name }`);
  }
}
```

### Splitting Across Files

애플리케이션이 커짐에 따라 코드를 여러 파일로 분할하여 유지보수성을 높이려고 한다.

#### Multi-file namespaces

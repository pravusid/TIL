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

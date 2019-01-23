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

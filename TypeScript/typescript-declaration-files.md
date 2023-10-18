# TypeScript Declaration Files (`d.ts`)

템플릿: <https://www.typescriptlang.org/docs/handbook/declaration-files/templates.html>

## 라이브러리 구조

선언파일 구조는 라이브러리가 어떻게 사용되는지에 달려있다.
자바스크립트에서 사용할 라이브러리를 제공하는 방법에는 여러가지가 있으며 이를 위해 선언파일을 작성해야 한다.

### 라이브러리 종류 식별

라이브러리의 구조를 확인하는 것은 선언파일을 작성하는 첫 단계이다.

#### 글로벌 라이브러리

글로벌 라이브러리는 전역범위(`import`를 사용하지 않고)에서 접근할 수 있는 라이브러리이다.

많은 라이브러리가 하나 이상의 전역변수를 노출한다.
예를 들어, jQuery를 사용했다면 단순히 `$` 변수를 참조하여 변수를 사용할 수 있다.

```ts
$(() => {
  console.log('hello!');
});
```

일반적으로 글로벌 라이브러리 문서에서 HTML 스크립트 태그로 어떻게 라이브러리를 사용하는지 볼 수 있다.

```html
<script src="http://great.cdn.for/someLib.js"></script>
```

오늘날 세계적으로 가장 많이 접근하는 라이브러리는 실제로 UMD 라이브리로 작성된다.
UMD 라이브러리 설명서는 글로벌 라이브러리 설명서와 구별하기 어렵다.
전역 선언 파일을 작성하기 전에 라이브러리가 실제로 UMD가 아닌지 확인해야한다.

##### 코드에서 글로벌 라이브러리 여부 확인

글로벌 라이브러리 코드는 일반적으로 매우 단순하다.
글로벌 "Hello, world" 라이브러리는 아마 이렇게 보일 것이다.

```ts
function createGreeting(s) {
  return 'Hello, ' + s;
}
```

또는

```ts
window.createGreeting = function (s) {
  return 'Hello, ' + s;
};
```

글로벌 라이브러리 코드를 살펴보면 일반적으로 다음 내용을 볼 수 있다.

- Top-level var statements or function declarations
- One or more assignments to window.someName
- Assumptions that DOM primitives like document or window exist

다음은 보이지 않는다

- Checks for, or usage of, module loaders like require or define
- CommonJS/Node.js-style imports of the form var fs = require("fs");
- Calls to define(...)
- Documentation describing how to require or import the library

##### 글로벌 라이브러리 예제

일반적으로 글로벌 라이브러리를 UMD 라이브러리로 변환하는 것은 쉽기 때문에 인기 있는 라이브러리들도 글로벌 스타일로 작성되지 않는다.
그러나 작거나 DOM이 필요하거나 종속성이 없는 라이브러리는 글로벌 스타일로 작성될 수 있다.

##### 글로벌 라이브러리 템플릿

[`global.d.ts`](https://www.typescriptlang.org/docs/handbook/declaration-files/templates/global-d-ts.html) 파일은 글로벌 라이브러리 `myLib`를 정의한다.

이름 충돌방지 각주를 확인해야 한다.

#### 모듈러 라이브러리

일부 라이브러리는 모듈로더 환경에서만 작동한다.
예를 들어 `express`는 Node.js에서만 작동하며 CommonJS `require` 함수를 사용하여 로드해야 한다.

ES6, CommonJS 및 RequireJS는 모듈 가져오기와 비슷한 개념을 가지고 있다.
예를 들어, JavaScript CommonJS(Node.js)에서

```js
var fs = require('fs');
```

TypeScript나 ES6에서 `import` 키워드를 동일한 목적으로 사용할 수 있다

```ts
import fs = require('fs');
```

일반적으로 모듈러 라이브러리 문서에 다음 중 하나가 포함되어 있다

```ts
var someLib = require('someLib');
// 또는
define(..., ['someLib'], function(someLib) {
  // ...
});
```

글로벌 모듈과 마찬가지로 UMD 모듈의 문서에서 이러한 예제를 확인할 수 있다

##### 코드에서 모듈 라이브러리 식별

모듈러 라이브러리는 일반적으로 다음 중 일부를 포함한다

- Unconditional calls to require or define
- Declarations like import \* as a from 'b'; or export c;
- Assignments to exports or module.exports

매우 드물게 다음을 포함한다

- Assignments to properties of window or global

##### 모듈러 라이브러리 예제

`express`, `request`와 같은 Node.js 라이브러리가 module family이다.

#### UMD

UMD 모듈은 (가져오기를 통한) 모듈 또는 (모듈 로더가 없는 환경으로 실해될 때) 글로벌로 사용할 수 있다.

Moment.js와 같은 인기있는 라이브러리는 이 방식으로 작성되었다.
예를 들어 Node.js 또는 RequireJS를 사용하면 다음과 같이 작성할 수 있다.

```ts
import moment = require('moment');
console.log(moment.format());
```

바닐라 브라우저 환경에서는 다음처럼 작성할 수 있다

```ts
console.log(moment.format());
```

##### UMD 라이브러리 식별

UMD 모듈은 모듈 로더 환경유무를 확인한다.

```ts
(function (root, factory) {
  if (typeof define === "function" && define.amd) {
    define(["libName"], factory);
  } else if (typeof module === "object" && module.exports) {
    module.exports = factory(require("libName"));
  } else {
    root.returnExports = factory(root.libName);
  }
}(this, function (b) {
```

라이브러리 코드, 특히 파일 상단의 `typeof define`, `typeof window` 또는 `typeof module`에 대한 테스트가 표시되면 거의 항상 UMD 라이브러리이다.

UMD 라이브러리 문서는 Node.js에서 `require`, 브라우저에서 `<script>`태그를 사용하는 형식을 자주 보여준다.

##### UMD 라이브러리 예제

JQuery, Moment.js, lodash와 같은 인기 라이브러리는 UMD 패키지로 제공된다.

##### 템플릿

모듈에 사용할 수 있는 템플릿은 `module.d.ts`, `module-class.d.ts`, `module-function.d.ts` 세 가지이다.

모듈이 함수처럼 호출된다면 `module-function.d.ts`를 사용한다

```ts
var x = require('foo');
// Note: calling 'x' as a function
var y = x(42);
```

모듈을 `new` 키워드로 생성할 수 있는 경우 `module-class.d.ts`를 사용한다

```ts
var x = require('bar');
// Note: using 'new' operator on the imported variable
var y = new x('hello');
```

모듈이 함수처럼 호출가능하거나 `new` 키워드로 생성가능하지 않다면 `module.d.ts`를 사용한다.

#### Module plugin 또는 UMD plugin

모듈 플러그인은 모듈(UMD 또는 모듈)의 모양을 변경한다.
예를 들어, Moment.js에서 `moment-range`는 `moment` 객체에 새로운 `range` 메소드를 추가한다.

선언 파일을 작성하기 위해 변경중인 모듈이 일반 모듈인지 또는 UMD 모듈인지에 관계없이 동일한 코드를 작성한다.

##### plugin 템플릿

`module-plugin.d.ts` 템플릿을 사용한다

#### Global plugin

글로벌 플러그인은 일부 글로벌의 형태를 변경하는 글로벌 코드이다
global-modifying 모듈과 마찬가지로 런타임 충돌가능성이 높아진다.

예를 들어, 일부 라이브러리는 `Array.prototype` 또는 `String.prototype`에 새로운 함수를 추가한다.

##### Global plugin 식별

글로벌 플러그인은 일반적으로 문서에서 쉽게 식별할 수 있다

```ts
var x = 'hello, world';
// Creates new methods on built-in types
console.log(x.startsWithHello());

var y = [1, 2, 3];
// Creates new methods on built-in types
console.log(y.reverseAndSort());
```

##### Global plugin 템플릿

`global-plugin.d.ts` 템플릿을 사용한다

#### Global-modifying modules

Global-modifying 모듈은 불러올 때 전역범위의 기존 값을 변경한다.
예를 들어, 불러오면 `String.prototype`에 새 멤버를 추가하는 라이브러리가 있을 수 있다.
이는 런타임 충돌 가능성 때문에 다소 위험하지만 이에 대한 선언파일을 작성할 수는 있다.

##### Global-modifying 모듈 식별

Global-modifying 모듈은 일반적으로 문서에서 쉽게 식별할 수 있다.
일반적으로 글로벌 플러그인과 비슷하지만 효과를 활성화하려면 `require` 호출이 필요하다.

```ts
// 'require' call that doesn't use its return value
var unused = require('magic-string-time');
/* or */
require('magic-string-time');

var x = 'hello, world';
// Creates new methods on built-in types
console.log(x.startsWithHello());

var y = [1, 2, 3];
// Creates new methods on built-in types
console.log(y.reverseAndSort());
```

##### Global-modifying 모듈 템플릿

`global-modifying-module.d.ts` 템플릿을 사용한다

### Consuming Dependencies

라이브러리에 있을 수 있는 여러 종류의 종속성이 있다.
이 섹션에서는 이를 선언 파일로 가져오는 방법을 알아본다.

#### 글로벌 라이브러리에 대한 의존성

라이브러리가 전역 라이브러리에 의존하는 경우 `reference types` 디렉티브를 사용한다.

```ts
/// <reference types="someLib" />
function getThing(): someLib.thing;
```

#### 모듈에 대한 의존성

라이브러리가 모듈에 의존하는 경우 `import` 문을 사용한다

```ts
import * as moment from 'moment';
function getThing(): moment;
```

#### UMD 라이브러리에 대한 의존성

##### From a Global Library

만약 글로벌 라이브러리가 UMD 모듈에 의존한다면 `reference types` 디렉티브를 사용한다.

```ts
/// <reference types="moment" />
function getThing(): moment;
```

##### From a Module or UMD Library

만약 모듈이나 UMD 라이브러리가 UMD 라이브러리에 의존한다면, `import` 문을 사용한다.

```ts
import * as someLib from 'someLib';
```

> UMD 라이브러리에 의존성을 선언하기 위해 `reference` 디렉티브를 사용하지 않는다

### Footnotes

#### 이름 충돌 피하기

전역 선언 파일을 작성할 때 전역 범위에서 많은 타입을 정의할 수 있다.
그러나 많은 선언 파일이 프로젝트에 있을 때 해결할 수 없는 이름 충돌이 발생할 수 있다.

따라야 할 간단한 규칙은 라이브러리가 정의하는 글로별 변수가 무엇이든 간에 하나의 네임스페이스 타입을 지정하는 것이다.

예를 들어, 라이브러리가 전역 값 'cats'를 정의하면

```ts
declare namespace cats {
  interface KittySettings {}
}
```

다음 처럼 하지 않는다

```ts
// at top-level
interface CatsKittySettings {}
```

또한 이 방식은 선언파일을 사용자들을 방해하지 않고 라이브러리를 UMD로 전환할 수 있음을 보장한다.

##### 모듈 플러그인에 대한 ES6의 영향

일부 플러그인은 기존 모듈에서 최상위 내보내기를 추가하거나 수정한다.
CommonJS 및 다른 로더에서는 이것이 허용되지만 ES6 모듈은 불변으로 간주되므로 이 패턴이 불가능하다.

TypeScript는 로더에 의존하지 않기 때문에 컴파일 타임에 정책에 대한 강제사항은 없지만 ES6 모듈 로더로 전환하려는 개발자는 이 사항을 알고 있어야 한다.

##### 모듈 호출 시그니처에 대한 ES6의 영향

여러 인기 라이브러리는 불러올 때 호출 가능 함수로 자신을 노출한다.
예를 들어, 일반적인 express 사용법은 다음과 같다.

```ts
import exp = require('express');
var app = exp();
```

ES6 모듈 로더에서 최상위 오브젝트(위에서 `exp`로 불러온)는 프로퍼티만 가질 수 있다.
최상위 모듈 객체는 절대 호출할 수 없다(never callable).

일반적인 해결책은 호출가능(callable)/생성가능(constructable) 객체를 기본으로 내보내는 것이다.
일부 모듈 로더 shim은 이 상황을 감지하여 최상위 객체를 기본 내보내기로 변경한다.

##### 라이브러리 파일 레이아웃

선언 파일에 대한 레이아웃은 라이브러리 레이아웃과 동일해야 한다.

라이브러리는 다음처럼 여러 모듈로 구성될 수 있다

```ts
myLib
  +---- index.js
  +---- foo.js
  +---- bar
         +---- index.js
         +---- baz.js
```

이는 다음처럼 불러올 수 있다

```ts
var a = require('myLib');
var b = require('myLib/foo');
var c = require('myLib/bar');
var d = require('myLib/bar/baz');
```

선언 파일 구조는 다음과 같아야 한다

```ts
@types/myLib
  +---- index.d.ts
  +---- foo.d.ts
  +---- bar
         +---- index.d.ts
         +---- baz.d.ts
```

## 예시 (By Example)

### Global Variables

전역 변수 `foo`는 존재하는 위젯 수를 포함한다

```ts
console.log('Half the number of widgets is ' + foo / 2);
```

> `declare var`를 사용하여 변수를 선언한다. 변수가 읽기 전용이면 `declare const`를 사용할 수 있다.
> 변수가 블록 범위인 경우 `declare let`을 사용할 수도 있다.

```ts
/** The number of widgets present */
declare var foo: number;
```

### Global Functions

`greet` 함수를 호출하여 사용자에게 인사를 한다

```ts
greet('hello, world');
```

> `declare funtion`을 사용하여 함수를 선언한다

```ts
declare function greet(greeting: string): void;
```

### Objects with Properties

전역 변수 `myLib`에는 인사 생성을 위한 `makeGreeting` 함수와 지금 까지 작성된 인사말 수를 나타내는 `numberOfGreetings` 프로퍼티가 있다.

```ts
let result = myLib.makeGreeting('hello, world');
console.log('The computed greeting is:' + result);

let count = myLib.numberOfGreetings;
```

> `declare namespace`를 사용하여 dotted notation으로 접근하는 타입이나 값을 표기한다.

```ts
declare namespace myLib {
  function makeGreeting(s: string): string;
  let numberOfGreetings: number;
}
```

### Overloaded Functions

`getWidget` 함수는 숫자를 받아서 위젯을, 문자열을 받아서 위젯 배열을 반환한다.

```ts
let x: Widget = getWidget(43);
let arr: Widget[] = getWidget('all of them');
```

> 다음과 같이 선언한다

```ts
declare function getWidget(n: number): Widget;
declare function getWidget(s: string): Widget[];
```

### Reusable Types (Interfaces)

인사말을 지정할 때 `GreetingSettings` 객체를 전달해야 한다.
이 객체에는 다음과 같은 프로퍼티들이 있다.

1. greeting: Mandatory string
2. duration: Optional length of time (in milliseconds)
3. color: Optional string, e.g. ‘#ff00ff’

```ts
greet({
  greeting: 'hello world',
  duration: 4000,
});
```

> 프로퍼티가 있는 타입을 정의하기 위해서 `interface`를 사용한다

```ts
interface GreetingSettings {
  greeting: string;
  duration?: number;
  color?: string;
}

declare function greet(setting: GreetingSettings): void;
```

### Reusable Types (Type Aliases)

인사말을 받는 곳에서 `문자열`, `문자열`을 반환하는 함수 또는 `Greeter` 인스턴스를 사용할 수 있다.

```ts
function getGreeting() {
  return 'howdy';
}
class MyGreeter extends Greeter {}

greet('hello');
greet(getGreeting);
greet(new MyGreeter());
```

> 타입 별칭을 사용하여 타입에 대한 shorthand를 만들 수 있다

```ts
type GreetingLike = string | (() => string) | MyGreeter;

declare function greet(g: GreetingLike): void;
```

### Organizing Types

`greeter` 객체는 파일에 기록하거나 경고를 표시할 수 있다.
`.log(...)`에 LogOptions을 사용하고 `.alert(...)`에 AlertOptions를 사용할 수 있다.

```ts
const g = new Greeter('Hello');
g.log({ verbose: true });
g.alert({ modal: false, title: 'Current Greeting' });
```

> organize types를 위해 namespace를 사용한다

```ts
declare namespace GreetingLib {
  interface LogOptions {
    verbose?: boolean;
  }
  interface AlertOptions {
    modal: boolean;
    title?: string;
    color?: string;
  }
}
```

> 하나의 선언에 중첩 namespace를 만들 수도 있다

```ts
declare namespace GreetingLib.Options {
  // Refer to via GreetingLib.Options.Log
  interface Log {
    verbose?: boolean;
  }
  interface Alert {
    modal: boolean;
    title?: string;
    color?: string;
  }
}
```

### Classes

`Greeter` 객체를 인스턴스화 하여 greeter를 생성하거나, 상속하여 사용자 정의 greeter를 만들 수 있다

```ts
const myGreeter = new Greeter('hello, world');
myGreeter.greeting = 'howdy';
myGreeter.showGreeting();

class SpecialGreeter extends Greeter {
  constructor() {
    super('Very special greetings');
  }
}
```

> `declare class`를 사용하여 클래스 또는 클래스와 유사한 객체를 나타낸다. 클래스는 생성자 뿐만 아니라 프로퍼티 및 메소드를 가질 수 있다.

```ts
declare class Greeter {
  constructor(greeting: string);

  greeting: string;
  showGreeting(): void;
}
```

## 해야 할 것과 하지 말아야 할 것

### 일반 타입

`Number`, `String`, `Boolean`, `Symbol` and `Object`

위와 같은 타입을 사용하지 말아야 한다.
이러한 타입은 JavaScript 코드에서 거의 사용되지 않는 non-primitive boxed object를 나타낸다.

대신 `number`, `string`, `boolean`, `symbol`을 사용한다.
`Object` 대신 non-primitive `object` 타입을 사용한다. (TS 2.2에서 추가됨)

```ts
/* WRONG */
function reverse(s: String): String;

/* OK */
function reverse(s: string): string;
```

#### Generics

타입 파라미터를 사용하지 않는다면 제네릭을 사용하지 않아야 한다.

<https://github.com/Microsoft/TypeScript/wiki/FAQ#why-doesnt-type-inference-work-on-this-interface-interface-foot-->

### 콜백 타입

#### 콜백의 반환타입

값이 무시 될 콜백에 대해 반환 타입 `any`를 사용하지 않아야 한다.
대신 값이 무시될 콜백에는 반환 타입 `void`를 사용한다.

```ts
/* WRONG */
function fn(x: () => any) {
  x();
}

/* OK */
function fn(x: () => void) {
  x();
}
```

> `void`를 사용하면 실수로 `x`의 반환 값을 확인 되지 않은 방식으로 사용하는 것을 방지할 수 있으므로 더 안전하다.

```ts
function fn(x: () => void) {
  var k = x(); // oops! meant to do something else
  k.doSomething(); // error, but would be OK if the return type had been 'any'
}
```

#### 콜백에서의 선택적 파라미터

실제로 사용하지 않는다면 콜백에서 선택적 파라미터를 사용하지 말아야 한다

이는 매우 구체적인 의미가 있다.
완료 콜백은 1개 혹은 2개의 인수로 호출될 수 있다.
원래 의도는 콜백이 `elapsedTime` 매개변수를 신경쓰지 않을 수도 있다는 것이었지만 이를 위해 선택적 파라미터를 사용할 필요는 없다.

더 적은 인수를 허용하는 콜백을 제공하는 것이 보다 합리적이다.

```ts
/* WRONG */
interface Fetcher {
  getObject(done: (data: any, elapsedTime?: number) => void): void;
}

/* OK */
interface Fetcher {
  getObject(done: (data: any, elapsedTime: number) => void): void;
}
```

#### 오버로드와 콜백

콜백만 다른 별도의 overload를 작성하지 않아야 한다.

> 콜백은 항상 매개변수를 무시할 수 있으므로 더 짧은 overload는 필요하지 않다.
> 짧은 콜백이 먼저 선언되면 잘못된 타입의 함수가 첫번 째 오버로드와 일치하게 되어 우선 전달된다.

```ts
/* WRONG */
declare function beforeAll(action: () => void, timeout?: number): void;
declare function beforeAll(action: (done: DoneFn) => void, timeout?: number): void;

/* OK */
declare function beforeAll(action: (done: DoneFn) => void, timeout?: number): void;
```

### 함수 오버로드

#### Ordering

보다 구체적인 오버로드 이전에 일반적인 오버로드를 선언하지 않아야 한다.
좀 더 특수한 시그니처 뒤에 일반적인 시그니처를 둔다.

> TypeScript는 함수 호출을 처리할 때 처음으로 일치하는 오버로드를 선택한다.
> 이전 오버로드가 나중의 오버로드보다 "더 일반적" 일 때, 이후 오버로드는 호출될 수 없다

```ts
/* WRONG */
declare function fn(x: any): any;
declare function fn(x: HTMLElement): number;
declare function fn(x: HTMLDivElement): string;
var myElem: HTMLDivElement;
var x = fn(myElem); // x: any, wat?

/* OK */
declare function fn(x: HTMLDivElement): string;
declare function fn(x: HTMLElement): number;
declare function fn(x: any): any;
var myElem: HTMLDivElement;
var x = fn(myElem); // x: string, :)
```

#### 선택적 파라미터 사용

후행 매개변수만 다른 여러개의 오버로드를 작성하면 안된다

```ts
/* WRONG */
interface Example {
  diff(one: string): number;
  diff(one: string, two: string): number;
  diff(one: string, two: string, three: boolean): number;
}

/* OK */
interface Example {
  diff(one: string, two?: string, three?: boolean): number;
}
```

이러한 방식은 모든 오버로드가 동일한 반환 타입을 가질때만 가능하다

> TypeScript는 소스의 인수들로 대상의 시그니처가 호출될 수 있는지를 확인하여 시그니처 호환성을 해결한다.

아래의 코드는 선택적 매개변수를 사용하여 시그니처가 올바르게 작성된 경우에만 버그를 노출한다.

```ts
function fn(x: (a: string, b: number, c: number) => void) {}
var x: Example;
// When written with overloads, OK -- used first overload
// When written with optionals, correctly an error
fn(x.diff);
```

> 함수를 사용하는 곳에서 "strict null checking"을 사용하는 경우 지정되지 않는 매개 변수는 JavaScript의 `undefined`로 나타나므로
> 선택적 인수가 있는 함수에 명시적으로 `undefined` 값을 전달하는 것이 좋다.

아래의 코드는 엄격한 null 확인에서 허용되어야 한다

```ts
var x: Example;
// When written with overloads, incorrectly an error because of passing 'undefined' to 'string'
// When written with optionals, correctly OK
x.diff('something', true ? undefined : 'hour');
```

#### Union 타입 사용

하나의 인수 위치에서 타입만 다른 오버로드를 사용하면 안된다.
가능하다면 union 타입을 사용한다.

```ts
/* WRONG */
interface Moment {
  utcOffset(): number;
  utcOffset(b: number): Moment;
  utcOffset(b: string): Moment;
}

/* OK */
interface Moment {
  utcOffset(): number;
  utcOffset(b: number | string): Moment;
}
```

> 이는 값을 함수에 전달하려는 사람에게 중요하다

```ts
function fn(x: string): void;
function fn(x: number): void;
function fn(x: number | string) {
  // 별도의 오버로드로 작성되었다면, 부정확함 -> 오류
  // union 타입으로 작성되었다면, 정확함 -> 허용됨
  return moment().utcOffset(x);
}
```

## Deep Dive

모듈 혹은 UMD에 중점을 둔 내용임

### Key Concepts

TypeScript 작동방식에 대한 몇 가지 주요개념을 이해하면 모든 형태의 정의를 만드는 방법을 이해할 수 있다

#### 타입

- 타입 별칭 선언 (`type sn = number | string;`)
- 인터페이스 선언 (`interface I { x: number[]; }`)
- 클래스 선언 (`class C {}`)
- Enum 선언 (`enum E { A, B, C }`)
- 타입을 가리키는 가져오기 선언

각 타입 선언은 새로운 타입 이름을 생성한다

#### 값

값은 표현식에서 참조할 수 있는 런타임 이름이다. 예를 들면 `let x = 5;`는 `x`라는 값을 생성한다.

- `let`, `const`, `var` 선언
- 네임스페이스 또는 모듈의 값을 포함한 선언
- Enum 선언
- 클래스 선언
- 값을 가리키는 가져오기 선언
- 함수 선언

#### 네임스페이스

네임스페이스에 타입이 존재할 수 있다. 예를 들어 `let x: A.B.C` 선언이 있는 경우 타입 `C`는 `A.B` 네임스페이스에서 온 것이다.

이 구별은 중요하다. `A.B`는 타입 또는 값일 필요가 없다.

### 단순 조합: One name, multiple meanings

이름 `A`가 주어지면 `A`에 대한 세 가지 다른 의미인 타입, 값, 네임스페이스를 찾을 수 있다.

이름이 해석되는 방법은 컨텍스트에 따라 다르다.
예를 들어, 선언 `m: A.A = A;`에서 `A`는 네임스페이스로 사용된 다음 타입이름으로 사용되고 값으로도 사용된다.

혼란스러울 수 있겟지만 지나친 오버로딩이 아니라면 편리하게 사용할 수 있다.

#### Built-in Combinations

예를 들어, 클래스는 타입 및 값 모두 될 수 있다.
`class C {}` 선언은 클래스 인스턴스를 나타내는 타입 `C`와 클래스 생성자 함수를 나타내는 값 `C`를 만든다.

Enum 선언도 비슷하게 동작한다.

#### User Combinations

`foo.d.ts` 모듈 파일을 작성했다고 가정하자

```ts
export var SomeVar: { a: SomeType };
export interface SomeType {
  count: number;
}
```

이를 사용한다

```ts
import * as foo from './foo';
let x: foo.SomeType = foo.SomeVar.a;
console.log(x.count);
```

위 코드는 잘 작동하지만 `SomeType`과 `SomeVar`는 같은 이름이므로 밀접한 연관관계가 있다고 생각할 수 있다.
동일한 이름 `Bar`를 사용하여 두 가지 다른 객체(값과 타입)를 결합할 수 있다.

```ts
export var Bar: { a: Bar };
export interface Bar {
  count: number;
}
```

코드를 사용할 때 구조분해를 사용하기 좋다

```ts
import { Bar } from './foo';
let x: Bar = Bar.a;
console.log(x.count);
```

여기에서도 `Bar`를 타입이면서 값으로 사용했다. `Bar` 값을 `Bar` 타입으로 선언하지 않아도 독립적으로 작동한다.

### Advanced Combinations

일부 선언은 여러 선언에 걸쳐 결합할 수 있다.
예를 들어, `class C {}`와 `interface C {}`는 공존할 수 있으며 둘다 `C` 타입에 프로퍼티를 제공한다.

충돌을 일으키지 않으면 허용된다.
일반적인 규칙은 `namespace s`로 선언되지 않는 한 동일한 이름의 값 끼리는 항상 충돌하며,
타입은 타입 별칭 선언 `type s = string`으로 선언되면 충돌하지만, 네입스페이스는 절대 충돌하지 않는다.

#### `interface`를 사용하여 추가

다른 `interface` 선언으로 `interface`에 멤버를 추가할 수 있다

```ts
interface Foo {
  x: number;
}
// ... elsewhere ...
interface Foo {
  y: number;
}
let a: Foo = ...;
console.log(a.x + a.y); // OK
```

이는 클래스에서도 적용된다

```ts
class Foo {
  x: number;
}
// ... elsewhere ...
interface Foo {
  y: number;
}
let a: Foo = ...;
console.log(a.x + a.y); // OK
```

인터페이스를 사용하여 타입 별칭 `type s = string;`에 추가할 수 없다

#### `namespace`를 사용하여 추가

네임스페이스 선언을 사용하면 충돌을 일으키지 않는 방식으로 새로운 타입, 값, 네임스페이스를 추가할 수 있다

예를 들어, 클래스에 static 멤버를 추가할 수 있다

```ts
class C {}
// ... elsewhere ...
namespace C {
  export let x: number;
}
let y = C.x; // OK
```

이 예제에서 네임스페이스 선언을 작성할 때 까지 네임스페이스 `C`는 없었다.
이후 선언된 네임스페이스 `C`는 클래스에 의해 생성된 `C`의 값이나 타입과 충돌하지 않는다.

마지막으로 네임스페이스 선언을 사용하여 다양한 병합을 수행할 수 있다.

```ts
namespace X {
  export interface Y {}
  export class Z {}
}

// ... elsewhere ...
namespace X {
  export var Y: number;
  export namespace Z {
    export class C {}
  }
}
type X = string;
```

이 예제에서 첫 블록은 다음 이름을 생성한다

- 값 `X` (네임스페이스 선언에 값 `Z`가 포함되어 있으므로)
- 네임스페이스 `X` (네임스페이스 선언에 타입 `Y`가 포함되어 있으므로)
- 네임스페이스 `X`의 타입 `Y`
- 네임스페이스 `X`의 타입 `Z` (클래스 인스턴스 형태)
- 값 `X`의 프로퍼티 값 `Z` (클래스 생성자 함수)

두 번째 블록은 다음 이름을 생성한다

- 값 `X`의 프로퍼티인 값 `Y`(`number` 타입)
- 네임스페이스 `Z`
- 값 `X`의 프로퍼티인 값 `Z`
- 네임스페이스 `X.Z`의 타입 `C`
- 값 `X.Z`의 값 `C`
- 타입 `X`

### Using with `export =` or `import`

`export`와 `import` 선언은 내보내기나 가져오기 대상 전체(_all meanings_)를 가져온다

## Publishing

선언 파일을 npm에 게시할 수 있는 주요 방법은 다음과 같다

- npm 패키지에 번들링 하거나
- npm의 [@types](https://www.npmjs.com/~types)에 게시

패키지가 TypeScript로 작성된 경우라면 컴파일러 옵션에서 `--declaration` 플래그를 사용하여 첫 번째 방법을 사용하면 된다.

### npm 패키지에 선언 포함

패키지에 main `.js` 파일이 있는경우 `package.json` 파일에도 기본 선언파일을 표시해야 한다.
동봉된 선언파일을 카리키는 `types` 속성을 설정한다.

```json
{
  "name": "awesome",
  "main": "./lib/main.js",
  "types": "./lib/main.d.ts"
}
```

`"typings"` 필드는 `"types"` 필드와 같으며 대신 사용할 수 있다.

또한 기본 선언파일의 이름이 `index.d.ts`이고 패키지 루트 (`index.js` 옆)에 있는 경우 `"types"` 속성을 표기할 필요는 없지만 되도록 표기하는 것이 좋다.

#### Dependencies

모든 의존성은 npm에 의해서 관리된다. 의존하는 패키지가 `package.json`에 제대로 표시되어 있는지 확인하여야 한다.

#### Red flags

선언파일에 `/// <reference path="..." />`를 사용하면 안된다. 대신 `/// <reference types="..." />`을 사용하여야 한다.

Type definitions이 다른 패키지에 의존하는 경우 현재 작성하는 패키지에 결합하지 않고 별개의 파일로 유지해야 한다

### `@types`에 게시

[types-publisher tool](https://github.com/Microsoft/types-publisher)을 사용하여
[DefinitelyTyped](https://github.com/DefinitelyTyped/DefinitelyTyped)에 풀리퀘스트를 요청하면 `@types` 스코프 패키지에 자동으로 게시된다.

## Consumption

### 다운로드

npm 패키지에 선언 파일이 없다면 `@types` 패키지를 다운로드 하면된다

```sh
npm install --save @types/lodash
```

### 사용

설치한 이후 불러와서 사용하기만 하면 된다

```ts
import * as _ from 'lodash';
_.padStart('Hello TypeScript!', 20, ' ');
```

### 검색

대부분의 경우 타입 선언패키지는 npm 패키지 이름과 같으며 접두사가 `@types/`이지만 필요한 경우 <https://aka.ms/types>에서 직접 찾을 수 있다.

## Examples

### process.env

- <https://stackoverflow.com/questions/45194598/using-process-env-in-typescript>
- <https://typescript.tv/errors/#TS2669>

`global.d.ts`

```ts
declare global {
  namespace NodeJS {
    interface ProcessEnv {
      FOO_ENV_VAR: string;
    }
  }
}

export {};
```

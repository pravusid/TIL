# TypeScript HandBook: Module System

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

`Validation.ts`

```ts
export interface StringValidator {
  isAcceptable(s: string): boolean;
}
```

`ZipCodeValidator.ts`

```ts
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

`ParseIntBasedZipCodeValidator.ts`

```ts
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

`JQuery.d.ts`

```ts
declare let $: JQuery;
export default $;
```

`App.ts`

```ts
import $ from "JQuery";
$("button.continue").html("Next Step...");
```

클래스나 함수 선언은 기본 내보내기로 직접 작성될 수 있다. 이름 선언은 선택사항이다.

`ZipCodeValidator.ts`

```ts
export default class ZipCodeValidator {
  static numberRegexp = /^[0-9]+$/;
  isAcceptable(s: string) {
    return s.length === 5 && ZipCodeValidator.numberRegexp.test(s);
  }
}
```

`Test.ts`

```ts
import validator from "./ZipCodeValidator";
let myValidator = new validator();
```

또는

`StaticZipCodeValidator.ts`

```ts
const numberRegexp = /^[0-9]+$/;
export default function(s: string) {
  return s.length === 5 && numberRegexp.test(s);
}
```

`Test.ts`

```ts
import validate from "./StaticZipCodeValidator";
let myValidator = new validator();
```

기본 내보내기는 단순히 값일 수도 있다

`OneTwoThree.ts`

```ts
export default "123";
```

`Log.ts`

```ts
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

`ZipCodeValidator.ts`

```ts
let numberRegexp = /^[0-9]+$/;
class ZipCodeValidator {
  isAcceptable(s: string) {
    return s.length === 5 && numberRegexp.test(s);
  }
}
export = ZipCodeValidator;
```

`Test.ts`

```ts
import zip = require("./ZipCodeValidator");
let validator = new zip();
```

### Code Generation for Modules

컴파일하는 동안 지정된 모듈 타겟에 따라 컴파일러는 Node.js(CommonJS), require.js(AMD), UMD, SystemJS 또는
ES6 모듈로드 시스템에 적합한 코드를 생성한다.

`SimpleModule.ts`

```ts
import m = require("mod");
export let t = m.something + 1;
```

`AMD / RequireJS SimpleModule.js`

```js
define(["require", "exports", "./mod"], function (require, exports, mod_1) {
  exports.t = mod_1.something + 1;
});
```

`CommonJS / Node SimpleModule.js`

```js
var mod_1 = require("./mod");
exports.t = mod_1.something + 1;
```

`UMD SimpleModule.js`

```js
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
```

`System SimpleModule.js`

```js
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
```

`Native ECMAScript 2015 modules SimpleModule.js`

```js
import { something } from "./mod";
export var t = something + 1;
```

### Working with Other JavaScript Libraries

타입스크립트로 작성되지 않은 라이브러리의 형태를 기술하려면, 라이브러리가 노출하는 API를 선언해야 한다.

#### Ambient Modules

최상위 수준의 내보내기 선언을 사용하여 각 모듈을 자체 `d.ts` 파일로 정의할 수 있지만, 더 큰 `d.ts` 파일로 작성하는 것이 편리하다.

`node.d.ts (simplified excerpt)`

```ts
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

`MyClass.ts`

```ts
export default class SomeType {
  constructor() { ... }
}
```

`MyFunc.ts`

```ts
export default function getThing() { return "thing"; }
```

`Consumer.ts`

```ts
import t from "./MyClass";
import f from "./MyFunc";
let x = new t();
console.log(f());
```

기본 내보내기를 사용하면 모듈을 사용하는 곳에서는 원하는 대로 유형을 지정할 수 있으며 개체의 이름을 위한 추가 문자열이 절약된다.

#### 여러 개체를 내보내는 경우 최상위 수준에 모두 배치하여야 함

`MyThings.ts`

```ts
export class SomeType {
  /* ... */
}
export function someFunc() {
  /* ... */
}
```

반대로 불러올 때는 명시적으로 이름을 나열한다

`Consumer.ts`

```ts
import { SomeType, someFunc } from "./MyThings";
let x = new SomeType();
let y = someFunc();
```

#### 많은 수의 항목을 가져오는 경우 네임스페이스 가져오기 패턴을 사용

`MyLargeModule.ts`

```ts
export class Dog { ... }
export class Cat { ... }
export class Tree { ... }
export class Flower { ... }
```

`Consumer.ts`

```ts
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
let validators: { [s: string]: Validation.StringValidator } = {};
validators["ZIP code"] = new Validation.ZipCodeValidator();
validators["Letters only"] = new Validation.LettersOnlyValidator();

// Show whether each string passed each validator
for (let s of strings) {
  for (let name in validators) {
    console.log(
      `"${s}" - ${
        validators[name].isAcceptable(s) ? "matches" : "does not match"
      } ${name}`
    );
  }
}
```

### Splitting Across Files (Multi-file namespaces)

애플리케이션이 커짐에 따라 코드를 여러 파일로 분할하여 유지보수성을 높이려고 한다.

`Validation` 네임스페이스를 여러 파일로 분리하여도 모든 파일이 한 곳에서 정의된 것처럼 사용할 수 있다.

파일 간에는 종속성이 있으므로 컴파일러에서 파일 사이 관계를 알 수 있도록 참조 태그를 추가한다.

`Validation.ts`

```ts
namespace Validation {
  export interface StringValidator {
    isAcceptable(s: string): boolean;
  }
}
```

`LettersOnlyValidator.ts`

```ts
/// <reference path="Validation.ts" />
namespace Validation {
  const lettersRegexp = /^[A-Za-z]+$/;
  export class LettersOnlyValidator implements StringValidator {
    isAcceptable(s: string) {
      return lettersRegexp.test(s);
    }
  }
}
```

`ZipCodeValidator.ts`

```ts
/// <reference path="Validation.ts" />
namespace Validation {
  const numberRegexp = /^[0-9]+$/;
  export class ZipCodeValidator implements StringValidator {
    isAcceptable(s: string) {
      return s.length === 5 && numberRegexp.test(s);
    }
  }
}
```

### Aliases

네임스페이스를 간단하게 사용하기 위한 다른 방법은 `import q = x.y.z`를 사용하여 짧은 이름을 만드는 것이다.

이는 모듈을 불러오는 `import x = require('name')` 구문과 혼동되기 쉽다.

```ts
namespace Shapes {
  export namespace Polygons {
    export class Triangle { }
    export class Square { }
  }
}

import polygons = Shapes.Polygons;
let sq = new polygons.Square(); // Same as 'new Shapes.Polygons.Square()'
```

### 네임스페이스와 JavaScript Libraries

TypeScript로 작성되지 않은 라이브러리는 외부 노출을 위한 API 선언이 필요하다.

대부분의 JavaScript 라이브러리는 몇 개의 최상위 수준 객체만 노출하므로 네임스페이스를 사용하는 것이 좋은 방식이다.

## Namespaces and Modules

### Using Namespaces

네임스페이스는 글로벌 네임스페이스에서 JavaScript 객체로 명명된다.

### Using Modules

네임스페이스와 마찬가지로 모듈에는 코드와 선언이 모두 포함될 수 있다. 가장 큰 차이점은 모듈이 의존선을 선언한다는 것이다.

또한 모듈은 모듈로더(CommonJs Require.js ...)에 종속된다.
이는 대규모 응용프로그램의 경우 장기적인 유지 관리의 이점이 있다.

### 네임스페이스와 모듈의 함정

#### `/// <reference>`-ing a module

일반적인 실수는 `import`문을 사용하는 대신 `/// <reference ... />` 구문을 사용하여 모듈 파일을 참조하려고 하는 것이다.

컴파일러는 적절한 경로로 `.ts`, `.tsx`, `.d.ts`를 찾으려고 한다.
특정 파일을 찾을 수 없으면 컴파일러는 주변 모듈 선언을 찾는다. 이는 `.d.ts` 파일에서 선언할 필요가 있다.

`myModules.d.ts`

```ts
// In a .d.ts file or .ts file that is not a module:
declare module "SomeModule" {
  export function fn(): string;
}
```

`myOtherModule.ts`

```ts
/// <reference path="myModules.d.ts" />
import * as m from "SomeModule";
```

여기에서 참조 태그는 주변 모듈에 대한 선언을 포함하는 선언 파일(`.d.ts`)을 찾을 수 있게 해준다.

#### Needless Namespacing

프로그램을 네임스페이스에서 모듈로 변환하는 경우 다음과 같이 할 수 있다.

`shapes.ts`

```ts
export namespace Shapes {
  export class Triangle { /* ... */ }
  export class Square { /* ... */ }
}
```

최상위 모듈 `Shapes`는 특별한 이유없이 `Triangle`과 `Square`를 감싸고 있다. 이는 사용시 혼란과 짜증을 불러온다.

`shapeConsumer.ts`

```ts
import * as shapes from "./shapes";
let t = new shapes.Shapes.Triangle(); // shapes.Shapes?
```

TypeScript 모듈의 핵심기능은 두 개의 다른 모듈이 동일한 범위에 이름을 제공하지 않는다는 것이다.
모듈 소비자는 할당 할 이름을 결정하기 때문에 네임스페이스에서 내보낸 심볼을 사전에 감쌀필요가 없다.

네임스페이스의 일반적인 개념은 논리적 그룹핑 구조를 제공하고 이름 충돌을 방지하는 것이다.
모듈 파일 자체는 이미 논리적 그룹이며 최상위 이름은 가져오는 코드에서 정의되므로 내보낸 객체에 추가적인 모듈 레이어를 사용할 필요가 없다.

따라서 위의 코드를 다시 써보면

`shapes.ts`

```ts
export class Triangle { /* ... */ }
export class Square { /* ... */ }
```

`shapeConsumer.ts`

```ts
import * as shapes from "./shapes";
let t = new shapes.Triangle();
```

## Module Resolution

Module Resolution은 import가 무엇을 의미하는지 파악하기 위해 컴파일러가 사용하는 프로세스이다.

`import { a } from 'moduleA'` 구문에서 `a`가 무엇을 나타내는지 알아내기 위해서 `moduleA`를 검사할 필요가 있다.

이때 `moduleA`는 `.ts` / `.tsx` 파일 중 하나 또는 `.d.ts`파일에 정의될 수 있다.

컴파일러는 불러온 모듈을 지정한 파일을 찾은 뒤 Classic 또는 Node라는 위치를 찾기위한 두 가지 전략 중 하나를 따른다.
두 방법이 작동하지 않고 모듈 이름이 상대적이지 않은 경우 컴파일러는 ambient module 선언을 찾으려고 시도한다.

마지막으로 컴파일러가 모듈을 찾을 수 없으면 오류가 출력된다.

> `error TS2307: Cannot find module 'moduleA'.`

### Relative vs. Non-relative module imports

모듈 가져오기는 모듈 참조가 상대적인지에 따라 다르게 처리된다.

상대경로 가져오기는 `/`, `./` 또는 `../`으로 시작한다.

```ts
import Entry from "./components/Entry";
import { DefaultHeaders } from "../constants/http";
import "/mod";
```

가져오기가 상대경로가 아닌 경우는 다음과 같다

```ts
import * as $ from "jquery";
import { Component } from "@angular/core";
```

상대경로 가져오기는 런타임시 상대위치를 유지하도록 보장되는 모듈에 대해서 사용해야 한다.

상대경로가 아닌 가져오기는 `baseUrl`을 기준으로 하거나 경로 매핑을 통해 사용할 수 있다.

### Module Resolution Strategies

모듈 분석 전략에는 Node와 Classic 두 가지가 있다.
값이 지정되지 않는다면 기본적으로 `--module AMD | System | ES2015`에 대해서는 **Classic**이며, 이외의 경우 **Node**이다.

#### Classic

타입스크립트의 기본 모듈 분석전략이었다. 최근에는 이전 버전과의 호환성을 위해 제공된다.

상대경로 가져오기는 가져오는 파일과 관련하여 처리된다.

소스파일 `/root/src/folder/A.ts`에서 `import { b } from "./moduleB"`를 사용한다면 다음 순서의 조회가 발생한다.

1. `/root/src/folder/moduleB.ts`
2. `/root/src/folder/moduleB.d.ts`

그러나 **상대경로가 아닌 가져오기**의 경우 컴파일러는 가져오기를 사용한 소스의 위치부터 시작하여 일치하는 정의 파일을 찾기위해 디렉토리 트리를 찾는다.

소스파일 `/root/src/folder/A.ts`에서 `import { b } from "./moduleB"`를 사용한다면 다음 순서의 조회가 발생한다.

1. `/root/src/folder/moduleB.ts`
2. `/root/src/folder/moduleB.d.ts`
3. `/root/src/moduleB.ts`
4. `/root/src/moduleB.d.ts`
5. `/root/moduleB.ts`
6. `/root/moduleB.d.ts`
7. `/moduleB.ts`
8. `/moduleB.d.ts`

#### Node

Node 분석 전략은 Node.js 모듈 분석 메커니즘을 모방한다. <https://nodejs.org/api/modules.html#modules_all_together>

일반적으로 Node.js의 가져오기는 `require`라는 함수를 호출하여 수행된다.
Node.js의 동작은 `require`에 상대 경로 또는 비 상대경로가 있는지에 따라 달라진다.

상대경로 가져오기는 소스파일 `/root/src/folder/A.ts`에서 `var x = require("./moduleB")`를 사용한다면 다음 순서의 조회가 발생한다.

1. `/root/src/moduleB.js` 파일이 존재하면 요청한다.
2. `/root/src/moduleB` 폴더가 `"main"` 모듈을 가리키는 `package.json` 파일을 포함하고 있는지 확인한다. 만약 `{ "main": "lib/mainModule.js" }` 코드를 포함하는 `package.json` 파일을 찾는다면 Node.js는 `/root/src/moduleB/lib/mainModule.js` 파일을 참조한다.
3. `/root/src/moduleB` 폴더가 `index.js` 파일을 포함하는지 확인한다. 이 파일은 암시적으로 폴더의 main 모듈로 간주된다.

그러나 **상대경로가 아닌 가져오기**의 경우 다르게 수행된다.

node는 `node_modules`라는 특수 디렉토리에서 모듈을 검색한다.
`node_modules` 디렉토리는 현재 파일과 동일 레벨이거나 디렉토리 체인에서 상위 레벨일 수 있다.

node는 불려오려고 시도한 모듈을 찾을 때 까지 각 `node_modules`를 조사하기위해 디렉토리 체인을 따라간다.

소스파일 `/root/src/moduleA.js`가 비상대경로를 사용하고 `import var x = require ( "moduleB")`가 있는 경우 다음 순서의 조회가 발생한다.

1. `/root/src/node_modules/moduleB.js`
2. `/root/src/node_modules/moduleB/package.json` (if it specifies a "main" property)
3. `/root/src/node_modules/moduleB/index.js`

4. `/root/node_modules/moduleB.js`
5. `/root/node_modules/moduleB/package.json` (if it specifies a "main" property)
6. `/root/node_modules/moduleB/index.js`

7. `/node_modules/moduleB.js`
8. `/node_modules/moduleB/package.json` (if it specifies a "main" property)
9. `/node_modules/moduleB/index.js`

4단계와 7단계에서 디렉토리를 건너뛰었다.

`node_modules`에서 모듈을 로딩하는 과정은 다음을 참조: <https://nodejs.org/api/modules.html#modules_loading_from_node_modules_folders>

#### How TypeScript resolves modules

TypeScript는 컴파일 타임에 모듈 정의 파일을 찾기 위해 Node.js 런타임 해석 전략을 모방한다.
이를 위해 TypeScript 원본 파일 확장명(.ts .tsk .d.ts)을 Node.js 전략위에 둔다.
또한 `package.json`의 `types`필드는 `main`을 미러링할 목적으로 사용한다.

`/root/src/moduleA.ts` 파일에서 `"./moduleB"에서 import {b}`를 사용하면 `"./moduleB"`를 찾기 위해 다음 위치를 시도한다.

1. `/root/src/moduleB.ts`
2. `/root/src/moduleB.tsx`
3. `/root/src/moduleB.d.ts`
4. `/root/src/moduleB/package.json` (if it specifies a "types" property)
5. `/root/src/moduleB/index.ts`
6. `/root/src/moduleB/index.tsx`
7. `/root/src/moduleB/index.d.ts`

비슷하게 상대경로 불러오기는 Node.js의 분석전략에 따라 파일을 찾은 다음 디렉토리를 찾는다.
따라서 `/root/src/moduleA.ts` 파일에서 `"./moduleB"에서 import {b}`를 사용하면 다음 과정으로 처리된다.

1. `/root/src/node_modules/moduleB.ts`
2. `/root/src/node_modules/moduleB.tsx`
3. `/root/src/node_modules/moduleB.d.ts`
4. `/root/src/node_modules/moduleB/package.json` (if it specifies a "types" property)
5. `/root/src/node_modules/@types/moduleB.d.ts`
6. `/root/src/node_modules/moduleB/index.ts`
7. `/root/src/node_modules/moduleB/index.tsx`
8. `/root/src/node_modules/moduleB/index.d.ts`

9. `/root/node_modules/moduleB.ts`
10. `/root/node_modules/moduleB.tsx`
11. `/root/node_modules/moduleB.d.ts`
12. `/root/node_modules/moduleB/package.json` (if it specifies a "types" property)
13. `/root/node_modules/@types/moduleB.d.ts`
14. `/root/node_modules/moduleB/index.ts`
15. `/root/node_modules/moduleB/index.tsx`
16. `/root/node_modules/moduleB/index.d.ts`

17. `/node_modules/moduleB.ts`
18. `/node_modules/moduleB.tsx`
19. `/node_modules/moduleB.d.ts`
20. `/node_modules/moduleB/package.json` (if it specifies a "types" property)
21. `/node_modules/@types/moduleB.d.ts`
22. `/node_modules/moduleB/index.ts`
23. `/node_modules/moduleB/index.tsx`
24. `/node_modules/moduleB/index.d.ts`

9단계와 17단계에서 디렉토리 건너뛰기가 발생한다.

### Additional module resolution flags

프로젝트 소스 레이아웃이 출력의 레이아웃과 일치하지 않는 경우가 있다.
`.ts` 파일을 `.js` 파일로 컴파일하고 여러 종속성을 단일 출력위치로 복사하는 작업이 포함된다.

즉, 최종 결과는 런타임에 모듈 정의파일과 다른 이름을 가지거나 최종 출력의 모듈 경로가 소스파일 경로와 일치하지 않을 수 있다.

TypeScript 컴파일러에는 최종 출력을 생성하기 위해 예상되는 변환을 컴파일러에 알리는 추가 플래그가 있다.
컴파일러는 이러한 변환을 직접 수행하지 않고, 모듈 가져오기를 definition 파일로 해석하는 프로세스를 안내한다.

#### Base URL

`baseUrl`을 사용하는 것은 모듈이 런타임에 단일 디렉토리에 배포되는 AMD 모듈 로더를 사용하는 응용 프로그램에서 일반적으로 사용된다.

`baseUrl`을 설정하면 모듈을 찾을 위치를 컴파일러에게 알린다.
비상대경로의 모든 모듈을 가져올 때 `baseUrl`에 상대적이라고 가정한다.

`baseUrl` 값은 다음 중 하나로 결정된다

- `baseUrl` 커맨드 라인 인자 값 (주어진 경로가 상대경로라면 현재 디렉토리를 기반으로 계산됨)
- `tsconfig.json`의 `baseUrl` 속성 값 (주어진 경로가 상대경로인 경우 `tsconfig.json` 위치를 기반으로 계산됨)

상대경로 모듈 가져오기는 `baseUrl` 설정에 영향을 받지 않으며 항상 불러오는 파일에 대해 상대적으로 처리된다.

#### Path mapping

때때로 모듈은 `baseUrl` 아래에 직접 위치하지 않는다.
예를 들어 `jquery`에 대한 가져오기는 런타임에 `node_modules/jquery/dist/jquery.slim.min.js`로 변환된다.

로더는 mapping 구성을 사용하여 런타임시 모듈 이름을 파일에 매핑한다.

TypeScript컴파일러는 `tsconfig.json` 파일의 `path` 속성을 사용하여 매핑을 선언한다.

```json
{
  "compilerOptions": {
    "baseUrl": ".", // This must be specified if "paths" is.
    "paths": {
      "jquery": ["node_modules/jquery/dist/jquery"] // This mapping is relative to "baseUrl"
    }
  }
}
```

경로는 `baseUrl`을 기준으로 처리된다. `baseUrl`을 `"."` 이외의 다른 값으로 설정하는 경우 그에 따라 매핑을 변경해야 한다.

예를 들어 `"baseUrl": "./src"`인 경우 `jquery`는 `"../node_modules/jquery/dist/jquery"`에 매핑되어야 한다.

`path`를 사용하면 다중 fallback 경로를 포함하여 보다 정교한 매핑을 할 수 있다.

한 위치에서 일부 모듈만 사용하고 나머지 모듈은 다른 위치에 있는 프로젝트 구성을 생각해보자.

```txt
projectRoot
├── folder1
│   ├── file1.ts (imports 'folder1/file2' and 'folder2/file3')
│   └── file2.ts
├── generated
│   ├── folder1
│   └── folder2
│       └── file3.ts
└── tsconfig.json
```

이에 대응하는 `tsconfig.json`은 다음과 같다

```json
{
  "compilerOptions": {
    "baseUrl": ".",
    "paths": {
      "*": [
        "*",
        "generated/*"
      ]
    }
  }
}
```

- `"*"`: 변경없이 같은이름 적용, 따라서 `<moduleName> => <baseUrl>/<moduleName>`으로 매핑됨
- `"generated/*"`: 모듈이름에 "generated"라는 prefix를 추가하여 매핑, 따라서 `<moduleName> => <baseUrl>/generated/<moduleName>`

이 논리에 따라 컴파일러는 두 가지 가져오기를 다음과 같이 처리한다

`import 'folder1/file2'`

- `"*"` 패턴 순서에서 와일드카드는 전체 모듈이름을 캡처한다
- 목록에서 첫 대체가 수행된다: `'*' -> folder1/file2`
- 대체의 결과물은 비상대 경로이므로 baseUrl과 결합한다: `projectRoot/folder1/file2.ts`
- 파일이 존재하고 끝난다

`import 'folder2/file3'`

- `"*"` 패턴 순서에서 와일드카드는 전체 모듈이름을 캡처한다
- 목록에서 첫 대체가 수행된다: `'*' -> folder2/file3`
- 대체의 결과물은 비상대 경로이므로 baseUrl과 결합한다: `projectRoot/folder2/file3.ts`
- 파일이 존재하지 않으면 두 번째의 대체를 수행한다
- 두 번째 대체가 수행된다: `'generated/*' -> generated/folder2/file3`
- 대체의 결과물은 비상대 경로이므로 baseUrl과 결합한다: `projectRoot/generated/folder2/file3.ts`
- 파일이 존재하고 끝난다

#### Virtual Directories with rootDirs

때로는 컴파일 타임에 여러 디렉토리의 프로젝트 소스가 모두 결합되어 하나의 출력 디렉토리가 생성된다.
이것은 일련의 소스 디렉토리가 가상(virtual) 디렉토리를 만드는 것으로 볼 수 있다.

`rootDirs`를 사용하면 가상 디레곹리를 구성하는 루트를 컴파일러에 알릴 수 있다.
따라서 컴파일러는 가상 디렉토리 내의 상대경로 모듈 가져오기를 하나의 티렉토리에 병합된 것처럼 해석할 수 있다.

다음과 같은 프로젝트 구조로 예를 들어 보면

```txt
src
└── views
    └── view1.ts (imports './template1')
    └── view2.ts

generated
└── templates
        └── views
            └── template1.ts (imports './view2')
```

빌드 단계에서 `/src/views`와 `/generated/templates/views` 파일 출력을 동일한 디렉토리에 하려고 한다.
런타임에 뷰는 템플릿이 동일 경로에 존재하는 것을 기대하므로 `"./template"`와 같이 상대경로를 사용하여 뷰를 불러온다.

이러한 관계를 컴파일러에게 알려주기 위해서 `rootDirs`를 사용한다.
`rootDirs`는 내용이 런타임에 병합될 것으로 예상되는 루트 목록을 지정한다.

이 경우 `tsconfig.json` 파일은 다음과 같다

```json
{
  "compilerOptions": {
    "rootDirs": [
      "src/views",
      "generated/templates/views"
    ]
  }
}
```

컴파일러는 `rootDirs` 중 하나의 하위폴더에서 상대경로 모듈 가져오기를 볼 때마다 `rootDirs`의 각 항목에서 가져오기를 찾는다.

`rootDirs`의 유연성은 논리적으로 병합된 실제 소스 디렉토리의 목록을 지정하는데에 국한되지 않는다.
디렉토리 목록에는 존재 여부와 관계없이 임의 숫자의 디렉토리 이름이 포함될 수 있다.
이를 통해 컴파일러는 조건부 포함 및 프로젝트 특정 로더 플러그인과 같은 정교한 번들 및 런타임 기능을 type safe 방식으로 캡쳐할 수 있다.

`./#{locale}/messages`의 상대 모듈 경로의 일부로 `#{locale}`와 같은 특수 경로 토큰을 삽입하여
빌드 도구가 locale별 번들을 자동으로 생성하는 국제화 시나리오를 생각해 보자.
이 경우 지원되는 locale을 열거하고 추상화된 경로를 `./ko/messages`, `./de/messages`등으로 매핑한다.

`rootDirs`를 활용하면 컴파일러에게 이 매핑을 알릴 수 있으므로 디렉토리가 존재하지 않더라도 `./#{locale}/messages`를 안전하게 해결할 수 있다.

```json
{
  "compilerOptions": {
    "rootDirs": [
      "src/ko",
      "src/de",
      "src/#{locale}"
    ]
  }
}
```

이제 컴파일러는 `import messages from './#{locale}/messages`를 `import messages from './ko/messages`로 해석하여
디자인 시점에서 지원하기로 협의되지 않은 것을 지원할 수 있도록 한다.

#### Tracing module resolution

앞에서 논의된 것 처럼 컴파일러는 모듈을 확인할 때 현재 디렉토리 외부의 파일을 확인할 수 있다.
이는 모듈이 처리되지 않거나 잘못된 정의로 해석될 때 분석을 어렵게한다.
`--traceResolution` 옵션을 사용하여 컴파일러 모듈 분석 추적을 활성화 하면 모듈 확인 프로세스 중 발생한 문제를 파악할 수 있다.

#### Using `--noResolve`

일반적으로 컴파일러는 컴파일 프로세스를 시작하기 전에 모든 모듈 가져오기를 처리하려고 시도한다.
파일에서 가져오기를 성공할 때마다 컴파일러에서 나중에 처리할 파일 집합에 파일이 추가된다.

`--noResolve` 컴파일러 옵션은 컴파일러가 명령줄에서 전달되지 않은 파일을 컴파일에 추가하지 않도록 처리한다.

`app.ts`

```ts
import * as A from "moduleA" // OK, 'moduleA' passed on the command-line
import * as B from "moduleB" // Error TS2307: Cannot find module 'moduleB'.
```

```sh
tsc app.ts moduleA.ts --noResolve
```

`--noResolve` 옵션을 사용하여 `app.ts`를 컴파일 하려고 하면

- 명령줄에서 전달된 `moduleA`는 올바르게 찾는다
- 명령줄에서 전달되지 않은 `modulesB`는 찾지 못한다

#### Common Questions

> Exclude 목록에 있는 모듈이 컴파일러에 의해 선택되는 경우는?

`tsconfig.json` 설정은 디렉토리를 **프로젝트**로 만든다.
*"excludes"*나 *"files"* 항목을 지정하지 않으면 `tsconfig.json` 경로의 모든 파일과 모든 하위 디렉토리가 컴파일에 포함된다.

일부 파일을 제외시키려면 *"excludes"*를, 컴파일러에게 특정한 파일만을 지정하여 처리하도록 하려면 *"files"*를 사용한다.

컴파일러에서 파일을 가져오기 대상으로 식별한 경우 이전단계에서 제외되었는지 관계없이 컴파일에 포함된다.

## Declaration Merging

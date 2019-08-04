# TypeScript Declaration Files (`d.ts`)

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
$(() => { console.log('hello!'); });
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
  return "Hello, " + s;
}
```

또는

```ts
window.createGreeting = function(s) {
  return "Hello, " + s;
}
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
var fs = require("fs");
```

TypeScript나 ES6에서 `import` 키워드를 동일한 목적으로 사용할 수 있다

```ts
import fs = require("fs");
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
- Declarations like import * as a from 'b'; or export c;
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
import moment = require("moment");
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
var x = require("foo");
// Note: calling 'x' as a function
var y = x(42);
```

모듈을 `new` 키워드로 생성할 수 있는 경우 `module-class.d.ts`를 사용한다

```ts
var x = require("bar");
// Note: using 'new' operator on the imported variable
var y = new x("hello");
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
var x = "hello, world";
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
예를 들어, 불러오면 `String.prototype`에 새 멤버를 추가하는 라이브러리가 있을 수 있다. 이는 런타임 충돌 가능성 때문에 다소 위험하지만 이에 대한 선언파일을 작성할 수는 있다.

##### Global-modifying 모듈 식별

Global-modifying 모듈은 일반적으로 문서에서 쉽게 식별할 수 있다.
일반적으로 글로벌 플러그인과 비슷하지만 효과를 활성화하려면 `require` 호출이 필요하다.

```ts
// 'require' call that doesn't use its return value
var unused = require("magic-string-time");
/* or */
require("magic-string-time");

var x = "hello, world";
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
import * as moment from "moment";
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
  interface KittySettings { }
}
```

다음 처럼 하지 않는다

```ts
// at top-level
interface CatsKittySettings { }
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
import exp = require("express");
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
var a = require("myLib");
var b = require("myLib/foo");
var c = require("myLib/bar");
var d = require("myLib/bar/baz");
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

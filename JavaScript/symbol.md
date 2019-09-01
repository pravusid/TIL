# JavaScript Symbol Type

- <https://developer.mozilla.org/en-US/docs/Glossary/Symbol>
- <https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Symbol>
- <http://hacks.mozilla.or.kr/2015/09/es6-in-depth-symbols/>

`Symbol`은 ES6에서 추가된 Primative type이다.

**심볼** 원시 데이터 타입(primitive data type)은,
클래스나 객체 형식(objet type)의 내부에서만 접근할 수 있도록 전용(private) 객체 속성의 키(key)로 사용된다.

## 필요성

객체를 구분하기 위해서 속성 구분자를 사용할 수도 있다. (예: 속성 중 `isKorean`)
그러나 속성을 추가했을 때 발생할 수 있는 문제가 있다.

- 프로퍼티를 순회하는 `for..in`, `Object.keys()`, `Object.entries()` 등을 사용했을 때 해당 프로퍼티가 잘못 사용될 수 있다
- 다른 라이브러리와 충돌할 수 있다
- 자바스크립트 표준에 해당 프로퍼티가 추가될 수 있다

이를 해결하기 위해 이름 충돌 위험없이 프로퍼티 key 값으로 사용할 수 있는 값이 Symbol 값이다.

예를 들어, `iterator`를 사용하려면 `for..of` 반복문을 사용하는데, 반복문에서는 `[Symbol.iterator]()` 메소드를 호출한다.
단순히 `iterator()` 메소드를 추가해서 사용하지 않고 심볼을 사용하면 기존 코드와 충돌을 회피할 수 있다(하위 호환성 보장).

## 사용

심볼은 고유하고 다른 심볼과 구별된다. (동일한 description이 있더라도 다른 심볼이다)
심볼은 생성되면 변경되지 않고, 속성을 부여할 수 없다. 또한 `string`으로 자동 변환되지 않는다.

- `Symbol()`
  - `new` 키워드로 생성자를 호출하지 않고, 함수만을 호출한다(class-like)
  - 호출마다 새 고유한 심볼(동적으로 익명의 고유한 값)을 반환한다

- `Symbol.for(string)`
  - symbol registry(global symbol table)를 참고하여 심볼을 반환한다
  - 레지스트리에 호출한 심볼이 존재하면 존재하는 심볼을 반환한다
  - 레지스트리에 호출한 심볼이 존재하지 않는다면 생성하여 레지스트리에 등록하고 반환한다

- `Symbol.keyFor(symbol)`
  - symbol을 symbol registry(global symbol table)에서 찾은 뒤 `shared symbol key`를 반환한다

- 표준 `Symbol` 사용
  - `Symbol.iterator`, `Symbol.asyncIterator`, `Symbol.search` 처럼 언어 내장 Symbol을 바로 사용할 수 있다
  - <https://tc39.es/ecma262/#sec-well-known-symbols>

## 가시성

- 비열거형이기 때문에 `for..in` 반복문 내에서 멤버로 사용될 수 없다
- 속성이 익명이기 때문에 `Object.getOwnPropertyNames()`가 반환하는 배열에 들어갈 수 없다
- 심볼 속성 접근은 알고있는 심볼 값을 직접 이용하거나 `Object.getOwnPropertySymbols()`가 반환하는 배열을 사용한다

## 언어상 사례

### `instanceof` 확장

ES6에서, `object instanceof constructor` 구문은 생성자(constructor)의 메소드인 `constructor[Symbol.hasInstance](object)`로 규정됨

### 새로운 종류의 문자열 검색(string-matching)을 지원

ES6에서 `str.match(myObject)` 코드는 우선 `myObject`가 `myObject[Symbol.match](str)` 메소드를 갖고 있는지 확인한다.

심볼로 정의된 메소드로 `RegExp` 객체가 사용되는 모든 곳에 커스텀 문자열 파싱(string-parsing)을 제공할 수 있다.

### 새로운 기능과 이전 코드의 충돌 방지

새로운 기능이 추가되었을 때동작 오류의 주요 원인은 동적 스코핑(dynamic scoping) 기능 때문이었다.

웹 표준은 `Symbol.unscopables`을 이용해서 특정 메소드들이 동적 스코핑에 관여되는 것을 방지한다.

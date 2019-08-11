# Parse Number (integer, float) in JavaScript

## `Number.parseInt()` (or `parseInt()`)

<https://www.ecma-international.org/ecma-262/10.0/index.html#sec-number.parseint>

문자열을 특정 진수를 사용한 정수로 변환한다: `Number.parseInt(string,[ radix])`

- 선행 공백은 무시한다
- 기호(+, -), 숫자(0-9), 소수점 또는 지수 이외의 문자를 발견하면, 그 전까지의 결과만 반환하고 해당 문자와 그 이후는 모두 무시한다
- 첫 번째 문자를 숫자로 변환할 수 없는 경우 `NaN`을 반환한다
- 숫자인 문자열을 해석할 때 지수를 지정할 수 있다
- 해석할 지수를 명시하지 않아도 `0x`로 시작하는 문자열은 16진수로 해석한다

## `Number.parseFloat()` (or `parseFloat()`)

<https://www.ecma-international.org/ecma-262/10.0/index.html#sec-number.parsefloat>

문자열을 부동 소수점 실수로 변환한다: `parseFloat(value)`

- 선행 공백은 무시한다
- 기호(+, -), 숫자(0-9), 소수점 또는 지수 이외의 문자를 발견하면, 그 전까지의 결과만 반환하고 해당 문자와 그 이후는 모두 무시한다
- 첫 번째 문자를 숫자로 변환할 수 없는 경우 `NaN`을 반환한다
- Infinity도 분석 및 반환할 수 있으며 `isFinite()` 함수를 사용해 구분할 수 있다 (Infinity, -Infinity, NaN이 아닌 수)
- `toString`이나 `valueOf` 메소드를 구현한 객체 입력해도 문자열을 전달한 것 처럼 작동한다

## `Number()`

<https://www.ecma-international.org/ecma-262/10.0/index.html#sec-tonumber>

Number 객체는 숫자 값으로 작업할 수 있게 해주는 래퍼wrapper 객체이다.
Number 객체는 `Number()` **생성자**를 사용하여 만든다. 원시 숫자 자료형은 `Number()` **함수**를 사용해 생성한다.

- 변환 규칙은 상단 링크를 참고
- `parseFloat`과 비슷하지만 trailing text를 허용하지 않고(`parse##` 보다 엄격함) 결과로 `NaN`을 반환함

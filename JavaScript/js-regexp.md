# JavaScript 정규표현식 (regexp)

- <https://developer.mozilla.org/ko/docs/Web/JavaScript/Guide/Regular_expressions>
- <https://github.com/ziishaned/learn-regex/blob/master/translations/README-ko.md>
- <https://en.wikipedia.org/wiki/Regular_expression>
- <https://www.regular-expressions.info/quickstart.html>

## 예약어 (특수문자)

`[ ] ( ) { } . * + ? ^ $ \ |`

## 정규표현식 규칙

<https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Regular_expressions/Cheatsheet>

정규표현식 리터럴은 `const regexp = /ab+c/;` 와 같이 슬래쉬(`/`)로 감싸는 패턴을 사용한다

정규표현식 생성자 함수를 호출할 수도 있다: `const regexp = new RegExp("ab+c");`

| 문자     | 일치                                | 예제                               |
| -------- | ----------------------------------- | ---------------------------------- |
| `^`      | 입력값의 시작                       | `/^This/` => **This** is end.      |
| `$`      | 입력값의 끝                         | `/end\.$/` => This is the **end.** |
| `*`      | 0번 이상 반복                       | `/se*/` => **seeee**d              |
| `+`      | 1번 이상 반복                       | `/ap+/` => **apple**               |
| `?`      | 0번 또는 1번 반복                   | `/ap?/` => **apple**               |
| `{n}`    | 정확히 n번 반복                     | `/ap{2}/` => **app**le             |
| `{n,}`   | n번 이상 반복                       | `/ap{1,}/` => **app**le            |
| `{n,m}`  | 최소 n번, 최대 m번                  | `/ap{2,4}/` => **apppp**ppple      |
| `x\|y`   | x 또는 y                            | `/p\|l/` => a**ppppl**e            |
| `[xyz]`  | 대괄호 안의 모든문자                | `/a[px]e/` => a**p**e or a**x**e   |
| `[^xyz]` | 대괄호 안의 문자를 제외한 모든문자  | `/a[^px]/` => axe or ape           |
| `.`      | 줄 바꿈을 제외한 모든문자           | `/.pp/` => **app**le               |
| `\b`     | 단어 경계                           | `/\bno/` => **no**nonono           |
| `\B`     | 단어 경계를 제외한 모든문자         | `/\Bno/` => no**nonono**           |
| `\d`     | 0부터 9까지의 숫자                  | `/\d{3}/` => Now in **123**        |
| `\D`     | 숫자를 제외한 모든문자              | `/\D{2,4}/` => **Now** in 123      |
| `\w`     | 단어문자 == 알파벳, 숫자, 밑줄(`_`) | `/\w/` => **j**avascript           |
| `\W`     | 단어문자가 아닌 문자                | `/\W/` => 100 **\$**               |
| `\n`     | 줄바꿈                              |                                    |
| `\s`     | 하나의 공백 문자                    |                                    |
| `\S`     | 공백 문자가 아닌 모든문자           |                                    |
| `\t`     | 탭                                  |                                    |

### Character classes

<https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Regular_expressions/Character_classes>

> 대부분의 정규표현식 처리기에서 character class 내부는 다른규칙이 적용되고, 사용할 수 있는 특수문자는 다음과 같다.
> `backslash \, caret ^, hyphen -`
>
> -- <https://www.regular-expressions.info/charclass.html>
> -- <https://stackoverflow.com/questions/19976018>

### capturing & non-capturing

| 문자      | 일치                                          | 예제                                                               |
| --------- | --------------------------------------------- | ------------------------------------------------------------------ |
| `(x)`     | x를 묶고(하위패턴), `\위치`에 대응            | `/(foo) (bar) \1 \2/`는 "foo bar foo bar"는 정규식 순서대로 대응함 |
| `(?:x)`   | x를 묶지만(하위패턴), `\위치`에 대응 하지않음 | `/(?:foo){1,2}/` 처럼 문자를 묶어 하위패턴을 정의할 때 사용        |
| `x(?=y)`  | y 앞의 x에 대응 (lookahead)                   | `[T\|t]he(?=\sfat)` => **The** fat cat sat on the mat.             |
| `x(?!y)`  | not y 앞의 x에 대응 (negated lookahead)       | `[T\|t]he(?!\sfat)` => The fat cat sat on **the** mat.             |
| `(?<=y)x` | y 뒤의 x에 대응 (lookbehind)                  | `(?<=[T\|t]he\s)(fat\|mat)` => The **fat** cat sat on the **mat**. |
| `(?<!y)x` | not y 뒤의 x에 대응 (negated lookbehind)      | `(?<![T\|t]he\s)(cat)` => The cat sat on **cat**.                  |

### flag

플래그는 정규표현식 리터럴 마지막에 추가한다: `const regexp = /pattern/flags;`

| Flag | Description                                                                              |
| ---- | ---------------------------------------------------------------------------------------- |
| `g`  | 전역 검색                                                                                |
| `i`  | 대소문자 구분 없는 검색                                                                  |
| `m`  | 다중행(multi-line) 검색                                                                  |
| `u`  | 유니코드; 패턴을 유니코드 코드 포인트의 나열로 취급합니다                                |
| `y`  | "sticky" 검색을 수행. 문자열의 현재 위치부터 검색을 수행합니다. sticky 문서를 확인하세요 |

### capturing

정규식 패턴 내부에서 capturing 사용: `\n`은 n번째 capturing 대응함

> 패턴 `/(foo) (bar) \1 \2/` 안의 '(foo)' 와 '(bar)'는 문자열 "foo bar foo bar"에서 앞 두단어에 대응하고, 패턴 내부의 `\1 \2`는 문자열의 뒷 두 단어에 대응한다.

정규식의 치환에서 capturing 사용: `$n`은 n번째 capturing 대응하고, `$&`는 capturing 전체 문자열을 가리킴

> `'bar foo'.replace( /(...) (...)/, '$2 $1')`를 처리한 결과는 'foo bar'가 된다.

## 예제

```js
// 줄바꿈
const linebreak = /[\r|\n|\r\n]$/;

// 아이디 체크
const regExpId = /^[a-z0-9_-]\w{5,20}$/;

// 비밀번호 길이 체크
const regExpPassword = /^\w[6,16]$/;

// 비밀번호 조합(영문, 숫자) 및 길이 체크
const regExpPassword = /^(?=.*[a-zA-Z])(?=.*[0-9]).{6,16}$/;

// 이메일 체크
const regExpEmail = /^([\w-]+(?:\.[\w-]+)*)@((?:[\w-]+\.)*\w[\w-]{0,66})\.([a-z]{2,6}(?:\.[a-z]{2})?)$/;

// 휴대폰번호
const regExpMobile = /^01([016789]?)-?([0-9]{3,4})-?([0-9]{4})$/;

// 숫자만 사용
const regExpNumber = /^\d+$/;
```

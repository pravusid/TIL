# JavaScript String

<https://developer.mozilla.org/ko/docs/Web/JavaScript/Reference/Global_Objects/String>

## String.split

<https://developer.mozilla.org/ko/docs/Web/JavaScript/Reference/Global_Objects/String/split>

### 정규표현식을 사용해 문자열 분해

<https://www.stefanjudis.com/today-i-learned/how-to-preserve-separators-in-the-result-of-string-prototype-split/>

```js
// Split the string on "-" and "_"
"Hello_party-people!".split(/[-_]/);
// Array(3) [ "Hello", "party", "people!" ]
```

문자열 분해에 사용하는 정규표현식을 capture 하는 경우 분해 결과에 `separator` 값도 포함된다

```js
// Split the string on "-" and "_"
// but include the separator in the result
"Hello_party-people!".split(/([-_])/);
// Array(5) [ "Hello", "_", "party", "-", "people!" ]
```

# 값의 진위여부 for JavaScript

참고: <https://github.com/denysdovhan/wtfjs>

```js
if (value) {
  // ...
}
```

다음 값은 `false`로 평가됨

- null
- undefined
- NaN
- empty string ("")
- 0
- false

그러나 할당은 되어 있고 값만 비어있다면 `true`로 평가됨

- empty object: `{}`
- empty array: `[]`

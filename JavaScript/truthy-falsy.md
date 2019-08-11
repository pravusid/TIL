# 값의 진위여부 for JavaScript

References

- <https://github.com/denysdovhan/wtfjs>
- <https://developer.mozilla.org/ko/docs/Glossary/Truthy>
- <https://developer.mozilla.org/ko/docs/Glossary/Falsy>

다음 값은 `false`로 평가됨

- `false`
- `null`
- `undefined`
- `NaN`
- empty string (`''`)
- `0`

그러나 할당은 되어 있고 값만 비어있다면 `true`로 평가됨

- empty object: `{}`
- empty array: `[]`

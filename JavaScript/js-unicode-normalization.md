# JavaScript 유니코드 정규화

<https://developer.mozilla.org/ko/docs/Web/JavaScript/Reference/Global_Objects/String/normalize>

## Diacritic (발음 구별 기호) 제거 (latin to ascii)

> 유니코드 정규형 분해 후 [Unicode character class escape](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Regular_expressions/Unicode_character_class_escape)정의를 사용해서 발음 구별 기호 제거
>
> --<https://stackoverflow.com/questions/990904/remove-accents-diacritics-in-a-string-in-javascript>

```js
const str = 'Crème Brûlée';
str.normalize('NFD').replace(/\p{Diacritic}/gu, '');
// "Creme Brulee"
```

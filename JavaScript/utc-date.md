# Javascript UTC Datetime

## UTC timezone offset 반영 시각 변환

```js
const date = new Date();
const result = new Date(date.getTime() - (date.getTimezoneOffset() * 60000)).toJSON();
```

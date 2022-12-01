# Node.js CLI 인자전달

`$ node server.js one two=three four`

위의 명령을 실행하면 다음과 같은 방법으로 인자 배열을 출력할 수 있다

```js
const args = process.argv;
console.log(args);
```

`args`를 출력해보면 다음과 같은 형태이다: (커맨드 전체를 출력함)

```js
[
  '/usr/bin/node',
  '/home/idpravus/server.js',
  'one',
  'two=three',
  'four'
]
```

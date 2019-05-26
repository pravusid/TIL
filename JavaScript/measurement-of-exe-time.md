# 실행시간 측정

Node.js의 `process.hrtime([time])` API 사용

## API

<https://nodejs.org/api/process.html#process_process_hrtime_time>

`process.hrtime([time]): integer[]`

The process.hrtime() method returns the current high-resolution real time in a **[seconds, nanoseconds]**

## 예제

```ts
const start = new Date();
const hrstart = process.hrtime();

setTimeout(() => {
  // execution time simulated with setTimeout function
  const end = new Date() - start;
  const hrend = process.hrtime(hrstart);

  const [sec, nano] = hrend;
  const result = (sec * 1e9 + nano) / 1e9;
  // or const result = sec + (nano / 1000000);
}, 1000);
```

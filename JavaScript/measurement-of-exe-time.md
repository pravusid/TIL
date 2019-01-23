# 실행시간 측정

Node.js의 `process.hrtime([time])` API 사용

## API

<https://nodejs.org/api/process.html#process_process_hrtime_time>

`process.hrtime([time]): integer[]`

The process.hrtime() method returns the current high-resolution real time in a **[seconds, nanoseconds]**

## 예제

```js
const start = new Date();
const hrstart = process.hrtime();

setTimeout(() => {
  // execution time simulated with setTimeout function
  const end = new Date() - start;
  const hrend = process.hrtime(hrstart);

  console.info("Execution time: %dms", end);
  console.info("Execution time (hr): %ds %dms", hrend[0], hrend[1] / 1000000);
}, 1000);
```

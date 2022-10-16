# JavaScript Async

## Callback, 콜백

<https://ko.javascript.info/callbacks>

## Promise, 프로미스

<https://ko.javascript.info/promise-basics>

## Async & Await

<https://ko.javascript.info/async-await>

### `async` 키워드

- `async` 키워드를 사용하였을 때 반환값이 Promise가 아닌경우 Resolved Promise로 감싸서 반환하게 된다

- `async` 키워드를 사용하였을 때 scope 내에서 `await`를 사용한 경우와 아닌경우 오류발생은 다른 결과를 가져온다

  - `await` 키워드 미사용: 동기함수에서 오류발생한 것과 동일한 오류발생
  - `await` 키워드 사용: `processTicksAndRejections` 마이크로태스크큐에서 오류발생 ([[nodejs-event-loop]] 참고)

### Async Stacktrace in V8

> The fundamental difference between await and vanilla promises is that await X() suspends execution of the current function,
> while promise.then(X) continues execution of the current function after adding the X call to the callback chain.
>
> -- <https://mathiasbynens.be/notes/async-stack-traces>

- <https://v8.dev/docs/stack-trace-api#async-stack-traces>
- <https://cloudreports.net/v8-zero-cost-async-stack-traces/>

<https://github.com/goldbergyoni/nodebestpractices/blob/master/sections/errorhandling/returningpromises.md>

> V8 엔진의 "zero-cost async stacktraces" 기능을 활용하기 위해서는
> 오류발생가능성이 있고 Promise 값을 반환하는 함수에 await을 사용하지 않고 바로 반환하면 안된다

`await` 키워드 없이 반환한 경우

```js
async function throwAsync(msg) {
  await null // need to await at least something to be truly async (see note #2)
  throw Error(msg)
}

async function returnWithoutAwait () {
  return throwAsync('missing returnWithoutAwait in the stacktrace')
}

// 👎 will NOT have returnWithoutAwait in the stacktrace
returnWithoutAwait().catch(console.log)
```

오류 발생하여도 stacktrace에 `throwAsync` 함수가 보이지 않는다

```txt
Error: missing returnWithoutAwait in the stacktrace
    at throwAsync ([...])
```

`await` 키워드를 사용하여 반환한 경우

```js
async function throwAsync(msg) {
  await null // need to await at least something to be truly async (see note #2)
  throw Error(msg)
}

async function returnWithAwait() {
  return await throwAsync('with all frames present')
}

// 👍 will have returnWithAwait in the stacktrace
returnWithAwait().catch(console.log)
```

오류 발생하면 stacktrace에 `throwAsync` 함수가 보인다

```txt
Error: with all frames present
    at throwAsync ([...])
    at async returnWithAwait ([...])
```

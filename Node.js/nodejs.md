# Node.js

## 개요

JavaScript Runtime

## 특징

1. google V8 engine : 구글이 개발한 V8 자바스크립트 엔진 기반, 자바스크립트 기반이므로 single thread로 작동함
2. Event driven : 순차적으로 작업을 실행하는 것이 아니라 작업에서 무엇을 처리해야하는지만 알려주고(callback) 시간이 걸리는 처리들을 EventListener로 위임한다.
3. Nonblocking I/O / Asynchronous : single thread로 작동하므로 다중 요청 처리를 위해서 비동기 nonblocking I/O로 처리가 이루어짐

## Node.js CLI

<https://nodejs.org/dist/latest-v20.x/docs/api/cli.html>

### `--enable-source-maps`

- Added in: v15.11.0, v14.18.0 (Stable)
- <https://nodejs.org/dist/latest-v20.x/docs/api/cli.html#--enable-source-maps>
- <https://github.com/evanw/node-source-map-support?tab=readme-ov-file#node-12120>

### `--import=module`

- Added in: v19.0.0 (Experimental)
- <https://nodejs.org/dist/latest-v20.x/docs/api/cli.html#--importmodule>

> Follows ECMAScript module resolution rules.
> Use --require to load a CommonJS module. Modules preloaded with --require will run before modules preloaded with --import.

### `--env-file=config`

- Added in: v20.6.0 (Active development)
- <https://nodejs.org/dist/latest-v20.x/docs/api/cli.html#--env-fileconfig>

```bash
# You can pass multiple --env-file arguments. Subsequent files override pre-existing variables defined in previous files.
node --env-file=.env --env-file=.development.env index.js
```

### `--watch`, `--watch-path`

- Added in: v18.11.0, v16.19.0 (Experimental)
- <https://nodejs.org/dist/latest-v20.x/docs/api/cli.html#--watch>
- <https://nodejs.org/dist/latest-v20.x/docs/api/cli.html#--watch-path>
- `--watch-preserve-output`: Disable the clearing of the console when watch mode restarts the process.

```bash
# This flag cannot be combined with --check, --eval, --interactive, --test, or the REPL.
node --watch-path=./src --watch-path=./tests index.js
```

### `--max-old-space-size=SIZE` (in megabytes)

<https://nodejs.org/dist/latest-v20.x/docs/api/cli.html#--max-old-space-sizesize-in-megabytes>

```bash
NODE_OPTIONS="--max-old-space-size=4096" node index.js
```

## http

<https://nodejs.org/ko/docs/guides/anatomy-of-an-http-transaction/>

nodejs 는 다양한 작업을 수행할 수 있지만, 웹 기반 Application에 적합하도록 많은 투자가 이루어지고 있다.

### 의존성

<https://nodejs.org/ko/docs/meta/topics/dependencies/#llhttp>

> HTTP 파싱은 llhttp라는 경량 C 라이브러리가 처리합니다.
> 이는 시스템 호출이나 할당을 하려고 만들어진 것이 아니므로 요청당 아주 작은 메모리 공간만 차지합니다.

### nodejs server 예제

```js
const http = require('http');

const hostname = '127.0.0.1';
const port = 3000;

const server = http.createServer((req, res) => {
  res.statusCode = 200;
  res.setHeader('Content-Type', 'text/plain');
  res.end('Hello World\n');
});

server.listen(port, hostname, () => {
  console.log(`Server running at http://${hostname}:${port}/`);
});
```

> 이 서버로 오는 HTTP 요청마다 createServer에 전달된 함수가 한 번씩 호출됩니다.
> 사실 createServer가 반환한 Server 객체는 EventEmitter이고 여기서는 server 객체를 생성하고 리스너를 추가하는 축약 문법을 사용한 것입니다.

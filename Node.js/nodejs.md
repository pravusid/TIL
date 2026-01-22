# Node.js

## 개요

JavaScript Runtime

## 특징

1. google V8 engine : 구글이 개발한 V8 자바스크립트 엔진 기반, 자바스크립트 기반이므로 single thread로 작동함
2. Event driven : 순차적으로 작업을 실행하는 것이 아니라 작업에서 무엇을 처리해야하는지만 알려주고(callback) 시간이 걸리는 처리들을 EventListener로 위임한다.
3. Nonblocking I/O / Asynchronous : single thread로 작동하므로 다중 요청 처리를 위해서 비동기 nonblocking I/O로 처리가 이루어짐

## Node.js CLI

<https://nodejs.org/docs/latest/api/cli.html>

### `--enable-source-maps`

- Added in: v15.11.0, v14.18.0 (Stable)
- <https://github.com/evanw/node-source-map-support?tab=readme-ov-file#node-12120>

### `--import=module`

- Added in: v19.0.0 (Experimental)

> Follows ECMAScript module resolution rules.
> Use --require to load a CommonJS module. Modules preloaded with --require will run before modules preloaded with --import.

### `--env-file=config`

- Added in: v20.6.0 (Active development)

```bash
# You can pass multiple --env-file arguments. Subsequent files override pre-existing variables defined in previous files.
node --env-file=.env --env-file=.development.env index.js
```

### `--watch`, `--watch-path`

- Added in: v18.11.0, v16.19.0 (Experimental)
- `--watch-preserve-output`: Disable the clearing of the console when watch mode restarts the process.

```bash
# This flag cannot be combined with --check, --eval, --interactive, --test, or the REPL.
node --watch-path=./src --watch-path=./tests index.js
```

### `--max-old-space-size=SIZE` (in megabytes)

```bash
node --max-old-space-size=1536 index.js
```

### `--openssl-legacy-provider`

- [Investigate loading legacy provider with OpenSSL 3.0](https://github.com/nodejs/node/issues/40455)
- <https://www.codingbeautydev.com/blog/node-err-ossl-evp-unsupported>

```bash
node --openssl-legacy-provider
```

### `--run`

- Added in: v22.0.0

```bash
# For example, the following command will run the test script of the package.json in the current folder:
node --run test

# You can also pass arguments to the command. Any argument after -- will be appended to the script:
node --run test -- --verbose
```

## Compile Cache

- Added in: v22.1.0
- <https://nodejs.org/api/cli.html#node_compile_cachedir>
- <https://nodejs.org/api/module.html#module-compile-cache>

```bash
NODE_COMPILE_CACHE=dir node index.js
```

## `NODE_OPTIONS`

<https://nodejs.org/api/cli.html#node_optionsoptions>

- 공백으로 구분된 명령행 옵션 여러 건을 환경변수로 정의할 수 있다
- 명령행 옵션보다 먼저 해석되고, 함께 적용되거나 overriding 된다

```bash
# 예시 1
node --max-old-space-size=1536 index.js
# is equivalent to:
NODE_OPTIONS="--max-old-space-size=1536" node index.js

# 예시 2
NODE_OPTIONS='--require "./a.js"' node --require "./b.js"
# is equivalent to:
node --require "./a.js" --require "./b.js"
```

> `.npmrc` 파일에 다음 내용을 추가해도 된다, npm script(`npm run ??`) 호출할 때 적용 됨
>
> --<https://docs.npmjs.com/cli/v10/using-npm/config#node-options>

```env
node-options="--max-old-space-size=1536"
```

## Node.js default max-old-space-size

- <https://github.com/nodejs/node/pull/25576#issuecomment-455737693>
- <https://github.com/nodejs/node/issues/35573>
- <https://github.com/nodejs/node/issues/39107>
- <https://github.com/nodejs/node/issues/43991>
- <https://stackoverflow.com/a/71337579>
- <https://source.chromium.org/chromium/chromium/src/+/main:v8/src/heap/heap.cc>

old generation heap size (64bit 기준)

- 512MB 미만: 256MB
- 512MB~4GB: 256MB~2048MB
- 4GB~15GB: 2048MB
- 15GB이상: 4096MB

## http

<https://nodejs.org/en/learn/modules/anatomy-of-an-http-transaction>

nodejs 는 다양한 작업을 수행할 수 있지만, 웹 기반 Application에 적합하도록 많은 투자가 이루어지고 있다.

### 의존성

<https://github.com/nodejs/llhttp>

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

## nodejs ESM

- <https://nodejs.org/docs/latest/api/esm.html>
- <https://nodejs.org/en/blog/release/v22.12.0#requireesm-is-now-enabled-by-default>
- [Pure ESM package](https://gist.github.com/sindresorhus/a39789f98801d908bbc7ff3ecc99d99c)
- [Move on to ESM-only](https://antfu.me/posts/move-on-to-esm-only)

## nodejs built-in support for TypeScript

> v22.6.0 에서 `--experimental-strip-types` 플래그 추가
> v23.6.0 부터 Type stripping 은 플래그 없이 기본기능으로 동작

- <https://nodejs.org/docs/latest/api/typescript.html#type-stripping>
- <https://github.com/nodejs/typescript> [Roadmap to stable strip-types](https://github.com/nodejs/typescript/issues/24)
- <https://github.com/nodejs/amaro>
- <https://nodejs.org/en/learn/typescript/run-natively>
- <https://github.com/bloomberg/ts-blank-space>
- <https://satanacchio.hashnode.dev/everything-you-need-to-know-about-nodejs-type-stripping>

### Determining module system

- `.ts` files will have their module system determined the same way as .js files. To use import and export syntax, add "type": "module" to the nearest parent package.json.
- `.mts` files will always be run as ES modules, similar to .mjs files.
- `.cts` files will always be run as CommonJS modules, similar to .cjs files.
- `.tsx` files are unsupported.

### TypeScript features

타입스크립트의 모든 기능을 지원하는 것은 아니며 아래의 기능이 포함된 코드의 경우 오류가 발생함

- Enum declarations
- namespace with runtime code
- legacy module with runtime code
- parameter properties
- import aliases

오류를 방지하기 위해서 두 가지 방법을 사용할 수 있음

- <https://devblogs.microsoft.com/typescript/announcing-typescript-5-8/#the---erasablesyntaxonly-option>
- [--experimental-transform-types](https://nodejs.org/docs/latest/api/cli.html#--experimental-transform-types)

### Importing types without type keyword

type stripping을 실행할 때 import 구문에서 타입을 불러올 때 type 키워드를 사용하지 않은 코드를 실행하는 경우 런타임 오류가 발생한다.
tsconfig의 [Verbatim Module Syntax](https://devblogs.microsoft.com/typescript/announcing-typescript-5-0/#--verbatimmodulesyntax) 옵션을 사용해서 오류를 방지할 수 있다.

### [[typescript-esm-nodejs]] 상호호환

type stripping에서는 기존 TypeScript-ESM 지원을 위해 작성한 `import ... from './foo.js'` 구문을 처리할 수 없다.
TypeScript 5.7에서 [Path Rewriting for Relative Paths](https://devblogs.microsoft.com/typescript/announcing-typescript-5-7/#path-rewriting-for-relative-paths) 옵션이 추가되었다.

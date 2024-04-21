# Express.js

> 신규 프로젝트라면 [fastify](https://www.fastify.io/) 사용을 고려하는 것이 좋음

## 배포

배포 후 환경변수를 production으로 변경해야 함: `NODE_ENV=production node app.js`

앱 내부에서 `process.env.NODE_ENV` 값에 할당되어 express 배포시 최적화 처리됨

## Middleware

> Express is a routing and middleware web framework that has minimal functionality of its own:
> An Express application is essentially a series of middleware function calls.

미들웨어의 `next` 함수는 다음과 같이 동작한다

- 미들웨어는 정의한 순서대로 순차 실행되며 `next` 함수로 흐름을 조절한다
- 미들웨어 파라미터 중 `next` 함수를 실행하면 다음 미들웨어를 호출한다
- `next` 함수에 `'route'` 문자열을 제외한 값을 전달하여 실행하는 경우 나머지 미들웨어를 건너뛰고 에러핸들러를 실행한다
- `next('route')`를 호출하는 경우 다음 라우트 핸들러를 호출한다: `app.get('/path', handlerA, handlerB)` 일때 A -> B
- `next` 함수를 호출하지 않으면 응답(`res`)을 처리해야 하며, 둘 다 하지 않으면 요청은 중단된 상태로 남는다 (left hanging)

## ErrorHandling

- Express.js 4버전에서는 Router에서 Promise처리를 지원하지 않음
- 따라서 Router에서 Async Function을 사용하려면 두 방법 중 하나를 선택해야 함

### 방법1. 라이브러리 사용

- [`express-async-errors`](https://github.com/davidbanham/express-async-errors)
- [`express-promise-router`](https://github.com/express-promise-router/express-promise-router)

### 방법2. Wrapping Route Functions

```ts
export const errorHandler = (error: Error, request: Request, response: Response, next: NextFunction) => {
  response.status(500).json({ message: error.message });
};

type AsyncHandler<Req extends Request, Res extends Response, Next extends NextFunction = NextFunction> = (
  req: Req,
  resp: Res,
  next: Next
) => Promise<unknown>;

export const asyncHandler: <Req extends Request, Res extends Response>(
  handler: AsyncHandler<Req, Res>
) => AsyncHandler<Req, Res> = (handler) => (request, response, next) =>
  Promise.resolve(handler(request, response, next)).catch((error: Error) =>
    errorHandler(error, request, response, next)
  );
```

다음처럼 사용한다

```ts
// app.ts
app.use(errorHandler);

// foo.controller.ts
this.routes.get(
  '/hello',
  asyncHandler((req, resp) => this.foobar(req, resp))
);
```

### Error Handling Middleware

<https://expressjs.com/en/guide/error-handling.html>

> 에러핸들러는 4개의 인자 `(err, req, res, next)`를 갖는 미들웨어이다

사용자정의 에러핸들러를 정의할 때 다음 사항에 유의해야 한다

- **반드시 미들웨어 및 라우트 호출을 정의한 뒤 마지막으로 정의해야 한다** (정의한 순서대로 실행되므로)
- express 미들웨어 마지막에는 항상 기본 에러핸들러가 실행된다 (별도로 정의하지 않아도 실행된다)
- 사용자정의 에러핸들러에서 `next(err)` 함수를 호출하는 경우 기본 에러핸들러가 실행된다
- 사용자정의 에러핸들러에서 `next` 함수를 호출하지 않는다면 반드시 응답(`res`객체)을 처리해야 한다

> 과거 nodejs는 비동기 호출을 위해 콜백을 사용했고, 콜백의 첫 파라미터는 항상 optional error 였기 때문에
> `next` 함수를 콜백으로 바로 넘기는 경우 오류가 발생한 경우 에러핸들러를 간단하게 호출할 수 있었을 거라 추측된다

## merging interfaces in TypeScript

- <https://github.com/DefinitelyTyped/DefinitelyTyped/blob/master/types/express-serve-static-core/index.d.ts>
- <https://github.com/DefinitelyTyped/DefinitelyTyped/blob/master/types/express/index.d.ts>

`types/express.d.ts`

```ts
import { User } from '../src/domain/user';

declare global {
  namespace Express {
    export interface Request {
      user: User;
    }
  }
}
```

`tsconfig.json`

```json
{
  "files": ["types/express.d.ts"]
}
// OR
{
  "include": ["src/**/*", "types/**/*"]
}
```

> ts-node에서 타입인식이 되지 않는 경우: <https://github.com/TypeStrong/ts-node#help-my-types-are-missing>

`tsconfig.json`

```json
{
  "compilerOptions": {
    "typeRoots": ["./node_modules/@types", "./types"]
  }
}
```

경로는 다음과 같음

```txt
<project_root>/
-- tsconfig.json
-- types/
  -- <module_name>/
    -- index.d.ts
```

## HTTPS

<https://nodejs.org/api/tls.html#tls_tls_createsecurecontext_options>

ca 옵션은 다음 사항이 적용된다

- 선택적으로 신뢰할 수 있는 CA 인증서를 대체할 수 있다. 기본값은 Mozilla가 선택한 잘 알려지고 믿을 수 있는 CA로 구성되어 있다.
- 잘 알려진 CA에 연결되지 않은(not chainable) 인증서를 사용하는 경우, 인증서의 CA(Intermediate certificate)를 명시하지 않으면 연결에 실패한다.
- 자체 서명 인증서를 사용하는 경우 자체 인증기관(own CA)이 명시되어야 한다.

```ts
import * as express from 'express';
import { readFileSync } from 'fs';
import { createServer } from 'https';

require('dotenv').config();

const app = express();

createServer(
  {
    ca: readFileSync('cert/chain.crt'), // 인증서 체인
    key: readFileSync('cert/server.key'), // 서버 비밀키
    cert: readFileSync('cert/server.crt'), // 서버 도메인 인증서
  },
  app
).listen(process.env.PORT || 3000, () => console.log('서버실행'));
```

<https://nodejs.org/api/tls.html#tls_tls_createserver_options_secureconnectionlistener>

`rejectUnauthorized`

- If not false the server will reject any connection which is not authorized with the list of supplied CAs.
- This option only has an effect if requestCert is true
- Default: `true`

`requestCert`

- If true the server will request a certificate from clients that connect and attempt to verify that certificate.
- Default: `false`

위의 두 옵션을 사용하고 자체 서명 인증서인 경우 클라이언트에서 인증서를 처리해야 한다

<https://nodejs.org/api/tls.html#tls_tls_connect_options_callback>

```ts
const agent = new https.Agent({
  // Necessary only if the server requires client certificate authentication.
  key: fs.readFileSync('client-key.pem'),
  cert: fs.readFileSync('client-cert.pem'),

  // Necessary only if the server uses a self-signed certificate.
  ca: [fs.readFileSync('server-cert.pem')],

  // Necessary only if the server's cert isn't for "localhost".
  checkServerIdentity: () => null,
});

// 참고: https Agent option
interface AgentOptions extends http.AgentOptions, tls.ConnectionOptions {
  rejectUnauthorized?: boolean;
  maxCachedSessions?: number;
}
```

## Graceful Shutdown

- <https://expressjs.com/en/advanced/healthcheck-graceful-shutdown.html>
- [[nodejs-graceful-shutdown]]

## 경로별 bodyParser 사용

> bodyParser (global)미들웨어를 경로별로 지정하려면 순서에 유의해야 함
>
> -- <https://github.com/expressjs/express/issues/3932>

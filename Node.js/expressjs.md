# Express.js

> 신규 프로젝트라면 [fastify](https://www.fastify.io/) 사용을 고려하는 것이 좋음

## 배포

배포 후 환경변수를 production으로 변경해야 함: `NODE_ENV=production node app.js`

앱 내부에서 `process.env.NODE_ENV` 값에 할당되어 express 배포시 최적화 처리됨

## ErrorHandling

Express.js 4버전에서는 Router에서 Promise처리를 지원하지 않음

따라서 Router에서 Async Function을 사용하려면 두 방법 중 하나를 선택해야 함

### 라이브러리 사용

- [`express-async-errors`](https://github.com/davidbanham/express-async-errors)
- [`express-promise-router`](https://github.com/express-promise-router/express-promise-router)

### Wrapping Route Functions

```ts
export const errorHandler = (
  error: Error,
  request: Request,
  response: Response,
  next: NextFunction
) => {
  response.status(500).json({ message: error.message });
  next();
};

type AsyncFunc = (
  req: Request,
  resp: Response,
  next: NextFunction
) => Promise<any>;

export const asyncHandler: (func: AsyncFunc) => AsyncFunc = (func) => {
  return (request, response, next) =>
    Promise.resolve(func(request, response, next)).catch((error: Error) =>
      errorHandler(error, request, response, next)
    );
};
```

다음처럼 사용한다

```ts
// app.ts
app.use(errorHandler);

// foo.controller.ts
this.routes.get(
  "/hello",
  asyncHandler((req, resp) => this.foobar(req, resp))
);
```

## merging interfaces

- <https://github.com/DefinitelyTyped/DefinitelyTyped/blob/master/types/express-serve-static-core/index.d.ts>
- <https://github.com/DefinitelyTyped/DefinitelyTyped/blob/master/types/express/index.d.ts>

`types/express.d.ts`

```ts
import { User } from "../src/domain/user";

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
import * as express from "express";
import { readFileSync } from "fs";
import { createServer } from "https";

require("dotenv").config();

const app = express();

createServer(
  {
    ca: readFileSync("cert/chain.crt"), // 인증서 체인
    key: readFileSync("cert/server.key"), // 서버 비밀키
    cert: readFileSync("cert/server.crt"), // 서버 도메인 인증서
  },
  app
).listen(process.env.PORT || 3000, () => console.log("서버실행"));
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
  key: fs.readFileSync("client-key.pem"),
  cert: fs.readFileSync("client-cert.pem"),

  // Necessary only if the server uses a self-signed certificate.
  ca: [fs.readFileSync("server-cert.pem")],

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

> <https://expressjs.com/en/advanced/healthcheck-graceful-shutdown.html>

nodejs 서버를 사용하는 다른 라이브러리에도 적용가능함

### terminus

[`@godaddy/terminus`](https://github.com/godaddy/terminus) 실행순서

- `beforeShutdown`
- `onSignal`
- `onShutdown`

### 다른 선택지

- <https://github.com/gquittet/graceful-server>
- <https://github.com/sebhildebrandt/http-graceful-shutdown>

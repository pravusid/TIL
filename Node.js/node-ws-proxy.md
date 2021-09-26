# WebSocket proxy

```ts
const httpProxy = require("http-proxy");
const fs = require("fs");

const server = httpProxy
  .createServer({
    target: "wss://localhost:10101",
    ws: true,
    // 개발환경 wss 사용을 위한 self-signed cert
    secure: false,
    ssl: {
      key: fs.readFileSync("cert/server.key", "utf8"),
      cert: fs.readFileSync("cert/server.crt", "utf8"),
    },
  })
  .listen($PORT);

console.log("proxy started");
```

ws-request (remote)

`wss://remote:$PORT` 주소로 요청함

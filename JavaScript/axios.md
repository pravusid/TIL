# HTTP client Library Axios

## 설치

`npm install --save axios`

## 사용

### 인스턴스 생성

```js
import axios from "axios";

// default configuration이 적용된 axios 인스턴스
const instance = axios;

// 사용자 지정 configuration과 default configuration을 mix-in 하여 인스턴스 생성
const instanceCustom = axios.create({
  baseURL: "https://some-domain.com/api/",
  timeout: 1000,
  headers: { "X-Custom-Header": "foobar" }
});
```

### 기본 method call

```js
axios.request(config);

axios.get(url[, config]);

axios.delete(url[, config]);

axios.head(url[, config]);

axios.options(url[, config]);

axios.post(url[, data[, config]]);

axios.put(url[, data[, config]]);

axios.patch(url[, data[, config]]);
```

## Response object key를 camelcase로 parsing

`axios/lib/defaults.js` 참조

```ts
import axios from "axios";

const camelize = (data: any) => {
  let result = data;
  if (typeof data === "string") {
    try {
      result = JSON.parse(data, (key, value) => {
        if (value && typeof value === "object") {
          for (const k in value) {
            if (/^[A-Z]/.test(k) && Object.hasOwnProperty.call(value, k)) {
              value[k.charAt(0).toLowerCase() + k.substring(1)] = value[k];
              delete value[k];
            }
          }
        }
        return value;
      });
    } catch (e) {
      /* Ignore */
    }
  }
  return result;
};

const createAxiosInstance = (() => {
  const instance = axios.create();
  instance.defaults.transformResponse = [camelize];
  return instance;
})();

export default createAxiosInstance;
```

## client TLS 설정

```js
const httpsAgent = new https.Agent({
  cert: fs.readFileSync("./usercert.pem"),
  key: fs.readFileSync("./key.pem"),
  passphrase: "some-passphrase"
});

axios.get(url, { httpsAgent });
```

## TLS 검증하지 않음

node 환경변수에서 TLS 미검증 선언

```js
process.env.NODE_TLS_REJECT_UNAUTHORIZED = 0;
```

또는 https agent 옵션사용

```js
import axios from 'axios';
import * as https from 'https';

axios.get(URL, { httpsAgent: new https.Agent({ rejectUnauthorized: false }) });
```

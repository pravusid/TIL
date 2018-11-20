# HTTP client Library Axios

## 설치

`npm install --save axios`

## 사용

### 인스턴스 생성

```js
import axios from 'axios';

// default configuration이 적용된 axios 인스턴스
const instance = axios;

// 사용자 지정 configuration과 default configuration을 mix-in 하여 인스턴스 생성
const instanceCustom = axios.create({
  baseURL: 'https://some-domain.com/api/',
  timeout: 1000,
  headers: {'X-Custom-Header': 'foobar'}
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

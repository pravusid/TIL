# OAuth2 JWT 인증

## 준비

- `axios`를 사용하여 http 요청을 보냄
- `query-string`을 사용하여 쿼리를 생성함

## 구현

토큰을 받았다면 이후 통신에서 http header의 Authorization에 token이 명시 되어야 함

`axios.defaults.headers.common.Authorization = 'bearer ${access_token}';`

브라우저를 새로고침 한다면 Token 변수가 초기화 되므로 브라우저 LocalStorage에 Token을 저장해야함

`Object`를 바로 저장할 수 없으므로 serialize 해서 저장함

`localStorage.user = qstr.stringify(token);`

### JWT Token만 발급받는 경우

로그인(토큰발급) API로 HTTP POST (username, password)를 보내고 response로 토큰을 받는다

### OAuth2 Grant Type: implicit

일반적으로 인증서버의 `요청-로그인-redirect` 페이지를 팝업으로 생성하고 진행 상황을 polling 하여 처리함

```js
import qstr from 'query-string';

function login() {
  const params = {
    response_type: 'token',
    client_id: 'client_id',
    redirect_uri: 'http://localhost:3000/login?success',
  };
  const originHost = 'http://localhost:3000';
  const url = `http://localhost:8080/oauth/authorize?${qstr.stringify(params)}`;
  const options = 'width=600, height=600';
  const popup = window.open(url, 'auth', options);
  popupWatcher(popup, originHost).then((param) => {
    // promise의 반환값인 token 변수 할당
  });
}

function popupWatcher(popup, exitUrl) {
  const parseUrl = document.createElement('a');
  parseUrl.href = exitUrl;
  return new Promise((resolve, reject) => {
    const polling = setInterval(() => {
      if (!popup || popup.closed || popup === undefined) {
        clearInterval(polling);
        reject(new Error('로그인 창 종료됨'));
      }
      try {
        if (popup.location.host === parseUrl.host) {
          const hash = qstr.parse(popup.location.hash.substring(1));
          if (hash.error) {
            reject(new Error(hash.error));
          } else {
            resolve(hash);
          }
          clearInterval(polling);
          popup.close();
        }
      } catch (error) {
        // cross origin frame exception
      }
    }, 250);
  });
}
```

### OAuth2 Grant Type: password

HTTP POST 요청에서 URL query string이 아닌 data로 grant_type, username, password 등을 보내면 오류 발생함
(다른 서버 Framework에서도 그런지는 확인이 필요함)

일반적으로 SPA에서 password Grant Type은 사용하지 않는다 (client_secret이 노출됨): implicit 사용

password Grant Type에서 client_secret 없이 (서버에서 `null` 값으로 설정) 사용할 수 있으나 OAuth2의 아키텍처에 따르면 이는 안티패턴임

```js
import qstr from 'query-string';

axios.post(`/oauth/token?grant_type=password&username=user&pasword=pwd`, null, {
    auth: {
      username: 'client_id',
      password: 'client_secret',
    },
  }).then((res) => {
    // response의 token 변수 할당
  }).catch((err) => {
    // error 처리
  });
```

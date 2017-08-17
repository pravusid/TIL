# Vue.js에서 기타 활용

## vue에서 jquery / bootstrap 사용

### JQuery

jQuery를 추가하려면 일단 외부 저장소에서 jQuery 라이브러리를 가져와야겠죠? 해당 프로젝트 루트에서 다음과 같은 명령어를 실행합니다.
`npm i --save-dev expose-loader`
`npm i --save jquery`

라이브러리를 설치하면서 package.json에 같이 추가하게 됩니다.
이 것으로 모든 준비는 끝났습니다. 코드 몇줄로 추가하고 그냥 사용하면 됩니다.
vue-cli로 프로젝트를 생성했다면 /project/src/main.js파일을 확인하실 수 있습니다. 해당 파일에 jQuery사용을 위해 한줄 코드를 추가합니다.
`import 'expose-loader?$!expose-loader?jQuery!jquery'`

### 부트스트랩

`npm i --save bootstrap`

```js
import 'expose-loader?$!expose-loader?jQuery!jquery'
// 위에서 추가했던 jQuery 밑에 코드를 작성하세요

import 'bootstrap'
import 'bootstrap/dist/css/bootstrap.min.css'
```

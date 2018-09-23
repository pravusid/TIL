# Vue.js 환경변수 / 모드

vue-cli-3의 scaffolding에는 의존성에 NODE_ENV가 포함되어 있다

```sh
.env                # loaded in all cases
.env.local          # loaded in all cases, ignored by git
.env.[mode]         # only loaded in specified mode
.env.[mode].local   # only loaded in specified mode, ignored by git
```

다음의 네이밍 컨벤션으로 환경번수를 분리할 수 있다.

## 환경변수 파일 (.env)

```env
VUE_APP_SECRET=secret
```

환경변수의 변수명은 반드시 `VUE_APP`으로 시작해야 번들링시 적용이 된다

## Mode 설정

vue-cli 3.0부터 `vue-cli-service`가 추가되었는데, 인자로 모드를 선언할 수 있다

`package.json` script를 다음과 같이 수정한다

`"dev-build": "vue-cli-service build --mode development",`

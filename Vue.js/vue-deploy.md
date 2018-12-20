# Vue.js 배포

## 환경변수

`vue-cli-service`에서는 `build` 명령어 실행시 기본환경변수를 불러옴

기본 환경변수: `.env`, `.env.production`, `.env.production.local`

## production 환경에서 source map 비활성화

프로젝트 루트 `vue.config.js` 파일 (webpack 설정을 override함)

```js
module.exports = {
  productionSourceMap: false
};
```

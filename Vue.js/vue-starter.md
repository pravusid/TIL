# Vue.js 시작하기

## 직접 CDN 링크

`<script src="https://unpkg.com/vue"></script>`

## vue-cli 사용

### vue와 vue-cli 패키지 설치

`npm install -g @vue/cli`

### vue-cli3으로 새로운 프로젝트 생성

<https://cli.vuejs.org/guide/>

`vue create <project-name>`

### modules 다운로드

`yarn` or `npm install`

## 개발환경 구축

### 크롬 확장기능 vue-devtools 사용

[vue-devtools](https://chrome.google.com/webstore/detail/vuejs-devtools/nhdogjmejiglipccpnnnanhbledajbpd)

### VSCode Extensions

- Vetur
- ESLint

VSCode Vue Formatting 관련 설정에 추가

```json
{
 "eslint.validate": [
    {
      "language": "vue",
      "autoFix": true
    },
    {
      "language": "html",
      "autoFix": true
    },
  ],
  "vetur.format.defaultFormatter.html": "prettyhtml",
}
```

## CSS

### import css (static file)

`App.vue`에 추가

```html
<style lang="css">
  @import './assets/yourstyles.css'
</style>
```

OR `index.html`에서 link도 가능

### css framework 추가

Bulma 설치 : `yarn add bulma`

Webpack sass-loader module 설치 : `yarn add node-sass sass-loader`

`src/assets/sass/main.scss` 파일을 생성해서 bulma를 불러온다

```scss
// ~는 Webpack/sass-loader가 node_modules 디렉토리로 인식
@import '~bulma/bulma'
```

`src/main.js` 파일에 추가

```js
// Require the main Sass manifest file
require('./assets/sass/main.scss');
```

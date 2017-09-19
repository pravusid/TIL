# Vue.js 시작하기

## 직접 CDN 링크

`<script src="https://unpkg.com/vue"></script>`

## vue-cli 사용

NPM으로 vue와 vue-cli 패키지 설치

`yarn global add vue vue-cli`

### vue-cli로 새로운 프로젝트 생성

`$ vue init <template-name> <project-name>`

템플릿으로는 webpack을 사용(starter)

```sh
vue init webpack vue-example-project
cd <project-name>
# install dependencies
npm install
# serve with hot reload at localhost:8080
npm run dev
# build for production with minification
npm run build
# build for production and view the bundle analyzer report
npm run build --report
# run unit tests
npm run unit
# run e2e tests
npm run e2e
# run all tests
npm test
```

### 프로젝트 생성시 설정

```sh
# Vue 빌드 선택 : 두 개 중에서 선택할 수 있으며, 기본선택은 Runtime + Compiler
# 두 번째 선택은 6KB의 가벼운 min+gzip으로 이루어져 있는 런타임전용. 템플릿은 .vue에서만 허용
? Vue build
- Runtime + Compiler: recommended for most users
- Runtime-only: about 6KB lighter min+gzip, but templates (or any Vue-specific HTML) are ONLY allowed in .vue files - render functions are required elsewhere
# vue-router사용여부
? Install vue-router? (Y/n)
# ESLint 적용여부
? Use ESLint to lint your code? (Y/n)
# ESLint 스타일 (AirBnB 사용 추천)
? Pick an ESLint preset
- Standard (https://github.com/feross/standard)
- Airbnb (https://github.com/airbnb/javascript)
- none (configure it yourself)
# 유닛테스트 Karma, Mocha 적용 여부
? Setup unit tests with Karma + Mocha? (Y/n)
# UI테스트 툴 Nightwatch 적용 여부
? Setup e2e tests with Nightwatch? (Y/n)
```

### modules 다운로드

`yarn` or `npm install`

### 디렉토리 구조

1. vue-cli로 만든 앱은 ./src 폴더의 main.js 에서 시작된다.
1. main.js에서 Vue 인스턴스를 생성한다.
1. 다른 라이브러리 (vue-router, vuex)를 추가하는 경우 여기에서 Vue 인스턴스 생성 전에 추가해 주면된다.

- 파일 / 디렉토리 구조
  - /src/main.js - Application entry point.
  - /src/App.vue - Base component
  - /src/router/index.js - Application routes
  - /src/components - All our components go here
  - /src/assets - All our image assets go here
  - /config - Application configuration information
  - /build - Build configuration information

#### `/src/main.js`

Entry Point and Root Vue Instance : import와 webpack entry

```js
import Vue from 'vue';
import App from './App';
import router from './router';

new Vue({
  el: '#app',
  router,
  template: '<App/>',
  components: { App },
});
```

- el : `index.html`의 application mount point (id가 app인 div의 DOM을 제어함)
- router : 라우터 사용 선언
- template : mount point에 적용될 템플릿 (template)
- components : 컴포넌트 (App.vue를 import해서 여기에서 사용함)

#### `/src/App.vue`

```html
<template>
  <div id="app">
    <img src="./assets/logo.png">
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  name: 'app',
};
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
```

#### `/dist/index.html`

빌드 결과물 시작점

```html
<!DOCTYPE html>
<html>

<head>
  <meta charset=utf-8>
  <title>Vue</title>
  <link href=/static/css/app.aa4a320cae73ccea747f91c11e724a37.css rel=stylesheet>
</head>

<body>
  <div id=app></div>
  <script type=text/javascript src=/static/js/manifest.0e912206617edce8a3e3.js></script>
  <script type=text/javascript src=/static/js/vendor.ae75c6b5bea60f5d8cec.js></script>
  <script type=text/javascript src=/static/js/app.12c3f867db1cd28fc91d.js></script>
</body>

</html>
```

## vue-cli 기본 dev-dependencies

vue-cli는 npm scripts 를 이용해서 필요한 기능을 실행할 수 있다.
최초 프로젝트를 만들때 선택한 내용들로 구성 되어 있고 나는 주로 ESLint에서 스타일은 airbnb로 선택한다.
npm 을 이용한 기능들은 아래와 같다.

- dev : 개발용 http 서버를 실행한다. 개발중에는 이 명령어로 실행하면 된다.
- build : build를 실행한다. 실행 후에 ./dist 에 완료된 파일들이 있다.
- unit : 유닛 테스트를 실행한다.
- e2e : 엔드 투 엔드 테스트를 실행한다.
- test : 위 두 테스트를 실행한다.
- lint : 소스코드에 대한 정적 테스트를 실행한다.

## 크롬 확장기능 vue-devtools 사용

[vue-devtools](https://chrome.google.com/webstore/detail/vuejs-devtools/nhdogjmejiglipccpnnnanhbledajbpd)

## import css (static file)

`App.vue`에 추가

```html
<style lang="css">
  @import './assets/yourstyles.css'
</style>
```

OR `index.html`에서 link도 가능

## css framework 추가

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

## eslint 사용

### 프로젝트에서 사용

```sh
sudo npm install --save-dev
eslint eslint-config-airbnb eslint-friendly-formatter eslint-loader eslint-plugin-html eslint-plugin-vue eslint-plugin-import eslint-plugin-node eslint-plugin-promise eslint-plugin-standard
```

### VS code 에서 global 사용

```sh
sudo npm -g install
eslint eslint-config-airbnb-base eslint-plugin-import eslint-plugin-html eslint-plugin-vue
```

.eslintrc.json 에 추가

```json
"extends": [
    "airbnb-base"
],
"plugins": [
    "vue/recommanded"
],
```

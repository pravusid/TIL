# Vue.js 시작하기

## 직접 CDN 링크

`<script src="https://unpkg.com/vue"></script>`

## vue-cli 사용

### vue와 vue-cli 패키지 설치

`npm install -g @vue/cli`

### vue-cli3으로 새로운 프로젝트 생성

<https://cli.vuejs.org/guide/>

`vue create <project-name>`

### 개발환경 port 변경

`package.json` 파일 수정

`"serve": "vue-cli-service serve --port 3000"`

### modules 다운로드

## 개발환경 구축

### 크롬 확장기능 vue-devtools 사용

[vue-devtools](https://chrome.google.com/webstore/detail/vuejs-devtools/nhdogjmejiglipccpnnnanhbledajbpd)

### VSCode Extensions

- ESLint
- Prettier
- Vetur

`.vscode/settings.json`

- vscode 내장 js, ts 문법 검사 비활성화
  - (babel-proposal 플러그인과 충돌)
  - <https://code.visualstudio.com/docs/languages/javascript#_how-do-i-disable-syntax-validation-when-using-nones6-constructs>
- vetur 내장 template, script 검사 비활성화 (prettier, eslint와 충돌)

```json
{
  "editor.formatOnSave": true,
  "javascript.validate.enable": false,
  "eslint.validate": [
    "javascript",
    "javascriptreact",
    "typescript",
    "typescriptreact",
    "html",
    "vue"
  ],
  "vetur.format.defaultFormatter.html": "prettier",
  "vetur.validation.template": false,
  "vetur.validation.script": false
}
```

eslint 설치

```sh
npm i -D eslint babel-eslint eslint-plugin-vue

npm i -D eslint-config-standard eslint-plugin-standard \
  eslint-plugin-promise eslint-plugin-import eslint-plugin-node

npm i -D prettier @vue/eslint-config-prettier eslint-plugin-prettier
```

`.eslintrc.js`

```js
module.exports = {
  root: true,
  env: {
    node: true
  },
  parser: 'vue-eslint-parser',
  parserOptions: {
    parser: 'babel-eslint'
  },
  extends: ['standard', 'plugin:vue/recommended', '@vue/prettier'],
  rules: {
    'no-console': process.env.NODE_ENV === 'production' ? 'error' : 'off',
    'no-debugger': process.env.NODE_ENV === 'production' ? 'error' : 'off',
    'space-before-function-paren': ['error', { anonymous: 'never', named: 'never', asyncArrow: 'always' }],
    'vue/max-attributes-per-line': 'off',
    'vue/singleline-html-element-content-newline': 'off'
  }
}
```

`.prettierrc`

```json
{
  "printWidth": 120,
  "semi": false,
  "singleQuote": true,
  "quoteProps": "consistent",
  "jsxSingleQuote": true,
  "endOfLine": "lf"
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

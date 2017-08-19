# Vue.js 시작하기

## Vue.js 설치

### 직접 CDN 링크

`<script src="https://unpkg.com/vue"></script>`

### vue-cli 사용

NPM으로 vue와 vue-cli 패키지 설치
  ```sh
  npm install vue
  # vue-cli 설치
  npm install --global vue-cli
  # "webpack" 템플릿을 이용해서 새 프로젝트 생성
  ```

#### vue-cli로 새로운 프로젝트 생성

`$ vue init <template-name> <project-name>`

템플릿으로는 webpack을 사용(starter), 이후 프로젝트 폴더에서 dependencies 다운로드
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

#### 프로젝트 생성시 설정

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

#### 디렉토리 구조

1. vue-cli로 만든 앱은 ./src 폴더의 main.js 에서 시작된다.
1. main.js에서 Vue 인스턴스를 생성한다.
1. 다른 라이브러리 (vue-router, vuex)를 추가하는 경우 여기에서 Vue 인스턴스 생성 전에 추가해 주면된다.

- 파일 / 디렉토리 구조
  - /src/main.js - Application entry point.
  - /src/App.vue - Base component
  - /router/index.js - Application routes
  - /src/components - All our components go here
  - /src/assets - All our image assets go here
  - /config - Application configuration information
  - /build - Build configuration information

- /src/main.js (Entry Point and Root Vue Instance) : import와 webpack entry, 빌드결과물이 여기로 부터 시작함

  root vue 인스턴스 선언부
  ```js
  new Vue({
    el: '#app',
    router,
    template: '<App/>',
  });
  ```

  main.js의 템플릿 구문을 사용해 선언적으로 DOM에 데이터를 렌더링한다
  ```html
  <div id="app">
      <App />
  </div>
  ```

### vue-cli 기능

vue-cli는 npm scripts 를 이용해서 필요한 기능을 실행할 수 있다.
최초 프로젝트를 만들때 선택한 내용들로 구성 되어 있고 나는 주로 ESLint에서 스타일을 none으로 선택하고 사용한다.
스타일을 none으로 하는 이유는 아직 standardjs와 airbnb 중 어떤 스타일이 더 나은지 결론을 내지 못했기 때문이다.
나머지 테스트에 대한 기능들은 모두 Yes를 선택한다.
npm 을 이용한 기능들은 아래와 같다.

- dev : 개발용 http 서버를 실행한다. 개발중에는 이 명령어로 실행하면 된다.
- build : build를 실행한다. 실행 후에 ./dist 에 완료된 파일들이 있다.
- unit : 유닛 테스트를 실행한다.
- e2e : 엔드 투 엔드 테스트를 실행한다.
- test : 위 두 테스트를 실행한다.
- lint : 소스코드에 대한 정적 테스트를 실행한다.

### 크롬 확장기능 vue-devtools 사용

[vue-devtools](https://chrome.google.com/webstore/detail/vuejs-devtools/nhdogjmejiglipccpnnnanhbledajbpd)

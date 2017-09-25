# 자바스크립트 변화 추세

## 자바스크립트 표준

- ECMAScript2015(ES6)
- 트랜스파일러 : 상위버전 JS 명세를 ECMAScript5(ES5)로 변환
  - Babel

## TypeScript, Flow

Typescript는 MS의 주도로 자바스크립트 superset으로 개발되고 있는 언어이다.

Flow는 Facebook 주도로 개발되고 있는 자바스크립트 정적타입 분석기이다.

## Node.js

자바스크립트 런타임 라이브러리, 클라이언트, 서버, 응용프로그램에 걸쳐 다양한 용도로 사용되고 있음.

## 패키지 관리자

- npm(Node package manager) : node.js 패키지 관리자였으나 현재는 사실상 자바스크립트 표준 패키지 관리자
- Bower : 성장 정체기, 아무래도 대세에서 멀어짐
- Yarn : 2016년 10월에 Facebook에서 새롭게 발표한 npm 클라이언트 == npm 저장소를 그대로 사용함 but, 별도의 자체 저장소도 보유

## 빌드툴 / 번들러

- 빌드툴 : 작업(task)을 기반으로 실행하여 결과물 출력
  - gulp
  - grunt
- 번들러 : 프로젝트에서 사용된 자원들의 상관관계를 정적으로 분석해 지정한 형태의 파일을 출력
  - webpack : 의존관계를 분석해 이것을 모은 엔트리파일을 만들고 엔트리파일을 컴파일하여 번들파일을 생성함
    - webpack loader : 변환에 사용됨 (javascript:babel, css:sass ...)

## 프레임워크 / 라이브러리

### 과거의 프레임워크 / 라이브러리

개인 개발자들의 프로젝트 위주였다

 - John Resig(jQuery)
 - Jeremy Ashkenas(Backbone, Underscore)
 - Thomas Fuchs(Zepto, script.aculo.us)
 - Mihai Baizon(Uglify)
 - Eric Schoffstall(Gulp)
 - Ben Alman(Grunt)

### 최근의 프레임워크 / 라이브러리

기업주도 오픈소스 프로젝트가 대세

- Angular
- React
- Vue.js

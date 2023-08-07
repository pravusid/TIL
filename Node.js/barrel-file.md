# Barrel File

<https://basarat.gitbook.io/typescript/main-1/barrel>

범위 내의 모듈을 re-export 하는 파일이다. 파일명으로 `index.js|ts`를 사용한다.

## 사용하지 말아야 할 이유

### 잘못된 디자인 결정

- [Node.js에 관해 후회하는 10가지 - Ryan Dahl - JSConf EU](https://youtu.be/M3BM9TB-8yA?t=835)
- <https://www.reddit.com/r/node/comments/128in46/why_is_indexjs_a_mistake_in_nodejs_ryan_dahl/>

> index.html을 생각하며 도입하였는데, 시간이 지나고 생각해보니 불필요한 기능이었다.

### 성능상의 문제

[Slow start times due to use of barrel files](https://github.com/jestjs/jest/issues/11234)

### 순환참조

[Why you should avoid index.js?](https://medium.com/@alonmiz1234/why-you-should-avoid-index-js-3321a9902120)

- import 할 때 원본 모듈이 아니라 barrel file 파일우선으로 import를 하다보면 순환참조가 발생하게 됨
- dynamic imports 사용할 때 의도하지 않은 import 동작이 발생할 수 있음

# Jest

Facebook 에서 주도하는 자바스크립트 테스트 툴 (React와 함께 성장)

## 기존의 JS 테스트 stack

- 테스트 프레임워크 : Mocha
- assertion(문법) : Chai
- mocking : Sinon
- coverage : Istanbul

## Jest for back-end

jest역시 서버용 테스트 툴로 사용가능하다.

### 설치

```sh
yarn add --dev jest eslint eslint-plugin-jest
yarn add --peer eslint-config-airbnb-base
```

### eslint 설정

`.eslintrc.js`

```js
module.exports = {
  extends: [
    'airbnb-base',
    'plugin:jest/recommanded',
  ],
  plugins: [
    'import',
    'jest',
  ],
  env: {
    node: true,
    'jest/globals': true,
  },
};
```

### jest설정

`packages.json`

```json
  "scripts": {
    "lint": "eslint src/**",
    "test": "jest src",
    "coverage": "jest --collectCoverageFrom=src/**.js --coverage src"
  },
```

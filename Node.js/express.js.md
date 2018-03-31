# Express.js

## 설치

### Express & middleware 설치

`yarn add express body-parser morgan rimraf`

* express: js server Framework
* body-paser: middleware for body parsing
* morgan: middleware for logging
* rimraf: rm -rf 명령을 위한 library

`yarn add --dev dotenv`

* dotenv: 환경변수 불러오기 위한 library

`yarn add --dev nodemon`

live reload 를 위한 nodemon

### eslint 설치

`yarn add --dev eslint eslint-plugin-import eslint-watch eslint-config-airbnb-base`

`eslint --init`

### Test Framework

```sh
yarn add --dev jest supertest
yarn add --dev eslint-plugin-jest
```

### babel 설치

```sh
yarn add babel-cli babel-plugin-transform-class-properties babel-plugin-transform-object-rest-spread babel-preset-env
yarn add --dev babel-eslint babel-jest babel-register
```

## 설정

### `packages.json`

```json
  "scripts": {
    "prestart": "yarn run -s build",
    "start": "node dist/index.js",
    "dev": "nodemon src/index.js --exec \"node -r dotenv/config -r babel-register\"",
    "clean": "rimraf dist",
    "build": "yarn run clean && mkdir -p dist && babel src -s -D -d dist",
    "test": "jest --watch",
    "lint": "esw -w src test"
  },
  "babel": {
    "presets": [
      [
        "env",
        {
          "targets": {
            "node": "current"
          }
        }
      ]
    ],
    "plugins": [
      "transform-object-rest-spread",
      "transform-class-properties"
    ]
  },
  "jest": {
    "testEnvironment": "node"
  }
```

### `.eslintrc.js`

```js
{
  parser: 'babel-eslint',
  plugins: ['import', 'jest'],
  parserOptions: {
    ecmaVersion: 2017,
    sourceType: 'module',
  },
  env: {
    node: true,
    jest: true,
  },
  extends: ['eslint:recommended'],
  rules: {
    'jest/no-focused-tests': 2,
    'jest/no-identical-title': 2,
  },
}
```

## template engine

### pug

`yarn add pug`

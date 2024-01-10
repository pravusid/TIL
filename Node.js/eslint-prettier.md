# Linting and Formatting

> <https://eslint.org/blog/2023/10/deprecating-formatting-rules/> [(번역)](https://velog.io/@typo/deprecation-of-formatting-rules)

## 차이

- [`prettier-eslint`](https://github.com/prettier/prettier-eslint)

  - 구분: A JavaScript module exporting a single function
  - 역할: Runs the code (string) through `prettier` then `eslint --fix`. The output is also a string.
  - 사용법: Either calling the function in your code or via [`prettier-eslint-cli`](https://github.com/prettier/prettier-eslint-cli) if you prefer the command line.
  - 최종결과물 Prettier 적용: Depends on your ESLint config
  - `prettier` 커맨드 별도 실행필요: No
  - 다른 것 사용필요: No

- [`eslint-plugin-prettier`](https://github.com/prettier/eslint-plugin-prettier)

  - 구분: An ESLint plugin
  - 역할: Plugins usually contain implementations for additional rules that ESLint will check for. This plugin uses Prettier under the hood and will raise ESLint errors when your code differs from Prettier's expected output.
  - 사용법: Add it to your `.eslintrc`.
  - 최종결과물 Prettier 적용: Yes
  - `prettier` 커맨드 별도 실행필요: No
  - 다른 것 사용필요: You may want to turn off conflicting rules using `eslint-config-prettier`.

- [`eslint-config-prettier`](https://github.com/prettier/eslint-config-prettier)

  - 구분: An ESLint configuration
  - 역할: This config turns off formatting-related rules that might conflict with Prettier, allowing you to use Prettier with other ESLint configs like [`eslint-config-airbnb`](https://www.npmjs.com/package/eslint-config-airbnb).
  - 사용법: Add it to your `.eslintrc`.
  - 최종결과물 Prettier 적용: Yes
  - `prettier` 커맨드 별도 실행필요: Yes
  - 다른 것 사용필요: No

## @javascript-eslint

```sh
npm i --save-dev eslint
npm i --save-dev babel-eslint
```

`.eslintrc.json`

<https://eslint.org/docs/user-guide/configuring#specifying-environments>

```json
{
  "parser": "babel-eslint",
  "parserOptions": {
    "ecmaVersion": 2017
  },
  "env": {
    "es6": true,
    "node": true
  },
  "extends": ["eslint:recommended"]
}
```

## @typescript-eslint

```sh
npm i --save-dev eslint
npm i --save-dev @typescript-eslint/parser @typescript-eslint/eslint-plugin
npm i --save-dev eslint-plugin-jest
```

- @typescript-eslint/parser: ESLint가 TypeScript 코드를 처리할 수 있게 함
- @typescript-eslint/eslint-plugin: TypeScript에 맞춘 ESLint rule의 모음

<https://github.com/typescript-eslint/typescript-eslint>

<https://github.com/typescript-eslint/typescript-eslint/tree/master/packages/eslint-plugin>

<https://github.com/typescript-eslint/typescript-eslint/tree/master/packages/eslint-plugin/src/configs>

- 실행: `npx eslint --ext .ts src`
- 비활성화: `// eslint-disable-next-line {rule}`

`.eslintrc.json`

```json
{
  "parser": "@typescript-eslint/parser",
  "parserOptions": {
    "project": "./tsconfig.json"
  },
  "env": {
    "node": true
  },
  "extends": [
    "eslint:recommended",
    "plugin:@typescript-eslint/recommended",
    "plugin:@typescript-eslint/recommended-requiring-type-checking"
  ],
  "rules": {
    "no-empty-function": "off",
    "@typescript-eslint/no-empty-function": ["error", { "allow": ["private-constructors", "protected-constructors"] }],
    "@typescript-eslint/no-namespace": "off"
  },
  "overrides": [
    {
      "files": ["src/**/*.spec.*", "*.js"],
      "parserOptions": {
        "project": "./tsconfig.spec.json"
      },
      "env": {
        "jest": true
      },
      "plugins": ["jest"],
      "rules": {}
    }
  ]
}
```

> `@typescript-eslint/eslint-recommended` config is automatically included if you use either the `recommended` or `recommended-requiring-type-checking` configs.

`recommended-requiring-type-checking` 사용하지 않는 경우 다음 규칙 추가

```json
{
  "rules": {
    // ...
    "no-unused-vars": "off",
    "require-await": "off",
    "no-return-await": "off",
    "@typescript-eslint/await-thenable": "warn",
    "@typescript-eslint/no-misused-promises": ["error", { "checksVoidReturn": false }],
    "@typescript-eslint/no-unused-vars": "warn",
    "@typescript-eslint/require-await": "warn",
    "@typescript-eslint/return-await": ["warn", "in-try-catch"]
    // ...
  }
}
```

- recommended without typecheck: `plugin:@typescript-eslint/recommended`
- recommended requiring typecheck: `plugin:@typescript-eslint/recommended-requiring-type-checking`

## Prettier

`npm i --save-dev prettier`

`.prettierrc`

```json
{
  "printWidth": 120,
  "tabWidth": 2,
  "useTabs": false,
  "semi": true,
  "singleQuote": true,
  "quoteProps": "consistent",
  "jsxSingleQuote": true,
  "trailingComma": "all",
  "bracketSpacing": true,
  "bracketSameLine": false,
  "arrowParens": "always",
  "endOfLine": "lf"
}
```

## ESLint Prettier Integration

### 방법1: prettier + (eslint | @typescript-eslint)

- <https://prettier.io/docs/en/eslint.html>
- <https://github.com/prettier/eslint-plugin-prettier#installation>

> add `plugin:prettier/recommended` as the **last extension** in your .eslintrc.json

plugin 사용만으로는 eslint formatting rules와 prettier rules가 충돌하므로, eslint-config-prettier를 함께 사용한다

```sh
npm i --save-dev prettier
npm i --save-dev eslint-plugin-prettier eslint-config-prettier
```

[8.0.0 이후 모든 플러그인 규칙을 통합하였다](https://github.com/prettier/eslint-config-prettier/blob/main/CHANGELOG.md#version-800-2021-02-21). 따라서 JS, TS 관계 없이 동일한 설정으로 사용가능하다.

> If you use `eslint-plugin-prettier`, all you need is [`plugin:prettier/recommended`](https://github.com/prettier/eslint-plugin-prettier#recommended-configuration)

`.eslintrc.json` 설정에 추가 (plugin + enable all the recommended rules at once)

```json
{
  "extends": ["plugin:prettier/recommended"]
}
```

### 방법2: prettier-eslint

eslint 설정이 끝난 상태에서 (prettier-config & plugin 설정을 하지 않았음) 설치함

`npm i --save-dev prettier prettier-eslint prettier-eslint-cli`

실행

`prettier-eslint 'src/**/*.js'`

## VSCode

### 설정

> 기본설정은 다음 문서 참고: [Visual Studio Code](../Tools/vs-code.md)

내장 formatter 대신 prettier 사용

```json
{
  "[javascript]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "[typescript]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "javascript.format.enable": false,
  "typescript.format.enable": false
}
```

eslint에서 js, ts 처리

```json
{
  "eslint.validate": ["javascript", "typescript"]
}
```

optional: prettier-eslint Integration

```json
{
  "prettier.eslintIntegration": true
}
```

### VSCode 확장 기능

- ESLint

  - eslint 포함하지 않으므로 global or local 설치 필요

- Prettier – Code Formatter

  - Prettier 확장은 prettier, prettier-eslint, prettier-tslint 포함
  - Prettier 확장은 configuration을 npm global에서 가져올 수 없고 local에서만 가져옴

## 참고

### airbnb 규칙에서 'ForOfStatement' is not allowed

<https://github.com/airbnb/javascript/issues/1271>

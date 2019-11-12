# Linting and Formatting

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

## @javascript/ESLint

`npm i --save-dev eslint eslint-config-airbnb-base eslint-plugin-import`

`npm i --save-dev babel-eslint`

`.eslintrc.json`

```json
{
  "parser": "babel-eslint",
  "env": {
    "node": true
  },
  "extends": ["airbnb-base"]
}
```

## @typescript/ESLint

`npm i --save-dev eslint eslint-config-airbnb-base eslint-plugin-import`

`npm i --save-dev @typescript-eslint/parser @typescript-eslint/eslint-plugin`

`npm i --save-dev eslint-plugin-jest`

- @typescript-eslint/parser: ESLint가 TypeScript 코드를 처리할 수 있게 함
- @typescript-eslint/eslint-plugin: TypeScript에 맞춘 ESLint rule의 모음

<https://github.com/typescript-eslint/typescript-eslint>

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
  "extends": ["airbnb-base", "plugin:@typescript-eslint/recommended"],
  "rules": {
    "class-methods-use-this": "off",
    "lines-between-class-members": ["error", "always", { "exceptAfterSingleLine": true }],
    "no-empty-function": "off",
    "no-undef": "off",
    "no-useless-constructor": "off",
    "import/no-unresolved": "off",
    "import/prefer-default-export": "off",
    "@typescript-eslint/explicit-function-return-type": "off",
    "@typescript-eslint/explicit-member-accessibility": "off",
    "@typescript-eslint/no-explicit-any": "off",
    "@typescript-eslint/no-namespace": "off",
    "@typescript-eslint/no-parameter-properties": "off",
    "@typescript-eslint/no-useless-constructor": ["error"]
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
      "rules": {
        "dot-notation": "off",
        "global-require": "off",
        "@typescript-eslint/no-non-null-assertion": "off",
        "@typescript-eslint/no-var-requires": "off"
      }
    }
  ]
}
```

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
  "jsxBracketSameLine": false,
  "arrowParens": "avoid",
  "endOfLine": "lf"
}
```

## ESLint Prettier Integration

### 방법1: prettier + eslint

#### prettier in @javascript/ESLint

<https://prettier.io/docs/en/eslint.html>

plugin 사용만으로는 eslint formatting rules와 prettier rules가 충돌하므로, eslint-config-prettier를 함께 사용한다

`npm i --save-dev prettier`

`npm i --save-dev eslint-plugin-prettier eslint-config-prettier`

`.eslintrc.json` 설정에 추가 (plugin + enable all the recommended rules at once)

```json
{
  "extends": ["plugin:prettier/recommended"]
}
```

#### prettier in @typescript/ESLint

`npm i --save-dev prettier`

`npm i --save-dev eslint-plugin-prettier eslint-config-prettier`

<https://github.com/typescript-eslint/typescript-eslint/tree/master/packages/eslint-plugin>

`.eslintrc.json`

```json
{
  "extends": ["plugin:prettier/recommended", "prettier/@typescript-eslint"]
}
```

### 방법2: prettier-eslint

eslint 설정이 끝난 상태에서(prettier-config & plugin 설정을 하지 않았음)

`npm i --save-dev prettier prettier-eslint prettier-eslint-cli`

`prettier-eslint 'src/**/*.js'`

## VSCode

### 설정

내장 formatter를 비활성화

```json
{
  "javascript.format.enable": false,
  "typescript.format.enable": false
}
```

eslint에서 typescript 처리

```json
{
  "eslint.validate": [
    { "language": "typescript", "autoFix": true },
    { "language": "typescriptreact", "autoFix": true }
  ]
}
```

prettier-eslint Integration

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

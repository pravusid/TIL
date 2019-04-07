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

## ESLint

`npm install --save-dev eslint eslint-config-airbnb-base eslint-plugin-import babel-eslint`

`.eslintrc.js`

```js
module.exports = {
  extends: ['airbnb-base'],
  parser: 'babel-eslint',
  env: {
    node: true,
  },
};
```

## Prettier

`npm install --save-dev prettier`

`.prettierrc`

```json
{
  "printWidth": 100,
  "tabWidth": 2,
  "singleQuote": true,
  "trailingComma": "all",
  "bracketSpacing": true,
  "semi": true,
  "useTabs": false,
  "arrowParens": "avoid",
  "endOfLine": "lf"
}
```

## ESLint Prettier Integration

### 방법1: prettier + eslint

#### @javascript/ESLint

<https://prettier.io/docs/en/eslint.html>

plugin 사용만으로는 eslint formatting rules와 prettier rules가 충돌하므로, eslint-config-prettier를 함께 사용한다

DevDependencies 추가: `npm i --save-dev prettier eslint-plugin-prettier eslint-config-prettier`

`.eslintrc.js` 설정에 추가

```js
{
  extends: ["plugin:prettier/recommended"],
}
```

#### @typescript/ESLint

`npm install --save-dev @typescript-eslint/parser @typescript-eslint/eslint-plugin`

`npm install --save-dev prettier eslint-config-prettier`

<https://github.com/typescript-eslint/typescript-eslint/tree/master/packages/parser>

<https://github.com/typescript-eslint/typescript-eslint/tree/master/packages/eslint-plugin>

- 실행: `npx eslint --ext .ts src`
- 비활성화: `// eslint-disable-next-line @typescript-eslint/no-for-in-array`

`.eslintrc.js`

```js
module.exports = {
  parser: '@typescript-eslint/parser',
  parserOptions: {
    project: './tsconfig.json',
  },
  env: {
    node: true,
  },
  extends: [
    'airbnb-base',
    'plugin:@typescript-eslint/recommended',
    'prettier',
    'prettier/@typescript-eslint',
  ],
  rules: {
    'import/prefer-default-export': false,
    '@typescript-eslint/explicit-function-return-type': 'off',
    '@typescript-eslint/no-explicit-any': 'off',
  },
  overrides: [
    {
      files: ['test/**/*.ts', 'test/**/*.tsx'],
      env: {
        jest: true,
      },
      plugins: ['jest'],
      rules: {
        '@typescript-eslint/no-non-null-assertion': 'off',
        '@typescript-eslint/no-object-literal-type-assertion': 'off',
      },
    },
  ],
};
```

### 방법2: prettier-eslint

eslint 설정이 끝난 상태에서(prettier-config | plugin 설정을 하지 않았음)

`npm i --save-dev prettier prettier-eslint prettier-eslint-cli`

`prettier-eslint 'src/**/*.js'`

## VSCode

### 설정

내장 formatter를 비활성화

```json
{
  "javascript.format.enable": false,
  "typescript.format.enable": false,
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

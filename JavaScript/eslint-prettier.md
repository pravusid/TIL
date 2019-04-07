# Linting and Formatting

## 차이

| | [`prettier-eslint`](https://github.com/prettier/prettier-eslint) | [`eslint-plugin-prettier`](https://github.com/prettier/eslint-plugin-prettier) | [`eslint-config-prettier`](https://github.com/prettier/eslint-config-prettier) |
| :---: | --- | --- | --- |
| 구분 | A JavaScript module exporting a single function | An ESLint plugin | An ESLint configuration |
| 역할 | Runs the code (string) through `prettier` then `eslint --fix`. The output is also a string. | Plugins usually contain implementations for additional rules that ESLint will check for. This plugin uses Prettier under the hood and will raise ESLint errors when your code differs from Prettier's expected output. | This config turns off formatting-related rules that might conflict with Prettier, allowing you to use Prettier with other ESLint configs like [`eslint-config-airbnb`](https://www.npmjs.com/package/eslint-config-airbnb). |
| 사용법 | Either calling the function in your code or via [`prettier-eslint-cli`](https://github.com/prettier/prettier-eslint-cli) if you prefer the command line. | Add it to your `.eslintrc`. | Add it to your `.eslintrc`. |
| 최종결과물의 Prettier 적용 여부 | Depends on your ESLint config | Yes | Yes |
| `prettier` 커맨드를 별도로 실행하는가? | No | No | Yes |
| 다른 것을 추가로 사용해야 하는가?| No | You may want to turn off conflicting rules using `eslint-config-prettier`. | No |

## ESLint

`npm install --save-dev eslint eslint-config-airbnb-base eslint-plugin-import`

## 방법1: Prettier + ESLint

### ESLint @javascript

<https://prettier.io/docs/en/eslint.html>

plugin 사용만으로는 eslint formatting rules와 prettier rules가 충돌하므로, eslint-config-prettier를 함께 사용한다

DevDependencies 추가: `npm i -D prettier eslint-plugin-prettier eslint-config-prettier`

`.eslintrc.json` 설정에 추가

```json
{
  "extends": ["plugin:prettier/recommended"]
}
```

### ESLint @typescript

`npm install --save-dev @typescript-eslint/parser @typescript-eslint/eslint-plugin`

`npm install --save-dev prettier eslint-config-prettier`

<https://github.com/typescript-eslint/typescript-eslint/tree/master/packages/parser>

<https://github.com/typescript-eslint/typescript-eslint/tree/master/packages/eslint-plugin>

- 실행: `npx eslint --ext .ts src`
- 비활성화: `// eslint-disable-next-line @typescript-eslint/no-for-in-array`

## 방법2: prettier-eslint

`npm i -D prettier prettier-eslint prettier-eslint-cli`

`prettier-eslint 'src/**/*.js'`

## VSCode

우선 내장 formatter를 비활성화 한다

```json
{
  "javascript.format.enable": false,
  "typescript.format.enable": false,
}
```

```json
{
  "eslint.validate": [
    { "language": "typescript", "autoFix": true },
    { "language": "typescriptreact", "autoFix": true }
  ]
}
```

VSCode 확장 기능 사용: Marketplace에서

- ESLint 설치

- Prettier – Code Formatter 설치
  - Prettier 확장은 prettier, prettier-eslint | prettier-tslint 포함
  - Prettier 확장은 configuration을 npm global에서 가져올 수 없고 local에서만 가져옴

```json
{
  "prettier.eslintIntegration": true
}
```

## 설정파일 예시

### `.eslintrc.json`

```js
module.exports = {
  extends: ['airbnb-base'],
  parser: 'babel-eslint',
  plugins: ['import'],
  parserOptions: {
    ecmaVersion: 2017,
    sourceType: 'module',
  },
  env: {
    node: true,
  },
};
```

### `.eslintrc.json` @typescript

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

### `.prettierrc`

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

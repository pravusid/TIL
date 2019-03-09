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

## 방법1: Prettier + (ESLint | TSLint)

<https://prettier.io/docs/en/eslint.html>

plugin 사용만으로는 eslint formatting rules와 prettier rules가 충돌하므로, eslint-config-prettier를 함께 사용한다

DevDependencies 추가: `npm i -D prettier eslint-plugin-prettier eslint-config-prettier`

`.eslintrc.json` 설정에 추가

```json
{
  "extends": ["plugin:prettier/recommended"]
}
```

DevDependencies 추가: `npm i -D prettier tslint-plugin-prettier tslint-config-prettier`

`.tslint.json` 설정에 추가

```json
{
  "extends": [
    "tslint-plugin-prettier",
    "tslint-config-prettier"
  ],
  "rules": {
    "prettier": true
  }
}
```

## 방법2: prettier-eslint | prettier-tslint

`prettier-eslint-cli` 또는 `prettier-tslint`의 CLI를 이용하거나 VSCode의 확장기능을 이용할 수 있다

VSCode 확장 기능 사용: Marketplace에서

- ESLint | TSLint 설치

- Prettier – Code Formatter 설치
  - Prettier 확장은 prettier, prettier-eslint | prettier-tslint 포함
  - Prettier 확장은 configuration을 npm global에서 가져올 수 없고 local에서만 가져옴

VSCode 설정에서 다음을 추가

```json
{
  "javascript.format.enable": false,
  "typescript.format.enable": false,

  "prettier.eslintIntegration": true,
  "prettier.tslintIntegration": true,
}
```

## 설정파일 예시

### `.eslintrc.json`

```json
{
  "extends": ["airbnb-base"],
  "parser": "babel-eslint",
  "plugins": ["import"],
  "parserOptions": {
    "ecmaVersion": 2017,
    "sourceType": "module"
  },
  "env": {
    "node": true
  }
}
```

### `tslint.json`

```json
{
  "defaultSeverity": "error",
  "extends": ["tslint-config-airbnb"],
  "jsRules": {},
  "rules": {},
  "rulesDirectory": []
}
```

### `tslint.json` w/ `prettier-tslint`

```json
{
  "defaultSeverity": "error",
  "extends": "tslint-config-airbnb",
  "jsRules": {},
  "rules": {
    "align": false,
    "max-line-length": [
      true,
      { "limit": 100, "ignore-pattern": "^import |^export {(.*?)}" }
    ]
  },
  "rulesDirectory": []
}
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

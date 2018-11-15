# Linting and Formatting

ESLint / TSLint configuration은 별도 설정해야 함 (eslint-config-airbnb-base ...)

## 차이

| | [`prettier-eslint`](https://github.com/prettier/prettier-eslint) | [`eslint-plugin-prettier`](https://github.com/prettier/eslint-plugin-prettier) | [`eslint-config-prettier`](https://github.com/prettier/eslint-config-prettier) |
|--|:--:|:--:|:--:|
| 구분 | A JavaScript module exporting a single function. | An ESLint plugin. | An ESLint configuration. |
| 역할 | Runs the code (string) through `prettier` then `eslint --fix`. The output is also a string. | Plugins usually contain implementations for additional rules that ESLint will check for. This plugin uses Prettier under the hood and will raise ESLint errors when your code differs from Prettier's expected output. | This config turns off formatting-related rules that might conflict with Prettier, allowing you to use Prettier with other ESLint configs like [`eslint-config-airbnb`](https://www.npmjs.com/package/eslint-config-airbnb). | 
| 사용 | Either calling the function in your code or via [`prettier-eslint-cli`](https://github.com/prettier/prettier-eslint-cli) if you prefer the command line. | Add it to your `.eslintrc`. | Add it to your `.eslintrc`. |
| Is the final output Prettier compliant? | Depends on your ESLint config | Yes | Yes |
| Do you need to run `prettier` command separately? | No | No | Yes |
| Do you need to use anything else? | No | You may want to turn off conflicting rules using `eslint-config-prettier`. | No |

## 방법1: Prettier + (ESLint / TSLint)

<https://prettier.io/docs/en/eslint.html>

플러그인 사용만으로는 eslint rules와 prettier rules가 충돌하므로, config-prettier를 함께 사용할 수 있다

DevDependencies 추가: `npm i -D eslint-plugin-prettier eslint-config-prettier`

`.eslintrc.json` 설정 변경

```json
{
  "extends": ["plugin:prettier/recommended"]
}
```

## 방법2: prettier-eslint / prettier-tslint

`prettier-eslint-cli` 또는 `prettier-tslint`의 CLI를 이용하거나 VSCode의 확장기능을 이용할 수 있다

VSCode 확장 기능 사용: Marketplace에서

1. ESLint 설치
2. TSLint 설치
3. Prettier – Code Formatter 설치

VSCode 설정에서 다음을 입력

```json
{
  "javascript.format.enable": false,
  "typescript.format.enable": false,
}
```

### VSCode extension: Prettier + (ESLint / TSLint)

```sh
npm i -D prettier-eslint
npm i -D prettier-tslint
```

VSC 설정에서 다음을 입력

```json
{
  "prettier.eslintIntegration": true,
  "prettier.tslintIntegration": true,
}
```

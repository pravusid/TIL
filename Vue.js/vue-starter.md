# Vue.js 시작하기

## vue-cli

<https://cli.vuejs.org/>

## 크롬 확장기능 vue-devtools 사용

[vue-devtools](https://chrome.google.com/webstore/detail/vuejs-devtools/nhdogjmejiglipccpnnnanhbledajbpd)

## VSCode Extensions

- [Vetur](https://marketplace.visualstudio.com/items?itemName=octref.vetur)
- [ESLint](https://marketplace.visualstudio.com/items?itemName=dbaeumer.vscode-eslint)
- [Prettier](https://marketplace.visualstudio.com/items?itemName=esbenp.prettier-vscode)

`.vscode/settings.json`

- vscode 내장 js, ts 문법 검사 비활성화
- vetur 내장 template 검사 비활성화 (eslint와 충돌): <https://vuejs.github.io/vetur/guide/linting-error.html#linting>

> 기본설정은 다음 문서 참고: [Visual Studio Code](../Tools/vs-code.md)

```json
{
  "eslint.validate": ["javascript", "javascriptreact", "typescript", "typescriptreact", "html", "vue"],
  "vetur.validation.template": false
}
```

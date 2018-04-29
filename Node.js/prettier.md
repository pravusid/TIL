# Prettier

## VSCode extension

VSCode 마켓플레이스에서 Prettier – Code Formatter 설치

## Prettier ESLint 연동

`yarn add --dev prettier-eslint`

작업 영역 설정에 다음을 입력

```json
{
  "editor.formatOnSave": true,
  "javascript.format.enable": false,
  "prettier.eslintIntegration": true
}
```

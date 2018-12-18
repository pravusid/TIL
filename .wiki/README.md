# 개인 Wiki

매일 학습한 내용을 기록한 [TIL](https://github.com/pravusid/TIL)을 Vue.js 정적 웹사이트 생성기인 vuepress로 출력하여 개인 wiki로 활용

## Git hooks

`.git/hooks/pre-push` 훅을 사용한다

```sh
 #!/bin/bash

echo "=== 개인 wiki 배포 스크립트 실행 ==="

cd .wiki
source ./deploy.sh
```

## deploy script

[deploy.sh](deploy.sh)

## Vuepress 설정

`docs/.vuepress/config.js`

```js
module.exports = {
  base: "/wiki/",
  title: "TIL wiki",
  themeConfig: {
    repo: 'pravusid/TIL',
    sidebar: 'auto',
    searchMaxSuggestions: 10,
  },
  markdown: {
    lineNumbers: true
  }
}
```

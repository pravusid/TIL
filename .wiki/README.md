# 개인 Wiki

> Markdown + Github Pages + Node.js + Vuepress + Git Hooks + Shell Script

매일 학습한 내용을 기록한 [TIL](https://github.com/pravusid/TIL)을 Vue.js 정적 웹사이트 생성기인 vuepress로 출력하여 개인 wiki로 활용

## Git hooks

`.git/hooks/pre-push` 훅을 사용한다

git hook을 실행하지 않으려면 `git push --no-verify` 옵션을 사용한다

```sh
 #!/bin/bash

echo "=== 개인 wiki 배포 스크립트 실행 ==="

cd .wiki
source ./deploy.sh
```

## Index 생성

[첫 페이지에서 문서를 링크하기 위해서 디렉토리 구조 파싱 후 Index 생성](create.index.js)

## deploy script

[deploy.sh](deploy.sh)

## Vuepress 설정

[config.js](docs/.vuepress/config.js)

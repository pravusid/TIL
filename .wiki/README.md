# 개인 Wiki

> Markdown + Node.js + Vuepress ++ Github Pages + (Shell Script | Github Actions)

매일 학습한 내용을 기록한 [TIL](https://github.com/pravusid/TIL)을 Vue.js 정적 웹사이트 생성기인 vuepress로 출력하여 개인 wiki로 활용

## Index 생성

[첫 페이지에서 문서를 링크하기 위해서 디렉토리 구조 파싱 후 Index 생성](create.index.js)

## Vuepress 설정

[config.js](docs/.vuepress/config.js)

## Deploy

직접 Github Pages를 배포하거나 Github Actions를 사용함

### Deployment Script

[deploy.sh](deploy.sh)

### Github Actions

Github Actions를 사용하여 원격 저장소에 Github Pages를 배포한다

[Github Actions Script](../.github/workflows/wiki.yml)

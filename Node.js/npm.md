# npm (node.js package manager)

## npm 기초명령어

### 의존성 추가

1. 'npm install': ./node_modules 패키지 설치
1. 'npm install --save / --save-dev` ./node_modules에 패키지 설치 + ./package.json 업데이트

운영이 아닌 개발 단계에서만 필요한 의존성 모듈들은 devDependencies 에 설치

peerDependencies는 현재 모듈과 다른 모듈간의 호환성 표시 (npm install시 node_modules에 설치되지 않음)

### npm삭제, npm global modules 모두 삭제

`sudo npm list -g --depth=0. | awk -F ' ' '{print $2}' | awk -F '@' '{print $1}'  | sudo xargs npm remove -g`

## yarn

### 설치

우분투/데비안에서 설치

```sh
curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | sudo apt-key add -
echo "deb https://dl.yarnpkg.com/debian/ stable main" | sudo tee /etc/apt/sources.list.d/yarn.list
sudo apt-get update && sudo apt-get install yarn
```

Note: Ubuntu 17.04 comes with cmdtest installed by default. If you’re getting errors from installing yarn, you may want to run sudo apt remove cmdtest first. Refer to this for more information.

**설치이후 global package binaries 경로를 환경변수로 설정해주어야 한다. 기본경로는 `~/.yarn`이고 경로를 변경하면 `경로/bin`에 바이너리가 연결된다**

yarn global bin will output the location where Yarn will install symlinks to your installed binaries. You can configure the base location with yarn config set prefix <filepath>.

### 기본 명령어

- `yarn` == `npm install` packages.json에 명시된 의존성 패키지를 다운로드/설치 한다.
- `yarn (global) add/bin/list/remove/upgrade [--prefix]`

### yarn add prefix

-  `--dev`(-D) : devDependencies
-  `--peer`(-P) : peerDependencies
-  `--optional`(-O) : optionalDependencies에
-  `--exact`(-E) : 명시한 버전과 정확한 경우에만 설치
-  `--tilde`(-T) : 명시한 버전과 같은 minor 버전의 최신버전 설치 (버전 세 번째 자리)


## packages

### axios

http 클라이언트 `npm install --save axios`

### bootstrap

`npm i --save bootstrap`

```js
import 'bootstrap'
import 'bootstrap/dist/css/bootstrap.min.css'
```

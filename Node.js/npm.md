# npm (node.js package manager)

## npm 기초명령어

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

`yarn` == `npm install` packages.json에 명시된 의존성 패키지를 다운로드/설치 한다.

`yarn (global) add/bin/list/remove/upgrade [--prefix]`

### yarn add prefix

 --dev(-D) : devDependencies

 --peer(-P) : peerDependencies

 --optional(-O) : optionalDependencies에

 --exact(-E) : 명시한 버전과 정확한 경우에만 설치

 --tilde(-T) : 명시한 버전과 같은 minor 버전의 최신버전 설치 (버전 세 번째 자리)

## packages

### axios

http 클라이언트 `npm install --save axios`

### JQuery

`npm i --save-dev expose-loader`
`npm i --save jquery`

라이브러리를 설치하면서 package.json에 같이 추가하게 됩니다.

vue-cli로 프로젝트를 생성했다면 `/project/src/main.js` 파일에서 import
`import 'expose-loader?$!expose-loader?jQuery!jquery'`

### bootstrap

`npm i --save bootstrap`

```js
import 'bootstrap'
import 'bootstrap/dist/css/bootstrap.min.css'
```
# npm (node.js package manager)

## npm

### npm install

`npm install --only=production` (or `--only=prod`): devDependencies 설치하지 않음

> The npm install command installs package files based on dependencies and devDependencies in package.json.

### npm ci

npm 5.7.0 이상 사용가능한 커맨드

> package-lock.json takes precedence over package.json for npm ci.

`npm ci`는 `package-lock.json` 파일을 우선으로 하여 패키지를 설치함

> When you run ci command, all of the node_modules folder installed by npm is deleted by default, and package is reinstalled based on package-lock.json.

### version in `package.json`

<https://docs.npmjs.com/misc/semver>

`[major, minor, patch]`

- `version`: 정확히 일치하는 버전
- `>version`: 특정 버전 초과
- `>=version`: 특정 버전 이상
- `<version`: 특정 버전 미만
- `<=version`: 특정 버전 이하
- `~version`: (tilde) minor 있으면 patch 변경가능, minor 없으면 minor 변경가능, m.m.p 인 경우 해당 버전 이상
- `^version`: (caret) m.m.p 버전에서 0이 아닌 가장 왼쪽 버전은 변경 불가, m.m.p 인 경우 해당 버전 이상
- `1.2.x`: 1.2.0, 1.2.1, etc ... 그러나 1.3.0 불가능
- `http://...`: <https://docs.npmjs.com/files/package.json#urls-as-dependencies>
- `*`: 아무 버전이나 가능
- `""`: 비어있는 문자열 `*`와 같음
- `version1 - version2`: between 버전1 and 버전2
- `range1 || range2`: 두 버전 범위중 하나
- `git...`: <https://docs.npmjs.com/files/package.json#git-urls-as-dependencies>
- `user/repo`: <https://docs.npmjs.com/files/package.json#github-urls>
- `tag`: <https://docs.npmjs.com/cli/dist-tag>
- `path/path/path`: <https://docs.npmjs.com/files/package.json#local-paths>

### npm global modules 확인

`npm ls -g --depth=0`

### npm 의존성 추가

운영이 아닌 개발 단계에서만 필요한 의존성 모듈들은 devDependencies로 설치

peerDependencies는 현재 모듈과 다른 모듈간의 호환성 표시 (npm install시 node_modules에 설치되지 않음)

peer dependencies 확인: `npm info "패키지명@버전" peerDependencies`

1. `npm install`: ./node_modules 패키지 설치
2. `npm install --save / --save-dev(-D)` ./node_modules에 패키지 설치 + ./package.json 업데이트

### npm 캐시삭제

`npm cache clean --force`

### npmrc

> npm gets its config settings from the command line, environment variables, and npmrc files.

- <https://docs.npmjs.com/cli/v9/configuring-npm/npmrc>
- <https://docs.npmjs.com/cli/v9/using-npm/config>

```env
fund=false
audit=false
```

## pnpm

### pre/post scripts

- [Don't run the pre/post scripts](https://github.com/pnpm/pnpm/issues/2891)
- <https://pnpm.io/cli/run#enable-pre-post-scripts>

## yarn

### yarn 설치

<https://yarnpkg.com/getting-started/install>

```sh
npm install -g yarn
```

> Using a single package manager across your system has always been a problem.
> To be stable, installs need to be run with the same package manager version across environments,
> otherwise there's a risk we introduce accidental breaking changes between versions - after all,
> that's why the concept of lockfile was introduced in the first place!
> And with Yarn being in a sense your very first project dependency,
> it should make sense to "lock it" as well.

### yarn berry (`>=v2`) 사용

프로젝트 yarn 버전 변경

- 해당프로젝트에 `.yarn/`, `.yarnrc.yml`, `.pnp.cjs` 등이 생성됨
- yarn 커맨드가 classic 대신 berry로 작동함

```sh
yarn set version berry
```

에디터 SDK 설정

yarn berry는 `node_modules` 대신 압축된 의존성이 [PnP](https://yarnpkg.com/features/pnp)로 작동하므로 별도의 에디터 설정이 필요

<https://yarnpkg.com/getting-started/editor-sdks>

```sh
# VSCode
yarn dlx @yarnpkg/sdks vscode
```

파일 구조

- `.pnp.cjs`: node가 패키지를 불러오는데 사용할 의존성 트리가 포함됨 (ZeroInstall 관련사항)

- `.yarn`

  - `cache`: 압축된 의존성이 위치함 (GlobalCache 및 ZeroInstall 관련사항)
  - `patches`: `yarn patch-commit` 명령어 결과물인 패치
  - `plugins`, `releases`: 저장소에서 사용하는 yarn 관련 파일들 (`yarn-berry.cjs` ...)
  - `sdks`: 에디터 통합등을 위한 sdk 위치함
  - `unplugged`
    - 의존성 압축을 풀어서 관리함
    - native 파일이 포함된 경우 자동으로 unplug 실행함 <https://github.com/yarnpkg/berry/pull/853>
  - `versions`: 버전 플러그인이 릴리즈 정의를 위해 사용하는 기록

- `yarn.lock`: lockfile

- `.yarnrc.yml`: yarn 설정파일

Zero Install

> `node_modules` 디렉토리와 용량 및 파일 숫자면에서 엄청난 변화가 있으므로,
> 버전관리에 포함시켜 CI 및 배포등의 과정에서 의존성 다운로드 및 설치를 생략할 수 있다

- <https://yarnpkg.com/getting-started/qa#which-files-should-be-gitignored>
- <https://yarnpkg.com/features/zero-installs/>

### yarn 기본 명령어

- `yarn` == `npm install` packages.json에 명시된 의존성 패키지를 다운로드/설치 한다.
- `yarn (global) add/bin/list/remove/upgrade [--prefix]`

### yarn add prefix

- `--dev`(-D) : devDependencies (개발용)
- `--peer`(-P) : peerDependencies (호환성이 있음을 명시: 의존성이 있는것은 아님)
- `--optional`(-O) : optionalDependencies에
- `--exact`(-E) : 명시한 버전과 정확한 경우에만 설치
- `--tilde`(-T) : 명시한 버전과 같은 minor 버전의 최신버전 설치 (버전 세 번째 자리)

### yarn workspace

[[monorepo]] 페이지 참조

### yarn 캐시삭제

```sh
yarn cache clean
```

## package publishing

TS 기준

<https://docs.npmjs.com/files/package.json>

`package.json`

```json
{
  "script": {
    // npm install 되고 난 후 실행됨
    "prepare": "npm run build",
    // npm publish 직전 실행됨
    "prepublishOnly": "npm run lint"
  },
  "main": "lib/index.js",
  "types": "lib/index.d.ts",
  // publishing후 module에 포함할 파일 경로
  "files": ["lib/**/*"],
  // cli를 포함하고 있다면
  "bin": {
    "my-package": "./cli.js"
  },
  "homepage": "https://pravusid.kr",
  "repository": {
    "type": "git",
    "url": "https://github.com/pravusid/my-package.git"
  },
  "bugs": {
    "url": "https://github.com/pravusid/my-package/issues"
  }
}
```

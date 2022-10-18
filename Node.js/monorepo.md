# Mono Repository (monorepo)

Mono Repository(이하 monorepo)는 관련 패키지들을 단일 저장소에서 함께 관리하는 방식이다.

## refs

- overview
  - <https://d2.naver.com/helloworld/0923884>
  - <https://d2.naver.com/helloworld/7553804>
- Nx
  - <https://nx.dev/getting-started/intro>
- TuboRepo
  - <https://turborepo.org/docs>
  - <https://engineering.linecorp.com/ko/blog/monorepo-with-turborepo>
- Rush
  - <https://rushjs.io/pages/intro/get_started/>
  - <https://medium.com/mildang/rush로-프론트엔드-모노레포-도입기-5da0c5bc9b30>
  - <https://dev.to/kkazala/series/17133>
- Lerna
  - <https://medium.com/jung-han/lerna-로-모노레포-해보러나-34c8e008106a>

## monorepo in nodejs

- Node.js 환경의 monorepo 툴은 여러가지가 있지만 주로 node package manager(npm, yarn, pnpm), lerna 등을 사용한다
- MS의 [Rush by TypeScript](https://rushjs.io/), Google의 [Bazel by Java](https://bazel.build/) 등을 사용할 수도 있다

mororepo 툴이 수행하는 주요기능은 다음과 같다

- 버전관리 (version)
- 배포관리 (publish)
- 스크립트 실행 (run)
- 의존성관리 (bootstrap / hoist, symlink)

> monorepo 내의 node_modules 중복이 많아질 수록 lerna의 성능이 좋지 않다
>
> -- <https://doppelmutzi.github.io/monorepo-lerna-yarn-workspaces/>

## pnpm workspace

<https://pnpm.io/workspaces>

## lerna

<https://github.com/lerna/lerna>

주요 명령어는 다음과 같다

### [init](https://github.com/lerna/lerna/tree/main/commands/init#readme)

> Create a new Lerna repo or upgrade an existing repo to the current version of Lerna

### [bootstrap](https://github.com/lerna/lerna/tree/main/commands/bootstrap#readme)

> Bootstrap the packages in the current Lerna repo. Installs all of their dependencies and links any cross-dependencies.

### [version](https://github.com/lerna/lerna/tree/main/commands/version#readme)

> Bump version of packages changed since the last release.
> Identifies packages that have been updated since the previous tagged release

### [publish](https://github.com/lerna/lerna/tree/main/commands/publish#readme)

> Publish packages in the current project

### [run](https://github.com/lerna/lerna/tree/main/commands/run#readme)

> Run an npm script in each package that contains that script

## monorepo 적용 (w/ lerna, github package registry)

<https://viewsource.io/publishing-and-installing-private-github-packages-using-yarn-and-lerna/>

```bash
yarn init
lerna init --independent
```

프로젝트 루트에 `package.json` 파일과 `lerna.json` 설정파일이 생성된다

### `package.json`

#### root

```json
{
  "name": "monorepo",
  "private": true,
  "workspaces": ["packages/core", "packages/*"],
  "scripts": {
    "build": "lerna run build --stream",
    "clean": "lerna run clean --parallel",
    "deps": "yarn exec --workspaces -- ncu",
    "deps:update": "yarn exec --workspaces -- ncu --target minor -u",
    "deps:clean": "find . -name 'node_modules' -type d -prune -print -exec rm -rf '{}' +",
    "version:all": "lerna version --no-changelog",
    "publish:all": "yarn run clean && yarn run build && lerna publish from-package"
  }
}
```

workspace 목록을 별도 표기하여 빌드 실행 우선순위를 지정할 수 있음
(이 경우 lerna 커맨드 실행시 패키지 중복오류가 발생할 수 있으므로 전체 경로(`*`)만 사용하고 스크립트 실행등은 lerna run 사용)

#### in packages

<https://docs.npmjs.com/cli/v6/configuring-npm/package-json#repository>

패키지 이름은 `@scope` + 워크스페이스 prefix 경로(`packages/*`)를 제외한 패키지로 정의하고

각 패키지의 `package.json` 파일에 다음 내용을 추가한다

```json
{
  ...
  "name": "@scope/package-name-without-workspaces-prefix",
    "repository": {
    "type": "git",
    "url": "ssh://git@github.com/idpravus/monorepo.git",
    "directory": "packages/package-name-without-workspaces-prefix"
  },
  "publishConfig": {
    "registry": "https://npm.pkg.github.com/"
  }
  ...
}
```

### `lerna.json`

다음 내용을 추가한다

```json
{
  ...
  "npmClient": "yarn",
  "useWorkspaces": true,
  "command": {
    "version": {
      "message": "release",
      "ignoreChanges": ["**/*.spec.ts", "**/*.md"]
    },
    "publish": {
      "registry": "https://npm.pkg.github.com",
      "allowBranch": "main"
    }
  }
}
```

### monorepo 내부의 다른 모듈 사용

다른 모듈을 사용하기 위해서는 `dependencies` 선언한 뒤 `import` 해서 사용한다

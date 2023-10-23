# Mono Repository (monorepo)

Mono Repository(이하 monorepo)는 관련 패키지들을 단일 저장소에서 함께 관리하는 방법론이다

## articles, examples

- overview

  - <https://monorepo.tools/>
  - <https://d2.naver.com/helloworld/0923884>
  - <https://d2.naver.com/helloworld/7553804>

- Nx

  - <https://github.com/ddd-by-examples/library-nestjs>

- Turborepo

  - <https://engineering.linecorp.com/ko/blog/monorepo-with-turborepo>

- Rush

  - <https://medium.com/mildang/rush로-프론트엔드-모노레포-도입기-5da0c5bc9b30>
  - <https://dev.to/kkazala/series/17133>

- Lerna

  - <https://medium.com/jung-han/lerna-로-모노레포-해보러나-34c8e008106a>

- Sharing TypeScript with Nx and Turborepo

  - [Part 1: An Introduction to Monorepos](https://javascript.plainenglish.io/d8d54b805e46)
  - [Part 2: Creating a Monorepo](https://javascript.plainenglish.io/347bf3194181)
  - [Part 3: Configuring a Monorepo](https://javascript.plainenglish.io/2e4608701964)
  - [Part 4: Managing a Monorepo](https://javascript.plainenglish.io/a48adc99823e)

- Pnpm and Nx monorepo

  - <https://www.javierbrea.com/blog/pnpm-nx-monorepo-01/>
  - <https://www.javierbrea.com/blog/pnpm-nx-monorepo-02/>
  - <https://www.javierbrea.com/blog/pnpm-nx-monorepo-03/>

## vids

- [우리는 하나다! 모노레포 with pnpm](https://www.youtube.com/watch?v=Bycg5w5qXfE)
- [일백개 패키지 모노레포 우아하게 운영하기](https://www.youtube.com/watch?v=Ix9gxqKOatY)
- [모노레포 마이크로 아키텍처를 지향하며](https://www.youtube.com/watch?v=CsbBuE_MF2U)
- [모노레포 희망편 (Feat.Polylith)](https://www.youtube.com/watch?v=CCo7T3m6LLM)

## monorepo features

mororepo 툴이 수행하는 주요기능은 다음과 같다

- 배포 관리 (publish)
- 스크립트 실행 (run)
- 의존성 관리 (root 레벨에서 monorepo 전체 의존성을 관리하는 경우)
- ~~버전 관리 (version)~~ (툴별로 지원여부가 다르다)

## Version Managers

- <https://github.com/changesets/changesets>
- <https://nx.dev/recipes/adopting-nx/lerna-and-nx#version-management-&-publishing>
- <https://turbo.build/repo/docs/handbook/publishing-packages/versioning-and-publishing>

## Package Managers

주요 패키지매니저(npm, pnpm, yarn)들은 워크스페이스 기능을 지원한다.
워크스페이스 옵션을 지정하면 지정한 경로의 패키지를 우선 참조하여, 패키지를 설치했을 때 `node_moudles`에 패키지 링크를 생성한다.

각 패키지에서 다른 패키지를 참조할 때 `"*"`(모든버전)으로 지정한다. (pnpm은 `"workspace:*"`)

```json
{
  "dependencies": {
    "other-package": "*"
  }
}
```

패키지매니저에서 별도 옵션을 사용하지 않는다면 루트 디렉토리의 `node_modules`에 전체 의존성이 관리되고
하위 패키지의 공통 의존성, 프로젝트 패키지간 참조등은 심볼릭 링크로 처리된다.

### pnpm workspace

[[npm#pnpm]]

<https://pnpm.io/workspaces>

[`pnpm-workspace.yaml`](https://pnpm.io/pnpm-workspace_yaml)

#### pnpm workspace deployment

pnpm에서는 다음 옵션을 고려해볼 수 있다

- <https://pnpm.io/cli/deploy>
- <https://pnpm.io/npmrc#shared-workspace-lockfile>

관련 문서

- [Bundling up project for deployment](https://github.com/pnpm/pnpm/issues/2198)
- [A deploy command](https://github.com/pnpm/pnpm/issues/4378)
- <https://github.com/vercel/next.js/issues/45258>

### npm workspace

[[npm#npm]]

<https://docs.npmjs.com/cli/v10/using-npm/workspaces>

[`package-json`](https://docs.npmjs.com/cli/v10/configuring-npm/package-json#workspaces)

## monorepo deployment

- 패키지 게시

  - lockfile을 사용하지 않으므로 빌드 결과물과 `package.json`을 배포하면 된다
  - 참고: [[npm#npm package publishing & lockfile]]

- 웹 프론트엔드 프로젝트

  - 번들러를 사용하므로 빌드 결과물을 배포하면 된다

- 웹 백엔드 프로젝트

  - 번들러를 사용했을 때 파일기반 기능(auto loader, file scan)이 정상작동하지 않을 수 있어 주로 `node_modules`를 함께 배포하게 된다 (docker 배포 역시 마찬가지)
  - 배포하려는 패키지의 의존성만 선택해서 포함하기 어려울 수 있다 (패키지매니저가 기능을 제공해야 한다)
  - 참고: [pnpm workspace deployment](#pnpm-workspace-deployment), [symlinks-in-node_modules](#symlinks-in-node_modules)

## Nx

<https://nx.dev/getting-started/intro>

### Nx Quickstart

- <https://turbo.build/repo/docs/getting-started/create-new#quickstart>
- <https://github.com/nrwl/nx-examples>
- <https://github.com/nrwl/nx-recipes>

```sh
npx create-nx-workspace --pm pnpm
```

### Nx Configuration

[`nx.json`](https://nx.dev/reference/nx-json)

## Turborepo

<https://turbo.build/repo/docs>

### Turborepo Quckstart

- <https://turbo.build/repo/docs/getting-started/create-new#quickstart>
- <https://turbo.build/repo/docs/getting-started/from-example>

```sh
npx create-turbo@latest
```

### Turborepo Configuration

[`turbo.json`](https://turbo.build/repo/docs/reference/configuration)

## lerna

<https://github.com/lerna/lerna>

> monorepo 내의 node_modules 중복이 많아질 수록 lerna의 성능이 좋지 않다
>
> -- <https://doppelmutzi.github.io/monorepo-lerna-yarn-workspaces/>

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

### monorepo 적용 (w/ lerna, github package registry)

<https://viewsource.io/publishing-and-installing-private-github-packages-using-yarn-and-lerna/>

```bash
yarn init
lerna init --independent
```

프로젝트 루트에 `package.json` 파일과 `lerna.json` 설정파일이 생성된다

#### `package.json` - root

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

#### `package.json` - packages

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

#### `lerna.json`

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

## Rush

<https://rushjs.io/pages/intro/get_started/>

## monorepo tools

### eslint

<https://typescript-eslint.io/linting/typed-linting/monorepos>

### jest

<https://jestjs.io/docs/next/configuration#projects-arraystring--projectconfig>

### vitest

<https://vitest.dev/guide/workspace.html>

## monorepo troubleshooting

### symlinks in node_modules

> `node_modules` 를 복사할 때는 symlink도 포함해야 함

- `tar` 는 별도 옵션 없어도 포함함
- `zip` 은 `-y, --symlinks` 옵션 사용

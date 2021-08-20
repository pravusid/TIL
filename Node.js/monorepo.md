# Mono Repository

Mono Repository(이하 monorepo)는 관련 패키지들을 단일 저장소에서 함께 관리하는 방식이다.

## monorepo in NodeJS

- NodeJS 환경의 monorepo 툴은 여러가지가 있지만 주로 lerna, yarn workspace를 사용한다
- MS의 [Rush by TypeScript](https://rushjs.io/), Google의 [Bazel by Java](https://bazel.build/) 등을 사용할 수도 있다

> npm도 v7 이후 workspace 지원함

둘은 고유 기능과 중복 기능이 있는데 중복기능은 상황에 맞는 패키지를 사용하면 된다.

- lerna

  - 버전관리 (version)
  - 배포관리 (publish)
  - 스크립트 실행 (run)
  - 의존성관리 (bootstrap / hoist, symlink)

- yarn workspace

  - 스크립트 실행 (run)
  - 의존성관리 (add / hoist, symlink)

monorepo 내의 node_modules 중복이 많아질 수록 yarn workspace의 성능이 좋다
(참고: <https://doppelmutzi.github.io/monorepo-lerna-yarn-workspaces/>)

> 의존성관리를 제외한 버전관리, 배포관리, 스크립트 실행은 lerna를 사용하는 것이 더 좋음 (yarn workspaces는 의존관계에 따른 순서 처리가 없음)

## Lerna

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

## Yarn workspace

<https://classic.yarnpkg.com/en/docs/workspaces/>

<https://classic.yarnpkg.com/en/docs/cli/workspace>

<https://classic.yarnpkg.com/en/docs/cli/workspaces>

## monorepo 적용 (w/ github package registry)

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
  "workspaces": ["packages/core", "packages/*"]
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

### TypeScript

타입스크립트 컴파일러 설정과 type export 경로 처리에 주의해야 한다

### 사용

다른 모듈을 사용하기 위해서는 `Peer Dependencies` 또는 `Dependencies` 선언한 뒤 `import` 해서 사용한다

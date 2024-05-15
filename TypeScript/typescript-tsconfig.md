# tsconfig.json

디렉토리에 `tsconfig.json` 파일이 있으면 TypeScript 프로젝트의 루트로 인식된다.

프로젝트는 다음 중 하나의 방식으로 컴파일 된다

- 입력 파일없이 tsc를 호출하면 컴파일러는 현재 디렉토리에서 시작하여 상위 디렉토리 방향의 체인으로 `tsconfig.json` 파일을 검색한다
- 입력 파일없이 tsc를 호출하고 `--project` 커맨드 라인 옵션과 함께 `tsconfig.json` 파일을 포함하는 디렉토리나 유효한 `.json` 파일 경로를 사용한다.

커맨드라인에 파일 입력이 지정되면 `tsconfig.json` 파일이 무시된다

## 참조

- <https://www.typescriptlang.org/tsconfig>
- <https://www.typescriptlang.org/docs/handbook/compiler-options.html>
- <https://www.typescriptlang.org/docs/handbook/project-references.html>
- <https://www.typescriptlang.org/docs/handbook/modules/guides/choosing-compiler-options.html>
- <https://www.totaltypescript.com/tsconfig-cheat-sheet>

## 예시

```jsonc
{
  "compilerOptions": {
    // ...
  },
  "files": ["foo.ts", "bar.ts"],
  "include": ["src/**/*"],
  "exclude": ["node_modules", "**/*.spec.ts"]
}
```

## 옵션 자세히 보기

`"compilerOptions"` 속성을 생략할 수 있으며 이 경우 컴파일러 기본값이 사용된다.
마찬가지로 `tsconfig.json` 파일을 완전히 비워두면 모든 설정값이 기본값으로 적용된다.

명령행에 지정된 컴파일러 옵션이 `tsconfig.json` 파일에서 지정된 설정을 대체한다.

### target

- <https://www.typescriptlang.org/tsconfig/#target>
- <https://github.com/tsconfig/bases#centralized-recommendations-for-tsconfig-bases>

컴파일 할 자바스크립트 버전을 선택한다. target 설정에 따라 lib 설정 값도 변경된다 (기본 값을 사용하고 있다면)

### lib

<https://www.typescriptlang.org/tsconfig/#lib>

- 타입스크립트에서 사용할 built-in JavaScript API 타입선언을 지정한다
- 필요에 따라 선언의 일부분만 사용할 경우가 있다 (DOM 타입선언을 사용하지 않는 경우, 일부 polyfill만 포함한 경우...)

> target 설정에 따른 lib 기본 값은 <https://github.com/microsoft/TypeScript/blob/main/src/lib/libs.json> 코드에서
> target 버전과 일치하는 Default libraries 설정 값이다. (`es20??.full`)
>
> 해당 타입선언을 열어보면 기본 값으로 설정되는 lib들이 선언되어 있다. (`<reference lib="es20??" /><reference lib="dom" />` ...)
>
> --<https://stackoverflow.com/questions/63943629/what-is-the-typescript-compilers-default-lib-value>

### 경로 관련 설정

#### baseUrl

- <https://www.typescriptlang.org/tsconfig#baseUrl>
- <https://www.typescriptlang.org/tsconfig#paths>

모듈 절대경로(non-relative) 설정(`paths`)에서 기준 경로를 설정할 때 사용하는 옵션

<https://www.typescriptlang.org/docs/handbook/module-resolution.html#base-url>

> 기본값: 없음

#### rootDir

<https://www.typescriptlang.org/tsconfig#rootDir>

컴파일러는 컴파일 후 rootDir 내부와 동일한 구조를 출력 디렉토리에 유지한 채 출력한다.

> 기본값: 선언파일을 제외한 입력파일에서 가장 긴 공통경로

rootDir은 컴파일러 입력 포함/제외에 영향을 미치지 않는다.(출력물의 경로 구조에만 영향)

그러나 outDir 외부에 컴파일 결과가 출력되지 않으므로, 결과적으로는 rootDir 외부 경로의 파일이 입력에 포함되지도 않는다.
(컴파일 중 rootDir 외부 경로의 모듈 참조가 있다면 오류가 발생할 것이다)

rootDir 옵션이 지정된 경우 `<configname>.tsbuildinfo` 파일(incremental build information)이 프로젝트 루트에 생성됨.
이 경우 직접 파일 경로를 컴파일러 옵션에서 변경할 수 있음.

```json
{
  "compilerOptions": {
    "tsBuildInfoFile": "./dist/tsconfig.tsbuildinfo",
    "rootDir": "./src",
    "outDir": "./dist"
  }
}
```

This is correct since the output is relative to rootDir when specified.
Since configFile is in parent directory relative to rootDir, the tsbuildinfo file goes in parent folder to outDir.

From [d53efdf](https://github.com/microsoft/TypeScript/commit/d53efdf38058e37d52e794b6650689294e69b185)

tsbuild info is generated at:

- If composite or incremental then only the .tsbuildinfo will be generated
- if --out or --outFile the file is outputFile.tsbuildinfo
- if rootDir and outDir then outdir/relativePathOfConfigFromRootDir/configname.tsbuildinfo
- if just outDir then outDir/configname.tsbuild
- otherwise config.tsbuildinfo next to configFile

<https://github.com/microsoft/TypeScript/issues/30925>

#### rootDirs

<https://www.typescriptlang.org/tsconfig#rootDirs>

프로젝트 루트는 여러 경로인데 출력은 병합해야 하는 경우, 컴파일러에서 상대 모듈 가져오기를 처리하기 위한 옵션

<https://www.typescriptlang.org/docs/handbook/module-resolution.html#virtual-directories-with-rootdirs>

#### outDir

<https://www.typescriptlang.org/tsconfig#outDir>

컴파일러 출력경로

> 기본값: 없음 == 소스 파일과 동일 경로에 컴파일된 파일이 생성됨

### `"files"`, `"include"`, `"exclude"`

`"files"` 속성은 상대경로 또는 절대경로 파일 목록을 가져온다.

`"include"`, `"exclude"` 속성은 glob-like 파일패턴 목록을 사용한다. 지원되는 glob 와일드 카드는 다음과 같다

- `*`: 0개 혹은 이상의 문자(디렉토리 구분자 제외)
- `?`: 임의의 1개 문자(디렉토리 구분자 제외)
- `**/`: 재귀적으로 하위 디렉토리 매칭

glob패턴 세그먼트에 `*` 또는 `.*`만 포함된 경우 지원하는 확장자를 가진 파일만 포함한다.

> e.g. `.ts`, `.tsx`, `.d.ts`, 만약 `allowJs` 옵션을 사용한다면 `.js`, `.jsx` 포함

`"files"`와 `"include"` 모두 지정되지 않은 경우 컴파일러는 기본적으로 `"exclude"` 속성에 정의된 파일은 제외하고
포함된 디렉토리와 서브디렉토리의 모든 TypeScript 파일을 포함한다.

`"files"` 또는 `"include"` 속성이 지정되면 컴파일러는 두 속성에 포함된 파일의 합집합을 포함한다.

`"files"` 또는 `"include"` 속성을 통해 포함된 파일이 참조하는 모든 파일도 포함된다.
마찬가지로 파일 `A.ts`에서 다른파일 `B.ts`를 참조하는 경우 파일 `A.ts`가 `"exclude"` 목록에 지정되어 있지 않으면 `B.ts`를 제외할 수 없다.

`"exclude"` 특성이 지정되지 않은 경우 `"outDir"` 컴파일러 옵션을 사용하여 지정한 디렉토리의 파일 및
`node_modules`, `bower_components`, `jspm_packages`를 기본적으로 제외한다.

`"include"`를 사용하여 포함한 파일은 `"exclude"` 속성을 사용하여 필터링 할 수 있다.

그러나 `"files"` 속성을 사용하여 명시적으로 포함한 파일은 `"exclude"` 속성에 관계없이 항상 포함된다.

컴파일러 입력에 `index.ts`가 포함되어 있으면 `index.d.ts`와 `index.js`는 제외된다.
따라서 확장자만 다르고 이름이 같은 파일은 같이 두지 않는 것이 좋다.

### `@types`, `typeRoots`, `types`

기본적으로 모든 가시적인 `@types` 패키지가 컴파일에 포함된다.
enclosing `node_modules/@types` 디렉토리는 가시적인 것으로 간주된다.

> enclosing 디렉토리를 풀어보면 `./node_modules/@types/`, `../node_modules/@types/`, `../../node_modules/@types/` 같은 규칙으로 순차 적용된다

`typeRoots` 속성이 지정되면 `typeRoots` 아래의 패키지만 포함된다

```json
{
  "compilerOptions": {
    "typeRoots": ["./typings"]
  }
}
```

이 설정에서는 `./typings` 아래의 모든 패키지가 포함되고, `./node_modules/@types` 아래의 패키지는 포함되지 않는다.

`types` 속성이 지정되면 나열된 패키지만 포함된다

```json
{
  "compilerOptions": {
    "types": ["node", "lodash", "express"]
  }
}
```

이 설정에서는 `./node_modules/@types/node`, `./node_modules/@types/lodash`, `./node_modules/@types/express`만 포함된다.
`./node_modules/@types/*` 경로의 다른 패키지는 포함되지 않는다.

types 패키지는 `index.d.ts` 파일이 있는 폴더 혹은 `package.json` 에서 `types` 필드에 명시된 폴더이다.

자동 포함은 모듈로 선언된 파일이 아니라 global 선언 파일을 사용할 때만 중요하다.
예를 들어, `import "foo"`를 사용하면 TypeScript는 여전히 `node_modules` 및 `node_modules/@types` 폴더에서 `foo` 패키지를 찾는다.

## `extends` 사용으로 설정 상속

`tsconfig.json` 파일은 `extends` 속성으로 다른 파일에서 설정을 상속할 수 있다.

`extends`는 `tsoncig.json` 파일의 최상위 속성(`compilerOptions`, `files`...)이다.

base file의 설정이 우선 로드된 다음 상속되는 파일의 설정으로 대체된다.

상속되는 파일에서 `files`, `include`, `exclude` 옵션은 base file 설정을 덮어쓴다.

설정파일의 모든 상대경로는 해당 설정파일의 위치에 따라 처리된다

`configs/base.json`

```json
{
  "compilerOptions": {
    "noImplicitAny": true,
    "strictNullChecks": true
  }
}
```

`tsconfig.json`

```json
{
  "extends": "./configs/base",
  "files": ["main.ts", "supplemental.ts"]
}
```

`tsconfig.nostrictnull.json`

```json
{
  "extends": "./tsconfig",
  "compilerOptions": {
    "strictNullChecks": false
  }
}
```

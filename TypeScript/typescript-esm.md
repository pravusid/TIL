# TypeScript ESM support

## Refs

- <https://github.com/microsoft/TypeScript/issues/46452>
- <https://devblogs.microsoft.com/typescript/announcing-typescript-4-5-beta/#esm-nodejs>
- <https://devblogs.microsoft.com/typescript/announcing-typescript-4-5-rc/#esm-nodejs>
- <https://devblogs.microsoft.com/typescript/announcing-typescript-4-7/#esm-nodejs>
- <https://www.typescriptlang.org/docs/handbook/esm-node.html>
- <https://dev.to/logto/migrate-a-60k-loc-typescript-nodejs-repo-to-esm-and-testing-become-4x-faster-12-5f82>
- <https://dev.to/logto/migrate-a-60k-loc-typescript-nodejs-repo-to-esm-and-testing-become-4x-faster-22-4a4k>

## 전환방법

### `tsconfig.json`

> TypeScript 4.7 adds this functionality with two new module settings: node16 and nodenext.

```json
{
  "compilerOptions": {
    "module": "nodenext"
  }
}
```

### `package.json`

```json
{
  "name": "my-package",
  "type": "module"
}
```

이 설정은 `.js`, `.d.ts` 파일을 ESM으로 인식하도록 설정하고 다음 사항들을 적용한다

- import/export statements 사용 가능
- Top-level await 사용 가능
- 불러오기의 상대경로는 확장자를 명시해야 함 (`import "./foo"` 대신 `import "./foo.js"`)
- `node_modules` 경로와는 불러오기가 다르게 작동할 수 있음 (commonjs는 [[typescript-handbook-module-system]] 참고)
- `require`, `module`, `__dirname` 같은 글로벌 변수를 직접 사용할 수 없음
- CommonJS modules은 정해진 규칙에 따라 불러올 수 있음

다음과 같은 명시적인 확장자를 사용할 수 있음

- CommonJS: `.cjs`, `.cts`, `.d.cts`
- ESM: `.mjs`, `.mts`, `.d.mts`

## `package.json` Exports, Imports, and Self-Referencing

- <https://github.com/microsoft/TypeScript/issues/33079>
- <https://www.typescriptlang.org/docs/handbook/esm-node.html#packagejson-exports-imports-and-self-referencing>
- <https://nodejs.org/api/packages.html>
- <https://antfu.me/posts/publish-esm-and-cjs>
- <https://toss.tech/article/commonjs-esm-exports-field>

## TypeScript’s Migration to Modules

타입스크립트 팀에서 코드베이스를 ESM으로 전환하면서 남긴 블로그 포스트이다

> In TypeScript 5.0, we restructured our entire codebase to use ECMAScript modules, and switched to a newer emit target
>
> -- <https://devblogs.microsoft.com/typescript/typescripts-migration-to-modules/>

## import 구문 확장자 관련

- <https://github.com/microsoft/TypeScript/issues/49083>
- <https://www.reddit.com/r/typescript/comments/uuivss/module_node16_should_support_extension_rewriting/>
- <https://github.com/nodejs/node/issues/46006>
- <https://github.com/microsoft/TypeScript/issues/37582>
- <https://devblogs.microsoft.com/typescript/announcing-typescript-5-0/#resolution-customization-flags>

## import default 관련

> ESM에서 CommonJS 모듈을 불러올 때 named export, default export 선언이 섞여있고 별도의 처리가 없다면 오류 발생할 수 있음

- <https://nodejs.org/docs/latest/api/esm.html#interoperability-with-commonjs>
- [TypeScript module "node16" does not work with CommonJS dependencies](https://github.com/microsoft/TypeScript/issues/49271)
- [This expression is not callable for ESM consuming CJS with default export](https://github.com/microsoft/TypeScript/issues/52086)

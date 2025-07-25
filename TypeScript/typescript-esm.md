# TypeScript ESM support

## Refs

- <https://www.typescriptlang.org/docs/handbook/modules/introduction.html>
- <https://www.typescriptlang.org/docs/handbook/esm-node.html>
- [Concerns with TypeScript 4.5's Node 12+ ESM Support](https://github.com/microsoft/TypeScript/issues/46452)
- <https://devblogs.microsoft.com/typescript/announcing-typescript-4-5-beta/#esm-nodejs>
- <https://devblogs.microsoft.com/typescript/announcing-typescript-4-5-rc/#esm-nodejs>
- <https://devblogs.microsoft.com/typescript/announcing-typescript-4-7/#esm-nodejs>
- <https://dev.to/logto/migrate-a-60k-loc-typescript-nodejs-repo-to-esm-and-testing-become-4x-faster-12-5f82>
- <https://dev.to/logto/migrate-a-60k-loc-typescript-nodejs-repo-to-esm-and-testing-become-4x-faster-22-4a4k>
- <https://nodejs.org/docs/latest/api/esm.html>

## 전환방법

### `tsconfig.json`

> TypeScript 4.7 adds this functionality with two new module settings: node16 and nodenext.

```json
{
  "compilerOptions": {
    "module": "nodenext",
    "rewriteRelativeImportExtensions": true
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

### `--module node20`

<https://github.com/microsoft/TypeScript/pull/61805>

|          | `target` | `moduleResolution` | `resolveJsonModule` | import assertions | import attributes | JSON imports        | require(esm) |
| -------- | -------- | ------------------ | ------------------- | ----------------- | ----------------- | ------------------- | ------------ |
| node16   | `es2022` | `node16`           | false               | 🚫                | 🚫                | no restrictions     | 🚫           |
| node18   | `es2022` | `node16`           | false               | ✅                | ✅                | needs `type "json"` | 🚫           |
| node20   | `es2023` | `node16`           | true                | 🚫                | ✅                | needs `type "json"` | ✅           |
| nodenext | `esnext` | `nodenext`         | true                | 🚫                | ✅                | needs `type "json"` | ✅           |

## `package.json` Exports, Imports, and Self-Referencing

- [Support for NodeJS 12.7+ package exports](https://github.com/microsoft/TypeScript/issues/33079)
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

- [`"module": "node16"` should support extension rewriting](https://github.com/microsoft/TypeScript/issues/49083)
- <https://www.reddit.com/r/typescript/comments/uuivss/module_node16_should_support_extension_rewriting/>
- [Adding support for ESM references without a .js extension](https://github.com/nodejs/node/issues/46006)
- [allow voluntary .ts suffix for import paths](https://github.com/microsoft/TypeScript/issues/37582)
- <https://devblogs.microsoft.com/typescript/announcing-typescript-5-0/#--moduleresolution-bundler>
- <https://devblogs.microsoft.com/typescript/announcing-typescript-5-0/#resolution-customization-flags>
- <https://devblogs.microsoft.com/typescript/announcing-typescript-5-2/#decorator-metadata>
- <https://devblogs.microsoft.com/typescript/announcing-typescript-5-7/#path-rewriting-for-relative-paths>
- <https://github.com/microsoft/TypeScript/pull/59767>
- <https://www.perplexity.ai/search/typescript-esm-support-eseo-im-p2J8Ki87TLi6sB2pUKsqEA>

확장자 관련 이슈는 5.7 버전에서 `--rewriteRelativeImportExtensions` 설정을 지원하면서 거의 정리된 것으로 보인다.
해당옵션을 사용하면 [`--allowImportingTsExtensions`](https://www.typescriptlang.org/vo/tsconfig/#allowImportingTsExtensions) 설정도 같이 활성화된다.

## import default 관련

> ESM에서 CommonJS 모듈을 불러올 때 named export, default export 선언이 섞여있고 별도의 처리가 없다면 오류 발생할 수 있음

- [TypeScript module "Node16" does not resolve types of CJS module](https://github.com/microsoft/TypeScript/issues/49160)
- [TypeScript module "node16" does not work with CommonJS dependencies](https://github.com/microsoft/TypeScript/issues/49271)
- [This expression is not callable for ESM consuming CJS with default export](https://github.com/microsoft/TypeScript/issues/52086)
- <https://www.npmjs.com/package/default-import>

## esm cjs interop

- <https://www.typescriptlang.org/docs/handbook/modules/appendices/esm-cjs-interop.html>
- <https://devblogs.microsoft.com/typescript/announcing-typescript-5-8/#support-for-require()-of-ecmascript-modules-in---module-nodenext>
- [[nodejs#nodejs ESM]]

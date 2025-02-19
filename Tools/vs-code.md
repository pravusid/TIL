# Visual Studio Code

- <https://github.com/microsoft/vscode-recipes>
- <https://code.visualstudio.com/docs>
- <https://code.visualstudio.com/docs/getstarted/settings#_default-settings>

## TypeScript

### TypeScript SDK

`.vscode/settings.json`

```json
{
  "typescript.enablePromptUseWorkspaceTsdk": true,
  "typescript.tsdk": "node_modules/typescript/lib"
}
```

> The typescript.tsdk workspace setting only tells VS Code that a workspace version of TypeScript exists.
> To actually start using the workspace version for IntelliSense, you must run the TypeScript:
> Select TypeScript Version command and select the workspace version.

### TS Server WatchOptions

options: <https://www.typescriptlang.org/docs/handbook/configuring-watch.html>

```jsonc
{
  "typescript.tsserver.watchOptions": {
    // ...options
  }
}
```

CodeGen + tsserver.watch

<https://github.com/dotansimha/graphql-code-generator/discussions/8345#discussioncomment-4028928>

```json
{
  "typescript.tsserver.watchOptions": {
    "watchDirectory": "useFsEvents",
    "fallbackPolling": "dynamicPriorityPolling",
    "watchFile": "useFsEventsOnParentDirectory",
    "synchronousWatchDirectory": true
  }
}
```

### Exclude Patterns for Auto-Imports

<https://devblogs.microsoft.com/typescript/announcing-typescript-5-6/#exclude-patterns-for-auto-imports>

> The same settings can be applied for JavaScript through `javascript.preferences.autoImportSpecifierExcludeRegexes` in VSCode.

```jsonc
{
  "typescript.preferences.autoImportSpecifierExcludeRegexes": [
    "^lodash$",
    "^node:",
    "^./lib/internal", // no escaping needed
    "/^.\\/lib\\/internal/", // escaping needed - note the leading and trailing slashes
    "/^.\\/lib\\/internal/i" // escaping needed - we needed slashes to provide the 'i' regex flag
  ]
}
```

## Formatting

`.vscode/settings.json`

```json
{
  "editor.formatOnSave": true,
  "editor.codeActionsOnSave": {
    "source.addMissingImports": "always",
    "source.organizeImports": "always"
  },
  "editor.pasteAs.preferences": ["text.updateImports"],
  "editor.defaultFormatter": "esbenp.prettier-vscode",
  "files.insertFinalNewline": true,
  "files.trimFinalNewlines": true,
  "files.trimTrailingWhitespace": true
}
```

for JavaScript, TypeScript

```json
{
  "javascript.format.enable": false,
  "typescript.format.enable": false,
  "javascript.validate.enable": false
}
```

for vuejs

```json
{
  "[vue]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  }
}
```

- <https://www.typescriptlang.org/docs/handbook/release-notes/typescript-5-0.html#case-insensitive-import-sorting-in-editors>
- <https://www.typescriptlang.org/docs/handbook/release-notes/typescript-4-9.html#remove-unused-imports-and-sort-imports-commands-for-editors>
- <https://www.typescriptlang.org/docs/handbook/release-notes/typescript-4-8.html#exclude-specific-files-from-auto-imports>
- <https://www.typescriptlang.org/docs/handbook/release-notes/typescript-4-7.html#group-aware-organize-imports>

## Linting

[[eslint-prettier#VSCode]] 참고

## Debugging

<https://code.visualstudio.com/docs/editor/debugging>

### Nodemon + TypeScript Debugging

attaching debugger to Nodemon

- <https://nodejs.org/en/learn/getting-started/debugging#command-line-options>
- <https://code.visualstudio.com/docs/nodejs/nodejs-debugging#_attaching-to-nodejs>
- <https://github.com/Microsoft/vscode-recipes/tree/main/nodemon>

npm script: `"debug": "nodemon --watch dist --exec 'NODE_ENV=debug node -r source-map-support/register --inspect=9229' dist/main.js"`

`.vscode/launch.json`

```json
{
  "version": "0.2.0",
  "configurations": [
    {
      "type": "node",
      "request": "attach",
      "name": "Debug: Nodemon, TypeScript",
      "port": 9229,
      "restart": true,
      "sourceMaps": true
    }
  ]
}
```

### TypeScript Debugging

<https://code.visualstudio.com/docs/typescript/typescript-debugging>

`.vscode/launch.json`

```json
{
  "version": "0.2.0",
  "configurations": [
    {
      "type": "node",
      "request": "launch",
      "name": "Debug: TypeScript in Node.js",
      "preLaunchTask": "npm: build",
      "program": "${workspaceFolder}/src/main.ts",
      "cwd": "${workspaceFolder}",
      "outFiles": ["${workspaceFolder}/dist/**/*.js"]
    }
  ]
}
```

`preLaunchTask`

- VSCode `Tasks: Run Task`에서 선택할 수 있는 작업 입력
- TypeScript 프로젝트(root 경로에 `tsconfig.json` 존재)는 tsc build/watch 작업이 기본 출력됨
- Node 프로젝트(root 경로에 `package.json` 존재)는 npm script 작업이 기본 출력됨
- 언어팩에 따라서 실행이 되지 않는 경우가 있으므로(build -> 빌드) `npm: build` 실행을 추천

> `main.ts` 경로와 `dist/` 경로는 설정에 따라변경

### node.js runtime 오류

[`nvm` 사용시 `PATH`에서 node runtime을 찾지 못하는 문제](https://code.visualstudio.com/docs/nodejs/nodejs-debugging#_multi-version-support)

`.vscode/launch.json`

```json
{
  "version": "0.2.0",
  "configurations": [
    {
      // ...
      "runtimeVersion": "<NODE_VESION_IN_NVM>"
      // ...
    }
  ]
}
```

> runtimeVersion을 지원하지 않는 version manager 사용 또는 위 옵션을 사용하지 않으려면
> `~/.profile`, `~/.zprofile`, `~/.zshenv` 같은 환경변수 설정에서 기본 node/bin PATH를 지정한다.

## [[monorepo]]

### Go to definition goes to `.d.ts`

[TypeScript Project References](https://www.typescriptlang.org/ko/docs/handbook/project-references.html#declarationmaps) 구성할 때 설정필요

> Can you please try adding `"declarationMap": true` to the **compilerOptions in your tsconfig**
> This will generated a map so that go to definition can just back to the original ts source instead of the generated d.ts
>
> -- <https://github.com/microsoft/vscode/issues/73201>
> -- <https://www.typescriptlang.org/tsconfig#declarationMap>

### importModuleSpecifier

[[monorepo#monorepo with VSCode]]

## HTML, JSX (Auto rename, Auto closing)

### Auto rename

> <https://www.roboleary.net/vscode/2023/05/08/auto-rename-tags-react-vue-svelte.html>

```json
{
  "editor.linkedEditing": true
}
```

- `javascript.preferences.renameMatchingJsxTags`: 기본값 true
- `typescript.preferences.renameMatchingJsxTags`: 기본값 true

### Auto closing

- `html.autoClosingTags` 기본값 true (HTML)
- `javascript.autoClosingTags` 기본값 true (JSX)
- `typescript.autoClosingTags` 기본값 true (JSX)

## Custom labels for open editors

> v1.88

- <https://github.com/microsoft/vscode/issues/208388>
- <https://www.reddit.com/r/nextjs/comments/1bzd0h7/vs_code_new_feature_implementation_for_next_js/>
- <https://gist.github.com/hAbuMustafa/88288a7fc2141c2a919a492ff3bf84cb>

### Custom labels example: Next.js

```json
{
  "workbench.editor.customLabels.patterns": {
    "**/components/**/index.{ts,tsx}": "#/${dirname}.${extname}",
    "**/components/**/*.{ts,tsx}": "#/${dirname}/${filename}.${extname}",
    "**/app/**/*.{ts,tsx}": "/${dirname}/${filename}.${extname}",
    "**/pages/**/index.tsx": "/${dirname}.${extname}",
    "**/pages/**/*.tsx": "/${dirname}/${filename}.${extname}",
    "**/pages/api/**/index.ts": "@api/${dirname}.${extname}",
    "**/pages/api/**/*.ts": "@api/${dirname}/${filename}.${extname}"
  }
}
```

## 참고사항

### Why are there duplicate paths in the terminal's $PATH environment variable and/or why are they reversed on macOS?

<https://code.visualstudio.com/docs/terminal/profiles#_why-are-there-duplicate-paths-in-the-terminals-path-environment-variable-andor-why-are-they-reversed-on-macos>

> [[env-variables#terminal emulator: macOS|macOS 터미널]]은 login shell을 실행한다

## Github Copilot

deny list

```json
{
  "github.copilot.enable": {
    "*": true,
    "dotenv": false,
    "properties": false
  }
}
```

allow list

```json
{
  "github.copilot.enable": {
    "*": false,
    "javascript": true,
    "javascriptreact": true,
    "typescript": true,
    "typescriptreact": true
  }
}
```

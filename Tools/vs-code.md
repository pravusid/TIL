# Visual Studio Code

<https://github.com/microsoft/vscode-recipes>

<https://code.visualstudio.com/docs>

## TypeScript SDK

`.vscode/settings.json`

```json
{
  "typescript.enablePromptUseWorkspaceTsdk": true,
  "typescript.tsdk": "node_modules/typescript/lib"
}
```

> The typescript.tsdk workspace setting only tells VS Code that a workspace version of TypeScript exists.
> To actually start using the workspace version for IntelliSense, you must run the TypeScript: Select TypeScript Version command and select the workspace version.

## Formatting

`.vscode/settings.json`

```json
{
  "editor.formatOnSave": true,
  "editor.codeActionsOnSave": ["source.addMissingImports", "source.organizeImports"],
  "files.insertFinalNewline": true,
  "files.trimFinalNewlines": true,
  "[javascript]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "[typescript]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "javascript.format.enable": false,
  "typescript.format.enable": false,
  "eslint.validate": ["javascript", "typescript"]
}
```

## Debugging

<https://code.visualstudio.com/docs/editor/debugging>

### Nodemon + TypeScript

[Nodemon과 debugger attaching](https://code.visualstudio.com/docs/nodejs/nodejs-debugging#_attaching-to-nodejs)

npm script: `"debug": "nodemon --watch src --exec \"node --inspect -r ts-node/register\" src/main.ts",`

`.vscode/launch.json`

```json
{
  "version": "0.2.0",
  "configurations": [
    {
      // https://github.com/Microsoft/vscode-recipes/tree/master/nodemon
      "type": "node",
      "request": "attach",
      "name": "Node: Nodemon",
      "internalConsoleOptions": "neverOpen",
      "protocol": "inspector",
      "processId": "${command:PickProcess}",
      "restart": true
    }
  ]
}
```

### TypeScript

<https://code.visualstudio.com/docs/typescript/typescript-debugging>

`.vscode/launch.json`

```json
{
  "version": "0.2.0",
  "configurations": [
    {
      "type": "node",
      "request": "launch",
      "name": "Debug TypeScript in Node.js",
      "preLaunchTask": "tsc: build - tsconfig.json",
      "program": "${workspaceFolder}/src/main.ts",
      "cwd": "${workspaceFolder}",
      "protocol": "inspector",
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
> `~/.profile` 혹은 `~/.zprofile` 같은 환경변수 설정에서 기본 node/bin PATH를 지정한다.

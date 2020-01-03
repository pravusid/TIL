# Visual Studio Code

<https://github.com/microsoft/vscode-recipes>

<https://code.visualstudio.com/docs>

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

> `main.ts` 경로와 `dist/` 경로는 설정에 따라변경

### Jest extension + NVM

[`nvm` 사용시 `PATH`에서 node runtime을 찾지 못하는 문제](https://code.visualstudio.com/docs/nodejs/nodejs-debugging#_multi-version-support)

> 옵션에 `"runtimeVersion": "<NODE_VESION_IN_NVM>"` 추가

`.vscode/launch.json`

```json
{
  "version": "0.2.0",
  "configurations": [
    {
      "type": "node",
      "name": "vscode-jest-tests",
      "request": "launch",
      "program": "${workspaceFolder}/node_modules/jest/bin/jest",
      "args": ["--runInBand"],
      "cwd": "${workspaceFolder}",
      "runtimeVersion": "<NODE_VESION_IN_NVM>",
      "console": "internalConsole",
      "internalConsoleOptions": "openOnSessionStart",
      "disableOptimisticBPs": true
    }
  ]
}
```

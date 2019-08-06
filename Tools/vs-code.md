# Visual Studio Code

## Debugging

<https://code.visualstudio.com/docs/editor/debugging>

### Node.js

[Nodemon과 debugger attaching](https://code.visualstudio.com/docs/nodejs/nodejs-debugging#_attaching-to-nodejs)

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
      "args": [
        "--runInBand"
      ],
      "cwd": "${workspaceFolder}",
      "runtimeVersion": "<NODE_VESION_IN_NVM>",
      "console": "internalConsole",
      "internalConsoleOptions": "openOnSessionStart",
      "disableOptimisticBPs": true
    }
  ]
}
```

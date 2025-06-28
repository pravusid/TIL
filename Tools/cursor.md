# Cursor Editor

<https://www.cursor.com/>

## Prompt

<https://docs.cursor.com/context/rules>

- <https://github.com/PatrickJS/awesome-cursorrules>
- <https://github.com/sanjeed5/awesome-cursor-rules-mdc>
- <https://cursor.directory/>
- [You are using Cursor AI incorrectly](https://ghuntley.com/stdlib/)
- [My Best Practices for MDC rules and troubleshooting](https://forum.cursor.com/t/my-best-practices-for-mdc-rules-and-troubleshooting/50526)

## MCP

<https://docs.cursor.com/context/model-context-protocol>

- [Those MCP totally 10x my Cursor workflow](https://www.youtube.com/watch?v=oAoigBWLZgE)

## Security

- <https://www.cursor.com/security>
- <https://forum.cursor.com/t/questions-on-gitignore-cursorignore-cursorban/34713>
- <https://forum.cursor.com/t/environment-secrets-and-code-security/14486>

<https://docs.cursor.com/context/ignore-files>

> Cursor will also ignore all files listed in the **.gitignore** file in your root directory and in the **Default Ignore List** provided below.

## 단축키 설정 ([[vs-code]] compatible)

- <https://forum.cursor.com/t/cmd-k-vs-cmd-r-keyboard-shortcuts-default/1172/9>
- <https://forum.cursor.com/t/comment-code-using-cursor-ai/11683>

<https://marketplace.visualstudio.com/items?itemName=YuTengjing.vscode-classic-experience> 확장 설치

확장 단축키 추가 변경: [Copilot Chat](https://code.visualstudio.com/docs/copilot/copilot-chat), [Copilot Edit](https://code.visualstudio.com/docs/copilot/copilot-edits) 단축키 참고

```jsonc
[
  //
  // Cursor
  //
  {
    "key": "cmd+i",
    "command": "aipopup.action.modal.generate",
    "when": "editorFocus && !composerBarIsVisible"
  },
  {
    "key": "shift+cmd+i",
    "command": "composerMode.agent"
  },
  {
    "key": "ctrl+cmd+i",
    "command": "composerMode.chat"
  },
  {
    "key": "cmd+i",
    "command": "cursorai.action.generateInTerminal",
    "when": "terminalFocus && terminalHasBeenCreated || terminalFocus && terminalProcessSupported || terminalFocus && terminalHasBeenCreated && terminalProcessSupported"
  },
  {
    "key": "cmd+e",
    "command": "-aipopup.action.modal.generate",
    "when": "editorFocus && !composerBarIsVisible"
  },
  {
    "key": "cmd+e",
    "command": "-cursorai.action.generateInTerminal",
    "when": "terminalFocus && terminalHasBeenCreated || terminalFocus && terminalProcessSupported || terminalFocus && terminalHasBeenCreated && terminalProcessSupported"
  },
  {
    "key": "cmd+]",
    "command": "-aichat.newchataction",
    "when": "!view.workbench.panel.aichat.view.visible"
  },
  {
    "key": "cmd+]",
    "command": "-aichat.close-sidebar",
    "when": "view.workbench.panel.aichat.view.visible"
  },
  {
    "key": "shift+cmd+]",
    "command": "-aichat.insertselectionintochat"
  },
  {
    "key": "alt+cmd+.",
    "command": "cmdk.togglePromptBarModel",
    "when": "editorHasPromptBar && editorPromptBarFocused"
  },
  {
    "key": "alt+cmd+.",
    "command": "composer.openModelToggle",
    "when": "composerFocused && !editorTextFocus"
  }
]
```

# Cursor Editor

<https://www.cursor.com/>

## Prompt

- <https://cursor.directory/>
- [React18 Component -> React19 Component Refactor Prompt](https://gist.github.com/ellemedit/46cb6ac6a8c65aa69e010b1c88f406c3)

## 단축키 설정 ([[vs-code]] compatible)

- <https://forum.cursor.com/t/cmd-k-vs-cmd-r-keyboard-shortcuts-default/1172/9>
- <https://forum.cursor.com/t/comment-code-using-cursor-ai/11683>

<https://marketplace.visualstudio.com/items?itemName=YuTengjing.vscode-classic-experience> 확장 설치

확장 단축키 추가 변경: [Copilot Chat](https://code.visualstudio.com/docs/copilot/copilot-chat), [Copilot Edit](https://code.visualstudio.com/docs/copilot/copilot-edits) 단축키 참고

```json
[
  {
    "key": "cmd+i",
    "command": "aipopup.action.modal.generate",
    "when": "editorFocus && !composerBarIsVisible && !composerControlPanelIsVisible",
    "args": { "invocationType": "new" }
  },
  {
    "key": "shift+cmd+i",
    "command": "aipopup.action.modal.generate",
    "when": "editorFocus && !composerBarIsVisible && !composerControlPanelIsVisible",
    "args": { "invocationType": "toggle" }
  },
  {
    "key": "cmd+e",
    "command": "-aipopup.action.modal.generate",
    "when": "editorFocus && !composerBarIsVisible && !composerControlPanelIsVisible"
  },
  {
    "key": "cmd+i",
    "command": "cursorai.action.generateInTerminal",
    "when": "terminalFocus && terminalHasBeenCreated || terminalFocus && terminalProcessSupported || terminalFocus && terminalHasBeenCreated && terminalProcessSupported"
  },
  {
    "key": "cmd+e",
    "command": "-cursorai.action.generateInTerminal",
    "when": "terminalFocus && terminalHasBeenCreated || terminalFocus && terminalProcessSupported || terminalFocus && terminalHasBeenCreated && terminalProcessSupported"
  },
  {
    "key": "cmd+e",
    "command": "-composer.startComposerPrompt",
    "when": "composerIsEnabled"
  },
  {
    "key": "ctrl+cmd+i",
    "command": "aichat.newchataction"
  },
  {
    "key": "cmd+]",
    "command": "-aichat.newchataction"
  },
  {
    "key": "shift+cmd+]",
    "command": "-aichat.insertselectionintochat"
  },
  {
    "key": "shift+cmd+e",
    "command": "-aichat.fixerrormessage",
    "when": "(arbitrary function)"
  },
  {
    "key": "ctrl+cmd+i",
    "command": "aichat.insertselectionintochat",
    "when": "terminalFocus && terminalHasBeenCreated || terminalFocus && terminalProcessSupported || terminalFocus && terminalHasBeenCreated && terminalProcessSupported"
  }
]
```

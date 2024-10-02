# Cursor Editor

<https://www.cursor.com/>

## 단축키 설정 ([[vs-code]] compatible)

- <https://forum.cursor.com/t/cmd-k-vs-cmd-r-keyboard-shortcuts-default/1172/9>
- <https://forum.cursor.com/t/comment-code-using-cursor-ai/11683>

<https://marketplace.visualstudio.com/items?itemName=YuTengjing.vscode-classic-experience> 확장 설치

확장 단축키 추가 변경

```json
[
  {
    "key": "ctrl+shift+i",
    "command": "aipopup.action.modal.generate",
    "when": "editorFocus && !composerBarIsVisible && !composerControlPanelIsVisible",
    "args": { "invocationType": "toggle" }
  },
  {
    "key": "ctrl+i",
    "command": "aipopup.action.modal.generate",
    "when": "editorFocus && !composerBarIsVisible && !composerControlPanelIsVisible",
    "args": { "invocationType": "new" }
  },
  {
    "key": "cmd+e",
    "command": "-aipopup.action.modal.generate",
    "when": "editorFocus && !composerBarIsVisible && !composerControlPanelIsVisible"
  },
  {
    "key": "ctrl+i",
    "command": "cursorai.action.generateInTerminal",
    "when": "terminalFocus && terminalHasBeenCreated || terminalFocus && terminalProcessSupported || terminalFocus && terminalHasBeenCreated && terminalProcessSupported"
  },
  {
    "key": "cmd+e",
    "command": "-cursorai.action.generateInTerminal",
    "when": "terminalFocus && terminalHasBeenCreated || terminalFocus && terminalProcessSupported || terminalFocus && terminalHasBeenCreated && terminalProcessSupported"
  },
  {
    "key": "ctrl+o",
    "command": "aichat.newchataction"
  },
  {
    "key": "cmd+]",
    "command": "-aichat.newchataction"
  }
]
```

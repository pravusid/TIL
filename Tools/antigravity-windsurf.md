# Antigravity | Windsurf Editor

- <https://antigravity.google/docs/get-started>
- <https://docs.windsurf.com/windsurf/getting-started>

## 단축키 설정 ([[vs-code]] compatible)

```jsonc
[
  //
  // Antigravity | Windsurf
  //
  {
    "key": "cmd+l",
    "command": "-antigravity.prioritized.chat.open",
    "when": "!terminalFocus"
  },
  {
    "key": "cmd+l",
    "command": "-antigravity.prioritized.chat.openFromTerminal",
    "when": "terminalFocus"
  },
  {
    "key": "ctrl+cmd+i",
    "command": "antigravity.prioritized.chat.open",
    "when": "!terminalFocus"
  },
  {
    "key": "ctrl+cmd+i",
    "command": "antigravity.prioritized.chat.openFromTerminal",
    "when": "terminalFocus"
  },
  {
    "key": "shift+cmd+l",
    "command": "-antigravity.prioritized.chat.openNewConversation"
  },
  {
    "key": "cmd+n",
    "command": "antigravity.prioritized.chat.openNewConversation",
    "when": "!editorFocus && !terminalFocus && !sideBarFocus && !auxiliaryBarFocus && !panelFocus"
  }
]
```

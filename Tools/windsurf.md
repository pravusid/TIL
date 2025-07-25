# Windsurf Editor

<https://docs.windsurf.com/windsurf/getting-started>

## 단축키 설정 ([[vs-code]] compatible)

```jsonc
[
  //
  // Windsurf
  //
  {
    "key": "cmd+l",
    "command": "-windsurf.prioritized.chat.open",
    "when": "!terminalFocus"
  },
  {
    "key": "cmd+l",
    "command": "-windsurf.prioritized.chat.openFromTerminal",
    "when": "terminalFocus"
  },
  {
    "key": "ctrl+cmd+i",
    "command": "windsurf.prioritized.chat.open",
    "when": "!terminalFocus"
  },
  {
    "key": "ctrl+cmd+i",
    "command": "windsurf.prioritized.chat.openFromTerminal",
    "when": "terminalFocus"
  },
  {
    "key": "shift+cmd+l",
    "command": "-windsurf.prioritized.chat.openNewConversation"
  },
  {
    "key": "cmd+n",
    "command": "windsurf.prioritized.chat.openNewConversation",
    "when": "!editorFocus && !terminalFocus && !sideBarFocus && !auxiliaryBarFocus && !panelFocus"
  }
]
```

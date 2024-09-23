# Shell History 복구

<https://unix.stackexchange.com/questions/491227/how-can-i-recover-a-corrupted-zsh-history-file-from-memory>

실수로 `$HISTFILE` 삭제한 경우, 세션 종료하기 전 다음 명령어 실행

```bash
fc -W .zsh_history.bak
```

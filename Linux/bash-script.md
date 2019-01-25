# Bash Shell Script

## 예제

### process id

```sh
#!/bin/bash

PROC_NAME=$1
PID=$(pgrep -f $PROC_NAME)
echo "$PID"
```

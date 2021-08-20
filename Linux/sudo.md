# SUDO

> superuser do, substitute user do

## sudoer 설정

```sh
export EDITOR=vim
visudo
```

> 등록한 사용자는 `sudo` 포함 명령 실행시 비밀번호를 입력하지 않는다

```txt
# User privilege specification
<USER>    ALL=(ALL:ALL) NOPASSWD:ALL
```

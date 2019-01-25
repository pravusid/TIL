# 작업관리

## 백그라운드 실행

프로그램 실행 시 끝에 &를 붙여 백그라운드로 실행 시킬 수 있음

`nohup` 명령어로 백그라운드 실행하면 hangup을 무시하며 로그아웃 후에도 실행상태로 남아있음

```sh
java –jar foo.jar &
nohup java -jar foo.jar
nohup java -jar foo.jar &
```

### CRTL + Z

프로그램 실행 중 `Ctrl + Z`를 누르면 백그라운드로 보냄 (일시정지)

1. `jobs`: 백그라운드 실행중인 프로세스 목록을 보여줌
2. `fg %[숫자]`: `jobs`의 해당 프로세스를 foreground로 보냄 (다시 화면에 출력)
3. `bg %[숫쟈]`: `jobs`의 해당 프로세스를 background에서 실행시킴 (일시정지->실행)

## 프로세스 종료

- 프로세스 목록: `ps –[옵션]`: `e`(모든프로세스), `f`(Full format), `l`(상세정보), `x`(보이지 않는 프로세스도)
- 특정 프로세스 목록: `ps –ef | grep [프로세스이름]`
- 특정 프로세스 목록: `pgrep -[옵션] [프로세스이름]` (위 명령어의 alias) (`f`: 패턴매칭, `l`: 프로세스 이름도 출력)
- 종료(pid) : `kill -[옵션] (pid)`: `l`(옵션목록)), `9`(강제종료), `15`(종료시그널)
- 종료(pname): `pkill -[옵션] [프로세스이름]`

### kill signal

주로 사용하는 시그널은 다음과 같다

- -1 (-HUP) - restart a process
- -2 (-INT) - terminate a process
- -9 (-KILL) - let the kernel kick the process out
- -11 (-SEGV) - have the program crash brutally
- -15 (-TERM) - the default, ask the program kindly to terminate.

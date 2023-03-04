# TimeZone

## 시스템 시간 확인

`date`

## 타임존 변경 (to KST)

- `sudo ln -sf /usr/share/zoneinfo/Asia/Seoul /etc/localtime`
- `timedatectl set-timezone Asia/Seoul`

> reboot 이후 유지됨

## 타임존 변경 (interactive)

`tzselect`

이 경우 실행후 출력되는 다음 형식의 환경변수를 설정해야 함

```bash
You can make this change permanent for yourself by appending the line
        TZ='Asia/Seoul'; export TZ
to the file '.profile' in your home directory; then log out and log in again.
```

> [[docker-command]] run env 설정, [[docker-dockerfile]] env 설정에 `TZ`를 선언하면 실행환경의 타임존을 설정할 수 있음

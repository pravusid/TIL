# TimeZone

## 시스템 시간 확인

`date`

## 타임존 변경(to KST)

- `timedatectl set-timezone Asia/Seoul`
- `ln -sf /usr/share/zoneinfo/Asia/Seoul /etc/localtime`
- `tzselect`

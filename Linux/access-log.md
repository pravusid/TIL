# 리눅스 접속 기록

## 조회

- `/var/run/utmp`: 현재 접속해있는 사용자 정보 (로그아웃시 삭제됨)
- `/var/log/wtmp`: 로그인/로그아웃 정보 (`last` 명령어)
- `/var/log/lastlog`: 가장 최근 로그인 정보 (로그인시 가장 처음 출력됨)

## 삭제

- `cat /dev/null > /var/log/wtmp`
- `cat /dev/null > /var/log/lastlog`

# LogRotate

일반적인 리눅스 배포판에서 LogRotate는 기본적으로 실행되고 있음

## 파일

- daemon: `/usr/sbin/logrotate`
- 설정: `/etc/logrotate.conf`
- 설정: `/etc/logrotate.d/`
- Logrotate 로그 : `/etc/cron.daily/logrotate`

## 설정

`/etc/logrotate.conf` 설정 파일에서 `include /etc/logrotate.d` 경로의 설정도 참조하여 적용한다.

따라서 세부설정(어플리케이션 별)은 `/etc/logrotate.d/`를 활용하는 것이 좋다.

```conf
/home/idpravus/myapp/*.log {
    su idpravus idpravus
    daily
    rotate 90
    compress
    dateext
    missingok
    notifempty
    size +4096k
    create 644 root root
    sharedscripts
    postrotate
      /etc/init.d/syslog reload
    endscript
}
```

- `su <user> <group>`: rotate 대상을 제어하기 위해서 사용자를 변경함(default: root가 명령을 실행함)
- rotation 주기: daily, weekly, monthly, yearly
- rotate: 유지할 파일의 개수
- compress || nocompress: 파일 압축 여부
- dateext: 파일 이름에 날짜 추가
- missingok: 파일이 없어도 오류를 발생시키지 않음
- notifempty: 파일이 비어있으면 rotate 하지 않음
- size: size 보다 로그파일이 클 경우 로테이션을 수행 (k -> kilobyte, m -> megabyte)
- postrotate: rotate 후 실행할 스크립트
  - rotate를 실행하면 원본파일을 백업 파일로 복사한 후에 삭제한다
  - 따라서 로그를 생성하던 애플리케이션은 파일에 대한 참조를 잃어버려서 로그를 쌓을 수 없게 된다
  - 이를 처리하기 위해 추가로 명령을 실행한다
- copytruncate
  - 애플리케이션이 로그파일을 새로 만들기 위한 시그널 처리 코드를 가지고 있지 않은 경우
  - copytruncate를 이용하면 원본파일을 지우지 않고 파일 크기를 0으로 만든다(truncate)
  - 파일을 복사하고 truncate 하는 순간 로그를 잃어버릴 수도 있다

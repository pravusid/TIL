# supervisor

<http://supervisord.org/index.html>

python으로 제작된 프로세스 모니터링/관리 도구이다

## 설치

`sudo apt-get install supervisor`

## 설정

`/etc/supervisor/conf.d/<프로그램명>.conf`

```text
[program:<프로그램명>]
environment=<변수1>=<내용1>,<변수2>=<내용2>
command=/bin/cat
process_name=%(program_name)s // process_name expr (default %(program_name)s)
numprocs=1 // number of processes copies to start (def 1)
user=ubuntu
```

## 실행

<http://supervisord.org/running.html?highlight=reread#supervisorctl-actions>

```shell
sudo supervisorctl reread
sudo supervisorctl update
sudo supervisorctl start <프로그램명>
sudo supervisorctl start all
sudo supervisorctl status
```

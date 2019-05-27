# Jenkins

## Jenkins 다른 user로 실행

```sh
sudo vim /etc/sysconfig/jenkins
sudo vim /etc/default/jenkins # in debian

# 다음 내용을 수정함
$JENKINS_USER="idpravus"

# 이미 설치되어 jenkins:jenkins 소유의 파일 소유권을 변경
chown -R idpravus:idpravus /var/lib/jenkins
chown -R idpravus:idpravus /var/cache/jenkins
chown -R idpravus:idpravus /var/log/jenkins

# Jenkins 재시작
sudo /etc/init.d/jenkins restart
```

## Jenkins Execute Shell에서 백그라운드 실행

```sh
if pgrep -f idpravus-app; then kill -15 $(pgrep -f idpravus-app); fi
npm ci
BUILD_ID=dontKillMe nohup npm run serve &
```

# 방화벽 정책

## CentOS

### 설정

1. 설정 파일: `/etc/firewalld/zones/public.xml`
2. 퍼블릭 포트 추가: `firewall-cmd --permanent --zone=public --add-port=8080/tcp`
3. 퍼블릭 포트 제거: `firewall-cmd --permanent --zone=public --remove-port=8080/tcp`
4. 임시 추가: 위의 명령에서 `--permanent` 제외
5. ftp 서비스 추가: `firewall-cmd --add-service=ftp`
6. ftp 서비스 제거: `firewall-cmd --add-port=21/tcp`
7. 방화벽 새로고침 (변경사항 적용): `firewall-cmd --reload`

### 관리

1. 방화벽 상태 확인: `firewall-cmd --state`
2. 서비스 리스트: `firewall-cmd --get-service`
3. default zone 목록 : `firewall-cmd --get-default-zone`
4. active zone 목록: `firewall-cmd --get-active-zones`
5. 사용가능한 서비스/포트 목록: `firewall-cmd --list-all`
6. public존의 사용가능한 서비스/포트 목록: `firewall-cmd --zone=public --list-all`
7. 특정 존에 있는 서비스 리스트: `firewall-cmd --zone=public --list-services`
8. 방화벽 daemon 켜고 끄기

  ```sh
  systemctl start firewalld
  systemctl enable firewalld
  systemctl stop firewalld
  systemctl disable firewalld
  ```

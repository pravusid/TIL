# 방화벽 정책

## iptables

설정이후 재시작: `service iptables restart`

### 명령어

`target = -j targetname [per-target-options]`

- 특정 IP 허용: `iptables -A INPUT -s {ip주소} -j ACCEPT`
- 특정 IP 차단: `iptables -A INPUT -s {ip주소} -j DROP`
- 특정 Port 허용: `iptables -A INPUT -p tcp –dport {포트} -j ACCEPT`
- 특정 Port 차단: `iptables -A INPUT -p tcp –dport {포트} -j DROP`
- 특정 IP와 Port 차단: `iptables -A INPUT -s {ip주소} -p tcp –dport {port} -j DROP`
- 로그 설정: `iptables -I INPUT -s {ip주소} -p tcp –dport {port} -j LOG –log-prefix “[denied]”`

## CentOS

- centos7 이전 iptables 사용: `/etc/sysconfig/iptables`
- centos7 부터 firewall 사용

### 설정

- 설정 파일: `/etc/firewalld/zones/public.xml`
- 퍼블릭 포트 추가: `firewall-cmd --permanent --zone=public --add-port=8080/tcp`
- 퍼블릭 포트 제거: `firewall-cmd --permanent --zone=public --remove-port=8080/tcp`
- 임시 추가: 위의 명령에서 `--permanent` 제외
- ftp 서비스 추가: `firewall-cmd --add-service=ftp`
- ftp 서비스 제거: `firewall-cmd --add-port=21/tcp`
- 방화벽 새로고침 (변경사항 적용): `firewall-cmd --reload`

### 관리

- 방화벽 상태 확인: `firewall-cmd --state`
- 서비스 리스트: `firewall-cmd --get-service`
- default zone 목록 : `firewall-cmd --get-default-zone`
- active zone 목록: `firewall-cmd --get-active-zones`
- 사용가능한 서비스/포트 목록: `firewall-cmd --list-all`
- public존의 사용가능한 서비스/포트 목록: `firewall-cmd --zone=public --list-all`
- 특정 존에 있는 서비스 리스트: `firewall-cmd --zone=public --list-services`
- 방화벽 daemon 켜고 끄기

  ```sh
  systemctl start firewalld
  systemctl enable firewalld
  systemctl stop firewalld
  systemctl disable firewalld
  ```

## Ubuntu

우분투의 기본 방화벽은 ufw이다. ufw의 gui버전인 Gufw를 사용할 수도 있다.

설정파일은 `/etc/ufw/` 경로에서 확인할 수 있다.

### 기본문법

- 활성화 / 비활성화
  - 활성화: `sudo ufw enable`
  - 비활성화: `sudo ufw disable`

- 기본정책 확인(in-모두거부, out-모두허용): `sudo ufw show raw`
  - 기본정책으로 허용: `sudo ufw default allow`
  - 기본정책으로 차단: `sudo ufw default deny`

- 상태확인: `sudo ufw status verbose`

- Allow / Deny
  - Allow: `sudo ufw allow <port>/<optional: protocol>`
    - 예제: To allow incoming tcp and udp packet on port 53: `sudo ufw allow 53`
    - 예제: To allow incoming tcp packets on port 53: `sudo ufw allow 53/tcp`
    - 예제: To allow incoming udp packets on port 53: `sudo ufw allow 53/udp`

  - Deny: `sudo ufw deny <port>/<optional: protocol>`
    - 예제: To deny tcp and udp packets on port 53: `sudo ufw deny 53`
    - 예제: To deny incoming tcp packets on port 53: `sudo ufw deny 53/tcp`
    - 예제: To deny incoming udp packets on port 53: `sudo ufw deny 53/udp`

- 룰 삭제
  - ufw deny 80/tcp 룰을 삭제하려면: `sudo ufw delete deny 80/tcp`

- 서비스: 서비스 이름을 방화벽에 등록할 수 있다 (서비스 목록: `/etc/services`)
  - Allow by Service Name: `sudo ufw allow <service name>`
    - 예제: to allow ssh by name: `sudo ufw allow ssh`
  - Deny by Service Name: `sudo ufw deny <service name>`
    - 예제: to deny ssh by name: `sudo ufw deny ssh`

- 방화벽 상태확인: `sudo ufw status`

- Logging
  - 로깅 on: `sudo ufw logging on`
  - 로깅 off: `sudo ufw logging off`

### 심화 문법

- 접근 허용
  - 특정 IP 허용: `sudo ufw allow from <ip address>`
    - 예제:To allow packets from 207.46.232.182: `sudo ufw allow from 207.46.232.182`

  - 특정 subnet 허용
    - netmask를 사용한다: `sudo ufw allow from 192.168.1.0/24`

  - 특정 Port와 IP를 허용: `sudo ufw allow from <target> to <destination> port <port number>`
    - 예제: allow IP address 192.168.0.4 access to port 22 for all protocols: `sudo ufw allow from 192.168.0.4 to any port 22`

  - 특정 Port와 IP와 프로토콜 허용: `sudo ufw allow from <target> to <destination> port <port number> proto <protocol name>`
    - 예제: allow IP address 192.168.0.4 access to port 22 using TCP: `sudo ufw allow from 192.168.0.4 to any port 22 proto tcp`

- 핑 허용
  - 기본적으로 UFW는 ping requests를 허용한다..
  - 핑 요청을 해제하기 위해서 `/etc/ufw/before.rules`의 다음 라인을 삭제해야 한다

  ```text
  # ok icmp codes
  -A ufw-before-input -p icmp --icmp-type destination-unreachable -j ACCEPT
  -A ufw-before-input -p icmp --icmp-type source-quench -j ACCEPT
  -A ufw-before-input -p icmp --icmp-type time-exceeded -j ACCEPT
  -A ufw-before-input -p icmp --icmp-type parameter-problem -j ACCEPT
  -A ufw-before-input -p icmp --icmp-type echo-request -j ACCEPT

  or change the "ACCEPT" to "DROP"

  # ok icmp codes
  -A ufw-before-input -p icmp --icmp-type destination-unreachable -j DROP
  -A ufw-before-input -p icmp --icmp-type source-quench -j DROP
  -A ufw-before-input -p icmp --icmp-type time-exceeded -j DROP
  -A ufw-before-input -p icmp --icmp-type parameter-problem -j DROP
  -A ufw-before-input -p icmp --icmp-type echo-request -j DROP
  ```

- 접근 거부
  - 특정 IP 거부: `sudo ufw deny from <ip address>`
    - 예제:To block packets from 207.46.232.182: `sudo ufw deny from 207.46.232.182`

  - 특정 Port와 IP를 거부: `sudo ufw deny from <ip address> to <protocol> port <port number>`
    - 예제: deny ip address 192.168.0.1 access to port 22 for all protocols: `sudo ufw deny from 192.168.0.1 to any port 22`

- status의 순서(룰의 id 번호)를 이용해서 수정할 수 있다
  - 숫자와 함께 룰 목록 출력: `sudo ufw status numbered`
  - 룰 삭제(번호): `sudo ufw delete 1`
  - 룰 추가(번호): `sudo ufw insert 1 allow from <ip address>`

### 심화 예제

192.168.0.1, 192.168.0.7으로 부터오는 22번 포트 요청을 막고싶지만, 다른 192.168.0.x 대역의 22번 포트 tcp 요청은 허용하려는 경우.

```sh
sudo ufw deny from 192.168.0.1 to any port 22
sudo ufw deny from 192.168.0.7 to any port 22
sudo ufw allow from 192.168.0.0/24 to any port 22 proto tcp
```

세부사항의 룰을 우선 선언하고, 일반 룰은 나중에 선언해야 한다. 룰을 해석하면 이에 matching 하는 다른 룰은 해석하지 않기 때문이다.

```sh
sudo ufw status

To                         Action  From
--                         ------  ----
22:tcp                     DENY    192.168.0.1
22:udp                     DENY    192.168.0.1
22:tcp                     DENY    192.168.0.7
22:udp                     DENY    192.168.0.7
22:tcp                     ALLOW   192.168.0.0/24
```

192.168.0.3 주소의 22번 포트요청을 추가로 막으려는 경우

```sh
sudo ufw delete allow from 192.168.0.0/24 to any port 22
sudo ufw status

To                         Action  From
--                         ------  ----
22:tcp                     DENY    192.168.0.1
22:udp                     DENY    192.168.0.1
22:tcp                     DENY    192.168.0.7
22:udp                     DENY    192.168.0.7
```

일반 룰을 우선 삭제하고 다시 세부룰을 추가한 뒤, 일반 룰을 다시 추가한다

```sh
sudo ufw deny 192.168.0.3 to any port 22
sudo ufw allow 192.168.0.0/24 to any port 22 proto tcp
sudo ufw status

To                         Action  From
--                         ------  ----
22:tcp                     DENY    192.168.0.1
22:udp                     DENY    192.168.0.1
22:tcp                     DENY    192.168.0.7
22:udp                     DENY    192.168.0.7
22:tcp                     DENY    192.168.0.3
22:udp                     DENY    192.168.0.3
22:tcp                     ALLOW   192.168.0.0/24
```

### 로그

#### Log Entries

Pseudo Log Entry

```text
Feb  4 23:33:37 hostname kernel: [ 3529.289825] [UFW BLOCK] IN=eth0 OUT= MAC=00:11:22:33:44:55:66:77:88:99:aa:bb:cc:dd SRC=444.333.222.111 DST=111.222.333.444 LEN=103 TOS=0x00 PREC=0x00 TTL=52 ID=0 DF PROTO=UDP SPT=53 DPT=36427 LEN=83
```

- Date: It's good practice to watch the dates and times. If things are out of order or blocks of time are missing then an attacker probably messed with your logs.
- Hostname: The server’s hostname
- Uptime: The time in seconds since boot.

### logged event 용어에 관한 설명

- IN: If set, then the event was an incoming event.
- OUT: If set, then the event was an outgoing event.
- MAC: This provides a 14-byte combination of the Destination MAC, Source MAC, and EtherType fields, following the order found in the Ethernet II header. See Ethernet frame and EtherType for more information.
- SRC: This indicates the source IP, who sent the packet initially. Some IPs are routable over the internet, some will only communicate over a LAN, and some will only route back to the source computer. See IP address for more information.
- DST: This indicates the destination IP, who is meant to receive the packet. You can use whois.net or the cli whois to determine the owner of the IP address.
- LEN: This indicates the length of the packet.
- TOS: I believe this refers to the TOS field of the IPv4 header. See TCP Processing of the IPv4 Precedence Field for more information.
- PREC: I believe this refers to the Precedence field of the IPv4 header.
- TTL: This indicates the “Time to live” for the packet. Basically each packet will only bounce through the given number of routers before it dies and disappears. If it hasn’t found its destination before the TTL expires, then the packet will evaporate. This field keeps lost packets from clogging the internet forever. See Time to live for more information.
- ID: Not sure what this one is, but it's not really important for reading logs. It might be ufw’s internal ID system, it might be the operating system’s ID.
- PROTO: This indicates the protocol of the packet - TCP or UDP. See TCP and UDP Ports Explained for more information.
- SPT: This indicates the source. I believe this is the port, which the SRC IP sent the IP packet over. See List of TCP and UDP port numbers for more information.
- DPT: This indicates the destination port. I believe this is the port, which the SRC IP sent its IP packet to, expecting a service to be running on this port.
- WINDOW: This indicates the size of packet the sender is willing to receive.
- RES: This bit is reserved for future use & is always set to 0. Basically it’s irrelevant for log reading purposes.
- SYN URGP: SYN indicates that this connection requires a three-way handshake, which is typical of TCP connections. URGP indicates whether the urgent pointer field is relevant. 0 means it's not. Doesn’t really matter for firewall log reading.

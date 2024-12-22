# SSH

시큐어 셸(Secure Shell, SSH)은 네트워크 상의 다른 컴퓨터에 로그인하거나 원격 시스템에서 명령을 실행하고 다른 시스템으로 파일을 복사할 수 있도록 해 주는 응용 프로그램 또는 그 프로토콜을 가리킨다.

기존의 rsh, rlogin, 텔넷 등을 대체하기 위해 설계되었으며, 강력한 인증 방법 및 안전하지 못한 네트워크에서 안전하게 통신을 할 수 있는 기능을 제공한다.

기본적으로는 22번 포트를 사용한다.

## 설치

`sudo apt install openssh-server`

## 기본명령어

- `-b bind_address`: ip가 두개이상인 경우
- `-E log_file`
- `-e escape_char`
- `-F configfile`
- `-i identity_file`
- `-J [user@]host[:port]` (ProxyJump)
- `-p port`

`ssh [user@]hostname [command]`

## 인증

### SSH 비대칭 키 발급

RSA(ed25519) 방식의 비대칭키를 생성함

```sh
ssh-keygen -t rsa -C "comment"
ssh-keygen -t ed25519 -C "comment"
```

Convert `BEGIN OPENSSH PRIVATE KEY` to `BEGIN RSA PRIVATE KEY`:

```sh
cd ~/.ssh
cp id_rsa id_rsa.pem
ssh-keygen -p -m PEM -f id_rsa.pem
```

### 비대칭 키를 사용해 인증 (비밀번호 입력 대신)

공개키를 서버의 `authorized_keys`에 등록: `ssh-copy-id user@host`

서버에서 권한 설정

> permissions are too open 오류

```sh
chmod 700 ~/.ssh
chmod 600 ~/.ssh/id_rsa
chmod 644 ~/.ssh/id_rsa.pub
chmod 644 ~/.ssh/authorized_keys
chmod 644 ~/.ssh/known_hosts
```

### ssh 연결을 비대칭 키만 사용(비밀번호 사용X)

`/etc/ssh/sshd_config` 설정파일 수정

```txt
PubkeyAuthentication yes
PasswordAuthentication no
```

### 공개키 복사

원격서버에서 패스워드 인증을 사용하는 경우

> `-i` 옵션을 사용하지 않으면 기본값은 `id_rsa.pub`임

`ssh-copy-id [-i pubkey_file] <user>@<host>`

원격서버에서 공개키 인증을 사용하는 경우

`ssh <user>@<host> 'cat >> ~/.ssh/authorized_keys' < ~/.ssh/id_rsa.pub`

## SSH config 설정

`.ssh/config`

```sh
Host *
    ServerAliveInterval 60

Host <host-alias1> [<host-alias2> ...]
    HostName <remote-host>
    User <username>
    IdentityFile ~/.ssh/my-identity.pem
```

Identity 파일을 지정하지 않으면 `.ssh/id_rsa`가 기본으로 사용된다

### Host Key Checking 비활성화

`~/.ssh/known_hosts` 등록시 키검증 생략 (MITM 위험성)

`.ssh/config`

```sh
Host *
    StrictHostKeyChecking no
```

### 보안관련 설정

`/etc/ssh/sshd_config`

```conf
# root 로그인 차단
PermitRootLogin no
PermitRootLogin prohibit-password

# 공개키 인증 사용
PubkeyAuthentication yes

# 패스워드 인증 차단 (기본 값: no)
PasswordAuthentication no
PermitEmptyPasswords no

# 설정 적용
sudo systemctl restart ssh
```

## Secure Copy (SCP)

```sh
scp <옵션> <원본 경로 및 파일명> <대상 경로 및 파일명>
```

- 옵션

  - `-P` : 포트번호
  - `-p` : 원본 파일 시간의 수정시간, 사용시간, 권한을 유지한다 (preserve)
  - `-r` : 하위 폴더/파일 모두 복사한다 (recursive)

- 사용 예

  - 보내기 : `scp -rp 파일명 user@host:~/다운로드/파일명`
  - 받기 : `scp -rp user@host:~/다운로드/파일명 로컬경로/파일명`

## SSH Port Forwarding (Tunneling)

<https://en.wikibooks.org/wiki/OpenSSH/Cookbook/Tunnels>

- `C` (압축하여 접속)
- `N` (shell 명령어 실행 금지)
- `f` (백그라운드로 실행)

### Local Forwarding

`ssh -L <PORT1>:remote:<PORT2> [user@]server`

다음의 연결이 성립한다: `localhost:<PORT1> -> server == tunnel -> remote:<PORT2>`

remote는 server에서 도달할 수 있는 `<hostname | ip address>`를 사용하여야 한다

`ssh -i ~/.ssh/id_rsa -N -L 8080:localhost:3000 -L 8081:192.168.0.100:3001 me@pravusid.kr`

`.ssh/config` 설정으로도 실행가능하다

```conf
Host pravusid
    HostName pravusid.kr
    User me
    LocalForward 8080 localhost:3000
    LocalForward 8081 192.168.0.100:3001
    ExitOnForwardFailure yes
```

### Remote Forwarding

`ssh -R [허용IP:]8080:localhost:80 public.example.com`

다음의 연결이 성립한다: `public.example.com:8080 -> localhost:80`

사용을 위해서 ssh 설정을 변경해야 한다

`/etc/ssh/sshd_config`

```sh
GatewayPorts yes
```

## SSH Proxy

### SSH Agent Forwarding vs SSH ProxyJump

- <https://rabexc.org/posts/pitfalls-of-ssh-agents>
- <https://www.infoworld.com/article/3619278/proxyjump-is-safer-than-ssh-agent-forwarding.html>

ProxyJump를 사용해야 하는 이유

- agent forwarding의 경우 인증체인에서 루트권한이 있다면, agent가 바인딩된 unix-domain 소켓을 사용하여 ssh-agent를 탈취할 수 있음
- forward agent를 사용하지 않는 프록시 점프 사용을 권장함 (OpenSSH 7.3 이상의 버전)
- ProxyJump는 로컬 클라이언트의 표준 입출력을 목적지 호스트로 포워딩한다

### ProxyCommand vs ProxyJump

- <https://en.wikibooks.org/wiki/OpenSSH/Cookbook/Proxies_and_Jump_Hosts>
- <https://goteleport.com/blog/ssh-proxyjump-ssh-proxycommand/>

ProxyJump 명령어는 OpenSSH 7.3버전 이상에서 지원하므로 버전 미만에서는 ProxyCommand를 사용해야 한다

<https://en.wikibooks.org/wiki/OpenSSH/Cookbook/Proxies_and_Jump_Hosts#Old_Methods_of_Passing_Through_Jump_Hosts>

#### 단일 Host 점프

프록시점프는 다음 명령어로 실행할 수 있다

`ssh -J <juser@jump_server:port> dev@192.168.200.200`

> 점프서버 포트 22는 생략가능

`.ssh/config` 설정은 다음과 같다

```conf
Host remote_server
    HostName 192.168.200.200
    User dev
    ProxyJump <juser@jump_server:port>
```

#### 여러 Host 점프

여러 호스트를 대상으로 점프할 수도 있다

`ssh -J <juser1@jump_server1:port>,<juser2@jump_server2:port>,<juser3@jump_server3:port> dev@192.168.200.200`

`.ssh/config` 설정은 다음과 같다

```conf
Host remote_server_multi
    HostName 192.168.200.200
    User dev
    ProxyJump <juser1@jump_server1:port>,<juser2@jump_server2:port>,<juser3@jump_server3:port>
```

#### 특정 조건에서만 ProxyJump를 실행

설정에서 Match를 사용하면 된다

<https://en.wikibooks.org/wiki/OpenSSH/Cookbook/Proxies_and_Jump_Hosts#Conditional_Use_of_Jump_Hosts>

```conf
# 현재 IP주소가 192.168.100.10 인 경우 프록시점프를 실행함
Match host server1 !exec "ifconfig en0 | grep 192.168.100.10"
    ProxyJump user@<jump server>

Host server1
    Hostname 192.168.100.20
```

설정을 수정한 경우 별도의 인자 없이 Proxy를 실행할 수 있다: `ssh remoteserver`

### Bastion Host 설정

- <https://goteleport.com/blog/ssh-bastion-host/>
- <https://goteleport.com/blog/security-hardening-ssh-bastion-best-practices/>

Bation Host는 외부에서 유일하게 연결할 수 있는 SSH 호스트로, 내부의 다른 호스트에 접근하는 기착지가 된다

#### Bastion Host 연결을 위한 클라이언트 설정

`.ssh/config`

```conf
Match User bastionuser
   PermitTTY no
   X11Forwarding no
   PermitTunnel no
   GatewayPorts no
   ForceCommand /usr/sbin/nologin

Host *-myhost
    ProxyJump bastionuser@bastion.example.com
```

#### Bastion Server 설정

`/etc/ssh/sshd_config`

```conf
# Prohibit regular SSH clients from allocating virtual terminals, forward X11, etc
PermitTTY no
X11Forwarding no
PermitTunnel no
GatewayPorts no

# Prohibit launching any remote commands
ForceCommand /usr/sbin/nologin
```

사용자 설정

<https://aws.amazon.com/ko/premiumsupport/knowledge-center/new-user-accounts-linux-instance/>

```sh
# 사용자 생성
sudo useradd -m bastionuser && sudo passwd -d bastionuser
# 사용자 생성 (ubuntu)
sudo useradd -m bastionuser --disabled-password

# 실행 컨텍스트 변경
sudo su - bastionuser

cd /home/bastionuser

mkdir .ssh
chmod 700 .ssh

touch .ssh/authorized_keys
chmod 600 .ssh/authorized_keys

cat >> .ssh/authorized_keys
# 퍼블릭키 입력이 끝나면 ctrl + d

# 사용자 shell 변경
sudo usermod -s /sbin/nologin bastionuser
```

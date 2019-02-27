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

RSA 방식의 비대칭키를 생성함. 비밀번호 대신 public key, private key를 활용해 인증한다.

`ssh-keygen -t rsa -C "comment"`

> permissions are too open 오류: 비공개 키의 권한을 변경한다: `chmod 400 ~/.ssh/id_rsa`

### 비대칭 키를 사용해 인증 (비밀번호 입력 대신)

공개키를 서버의 `authorized_keys`에 등록: `ssh-copy-id user@host`

서버에서 권한 설정

```sh
chmod 700 ~/.ssh
chmod 600 ~/.ssh/id_rsa
chmod 644 ~/.ssh/id_rsa.pub
chmod 644 ~/.ssh/authorized_keys
chmod 644 ~/.ssh/known_hosts
```

### ssh 연결을 비대칭 키만 사용(비밀번호 사용X)

`/etc/ssh/sshd_config` 설정파일 수정

```text
# Change to no to disable tunnelled clear text passwords
PasswordAuthentication no
```

### 공개키 복사

`ssh-copy-id <user>@<host>`

## SSH config 설정

`.ssh/config`

```sh
Host <host-alias>
    HostName <remote-host>
    User <username>
    IdentityFile <~/.ssh/my-identity.pem>
```

Identity 파일을 지정하지 않으면 `.ssh/id_rsa`가 기본으로 사용된다

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

- C(압축하여 접속)
- N(shell 명령어 실행 금지)
- f(백그라운드로 실행)

### Local Forwarding

`ssh -L <PORT1>:remote:<PORT2> [user@]server`

다음의 연결이 성립한다: `localhost:<PORT1> -> server == tunnel -> remote:<PORT2>`

remote는 server에서 도달할 수 있는 `<hostname | ip address>`를 사용하여야 한다

### Remote Forwarding

`ssh -R [허용IP:]8080:localhost:80 public.example.com`

다음의 연결이 성립한다: `public.example.com:8080 -> localhost:80`

사용을 위해서 ssh 옵션을 변경해야 한다

`/etc/ssh/sshd_config`

```sh
GatewayPorts yes
```

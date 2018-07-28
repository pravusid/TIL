# SSH

시큐어 셸(Secure Shell, SSH)은 네트워크 상의 다른 컴퓨터에 로그인하거나 원격 시스템에서 명령을 실행하고 다른 시스템으로 파일을 복사할 수 있도록 해 주는 응용 프로그램 또는 그 프로토콜을 가리킨다. 기존의 rsh, rlogin, 텔넷 등을 대체하기 위해 설계되었으며, 강력한 인증 방법 및 안전하지 못한 네트워크에서 안전하게 통신을 할 수 있는 기능을 제공한다. 기본적으로는 22번 포트를 사용한다.

## SSH 설치

`sudo apt install openssh-server`

## SSH로 명령실행

`ssh user@host 명령어`

## SSH 비대칭 키 발급

RSA 방식의 비대칭키를 생성함. 비밀번호 대신 public key, private key를 활용해 인증한다.

`ssh-keygen -t rsa`

## 비대칭 키를 사용해 인증 (비밀번호 입력 대신))

공개키를 서버의 `authorized_keys`에 등록: `ssh-copy-id user@host`

서버에서 권한 설정

```sh
chmod 700 ~/.ssh
chmod 600 ~/.ssh/authorized_keys
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

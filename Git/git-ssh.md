# Git에서 SSH 프로토콜로 전송 설정

## 비대칭키 생성

1. Git bash를 연다
1. `ssh-keygen -t rsa -b 4096 -C "your@email.com"` 이메일을 넣고 다음 내용을 입력한다.
1. public/private rsa key 쌍이 생성된다.
1. "Enter a file in which to save the key,"에서 `Enter`를 누르면 `~/.ssh`에 저장된다.
1. `Enter passphrase (empty for no passphrase): [Type a passphrase]` 에서 비밀번호를 사용할 수 있다.

## Github 설정

1. Settings
1. SSH and GPG keys
1. SSH keys
1. New SSH key를 눌러 public key를 입력한다.

## 사용

최초 사용시 콘솔에서 `push`, `pull` 작업을 해서 known_hosts에 github를 등록한다.

## 권한오류

`github sign_and_send_pubkey: signing failed: agent refused operation` 오류 발생시

`.ssh` 경로의 private key 권한을 `400`으로 설정한다

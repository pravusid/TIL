# GnuPG

## 설치

`sudo pacman -S gnupg`

<https://wiki.archlinux.org/index.php/GnuPG>

> 명령어 확인: `gpg -h`

## 생성

- `gpg --full-generate-key`
- 알고리즘 `RSA and RSA (default)` 선택
- 키 사이즈 `2048bit` 이상
- 유효기간 2년 이하
- 이름, 이메일, 코멘트 설정

## 권한

- 디렉토리(`~/.gnupg`): 700
- 파일: 600

## 키 목록

- 공개키 목록: `gpg -k`

  - `pub`: 공개키
  - `uid`: user id
  - `sub`: sub key
  - `[SC]`: for Signing & Certificate
  - `[E]`: for Encryption
  - `[expires: yyyy-mm-dd]`: 만료일
  - `trust values`: ultimate, full, marginal, never, undefined, expired, unknown

- 비밀키 목록: `gpg -K`

  - `sec`: 비밀키
  - `ssb`: secret sub key

## 키 관리

- public key 내보내기: `gpg --armor --export <key-id|uid>`
- secret key 내보내기: `gpg --armor --export-secret-keys <key-id|uid>`
- secret key 정보 변경: `gpg --edit-key <key-id|uid>` (help 명령어로 작업확인)
- key 불러오기: `gpg --import <key-file>`
- secret key 삭제: `gpg --delete-secret-key <key-id|uid>`
- public key 삭제: `gpg --delete-key <key-id|uid>`

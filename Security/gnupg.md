# GnuPG

## 설치

- <https://formulae.brew.sh/formula/gnupg>
- <https://wiki.archlinux.org/index.php/GnuPG>

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

- public|secret key 불러오기: `gpg --import <key-file>`
- public key 내보내기: `gpg --armor --export <key-id|uid>`
- secret key 내보내기: `gpg --armor --export-secret-keys <key-id|uid>`
- public key 삭제: `gpg --delete-key <key-id|uid>`
- secret key 삭제: `gpg --delete-secret-key <key-id|uid>`

## 키 정보 변경

```sh
gpg --edit-key <key-id|uid>

# 도움말
gpg> help

# 비밀번호 변경
gpg> passwd
gpg> save
```

### 키 갱신 (유효기간 변경)

<https://unix.stackexchange.com/questions/552707/how-to-renew-an-expired-encryption-subkey-with-gpg>

```sh
gpg --edit-key <key-id|uid>

# 유효기간 변경
gpg> expire
gpg> <유효기간>
gpg> save

# subkey 선택: 1~n, 0(선택해제), *(모두선택)
gpg> key <number>
```

갱신 명령어 사용

```sh
gpg --quick-set-expire <key-id> <period> [subkeys]

# key, subkey 일괄갱신
# gpg --quick-set-expire <key-id> 2y && gpg --quick-set-expire <key-id> 2y "*"
```

## 관련 파일 확장자

> <https://stackoverflow.com/questions/58929260/what-are-the-meaningful-differences-between-gpg-sig-asc>

- `.gpg`: GNU Privacy Guard public keyring file, binary format
- `.sig`: GPG signed document file, binary format
- `.asc`: ASCII-armored signature with or without wrapped document, plain text format

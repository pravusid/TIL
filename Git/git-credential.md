# Git Credential

## 유형

<https://git-scm.com/book/en/v2/Git-Tools-Credential-Storage>

HTTP 프로토콜을 사용하는 경우 매전 인증정보를 입력해야 하지만, Git은 인증정보를 저장해둘 수 있다.

- 아무런 설정도 하지 않으면 어떤 암호도 저장하지 않는다

- `cache` 모드
  - 일정 시간 동안 메모리에 사용자이름과 암호 같은 인증정보를 기억한다
  - 이 정보를 Disk에 저장하지는 않으며 메모리에서도 15분 까지만 유지한다

- `store` 모드
  - 인증정보를 Disk의 텍스트 파일로 저장하며 계속 유지한다
  - 인증정보가 사용자 홈 디렉토리 아래에 일반 텍스트 파일로 저장된다

- `osxkeychain` 모드
  - Mac을 사용하는 경우 Mac에서 제공하는 Keychain 시스템에 사용자이름과 암호를 현재 로그인 계정 저장한다

- `Git Credential Manager for Windows`
  - Windows 환경의 인증 Helper이다
  - `osxkeychain`과 비슷하게 Windows Credential Store를 사용하여 안전하게 인증정보를 저장한다
  - <https://github.com/Microsoft/Git-Credential-Manager-for-Windows>

## 설정

`git config --global credential.helper <MODE>`

`git config --global credential.helper 'store --file ~/.my-credentials'`

Helper를 여러개 사용할 수도 있다.
인증시에는 Helper를 순차적으로 사용하며, 저장할 때는 모든 Helper에 저장한다.

```conf
[credential]
    helper = store --file /mnt/thumbdrive/.git-credentials
    helper = cache --timeout 30000
```

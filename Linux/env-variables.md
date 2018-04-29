# 리눅스 환경변수

## ~/.profile 환경변수 설정 (shell에서만 적용) : bin연결에 사용

  ```sh
  export PATH="$PATH:$APP_HOME:$APP_HOME/bin"
  export APP_HOME="/APP경로"
  ```

zsh를 사용중이면 `.zshenv` 파일에 환경변수를 설정해도 됨 (.profile로 안된다면)

export를 붙이면 shell variable이 아니라 environment variable로 사용하겠다는 의미

## 리눅스 자바 버전 변경

- `sudo update-alternatives --config java`

## alias 설정

shell 설정 파일 (.bashrc / .zshrc)에 다음 내용을 추가한다

```sh
# set alias
alias altjava='sudo update-alternatives --config java'
alias altterm='sudo update-alternatives --config x-terminal-emulator'
alias python='/usr/bin/python3'
alias pip='/usr/bin/pip3'
```

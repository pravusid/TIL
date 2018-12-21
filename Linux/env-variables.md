# 리눅스 환경변수

`env` 명령어로 출력 및 관리 할 수 있다

## 계층

- Single user mode
  - /etc/environment

- Multi user mode
  - login shell
    - /etc/profile
    - ~/.bash_profile: bash쉘을 통해 사용자의 로그인 세션이 열릴때 호출된다
    - ~/.profile: 사용자의 로그인 세션이 열릴 때 호출된다
  - non login shell
    - /etc/bash.bashrc
    - ~/.bashrc: 사용자에게만 적용되고, 리눅스 기본 쉘인 bash 쉘 세션이 생성될 때마다 로드된다
    - ~/.zshrc: 사용자에게만 적용되고, zsh쉘 세션이 생성될 때 마다 로드된다
    - 세션 환경변수는 현재 쉘에서 지정된 값으로 `set/unset 변수=값`을 활용해서 지정/해제한다

## 예시: ~/.profile 환경변수 설정 (shell에서만 적용)

  ```sh
  export PATH="$PATH:$APP_HOME:$APP_HOME/bin"
  export APP_HOME="/APP경로"
  ```

zsh를 사용중이면 `.zshenv` 파일에 환경변수를 설정해도 됨 (.profile로 안된다면)

export를 붙이면 shell variable이 아니라 environment variable로 사용하겠다는 의미

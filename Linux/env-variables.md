# 리눅스 환경변수 & SHELL

`env` 명령어로 출력 및 관리 할 수 있다

## 계층

### Single user mode

- `/etc/environment`

### Multi user mode

- login shell
  - `/etc/profile`
  - `~/.bash_profile`: bash쉘을 통해 사용자의 로그인 세션이 열릴때 호출된다
  - `~/.profile`: 사용자의 로그인 세션이 열릴 때 호출된다

- non login shell
  - `/etc/bash.bashrc`
  - `~/.bashrc`: 사용자에게만 적용되고, 리눅스 기본 쉘인 bash 쉘 세션이 생성될 때마다 로드된다
  - `~/.zshrc`: 사용자에게만 적용되고, zsh쉘 세션이 생성될 때 마다 로드된다
  - 세션 환경변수는 현재 쉘에서 지정된 값으로 `set/unset 변수=값`을 활용해서 지정/해제한다

## 참고사항

### login shell / non login shell

`~/.profile` 파일은 login shell에 의해 로드된다. login shell은 text mode에서 로그인 하면 수행되는 최초 과정이다.
대부분의 리눅스에서는 기본 login shell은 `bash`이고 이는 `/etc/passwd`에서 확인할 수 있다.

login shell에서 `bash`는 `~/.bash_profile` 파일과 `~/.profile` 파일이 존재하면 읽는다.
반면, `zsh`는 `~/.zprofile` 파일만 읽는다. (이는 zsh 문법이 기본 bourne shell 계통과 완전한 호환성을 보장하지 않기 때문이다)

`/bin/sh`을 login shell로 하고 `~/.profile`에 `export SHELL=/bin/zsh` 코드를 포함한다면,
터미널을 열었을 때 터미널은 `zsh`를 실행할 것이다. (일부 터미널은 `$SHELL`을 따르지 않는다)
> 이경우 여전히 login shell은 `sh` 이다.

대부분의 설정에서 `~/.profile` 파일은 그래픽 디스플레이 매니저로 로그인 할 때 **X session startup scripts**에 의해서 로드된다.

### terminal emulator (gnome terminal...)

터미널 에뮬레이터를 시작하여 shell prompt를 얻거나 shell을 명시적으로 시작하면, 해당 shell은 **non login shell**이다.

`~/.profile` (또는 `~/.zprofile`) 파일은 로그인 했을 때 실행하는 명령들 이므로 non-login shell에서는 해당 파일을 읽지 않는다.

- `zsh`는 login shell이든 non-login shell이든 관계없이 모든 대화형 shell에서 `~/.zshrc` 파일을 읽는다
- `bash`는 login shell에서, `~/.bashrc` 파일을 읽지 않는다

일반적으로 `~/.profile` 파일은 환경변수 정의를 담고있으므로, 로그인 시 혹은 전체 세션동안 한번만 실행하려는 프로그램일 것이다.

`~/.zshrc` 파일은 모든 대화형 shell instance에서 실행되어야 하는 것을 포함해야 한다.
예를 들면, alias, 함수 정의, shell 옵선 설정, 자동완성 설정, 프롬프트 설정 키 바인딩 등의 설정이다.

### zsh

다음은 `zsh`와 관련한 환경변수 파일이다

#### `.zshenv`

> shell 실행시 항상 읽음

- 변경가능성이 있는 (혹은 자주 변하는) 환경변수에 적합
- *PATH*를 포함하는 것도 적합하다
  - 경로가 변경될 때마다 업데이트를 위해 전체 세션을 재시작하는 것은 불편하기 때문이다
  - 새로운 shell 인스턴스를 실행하면 (i.e. 터미널 에뮬레이터 재실행) 변경한 설정이 적용된다

이 파일은 단일 명령어 실행(command string option == `zsh -c`)에도 적용된다는 사실을 잊으면 안된다. (대화형 쉘이 아닌 경우 포함)
즉, 이 파일에 선언된 설정이나 alias등이 명령어에 영향을 줄 수 있고, 이를 염두에 두고 `.zshenv` 설정이나 단일 명령어 실행을 사용해야 한다.

#### `.zprofile`

> 로그인 시 읽음

- 툴과 관련한 환경변수
- 명령어 실행과 관련한 환경설정 (i.e. `export FZF_DEFAULT_COMMAND="fd --type f"`)

이 파일을 수정하면 새로운 login shell을 실행하여 변경한 설정이 적용된 shell을 사용할 수 있다: `exec zsh --login`

#### `.zshrc`

> **대화형** shell 실행시 항상 읽음

대화형 쉘에서 사용할 내용을 입력하는 것이 좋음

- prompt
- output coloring
- aliases
- key bindings
- command completion/suggestion/highlighting
- commands history management
- miscellaneous...

#### `.zlogin`

> 로그인 시 읽음

`.zprofile` 파일과 유사하나 `.zshrc` 파일 이후 읽는다.

#### `.zlogout`

> 로그인 쉘에서 로그아웃 시 읽음

## 예시: `~/.profile` 환경변수 설정

zsh를 사용중이면 `.zshenv` 파일에 환경변수를 설정해도 됨

`export`를 붙이면 shell variable이 아니라 **environment variable**로 사용하겠다는 의미

```sh
export PATH="$PATH:$APP_HOME:$APP_HOME/bin"
export APP_HOME="/APP경로"
```

## references

- <https://unix.stackexchange.com/questions/71253/what-should-shouldnt-go-in-zshenv-zshrc-zlogin-zprofile-zlogout>
- <https://superuser.com/questions/187639/zsh-not-hitting-profile>

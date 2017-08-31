# 리눅스 환경변수

## ~/.profile 환경변수 설정 (shell에서만 적용) : bin연결에 사용

  ```sh
  export PATH="$PATH:$APP_HOME:$APP_HOME/bin"
  export APP_HOME="/APP경로"
  ```

## 리눅스 자바 버전 변경

- `sudo update-alternatives --config java`
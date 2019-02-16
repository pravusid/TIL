# Dockerfile

Dockerfile을 이용해서 배포용 이미지 파일을 생성한다

## Dockerfile 기본 명령어

<https://docs.docker.com/engine/reference/builder/>

### FROM

```dockerfile
FROM <image> [AS <name>]
FROM <image>[:<tag>] [AS <name>]
FROM <image>[@<digest>] [AS <name>]
```

베이스 이미지를 반드시 지정해야 하며 기본값으로 쓸 수 있는 베이스 이미지는 [Docker hub](https://hub.docker.com/explore/)에서 확인가능

### COPY

```dockerfile
COPY [--chown=<user>:<group>] <src>... <dest>
COPY [--chown=<user>:<group>] ["<src>",... "<dest>"]
```

파일이나 디렉토리를 이미지로 복사. 일반적으로 소스를 복사함. `target`디렉토리가 없으면 자동생성

### ADD

```dockerfile
ADD [--chown=<user>:<group>] <src>... <dest>
ADD [--chown=<user>:<group>] ["<src>",... "<dest>"]
```

`COPY`명령어와 매우 유사하나 차이점이 있음.
`src`에 파일 대신 URL을 입력할 수 있고 `src`에 압축 파일을 입력하는 경우 자동으로 압축을 해제하면서 복사됨.

### RUN

```dockerfile
RUN <command>
RUN ["executable", "param1", "param2"]
RUN executable param1 param2
```

명령어를 그대로 실행함. `/bin/sh -c` 뒤에 명령어를 실행하는 방식

### CMD

```dockerfile
CMD ["executable","param1","param2"]
CMD executable param1 param2
```

도커 컨테이너가 실행되었을 때 작동하는 명령어. 빌드할 때는 실행되지 않으며 여러 개의 `CMD`가 존재할 경우 가장 마지막 `CMD`만 실행 됨. 한꺼번에 여러 개의 프로그램을 실행하고 싶은 경우에는 `run.sh`파일을 작성하여 데몬으로 실행 함.

### WORKDIR

```dockerfile
WORKDIR /path/to/workdir
```

각 명령어의 현재 디렉토리는 한 줄 한 줄마다 초기화되기 때문에  같은 디렉토리에서 계속 작업하기 위해서 `WORKDIR`을 사용함

### EXPOSE

```dockerfile
EXPOSE <port> [<port>/<protocol>...]
```

도커 컨테이너가 실행되었을 때 요청을 기다리고 있는(Listen) 포트. 여러개의 포트를 지정가능

### VOLUME

```dockerfile
VOLUME ["/data"]
```

컨테이너 외부의 파일시스템을 컨테이너 내부로 마운트 할 때 사용

### ENV

```dockerfile
ENV <key> <value>
ENV <key>=<value> ...
```

컨테이너에서 사용할 환경변수. 컨테이너를 실행할 때 `-e`옵션을 사용하면 기존 값을 오버라이딩

### ENTRYPOINT

```dockerfile
ENTRYPOINT ["executable", "param1", "param2"]
ENTRYPOINT command param1 param2
```

컨테이너가 시작되었을 때(docker run / docker start) 실행할 명령. Dockerfile에서 한 번만 지정할 수 있다.

### LABEL

이미지의 metadata를 명시

```dockerfile
LABEL "com.example.vendor"="ACME Incorporated"
LABEL com.example.label-with-value="foo"
LABEL version="1.0"
LABEL description="This text illustrates \
that label-values can span multiple lines."
```

### USER

```dockerfile
USER <user>[:<group>] or
USER <UID>[:<GID>]
```

이미지와 Dockerfile에 지정된 RUN/CMD/ENTRYPOINT 명령을 실행할 때의 유저와 유저그룹을 지정한다

> user의 primary group이 없으면 root group으로 실행됨

## 이미지 생성 및 사용

- 절대경로지정: `docker build -f /path/to/a/Dockerfile .`

- 이름:태그 지정: `--tag`(`-t`):
  - `docker build -t shykes/myapp .`
  - `docker build -t -t shykes/myapp:latest shykes/myapp:1.0.2 .`

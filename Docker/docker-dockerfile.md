# Dockerfile

Dockerfile을 이용해서 배포용 이미지 파일을 생성한다

## 공통사항

exec form(`["foo", "bar"]`)은 JSON 배열로 인식되므로, 반드시 double quotes를 사용해야 한다

exec form은 shell form과 달리 명령 쉘을 호출하지 않는다.

> shell을 동반한 실행을 원하면 shell form을 사용하거나 직접 쉘을 실행해야 한다.
> (예 : `CMD [ "sh", "-c", "echo $ HOME"]`)

shell form은 다음 명령의 subcommand로 작동한다

- linux: `/bin/sh -c`
- windows: `cmd /S /C`

## Dockerfile 기본 명령어

<https://docs.docker.com/engine/reference/builder/>

### FROM

```dockerfile
FROM <image> [AS <name>]
FROM <image>[:<tag>] [AS <name>]
FROM <image>[@<digest>] [AS <name>]
```

베이스 이미지를 반드시 지정해야 하며 기본값으로 쓸 수 있는 베이스 이미지는 [Docker hub](https://hub.docker.com/explore/)에서 확인가능

### RUN

```dockerfile
RUN <command> (shell form, the command is run in a shell, which by default is /bin/sh -c on Linux or cmd /S /C on Windows)
RUN ["executable", "param1", "param2"] (exec form)
```

RUN 명령은 현재 이미지 위에 새 레이어에서 명령을 실행한 뒤 결과를 커밋한다.
커밋 된 결과 이미지는 Dockerfile의 다음 단계에 사용된다.

명령을 실행하면서 레이어가 쌓이는 것은 VCS(version control system)와 비슷하다.
적층 레이어는 커밋 비용은 저렴하고, 이미지 히스토리의 어느 시점에서든 컨테이너를 만들 수 있다는 Docker의 핵심 개념을 따른다.

> `SHELL` 명령을 사용해서 기본 쉘을 변경할 수 있다.

backslash(`\`)를 사용하여 여러행에 걸쳐 명령을 정의할 수 있다.

```dockerfile
RUN /bin/bash -c 'source $HOME/.bashrc; \
echo $HOME'
```

exec form을 사용해서 다른 쉘에서 명령을 실행할 수 있다

```dockerfile
RUN ["/bin/bash", "-c", "echo hello"]
```

다음 빌드시, RUN 수행결과의 캐시는 무효화 되지 않으므로, 캐시를 사용하지 않으려면 `docker build --no-cache` 플래그를 사용해야 한다.

### CMD

```dockerfile
CMD ["executable","param1","param2"] (exec form, this is the preferred form)
CMD ["param1","param2"] (as default parameters to ENTRYPOINT)
CMD command param1 param2 (shell form)
```

CMD의 주요 목적은 실행 컨테이너에 대한 기본값을 제공하는 것이다.

CMD에 실행 파일 포함여부는 선택가능하다.
그러나, 실행 파일이 생략된 경우 ENTRYPOINT 명령을 반드시 지정해야 한다.
(이 경우 ENTRYPOINT에 파라미터를 전달하는 형식으로 사용하게 됨: case2)

> 사용자가 `docker run`에 인수를 지정하면 CMD에 지정된 기본값을 무시한다

exec form으로 명령을 실행할 때는 명령의 전체 경로가 제공되어야 함.

> 예시: `CMD [ "/ usr / bin / wc", "-help"]`

**빌드할 때는 실행되지 않으며** 여러 개의 `CMD`가 존재할 경우 가장 마지막 `CMD`만 실행 됨.

한꺼번에 여러 개의 프로그램을 실행하고 싶은 경우에는 `run.sh`파일을 작성하여 데몬으로 실행 함.

### LABEL

이미지의 metadata를 명시

```dockerfile
LABEL "com.example.vendor"="ACME Incorporated"
LABEL com.example.label-with-value="foo"
LABEL version="1.0"
LABEL description="This text illustrates \
that label-values can span multiple lines."
```

### EXPOSE

```dockerfile
EXPOSE <port> [<port>/<protocol>...]
```

도커 컨테이너가 실행되었을 때 요청을 기다리고 있는(Listen) 포트. 여러개의 포트를 지정가능

### ENV

```dockerfile
ENV <key> <value>
ENV <key>=<value> ...
```

컨테이너에서 사용할 환경변수. 컨테이너를 실행할 때 `-e`옵션을 사용하면 기존 값을 오버라이딩

### ADD

```dockerfile
ADD [--chown=<user>:<group>] <src>... <dest>
ADD [--chown=<user>:<group>] ["<src>",... "<dest>"] (this form is required for paths containing whitespace)
```

`COPY`명령어와 매우 유사하나 차이점이 있음.
`src`에 파일 대신 URL을 입력할 수 있고 `src`에 압축 파일을 입력하는 경우 자동으로 압축을 해제하면서 복사됨.

### COPY

```dockerfile
COPY [--chown=<user>:<group>] <src>... <dest>
COPY [--chown=<user>:<group>] ["<src>",... "<dest>"]
```

파일이나 디렉토리를 이미지로 복사. 일반적으로 소스를 복사함. `target`디렉토리가 없으면 자동생성

### ENTRYPOINT

```dockerfile
ENTRYPOINT ["executable", "param1", "param2"] (exec form, preferred)
ENTRYPOINT command param1 param2 (shell form)
```

`docker run <image>` 뒤에 따라오는 모든 명령행 인수들은 `ENTRYPOINT` exec form에 더해지며, `CMD`에서 정의한 인수(case2)를 대체한다.

`docker run --entrypoint` 플래그를 사용해서 `ENTRYPOINT` 명령을 무시할 수 있다.

shell form은 `CMD` 혹은 `run` 명령행 인수들이 사용되는 것을 방지하지만,
`ENTRYPOINT`가 `/bin/sh -c`의 subcommand로 실행되기 때문에 실행파일은 컨테이너의 PID 1이 아니게 되어 SIGTERM과 같은 Unix signal을 수신하지 못하게 된다.

ENTRYPOINT를 중복선언해도 DockerFile에서 마지막으로 선언한 하나만 적용된다.

#### exec form ENTRYPOINT 예제

```dockerfile
FROM ubuntu
ENTRYPOINT ["top", "-b"]
CMD ["-c"]
```

위의 dockerfile로 이미지를 생성

- 실행: `docker run -it --rm --name test top -H`
- 확인: `docker exec`
- 결과: `top -b -H`, PID == 1

### VOLUME

```dockerfile
VOLUME ["/data"]
```

컨테이너에서 사용할 volume을 지정한다.

> 예시: mysql docker 이미지의 볼륨은 `VOLUME /var/lib/mysql`이다

여러 개를 지정하면 컨테이너 실행시 여러개의 볼륨이 컨테이너에 마운트 된다.
마운트 되는 볼륨의 실제데이터는 호스트의 `/var/lib/docker/volumes` 경로에 임의의 이름으로 존재한다.

### USER

```dockerfile
USER <user>[:<group>] or
USER <UID>[:<GID>]
```

이미지와 Dockerfile에 지정된 RUN/CMD/ENTRYPOINT 명령을 실행할 때의 유저와 유저그룹을 지정한다

> user의 primary group이 없으면 root group으로 실행됨

### WORKDIR

```dockerfile
WORKDIR /path/to/workdir
```

각 명령어의 현재 디렉토리는 매 라인마다 초기화되기 때문에 같은 디렉토리에서 계속 작업하기 위해서 `WORKDIR`을 사용함

## 이미지 생성 및 사용

- 절대경로지정: `docker build -f /path/to/a/Dockerfile .`

- 이름:태그 지정: `--tag`(`-t`):
  - `docker build -t shykes/myapp .`
  - `docker build -t shykes/myapp:latest -t shykes/myapp:1.0.2 .`

# Dockerfile

Dockerfile을 이용해서 배포용 이미지 파일을 생성한다

## Dockerfile 기본 명령어

<https://docs.docker.com/engine/reference/builder/>

### FROM

```sh
FROM <image>:<tag>
FROM ubuntu:16.04
```

베이스 이미지를 반드시 지정해야 하며 기본값으로 쓸 수 있는 베이스 이미지는 [Docker hub](https://hub.docker.com/explore/)에서 확인가능

### COPY

```sh
COPY <src>... <dest>
COPY . /usr/src/app
```

파일이나 디렉토리를 이미지로 복사. 일반적으로 소스를 복사함. `target`디렉토리가 없으면 자동생성

### ADD

```sh
ADD <src>... <dest>
ADD . /usr/src/app
```

`COPY`명령어와 매우 유사하나 차이점이 있음.  `src`에 파일 대신 URL을 입력할 수 있고  `src`에 압축 파일을 입력하는 경우 자동으로 압축을 해제하면서 복사됨.

### RUN

```sh
RUN <command>
RUN ["executable", "param1", "param2"]
RUN bundle install
```

명령어를 그대로 실행함. `/bin/sh -c` 뒤에 명령어를 실행하는 방식

### CMD

```sh
CMD ["executable","param1","param2"]
CMD command param1 param2
CMD bundle exec ruby app.rb
```

도커 컨테이너가 실행되었을 때 작동하는 명령어. 빌드할 때는 실행되지 않으며 여러 개의 `CMD`가 존재할 경우 가장 마지막 `CMD`만 실행 됨. 한꺼번에 여러 개의 프로그램을 실행하고 싶은 경우에는 `run.sh`파일을 작성하여 데몬으로 실행 함.

### WORKDIR

```sh
WORKDIR /path/to/workdir
```

각 명령어의 현재 디렉토리는 한 줄 한 줄마다 초기화되기 때문에  같은 디렉토리에서 계속 작업하기 위해서 `WORKDIR`을 사용함

### EXPOSE

```sh
EXPOSE <port> [<port>...]
EXPOSE 4567
```

도커 컨테이너가 실행되었을 때 요청을 기다리고 있는(Listen) 포트. 여러개의 포트를 지정가능

### VOLUME

```sh
VOLUME ["/data"]
```

컨테이너 외부에 파일시스템을 마운트 할 때 사용

### ENV

```sh
ENV <key> <value>
ENV <key>=<value> ...
ENV DB_URL mysql
```

컨테이너에서 사용할 환경변수. 컨테이너를 실행할 때 `-e`옵션을 사용하면 기존 값을 오버라이딩

## 이미지 생성 및 사용

Traditionally, the Dockerfile is called Dockerfile and located in the root of the context.
You use the -f flag with docker build to point to a Dockerfile anywhere in your file system.

`$ docker build -f /path/to/a/Dockerfile .`

`$ docker build --tag my-image-name .`

`$ docker build -t shykes/myapp .`

To tag the image into multiple repositories after the build, add multiple -t parameters when you run the build command:

`$ docker build -t shykes/myapp:1.0.2 -t shykes/myapp:latest .`

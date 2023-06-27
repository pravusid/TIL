# Docker Compose

<https://docs.docker.com/compose/>

## 설치

- <https://github.com/pravusid/dotfiles#prerequisite>
- <https://docs.docker.com/compose/install/>

## Compose CLI

```sh
Define and run multi-container applications with Docker.

Usage:
  docker-compose [-f <arg>...] [options] [COMMAND] [ARGS...]
  docker-compose [COMMAND] -h|--help

Options:
  -f, --file FILE             Specify an alternate compose file
                              (default: docker-compose.yml)
                              (여러 파일을 지정할 수도 있음)
  -p, --project-name NAME     Specify an alternate project name
                              (default: directory name)
  --verbose                   Show more output
  --log-level LEVEL           Set log level (DEBUG, INFO, WARNING, ERROR, CRITICAL)
  --no-ansi                   Do not print ANSI control characters
  -v, --version               Print version and exit
  -H, --host HOST             Daemon socket to connect to

  --tls                       Use TLS; implied by --tlsverify
  --tlscacert CA_PATH         Trust certs signed only by this CA
  --tlscert CLIENT_CERT_PATH  Path to TLS certificate file
  --tlskey TLS_KEY_PATH       Path to TLS key file
  --tlsverify                 Use TLS and verify the remote
  --skip-hostname-check       Don't check the daemon's hostname against the
                              name specified in the client certificate
  --project-directory PATH    Specify an alternate working directory
                              (default: the path of the Compose file)
  --compatibility             If set, Compose will attempt to convert deploy
                              keys in v3 files to their non-Swarm equivalent

Commands:
  build         Services are built once and then tagged, by default as project_service.
  bundle        Generate a Distributed Application Bundle (DAB) from the Compose file.
  config        Validate and view the Compose file.
  create        Creates containers for a service.
  down          Stops containers and removes containers, networks, volumes, and images created by up.
  events        Stream container events for every container in the project.
  exec          Execute a command in a running container.
  help          Get help on a command.
  images        List images.
  kill          Forces running containers to stop by sending a SIGKILL signal.
  logs          View output from containers.
  pause         Pauses running containers of a service. They can be unpaused with docker-compose unpause.
  port          Print the public port for a port binding
  ps            List containers
  pull          Pull service images
  push          Push service images
  restart       Restarts all stopped and running services.
  rm            Removes stopped service containers.
  run           Runs a one-time command against a service.
  scale         Set number of containers for a service
  start         Starts existing containers for a service.
  stop          Stops running containers without removing them. They can be started again with docker-compose start.
  top           Display the running processes
  unpause       Unpause services
  up            Builds, (re)creates, starts, and attaches to containers for a service
  version       Show the Docker-Compose version information
```

## Configuration reference

compose 파일은 service, networks, volumes를 정의하는 yaml 파일이다.
compose 파일의 기본 경로는 `./docker-compose.yml`이다.

서비스 정의는 각 컨테이너가 시작할 때 서비스에 적용되는 설정이며, 이는 `docker container create` 명령 매개변수에 전달하는 것과 유사하다.
마찬가지로 네트워크, 볼륨 정의는 `docker network create` 및 `docker volumn create`에 매개변수를 전달하는 것과 유사하다.

`docker container create`와 마찬가지로 `CMD`, `EXPOSE`, `VOLUME`, `ENV`와 같은 Dockerfile에 지정된 옵션은
기본적으로 인정되므로 `docker-compose.yml`에서 다시 지정할 필요가 없다.

## Build 사용하는 경우

<https://docs.docker.com/compose/compose-file/build/>

## 실행 순서 처리

<https://docs.docker.com/compose/startup-order/>

## Compose File (v3) example

<https://docs.docker.com/compose/compose-file/>

```yml
version: "3.7"
services:

  redis:
    image: redis:alpine
    ports:
      - "6379"
    networks:
      - frontend
    deploy:
      replicas: 2
      update_config:
        parallelism: 2
        delay: 10s
      restart_policy:
        condition: on-failure

  db:
    image: postgres:9.4
    volumes:
      - db-data:/var/lib/postgresql/data
    networks:
      - backend
    deploy:
      placement:
        constraints: [node.role == manager]

  vote:
    image: dockersamples/examplevotingapp_vote:before
    ports:
      - "5000:80"
    networks:
      - frontend
    depends_on:
      - redis
    deploy:
      replicas: 2
      update_config:
        parallelism: 2
      restart_policy:
        condition: on-failure

  result:
    image: dockersamples/examplevotingapp_result:before
    ports:
      - "5001:80"
    networks:
      - backend
    depends_on:
      - db
    deploy:
      replicas: 1
      update_config:
        parallelism: 2
        delay: 10s
      restart_policy:
        condition: on-failure

  worker:
    image: dockersamples/examplevotingapp_worker
    networks:
      - frontend
      - backend
    deploy:
      mode: replicated
      replicas: 1
      labels: [APP=VOTING]
      restart_policy:
        condition: on-failure
        delay: 10s
        max_attempts: 3
        window: 120s
      placement:
        constraints: [node.role == manager]

  visualizer:
    image: dockersamples/visualizer:stable
    ports:
      - "8080:8080"
    stop_grace_period: 1m30s
    volumes:
      - "/var/run/docker.sock:/var/run/docker.sock"
    deploy:
      placement:
        constraints: [node.role == manager]

networks:
  frontend:
  backend:

volumes:
  db-data:
```

[[example-mysql]] & [[example-redis]]

```yml
version: '3'

services:
  mysql:
    image: mysql:8
    container_name: dco_mysql
    ports:
      - 3306:3306
    volumes:
      - ./initdb/:/docker-entrypoint-initdb.d/
    environment:
      - MYSQL_ALLOW_EMPTY_PASSWORD=true
      - MYSQL_USER=idpravus
      - MYSQL_PASSWORD=idpravus@mysql
      - LANG=C.UTF-8
    command:
      - --character-set-server=utf8mb4
      - --collation-server=utf8mb4_unicode_ci

  redis:
    image: redis:6
    container_name: dco_redis
    ports:
      - 6379:6379
    command:
      - --requirepass idpravus@redis
```

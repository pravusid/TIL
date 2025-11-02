# Docker Compose

<https://docs.docker.com/compose/>

## 설치

- <https://github.com/pravusid/dotfiles#prerequisite>
- <https://docs.docker.com/compose/install/>

## Compose CLI

```sh
Define and run multi-container applications with Docker.

Usage:  docker compose [OPTIONS] COMMAND

Define and run multi-container applications with Docker

Options:
      --all-resources              Include all resources, even those not used by services
      --ansi string                Control when to print ANSI control characters
                                   ("never"|"always"|"auto") (default "auto")
      --compatibility              Run compose in backward compatibility mode
      --dry-run                    Execute command in dry run mode
      --env-file stringArray       Specify an alternate environment file
  -f, --file stringArray           Compose configuration files
      --parallel int               Control max parallelism, -1 for unlimited (default -1)
      --profile stringArray        Specify a profile to enable
      --progress string            Set type of progress output (auto, tty, plain,
                                   json, quiet)
      --project-directory string   Specify an alternate working directory
                                   (default: the path of the, first specified,
                                   Compose file)
  -p, --project-name string        Project name

Management Commands:
  bridge      Convert compose files into another model

Commands:
  attach      Attach local standard input, output, and error streams to a service's running container
  build       Build or rebuild services
  commit      Create a new image from a service container's changes
  config      Parse, resolve and render compose file in canonical format
  cp          Copy files/folders between a service container and the local filesystem
  create      Creates containers for a service
  down        Stop and remove containers, networks
  events      Receive real time events from containers
  exec        Execute a command in a running container
  export      Export a service container's filesystem as a tar archive
  images      List images used by the created containers
  kill        Force stop service containers
  logs        View output from containers
  ls          List running compose projects
  pause       Pause services
  port        Print the public port for a port binding
  ps          List containers
  publish     Publish compose application
  pull        Pull service images
  push        Push service images
  restart     Restart service containers
  rm          Removes stopped service containers
  run         Run a one-off command on a service
  scale       Scale services
  start       Start services
  stats       Display a live stream of container(s) resource usage statistics
  stop        Stop services
  top         Display the running processes
  unpause     Unpause services
  up          Create and start containers
  version     Show the Docker Compose version information
  wait        Block until containers of all (or specified) services stop.
  watch       Watch build context for service and rebuild/refresh containers when files are updated
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

> The top-level version property is defined by the Compose Specification for backward compatibility.
> It is only informative and you'll receive a warning message that it is obsolete if used.

```yml
name: my_project_name

services:
  redis:
    image: redis:alpine
    ports:
      - '6379'
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
      - '5000:80'
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
      - '5001:80'
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
      - '8080:8080'
    stop_grace_period: 1m30s
    volumes:
      - '/var/run/docker.sock:/var/run/docker.sock'
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
name: my_project_name

services:
  mysql:
    image: mysql:8
    container_name: dco_mysql
    ports:
      - '3306:3306'
    volumes:
      - ./initdb/:/docker-entrypoint-initdb.d/
    environment:
      - MYSQL_ROOT_PASSWORD=root@mysql
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
      - '6379:6379'
    command:
      - --requirepass idpravus@redis
```

<https://www.baeldung.com/ops/docker-compose-mysql-connection-ready-wait>

```yml
services:
  web:
    image: alpine:latest
    depends_on:
      db:
        condition: service_healthy
  db:
    image: mysql:5.7
    ports:
      - "3306:3306"
    environment:
      MYSQL_ROOT_PASSWORD: root
    healthcheck:
      test: ["CMD", "mysqladmin", "ping", "-h", "localhost"]
      interval: 10s
      timeout: 5s
      retries: 5
```

## logging

<https://docs.docker.com/engine/logging/>

```yml
services:
  app:
    image: myapp:latest
    # ...
    logging:
      driver: json-file
      options:
        max-size: "10m"
        max-file: "3"
```

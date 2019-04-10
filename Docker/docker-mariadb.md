# Using MariaDB via docker

## 설치

`docker pull mariadb:latest`

## 설정 & 실행

run 명령어에서 환경 설정을 하고 컨테이너를 실행함

```sh
docker \
  run \
  # 백그라운드에서 작동
  --detach \
  # db를 컨테이너 밖으로 연결 / 현재 경로를 지정한다면 {데이터경로} == $(pwd)/data
  --volume <데이터경로>:/var/lib/mysql \
  # 관리자 비밀번호 (설정 or empty 택 1)
  --env MYSQL_ROOT_PASSWORD=<MYSQL_ROOT_PASSWORD> \
  --env MYSQL_ALLOW_EMPTY_PASSWORD=true \
  # user, password, database
  --env MYSQL_USER=<MYSQL_USER> \
  --env MYSQL_PASSWORD=<MYSQL_PASSWORD> \
  --env MYSQL_DATABASE=<MYSQL_DATABASE> \
  # 컨테이너 이름과 포트연결
  --name <MYSQL_CONTAINER_NAME> \
  --publish 3306:3306 \
  # 실행할 이미지
  mariadb:latest \
  # 인코딩 설정 (args)
  --character-set-server=utf8mb4 \
  --collation-server=utf8mb4_unicode_ci;
```

bash shell로 container 실행

`docker exec -it <CONTAINER_NAME> /bin/bash`

명령어로 직접 실행

`docker exec -it <CONTAINER_NAME> mysql -uroot -p`

## Dockerfile

```dockerfile
FROM mariadb:latest

ENV MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD}
ENV MYSQL_USER=${MYSQL_USER}
ENV MYSQL_PASSWORD=${MYSQL_PASSWORD}
ENV MYSQL_DATABASE=${MYSQL_DATABASE}

CMD ["mysqld", "--character-set-server=utf8mb4", "--collation-server=utf8mb4_unicode_ci"]

ADD script.sql /docker-entrypoint-initdb.d/
RUN chown -R mysql:mysql /docker-entrypoint-initdb.d/

EXPOSE 3306
```

- prefix MYSQL_ 환경변수는 하나씩만 입력가능(배열... 불가)
- `entrypoint-initdb.d`에 넣은 초기화 스크립트는 도커 인스턴스 최초 실행시 파일이름 오름차순으로 구동됨

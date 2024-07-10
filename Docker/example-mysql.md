# Docker Example: MySQL(MariaDB)

- <https://hub.docker.com/_/mysql>
- <https://hub.docker.com/_/mariadb>

## 설치

`docker pull mysql:latest`

## 설정 & 실행

run 명령어에서 환경 설정을 하고 컨테이너를 실행함

```sh
docker \
  run \
  # 백그라운드에서 작동
  --detach \
  # db를 컨테이너 밖으로 연결 / 현재 경로를 지정한다면 {데이터경로} == $(pwd)/data
  --volume <데이터경로>:/var/lib/mysql \
  # root host 설정 (기본값 '%')
  --env MYSQL_ROOT_HOST=localhost \
  # root 비밀번호 (설정 or random or empty 택 1)
  --env MYSQL_ROOT_PASSWORD=<MYSQL_ROOT_PASSWORD> \
  --env MYSQL_RANDOM_ROOT_PASSWORD=true \
  --env MYSQL_ALLOW_EMPTY_PASSWORD=true \
  # user, password, database
  --env MYSQL_USER=<MYSQL_USER> \
  --env MYSQL_PASSWORD=<MYSQL_PASSWORD> \
  --env MYSQL_DATABASE=<MYSQL_DATABASE> \
  # 컨테이너 이름과 포트연결
  --name <MYSQL_CONTAINER_NAME> \
  --publish 3306:3306 \
  # 실행할 이미지
  mysql:latest \
  # 인코딩 설정 (args)
  --character-set-server=utf8mb4 \
  --collation-server=utf8mb4_unicode_ci
```

bash shell로 container 실행

`docker exec -it <CONTAINER_NAME> /bin/bash`

명령어로 직접 실행

`docker exec -it <CONTAINER_NAME> mysql -uroot -p`

run 예시

```sh
#!/bin/bash

docker stop idpravus
docker rm idpravus

docker \
  run \
  -d \
  --name idpravus \
  -p 5000:3306 \
  -v $(pwd)/scripts/:/docker-entrypoint-initdb.d/ \
  -e MYSQL_ALLOW_EMPTY_PASSWORD=true \
  -e MYSQL_USER=idpravus \
  -e MYSQL_PASSWORD=idpravus@test \
  -e LANG=C.UTF-8 \
  mysql:5.7 \
  --character-set-server=utf8mb4 \
  --collation-server=utf8mb4_unicode_ci \
  --max_allowed_packet=16M

docker logs -f idpravus
```

초기화 스크립트 `scripts/_init.sql` 예시

> 초기화를 위한 `docker-entrypoint-initdb.d` 호출 코드: <https://github.com/docker-library/mysql/blob/8a3178fd2f84b610693bd4ba9d6bdf26215b04b8/8.0/docker-entrypoint.sh#L399>

```sql
CREATE DATABASE idpravus_db;

GRANT all privileges on idpravus_db.* to 'idpravus'@'%' identified by 'passwd@idpravus';
flush privileges;
```

## Dockerfile

```dockerfile
FROM mysql:latest

ENV LANG=C.UTF-8
ENV MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD}
ENV MYSQL_USER=${MYSQL_USER}
ENV MYSQL_PASSWORD=${MYSQL_PASSWORD}
ENV MYSQL_DATABASE=${MYSQL_DATABASE}

CMD ["mysqld", "--character-set-server=utf8mb4", "--collation-server=utf8mb4_unicode_ci"]
# CMD ["mysqld", "--sql-mode=STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION"]

ADD script.sql /docker-entrypoint-initdb.d/
RUN chown -R mysql:mysql /docker-entrypoint-initdb.d/

EXPOSE 3306
```

- prefix MYSQL_ 환경변수는 하나씩만 입력가능(배열... 불가)
- `entrypoint-initdb.d`에 넣은 초기화 스크립트는 도커 인스턴스 최초 실행시 파일이름 오름차순으로 구동됨

## Troubleshooting

### UTF8 Encoded SQL Scripts in initdb

<https://github.com/docker-library/mysql/issues/131>

`docker-entrypoint`를 사용해 UTF8 인코딩으로 되어있는 sql 스크립트 실행시켜 DB 초기화를 하려는 경우
shell에서 스크립트를 실행시키는데 docker 인스턴스의 기본 locale codeset은 UTF가 아니다

해결책1: `locale` 변경

```Dockerfile
RUN apt-get update && apt-get install -y locales && rm -rf /var/lib/apt/lists/* $ \
  && localedef -i en_US -c -f UTF-8 -A /usr/share/locale/locale.alias en_US.UTF-8

ENV LANG=C.UTF-8
```

해결책2: `/etc/mysql/conf.d/utf8.cnf` 변경

client -> `default-character-set = utf8` 설정으로 해결 가능

```conf
[mysqld]
init_connect=‘SET collation_connection = utf8_unicode_ci’
character-set-server = utf8
collation-server = utf8_unicode_ci

[client]
default-character-set = utf8
```

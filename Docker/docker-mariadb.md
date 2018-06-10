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
  # db를 컨테이너 밖으로 연결
  --volume [데이터경로]:/var/lib/mysql \
  # 관리자 비밀번호 (설정 or empty 택 1)
  --env MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD} \
  --env MYSQL_ALLOW_EMPTY_PASSWORD=true \
  # user, password, database
  --env MYSQL_USER=${MYSQL_USER} \
  --env MYSQL_PASSWORD=${MYSQL_PASSWORD} \
  --env MYSQL_DATABASE=${MYSQL_DATABASE} \
  # 컨테이너 이름과 포트연결
  --name ${MYSQL_CONTAINER_NAME} \
  --publish 3306:3306 \
  # 인코딩 설정
  --character-set-server=utf8mb4
  --collation-server=utf8mb4_unicode_ci
  # 실행할 이미지
  mariadb:latest;
```

bash shell로 container 실행

`docker exec -it [CONTAINER_NAME] /bin/bash`

명령어로 직접 실행

`docker exec -it [CONTAINER_NAME] mysql -uroot -p`
# Docker 명령어

`docker version` : 버전확인

## docker run

`docker run [OPTIONS] IMAGE[:TAG|@DIGEST][COMMAND] [ARG...]`

| 옵션  | 설명                                                   |
| ----- | ------------------------------------------------------ |
| -d    | detached mode 흔히 말하는 백그라운드 모드              |
| -p    | 호스트와 컨테이너의 포트를 연결 (포워딩)               |
| -v    | 호스트와 컨테이너의 디렉토리를 연결 (마운트)           |
| -e    | 컨테이너 내에서 사용할 환경변수 설정                   |
| –name | 컨테이너 이름 설정                                     |
| –rm   | 프로세스 종료시 컨테이너 자동 제거                     |
| -it   | -i와 -t를 동시에 사용한 것으로 터미널 입력을 위한 옵션 |
| –link | 컨테이너 연결 [컨테이너명:별칭]                        |

### bash shell로 Ubuntu 이미지 시작

`docker run --rm -it ubuntu:16.04 /bin/bash`

### -v 옵션으로 데이터 볼륨 지정

```sh
docker run -d -p 3306:3306 \
  -e MYSQL_ALLOW_EMPTY_PASSWORD=true \
  --name mysql \
  -v /my/own/datadir:/var/lib/mysql \ # <- volume mount
  mysql:5.7
```

## docker ps

실행중인 컨테이너 목록을 출력

| 옵션      | 설명           |
| --------- | -------------- |
| -a(--all) | 전체 목록 출력 |

## docker start

`docker start [OPTIONS] CONTAINER [CONTAINER...]`

## docker stop

`docker stop [OPTIONS] CONTAINER [CONTAINER...]`

## docker rm

`docker rm [OPTIONS] CONTAINER [CONTAINER...]`

중지된 컨테이너 모두 삭제 : `docker rm -v $(docker ps -a -q -f status=exited)`

## docker images

다운로드한 이미지 목록 출력

`docker images [OPTIONS][REPOSITORY[:TAG]]`

## docker pull

이미지 다운로드

`docker pull [OPTIONS] NAME[:TAG|@DIGEST]`

## docker rmi

다운로드 한 이미지 삭제

`docker rmi [OPTIONS] IMAGE [IMAGE...]`

## docker logs

`docker logs [OPTIONS] CONTAINER`

| 옵션                 | 설명                             |
| -------------------- | -------------------------------- |
| -f `${CONTAINER_ID}` | 해당 컨테이너의 실시간 로그 출력 |
| --tail n             | 마지막 n줄의 로그 출력           |

## docker exec

컨테이너 내의 파일을 실행하려고 할 때

`docker exec [OPTIONS] CONTAINER COMMAND [ARG...]`

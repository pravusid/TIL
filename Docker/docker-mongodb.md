# Docker Mongodb

## 설치

`docker pull mongo`

```sh
docker \
  run \
  -d \
  -v /home/idpravus/docker/mongo:/data/db \
  -p 27017:27017 \
  --env MONGO_INITDB_ROOT_USERNAME=admin \
  --env MONGO_INITDB_ROOT_PASSWORD=4321 \
  --name mongodb \
  mongo:latest \
  mongod --auth
```

bash shell로 컨테이너 실행

`docker exec -it mongodb /bin/bash`

# Docker Example: Mongodb

## 설치

`docker pull mongo:latest`

```sh
docker \
  run \
  -d \
  -v /home/idpravus/docker/mongo:/data/db \
  -p 27017:27017 \
  --env MONGO_INITDB_ROOT_USERNAME=root \
  --env MONGO_INITDB_ROOT_PASSWORD=password \
  --name mongodb \
  mongo:latest \
  mongod --auth
```

bash shell로 컨테이너 실행

`docker exec -it mongodb /bin/bash`

## Dockerfile example

```dockerfile
FROM mongo:latest

EXPOSE 27017

ENV MONGO_INITDB_ROOT_USERNAME=root
ENV MONGO_INITDB_ROOT_PASSWORD=password

COPY mongo-config.js /docker-entrypoint-initdb.d/
```

초기화시 `js`나 `sh`로 실행할 내용을 지정할 수 있다.

```js
let error = true

let res = [
db.idpravus.createUser({ user: "testUser", pwd: "test", roles: ["readWrite", "dbAdmin"] })
]

printjson(res)

if (error) {
print('Error, exiting')
quit(1)
}
```

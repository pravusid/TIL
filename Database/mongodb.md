# MONGODB

## 설치

- Mongodb GUI : <https://robomongo.org>
- 몽고DB-서버 설치 : `sudo apt install mongodb-server`

## 설정

- ip 바인딩 설정
  - /etc/mongodb.conf : `bind_ip = 0.0.0.0`
  - /etc/ : `sudo rm -rf *.sock`

## 실행/종료

- 서버 실행 : `mongod --dbpath /home/sist/springDev/mongodb/data`
- 서버 종료

  ```sh
  mongo
  use admin
  db.shutdownServer()
  ```

## 사용자 설정

### 관리자 계정 및 권한 추가

```sh
use admin
db.createUser( {
    user: "admin",
    pwd: "1234",
    roles: [ "userAdminAnyDatabase" ] } )

use admin
db.createUser( {
    user: "dbadmin",
    pwd: "1234",
    roles: ["readWriteNayDatabase", "dbAdminAnyDatabase", "clusterAdmin"] } )
```

### 사용자 관리

```sh
use dbname
db.createUser({
  user: "testUser",
  pwd: "test",
  roles: ["readWrite", "dbAdmin"]
})

db.dropUser("<username>")
```

### 다른 db에 권한을 가진 사용자

다음 명령은 read 권한만 갖고 있는 동일한 사용자를 admin 데이터베이스에 추가하고 testDB 데이터베이스에 대한 readWrite 권한을 부여한다.

```sh
use admin
db.createUser( {user: "testUser",
    userSource: "test",
    roles: ["read"],
    otherDBRoles:{ testDB: ["readWrite"] } } )
```

## 인증모드 활성화

`--auth` 파라미터로 인증모드를 활성화 한다

`mongod -dbpath D:\mongodb\data --auth`

### 관리자로 로그인

mongo shell에 접속한 이후

```sh
use admin
db.auth("useradmin", "test")

show users
```

mongo shell에 접속하기 전

`mongo admin --username "useradmin" --password "test"`

## 기본 명령어

- db생성 : `use db이름`
- table생성 : `db.테이블이름`
- insert : `db.member.insert({no:1,name:"hong",age:10})`
- select : `db.member.find()`
- drop : `db.member.drop()`
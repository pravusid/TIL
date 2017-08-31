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

## 기본 명령어

- db생성 : `use db이름`
- table생성 : `db.테이블이름`
- insert : `db.member.insert({no:1,name:"hong",age:10})`
- select : `db.member.find()`
- drop : `db.member.drop()`
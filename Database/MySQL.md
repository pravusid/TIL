# MySQL / MariaDB

## DB 정의

### user 조회

```sql
SELECT host, user, password from mysql.user;
```

### db, user 생성

```sql
-- UTF8로 DB생성
CREATE DATABASE `dbname` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

-- 유저 생성
CREATE USER 'userid'@'%' IDENTIFIED BY 'password';
INSERT INTO mysql.user VALUES('%', 'userid', PASSWORD('password'));

-- 유저에게 특정 DB 권한 주기
GRANT all privileges on dbname.* to userid@localhost identified by 'password';
GRANT all privileges on dbname.* to 'userid'@'%' identified by 'password';

-- 권한변경 반영
flush privileges;

-- 외부접속 권한 삭제
DELETE FROM mysql.user WHERE Host='%' AND User='root';
FLUSH PRIVILEGES;
```

## 테이블 인코딩 설정 생성

```sql
create table test(
  title varchar(20)
  ) default charset utf8;
```

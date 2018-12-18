# PostgreSQL

<http://postgresql.kr/docs/9.6/index.html>

## 설치

## 데이터베이스 생성

shell 상에서 다음명령어를 실행 (postgresql/bin 아래에 createdb, dropdb 존재)

데이터베이스 만들기 : `createdb mydb`

데이터베이스 삭제 : `dropdb mydb`

## psql 사용

pgAdmin이나 ODBC나 JDBC를 이용한 client를 사용할 수도 있음

mydb에서 대화형 터미널 프로그램 실행 : `psql mydb`

### 내장명령

`\h` : help

`\q` : quit

`\?` : 내장명령 확인



## 명령어

### 유저생성

`create user wwwi with password 'wwwi'`

### 데이터베이스  만들기

`create database wwwi`

### 자료형

|                               |                        |
| ----------------------------- | ---------------------- |
| smailint                      | 2 바이트 정수               |
| integer                       | 4 바이트 정수               |
| bigint                        | 8 바이트 정수               |
| decimal(a, a)/numeric(a, s)   | 10진수형                  |
| real                          | 6자리 단정도 부동소수점          |
| double precision              | 15 자리 배정도 부동소수점        |
| serial                        | 4 바이트 일련번호             |
| bigserial                     | 8 바이트 일련번호             |
| date                          | 일자                     |
| time                          | 시간                     |
| timestamp                     | 일자시간                   |
| char(문자수)/character           | 고정길이 문자열  (최대 4096 문자) |
| varchar(문자수)/charcter varying | 가변길이 문자열  (최대 4096 문자) |
| text                          | 무제한 텍스트                |
| Large Object                  | oid형                   |
| boolean/bool                  | true/false             |

### 테이블 만들기

```sql
create table test (
  key char(16) primary key,
  val1 integer,
  val2 integer,
);
```

#### 제약조건

```sql
create table test (
  key char(16) primary key,
  val1 integer not null,
  val2 integer unique,
  val3 integer default 0 not null
);

-- constraint으로 Primary Key를 설정하거나 복수의 Primary Key는 아래와 같은 방법으로 설정

create table test (
  key char(16),
  val1 integer,
  val2 integer,
  constraint PK_NAME primary key(key, val1)
);
```

### 테이블 삭제

`drop table testa`

### 인덱스 생성

```sql
create unique index PK_NAME on test (
  key,
  val1
);

create index PK_NAME on test (
  key,
  val1
);

```

### 인덱스 삭제

`drop index PK_NAME`

## 일렬번호 만들기

### serial 사용

```sql
create table test (
  key char(16),
  val1 serial,
  val2 integer,
);
```

> auto_increment 처럼 사용하는듯

### sequence 사용

```sql
create sequence seq
  increment 10
  minvalue 10
  maxvalue 1000000
  start 10
  cache 100
  cycle
;
```

- `nextval` : 마지막 sequence 다음 수
- `currval` : 현재 sequence

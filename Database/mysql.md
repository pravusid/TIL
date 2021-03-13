# MySQL

## Case Sensitivity

<https://dev.mysql.com/doc/refman/8.0/en/identifier-case-sensitivity.html>

### not case-sensitive

partition, subpartition, column, index, stored routine, event, and resource group names, column aliases

### case-sensitive

names of logfile groups

### by Platform

database, table, table aliases and trigger names

These are case-sensitive on Unix, but not so on Windows or macOS.

### `lower_case_table_names` system variable

On Unix, the default value of lower_case_table_names is 0. On Windows, the default value is 1. On macOS, the default value is 2.

> This behavior applies to database names, table names and table aliases

- 값 0: 이름은 생성시 지정된 대소문자로 저장되며 이름비교는 대소문자를 구분하여 실행함, case-sensitive인 경우만 적용됨, 대소문자를 구분하여 저장한 데이터가 있다면 이 옵션을 변경하지 않아야 함
- 값 1: 이름은 소문자로 디스크에 저장되며 이름 비교는 대소문자 구분없이 실행됨
- 값 2: 이름은 생성시 지정된 대소문자로 저장되며 이름 비교는 모두 소문자로 변환하여 실행함, 파일시스템이 not case-sensitive인 경우만 적용됨

## 시스템 설정

서버 시스템 변수 확인

- `HOW [GLOBAL | SESSION] VARIABLES`
- <https://dev.mysql.com/doc/refman/8.0/en/show-variables.html>

서버 시스템 변수 정보

<https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html>

설정파일

<https://dev.mysql.com/doc/refman/8.0/en/option-files.html>

### 변수 설정

<https://dev.mysql.com/doc/refman/8.0/en/program-variables.html>

명령행

```sh
mysql --max_allowed_packet=16777216
mysql --max_allowed_packet=16M
```

설정파일

```conf
[mysql]
max_allowed_packet=16777216
max_allowed_packet=16M
```

런타임 설정

```sql
mysql> SET GLOBAL max_allowed_packet=16M; -- 오류발생
mysql> SET GLOBAL max_allowed_packet=16*1024*1024;
```

### Packet Too Large

- <https://dev.mysql.com/doc/refman/8.0/en/packet-too-large.html>
- <https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_max_allowed_packet>

최대 1GB (MySQL 8.0 server)

`max_allowed_packet` 기본 크기

- (>= 8.0.3) 64MB
- (<= 8.0.2) 4MB
- (< 5.6) 1MB

### Slow query

- By default, the slow query log is disabled.
- `--slow_query_log[={0|1}]`
- `--slow_query_log_file=/path/to/file.log`

## SYSTEM

<https://dev.mysql.com/doc/refman/8.0/en/sql-syntax-server-administration.html>

### character set 확인

```sql
-- 시스템변수에서 인코딩 설정 확인
show variables like 'char%';

-- 데이터베이스 인코딩 설정 확인
SELECT S.SCHEMA_NAME, default_character_set_name
FROM information_schema.SCHEMATA S;

-- 테이블 인코딩 설정 확인
SELECT
  T.TABLE_NAME, CCSA.character_set_name, T.TABLE_COLLATION
FROM
  information_schema.`TABLES` T,
  information_schema.`COLLATION_CHARACTER_SET_APPLICABILITY` CCSA
WHERE
  CCSA.collation_name = T.table_collation and T.TABLE_SCHEMA = '<데이터베이스이름>';
```

### db 조회

```sql
show databases;
```

### user 조회

```sql
SELECT host, user from mysql.user;
```

### db, user 생성

```sql
-- UTF8로 DB생성
CREATE DATABASE `dbname` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- 유저 생성
CREATE USER 'userid'@'%' IDENTIFIED BY 'password';
INSERT INTO mysql.user VALUES('%', 'userid', PASSWORD('password'));
```

### 권한

권한확인

```sql
SHOW GRANTS FOR CURRENT_USER;
```

권한 부여

```sql
-- 유저에게 특정 DB 권한 주기
GRANT {권한} on {db|*}.{table|*} to 'user'@'localhost' identified by 'password';
GRANT {권한} on {db|*}.{table|*} to 'user'@'%' identified by 'password';
-- 권한변경 반영
FLUSH PRIVILEGES;
```

권한 제거

```sql
-- 특정권한 제거
REVOKE {권한} on {db|*}.{table|*} FROM 'some_user'@'some_host';
-- 전체권한 제거
REVOKE ALL PRIVILEGES, GRANT OPTION FROM 'some_user'@'some_host';
FLUSH PRIVILEGES;
```

HOST 변경

```sql
UPDATE mysql.user SET host = '172.31.%' WHERE host = 'localhost' AND user = 'username';
UPDATE mysql.db SET host = '172.31.%' WHERE host = 'localhost' AND user = 'username';
FLUSH PRIVILEGES;
```

외부접속 권한 삭제

```sql
DELETE FROM mysql.user WHERE Host='%' AND User='user';
FLUSH PRIVILEGES;
```

#### 운영용 계정 권한 제한

최소 권한만 부여하는 것이 좋음

```sql
GRANT SELECT, INSERT, UPDATE, DELETE, LOCK TABLES, CREATE TEMPORARY TABLES ON *.* TO 'some_user'@'some_host';
FLUSH PRIVILEGES;
```

#### 관리용 계정 권한 제한

운영용 계정 + 추가권한

```sql
GRANT ALTER, CREATE, DROP, INDEX, REFERENCES ON *.* TO 'some_user'@'some_host';
FLUSH PRIVILEGES;
```

#### 권한 종류

<https://dev.mysql.com/doc/refman/8.0/en/privileges-provided.html>

| Privilege               | Grant Table Column           | Context                               |
| ----------------------- | ---------------------------- | ------------------------------------- |
| ALL [PRIVILEGES]        | Synonym for “all privileges” | Server administration                 |
| ALTER                   | Alter_priv                   | Tables                                |
| ALTER ROUTINE           | Alter_routine_priv           | Stored routines                       |
| CREATE                  | Create_priv                  | Databases, tables, or indexes         |
| CREATE ROUTINE          | Create_routine_priv          | Stored routines                       |
| CREATE TABLESPACE       | Create_tablespace_priv       | Server administration                 |
| CREATE TEMPORARY TABLES | Create_tmp_table_priv        | Tables                                |
| CREATE USER             | Create_user_priv             | Server administration                 |
| CREATE VIEW             | Create_view_priv             | Views                                 |
| DELETE                  | Delete_priv                  | Tables                                |
| DROP                    | Drop_priv                    | Databases, tables, or views           |
| EVENT                   | Event_priv                   | Databases                             |
| EXECUTE                 | Execute_priv                 | Stored routines                       |
| FILE                    | File_priv                    | File access on server host            |
| GRANT OPTION            | Grant_priv                   | Databases, tables, or stored routines |
| INDEX                   | Index_priv                   | Tables                                |
| INSERT                  | Insert_priv                  | Tables or columns                     |
| LOCK TABLES             | Lock_tables_priv             | Databases                             |
| PROCESS                 | Process_priv                 | Server administration                 |
| PROXY                   | See proxies_priv table       | Server administration                 |
| REFERENCES              | References_priv              | Databases or tables                   |
| RELOAD                  | Reload_priv                  | Server administration                 |
| REPLICATION CLIENT      | Repl_client_priv             | Server administration                 |
| REPLICATION SLAVE       | Repl_slave_priv              | Server administration                 |
| SELECT                  | Select_priv                  | Tables or columns                     |
| SHOW DATABASES          | Show_db_priv                 | Server administration                 |
| SHOW VIEW               | Show_view_priv               | Views                                 |
| SHUTDOWN                | Shutdown_priv                | Server administration                 |
| SUPER                   | Super_priv                   | Server administration                 |
| TRIGGER                 | Trigger_priv                 | Tables                                |
| UPDATE                  | Update_priv                  | Tables or columns                     |
| USAGE                   | Synonym for “no privileges”  | Server administration                 |

## TIMEZONE

<https://dev.mysql.com/doc/refman/8.0/en/time-zone-support.html>

`default-time-zone='timezone'` 서버 타임존을 지정할 수 있다.

세션 타임존은 기본적으로 서버 타임존을 가져오지만 다음의 명령으로 세션 시간대를 설정할 수 있다: `SET time_zone = timezone;`

전역 시간대와 세션 시간대는 다음으로 검색할 수 있다: `SELECT @@GLOBAL.time_zone, @@SESSION.time_zone;`

### 시각/날짜 처리

MySQL은 TIMESTAMP 값을 현재 시간대에서 UTC로 저장하고, 저장된 UTC값을 현재 시간대로 변환하여 출력한다.
(DATETIME: `YYYY-MM-DD hh:mm:ss[.fraction]`, DATE, TIME은 변환과정이 없다)

기본적으로 각 연결의 시간대는 서버의 설정값이지만, 시간대는 연결별로 설정할 수 있다.

- DATETIME: 8 bytes(< 5.6.4) -> 5 bytes + fractional seconds storage
- TIMESTAMP: 4 bytes(< 5.6.4) -> 4 bytes + fractional seconds storage

<https://dev.mysql.com/doc/refman/8.0/en/date-and-time-functions.html#function_date-format>

#### 경우의 수

- 고정 TimeZone

  - local datetime을 (10:00+09:00) DATETIME으로 저장하고 +09 timezone으로 연결
  - UTC datetime을 (01:00Z) DATETIME으로 저장하고 +00 timezone으로 연결

- 변동 TimeZone

  - local datetime값을 가지고 있는 DATETIME을 `CONVERT_TZ` 사용하여 변환하고 timezone은 필요한 대로
  - UTC datetime을 (01:00Z) TIMESTAMP로 저장하고 timezone은 필요한 대로

## DDL

<https://dev.mysql.com/doc/refman/8.0/en/sql-syntax-data-definition.html>

### CREATE 테이블 인코딩 설정 생성

```sql
CREATE TABLE test(
  title varchar(20)
) DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

### 테이블 복사

```sql
-- 테이블 생성하여 복사
CREATE TABLE <TARGET> DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
SELECT * FROM <SOURCE> [WHERE <CONDITION>];

-- 생성한 테이블에 복사
INSERT INTO <TARGET> SELECT * FROM <SOURCE> [WHERE <CONDITION>];

-- 특정 컬럼만 복사
-- target, source 테이블 컬럼의 순서가 매우 중요함! -> 해당 idx 데이터가 입력되는 것임
INSERT INTO <TARGET_TABLE (col1, col2, ...)>
SELECT <col1, col2, ...> FROM <SOURCE> [WHERE <CONDITION>];
```

## ALTER TABLE

PK 변경

```sql
ALTER TABLE `my_table` DROP PRIMARY KEY, ADD PRIMARY KEY (`id`);
-- 또는
ALTER TABLE `my_table` DROP KEY `key_name`, ADD PRIMARY KEY (`id`);
```

COLUMN 변경

```sql
ALTER TABLE <TABLE> ADD <COLUMN_NAME> <COLUMN_TYPE> [CONSTRAINTS] [COMMENT '설명'] [FIRST | AFTER <COLUMN_NAME>];
ALTER TABLE <TABLE> MODIFY COLUMN <COLUMN_NAME> <COLUMN_TYPE> [CONSTRAINTS] [COMMENT '설명'];
ALTER TABLE <TABLE> RENAME COLUMN <COLUMN_NAME> TO <NEW_COLUMN_NAME>;
ALTER TABLE <TABLE> DROP COLUMN <COLUMN_NAME>;
```

COLUMN 순서만 변경

```sql
-- 다른COLUMN 다음으로 이동
ALTER TABLE <테이블명> MODIFY COLUMN <컬럼명> <자료형> [CONSTRAINTS] [COMMENT '설명'] AFTER <다른컬럼이름>;
-- 첫번째 위치로 이동
ALTER TABLE <테이블명> MODIFY COLUMN <컬럼명> <자료형> [CONSTRAINTS] [COMMENT '설명'] FIRST;
```

## DML

<https://dev.mysql.com/doc/refman/8.0/en/sql-data-manipulation-statements.html>

### SELECT

#### query cache

```sql
SET SESSION query_cache_type = OFF;
SHOW VARIABLES LIKE 'query_cache_type';

-- select 실행시 캐시 없이 할 수도 있음
SELECT SQL_NO_CACHE * FROM <TABLE_NAME>;
```

### INSERT SELECT

```sql
INSERT INTO tbl_temp2 (fld_id)
  SELECT tbl_temp1.fld_order_id
  FROM tbl_temp1
  WHERE tbl_temp1.fld_order_id > 100;
```

### transaction

```sql
START TRANSACTION [transaction_characteristic [, transaction_characteristic] ...]

BEGIN [WORK]
COMMIT [WORK] [AND [NO] CHAIN] [[NO] RELEASE]
ROLLBACK [WORK] [AND [NO] CHAIN] [[NO] RELEASE]
```

transaction_characteristic

- `WITH CONSISTENT SNAPSHOT`

  - The WITH CONSISTENT SNAPSHOT modifier starts a consistent read for storage engines that are capable of it.
  - This applies only to InnoDB. The effect is the same as issuing a START TRANSACTION followed by a SELECT from any InnoDB table.
  - The WITH CONSISTENT SNAPSHOT modifier does not change the current transaction isolation level
  - so it provides a consistent snapshot only if the current isolation level is one that permits a consistent read.
  - The only isolation level that permits a consistent read is REPEATABLE READ.
  - For all other isolation levels, the WITH CONSISTENT SNAPSHOT clause is ignored.
  - A warning is generated when the WITH CONSISTENT SNAPSHOT clause is ignored.

- access mode

  - `READ WRITE`
  - `READ ONLY`

auto commit 조회/수정

```sql
select @@AUTOCOMMIT;
SET AUTOCOMMIT=FALSE;
```

#### isolation level

- READ UNCOMMITTED
- READ COMMITTED
- REPEATABLE READ (기본값)
- SERIALIZABLE

isolation level 조회

```sql
SHOW VARIABLES WHERE VARIABLE_NAME='tx_isolation';
```

### UPDATE FROM SELECT

```sql
UPDATE
  tablename1 AS t1
  JOIN tablename2 AS t2 ON join_condition
SET assignment_list
[WHERE where_condition];
```

## 집계함수

<https://dev.mysql.com/doc/refman/8.0/en/group-by-functions.html>

숫자 인수의 경우 분산 및 표준편차 함수의 경우 DOUBLE 타입 값을 반환한다.

SUM(), AVG() 함수는 정확한 값의 인수(정수 또는 DECIMAL)에 대해서는 DECIMAL 타입 값을,
근사값 인수(FLOAT, DOUBLE)에 대해서는 DOUBLE 타입 값을 리턴한다.

## DB Dump / Import

### Dump

<https://dev.mysql.com/doc/refman/8.0/en/mysqldump.html>

- 전체: `mysqldump [-h host] -u user -p -A > dump.sql`
- Table: `mysqldump [-h host] -u user -p db table1 [table2 ...] > dump.sql`
- Database `USE DB`: `mysqldump [-h host] -u user -p --databases db1 [db2 ...] > dump.sql`

데이터베이스 인수가 하나일 때는 `--databases` 키워드를 생략할 수 있음.
키워드를 생략하게 되면 덤프파일에 `CREATE DATABASE` / `USE DATABASE`가 생략됨

특정 조건의 데이터만 dump 하려면 `--where`(`-w`) 키워드를 사용하면 된다: `-w 'condition'`

### Import

<https://dev.mysql.com/doc/refman/8.0/en/mysqlimport.html>

- 특정 database로 복구: `mysql [-h host] -u user -p db < dump.sql`

## sql_mode

작동방식 옵션을 설정한다.

조회: `mysql> SELECT @@sql_mode;`

### ONLY_FULL_GROUP_BY

`GROUP BY`를 사용하는 경우 `GROUP BY`의 조건이 되는 컬럼과 집계 함수(Aggregation Function)만 `SELECT` 할 수 있다.

```sql
SELECT job, COUNT(empno) as "인원수", AVG(sal) as "평균급여액"
FROM emp
GROUP BY job;
```

MySQL의 `GROUP BY`는 표준과 다르게 작동해서 조건이 아닌 컬럼도 `SELECT` 할 수 있다.
이를 방지하기 위해서 `sql_mode`에 `ONLY_FULL_GROUP_BY` 옵션을 추가할 수 있다. (버전이 올라가며 기본옵션일 수도 있다)

`mysql> SET sql_mode = 'ONLY_FULL_GROUP_BY';`

반대로 `ONLY_FULL_GROUP_BY`이 활성화 되어 있다면 해제할 수도 있다.

- `mysql> SELECT REPLACE(@@sql_mode,'ONLY_FULL_GROUP_BY','')`
- 직접 설정파일(`/etc/mysql/my.cnf`)의 `sql_mode`를 수정하여도 된다.

설정을 변경하였다면 `mysqld`를 재시작한다

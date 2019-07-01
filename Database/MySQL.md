# MySQL / MariaDB

<https://dev.mysql.com/doc/refman/5.7/en/sql-syntax-data-definition.html>

## 조회

### db 조회

```sql
show databases;
```

### user 조회

```sql
SELECT host, user from mysql.user;
```

## 생성

### db, user 생성

```sql
-- UTF8로 DB생성
CREATE DATABASE `dbname` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

-- 유저 생성
CREATE USER 'userid'@'%' IDENTIFIED BY 'password';
INSERT INTO mysql.user VALUES('%', 'userid', PASSWORD('password'));
```

### 권한 부여

```sql
-- 유저에게 특정 DB 권한 주기
GRANT {권한} on {db|*}.{table|*} to 'user'@'localhost' identified by 'password';
GRANT {권한} on {db|*}.{table|*} to 'user'@'%' identified by 'password';

-- 권한변경 반영
flush privileges;

-- 외부접속 권한 삭제
DELETE FROM mysql.user WHERE Host='%' AND User='user';
FLUSH PRIVILEGES;
```

#### 권한 종류

<https://dev.mysql.com/doc/refman/5.7/en/privileges-provided.html>

| Privilege               | Grant Table Column             | Context                               |
| ----------------------- | ------------------------------ | ------------------------------------- |
| ALL [PRIVILEGES]        | Synonym for “all privileges”   | Server administration                 |
| ALTER                   | Alter_priv                     | Tables                                |
| ALTER ROUTINE           | Alter_routine_priv             | Stored routines                       |
| CREATE                  | Create_priv                    | Databases, tables, or indexes         |
| CREATE ROUTINE          | Create_routine_priv            | Stored routines                       |
| CREATE TABLESPACE       | Create_tablespace_priv         | Server administration                 |
| CREATE TEMPORARY TABLES | Create_tmp_table_priv          | Tables                                |
| CREATE USER             | Create_user_priv               | Server administration                 |
| CREATE VIEW             | Create_view_priv               | Views                                 |
| DELETE                  | Delete_priv                    | Tables                                |
| DROP                    | Drop_priv                      | Databases, tables, or views           |
| EVENT                   | Event_priv                     | Databases                             |
| EXECUTE                 | Execute_priv                   | Stored routines                       |
| FILE                    | File_priv                      | File access on server host            |
| GRANT OPTION            | Grant_priv                     | Databases, tables, or stored routines |
| INDEX                   | Index_priv                     | Tables                                |
| INSERT                  | Insert_priv                    | Tables or columns                     |
| LOCK TABLES             | Lock_tables_priv               | Databases                             |
| PROCESS                 | Process_priv                   | Server administration                 |
| PROXY                   | See proxies_priv table         | Server administration                 |
| REFERENCES              | References_priv                | Databases or tables                   |
| RELOAD                  | Reload_priv                    | Server administration                 |
| REPLICATION CLIENT      | Repl_client_priv               | Server administration                 |
| REPLICATION SLAVE       | Repl_slave_priv                | Server administration                 |
| SELECT                  | Select_priv                    | Tables or columns                     |
| SHOW DATABASES          | Show_db_priv                   | Server administration                 |
| SHOW VIEW               | Show_view_priv                 | Views                                 |
| SHUTDOWN                | Shutdown_priv                  | Server administration                 |
| SUPER                   | Super_priv                     | Server administration                 |
| TRIGGER                 | Trigger_priv                   | Tables                                |
| UPDATE                  | Update_priv                    | Tables or columns                     |
| USAGE                   | Synonym for “no privileges”    | Server administration                 |

### 테이블 인코딩 설정 생성

```sql
CREATE TABLE test(
  title varchar(20)
) default charset utf8;
```

## 날짜 처리

<https://dev.mysql.com/doc/refman/5.7/en/date-and-time-functions.html#function_date-format>

## DB Dump / Import

### Dump

<https://dev.mysql.com/doc/refman/5.7/en/mysqldump.html>

- 전체: `mysqldump [-h host] -u user -p -A > dump.sql`
- Table: `mysqldump [-h host] -u user -p db table1 [table2 ...] > dump.sql`
- Database `USE DB`: `mysqldump [-h host] -u user -p --databases db1 [db2 ...] > dump.sql`

데이터베이스 인수가 하나일 때는 `--databases` 키워드를 생략할 수 있음.
키워드를 생략하게 되면 덤프파일에 `CREATE DATABASE` / `USE DATABASE`가 생략됨

특정 조건의 데이터만 dump 하려면 `--where`(`-w`) 키워드를 사용하면 된다: `-w 'condition'`

### Import

<https://dev.mysql.com/doc/refman/5.7/en/mysqlimport.html>

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

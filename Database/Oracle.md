# ORACLE DB

## 테이블 스페이스 생성

```sql
CREATE tablespace [tablespace_name]
datafile '/home/oracle/oradata/orcl/[file_name].dbf'
size 500m;

ALTER DATABASE DATAFILE '데이터파일경로' AUTOEXTEND ON NEXT 100M MAXSIZE UNLIMITED;
ALTER DATABASE DATAFILE '데이터파일경로' resize 500M;
```

## 유저 생성 및 테이블 QUOTA 부여

```sql
CREATE USER [user_name]
IDENTIFIED BY [password]
DEFAULT TABLESPACE [tablespace_name]
TEMPORARY TABLESPACE TEMP;

QUOTA 100M ON [tableA]
QUOTA 500M ON [tableB]
QUOTA UNLIMITED ON [table];
-- 참조
-- QUOTA UNLIMITED 옵션을 줄 경우 무제한 사용가능,
-- QUOTA 0M 의 경우 해당 테이블스페이스에는 테이블만 생성가능

-- Quota 정보 확인
SELECT * FROM dba_ts_quotas;
-- 무제한 사용권한을 부여
GRANT unlimited tablespace to test;
-- 특정 테이블스페이스에 대한 Quota 제거
ALTER USER test quota unlimited on test_tablespace;
```

## 유저 권한

```sql
GRANT connect, resource, dba TO [user_name]; -- 모든 권한 주기

GRANT CREATE SESSION TO 유저명         -- 데이터베이스에 접근할 수 있는 권한
GRANT CREATE DATABASE LINK TO 유저명
GRANT CREATE MATERIALIZED VIEW TO 유저명
GRANT CREATE PROCEDURE TO 유저명
GRANT CREATE PUBLIC SYNONYM TO 유저명
GRANT CREATE ROLE TO 유저명
GRANT CREATE SEQUENCE TO 유저명
GRANT CREATE SYNONYM TO 유저명
GRANT CREATE TABLE TO 유저명             -- 테이블을 생성할 수 있는 권한
GRANT DROP ANY TABLE TO 유저명         -- 테이블을 제거할 수 있는 권한
GRANT CREATE TRIGGER TO 유저명
GRANT CREATE TYPE TO 유저명
GRANT CREATE VIEW TO 유저명

GRANT CREATE SESSION
,CREATE TABLE
,CREATE SEQUENCE
,CREATE VIEW
TO 유저명;
```

## 제약조건

- 제약조건
  - NN : not null
  - PK : primary key
  - FK : foreign key
  - UK : unique (후보키)
  - CK : check

테이블 생성 후, 레코드 입력시 데이터의 중복이 발생할 수 있음

해결책) 데이터를 무조건 받아들이는 것이 아니라, 컬럼에 까다로운 제한을 걸어놓을 수 있다. 이렇게 컬럼에 지정할 수 있는 제한, 제약사항을 제약조건이라 한다.
데이터 중복을 허용하지 않겠다 라는 제약조건 unique

## 오라클 DB 개요

- 데이터형
  - 문자 : CHAR, VARCHAR2, CLOB
  - 숫자 : NUMBER
  - 날짜 : DATE, TIMESTAMP
  - 기타 : BFILE, BLOB

- DDL(Database Definition Lang)
  - CREATE / REPLACE (ALTER가 안되는 경우에 지원)
    - TABLE : 데이터형, 제약조건
    - SEQUENCE : START WITH, INCREMENT BY, CYCLE|NOCYCLE, CACHE|NOCACHE
    - VIEW : 가상테이블
      - 단순뷰
      - 복합뷰(JOIN, SUBQUERY)
      - 보안, 쿼리문장 단순화
    - INDEX : 검색속도 최적화
    - PROCEDURE : 반복수행, 여러개의 SQL문장을 동시처리
    - FUNCTION : 결과 값
    - PACKAGE : 관련성 있는 것을 모아놓음
    - TRIGGER
  - DROP
  - ALTER
  - TRUNCATE : 데이터 잘라내기(구조는 유지하고 자료만 삭제)
  - RENAME : table, column 이름 변경

- DML(Database Manipulation Lang)
  - SELECT
    - 출력 조건
      - WHERE
      - GROUP BY ~ HAVING
      - ORDER BY
    - JOIN : ANSI JOIN / ORACLE JOIN
      - INNER JOIN (null은 처리못함)
        - EQUI_JOIN : JOIN~USING, NATURAL JOIN : =
        - NON EQUI_JOIN : = 이외 (비교, 논리연산 등...)
      - OUTER JOIN
        - LEFT OUTER JOIN : `A.col = B.col(+)`
        - RIGHT OUTER JOIN : `(+)A.col = B.col`
        - FULL OUTER JOIN
    - SUBQUERY
      - 단일행
      - 다중행 : ANY/SOME, ALL, IN(특정 값들)
      - 다중칼럼 : WHERE (empno, job)
      - COLUMN(스칼라 서브쿼리), TABLE 조건 위치에 사용
    - 연산자
      - 산술연산
      - 논리연산
      - 비교연산
      - IN
      - BETWEEN ~ AND ( >= AND <= )
      - LIKE : _(한 글자) %(여러글자)
      - NOT
      - NULL : IS NULL / IS NOT NULL
      - EXIST
  - INSERT
    - INSERT
    - INSERT ALL
  - DELETE
  - UPDATE
  - UNION, UNION ALL(합집합), INTERSECT(교집합), MINUS(차집합)
  - MERGE

- DCL(Database Control Lang)
  - GRANT
  - REVOKE

- TCL
  - COMMIT
  - ROLLBACK

## 오라클 내장함수

- 숫자관련
  - ROUND() : 반올림
  - TRUNC() : 버림
  - CEIL() : 올림
  - MOD() : 나머지
- 문자관련
  - SUBSTR() : 문자 분해
  - INSTR() : 문자 위치 확인
  - RPAD()
- 변환함수
  - TO_CHAR() : 문자로 변환(시간출력에 주로 사용)
  - TO_DATE()
- 날짜함수
  - SYSDATE
  - MONTHS_BETWEEN
  - ADD_MONTH
- 기타함수
  - DECODE, CASE -> 조건문 : Trigger, Procedure에 쓰임
  - NVL : Null 값이면 대체
  - CASE
- 집합함수
  - COUNT()
  - MAX()
  - SUM()
  - AVG()
  - RANK()
  - ROLLUP()
  - CUBE()

## join

정규화에 의해 분리된 테이블을 마치 하나처럼 보여주는 쿼리

- inner join :
  join대상이 되는 테이블에 공통된 레코드만 가져온다

- outer join :
  join쿼리 수행시 반드시 모두 가져올 레코드를 보유한 테이블을 명시할 수 있어서, 공통되지 않은 레코드도 조회가 가능하다

## 테이블 조회

```sql
select table_name from user_tables;
```

## Oracle to Java

- CLOB
  - 8i, 9i : InputStream
  - 10g, 11g, 12C : String
- NUMBER
  - NUMBER : int, double, long
  - NUMBER(4) : int
  - NUMBER(7,2) : double
- DATE, TIMESTAMP : Date(java.util), TimeStamp
- BFILE, BLOB : ~ 4Gb : InputStream

## PL/SQL

- PL : 프로시저 제작 언어 : SQL 사용하면 PL/SQL
  - PROCEDURE
  - FUNCTION
  - PACKAGE
  - TRIGGER
  - INDEX
- 쿼리 문장이 복잡 (JOIN, SUBQUERY) : VIEW
- 검색속도 최적화 : Index(rowid)
- PROCEDURE : 절차적 중복이 많은 경우

- 선언부 : 변수선언
  - DELCARE
  - 스칼라 변수 :  no NUMBER, name VARCHAR2(10) ...
  - 참조 변수
    - ename -> emp.ename%TYPE(0)
    - emp -> emp%ROWTYPE -> VO
    - RECORD : 사용자 지정(필요한 데이터) : RECORD(ename,dname...)
    - CURSOR : ResultSet
- 구현부 : SQL 문장
  - BEGIN
    - `SELECT ename,job INTO vename,vjob FROM emp`
    - DBMS_OUTPUT.PUT_LINE(vename) : System.out.print
    - 제어문, 연산자 : `IF, IF~ELSE, IF~ELSIF, WHILE, FOR, LOOP`
  - END;

- 예외처리부 : EXCEPTION

- SQLPLUS에서 입력받기 (scan)
  ```sql
  SET serveroutput ON
  ACCEPT pempno PROMPT '사번:'
  -- 변수에 대입 (:=&pempno ), 사용 (&pempno)
  ```

# MySQL DML

<https://dev.mysql.com/doc/refman/8.0/en/sql-data-manipulation-statements.html>

## SELECT

### 집계함수

<https://dev.mysql.com/doc/refman/8.0/en/group-by-functions.html>

숫자 인수의 경우 분산 및 표준편차 함수의 경우 DOUBLE 타입 값을 반환한다.

SUM(), AVG() 함수는 정확한 값의 인수(정수 또는 DECIMAL)에 대해서는 DECIMAL 타입 값을,
근사값 인수(FLOAT, DOUBLE)에 대해서는 DOUBLE 타입 값을 리턴한다.

### query cache

```sql
SET SESSION query_cache_type = OFF;
SHOW VARIABLES LIKE 'query_cache_type';

-- select 실행시 캐시 없이 할 수도 있음
SELECT SQL_NO_CACHE * FROM <TABLE_NAME>;
```

## INSERT SELECT

```sql
INSERT INTO tbl_temp2 (fld_id)
  SELECT tbl_temp1.fld_order_id
  FROM tbl_temp1
  WHERE tbl_temp1.fld_order_id > 100;
```

## UPDATE FROM SELECT

```sql
UPDATE
  tablename1 AS t1
  JOIN tablename2 AS t2 ON join_condition
SET assignment_list
[WHERE where_condition];
```

## DELETE FROM multiple tables

> <https://dev.mysql.com/doc/refman/8.0/en/delete.html> > Multiple-Table Syntax

여러 테이블에 걸친 데이터를 한 번에 삭제할 수 있다

```sql
delete t, u
from t
inner join u
  on t.id = u.t_id
where
  t.id < 5000;

-- u 테이블 값만 삭제하는 경우
delete u
from t
inner join u
  on t.id = u.t_id
where
  t.id < 5000;
```

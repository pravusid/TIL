# SQL 예제

## 페이징 (Paging)

### 오라클 (Oracle)

```sql
SELECT *
FROM (
  SELECT x.*, ROWNUM rnum
  FROM (
    SELECT m.id AS id, m.username AS username
    FROM member m
    ORDER BY m.id
  ) x
  WHERE ROWNUM <= ?
)
WHERE rnum > ?
```

### MySQL

```sql
SELECT id, username
FROM member
ORDER BY id
LIMIT ?, ? -- 시작위치, 가져올 row 숫자
```

### PostgreSQL

```sql
SELECT id, username
FROM member
ORDER BY id
LIMIT ? OFFSET ? -- 가져올 row 숫자, 시작위치
```

## 검색

## 데이터 입력

### 예제 자료 입력시 FK 제약조건 검사피하기

```sql
set foreign_key_checks = 0;

-- 입력할 내용 ...

set foreign_key_checks = 1;
```

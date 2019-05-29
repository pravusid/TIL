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
LIMIT :page * :size, :size -- 시작위치, 가져올 row 숫자

SELECT id, username
FROM member
ORDER BY id
LIMIT :size OFFSET :page * :size

SELECT id, username, createdAt, @ROWNUM := @ROWNUM + 1 AS rnum
FROM member, (SELECT @ROWNUM := 0) x
ORDER BY createdAt
LIMIT :page * :size, :size
```

### PostgreSQL

```sql
SELECT id, username
FROM member
ORDER BY id
LIMIT ? OFFSET ? -- 가져올 row 숫자, 시작위치
```

### Paging with OFFSET

일반적으로 데이터베이스의 인덱스는 B+Tree 자료구조를 사용한다.
효율적인 위치 탐색이 가능하고 Leaf node끼리 연결되어 있으므로 범위 검색도 수행할 수 있다.

그러나 OFFSET은 순차적으로 탐색(OFFSET n 이면 n까지 탐색한 다음 앞의 n개를 버림)하므로 실행시간이 선형으로 증가한다.

따라서 페이징을 한다면 순서가 있는 key를 기반으로(인덱스가 적용된) 범위를 지정해야 한다.
이 경우 leaf node 위치 탐색(`O(log n)`)을 두 번 수행하는 것이므로 일정한 시간을 보장할 수 있다.

## 검색

### `IN` Clause with NULL

`IN`문은 `col = val1 or col = val2 or col = val3`의 형태로 해석되기 때문에 `val = NULL`의 형태가 되는 `NULL`을 사용할 수 없다.

`NOT IN`의 경우도 동일하다.

### LIKE

LIKE 연산자를 사용하여 검색을 하는 경우 인덱스를 사용하기 위해서는 전방 일치검색을 해야한다. (`LIKE '단어%'`)

## 데이터 입력

### 예제 자료 입력시 FK 제약조건 검사피하기

```sql
set foreign_key_checks = 0;

-- 입력할 내용 ...

set foreign_key_checks = 1;
```

## Update Join

### My sql

```sql
UPDATE
  emp e INNER JOIN dept d
  ON d.deptno = e.deptno
SET e.dname = d.dname;
```

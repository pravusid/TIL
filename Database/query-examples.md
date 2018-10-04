# SQL 예제

## 데이터 출력

### 페이징 (Paging)

#### 오라클 (Oracle)

```sql
SELECT *
FROM (
  SELECT X.*, ROWNUM rnum
  FROM (
    SELECT M.ID AS ID, M.AGE AS AGE, M.NAME AS NAME
    FROM MEMBER M
    ORDER BY M.NAME
    ) X
  WHERE ROWNUM <= ?
  )
WHERE rnum > ?
```

## 검색

## 데이터 입력
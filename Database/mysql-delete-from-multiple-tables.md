# MySQL: delete from multiple tables

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

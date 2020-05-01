# InnoDB

## InnoDB Lock

### Lock 개요

<https://dev.mysql.com/doc/refman/8.0/en/innodb-locking.html>

- 공유 (S) 락은 트랜잭션에게 레코드 읽기를 허용한다
- 배타적 (X) 락은 트랜잭션에게 레코드 업데이트 또는 삭제를 허용한다

인텐션은 테이블에 있는 행에 대해서 나중에 요청되는 것이 어떤 형태의 락(S/X)인지를 나타내는 테이블 수준의 잠금이다.
InnoDB에서는 두 가지 타입의 인텐션 락이 사용된다 (트랜잭션 T는 테이블 R에서 가리키는 타입에 대해서 요청되는 락을 가지고 있다)

- 인텐션 공유 (IS): 트랜잭션 T는 테이블 R에 있는 개별적인 열에 S 락을 설정하려고 한다.
- 인텐션 배타적 (IX): 트랜잭션 T는 이러한 열에 대해서 X 락을 설정하려고 한다.

인텐션락 프로토콜은 다음과 같다

- 트랜잭션에서 레코드에 대한 S 락을 얻기위해, 테이블에 대해 IS 락 혹은 더 강한 락 권한을 얻어야 한다
- 트랜잭션에서 레코드에 대한 X 락을 얻기위해, 테이블에 대해 IX 락을 얻어야 한다

> `IS`, `IX` 락은 여러 트랜잭션에서 동시에 접근이 가능하다 (접근 순서를 획득한다)

|     | X        | IX         | S          | IS         |
| --- | -------- | ---------- | ---------- | ---------- |
| X   | Conflict | Conflict   | Conflict   | Conflict   |
| IX  | Conflict | Compatible | Conflict   | Compatible |
| S   | Conflict | Conflict   | Compatible | Compatible |
| IS  | Conflict | Compatible | Compatible | Compatible |

### Lock 유형

- Record Locks
- Gap Locks
- Next-Key Locks
- Insert Intention Locks
- AUTO-INC Locks
- Predicate Locks for Spatial Indexes

### 상황별 Lock 설정

<https://dev.mysql.com/doc/refman/8.0/en/innodb-locks-set.html>

## 존재하지 않는행 `SELECT... FOR UPDATE`

TL;DR

InnoDB `FOR UPDATE` 잠금은 row 단위 인데(unique 조건) 만약 행이 존재하지 않는다면
해당 테이블에 [next-key-locks](https://dev.mysql.com/doc/refman/8.0/en/innodb-locking.html#innodb-next-key-locks) 적용 될 수 있음

또한 여러 트랜잭션이 존재하지 않는 행에 대해 `SELECT... FOR UPDATE` 요청을 동시에 하면, 명확한 lock 대상이 없으므로 deadlock 발생가능

- <https://mysqlquicksand.wordpress.com/2019/12/20/select-for-update-on-non-existent-rows/>
- <https://fastmail.blog/2017/12/09/mysql-lock-nonexistent-row/>
- <https://stackoverflow.com/questions/17068686/how-do-i-lock-on-an-innodb-row-that-doesnt-exist-yet>

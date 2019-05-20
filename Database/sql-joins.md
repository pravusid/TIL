# SQL Joins

![SQL Joins](https://raw.githubusercontent.com/pravusid/TIL/master/Database/img/sql-joins.jpg)

## Join

관계형 데이터베이스는 데이터를 효율적으로 저장하기 위해 정규화 과정을 거친 데이터를 보관하는데,
필요에 따라 관계있는 데이터를 함께 연결하여 조회하기 위해서 Join을 사용한다.

## INNER JOIN

### EQUI JOIN

INNER JOIN의 가장 일반적인 활용방식

여러 테이블에 존재하는 공통 컬럼의 공통 값을 조건으로 결과를 출력함

Explicit
> SELECT * FROM emp INNER JOIN dept ON emp.deptno = dept.deptno

Implicit
> SELECT * FROM emp, dept WHERE emp.deptno = dept.deptno

#### NATURAL JOIN

EQUI JOIN과 거의 유사하지만, 대상 테이블들의 모든 컬럼을 비교하여, 같은 컬럼명끼리 조인을 수행한다.

같은 이름을 가진 컬럼은 한 번만 출력한다.

> SELECT * FROM emp NATURAL JOIN dept

## (OUTER) JOIN

INNER JOIN은 공통 컬럼의 공통 값을 기반으로 결과를 출력한다.

하지만 JOIN 대상중 특정 테이블의 데이터가 모두 필요한(공통 값이 없더라도) 경우가 있다.
이때, OUTER JOIN을 사용한다.

### LEFT (OUTER) JOIN

LEFT OUTER JOIN은 좌측에 위치한 테이블 데이터는 모두 출력한다.

> SELECT * FROM emp LEFT OUTER JOIN dept ON emp.deptno = dept.deptno;

### RIGHT (OUTER) JOIN

LEFT OUTER JOIN은 우측에 위치한 테이블 데이터는 모두 출력한다.

> SELECT * FROM emp RIGHT OUTER JOIN dept ON emp.deptno = dept.deptno;

### FULL (OUTER) JOIN

FULL OUTER JOIN은 대상 테이블 데이터가 모두 필요한 경우 사용한다

MySQL, MariaDB에서는 FULL OUTER JOIN을 지원하지 않으므로 대신 UNION을 사용한다

> SELECT * FROM emp FULL OUTER JOIN dept ON emp.deptno = dept.deptno;

## SELF JOIN

하나의 테이블을 여러번 활용하여 JOIN 결과물을 출력함

Explicit
> SELECT e.empno, e.ename, e.job, e.hiredate, e.sal, m.empno, m.ename, m.job FROM emp E INNER JOIN emp M ON E.mgr = M.empno

Implicit
> SELECT e.empno, e.ename, e.job, e.hiredate, e.sal, m.empno, m.ename, m.job FROM emp E, emp M WHERE E.mgr = M.empno

## ANTI JOIN

테이블에서 JOIN 대상 테이블과 일치하지 않는 데이터를 출력한다 (NOT IN / NOT EXIST)

```sql
SELECT *
FROM Employee
WHERE DeptName NOT IN (
  SELECT DeptName FROM Dept
);

SELECT *
FROM Employee
WHERE NOT EXISTS (
  SELECT * FROM Dept
  WHERE Employee.DeptName = Dept.DeptName
)
```

## SEMI JOIN

서브쿼리를 사용했을 때 메인쿼리와의 연결을 처리한다 (IN / EXIST)

```sql
SELECT *
FROM Employee
WHERE DeptName IN (
  SELECT DeptName FROM Dept
);

SELECT *
FROM Employee
WHERE EXISTS (
  SELECT * FROM Dept
  WHERE Employee.DeptName = Dept.DeptName
);
```

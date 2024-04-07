# AWS RDS

관계형 데이터베이스 서비스

## 설정 (for MySQL)

- <https://aws.amazon.com/blogs/database/best-practices-for-configuring-parameters-for-amazon-rds-for-mysql-part-1-parameters-related-to-performance/>
- <https://aws.amazon.com/blogs/database/best-practices-for-configuring-parameters-for-amazon-rds-for-mysql-part-2-parameters-related-to-replication/>
- <https://aws.amazon.com/blogs/database/best-practices-for-configuring-parameters-for-amazon-rds-for-mysql-part-3-parameters-related-to-security-operational-manageability-and-connectivity-timeout/>

## TLS 사용한 연결

> AWS 인증서: <https://docs.aws.amazon.com/ko_kr/AmazonRDS/latest/UserGuide/UsingWithRDS.SSL.html>

## dump

To export db from RDS

`mysqldump -h rds.host.name -u remote_user_name -p remote_db > remote_db.sql`

To import db on RDS

`mysql -h rds.host.name -u remote_user_name -p remote_db < remote_db.sql`

## Troubleshooting

### 문자 인코딩 관련

- 최초 인코딩셋이 latin1이므로 RDS 대시보드 파라미터 그룹으로 들어감

- 새로운 파라미터 그룹을 생성하고 character set 관련 항목을 utf8로 변경
- 인스턴스에 새 파라미터 그룹 적용후 리부팅
- 확인: `show variables like 'c%';`

- 기존에 생성된 database의 문자열셋을 변경함

  ```sql
  ALTER DATABASE <database-name>
  CHARACTER SET = 'utf8'
  COLLATE = 'utf8_general_ci';
  ```

## 최대 연결 수 (기본설정)

<https://docs.aws.amazon.com/ko_kr/AmazonRDS/latest/UserGuide/CHAP_Limits.html#RDS_Limits.MaxConnections>

### RDS MySQL

공식에 따라 인스턴스 유형별 최대 연결 수 기본 값을 계산하면 다음과 같다

> `{DBInstanceClassMemory/12582880}`
> 값이 16,000보다 큰 경우 Amazon RDS는 MariaDB 및 MySQL DB 인스턴스에 대한 제한을 16,000으로 설정

- t2.micro: 66
- t2.small: 150
- m3.medium: 296
- t2.medium: 312
- M3.large: 609
- t2.large: 648
- M4.large: 648
- M3.xlarge: 1237
- R3.large: 1258
- M4.xlarge: 1320
- M2.xlarge: 1412
- M3.2xlarge: 2492
- R3.xlarge: 2540

> <https://support.bespinglobal.com/ko/support/solutions/articles/73000524758>

### RDS Aurora

- Aurora MySQL: <https://docs.aws.amazon.com/ko_kr/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Managing.Performance.html#AuroraMySQL.Managing.MaxConnections>
- Aurora PostgreSQL: <https://docs.aws.amazon.com/ko_kr/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Managing.html#AuroraPostgreSQL.Managing.MaxConnections>

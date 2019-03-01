# AWS RDS

관계형 데이터베이스 서비스

새 보안 그룹을 생성하면서 EC2와 통신을 위해 EC2인스턴스에 적용한 보안 그룹 ID를 가져와 인바운드 소스로 적용한다

문자 인코딩 관련

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

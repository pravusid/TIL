# Amazon Web Services

Cloud Computing Services

Free Tier는 최초 12개월간 허가된 인스턴스 사용시간 총량을 월별 750시간 제공함
사용량은 시간단위로 계산됨 (1분을 켜도 1시간)

Free Tier 관련 내용은 <https://aws.amazon.com/ko/free/> 참고

Free Tier 주의사항 관련 모음: <https://librewiki.net/wiki/%EC%95%84%EB%A7%88%EC%A1%B4_%EC%9B%B9_%EC%84%9C%EB%B9%84%EC%8A%A4/Free_Tier_%EC%A3%BC%EC%9D%98%EC%82%AC%ED%95%AD>

## Elastic Compute Cloud (EC2)

인스턴스 생성시 유의점: 보안그룹 확인

- ssh 인바운드는 지정된 IP에서만
- 웹서비스용 80포트 443포트는 열어야함
- 인스턴스에 접근하기 위한 pem키를 잃어버리면 인스턴스를 다시 생성해야 함

ssh에서 EC2 인스턴스 접속을 위한 호스트설정을 한다: `~/.ssh/config`

```text
Host ec2-idpravus
    HostName <EIP||퍼블릭DNS>
    User <ec2-user||ubuntu>
    IdentityFile ~/.ssh/<pem파일>
```

기본 타임존은 UTC이므로 UTC+9로 변경한다

```shell
sudo rm /etc/localtime
sudo ln -s /usr/share/zoneinfo/Asia/Seoul /etc/localtime
```

## RDS

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

## Simple Storage Service (S3)

스토리지 서비스, Bucket을 생성하여 key-object로 데이터를 저장할 수 있다

## DynamoDB

종합 관리형 NoSQL 데이터베이스 서비스

테이블 / 항목 / 속성으로 구성됨

### 기본키

테이블에는 기본키가 필요하고(고유식별자) 두 가지를 지원함

- 파티션 키(단순 기본키)
- 파티션 키 및 정렬 키(복합 기본키)

### 보조 인덱스

테이블에서 종류별 최대 5개의 보조 인덱스를 생성할 수 있다

- Global secondary index: 파티션 키 및 정렬 키가 테이블의 파티션 키 및 정렬 키와 다를 수 있는 인덱스
- Local secondary index: 테이블과 파티션 키는 동일하지만 정렬 키는 다른 인덱스

### DynamoDB 스트림

DynamoDB 테이블의 데이터 수정 이벤트가 발생한 순서대로 거의 실시간으로 스트림에 표시됨

AWS Lambda와 연결하여 거의 실시간으로 관심있는 정보를 처리할 수 있다

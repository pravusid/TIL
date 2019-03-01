# Amazon Web Services

Cloud Computing Services

Free Tier는 최초 12개월간 허가된 인스턴스 사용시간 총량을 월별 750시간 제공함

사용량은 시간단위로 계산됨 (1분을 켜도 1시간)

Free Tier 관련 내용은 <https://aws.amazon.com/ko/free/> 참고

## Free Tier 주의사항

### EC2

- 데이터 스토리지 무료 제공량 30GB
- ElasticIP(고정IP)를 생성하고 하고 인스턴스에 할당하지 않으면(혹은 인스턴스가 정지중이면) 과금
- ElasticIP를 통해 전송되는 데이터는 소액의 전송료가 발생
- General Purpose(SSD)를 사용할 것. Provisioned SSD는 사용에 주의

### RDS

- 프리티어에서는 다중 AZ배포를 사용하지 않을 것
- General Purpose(SSD)를 사용할 것. Provisioned SSD는 사용에 주의
- 자동 백업의 유지 기간은 0(disable)으로 설정해야 Snapshot 보관료를 내지 않음

### VPC

VPN이나 프라이빗 게이트웨이를 사용시 유료

### CloudWatch

프리 티어에서는 메트릭(기본 모니터링) 10개, 경보 10개 및 API 요청 1,000,000개 사용가능

## CloudFront

- 50GB 사용량
- 다른 리전으로 전송하는 데이터까지 포함한 것임

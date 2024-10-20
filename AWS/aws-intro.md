# Amazon Web Services (AWS)

Cloud Computing Services

## CLI

<https://docs.aws.amazon.com/ko_kr/cli/latest/userguide/cli-chap-welcome.html>

## AWS Console-to-Code

<https://aws.amazon.com/ko/blogs/korea/convert-aws-console-actions-to-reusable-code-with-aws-console-to-code-now-generally-available/>

## AWS SDK for JavaScript

<https://docs.aws.amazon.com/AWSJavaScriptSDK/v3/latest/>

> Default credential provider is how SDK resolve the AWS credential if you DO NOT provide one explicitly.

`v2`: CredentialProviderChain in Node.js resolves credential from sources as following order:

- [environmental variable](https://docs.aws.amazon.com/sdkref/latest/guide/environment-variables.html)
- shared credentials file
- ECS container credentials
- spawning external process
- OIDC token from specified file
- EC2 instance metadata

If one of the credential providers above fails to resolve the AWS credential, the chain falls back to next provider until a valid credential is resolved, or throw error when all of them fail.
In Browsers and ReactNative, the chain is empty, meaning you always need supply credentials explicitly.

`v3`: defaultProvider The credential sources and fallback order does not change in v3. It also supports AWS Single Sign-On credentials.

### aws-sdk-js-v3: `Error: Region is missing`

[client-cloudfront Region should not be required](https://github.com/aws/aws-sdk-js-v3/issues/3035)

## FreeTier

- FreeTier는 최초 12개월간 허가된 인스턴스 사용시간 총량을 월별 750시간 제공함
- 사용량은 시간단위로 계산됨 (1분을 켜도 1시간)
- 자세한 내용은 <https://aws.amazon.com/ko/free/> 참고

### EC2 (FreeTier)

- 데이터 스토리지 무료 제공량 30GB
- ElasticIP(고정IP)를 생성하고 하고 인스턴스에 할당하지 않으면(혹은 인스턴스가 정지중이면) 과금
- ElasticIP를 통해 전송되는 데이터는 소액의 전송료가 발생
- General Purpose(SSD)를 사용할 것. Provisioned SSD는 사용에 주의

### RDS (FreeTier)

- 프리티어에서는 다중 AZ배포를 사용하지 않을 것
- General Purpose(SSD)를 사용할 것. Provisioned SSD는 사용에 주의
- 자동 백업의 유지 기간은 0(disable)으로 설정해야 Snapshot 보관료를 내지 않음

### VPC (FreeTier)

VPN이나 프라이빗 게이트웨이를 사용시 유료

### CloudWatch (FreeTier)

프리 티어에서는 메트릭(기본 모니터링) 10개, 경보 10개 및 API 요청 1,000,000개 사용가능

### CloudFront (FreeTier)

- 50GB 사용량
- 다른 리전으로 전송하는 데이터까지 포함한 것임

## Services

### Systems Manager Parameter Store

<https://docs.aws.amazon.com/ko_kr/systems-manager/latest/userguide/systems-manager-parameter-store.html>

환경변수 같은 데이터를 AWS에서 관리하기 위해 사용하는 서비스

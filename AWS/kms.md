# AWS KMS

AWS Key Management System

<https://docs.aws.amazon.com/ko_kr/kms/latest/developerguide/overview.html>

암호화키 관리 서비스

- 사용자가 접근할 수 없는 관리형 키로 AWS의 다양한 서비스 (RDS, DynamoDB ...) 암호화를 자동실행할 수 있음
- 사용자가 관리하는 키(CMK)로 직접 암/복호화를 수행할 수 있음

## KMS 키 계층구조

- CMK (Client Master Key)

  - HSM(Hardware Security Module): 평문 CMK생성 -> 도메인키로 암호화
  - KMS Host(스토리지): 암호화 CMK 저장

- DataKey
  
  - HSM: 평문 데이터키 생성 -> 평문 CMK로 암호화
  - 클라이언트: 평문 데이터키 + 암호화 데이터키 수신

- 직접 KMS 암복호화

  - HSM에서 암복호화 수행
  - 원본 데이터는 4KB로 제한됨

## 대칭 DataKey 암복호화

- `GenerateDataKey`: 평문 및 암호화 DataKey 생성

  - 평문키: 데이터 암호화에 사용 후 폐기
  - 암호화키: 복호화를 위해서 암호화 된 데이터와 함께 보관

- `Decrypt`: 암호화 데이터키 복호화 -> 평문 데이터키

## 비대칭 DataKey 암복호화

- `GenerateDateKeyPair`: 암호화 비밀키, 비밀키, 공개키 생성
- `Decrypt`: 암호화 비밀키 복호화 -> 평문 비밀키

## 직접 KMS 암복호화

- `Encrypt`: 평문 암호화 요청
- `Decrypt`: 암호문 복호화 요청

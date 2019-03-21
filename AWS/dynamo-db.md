# Dynamo DB

종합 관리형 NoSQL 데이터베이스 서비스

테이블 / 항목 / 속성으로 구성됨

## 구조

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

## 쿼리

<https://docs.aws.amazon.com/ko_kr/amazondynamodb/latest/developerguide/WorkingWithItems.html>

### 비교연산자 / 조건식

<https://docs.aws.amazon.com/ko_kr/amazondynamodb/latest/developerguide/Expressions.OperatorsAndFunctions.html>

## Index

### GSI

<https://docs.aws.amazon.com/ko_kr/amazondynamodb/latest/developerguide/GSI.html>

### LSI

<https://docs.aws.amazon.com/ko_kr/amazondynamodb/latest/developerguide/LSI.html>

## 에러처리

<https://docs.aws.amazon.com/ko_kr/amazondynamodb/latest/developerguide/Programming.Errors.html>

JS error object 예시

```json
{
  "code": "ConditionalCheckFailedException",
  "time": "2019-03-21T14:40:32.171Z",
  "requestId": "94033QGK80J454DT5CR78O862NVV4KQNSO5AEMVJF66Q9ASUAAJG",
  "statusCode": 400,
  "retryable": false,
  "retryDelay": 13.794078870110315
}
```
